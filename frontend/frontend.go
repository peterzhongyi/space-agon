// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb2 "github.com/googleforgames/open-match2/v2/pkg/pb"
	"github.com/googleforgames/space-agon/game/protostream"
	"github.com/googleforgames/space-agon/omclient"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Println("Serving index page.")
			http.ServeFile(w, r, "/app/static/index.html")
		} else {
			log.Println("404 on", r.URL.Path)
			http.NotFound(w, r)
		}
	})

	http.Handle("/matchmake/", websocket.Handler(matchmake))

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func matchmake(ws *websocket.Conn) {
	ws.PayloadType = 2 // Sets sent payloads to binary
	stream := protostream.NewProtoStream(ws)

	ctx, cancel := context.WithCancel(ws.Request().Context())
	defer cancel()
	assignments := make(chan *pb2.Assignment)
	errs := make(chan error)

	go streamAssignments(ctx, assignments, errs)

	for {
		select {
		case err := <-errs:
			log.Println("Error getting assignment:", err)
			err = stream.Send(&pb2.Assignment{})
			if err != nil {
				log.Println("Error sending error:", err)
			}
			return
		case assigment := <-assignments:
			err := stream.Send(assigment)
			if err != nil {
				log.Println("Error sending updated assignment:", err)
				cancel()
				return
			}
		}
	}
}

func streamAssignments(ctx context.Context, assignments chan *pb2.Assignment, errs chan error) {
	omClient := omclient.CreateOMClient()

	ticketId, err := omClient.CreateTicket(&pb2.Ticket{})
	if err != nil {
		errs <- fmt.Errorf("error creating open match ticket: %w", err)
		return
	}

	ticketIdsToActivate := make(chan string)

	go func() {
		omClient.ActivateTickets(ctx, ticketIdsToActivate)
	}()

	ticketIdsToActivate <- ticketId

	waReq := &pb2.WatchAssignmentsRequest{
		TicketIds: []string{ticketId},
	}

	assignmentsResultChan := make(chan *pb2.StreamedWatchAssignmentsResponse)

	go omClient.WatchAssignments(context.Background(), waReq, assignmentsResultChan)

	resp := <-assignmentsResultChan
	log.Printf("Got assignment: %v", resp.Assignment)
	assignments <- resp.Assignment
}
