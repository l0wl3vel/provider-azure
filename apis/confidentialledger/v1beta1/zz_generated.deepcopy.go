//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureadBasedServicePrincipalObservation) DeepCopyInto(out *AzureadBasedServicePrincipalObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureadBasedServicePrincipalObservation.
func (in *AzureadBasedServicePrincipalObservation) DeepCopy() *AzureadBasedServicePrincipalObservation {
	if in == nil {
		return nil
	}
	out := new(AzureadBasedServicePrincipalObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureadBasedServicePrincipalParameters) DeepCopyInto(out *AzureadBasedServicePrincipalParameters) {
	*out = *in
	if in.LedgerRoleName != nil {
		in, out := &in.LedgerRoleName, &out.LedgerRoleName
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureadBasedServicePrincipalParameters.
func (in *AzureadBasedServicePrincipalParameters) DeepCopy() *AzureadBasedServicePrincipalParameters {
	if in == nil {
		return nil
	}
	out := new(AzureadBasedServicePrincipalParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateBasedSecurityPrincipalObservation) DeepCopyInto(out *CertificateBasedSecurityPrincipalObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateBasedSecurityPrincipalObservation.
func (in *CertificateBasedSecurityPrincipalObservation) DeepCopy() *CertificateBasedSecurityPrincipalObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateBasedSecurityPrincipalObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateBasedSecurityPrincipalParameters) DeepCopyInto(out *CertificateBasedSecurityPrincipalParameters) {
	*out = *in
	if in.LedgerRoleName != nil {
		in, out := &in.LedgerRoleName, &out.LedgerRoleName
		*out = new(string)
		**out = **in
	}
	if in.PemPublicKey != nil {
		in, out := &in.PemPublicKey, &out.PemPublicKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateBasedSecurityPrincipalParameters.
func (in *CertificateBasedSecurityPrincipalParameters) DeepCopy() *CertificateBasedSecurityPrincipalParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateBasedSecurityPrincipalParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ledger) DeepCopyInto(out *Ledger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ledger.
func (in *Ledger) DeepCopy() *Ledger {
	if in == nil {
		return nil
	}
	out := new(Ledger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ledger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LedgerList) DeepCopyInto(out *LedgerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ledger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LedgerList.
func (in *LedgerList) DeepCopy() *LedgerList {
	if in == nil {
		return nil
	}
	out := new(LedgerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LedgerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LedgerObservation) DeepCopyInto(out *LedgerObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdentityServiceEndpoint != nil {
		in, out := &in.IdentityServiceEndpoint, &out.IdentityServiceEndpoint
		*out = new(string)
		**out = **in
	}
	if in.LedgerEndpoint != nil {
		in, out := &in.LedgerEndpoint, &out.LedgerEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LedgerObservation.
func (in *LedgerObservation) DeepCopy() *LedgerObservation {
	if in == nil {
		return nil
	}
	out := new(LedgerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LedgerParameters) DeepCopyInto(out *LedgerParameters) {
	*out = *in
	if in.AzureadBasedServicePrincipal != nil {
		in, out := &in.AzureadBasedServicePrincipal, &out.AzureadBasedServicePrincipal
		*out = make([]AzureadBasedServicePrincipalParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CertificateBasedSecurityPrincipal != nil {
		in, out := &in.CertificateBasedSecurityPrincipal, &out.CertificateBasedSecurityPrincipal
		*out = make([]CertificateBasedSecurityPrincipalParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LedgerType != nil {
		in, out := &in.LedgerType, &out.LedgerType
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LedgerParameters.
func (in *LedgerParameters) DeepCopy() *LedgerParameters {
	if in == nil {
		return nil
	}
	out := new(LedgerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LedgerSpec) DeepCopyInto(out *LedgerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LedgerSpec.
func (in *LedgerSpec) DeepCopy() *LedgerSpec {
	if in == nil {
		return nil
	}
	out := new(LedgerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LedgerStatus) DeepCopyInto(out *LedgerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LedgerStatus.
func (in *LedgerStatus) DeepCopy() *LedgerStatus {
	if in == nil {
		return nil
	}
	out := new(LedgerStatus)
	in.DeepCopyInto(out)
	return out
}
