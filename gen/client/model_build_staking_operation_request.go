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

// checks if the BuildStakingOperationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildStakingOperationRequest{}

// BuildStakingOperationRequest struct for BuildStakingOperationRequest
type BuildStakingOperationRequest struct {
	// The ID of the blockchain network
	NetworkId string `json:"network_id"`
	// The ID of the asset being staked
	AssetId string `json:"asset_id"`
	// The onchain address from which the staking transaction originates and is responsible for signing the transaction.
	AddressId string `json:"address_id"`
	// The type of staking operation
	Action string `json:"action"`
	// Additional options for the staking operation.
	Options map[string]string `json:"options"`
}

type _BuildStakingOperationRequest BuildStakingOperationRequest

// NewBuildStakingOperationRequest instantiates a new BuildStakingOperationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildStakingOperationRequest(networkId string, assetId string, addressId string, action string, options map[string]string) *BuildStakingOperationRequest {
	this := BuildStakingOperationRequest{}
	this.NetworkId = networkId
	this.AssetId = assetId
	this.AddressId = addressId
	this.Action = action
	this.Options = options
	return &this
}

// NewBuildStakingOperationRequestWithDefaults instantiates a new BuildStakingOperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildStakingOperationRequestWithDefaults() *BuildStakingOperationRequest {
	this := BuildStakingOperationRequest{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *BuildStakingOperationRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *BuildStakingOperationRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *BuildStakingOperationRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetAssetId returns the AssetId field value
func (o *BuildStakingOperationRequest) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *BuildStakingOperationRequest) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *BuildStakingOperationRequest) SetAssetId(v string) {
	o.AssetId = v
}

// GetAddressId returns the AddressId field value
func (o *BuildStakingOperationRequest) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *BuildStakingOperationRequest) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *BuildStakingOperationRequest) SetAddressId(v string) {
	o.AddressId = v
}

// GetAction returns the Action field value
func (o *BuildStakingOperationRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BuildStakingOperationRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BuildStakingOperationRequest) SetAction(v string) {
	o.Action = v
}

// GetOptions returns the Options field value
func (o *BuildStakingOperationRequest) GetOptions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *BuildStakingOperationRequest) GetOptionsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *BuildStakingOperationRequest) SetOptions(v map[string]string) {
	o.Options = v
}

func (o BuildStakingOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildStakingOperationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["network_id"] = o.NetworkId
	toSerialize["asset_id"] = o.AssetId
	toSerialize["address_id"] = o.AddressId
	toSerialize["action"] = o.Action
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

func (o *BuildStakingOperationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network_id",
		"asset_id",
		"address_id",
		"action",
		"options",
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

	varBuildStakingOperationRequest := _BuildStakingOperationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildStakingOperationRequest)

	if err != nil {
		return err
	}

	*o = BuildStakingOperationRequest(varBuildStakingOperationRequest)

	return err
}

type NullableBuildStakingOperationRequest struct {
	value *BuildStakingOperationRequest
	isSet bool
}

func (v NullableBuildStakingOperationRequest) Get() *BuildStakingOperationRequest {
	return v.value
}

func (v *NullableBuildStakingOperationRequest) Set(val *BuildStakingOperationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildStakingOperationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildStakingOperationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildStakingOperationRequest(val *BuildStakingOperationRequest) *NullableBuildStakingOperationRequest {
	return &NullableBuildStakingOperationRequest{value: val, isSet: true}
}

func (v NullableBuildStakingOperationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildStakingOperationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


