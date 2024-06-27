/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type AccountServiceV1150BasicAuthState string

// List of AccountServiceV1150BasicAuthState
const (
	ACCOUNTSERVICEV1150BASICAUTHSTATE_ENABLED      AccountServiceV1150BasicAuthState = "Enabled"
	ACCOUNTSERVICEV1150BASICAUTHSTATE_UNADVERTISED AccountServiceV1150BasicAuthState = "Unadvertised"
	ACCOUNTSERVICEV1150BASICAUTHSTATE_DISABLED     AccountServiceV1150BasicAuthState = "Disabled"
)

// AssertAccountServiceV1150BasicAuthStateRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150BasicAuthStateRequired(obj AccountServiceV1150BasicAuthState) error {
	return nil
}

// AssertAccountServiceV1150BasicAuthStateConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150BasicAuthStateConstraints(obj AccountServiceV1150BasicAuthState) error {
	return nil
}
