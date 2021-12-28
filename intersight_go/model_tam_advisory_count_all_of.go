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

// TamAdvisoryCountAllOf Definition of the list of properties defined in 'tam.AdvisoryCount', excluding properties defined in parent classes.
type TamAdvisoryCountAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of advisories affecting the account.
	AdvisoryCount        *int64                  `json:"AdvisoryCount,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryCountAllOf TamAdvisoryCountAllOf

// NewTamAdvisoryCountAllOf instantiates a new TamAdvisoryCountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryCountAllOf(classId string, objectType string) *TamAdvisoryCountAllOf {
	this := TamAdvisoryCountAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamAdvisoryCountAllOfWithDefaults instantiates a new TamAdvisoryCountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryCountAllOfWithDefaults() *TamAdvisoryCountAllOf {
	this := TamAdvisoryCountAllOf{}
	var classId string = "tam.AdvisoryCount"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryCount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryCountAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryCountAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryCountAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryCountAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdvisoryCount returns the AdvisoryCount field value if set, zero value otherwise.
func (o *TamAdvisoryCountAllOf) GetAdvisoryCount() int64 {
	if o == nil || o.AdvisoryCount == nil {
		var ret int64
		return ret
	}
	return *o.AdvisoryCount
}

// GetAdvisoryCountOk returns a tuple with the AdvisoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetAdvisoryCountOk() (*int64, bool) {
	if o == nil || o.AdvisoryCount == nil {
		return nil, false
	}
	return o.AdvisoryCount, true
}

// HasAdvisoryCount returns a boolean if a field has been set.
func (o *TamAdvisoryCountAllOf) HasAdvisoryCount() bool {
	if o != nil && o.AdvisoryCount != nil {
		return true
	}

	return false
}

// SetAdvisoryCount gets a reference to the given int64 and assigns it to the AdvisoryCount field.
func (o *TamAdvisoryCountAllOf) SetAdvisoryCount(v int64) {
	o.AdvisoryCount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TamAdvisoryCountAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryCountAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TamAdvisoryCountAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TamAdvisoryCountAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TamAdvisoryCountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdvisoryCount != nil {
		toSerialize["AdvisoryCount"] = o.AdvisoryCount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryCountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamAdvisoryCountAllOf := _TamAdvisoryCountAllOf{}

	if err = json.Unmarshal(bytes, &varTamAdvisoryCountAllOf); err == nil {
		*o = TamAdvisoryCountAllOf(varTamAdvisoryCountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdvisoryCount")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamAdvisoryCountAllOf struct {
	value *TamAdvisoryCountAllOf
	isSet bool
}

func (v NullableTamAdvisoryCountAllOf) Get() *TamAdvisoryCountAllOf {
	return v.value
}

func (v *NullableTamAdvisoryCountAllOf) Set(val *TamAdvisoryCountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryCountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryCountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryCountAllOf(val *TamAdvisoryCountAllOf) *NullableTamAdvisoryCountAllOf {
	return &NullableTamAdvisoryCountAllOf{value: val, isSet: true}
}

func (v NullableTamAdvisoryCountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryCountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
