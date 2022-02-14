/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2022-02-14T08:05:27Z.
 *
 * API version: 0.0.1-38392
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TamEolSeverityAllOf Definition of the list of properties defined in 'tam.EolSeverity', excluding properties defined in parent classes.
type TamEolSeverityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Severity level associated with an end-of-life (EOL) milestone advisory. * `info` - The end-of-life (EOL) milestone is at info level. * `critical` - The end-of-life (EOL) milestone is at critical level. It usually hints 'red' in a color-map. * `high` - The end-of-life (EOL) milestone is at high level. * `medium` - The end-of-life (EOL) milestone is at medium level.
	Level                *string `json:"Level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamEolSeverityAllOf TamEolSeverityAllOf

// NewTamEolSeverityAllOf instantiates a new TamEolSeverityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamEolSeverityAllOf(classId string, objectType string) *TamEolSeverityAllOf {
	this := TamEolSeverityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var level string = "info"
	this.Level = &level
	return &this
}

// NewTamEolSeverityAllOfWithDefaults instantiates a new TamEolSeverityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamEolSeverityAllOfWithDefaults() *TamEolSeverityAllOf {
	this := TamEolSeverityAllOf{}
	var classId string = "tam.EolSeverity"
	this.ClassId = classId
	var objectType string = "tam.EolSeverity"
	this.ObjectType = objectType
	var level string = "info"
	this.Level = &level
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamEolSeverityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamEolSeverityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamEolSeverityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamEolSeverityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamEolSeverityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamEolSeverityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TamEolSeverityAllOf) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamEolSeverityAllOf) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TamEolSeverityAllOf) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TamEolSeverityAllOf) SetLevel(v string) {
	o.Level = &v
}

func (o TamEolSeverityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamEolSeverityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamEolSeverityAllOf := _TamEolSeverityAllOf{}

	if err = json.Unmarshal(bytes, &varTamEolSeverityAllOf); err == nil {
		*o = TamEolSeverityAllOf(varTamEolSeverityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamEolSeverityAllOf struct {
	value *TamEolSeverityAllOf
	isSet bool
}

func (v NullableTamEolSeverityAllOf) Get() *TamEolSeverityAllOf {
	return v.value
}

func (v *NullableTamEolSeverityAllOf) Set(val *TamEolSeverityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamEolSeverityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamEolSeverityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamEolSeverityAllOf(val *TamEolSeverityAllOf) *NullableTamEolSeverityAllOf {
	return &NullableTamEolSeverityAllOf{value: val, isSet: true}
}

func (v NullableTamEolSeverityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamEolSeverityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
