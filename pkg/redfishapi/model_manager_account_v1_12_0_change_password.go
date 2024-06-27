/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ManagerAccountV1120ChangePassword - This action changes the account password.
type ManagerAccountV1120ChangePassword struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertManagerAccountV1120ChangePasswordRequired checks if the required fields are not zero-ed
func AssertManagerAccountV1120ChangePasswordRequired(obj ManagerAccountV1120ChangePassword) error {
	return nil
}

// AssertManagerAccountV1120ChangePasswordConstraints checks if the values respects the defined constraints
func AssertManagerAccountV1120ChangePasswordConstraints(obj ManagerAccountV1120ChangePassword) error {
	return nil
}
