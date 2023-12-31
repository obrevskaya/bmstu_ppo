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

// checks if the DeleteElemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteElemResponse{}

// DeleteElemResponse struct for DeleteElemResponse
type DeleteElemResponse struct {
	Deleted bool `json:"deleted"`
}

// NewDeleteElemResponse instantiates a new DeleteElemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteElemResponse(deleted bool) *DeleteElemResponse {
	this := DeleteElemResponse{}
	this.Deleted = deleted
	return &this
}

// NewDeleteElemResponseWithDefaults instantiates a new DeleteElemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteElemResponseWithDefaults() *DeleteElemResponse {
	this := DeleteElemResponse{}
	return &this
}

// GetDeleted returns the Deleted field value
func (o *DeleteElemResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteElemResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteElemResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o DeleteElemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteElemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

type NullableDeleteElemResponse struct {
	value *DeleteElemResponse
	isSet bool
}

func (v NullableDeleteElemResponse) Get() *DeleteElemResponse {
	return v.value
}

func (v *NullableDeleteElemResponse) Set(val *DeleteElemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteElemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteElemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteElemResponse(val *DeleteElemResponse) *NullableDeleteElemResponse {
	return &NullableDeleteElemResponse{value: val, isSet: true}
}

func (v NullableDeleteElemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteElemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
