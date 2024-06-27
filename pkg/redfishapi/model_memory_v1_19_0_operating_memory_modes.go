/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type MemoryV1190OperatingMemoryModes string

// List of MemoryV1190OperatingMemoryModes
const (
	MEMORYV1190OPERATINGMEMORYMODES_VOLATILE MemoryV1190OperatingMemoryModes = "Volatile"
	MEMORYV1190OPERATINGMEMORYMODES_PMEM     MemoryV1190OperatingMemoryModes = "PMEM"
	MEMORYV1190OPERATINGMEMORYMODES_BLOCK    MemoryV1190OperatingMemoryModes = "Block"
)

// AssertMemoryV1190OperatingMemoryModesRequired checks if the required fields are not zero-ed
func AssertMemoryV1190OperatingMemoryModesRequired(obj MemoryV1190OperatingMemoryModes) error {
	return nil
}

// AssertMemoryV1190OperatingMemoryModesConstraints checks if the values respects the defined constraints
func AssertMemoryV1190OperatingMemoryModesConstraints(obj MemoryV1190OperatingMemoryModes) error {
	return nil
}
