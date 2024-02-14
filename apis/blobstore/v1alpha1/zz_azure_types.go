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

type AuthenticationInitParameters struct {

	// (String) The account key. Required if authentication_method is ACCOUNTKEY
	// The account key. Required if `authentication_method` is `ACCOUNTKEY`
	AccountKey *string `json:"accountKey,omitempty" tf:"account_key,omitempty"`

	// (String) The type of Azure authentication to use. Possible values: ACCOUNTKEY and MANAGEDIDENTITY
	// The type of Azure authentication to use. Possible values: `ACCOUNTKEY` and `MANAGEDIDENTITY`
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`
}

type AuthenticationObservation struct {

	// (String) The account key. Required if authentication_method is ACCOUNTKEY
	// The account key. Required if `authentication_method` is `ACCOUNTKEY`
	AccountKey *string `json:"accountKey,omitempty" tf:"account_key,omitempty"`

	// (String) The type of Azure authentication to use. Possible values: ACCOUNTKEY and MANAGEDIDENTITY
	// The type of Azure authentication to use. Possible values: `ACCOUNTKEY` and `MANAGEDIDENTITY`
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`
}

type AuthenticationParameters struct {

	// (String) The account key. Required if authentication_method is ACCOUNTKEY
	// The account key. Required if `authentication_method` is `ACCOUNTKEY`
	// +kubebuilder:validation:Optional
	AccountKey *string `json:"accountKey,omitempty" tf:"account_key,omitempty"`

	// (String) The type of Azure authentication to use. Possible values: ACCOUNTKEY and MANAGEDIDENTITY
	// The type of Azure authentication to use. Possible values: `ACCOUNTKEY` and `MANAGEDIDENTITY`
	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod" tf:"authentication_method,omitempty"`
}

type AzureInitParameters struct {

	// (Block List, Min: 1, Max: 1) The Azure specific configuration details for the Azure object that'll contain the blob store (see below for nested schema)
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration []BucketConfigurationInitParameters `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []SoftQuotaInitParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type AzureObservation struct {

	// (Number) Count of blobs
	// Count of blobs
	BlobCount *float64 `json:"blobCount,omitempty" tf:"blob_count,omitempty"`

	// (Block List, Min: 1, Max: 1) The Azure specific configuration details for the Azure object that'll contain the blob store (see below for nested schema)
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration []BucketConfigurationObservation `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []SoftQuotaObservation `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`

	// (Number) The total size of the blobstore in Bytes
	// The total size of the blobstore in Bytes
	TotalSizeInBytes *float64 `json:"totalSizeInBytes,omitempty" tf:"total_size_in_bytes,omitempty"`
}

type AzureParameters struct {

	// (Block List, Min: 1, Max: 1) The Azure specific configuration details for the Azure object that'll contain the blob store (see below for nested schema)
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	// +kubebuilder:validation:Optional
	BucketConfiguration []BucketConfigurationParameters `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	// +kubebuilder:validation:Optional
	SoftQuota []SoftQuotaParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type BucketConfigurationInitParameters struct {

	// (String) Account name found under Access keys for the storage account
	// Account name found under Access keys for the storage account
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// (Block List, Min: 1, Max: 1) The Azure specific authentication details (see below for nested schema)
	// The Azure specific authentication details
	Authentication []AuthenticationInitParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// (String) The name of an existing container to be used for storage
	// The name of an existing container to be used for storage
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`
}

type BucketConfigurationObservation struct {

	// (String) Account name found under Access keys for the storage account
	// Account name found under Access keys for the storage account
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// (Block List, Min: 1, Max: 1) The Azure specific authentication details (see below for nested schema)
	// The Azure specific authentication details
	Authentication []AuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// (String) The name of an existing container to be used for storage
	// The name of an existing container to be used for storage
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`
}

type BucketConfigurationParameters struct {

	// (String) Account name found under Access keys for the storage account
	// Account name found under Access keys for the storage account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// (Block List, Min: 1, Max: 1) The Azure specific authentication details (see below for nested schema)
	// The Azure specific authentication details
	// +kubebuilder:validation:Optional
	Authentication []AuthenticationParameters `json:"authentication" tf:"authentication,omitempty"`

	// (String) The name of an existing container to be used for storage
	// The name of an existing container to be used for storage
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`
}

type SoftQuotaInitParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SoftQuotaObservation struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SoftQuotaParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// AzureSpec defines the desired state of Azure
type AzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureParameters `json:"forProvider"`
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
	InitProvider AzureInitParameters `json:"initProvider,omitempty"`
}

// AzureStatus defines the observed state of Azure.
type AzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Azure is the Schema for the Azures API. ~> PRO Feature Use this resource to create a Nexus Azure blobstore.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type Azure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketConfiguration) || (has(self.initProvider) && has(self.initProvider.bucketConfiguration))",message="spec.forProvider.bucketConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AzureSpec   `json:"spec"`
	Status AzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureList contains a list of Azures
type AzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Azure `json:"items"`
}

// Repository type metadata.
var (
	Azure_Kind             = "Azure"
	Azure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Azure_Kind}.String()
	Azure_KindAPIVersion   = Azure_Kind + "." + CRDGroupVersion.String()
	Azure_GroupVersionKind = CRDGroupVersion.WithKind(Azure_Kind)
)

func init() {
	SchemeBuilder.Register(&Azure{}, &AzureList{})
}
