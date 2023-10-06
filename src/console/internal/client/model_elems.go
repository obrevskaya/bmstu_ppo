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

// checks if the Elems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Elems{}

// Elems struct for Elems
type Elems struct {
	Elems []Elem `json:"elems"`
}

// NewElems instantiates a new Elems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElems(elems []Elem) *Elems {
	this := Elems{}
	this.Elems = elems
	return &this
}

// NewElemsWithDefaults instantiates a new Elems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElemsWithDefaults() *Elems {
	this := Elems{}
	return &this
}

// GetElems returns the Elems field value
func (o *Elems) GetElems() []Elem {
	if o == nil {
		var ret []Elem
		return ret
	}

	return o.Elems
}

// GetElemsOk returns a tuple with the Elems field value
// and a boolean to check if the value has been set.
func (o *Elems) GetElemsOk() ([]Elem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elems, true
}

// SetElems sets field value
func (o *Elems) SetElems(v []Elem) {
	o.Elems = v
}

func (o Elems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Elems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["elems"] = o.Elems
	return toSerialize, nil
}

type NullableElems struct {
	value *Elems
	isSet bool
}

func (v NullableElems) Get() *Elems {
	return v.value
}

func (v *NullableElems) Set(val *Elems) {
	v.value = val
	v.isSet = true
}

func (v NullableElems) IsSet() bool {
	return v.isSet
}

func (v *NullableElems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElems(val *Elems) *NullableElems {
	return &NullableElems{value: val, isSet: true}
}

func (v NullableElems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}