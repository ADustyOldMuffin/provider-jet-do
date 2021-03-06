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

type DatabaseConnectionPoolObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	PrivateHost *string `json:"privateHost,omitempty" tf:"private_host,omitempty"`
}

type DatabaseConnectionPoolParameters struct {

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Required
	DBName *string `json:"dbName" tf:"db_name,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// DatabaseConnectionPoolSpec defines the desired state of DatabaseConnectionPool
type DatabaseConnectionPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseConnectionPoolParameters `json:"forProvider"`
}

// DatabaseConnectionPoolStatus defines the observed state of DatabaseConnectionPool.
type DatabaseConnectionPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseConnectionPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseConnectionPool is the Schema for the DatabaseConnectionPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,digitaloceanjet}
type DatabaseConnectionPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseConnectionPoolSpec   `json:"spec"`
	Status            DatabaseConnectionPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseConnectionPoolList contains a list of DatabaseConnectionPools
type DatabaseConnectionPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseConnectionPool `json:"items"`
}

// Repository type metadata.
var (
	DatabaseConnectionPool_Kind             = "DatabaseConnectionPool"
	DatabaseConnectionPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseConnectionPool_Kind}.String()
	DatabaseConnectionPool_KindAPIVersion   = DatabaseConnectionPool_Kind + "." + CRDGroupVersion.String()
	DatabaseConnectionPool_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseConnectionPool_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseConnectionPool{}, &DatabaseConnectionPoolList{})
}
