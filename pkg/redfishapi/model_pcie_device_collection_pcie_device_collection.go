/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// PcieDeviceCollectionPcieDeviceCollection - The collection of `PCIeDevice` resource instances.
type PcieDeviceCollectionPcieDeviceCollection struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []OdataV4IdRef `json:"Members"`

	// The number of items in a collection.
	MembersodataCount int64 `json:"Members@odata.count"`

	// The URI to the resource containing the next set of partial members.
	MembersodataNextLink string `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertPcieDeviceCollectionPcieDeviceCollectionRequired checks if the required fields are not zero-ed
func AssertPcieDeviceCollectionPcieDeviceCollectionRequired(obj PcieDeviceCollectionPcieDeviceCollection) error {
	elements := map[string]interface{}{
		"@odata.id":           obj.OdataId,
		"@odata.type":         obj.OdataType,
		"Members":             obj.Members,
		"Members@odata.count": obj.MembersodataCount,
		"Name":                obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Members {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPcieDeviceCollectionPcieDeviceCollectionConstraints checks if the values respects the defined constraints
func AssertPcieDeviceCollectionPcieDeviceCollectionConstraints(obj PcieDeviceCollectionPcieDeviceCollection) error {
	return nil
}
