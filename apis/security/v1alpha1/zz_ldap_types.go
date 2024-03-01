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

type LdapInitParameters struct {

	// (String) The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm *string `json:"authRealm,omitempty" tf:"auth_realm,omitempty"`

	// (String) Authentication scheme used for connecting to LDAP server
	// Authentication scheme used for connecting to LDAP server
	AuthSchema *string `json:"authSchema,omitempty" tf:"auth_schema,omitempty"`

	// (String) This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	AuthUsername *string `json:"authUsername,omitempty" tf:"auth_username,omitempty"`

	// (Number) How long to wait before retrying
	// How long to wait before retrying
	ConnectionRetryDelaySeconds *float64 `json:"connectionRetryDelaySeconds,omitempty" tf:"connection_retry_delay_seconds,omitempty"`

	// (Number) How long to wait before timeout
	// How long to wait before timeout
	ConnectionTimeoutSeconds *float64 `json:"connectionTimeoutSeconds,omitempty" tf:"connection_timeout_seconds,omitempty"`

	// (String) The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	GroupBaseDn *string `json:"groupBaseDn,omitempty" tf:"group_base_dn,omitempty"`

	// (String) This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIDAttribute *string `json:"groupIdAttribute,omitempty" tf:"group_id_attribute,omitempty"`

	// (String) LDAP attribute containing the usernames for the group. Required if groupType is static
	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	// (String) The format of user ID stored in the group member attribute. Required if groupType is static
	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat *string `json:"groupMemberFormat,omitempty" tf:"group_member_format,omitempty"`

	// (String) LDAP class for group objects. Required if groupType is static
	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass *string `json:"groupObjectClass,omitempty" tf:"group_object_class,omitempty"`

	// (Boolean) Are groups located in structures below the group base DN
	// Are groups located in structures below the group base DN
	GroupSubtree *bool `json:"groupSubtree,omitempty" tf:"group_subtree,omitempty"`

	// (String) Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	GroupType *string `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) LDAP server connection hostname
	// LDAP server connection hostname
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (Boolean) Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LdapGroupsAsRoles *bool `json:"ldapGroupsAsRoles,omitempty" tf:"ldap_groups_as_roles,omitempty"`

	// (Number) How many retry attempts
	// How many retry attempts
	MaxIncidentCount *float64 `json:"maxIncidentCount,omitempty" tf:"max_incident_count,omitempty"`

	// (String) LDAP server name
	// LDAP server name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) LDAP server connection port to use
	// LDAP server connection port to use
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) LDAP server connection Protocol to use
	// LDAP server connection Protocol to use
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) LDAP location to be added to the connection URL
	// LDAP location to be added to the connection URL
	SearchBase *string `json:"searchBase,omitempty" tf:"search_base,omitempty"`

	// (Boolean) Whether to use certificates stored in Nexus Repository Manager's truststore
	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// (String) The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	UserBaseDn *string `json:"userBaseDn,omitempty" tf:"user_base_dn,omitempty"`

	// (String) This is used to find an email address given the user ID
	// This is used to find an email address given the user ID
	UserEmailAddressAttribute *string `json:"userEmailAddressAttribute,omitempty" tf:"user_email_address_attribute,omitempty"`

	// (String) This is used to find a user given its user ID
	// This is used to find a user given its user ID
	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	// (String) LDAP search filter to limit user search
	// LDAP search filter to limit user search
	UserLdapFilter *string `json:"userLdapFilter,omitempty" tf:"user_ldap_filter,omitempty"`

	// (String) Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute *string `json:"userMemberOfAttribute,omitempty" tf:"user_member_of_attribute,omitempty"`

	// (String) LDAP class for user objects
	// LDAP class for user objects
	UserObjectClass *string `json:"userObjectClass,omitempty" tf:"user_object_class,omitempty"`

	// (String) If this field is blank the user will be authenticated against a bind with the LDAP server
	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute *string `json:"userPasswordAttribute,omitempty" tf:"user_password_attribute,omitempty"`

	// (String) This is used to find a real name given the user ID
	// This is used to find a real name given the user ID
	UserRealNameAttribute *string `json:"userRealNameAttribute,omitempty" tf:"user_real_name_attribute,omitempty"`

	// (Boolean) Are users located in structures below the user base DN?
	// Are users located in structures below the user base DN?
	UserSubtree *bool `json:"userSubtree,omitempty" tf:"user_subtree,omitempty"`
}

