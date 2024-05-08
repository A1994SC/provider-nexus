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

type YumHostedCleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type YumHostedCleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type YumHostedCleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type YumHostedComponentInitParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type YumHostedComponentObservation struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type YumHostedComponentParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// +kubebuilder:validation:Optional
	ProprietaryComponents *bool `json:"proprietaryComponents" tf:"proprietary_components,omitempty"`
}

type YumHostedInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []YumHostedCleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []YumHostedComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Validate that all paths are RPMs or yum metadata. Possible values: STRICT or PERMISSIVE
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy *string `json:"deployPolicy,omitempty" tf:"deploy_policy,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// 5
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth *float64 `json:"repodataDepth,omitempty" tf:"repodata_depth,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []YumHostedStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type YumHostedObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []YumHostedCleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []YumHostedComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Validate that all paths are RPMs or yum metadata. Possible values: STRICT or PERMISSIVE
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy *string `json:"deployPolicy,omitempty" tf:"deploy_policy,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// 5
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth *float64 `json:"repodataDepth,omitempty" tf:"repodata_depth,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []YumHostedStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type YumHostedParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []YumHostedCleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	// +kubebuilder:validation:Optional
	Component []YumHostedComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Validate that all paths are RPMs or yum metadata. Possible values: STRICT or PERMISSIVE
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	// +kubebuilder:validation:Optional
	DeployPolicy *string `json:"deployPolicy,omitempty" tf:"deploy_policy,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	// +kubebuilder:validation:Optional
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// 5
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	// +kubebuilder:validation:Optional
	RepodataDepth *float64 `json:"repodataDepth,omitempty" tf:"repodata_depth,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	// +kubebuilder:validation:Optional
	Storage []YumHostedStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type YumHostedStorageInitParameters struct {

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

type YumHostedStorageObservation struct {

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

type YumHostedStorageParameters struct {

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

// YumHostedSpec defines the desired state of YumHosted
type YumHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     YumHostedParameters `json:"forProvider"`
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
	InitProvider YumHostedInitParameters `json:"initProvider,omitempty"`
}

// YumHostedStatus defines the observed state of YumHosted.
type YumHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        YumHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// YumHosted is the Schema for the YumHosteds API. Use this resource to create a hosted yum repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type YumHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   YumHostedSpec   `json:"spec"`
	Status YumHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// YumHostedList contains a list of YumHosteds
type YumHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []YumHosted `json:"items"`
}

// Repository type metadata.
var (
	YumHosted_Kind             = "YumHosted"
	YumHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: YumHosted_Kind}.String()
	YumHosted_KindAPIVersion   = YumHosted_Kind + "." + CRDGroupVersion.String()
	YumHosted_GroupVersionKind = CRDGroupVersion.WithKind(YumHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&YumHosted{}, &YumHostedList{})
}
