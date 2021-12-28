/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-27T12:26:28Z.
 *
 * API version: 0.0.1-37434
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TestcryptCredentialAllOf Definition of the list of properties defined in 'testcrypt.Credential', excluding properties defined in parent classes.
type TestcryptCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Password associated with username, to be stored in the database encrypted.
	Password *string `json:"Password,omitempty"`
	// The name of the user whose credentials are being set.
	Username             *string                 `json:"Username,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestcryptCredentialAllOf TestcryptCredentialAllOf

// NewTestcryptCredentialAllOf instantiates a new TestcryptCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestcryptCredentialAllOf(classId string, objectType string) *TestcryptCredentialAllOf {
	this := TestcryptCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewTestcryptCredentialAllOfWithDefaults instantiates a new TestcryptCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestcryptCredentialAllOfWithDefaults() *TestcryptCredentialAllOf {
	this := TestcryptCredentialAllOf{}
	var classId string = "testcrypt.Credential"
	this.ClassId = classId
	var objectType string = "testcrypt.Credential"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *TestcryptCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TestcryptCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TestcryptCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TestcryptCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *TestcryptCredentialAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *TestcryptCredentialAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *TestcryptCredentialAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *TestcryptCredentialAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *TestcryptCredentialAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *TestcryptCredentialAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TestcryptCredentialAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TestcryptCredentialAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TestcryptCredentialAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TestcryptCredentialAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptCredentialAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TestcryptCredentialAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TestcryptCredentialAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TestcryptCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestcryptCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTestcryptCredentialAllOf := _TestcryptCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varTestcryptCredentialAllOf); err == nil {
		*o = TestcryptCredentialAllOf(varTestcryptCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestcryptCredentialAllOf struct {
	value *TestcryptCredentialAllOf
	isSet bool
}

func (v NullableTestcryptCredentialAllOf) Get() *TestcryptCredentialAllOf {
	return v.value
}

func (v *NullableTestcryptCredentialAllOf) Set(val *TestcryptCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTestcryptCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTestcryptCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestcryptCredentialAllOf(val *TestcryptCredentialAllOf) *NullableTestcryptCredentialAllOf {
	return &NullableTestcryptCredentialAllOf{value: val, isSet: true}
}

func (v NullableTestcryptCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestcryptCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
