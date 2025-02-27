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

type VariableStringObservation struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The description of the Automation Variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ID of the Automation Variable.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The value of the Automation Variable as a string.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VariableStringParameters struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The description of the Automation Variable.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The value of the Automation Variable as a string.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// VariableStringSpec defines the desired state of VariableString
type VariableStringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableStringParameters `json:"forProvider"`
}

// VariableStringStatus defines the observed state of VariableString.
type VariableStringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableStringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VariableString is the Schema for the VariableStrings API. Manages a string variable in Azure Automation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VariableString struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VariableStringSpec   `json:"spec"`
	Status            VariableStringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableStringList contains a list of VariableStrings
type VariableStringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VariableString `json:"items"`
}

// Repository type metadata.
var (
	VariableString_Kind             = "VariableString"
	VariableString_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VariableString_Kind}.String()
	VariableString_KindAPIVersion   = VariableString_Kind + "." + CRDGroupVersion.String()
	VariableString_GroupVersionKind = CRDGroupVersion.WithKind(VariableString_Kind)
)

func init() {
	SchemeBuilder.Register(&VariableString{}, &VariableStringList{})
}