type LdapObservation struct {

	// (String) The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm *string `json:"authRealm,omitempty" tf:"auth_realm,omitempty"`

	// (String) Authentication scheme used for connecting to LDAP server
	// Authentication scheme used for connecting to LDAP server
	AuthSchema *string `json:"authSchema,omitempty" tf:"auth_schema,omitempty"`

	// (String) This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	AuthUsername *string `json:"authUsername,omitempty" tf:"auth_username,omitempty"`

	// (Number) How long to wait before retrying
	// How long to wait before retrying
	ConnectionRetryDelaySeconds *float64 `json:"connectionRetryDelaySeconds,omitempty" tf:"connection_retry_delay_seconds,omitempty"`

	// (Number) How long to wait before timeout
	// How long to wait before timeout
	ConnectionTimeoutSeconds *float64 `json:"connectionTimeoutSeconds,omitempty" tf:"connection_timeout_seconds,omitempty"`

	// (String) The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	GroupBaseDn *string `json:"groupBaseDn,omitempty" tf:"group_base_dn,omitempty"`

	// (String) This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIDAttribute *string `json:"groupIdAttribute,omitempty" tf:"group_id_attribute,omitempty"`

	// (String) LDAP attribute containing the usernames for the group. Required if groupType is static
	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	// (String) The format of user ID stored in the group member attribute. Required if groupType is static
	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat *string `json:"groupMemberFormat,omitempty" tf:"group_member_format,omitempty"`

	// (String) LDAP class for group objects. Required if groupType is static
	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass *string `json:"groupObjectClass,omitempty" tf:"group_object_class,omitempty"`

	// (Boolean) Are groups located in structures below the group base DN
	// Are groups located in structures below the group base DN
	GroupSubtree *bool `json:"groupSubtree,omitempty" tf:"group_subtree,omitempty"`

	// (String) Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	GroupType *string `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) LDAP server connection hostname
	// LDAP server connection hostname
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LdapGroupsAsRoles *bool `json:"ldapGroupsAsRoles,omitempty" tf:"ldap_groups_as_roles,omitempty"`

	// (Number) How many retry attempts
	// How many retry attempts
	MaxIncidentCount *float64 `json:"maxIncidentCount,omitempty" tf:"max_incident_count,omitempty"`

	// (String) LDAP server name
	// LDAP server name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) LDAP server connection port to use
	// LDAP server connection port to use
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) LDAP server connection Protocol to use
	// LDAP server connection Protocol to use
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) LDAP location to be added to the connection URL
	// LDAP location to be added to the connection URL
	SearchBase *string `json:"searchBase,omitempty" tf:"search_base,omitempty"`

	// (Boolean) Whether to use certificates stored in Nexus Repository Manager's truststore
	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// (String) The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	UserBaseDn *string `json:"userBaseDn,omitempty" tf:"user_base_dn,omitempty"`

	// (String) This is used to find an email address given the user ID
	// This is used to find an email address given the user ID
	UserEmailAddressAttribute *string `json:"userEmailAddressAttribute,omitempty" tf:"user_email_address_attribute,omitempty"`

	// (String) This is used to find a user given its user ID
	// This is used to find a user given its user ID
	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	// (String) LDAP search filter to limit user search
	// LDAP search filter to limit user search
	UserLdapFilter *string `json:"userLdapFilter,omitempty" tf:"user_ldap_filter,omitempty"`

	// (String) Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute *string `json:"userMemberOfAttribute,omitempty" tf:"user_member_of_attribute,omitempty"`

	// (String) LDAP class for user objects
	// LDAP class for user objects
	UserObjectClass *string `json:"userObjectClass,omitempty" tf:"user_object_class,omitempty"`

	// (String) If this field is blank the user will be authenticated against a bind with the LDAP server
	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute *string `json:"userPasswordAttribute,omitempty" tf:"user_password_attribute,omitempty"`

	// (String) This is used to find a real name given the user ID
	// This is used to find a real name given the user ID
	UserRealNameAttribute *string `json:"userRealNameAttribute,omitempty" tf:"user_real_name_attribute,omitempty"`

	// (Boolean) Are users located in structures below the user base DN?
	// Are users located in structures below the user base DN?
	UserSubtree *bool `json:"userSubtree,omitempty" tf:"user_subtree,omitempty"`
}

