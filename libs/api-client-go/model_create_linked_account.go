/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateLinkedAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLinkedAccount{}

// CreateLinkedAccount struct for CreateLinkedAccount
type CreateLinkedAccount struct {
	// The authentication provider of the secondary account
	Provider string `json:"provider"`
	// The user ID of the secondary account
	UserId string `json:"userId"`
}

type _CreateLinkedAccount CreateLinkedAccount

// NewCreateLinkedAccount instantiates a new CreateLinkedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkedAccount(provider string, userId string) *CreateLinkedAccount {
	this := CreateLinkedAccount{}
	this.Provider = provider
	this.UserId = userId
	return &this
}

// NewCreateLinkedAccountWithDefaults instantiates a new CreateLinkedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkedAccountWithDefaults() *CreateLinkedAccount {
	this := CreateLinkedAccount{}
	return &this
}

// GetProvider returns the Provider field value
func (o *CreateLinkedAccount) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateLinkedAccount) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateLinkedAccount) SetProvider(v string) {
	o.Provider = v
}

// GetUserId returns the UserId field value
func (o *CreateLinkedAccount) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateLinkedAccount) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateLinkedAccount) SetUserId(v string) {
	o.UserId = v
}

func (o CreateLinkedAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLinkedAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *CreateLinkedAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"userId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateLinkedAccount := _CreateLinkedAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLinkedAccount)

	if err != nil {
		return err
	}

	*o = CreateLinkedAccount(varCreateLinkedAccount)

	return err
}

type NullableCreateLinkedAccount struct {
	value *CreateLinkedAccount
	isSet bool
}

func (v NullableCreateLinkedAccount) Get() *CreateLinkedAccount {
	return v.value
}

func (v *NullableCreateLinkedAccount) Set(val *CreateLinkedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkedAccount(val *CreateLinkedAccount) *NullableCreateLinkedAccount {
	return &NullableCreateLinkedAccount{value: val, isSet: true}
}

func (v NullableCreateLinkedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLinkedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
