/*
Composable Fabric Manager Service OpenAPI

This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ServiceResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceResource{}

// ServiceResource struct for ServiceResource
type ServiceResource struct {
	// A full path to the resource with id as the last component
	Uri string `json:"uri"`
	// The service(s) available for the specified URI
	Methods string `json:"methods"`
	// The description of service(s) offered by the URI
	Description string `json:"description"`
}

// NewServiceResource instantiates a new ServiceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResource(uri string, methods string, description string) *ServiceResource {
	this := ServiceResource{}
	this.Uri = uri
	this.Methods = methods
	this.Description = description
	return &this
}

// NewServiceResourceWithDefaults instantiates a new ServiceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResourceWithDefaults() *ServiceResource {
	this := ServiceResource{}
	return &this
}

// GetUri returns the Uri field value
func (o *ServiceResource) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ServiceResource) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ServiceResource) SetUri(v string) {
	o.Uri = v
}

// GetMethods returns the Methods field value
func (o *ServiceResource) GetMethods() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *ServiceResource) GetMethodsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Methods, true
}

// SetMethods sets field value
func (o *ServiceResource) SetMethods(v string) {
	o.Methods = v
}

// GetDescription returns the Description field value
func (o *ServiceResource) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceResource) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceResource) SetDescription(v string) {
	o.Description = v
}

func (o ServiceResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	toSerialize["methods"] = o.Methods
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableServiceResource struct {
	value *ServiceResource
	isSet bool
}

func (v NullableServiceResource) Get() *ServiceResource {
	return v.value
}

func (v *NullableServiceResource) Set(val *ServiceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResource(val *ServiceResource) *NullableServiceResource {
	return &NullableServiceResource{value: val, isSet: true}
}

func (v NullableServiceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
