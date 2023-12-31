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

// checks if the UserWine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserWine{}

// UserWine struct for UserWine
type UserWine struct {
	IdUser string `json:"idUser"`
	IdWine string `json:"idWine"`
}

// NewUserWine instantiates a new UserWine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWine(idUser string, idWine string) *UserWine {
	this := UserWine{}
	this.IdUser = idUser
	this.IdWine = idWine
	return &this
}

// NewUserWineWithDefaults instantiates a new UserWine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWineWithDefaults() *UserWine {
	this := UserWine{}
	return &this
}

// GetIdUser returns the IdUser field value
func (o *UserWine) GetIdUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdUser
}

// GetIdUserOk returns a tuple with the IdUser field value
// and a boolean to check if the value has been set.
func (o *UserWine) GetIdUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdUser, true
}

// SetIdUser sets field value
func (o *UserWine) SetIdUser(v string) {
	o.IdUser = v
}

// GetIdWine returns the IdWine field value
func (o *UserWine) GetIdWine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdWine
}

// GetIdWineOk returns a tuple with the IdWine field value
// and a boolean to check if the value has been set.
func (o *UserWine) GetIdWineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdWine, true
}

// SetIdWine sets field value
func (o *UserWine) SetIdWine(v string) {
	o.IdWine = v
}

func (o UserWine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["idUser"] = o.IdUser
	toSerialize["idWine"] = o.IdWine
	return toSerialize, nil
}

type NullableUserWine struct {
	value *UserWine
	isSet bool
}

func (v NullableUserWine) Get() *UserWine {
	return v.value
}

func (v *NullableUserWine) Set(val *UserWine) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWine) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWine(val *UserWine) *NullableUserWine {
	return &NullableUserWine{value: val, isSet: true}
}

func (v NullableUserWine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
