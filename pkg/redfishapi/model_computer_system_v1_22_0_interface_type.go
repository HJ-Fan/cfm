/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type ComputerSystemV1220InterfaceType string

// List of ComputerSystemV1220InterfaceType
const (
	COMPUTERSYSTEMV1220INTERFACETYPE_TPM1_2 ComputerSystemV1220InterfaceType = "TPM1_2"
	COMPUTERSYSTEMV1220INTERFACETYPE_TPM2_0 ComputerSystemV1220InterfaceType = "TPM2_0"
	COMPUTERSYSTEMV1220INTERFACETYPE_TCM1_0 ComputerSystemV1220InterfaceType = "TCM1_0"
)

// AssertComputerSystemV1220InterfaceTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220InterfaceTypeRequired(obj ComputerSystemV1220InterfaceType) error {
	return nil
}

// AssertComputerSystemV1220InterfaceTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220InterfaceTypeConstraints(obj ComputerSystemV1220InterfaceType) error {
	return nil
}