type LdapParameters struct {

	// (String, Sensitive) The password to bind with. Required if authScheme other than none.
	// The password to bind with. Required if authScheme other than none.
	// +kubebuilder:validation:Optional
	AuthPasswordSecretRef *v1.SecretKeySelector `json:"authPasswordSecretRef,omitempty" tf:"-"`

	// (String) The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	// +kubebuilder:validation:Optional
	AuthRealm *string `json:"authRealm,omitempty" tf:"auth_realm,omitempty"`

	// (String) Authentication scheme used for connecting to LDAP server
	// Authentication scheme used for connecting to LDAP server
	// +kubebuilder:validation:Optional
	AuthSchema *string `json:"authSchema,omitempty" tf:"auth_schema,omitempty"`

	// (String) This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	// +kubebuilder:validation:Optional
	AuthUsername *string `json:"authUsername,omitempty" tf:"auth_username,omitempty"`

	// (Number) How long to wait before retrying
	// How long to wait before retrying
	// +kubebuilder:validation:Optional
	ConnectionRetryDelaySeconds *float64 `json:"connectionRetryDelaySeconds,omitempty" tf:"connection_retry_delay_seconds,omitempty"`

	// (Number) How long to wait before timeout
	// How long to wait before timeout
	// +kubebuilder:validation:Optional
	ConnectionTimeoutSeconds *float64 `json:"connectionTimeoutSeconds,omitempty" tf:"connection_timeout_seconds,omitempty"`

	// (String) The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	// +kubebuilder:validation:Optional
	GroupBaseDn *string `json:"groupBaseDn,omitempty" tf:"group_base_dn,omitempty"`

	// (String) This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	// +kubebuilder:validation:Optional
	GroupIDAttribute *string `json:"groupIdAttribute,omitempty" tf:"group_id_attribute,omitempty"`

	// (String) LDAP attribute containing the usernames for the group. Required if groupType is static
	// LDAP attribute containing the usernames for the group. Required if groupType is static
	// +kubebuilder:validation:Optional
	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty" tf:"group_member_attribute,omitempty"`

	// (String) The format of user ID stored in the group member attribute. Required if groupType is static
	// The format of user ID stored in the group member attribute. Required if groupType is static
	// +kubebuilder:validation:Optional
	GroupMemberFormat *string `json:"groupMemberFormat,omitempty" tf:"group_member_format,omitempty"`

	// (String) LDAP class for group objects. Required if groupType is static
	// LDAP class for group objects. Required if groupType is static
	// +kubebuilder:validation:Optional
	GroupObjectClass *string `json:"groupObjectClass,omitempty" tf:"group_object_class,omitempty"`

	// (Boolean) Are groups located in structures below the group base DN
	// Are groups located in structures below the group base DN
	// +kubebuilder:validation:Optional
	GroupSubtree *bool `json:"groupSubtree,omitempty" tf:"group_subtree,omitempty"`

	// (String) Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	// +kubebuilder:validation:Optional
	GroupType *string `json:"groupType,omitempty" tf:"group_type,omitempty"`

	// (String) LDAP server connection hostname
	// LDAP server connection hostname
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// (Boolean) Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	// +kubebuilder:validation:Optional
	LdapGroupsAsRoles *bool `json:"ldapGroupsAsRoles,omitempty" tf:"ldap_groups_as_roles,omitempty"`

	// (Number) How many retry attempts
	// How many retry attempts
	// +kubebuilder:validation:Optional
	MaxIncidentCount *float64 `json:"maxIncidentCount,omitempty" tf:"max_incident_count,omitempty"`

	// (String) LDAP server name
	// LDAP server name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) LDAP server connection port to use
	// LDAP server connection port to use
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (String) LDAP server connection Protocol to use
	// LDAP server connection Protocol to use
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (String) LDAP location to be added to the connection URL
	// LDAP location to be added to the connection URL
	// +kubebuilder:validation:Optional
	SearchBase *string `json:"searchBase,omitempty" tf:"search_base,omitempty"`

	// (Boolean) Whether to use certificates stored in Nexus Repository Manager's truststore
	// Whether to use certificates stored in Nexus Repository Manager's truststore
	// +kubebuilder:validation:Optional
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// (String) The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	// +kubebuilder:validation:Optional
	UserBaseDn *string `json:"userBaseDn,omitempty" tf:"user_base_dn,omitempty"`

	// (String) This is used to find an email address given the user ID
	// This is used to find an email address given the user ID
	// +kubebuilder:validation:Optional
	UserEmailAddressAttribute *string `json:"userEmailAddressAttribute,omitempty" tf:"user_email_address_attribute,omitempty"`

	// (String) This is used to find a user given its user ID
	// This is used to find a user given its user ID
	// +kubebuilder:validation:Optional
	UserIDAttribute *string `json:"userIdAttribute,omitempty" tf:"user_id_attribute,omitempty"`

	// (String) LDAP search filter to limit user search
	// LDAP search filter to limit user search
	// +kubebuilder:validation:Optional
	UserLdapFilter *string `json:"userLdapFilter,omitempty" tf:"user_ldap_filter,omitempty"`

	// (String) Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	// +kubebuilder:validation:Optional
	UserMemberOfAttribute *string `json:"userMemberOfAttribute,omitempty" tf:"user_member_of_attribute,omitempty"`

	// (String) LDAP class for user objects
	// LDAP class for user objects
	// +kubebuilder:validation:Optional
	UserObjectClass *string `json:"userObjectClass,omitempty" tf:"user_object_class,omitempty"`

	// (String) If this field is blank the user will be authenticated against a bind with the LDAP server
	// If this field is blank the user will be authenticated against a bind with the LDAP server
	// +kubebuilder:validation:Optional
	UserPasswordAttribute *string `json:"userPasswordAttribute,omitempty" tf:"user_password_attribute,omitempty"`

	// (String) This is used to find a real name given the user ID
	// This is used to find a real name given the user ID
	// +kubebuilder:validation:Optional
	UserRealNameAttribute *string `json:"userRealNameAttribute,omitempty" tf:"user_real_name_attribute,omitempty"`

	// (Boolean) Are users located in structures below the user base DN?
	// Are users located in structures below the user base DN?
	// +kubebuilder:validation:Optional
	UserSubtree *bool `json:"userSubtree,omitempty" tf:"user_subtree,omitempty"`
}

