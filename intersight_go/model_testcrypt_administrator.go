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
	"reflect"
	"strings"
)

// TestcryptAdministrator Mo with complextype with secure properties.
type TestcryptAdministrator struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	Admin                NullableTestcryptUser   `json:"Admin,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestcryptAdministrator TestcryptAdministrator

// NewTestcryptAdministrator instantiates a new TestcryptAdministrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestcryptAdministrator(classId string, objectType string) *TestcryptAdministrator {
	this := TestcryptAdministrator{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTestcryptAdministratorWithDefaults instantiates a new TestcryptAdministrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestcryptAdministratorWithDefaults() *TestcryptAdministrator {
	this := TestcryptAdministrator{}
	var classId string = "testcrypt.Administrator"
	this.ClassId = classId
	var objectType string = "testcrypt.Administrator"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TestcryptAdministrator) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TestcryptAdministrator) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TestcryptAdministrator) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TestcryptAdministrator) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TestcryptAdministrator) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TestcryptAdministrator) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdmin returns the Admin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestcryptAdministrator) GetAdmin() TestcryptUser {
	if o == nil || o.Admin.Get() == nil {
		var ret TestcryptUser
		return ret
	}
	return *o.Admin.Get()
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestcryptAdministrator) GetAdminOk() (*TestcryptUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Admin.Get(), o.Admin.IsSet()
}

// HasAdmin returns a boolean if a field has been set.
func (o *TestcryptAdministrator) HasAdmin() bool {
	if o != nil && o.Admin.IsSet() {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given NullableTestcryptUser and assigns it to the Admin field.
func (o *TestcryptAdministrator) SetAdmin(v TestcryptUser) {
	o.Admin.Set(&v)
}

// SetAdminNil sets the value for Admin to be an explicit nil
func (o *TestcryptAdministrator) SetAdminNil() {
	o.Admin.Set(nil)
}

// UnsetAdmin ensures that no value is present for Admin, not even an explicit nil
func (o *TestcryptAdministrator) UnsetAdmin() {
	o.Admin.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TestcryptAdministrator) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestcryptAdministrator) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TestcryptAdministrator) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TestcryptAdministrator) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TestcryptAdministrator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Admin.IsSet() {
		toSerialize["Admin"] = o.Admin.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestcryptAdministrator) UnmarshalJSON(bytes []byte) (err error) {
	type TestcryptAdministratorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		Admin      NullableTestcryptUser   `json:"Admin,omitempty"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varTestcryptAdministratorWithoutEmbeddedStruct := TestcryptAdministratorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTestcryptAdministratorWithoutEmbeddedStruct)
	if err == nil {
		varTestcryptAdministrator := _TestcryptAdministrator{}
		varTestcryptAdministrator.ClassId = varTestcryptAdministratorWithoutEmbeddedStruct.ClassId
		varTestcryptAdministrator.ObjectType = varTestcryptAdministratorWithoutEmbeddedStruct.ObjectType
		varTestcryptAdministrator.Admin = varTestcryptAdministratorWithoutEmbeddedStruct.Admin
		varTestcryptAdministrator.Account = varTestcryptAdministratorWithoutEmbeddedStruct.Account
		*o = TestcryptAdministrator(varTestcryptAdministrator)
	} else {
		return err
	}

	varTestcryptAdministrator := _TestcryptAdministrator{}

	err = json.Unmarshal(bytes, &varTestcryptAdministrator)
	if err == nil {
		o.MoBaseMo = varTestcryptAdministrator.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Admin")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestcryptAdministrator struct {
	value *TestcryptAdministrator
	isSet bool
}

func (v NullableTestcryptAdministrator) Get() *TestcryptAdministrator {
	return v.value
}

func (v *NullableTestcryptAdministrator) Set(val *TestcryptAdministrator) {
	v.value = val
	v.isSet = true
}

func (v NullableTestcryptAdministrator) IsSet() bool {
	return v.isSet
}

func (v *NullableTestcryptAdministrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestcryptAdministrator(val *TestcryptAdministrator) *NullableTestcryptAdministrator {
	return &NullableTestcryptAdministrator{value: val, isSet: true}
}

func (v NullableTestcryptAdministrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestcryptAdministrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
