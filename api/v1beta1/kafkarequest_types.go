/*
Copyright 2022.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KafkaRequestSpec defines the desired state of KafkaRequest
type KafkaRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Edit kafkarequest_types.go to remove/update
	CloudProvider           string `json:"cloudProvider,omitempty"`
	Name                    string `json:"name,omitempty"`
	ReauthenticationEnabled *bool  `json:"reauthenticationEnabled,omitempty"`
	Region                  string `json:"region,omitempty"`
	MultiAz                 *bool  `json:"multiAz,omitempty"`
}

// KafkaRequestStatus defines the observed state of KafkaRequest
type KafkaRequestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	CreatedAt           metav1.Time `json:"region,omitempty"`
	Href                string      `json:"href,omitempty"`
	Id                  string      `json:"id,omitempty"`
	InstanceType        string      `json:"instanceType,omitempty"`
	Kind                string      `json:"kind,omitempty"`
	Owner               string      `json:"owner,omitempty"`
	Message             string      `json:"message,omitempty"`
	UpdatedAt           metav1.Time `json:"updatedAt,omitempty"`
	Version             string      `json:"version,omitempty"`
	BootstrapServerHost string      `json:"bootstrapServerHost,omitempty"`
	Phase               KafkaPhase  `json:"phase,omitempty"`
}

type KafkaPhase string

// These are valid phases of a Pod
const (
	KafkaUnknown      KafkaPhase = "Unknown"
	KafkaReady        KafkaPhase = "Ready"
	KafkaFailed       KafkaPhase = "Failed"
	KafkaAccepted     KafkaPhase = "Accepted"
	KafkaPreparing    KafkaPhase = "Preparing"
	KafkaProvisioning KafkaPhase = "Provisioning"
	KafkaDeleting     KafkaPhase = "Deleting"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Region",type=string,JSONPath=`.spec.region`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="BootstrapServerHost",type=string,JSONPath=`.status.bootstrapServerHost`
// KafkaRequest is the Schema for the kafkarequests API
type KafkaRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaRequestSpec   `json:"spec,omitempty"`
	Status KafkaRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KafkaRequestList contains a list of KafkaRequest
type KafkaRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KafkaRequest{}, &KafkaRequestList{})
}
