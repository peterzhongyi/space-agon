// Copyright 2024 Google LLC All Rights Reserved.
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

// This code was autogenerated. Do not edit directly.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FleetAutoscalerStatusApplyConfiguration represents an declarative configuration of the FleetAutoscalerStatus type for use
// with apply.
type FleetAutoscalerStatusApplyConfiguration struct {
	CurrentReplicas *int32   `json:"currentReplicas,omitempty"`
	DesiredReplicas *int32   `json:"desiredReplicas,omitempty"`
	LastScaleTime   *v1.Time `json:"lastScaleTime,omitempty"`
	AbleToScale     *bool    `json:"ableToScale,omitempty"`
	ScalingLimited  *bool    `json:"scalingLimited,omitempty"`
}

// FleetAutoscalerStatusApplyConfiguration constructs an declarative configuration of the FleetAutoscalerStatus type for use with
// apply.
func FleetAutoscalerStatus() *FleetAutoscalerStatusApplyConfiguration {
	return &FleetAutoscalerStatusApplyConfiguration{}
}

// WithCurrentReplicas sets the CurrentReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentReplicas field is set to the value of the last call.
func (b *FleetAutoscalerStatusApplyConfiguration) WithCurrentReplicas(value int32) *FleetAutoscalerStatusApplyConfiguration {
	b.CurrentReplicas = &value
	return b
}

// WithDesiredReplicas sets the DesiredReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredReplicas field is set to the value of the last call.
func (b *FleetAutoscalerStatusApplyConfiguration) WithDesiredReplicas(value int32) *FleetAutoscalerStatusApplyConfiguration {
	b.DesiredReplicas = &value
	return b
}

// WithLastScaleTime sets the LastScaleTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastScaleTime field is set to the value of the last call.
func (b *FleetAutoscalerStatusApplyConfiguration) WithLastScaleTime(value v1.Time) *FleetAutoscalerStatusApplyConfiguration {
	b.LastScaleTime = &value
	return b
}

// WithAbleToScale sets the AbleToScale field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AbleToScale field is set to the value of the last call.
func (b *FleetAutoscalerStatusApplyConfiguration) WithAbleToScale(value bool) *FleetAutoscalerStatusApplyConfiguration {
	b.AbleToScale = &value
	return b
}

// WithScalingLimited sets the ScalingLimited field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScalingLimited field is set to the value of the last call.
func (b *FleetAutoscalerStatusApplyConfiguration) WithScalingLimited(value bool) *FleetAutoscalerStatusApplyConfiguration {
	b.ScalingLimited = &value
	return b
}