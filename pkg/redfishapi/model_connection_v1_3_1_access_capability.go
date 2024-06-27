/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ConnectionV131AccessCapability string

// List of ConnectionV131AccessCapability
const (
	CONNECTIONV131ACCESSCAPABILITY_READ  ConnectionV131AccessCapability = "Read"
	CONNECTIONV131ACCESSCAPABILITY_WRITE ConnectionV131AccessCapability = "Write"
)

// AssertConnectionV131AccessCapabilityRequired checks if the required fields are not zero-ed
func AssertConnectionV131AccessCapabilityRequired(obj ConnectionV131AccessCapability) error {
	return nil
}

// AssertConnectionV131AccessCapabilityConstraints checks if the values respects the defined constraints
func AssertConnectionV131AccessCapabilityConstraints(obj ConnectionV131AccessCapability) error {
	return nil
}
