/*
API for ppo project

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetByOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetByOrderRequest{}

// GetByOrderRequest struct for GetByOrderRequest
type GetByOrderRequest struct {
	Id string `json:"id"`
}

// NewGetByOrderRequest instantiates a new GetByOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetByOrderRequest(id string) *GetByOrderRequest {
	this := GetByOrderRequest{}
	this.Id = id
	return &this
}

// NewGetByOrderRequestWithDefaults instantiates a new GetByOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetByOrderRequestWithDefaults() *GetByOrderRequest {
	this := GetByOrderRequest{}
	return &this
}

// GetId returns the Id field value
func (o *GetByOrderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetByOrderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetByOrderRequest) SetId(v string) {
	o.Id = v
}

func (o GetByOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetByOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGetByOrderRequest struct {
	value *GetByOrderRequest
	isSet bool
}

func (v NullableGetByOrderRequest) Get() *GetByOrderRequest {
	return v.value
}

func (v *NullableGetByOrderRequest) Set(val *GetByOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetByOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetByOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetByOrderRequest(val *GetByOrderRequest) *NullableGetByOrderRequest {
	return &NullableGetByOrderRequest{value: val, isSet: true}
}

func (v NullableGetByOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetByOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
