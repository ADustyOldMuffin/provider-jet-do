/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseUserObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DatabaseUserParameters struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	MySQLAuthPlugin *string `json:"mysqlAuthPlugin,omitempty" tf:"mysql_auth_plugin,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// DatabaseUserSpec defines the desired state of DatabaseUser
type DatabaseUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseUserParameters `json:"forProvider"`
}

// DatabaseUserStatus defines the observed state of DatabaseUser.
type DatabaseUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseUser is the Schema for the DatabaseUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,digitaloceanjet}
type DatabaseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseUserSpec   `json:"spec"`
	Status            DatabaseUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseUserList contains a list of DatabaseUsers
type DatabaseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseUser `json:"items"`
}

// Repository type metadata.
var (
	DatabaseUser_Kind             = "DatabaseUser"
	DatabaseUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseUser_Kind}.String()
	DatabaseUser_KindAPIVersion   = DatabaseUser_Kind + "." + CRDGroupVersion.String()
	DatabaseUser_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseUser_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseUser{}, &DatabaseUserList{})
}
