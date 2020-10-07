// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VPCLinkSpec defines the desired state of VPCLink
type VPCLinkSpec struct {
	// +kubebuilder:validation:Required
	Name             *string   `json:"name"`
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	// +kubebuilder:validation:Required
	SubnetIDs []*string          `json:"subnetIDs"`
	Tags      map[string]*string `json:"tags,omitempty"`
}

// VPCLinkStatus defines the observed state of VPCLink
type VPCLinkStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions           []*ackv1alpha1.Condition `json:"conditions"`
	CreatedDate          *metav1.Time             `json:"createdDate,omitempty"`
	VPCLinkID            *string                  `json:"vpcLinkID,omitempty"`
	VPCLinkStatus        *string                  `json:"vpcLinkStatus,omitempty"`
	VPCLinkStatusMessage *string                  `json:"vpcLinkStatusMessage,omitempty"`
	VPCLinkVersion       *string                  `json:"vpcLinkVersion,omitempty"`
}

// VPCLink is the Schema for the VPCLinks API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VPCLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCLinkSpec   `json:"spec,omitempty"`
	Status            VPCLinkStatus `json:"status,omitempty"`
}

// VPCLinkList contains a list of VPCLink
// +kubebuilder:object:root=true
type VPCLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCLink{}, &VPCLinkList{})
}
