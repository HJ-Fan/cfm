/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// FabricV131Fabric - The Fabric schema represents a simple fabric consisting of one or more switches, zero or more endpoints, and zero or more zones.
type FabricV131Fabric struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions FabricV131Actions `json:"Actions,omitempty"`

	AddressPools OdataV4IdRef `json:"AddressPools,omitempty"`

	Connections OdataV4IdRef `json:"Connections,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	EndpointGroups OdataV4IdRef `json:"EndpointGroups,omitempty"`

	Endpoints OdataV4IdRef `json:"Endpoints,omitempty"`

	FabricType ProtocolProtocol `json:"FabricType,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links FabricV131Links `json:"Links,omitempty"`

	// The maximum number of zones the switch can currently configure.
	MaxZones *int64 `json:"MaxZones,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	Switches OdataV4IdRef `json:"Switches,omitempty"`

	UUID string `json:"UUID,omitempty"`

	Zones OdataV4IdRef `json:"Zones,omitempty"`
}

// AssertFabricV131FabricRequired checks if the required fields are not zero-ed
func AssertFabricV131FabricRequired(obj FabricV131Fabric) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertFabricV131ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.AddressPools); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Connections); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.EndpointGroups); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Endpoints); err != nil {
		return err
	}
	if err := AssertFabricV131LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Switches); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Zones); err != nil {
		return err
	}
	return nil
}

// AssertFabricV131FabricConstraints checks if the values respects the defined constraints
func AssertFabricV131FabricConstraints(obj FabricV131Fabric) error {
	return nil
}