// LdapSpec defines the desired state of Ldap
type LdapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LdapParameters `json:"forProvider"`
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
	InitProvider LdapInitParameters `json:"initProvider,omitempty"`
}

// LdapStatus defines the observed state of Ldap.
type LdapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LdapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ldap is the Schema for the Ldaps API. Use this resource to create a Nexus Security LDAP configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type Ldap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authSchema) || (has(self.initProvider) && has(self.initProvider.authSchema))",message="spec.forProvider.authSchema is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authUsername) || (has(self.initProvider) && has(self.initProvider.authUsername))",message="spec.forProvider.authUsername is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionRetryDelaySeconds) || (has(self.initProvider) && has(self.initProvider.connectionRetryDelaySeconds))",message="spec.forProvider.connectionRetryDelaySeconds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionTimeoutSeconds) || (has(self.initProvider) && has(self.initProvider.connectionTimeoutSeconds))",message="spec.forProvider.connectionTimeoutSeconds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupType) || (has(self.initProvider) && has(self.initProvider.groupType))",message="spec.forProvider.groupType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.host) || (has(self.initProvider) && has(self.initProvider.host))",message="spec.forProvider.host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxIncidentCount) || (has(self.initProvider) && has(self.initProvider.maxIncidentCount))",message="spec.forProvider.maxIncidentCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.searchBase) || (has(self.initProvider) && has(self.initProvider.searchBase))",message="spec.forProvider.searchBase is a required parameter"
	Spec   LdapSpec   `json:"spec"`
	Status LdapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LdapList contains a list of Ldaps
type LdapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ldap `json:"items"`
}

// Repository type metadata.
var (
	Ldap_Kind             = "Ldap"
	Ldap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ldap_Kind}.String()
	Ldap_KindAPIVersion   = Ldap_Kind + "." + CRDGroupVersion.String()
	Ldap_GroupVersionKind = CRDGroupVersion.WithKind(Ldap_Kind)
)

func init() {
	SchemeBuilder.Register(&Ldap{}, &LdapList{})
}
