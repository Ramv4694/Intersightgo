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

// TestcryptShadowCredentialAllOf Definition of the list of properties defined in 'testcrypt.ShadowCredential', excluding properties defined in parent classes.
type TestcryptShadowCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Password associated with username.
	Password *string `json:"Password,omitempty"`
	// The username of user, whose password marked as 'secure'.
	Username             *string                 `json:"Username,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestcryptShadowCredentialAllOf TestcryptShadowCredentialAllOf

// NewTestcryptShadowCredentialAllOf instantiates a new TestcryptShadowCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestcryptShadowCredentialAllOf(classId string, objectType string) *TestcryptShadowCredentialAllOf {
	this := TestcryptShadowCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTestcryptShadowCredentialAllOfWithDefaults instantiates a new TestcryptShadowCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestcryptShadowCredentialAllOfWithDefaults() *TestcryptShadowCredentialAllOf {
	this := TestcryptShadowCredentialAllOf{}
	var classId string = "testcrypt.ShadowCredential"
	this.ClassId = classId
	var objectType string = "testcrypt.ShadowCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TestcryptShadowCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TestcryptShadowCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TestcryptShadowCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TestcryptShadowCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TestcryptShadowCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TestcryptShadowCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *TestcryptShadowCredentialAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptShadowCredentialAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *TestcryptShadowCredentialAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *TestcryptShadowCredentialAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TestcryptShadowCredentialAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptShadowCredentialAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TestcryptShadowCredentialAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TestcryptShadowCredentialAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TestcryptShadowCredentialAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptShadowCredentialAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TestcryptShadowCredentialAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TestcryptShadowCredentialAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TestcryptShadowCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
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

func (o *TestcryptShadowCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTestcryptShadowCredentialAllOf := _TestcryptShadowCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varTestcryptShadowCredentialAllOf); err == nil {
		*o = TestcryptShadowCredentialAllOf(varTestcryptShadowCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestcryptShadowCredentialAllOf struct {
	value *TestcryptShadowCredentialAllOf
	isSet bool
}

func (v NullableTestcryptShadowCredentialAllOf) Get() *TestcryptShadowCredentialAllOf {
	return v.value
}

func (v *NullableTestcryptShadowCredentialAllOf) Set(val *TestcryptShadowCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTestcryptShadowCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTestcryptShadowCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestcryptShadowCredentialAllOf(val *TestcryptShadowCredentialAllOf) *NullableTestcryptShadowCredentialAllOf {
	return &NullableTestcryptShadowCredentialAllOf{value: val, isSet: true}
}

func (v NullableTestcryptShadowCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestcryptShadowCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
