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

type MongoDatabaseAutoscaleSettingsObservation struct {

	// The maximum throughput of the MongoDB database (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoDatabaseAutoscaleSettingsParameters struct {

	// The maximum throughput of the MongoDB database (RU/s). Must be between 1,000 and 1,000,000. Must be set in increments of 1,000. Conflicts with throughput.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoDatabaseObservation struct {

	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// An autoscale_settings block as defined below.
	AutoscaleSettings []MongoDatabaseAutoscaleSettingsObservation `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The ID of the Cosmos DB Mongo Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The throughput of the MongoDB database (RU/s). Must be set in increments of 100. The minimum value is 400.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

type MongoDatabaseParameters struct {

	// The name of the Cosmos DB Mongo Database to create the table within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// An autoscale_settings block as defined below.
	// +kubebuilder:validation:Optional
	AutoscaleSettings []MongoDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// The name of the resource group in which the Cosmos DB Mongo Database is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The throughput of the MongoDB database (RU/s). Must be set in increments of 100. The minimum value is 400.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// MongoDatabaseSpec defines the desired state of MongoDatabase
type MongoDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoDatabaseParameters `json:"forProvider"`
}

// MongoDatabaseStatus defines the observed state of MongoDatabase.
type MongoDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDatabase is the Schema for the MongoDatabases API. Manages a Mongo Database within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MongoDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoDatabaseSpec   `json:"spec"`
	Status            MongoDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDatabaseList contains a list of MongoDatabases
type MongoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDatabase `json:"items"`
}

// Repository type metadata.
var (
	MongoDatabase_Kind             = "MongoDatabase"
	MongoDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongoDatabase_Kind}.String()
	MongoDatabase_KindAPIVersion   = MongoDatabase_Kind + "." + CRDGroupVersion.String()
	MongoDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MongoDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MongoDatabase{}, &MongoDatabaseList{})
}
