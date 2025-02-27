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

type SpringCloudActiveDeploymentObservation struct {

	// Specifies the name of Spring Cloud Deployment which is going to be active.
	DeploymentName *string `json:"deploymentName,omitempty" tf:"deployment_name,omitempty"`

	// The ID of the Spring Cloud Active Deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudActiveDeploymentParameters struct {

	// Specifies the name of Spring Cloud Deployment which is going to be active.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudJavaDeployment
	// +kubebuilder:validation:Optional
	DeploymentName *string `json:"deploymentName,omitempty" tf:"deployment_name,omitempty"`

	// Reference to a SpringCloudJavaDeployment in appplatform to populate deploymentName.
	// +kubebuilder:validation:Optional
	DeploymentNameRef *v1.Reference `json:"deploymentNameRef,omitempty" tf:"-"`

	// Selector for a SpringCloudJavaDeployment in appplatform to populate deploymentName.
	// +kubebuilder:validation:Optional
	DeploymentNameSelector *v1.Selector `json:"deploymentNameSelector,omitempty" tf:"-"`

	// Specifies the id of the Spring Cloud Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`
}

// SpringCloudActiveDeploymentSpec defines the desired state of SpringCloudActiveDeployment
type SpringCloudActiveDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudActiveDeploymentParameters `json:"forProvider"`
}

// SpringCloudActiveDeploymentStatus defines the observed state of SpringCloudActiveDeployment.
type SpringCloudActiveDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudActiveDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudActiveDeployment is the Schema for the SpringCloudActiveDeployments API. Manages an Active Azure Spring Cloud Deployment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudActiveDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudActiveDeploymentSpec   `json:"spec"`
	Status            SpringCloudActiveDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudActiveDeploymentList contains a list of SpringCloudActiveDeployments
type SpringCloudActiveDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudActiveDeployment `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudActiveDeployment_Kind             = "SpringCloudActiveDeployment"
	SpringCloudActiveDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudActiveDeployment_Kind}.String()
	SpringCloudActiveDeployment_KindAPIVersion   = SpringCloudActiveDeployment_Kind + "." + CRDGroupVersion.String()
	SpringCloudActiveDeployment_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudActiveDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudActiveDeployment{}, &SpringCloudActiveDeploymentList{})
}
