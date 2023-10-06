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

// checks if the PayBillResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayBillResponse{}

// PayBillResponse struct for PayBillResponse
type PayBillResponse struct {
	Payed bool `json:"payed"`
}

// NewPayBillResponse instantiates a new PayBillResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayBillResponse(payed bool) *PayBillResponse {
	this := PayBillResponse{}
	this.Payed = payed
	return &this
}

// NewPayBillResponseWithDefaults instantiates a new PayBillResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayBillResponseWithDefaults() *PayBillResponse {
	this := PayBillResponse{}
	return &this
}

// GetPayed returns the Payed field value
func (o *PayBillResponse) GetPayed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Payed
}

// GetPayedOk returns a tuple with the Payed field value
// and a boolean to check if the value has been set.
func (o *PayBillResponse) GetPayedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payed, true
}

// SetPayed sets field value
func (o *PayBillResponse) SetPayed(v bool) {
	o.Payed = v
}

func (o PayBillResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayBillResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payed"] = o.Payed
	return toSerialize, nil
}

type NullablePayBillResponse struct {
	value *PayBillResponse
	isSet bool
}

func (v NullablePayBillResponse) Get() *PayBillResponse {
	return v.value
}

func (v *NullablePayBillResponse) Set(val *PayBillResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayBillResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayBillResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayBillResponse(val *PayBillResponse) *NullablePayBillResponse {
	return &NullablePayBillResponse{value: val, isSet: true}
}

func (v NullablePayBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayBillResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
