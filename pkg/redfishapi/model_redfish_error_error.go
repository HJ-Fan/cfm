/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// RedfishErrorError - The properties that describe an error from a Redfish service.
type RedfishErrorError struct {

	// An array of messages describing one or more error messages.
	MessageExtendedInfo []MessageV120Message `json:"@Message.ExtendedInfo,omitempty"`

	// A string indicating a specific MessageId from a message registry.
	Code string `json:"code"`

	// A human-readable error message corresponding to the message in a message registry.
	Message string `json:"message"`
}

// AssertRedfishErrorErrorRequired checks if the required fields are not zero-ed
func AssertRedfishErrorErrorRequired(obj RedfishErrorError) error {
	elements := map[string]interface{}{
		"code":    obj.Code,
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.MessageExtendedInfo {
		if err := AssertMessageV120MessageRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRedfishErrorErrorConstraints checks if the values respects the defined constraints
func AssertRedfishErrorErrorConstraints(obj RedfishErrorError) error {
	return nil
}
