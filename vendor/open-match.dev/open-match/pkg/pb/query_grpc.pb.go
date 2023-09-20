// Copyright 2019 Google LLC
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: api/query.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QueryService_QueryTickets_FullMethodName   = "/openmatch.QueryService/QueryTickets"
	QueryService_QueryTicketIds_FullMethodName = "/openmatch.QueryService/QueryTicketIds"
	QueryService_QueryBackfills_FullMethodName = "/openmatch.QueryService/QueryBackfills"
)

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	//
	// QueryTickets pages the Tickets by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error)
	// QueryTicketIds gets the list of TicketIDs that meet all the filtering criteria requested by the pool.
	//   - If the Pool contains no Filters, QueryTicketIds will return all TicketIDs in the state storage.
	//
	// QueryTicketIds pages the TicketIDs by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTicketIds(ctx context.Context, in *QueryTicketIdsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketIdsClient, error)
	// QueryBackfills gets a list of Backfills.
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	QueryBackfills(ctx context.Context, in *QueryBackfillsRequest, opts ...grpc.CallOption) (QueryService_QueryBackfillsClient, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryTickets(ctx context.Context, in *QueryTicketsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryService_ServiceDesc.Streams[0], QueryService_QueryTickets_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryTicketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryTicketsClient interface {
	Recv() (*QueryTicketsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryTicketsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryTicketsClient) Recv() (*QueryTicketsResponse, error) {
	m := new(QueryTicketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) QueryTicketIds(ctx context.Context, in *QueryTicketIdsRequest, opts ...grpc.CallOption) (QueryService_QueryTicketIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryService_ServiceDesc.Streams[1], QueryService_QueryTicketIds_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryTicketIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryTicketIdsClient interface {
	Recv() (*QueryTicketIdsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryTicketIdsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryTicketIdsClient) Recv() (*QueryTicketIdsResponse, error) {
	m := new(QueryTicketIdsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) QueryBackfills(ctx context.Context, in *QueryBackfillsRequest, opts ...grpc.CallOption) (QueryService_QueryBackfillsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryService_ServiceDesc.Streams[2], QueryService_QueryBackfills_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceQueryBackfillsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_QueryBackfillsClient interface {
	Recv() (*QueryBackfillsResponse, error)
	grpc.ClientStream
}

type queryServiceQueryBackfillsClient struct {
	grpc.ClientStream
}

func (x *queryServiceQueryBackfillsClient) Recv() (*QueryBackfillsResponse, error) {
	m := new(QueryBackfillsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations should embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	// QueryTickets gets a list of Tickets that match all Filters of the input Pool.
	//   - If the Pool contains no Filters, QueryTickets will return all Tickets in the state storage.
	//
	// QueryTickets pages the Tickets by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTickets(*QueryTicketsRequest, QueryService_QueryTicketsServer) error
	// QueryTicketIds gets the list of TicketIDs that meet all the filtering criteria requested by the pool.
	//   - If the Pool contains no Filters, QueryTicketIds will return all TicketIDs in the state storage.
	//
	// QueryTicketIds pages the TicketIDs by `queryPageSize` and stream back responses.
	//   - queryPageSize is default to 1000 if not set, and has a minimum of 10 and maximum of 10000.
	QueryTicketIds(*QueryTicketIdsRequest, QueryService_QueryTicketIdsServer) error
	// QueryBackfills gets a list of Backfills.
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	QueryBackfills(*QueryBackfillsRequest, QueryService_QueryBackfillsServer) error
}

// UnimplementedQueryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) QueryTickets(*QueryTicketsRequest, QueryService_QueryTicketsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTickets not implemented")
}
func (UnimplementedQueryServiceServer) QueryTicketIds(*QueryTicketIdsRequest, QueryService_QueryTicketIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTicketIds not implemented")
}
func (UnimplementedQueryServiceServer) QueryBackfills(*QueryBackfillsRequest, QueryService_QueryBackfillsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryBackfills not implemented")
}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_QueryTickets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryTicketsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryTickets(m, &queryServiceQueryTicketsServer{stream})
}

type QueryService_QueryTicketsServer interface {
	Send(*QueryTicketsResponse) error
	grpc.ServerStream
}

type queryServiceQueryTicketsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryTicketsServer) Send(m *QueryTicketsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_QueryTicketIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryTicketIdsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryTicketIds(m, &queryServiceQueryTicketIdsServer{stream})
}

type QueryService_QueryTicketIdsServer interface {
	Send(*QueryTicketIdsResponse) error
	grpc.ServerStream
}

type queryServiceQueryTicketIdsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryTicketIdsServer) Send(m *QueryTicketIdsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_QueryBackfills_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryBackfillsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).QueryBackfills(m, &queryServiceQueryBackfillsServer{stream})
}

type QueryService_QueryBackfillsServer interface {
	Send(*QueryBackfillsResponse) error
	grpc.ServerStream
}

type queryServiceQueryBackfillsServer struct {
	grpc.ServerStream
}

func (x *queryServiceQueryBackfillsServer) Send(m *QueryBackfillsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryTickets",
			Handler:       _QueryService_QueryTickets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryTicketIds",
			Handler:       _QueryService_QueryTicketIds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryBackfills",
			Handler:       _QueryService_QueryBackfills_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/query.proto",
}
