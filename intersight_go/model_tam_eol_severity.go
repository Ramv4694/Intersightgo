/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-24T09:42:08Z.
 *
 * API version: 0.0.1-37430
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TamEolSeverity Severity level of an end-of-life (EOL) milestone advisory.
type TamEolSeverity struct {
	TamSeverity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Severity level associated with an end-of-life (EOL) milestone advisory. * `info` - The end-of-life (EOL) milestone is at info level. * `critical` - The end-of-life (EOL) milestone is at critical level. It usually hints 'red' in a color-map. * `high` - The end-of-life (EOL) milestone is at high level. * `medium` - The end-of-life (EOL) milestone is at medium level.
	Level                *string `json:"Level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamEolSeverity TamEolSeverity

// NewTamEolSeverity instantiates a new TamEolSeverity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamEolSeverity(classId string, objectType string) *TamEolSeverity {
	this := TamEolSeverity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var level string = "info"
	this.Level = &level
	return &this
}

// NewTamEolSeverityWithDefaults instantiates a new TamEolSeverity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamEolSeverityWithDefaults() *TamEolSeverity {
	this := TamEolSeverity{}
	var classId string = "tam.EolSeverity"
	this.ClassId = classId
	var objectType string = "tam.EolSeverity"
	this.ObjectType = objectType
	var level string = "info"
	this.Level = &level
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamEolSeverity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamEolSeverity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamEolSeverity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamEolSeverity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamEolSeverity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamEolSeverity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *TamEolSeverity) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamEolSeverity) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *TamEolSeverity) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *TamEolSeverity) SetLevel(v string) {
	o.Level = &v
}

func (o TamEolSeverity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamSeverity, errTamSeverity := json.Marshal(o.TamSeverity)
	if errTamSeverity != nil {
		return []byte{}, errTamSeverity
	}
	errTamSeverity = json.Unmarshal([]byte(serializedTamSeverity), &toSerialize)
	if errTamSeverity != nil {
		return []byte{}, errTamSeverity
	}
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

func (o *TamEolSeverity) UnmarshalJSON(bytes []byte) (err error) {
	type TamEolSeverityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Severity level associated with an end-of-life (EOL) milestone advisory. * `info` - The end-of-life (EOL) milestone is at info level. * `critical` - The end-of-life (EOL) milestone is at critical level. It usually hints 'red' in a color-map. * `high` - The end-of-life (EOL) milestone is at high level. * `medium` - The end-of-life (EOL) milestone is at medium level.
		Level *string `json:"Level,omitempty"`
	}

	varTamEolSeverityWithoutEmbeddedStruct := TamEolSeverityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamEolSeverityWithoutEmbeddedStruct)
	if err == nil {
		varTamEolSeverity := _TamEolSeverity{}
		varTamEolSeverity.ClassId = varTamEolSeverityWithoutEmbeddedStruct.ClassId
		varTamEolSeverity.ObjectType = varTamEolSeverityWithoutEmbeddedStruct.ObjectType
		varTamEolSeverity.Level = varTamEolSeverityWithoutEmbeddedStruct.Level
		*o = TamEolSeverity(varTamEolSeverity)
	} else {
		return err
	}

	varTamEolSeverity := _TamEolSeverity{}

	err = json.Unmarshal(bytes, &varTamEolSeverity)
	if err == nil {
		o.TamSeverity = varTamEolSeverity.TamSeverity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")

		// remove fields from embedded structs
		reflectTamSeverity := reflect.ValueOf(o.TamSeverity)
		for i := 0; i < reflectTamSeverity.Type().NumField(); i++ {
			t := reflectTamSeverity.Type().Field(i)

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

type NullableTamEolSeverity struct {
	value *TamEolSeverity
	isSet bool
}

func (v NullableTamEolSeverity) Get() *TamEolSeverity {
	return v.value
}

func (v *NullableTamEolSeverity) Set(val *TamEolSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableTamEolSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableTamEolSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamEolSeverity(val *TamEolSeverity) *NullableTamEolSeverity {
	return &NullableTamEolSeverity{value: val, isSet: true}
}

func (v NullableTamEolSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamEolSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}