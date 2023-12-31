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

// checks if the PlaceOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaceOrderResponse{}

// PlaceOrderResponse struct for PlaceOrderResponse
type PlaceOrderResponse struct {
	Placed bool `json:"placed"`
}

// NewPlaceOrderResponse instantiates a new PlaceOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceOrderResponse(placed bool) *PlaceOrderResponse {
	this := PlaceOrderResponse{}
	this.Placed = placed
	return &this
}

// NewPlaceOrderResponseWithDefaults instantiates a new PlaceOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceOrderResponseWithDefaults() *PlaceOrderResponse {
	this := PlaceOrderResponse{}
	return &this
}

// GetPlaced returns the Placed field value
func (o *PlaceOrderResponse) GetPlaced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Placed
}

// GetPlacedOk returns a tuple with the Placed field value
// and a boolean to check if the value has been set.
func (o *PlaceOrderResponse) GetPlacedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placed, true
}

// SetPlaced sets field value
func (o *PlaceOrderResponse) SetPlaced(v bool) {
	o.Placed = v
}

func (o PlaceOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["placed"] = o.Placed
	return toSerialize, nil
}

type NullablePlaceOrderResponse struct {
	value *PlaceOrderResponse
	isSet bool
}

func (v NullablePlaceOrderResponse) Get() *PlaceOrderResponse {
	return v.value
}

func (v *NullablePlaceOrderResponse) Set(val *PlaceOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceOrderResponse(val *PlaceOrderResponse) *NullablePlaceOrderResponse {
	return &NullablePlaceOrderResponse{value: val, isSet: true}
}

func (v NullablePlaceOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
