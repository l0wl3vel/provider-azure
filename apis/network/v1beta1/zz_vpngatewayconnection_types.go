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

type CustomBGPAddressObservation struct {

	// The custom bgp ip address which belongs to the IP Configuration.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The ID of the IP Configuration which belongs to the VPN Gateway.
	IPConfigurationID *string `json:"ipConfigurationId,omitempty" tf:"ip_configuration_id,omitempty"`
}

type CustomBGPAddressParameters struct {

	// The custom bgp ip address which belongs to the IP Configuration.
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The ID of the IP Configuration which belongs to the VPN Gateway.
	// +kubebuilder:validation:Required
	IPConfigurationID *string `json:"ipConfigurationId" tf:"ip_configuration_id,omitempty"`
}

type VPNGatewayConnectionObservation struct {

	// The ID of the VPN Gateway Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether Internet Security is enabled for this VPN Connection. Defaults to false.
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
	RemoteVPNSiteID *string `json:"remoteVpnSiteId,omitempty" tf:"remote_vpn_site_id,omitempty"`

	// A routing block as defined below. If this is not specified, there will be a default route table created implicitly.
	Routing []VPNGatewayConnectionRoutingObservation `json:"routing,omitempty" tf:"routing,omitempty"`

	// One or more traffic_selector_policy blocks as defined below.
	TrafficSelectorPolicy []VPNGatewayConnectionTrafficSelectorPolicyObservation `json:"trafficSelectorPolicy,omitempty" tf:"traffic_selector_policy,omitempty"`

	// The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// One or more vpn_link blocks as defined below.
	VPNLink []VPNLinkObservation `json:"vpnLink,omitempty" tf:"vpn_link,omitempty"`
}

type VPNGatewayConnectionParameters struct {

	// Whether Internet Security is enabled for this VPN Connection. Defaults to false.
	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The ID of the remote VPN Site, which will connect to the VPN Gateway. Changing this forces a new VPN Gateway Connection to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VPNSite
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RemoteVPNSiteID *string `json:"remoteVpnSiteId,omitempty" tf:"remote_vpn_site_id,omitempty"`

	// Reference to a VPNSite in network to populate remoteVpnSiteId.
	// +kubebuilder:validation:Optional
	RemoteVPNSiteIDRef *v1.Reference `json:"remoteVpnSiteIdRef,omitempty" tf:"-"`

	// Selector for a VPNSite in network to populate remoteVpnSiteId.
	// +kubebuilder:validation:Optional
	RemoteVPNSiteIDSelector *v1.Selector `json:"remoteVpnSiteIdSelector,omitempty" tf:"-"`

	// A routing block as defined below. If this is not specified, there will be a default route table created implicitly.
	// +kubebuilder:validation:Optional
	Routing []VPNGatewayConnectionRoutingParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// One or more traffic_selector_policy blocks as defined below.
	// +kubebuilder:validation:Optional
	TrafficSelectorPolicy []VPNGatewayConnectionTrafficSelectorPolicyParameters `json:"trafficSelectorPolicy,omitempty" tf:"traffic_selector_policy,omitempty"`

	// The ID of the VPN Gateway that this VPN Gateway Connection belongs to. Changing this forces a new VPN Gateway Connection to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in network to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in network to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`

	// One or more vpn_link blocks as defined below.
	// +kubebuilder:validation:Optional
	VPNLink []VPNLinkParameters `json:"vpnLink,omitempty" tf:"vpn_link,omitempty"`
}

