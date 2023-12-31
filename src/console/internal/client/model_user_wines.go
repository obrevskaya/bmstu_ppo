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

// checks if the UserWines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserWines{}

// UserWines struct for UserWines
type UserWines struct {
	UserWines []UserWine `json:"userWines"`
}

// NewUserWines instantiates a new UserWines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWines(userWines []UserWine) *UserWines {
	this := UserWines{}
	this.UserWines = userWines
	return &this
}

// NewUserWinesWithDefaults instantiates a new UserWines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWinesWithDefaults() *UserWines {
	this := UserWines{}
	return &this
}

// GetUserWines returns the UserWines field value
func (o *UserWines) GetUserWines() []UserWine {
	if o == nil {
		var ret []UserWine
		return ret
	}

	return o.UserWines
}

// GetUserWinesOk returns a tuple with the UserWines field value
// and a boolean to check if the value has been set.
func (o *UserWines) GetUserWinesOk() ([]UserWine, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserWines, true
}

// SetUserWines sets field value
func (o *UserWines) SetUserWines(v []UserWine) {
	o.UserWines = v
}

func (o UserWines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userWines"] = o.UserWines
	return toSerialize, nil
}

type NullableUserWines struct {
	value *UserWines
	isSet bool
}

func (v NullableUserWines) Get() *UserWines {
	return v.value
}

func (v *NullableUserWines) Set(val *UserWines) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWines) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWines(val *UserWines) *NullableUserWines {
	return &NullableUserWines{value: val, isSet: true}
}

func (v NullableUserWines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
