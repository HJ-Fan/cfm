/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

import (
	"errors"
)

// StorageStorageController - The StorageController schema describes a storage controller and its properties.  A storage controller represents a physical or virtual storage device that produces volumes.
type StorageStorageController struct {

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	Actions StorageV1160StorageControllerActions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this storage controller.
	AssetTag *string `json:"AssetTag,omitempty"`

	CacheSummary StorageV1160CacheSummary `json:"CacheSummary,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	ControllerRates StorageV1160Rates `json:"ControllerRates,omitempty"`

	// The firmware version of this storage controller.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`

	// The durable names for the storage controller.
	Identifiers []ResourceIdentifier `json:"Identifiers,omitempty"`

	Links StorageV1160StorageControllerLinks `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// The manufacturer of this storage controller.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	// Deprecated
	Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`

	// The unique identifier for the member within an array.
	MemberId string `json:"MemberId"`

	// The model number for the storage controller.
	Model *string `json:"Model,omitempty"`

	// The name of the storage controller.
	Name *string `json:"Name,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeInterface PcieDevicePcieInterface `json:"PCIeInterface,omitempty"`

	// The part number for this storage controller.
	PartNumber *string `json:"PartNumber,omitempty"`

	Ports OdataV4IdRef `json:"Ports,omitempty"`

	// The SKU for this storage controller.
	SKU *string `json:"SKU,omitempty"`

	// The serial number for this storage controller.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The maximum speed of the storage controller's device interface.
	SpeedGbps *float32 `json:"SpeedGbps,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The supported set of protocols for communicating with this storage controller.
	SupportedControllerProtocols []ProtocolProtocol `json:"SupportedControllerProtocols,omitempty"`

	// The protocols that the storage controller can use to communicate with attached devices.
	SupportedDeviceProtocols []ProtocolProtocol `json:"SupportedDeviceProtocols,omitempty"`

	// The set of RAID types supported by the storage controller.
	SupportedRAIDTypes []StorageStorageControllerSupportedRaidTypesInner `json:"SupportedRAIDTypes,omitempty"`
}

// AssertStorageStorageControllerRequired checks if the required fields are not zero-ed
func AssertStorageStorageControllerRequired(obj StorageStorageController) error {
	elements := map[string]interface{}{
		"@odata.id": obj.OdataId,
		"MemberId":  obj.MemberId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertStorageV1160StorageControllerActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertStorageV1160CacheSummaryRequired(obj.CacheSummary); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	if err := AssertStorageV1160RatesRequired(obj.ControllerRates); err != nil {
		return err
	}
	for _, el := range obj.Identifiers {
		if err := AssertResourceIdentifierRequired(el); err != nil {
			return err
		}
	}
	if err := AssertStorageV1160StorageControllerLinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPcieDevicePcieInterfaceRequired(obj.PCIeInterface); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Ports); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	for _, el := range obj.SupportedRAIDTypes {
		if err := AssertStorageStorageControllerSupportedRaidTypesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertStorageStorageControllerConstraints checks if the values respects the defined constraints
func AssertStorageStorageControllerConstraints(obj StorageStorageController) error {
	if obj.SpeedGbps != nil && *obj.SpeedGbps < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