type VPNGatewayConnectionRoutingObservation struct {

	// The ID of the Route Table associated with this VPN Connection.
	AssociatedRouteTable *string `json:"associatedRouteTable,omitempty" tf:"associated_route_table,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for inbound learned routes.
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for outbound advertised routes.
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	PropagatedRouteTable []VPNGatewayConnectionRoutingPropagatedRouteTableObservation `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type VPNGatewayConnectionRoutingParameters struct {

	// The ID of the Route Table associated with this VPN Connection.
	// +kubebuilder:validation:Required
	AssociatedRouteTable *string `json:"associatedRouteTable" tf:"associated_route_table,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for inbound learned routes.
	// +kubebuilder:validation:Optional
	InboundRouteMapID *string `json:"inboundRouteMapId,omitempty" tf:"inbound_route_map_id,omitempty"`

	// The resource ID of the Route Map associated with this Routing Configuration for outbound advertised routes.
	// +kubebuilder:validation:Optional
	OutboundRouteMapID *string `json:"outboundRouteMapId,omitempty" tf:"outbound_route_map_id,omitempty"`

	// A propagated_route_table block as defined below.
	// +kubebuilder:validation:Optional
	PropagatedRouteTable []VPNGatewayConnectionRoutingPropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`
}

type VPNGatewayConnectionRoutingPropagatedRouteTableObservation struct {

	// A list of labels to assign to this route table.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of Route Table IDs to associated with this VPN Gateway Connection.
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`
}

type VPNGatewayConnectionRoutingPropagatedRouteTableParameters struct {

	// A list of labels to assign to this route table.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of Route Table IDs to associated with this VPN Gateway Connection.
	// +kubebuilder:validation:Required
	RouteTableIds []*string `json:"routeTableIds" tf:"route_table_ids,omitempty"`
}

type VPNGatewayConnectionTrafficSelectorPolicyObservation struct {

	// A list of local address spaces in CIDR format for this VPN Gateway Connection.
	LocalAddressRanges []*string `json:"localAddressRanges,omitempty" tf:"local_address_ranges,omitempty"`

	// A list of remote address spaces in CIDR format for this VPN Gateway Connection.
	RemoteAddressRanges []*string `json:"remoteAddressRanges,omitempty" tf:"remote_address_ranges,omitempty"`
}

type VPNGatewayConnectionTrafficSelectorPolicyParameters struct {

	// A list of local address spaces in CIDR format for this VPN Gateway Connection.
	// +kubebuilder:validation:Required
	LocalAddressRanges []*string `json:"localAddressRanges" tf:"local_address_ranges,omitempty"`

	// A list of remote address spaces in CIDR format for this VPN Gateway Connection.
	// +kubebuilder:validation:Required
	RemoteAddressRanges []*string `json:"remoteAddressRanges" tf:"remote_address_ranges,omitempty"`
}

type VPNLinkIpsecPolicyObservation struct {

	// The DH Group used in IKE Phase 1 for initial SA. Possible values are None, DHGroup1, DHGroup2, DHGroup14, DHGroup24, DHGroup2048, ECP256, ECP384.
	DhGroup *string `json:"dhGroup,omitempty" tf:"dh_group,omitempty"`

	// The IPSec encryption algorithm (IKE phase 1). Possible values are AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256, None.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// The IKE encryption algorithm (IKE phase 2). Possible values are DES, DES3, AES128, AES192, AES256, GCMAES128, GCMAES256.
	IkeEncryptionAlgorithm *string `json:"ikeEncryptionAlgorithm,omitempty" tf:"ike_encryption_algorithm,omitempty"`

	// The IKE integrity algorithm (IKE phase 2). Possible values are MD5, SHA1, SHA256, SHA384, GCMAES128, GCMAES256.
	IkeIntegrityAlgorithm *string `json:"ikeIntegrityAlgorithm,omitempty" tf:"ike_integrity_algorithm,omitempty"`

	// The IPSec integrity algorithm (IKE phase 1). Possible values are MD5, SHA1, SHA256, GCMAES128, GCMAES192, GCMAES256.
	IntegrityAlgorithm *string `json:"integrityAlgorithm,omitempty" tf:"integrity_algorithm,omitempty"`

	// The Pfs Group used in IKE Phase 2 for the new child SA. Possible values are None, PFS1, PFS2, PFS14, PFS24, PFS2048, PFSMM, ECP256, ECP384.
	PfsGroup *string `json:"pfsGroup,omitempty" tf:"pfs_group,omitempty"`

	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for the site to site VPN tunnel.
	SaDataSizeKb *float64 `json:"saDataSizeKb,omitempty" tf:"sa_data_size_kb,omitempty"`

	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for the site to site VPN tunnel.
	SaLifetimeSec *float64 `json:"saLifetimeSec,omitempty" tf:"sa_lifetime_sec,omitempty"`
}

type VPNLinkIpsecPolicyParameters struct {

	// The DH Group used in IKE Phase 1 for initial SA. Possible values are None, DHGroup1, DHGroup2, DHGroup14, DHGroup24, DHGroup2048, ECP256, ECP384.
	// +kubebuilder:validation:Required
	DhGroup *string `json:"dhGroup" tf:"dh_group,omitempty"`

	// The IPSec encryption algorithm (IKE phase 1). Possible values are AES128, AES192, AES256, DES, DES3, GCMAES128, GCMAES192, GCMAES256, None.
	// +kubebuilder:validation:Required
	EncryptionAlgorithm *string `json:"encryptionAlgorithm" tf:"encryption_algorithm,omitempty"`

	// The IKE encryption algorithm (IKE phase 2). Possible values are DES, DES3, AES128, AES192, AES256, GCMAES128, GCMAES256.
	// +kubebuilder:validation:Required
	IkeEncryptionAlgorithm *string `json:"ikeEncryptionAlgorithm" tf:"ike_encryption_algorithm,omitempty"`

	// The IKE integrity algorithm (IKE phase 2). Possible values are MD5, SHA1, SHA256, SHA384, GCMAES128, GCMAES256.
	// +kubebuilder:validation:Required
	IkeIntegrityAlgorithm *string `json:"ikeIntegrityAlgorithm" tf:"ike_integrity_algorithm,omitempty"`

	// The IPSec integrity algorithm (IKE phase 1). Possible values are MD5, SHA1, SHA256, GCMAES128, GCMAES192, GCMAES256.
	// +kubebuilder:validation:Required
	IntegrityAlgorithm *string `json:"integrityAlgorithm" tf:"integrity_algorithm,omitempty"`

	// The Pfs Group used in IKE Phase 2 for the new child SA. Possible values are None, PFS1, PFS2, PFS14, PFS24, PFS2048, PFSMM, ECP256, ECP384.
	// +kubebuilder:validation:Required
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group,omitempty"`

	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for the site to site VPN tunnel.
	// +kubebuilder:validation:Required
	SaDataSizeKb *float64 `json:"saDataSizeKb" tf:"sa_data_size_kb,omitempty"`

	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for the site to site VPN tunnel.
	// +kubebuilder:validation:Required
	SaLifetimeSec *float64 `json:"saLifetimeSec" tf:"sa_lifetime_sec,omitempty"`
}

type VPNLinkObservation struct {

	// Should the BGP be enabled? Defaults to false. Changing this forces a new VPN Gateway Connection to be created.
	BGPEnabled *bool `json:"bgpEnabled,omitempty" tf:"bgp_enabled,omitempty"`

	// The expected connection bandwidth in MBPS. Defaults to 10.
	BandwidthMbps *float64 `json:"bandwidthMbps,omitempty" tf:"bandwidth_mbps,omitempty"`

	// The connection mode of this VPN Link. Possible values are Default, InitiatorOnly and ResponderOnly. Defaults to Default.
	ConnectionMode *string `json:"connectionMode,omitempty" tf:"connection_mode,omitempty"`

	// One or more custom_bgp_address blocks as defined below.
	CustomBGPAddress []CustomBGPAddressObservation `json:"customBgpAddress,omitempty" tf:"custom_bgp_address,omitempty"`

	// A list of the egress NAT Rule Ids.
	EgressNATRuleIds []*string `json:"egressNatRuleIds,omitempty" tf:"egress_nat_rule_ids,omitempty"`

	// A list of the ingress NAT Rule Ids.
	IngressNATRuleIds []*string `json:"ingressNatRuleIds,omitempty" tf:"ingress_nat_rule_ids,omitempty"`

	// One or more ipsec_policy blocks as defined above.
	IpsecPolicy []VPNLinkIpsecPolicyObservation `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// Whether to use local Azure IP to initiate connection? Defaults to false.
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled,omitempty"`

	// The name which should be used for this VPN Link Connection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to enable policy-based traffic selectors? Defaults to false.
	PolicyBasedTrafficSelectorEnabled *bool `json:"policyBasedTrafficSelectorEnabled,omitempty" tf:"policy_based_traffic_selector_enabled,omitempty"`

	// The protocol used for this VPN Link Connection. Possible values are IKEv1 and IKEv2. Defaults to IKEv2.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Should the rate limit be enabled? Defaults to false.
	RatelimitEnabled *bool `json:"ratelimitEnabled,omitempty" tf:"ratelimit_enabled,omitempty"`

	// Routing weight for this VPN Link Connection. Defaults to 0.
	RouteWeight *float64 `json:"routeWeight,omitempty" tf:"route_weight,omitempty"`

	// SharedKey for this VPN Link Connection.
	SharedKey *string `json:"sharedKey,omitempty" tf:"shared_key,omitempty"`

	// The ID of the connected VPN Site Link. Changing this forces a new VPN Gateway Connection to be created.
	VPNSiteLinkID *string `json:"vpnSiteLinkId,omitempty" tf:"vpn_site_link_id,omitempty"`
}

type VPNLinkParameters struct {

	// Should the BGP be enabled? Defaults to false. Changing this forces a new VPN Gateway Connection to be created.
	// +kubebuilder:validation:Optional
	BGPEnabled *bool `json:"bgpEnabled,omitempty" tf:"bgp_enabled,omitempty"`

	// The expected connection bandwidth in MBPS. Defaults to 10.
	// +kubebuilder:validation:Optional
	BandwidthMbps *float64 `json:"bandwidthMbps,omitempty" tf:"bandwidth_mbps,omitempty"`

	// The connection mode of this VPN Link. Possible values are Default, InitiatorOnly and ResponderOnly. Defaults to Default.
	// +kubebuilder:validation:Optional
	ConnectionMode *string `json:"connectionMode,omitempty" tf:"connection_mode,omitempty"`

	// One or more custom_bgp_address blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomBGPAddress []CustomBGPAddressParameters `json:"customBgpAddress,omitempty" tf:"custom_bgp_address,omitempty"`

	// A list of the egress NAT Rule Ids.
	// +kubebuilder:validation:Optional
	EgressNATRuleIds []*string `json:"egressNatRuleIds,omitempty" tf:"egress_nat_rule_ids,omitempty"`

	// A list of the ingress NAT Rule Ids.
	// +kubebuilder:validation:Optional
	IngressNATRuleIds []*string `json:"ingressNatRuleIds,omitempty" tf:"ingress_nat_rule_ids,omitempty"`

	// One or more ipsec_policy blocks as defined above.
	// +kubebuilder:validation:Optional
	IpsecPolicy []VPNLinkIpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// Whether to use local Azure IP to initiate connection? Defaults to false.
	// +kubebuilder:validation:Optional
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled,omitempty"`

	// The name which should be used for this VPN Link Connection.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether to enable policy-based traffic selectors? Defaults to false.
	// +kubebuilder:validation:Optional
	PolicyBasedTrafficSelectorEnabled *bool `json:"policyBasedTrafficSelectorEnabled,omitempty" tf:"policy_based_traffic_selector_enabled,omitempty"`

	// The protocol used for this VPN Link Connection. Possible values are IKEv1 and IKEv2. Defaults to IKEv2.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Should the rate limit be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	RatelimitEnabled *bool `json:"ratelimitEnabled,omitempty" tf:"ratelimit_enabled,omitempty"`

	// Routing weight for this VPN Link Connection. Defaults to 0.
	// +kubebuilder:validation:Optional
	RouteWeight *float64 `json:"routeWeight,omitempty" tf:"route_weight,omitempty"`

	// SharedKey for this VPN Link Connection.
	// +kubebuilder:validation:Optional
	SharedKey *string `json:"sharedKey,omitempty" tf:"shared_key,omitempty"`

	// The ID of the connected VPN Site Link. Changing this forces a new VPN Gateway Connection to be created.
	// +kubebuilder:validation:Required
	VPNSiteLinkID *string `json:"vpnSiteLinkId" tf:"vpn_site_link_id,omitempty"`
}

// VPNGatewayConnectionSpec defines the desired state of VPNGatewayConnection
type VPNGatewayConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayConnectionParameters `json:"forProvider"`
}

// VPNGatewayConnectionStatus defines the observed state of VPNGatewayConnection.
type VPNGatewayConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayConnection is the Schema for the VPNGatewayConnections API. Manages a VPN Gateway Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VPNGatewayConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vpnLink)",message="vpnLink is a required parameter"
	Spec   VPNGatewayConnectionSpec   `json:"spec"`
	Status VPNGatewayConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayConnectionList contains a list of VPNGatewayConnections
type VPNGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGatewayConnection `json:"items"`
}

// Repository type metadata.
var (
	VPNGatewayConnection_Kind             = "VPNGatewayConnection"
	VPNGatewayConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGatewayConnection_Kind}.String()
	VPNGatewayConnection_KindAPIVersion   = VPNGatewayConnection_Kind + "." + CRDGroupVersion.String()
	VPNGatewayConnection_GroupVersionKind = CRDGroupVersion.WithKind(VPNGatewayConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGatewayConnection{}, &VPNGatewayConnectionList{})
}
