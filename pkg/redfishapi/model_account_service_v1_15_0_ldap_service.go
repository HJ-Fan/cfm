/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// AccountServiceV1150LdapService - The settings required to parse a generic LDAP service.
type AccountServiceV1150LdapService struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	SearchSettings AccountServiceV1150LdapSearchSettings `json:"SearchSettings,omitempty"`
}

// AssertAccountServiceV1150LdapServiceRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150LdapServiceRequired(obj AccountServiceV1150LdapService) error {
	if err := AssertAccountServiceV1150LdapSearchSettingsRequired(obj.SearchSettings); err != nil {
		return err
	}
	return nil
}

// AssertAccountServiceV1150LdapServiceConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150LdapServiceConstraints(obj AccountServiceV1150LdapService) error {
	return nil
}
