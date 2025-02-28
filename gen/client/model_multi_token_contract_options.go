/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MultiTokenContractOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiTokenContractOptions{}

// MultiTokenContractOptions Options for multi-token contract creation
type MultiTokenContractOptions struct {
	// The URI for all token metadata
	Uri string `json:"uri"`
}

type _MultiTokenContractOptions MultiTokenContractOptions

// NewMultiTokenContractOptions instantiates a new MultiTokenContractOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiTokenContractOptions(uri string) *MultiTokenContractOptions {
	this := MultiTokenContractOptions{}
	this.Uri = uri
	return &this
}

// NewMultiTokenContractOptionsWithDefaults instantiates a new MultiTokenContractOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiTokenContractOptionsWithDefaults() *MultiTokenContractOptions {
	this := MultiTokenContractOptions{}
	return &this
}

// GetUri returns the Uri field value
func (o *MultiTokenContractOptions) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *MultiTokenContractOptions) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *MultiTokenContractOptions) SetUri(v string) {
	o.Uri = v
}

func (o MultiTokenContractOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiTokenContractOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *MultiTokenContractOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uri",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMultiTokenContractOptions := _MultiTokenContractOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultiTokenContractOptions)

	if err != nil {
		return err
	}

	*o = MultiTokenContractOptions(varMultiTokenContractOptions)

	return err
}

type NullableMultiTokenContractOptions struct {
	value *MultiTokenContractOptions
	isSet bool
}

func (v NullableMultiTokenContractOptions) Get() *MultiTokenContractOptions {
	return v.value
}

func (v *NullableMultiTokenContractOptions) Set(val *MultiTokenContractOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiTokenContractOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiTokenContractOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiTokenContractOptions(val *MultiTokenContractOptions) *NullableMultiTokenContractOptions {
	return &NullableMultiTokenContractOptions{value: val, isSet: true}
}

func (v NullableMultiTokenContractOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiTokenContractOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


