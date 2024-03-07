// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DockerHostedCleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type DockerHostedCleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type DockerHostedCleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type DockerHostedComponentInitParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type DockerHostedComponentObservation struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type DockerHostedComponentParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// +kubebuilder:validation:Optional
	ProprietaryComponents *bool `json:"proprietaryComponents" tf:"proprietary_components,omitempty"`
}

type DockerHostedDockerInitParameters struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth *bool `json:"forceBasicAuth,omitempty" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled *bool `json:"v1Enabled,omitempty" tf:"v1_enabled,omitempty"`
}

type DockerHostedDockerObservation struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth *bool `json:"forceBasicAuth,omitempty" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled *bool `json:"v1Enabled,omitempty" tf:"v1_enabled,omitempty"`
}

type DockerHostedDockerParameters struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	// +kubebuilder:validation:Optional
	ForceBasicAuth *bool `json:"forceBasicAuth" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	// +kubebuilder:validation:Optional
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	// +kubebuilder:validation:Optional
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	// +kubebuilder:validation:Optional
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	// +kubebuilder:validation:Optional
	V1Enabled *bool `json:"v1Enabled" tf:"v1_enabled,omitempty"`
}

type DockerHostedInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []DockerHostedCleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []DockerHostedComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	Docker []DockerHostedDockerInitParameters `json:"docker,omitempty" tf:"docker,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []DockerHostedStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerHostedObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []DockerHostedCleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []DockerHostedComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	Docker []DockerHostedDockerObservation `json:"docker,omitempty" tf:"docker,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []DockerHostedStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerHostedParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []DockerHostedCleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	// +kubebuilder:validation:Optional
	Component []DockerHostedComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	// +kubebuilder:validation:Optional
	Docker []DockerHostedDockerParameters `json:"docker,omitempty" tf:"docker,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	// +kubebuilder:validation:Optional
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	// +kubebuilder:validation:Optional
	Storage []DockerHostedStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerHostedStorageInitParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

type DockerHostedStorageObservation struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

type DockerHostedStorageParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	// +kubebuilder:validation:Optional
	BlobStoreName *string `json:"blobStoreName" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	// +kubebuilder:validation:Optional
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	// +kubebuilder:validation:Optional
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

// DockerHostedSpec defines the desired state of DockerHosted
type DockerHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DockerHostedParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DockerHostedInitParameters `json:"initProvider,omitempty"`
}

// DockerHostedStatus defines the observed state of DockerHosted.
type DockerHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DockerHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DockerHosted is the Schema for the DockerHosteds API. Use this resource to create a hosted docker repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type DockerHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.docker) || (has(self.initProvider) && has(self.initProvider.docker))",message="spec.forProvider.docker is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   DockerHostedSpec   `json:"spec"`
	Status DockerHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DockerHostedList contains a list of DockerHosteds
type DockerHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerHosted `json:"items"`
}

// Repository type metadata.
var (
	DockerHosted_Kind             = "DockerHosted"
	DockerHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DockerHosted_Kind}.String()
	DockerHosted_KindAPIVersion   = DockerHosted_Kind + "." + CRDGroupVersion.String()
	DockerHosted_GroupVersionKind = CRDGroupVersion.WithKind(DockerHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&DockerHosted{}, &DockerHostedList{})
}
