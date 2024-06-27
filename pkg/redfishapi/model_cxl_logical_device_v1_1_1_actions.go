/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// CxlLogicalDeviceV111Actions - The available actions for this resource.
type CxlLogicalDeviceV111Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCxlLogicalDeviceV111ActionsRequired checks if the required fields are not zero-ed
func AssertCxlLogicalDeviceV111ActionsRequired(obj CxlLogicalDeviceV111Actions) error {
	return nil
}

// AssertCxlLogicalDeviceV111ActionsConstraints checks if the values respects the defined constraints
func AssertCxlLogicalDeviceV111ActionsConstraints(obj CxlLogicalDeviceV111Actions) error {
	return nil
}
