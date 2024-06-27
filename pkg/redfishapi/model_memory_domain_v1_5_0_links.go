/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// MemoryDomainV150Links - The links to other Resources that are related to this Resource.
type MemoryDomainV150Links struct {

	// An array of links to the CXL logical devices associated with this memory domain.
	CXLLogicalDevices []OdataV4IdRef `json:"CXLLogicalDevices,omitempty"`

	// The number of items in a collection.
	CXLLogicalDevicesodataCount int64 `json:"CXLLogicalDevices@odata.count,omitempty"`

	// An array of links to the fabric adapters that present this memory domain to a fabric.
	FabricAdapters []OdataV4IdRef `json:"FabricAdapters,omitempty"`

	// The number of items in a collection.
	FabricAdaptersodataCount int64 `json:"FabricAdapters@odata.count,omitempty"`

	// An array of links to the media controllers for this memory domain.
	// Deprecated
	MediaControllers []OdataV4IdRef `json:"MediaControllers,omitempty"`

	// The number of items in a collection.
	MediaControllersodataCount int64 `json:"MediaControllers@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe functions representing this memory domain.
	PCIeFunctions []OdataV4IdRef `json:"PCIeFunctions,omitempty"`

	// The number of items in a collection.
	PCIeFunctionsodataCount int64 `json:"PCIeFunctions@odata.count,omitempty"`
}

// AssertMemoryDomainV150LinksRequired checks if the required fields are not zero-ed
func AssertMemoryDomainV150LinksRequired(obj MemoryDomainV150Links) error {
	for _, el := range obj.CXLLogicalDevices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FabricAdapters {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MediaControllers {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PCIeFunctions {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMemoryDomainV150LinksConstraints checks if the values respects the defined constraints
func AssertMemoryDomainV150LinksConstraints(obj MemoryDomainV150Links) error {
	return nil
}
