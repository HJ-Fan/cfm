/*
 * Composable Fabric Manager Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

type Credentials struct {

	// The username credentials used to communicate with another service
	Username string `json:"username"`

	// The password credentials used to communicate with another service
	Password string `json:"password"`

	// The IP Address in dot notation of the service
	IpAddress string `json:"ipAddress"`

	Port int32 `json:"port"`

	// Set to true if an insecure connection should be used
	Insecure bool `json:"insecure,omitempty"`

	// Examples of http vs https
	Protocol string `json:"protocol,omitempty"`

	// Examples of backend
	Backend string `json:"backend,omitempty"`

	// A user-defined string to uniquely identify an individual endpoint device.
	CustomId string `json:"customId,omitempty"`
}

// AssertCredentialsRequired checks if the required fields are not zero-ed
func AssertCredentialsRequired(obj Credentials) error {
	elements := map[string]interface{}{
		"username":  obj.Username,
		"password":  obj.Password,
		"ipAddress": obj.IpAddress,
		"port":      obj.Port,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCredentialsConstraints checks if the values respects the defined constraints
func AssertCredentialsConstraints(obj Credentials) error {
	if obj.Port < 80 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Port > 65535 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
