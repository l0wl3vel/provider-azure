/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NamespaceSchemaGroupObservation struct {

	// The ID of the EventHub Namespace Schema Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Specifies the compatibility of this schema group. Possible values are None, Backward, Forward. Changing this forces a new resource to be created.
	SchemaCompatibility *string `json:"schemaCompatibility,omitempty" tf:"schema_compatibility,omitempty"`

	// Specifies the Type of this schema group. Possible values are Avro, Unknown. Changing this forces a new resource to be created.
	SchemaType *string `json:"schemaType,omitempty" tf:"schema_type,omitempty"`
}

type NamespaceSchemaGroupParameters struct {

	// Specifies the ID of the EventHub Namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// Specifies the compatibility of this schema group. Possible values are None, Backward, Forward. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SchemaCompatibility *string `json:"schemaCompatibility,omitempty" tf:"schema_compatibility,omitempty"`

	// Specifies the Type of this schema group. Possible values are Avro, Unknown. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SchemaType *string `json:"schemaType,omitempty" tf:"schema_type,omitempty"`
}

// NamespaceSchemaGroupSpec defines the desired state of NamespaceSchemaGroup
type NamespaceSchemaGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamespaceSchemaGroupParameters `json:"forProvider"`
}

// NamespaceSchemaGroupStatus defines the observed state of NamespaceSchemaGroup.
type NamespaceSchemaGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamespaceSchemaGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceSchemaGroup is the Schema for the NamespaceSchemaGroups API. Manages a Schema Group for a EventHub Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NamespaceSchemaGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.schemaCompatibility)",message="schemaCompatibility is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.schemaType)",message="schemaType is a required parameter"
	Spec   NamespaceSchemaGroupSpec   `json:"spec"`
	Status NamespaceSchemaGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceSchemaGroupList contains a list of NamespaceSchemaGroups
type NamespaceSchemaGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceSchemaGroup `json:"items"`
}

// Repository type metadata.
var (
	NamespaceSchemaGroup_Kind             = "NamespaceSchemaGroup"
	NamespaceSchemaGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamespaceSchemaGroup_Kind}.String()
	NamespaceSchemaGroup_KindAPIVersion   = NamespaceSchemaGroup_Kind + "." + CRDGroupVersion.String()
	NamespaceSchemaGroup_GroupVersionKind = CRDGroupVersion.WithKind(NamespaceSchemaGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NamespaceSchemaGroup{}, &NamespaceSchemaGroupList{})
}
