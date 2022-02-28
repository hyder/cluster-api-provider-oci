// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// UpdateVirtualCircuitDetails The representation of UpdateVirtualCircuitDetails
type UpdateVirtualCircuitDetails struct {

	// The provisioned data rate of the connection. To get a list of the
	// available bandwidth levels (that is, shapes), see
	// ListFastConnectProviderVirtualCircuitBandwidthShapes.
	// To be updated only by the customer who owns the virtual circuit.
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName"`

	// An array of mappings, each containing properties for a cross-connect or
	// cross-connect group associated with this virtual circuit.
	// The customer and provider can update different properties in the mapping
	// depending on the situation. See the description of the
	// CrossConnectMapping.
	CrossConnectMappings []CrossConnectMapping `mandatory:"false" json:"crossConnectMappings"`

	// The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit.
	// Policies available are: `ORACLE_SERVICE_NETWORK`, `REGIONAL`, `MARKET_LEVEL`, and `GLOBAL`.
	// See Route Filtering (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/routingonprem.htm#route_filtering) for details.
	// By default, routing information is shared for all routes in the same market.
	RoutingPolicy []UpdateVirtualCircuitDetailsRoutingPolicyEnum `mandatory:"false" json:"routingPolicy,omitempty"`

	// Deprecated. Instead use `customerAsn`.
	// If you specify values for both, the request will be rejected.
	CustomerBgpAsn *int `mandatory:"false" json:"customerBgpAsn"`

	// The BGP ASN of the network at the other end of the BGP
	// session from Oracle.
	// If the BGP session is from the customer's edge router to Oracle, the
	// required value is the customer's ASN, and it can be updated only
	// by the customer.
	// If the BGP session is from the provider's edge router to Oracle, the
	// required value is the provider's ASN, and it can be updated only
	// by the provider.
	// Can be a 2-byte or 4-byte ASN. Uses "asplain" format.
	CustomerAsn *int64 `mandatory:"false" json:"customerAsn"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Drg
	// that this private virtual circuit uses.
	// To be updated only by the customer who owns the virtual circuit.
	GatewayId *string `mandatory:"false" json:"gatewayId"`

	// The provider's state in relation to this virtual circuit. Relevant only
	// if the customer is using FastConnect via a provider. ACTIVE
	// means the provider has provisioned the virtual circuit from their
	// end. INACTIVE means the provider has not yet provisioned the virtual
	// circuit, or has de-provisioned it.
	// To be updated only by the provider.
	ProviderState UpdateVirtualCircuitDetailsProviderStateEnum `mandatory:"false" json:"providerState,omitempty"`

	// The service key name offered by the provider (if the customer is connecting via a provider).
	ProviderServiceKeyName *string `mandatory:"false" json:"providerServiceKeyName"`

	// Provider-supplied reference information about this virtual circuit.
	// Relevant only if the customer is using FastConnect via a provider.
	// To be updated only by the provider.
	ReferenceComment *string `mandatory:"false" json:"referenceComment"`

	// The layer 3 IP MTU to use on this virtual circuit.
	IpMtu VirtualCircuitIpMtuEnum `mandatory:"false" json:"ipMtu,omitempty"`
}

func (m UpdateVirtualCircuitDetails) String() string {
	return common.PointerString(m)
}

// UpdateVirtualCircuitDetailsRoutingPolicyEnum Enum with underlying type: string
type UpdateVirtualCircuitDetailsRoutingPolicyEnum string

// Set of constants representing the allowable values for UpdateVirtualCircuitDetailsRoutingPolicyEnum
const (
	UpdateVirtualCircuitDetailsRoutingPolicyOracleServiceNetwork UpdateVirtualCircuitDetailsRoutingPolicyEnum = "ORACLE_SERVICE_NETWORK"
	UpdateVirtualCircuitDetailsRoutingPolicyRegional             UpdateVirtualCircuitDetailsRoutingPolicyEnum = "REGIONAL"
	UpdateVirtualCircuitDetailsRoutingPolicyMarketLevel          UpdateVirtualCircuitDetailsRoutingPolicyEnum = "MARKET_LEVEL"
	UpdateVirtualCircuitDetailsRoutingPolicyGlobal               UpdateVirtualCircuitDetailsRoutingPolicyEnum = "GLOBAL"
)

var mappingUpdateVirtualCircuitDetailsRoutingPolicy = map[string]UpdateVirtualCircuitDetailsRoutingPolicyEnum{
	"ORACLE_SERVICE_NETWORK": UpdateVirtualCircuitDetailsRoutingPolicyOracleServiceNetwork,
	"REGIONAL":               UpdateVirtualCircuitDetailsRoutingPolicyRegional,
	"MARKET_LEVEL":           UpdateVirtualCircuitDetailsRoutingPolicyMarketLevel,
	"GLOBAL":                 UpdateVirtualCircuitDetailsRoutingPolicyGlobal,
}

// GetUpdateVirtualCircuitDetailsRoutingPolicyEnumValues Enumerates the set of values for UpdateVirtualCircuitDetailsRoutingPolicyEnum
func GetUpdateVirtualCircuitDetailsRoutingPolicyEnumValues() []UpdateVirtualCircuitDetailsRoutingPolicyEnum {
	values := make([]UpdateVirtualCircuitDetailsRoutingPolicyEnum, 0)
	for _, v := range mappingUpdateVirtualCircuitDetailsRoutingPolicy {
		values = append(values, v)
	}
	return values
}

// UpdateVirtualCircuitDetailsProviderStateEnum Enum with underlying type: string
type UpdateVirtualCircuitDetailsProviderStateEnum string

// Set of constants representing the allowable values for UpdateVirtualCircuitDetailsProviderStateEnum
const (
	UpdateVirtualCircuitDetailsProviderStateActive   UpdateVirtualCircuitDetailsProviderStateEnum = "ACTIVE"
	UpdateVirtualCircuitDetailsProviderStateInactive UpdateVirtualCircuitDetailsProviderStateEnum = "INACTIVE"
)

var mappingUpdateVirtualCircuitDetailsProviderState = map[string]UpdateVirtualCircuitDetailsProviderStateEnum{
	"ACTIVE":   UpdateVirtualCircuitDetailsProviderStateActive,
	"INACTIVE": UpdateVirtualCircuitDetailsProviderStateInactive,
}

// GetUpdateVirtualCircuitDetailsProviderStateEnumValues Enumerates the set of values for UpdateVirtualCircuitDetailsProviderStateEnum
func GetUpdateVirtualCircuitDetailsProviderStateEnumValues() []UpdateVirtualCircuitDetailsProviderStateEnum {
	values := make([]UpdateVirtualCircuitDetailsProviderStateEnum, 0)
	for _, v := range mappingUpdateVirtualCircuitDetailsProviderState {
		values = append(values, v)
	}
	return values
}
