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

type PrivateDNSSRVRecordObservation struct {

	// The FQDN of the DNS SRV Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS SRV Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more record blocks as defined below.
	Record []PrivateDNSSRVRecordRecordObservation `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type PrivateDNSSRVRecordParameters struct {

	// One or more record blocks as defined below.
	// +kubebuilder:validation:Optional
	Record []PrivateDNSSRVRecordRecordParameters `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

type PrivateDNSSRVRecordRecordObservation struct {

	// The Port the service is listening on.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The priority of the SRV record.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The FQDN of the service.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The Weight of the SRV record.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type PrivateDNSSRVRecordRecordParameters struct {

	// The Port the service is listening on.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The priority of the SRV record.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// The FQDN of the service.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// The Weight of the SRV record.
	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// PrivateDNSSRVRecordSpec defines the desired state of PrivateDNSSRVRecord
type PrivateDNSSRVRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSSRVRecordParameters `json:"forProvider"`
}

// PrivateDNSSRVRecordStatus defines the observed state of PrivateDNSSRVRecord.
type PrivateDNSSRVRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSSRVRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSSRVRecord is the Schema for the PrivateDNSSRVRecords API. Manages a Private DNS SRV Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSSRVRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.record)",message="record is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ttl)",message="ttl is a required parameter"
	Spec   PrivateDNSSRVRecordSpec   `json:"spec"`
	Status PrivateDNSSRVRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSSRVRecordList contains a list of PrivateDNSSRVRecords
type PrivateDNSSRVRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSSRVRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSSRVRecord_Kind             = "PrivateDNSSRVRecord"
	PrivateDNSSRVRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSSRVRecord_Kind}.String()
	PrivateDNSSRVRecord_KindAPIVersion   = PrivateDNSSRVRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSSRVRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSSRVRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSSRVRecord{}, &PrivateDNSSRVRecordList{})
}
