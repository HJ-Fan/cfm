/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// SwitchV192Links - The links to other resources that are related to this resource.
type SwitchV192Links struct {
	Chassis OdataV4IdRef `json:"Chassis,omitempty"`

	// An array of links to the endpoints that connect to this switch.
	Endpoints []OdataV4IdRef `json:"Endpoints,omitempty"`

	// The number of items in a collection.
	EndpointsodataCount int64 `json:"Endpoints@odata.count,omitempty"`

	// An array of links to the managers that manage this switch.
	ManagedBy []OdataV4IdRef `json:"ManagedBy,omitempty"`

	// The number of items in a collection.
	ManagedByodataCount int64 `json:"ManagedBy@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeDevice OdataV4IdRef `json:"PCIeDevice,omitempty"`
}

// AssertSwitchV192LinksRequired checks if the required fields are not zero-ed
func AssertSwitchV192LinksRequired(obj SwitchV192Links) error {
	if err := AssertOdataV4IdRefRequired(obj.Chassis); err != nil {
		return err
	}
	for _, el := range obj.Endpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ManagedBy {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.PCIeDevice); err != nil {
		return err
	}
	return nil
}

// AssertSwitchV192LinksConstraints checks if the values respects the defined constraints
func AssertSwitchV192LinksConstraints(obj SwitchV192Links) error {
	return nil
}
