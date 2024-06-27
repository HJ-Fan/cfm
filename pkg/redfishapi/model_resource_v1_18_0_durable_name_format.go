/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ResourceV1180DurableNameFormat string

// List of ResourceV1180DurableNameFormat
const (
	RESOURCEV1180DURABLENAMEFORMAT_NAA         ResourceV1180DurableNameFormat = "NAA"
	RESOURCEV1180DURABLENAMEFORMAT_I_QN        ResourceV1180DurableNameFormat = "iQN"
	RESOURCEV1180DURABLENAMEFORMAT_FC_WWN      ResourceV1180DurableNameFormat = "FC_WWN"
	RESOURCEV1180DURABLENAMEFORMAT_UUID        ResourceV1180DurableNameFormat = "UUID"
	RESOURCEV1180DURABLENAMEFORMAT_EUI         ResourceV1180DurableNameFormat = "EUI"
	RESOURCEV1180DURABLENAMEFORMAT_NQN         ResourceV1180DurableNameFormat = "NQN"
	RESOURCEV1180DURABLENAMEFORMAT_NSID        ResourceV1180DurableNameFormat = "NSID"
	RESOURCEV1180DURABLENAMEFORMAT_NGUID       ResourceV1180DurableNameFormat = "NGUID"
	RESOURCEV1180DURABLENAMEFORMAT_MAC_ADDRESS ResourceV1180DurableNameFormat = "MACAddress"
	RESOURCEV1180DURABLENAMEFORMAT_GCXLID      ResourceV1180DurableNameFormat = "GCXLID"
)

// AssertResourceV1180DurableNameFormatRequired checks if the required fields are not zero-ed
func AssertResourceV1180DurableNameFormatRequired(obj ResourceV1180DurableNameFormat) error {
	return nil
}

// AssertResourceV1180DurableNameFormatConstraints checks if the values respects the defined constraints
func AssertResourceV1180DurableNameFormatConstraints(obj ResourceV1180DurableNameFormat) error {
	return nil
}
