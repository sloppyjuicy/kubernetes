/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	storagemigrationv1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MigrationConditionApplyConfiguration represents a declarative configuration of the MigrationCondition type for use
// with apply.
type MigrationConditionApplyConfiguration struct {
	Type           *storagemigrationv1alpha1.MigrationConditionType `json:"type,omitempty"`
	Status         *v1.ConditionStatus                              `json:"status,omitempty"`
	LastUpdateTime *metav1.Time                                     `json:"lastUpdateTime,omitempty"`
	Reason         *string                                          `json:"reason,omitempty"`
	Message        *string                                          `json:"message,omitempty"`
}

// MigrationConditionApplyConfiguration constructs a declarative configuration of the MigrationCondition type for use with
// apply.
func MigrationCondition() *MigrationConditionApplyConfiguration {
	return &MigrationConditionApplyConfiguration{}
}
func (b MigrationConditionApplyConfiguration) IsApplyConfiguration() {}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *MigrationConditionApplyConfiguration) WithType(value storagemigrationv1alpha1.MigrationConditionType) *MigrationConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *MigrationConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *MigrationConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *MigrationConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *MigrationConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *MigrationConditionApplyConfiguration) WithReason(value string) *MigrationConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *MigrationConditionApplyConfiguration) WithMessage(value string) *MigrationConditionApplyConfiguration {
	b.Message = &value
	return b
}
