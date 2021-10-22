/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-10-09T21:18:32Z.
 *
 * API version: 1.0.9-4809
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicEthRxQueueSettings Receive Queue resource settings.
type VnicEthRxQueueSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of queue resources to allocate.
	Count *int64 `json:"Count,omitempty"`
	// The number of descriptors in each queue.
	RingSize             *int64 `json:"RingSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthRxQueueSettings VnicEthRxQueueSettings

// NewVnicEthRxQueueSettings instantiates a new VnicEthRxQueueSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthRxQueueSettings(classId string, objectType string) *VnicEthRxQueueSettings {
	this := VnicEthRxQueueSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var count int64 = 4
	this.Count = &count
	var ringSize int64 = 512
	this.RingSize = &ringSize
	return &this
}

// NewVnicEthRxQueueSettingsWithDefaults instantiates a new VnicEthRxQueueSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthRxQueueSettingsWithDefaults() *VnicEthRxQueueSettings {
	this := VnicEthRxQueueSettings{}
	var classId string = "vnic.EthRxQueueSettings"
	this.ClassId = classId
	var objectType string = "vnic.EthRxQueueSettings"
	this.ObjectType = objectType
	var count int64 = 4
	this.Count = &count
	var ringSize int64 = 512
	this.RingSize = &ringSize
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthRxQueueSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthRxQueueSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthRxQueueSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthRxQueueSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicEthRxQueueSettings) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicEthRxQueueSettings) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicEthRxQueueSettings) SetCount(v int64) {
	o.Count = &v
}

// GetRingSize returns the RingSize field value if set, zero value otherwise.
func (o *VnicEthRxQueueSettings) GetRingSize() int64 {
	if o == nil || o.RingSize == nil {
		var ret int64
		return ret
	}
	return *o.RingSize
}

// GetRingSizeOk returns a tuple with the RingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthRxQueueSettings) GetRingSizeOk() (*int64, bool) {
	if o == nil || o.RingSize == nil {
		return nil, false
	}
	return o.RingSize, true
}

// HasRingSize returns a boolean if a field has been set.
func (o *VnicEthRxQueueSettings) HasRingSize() bool {
	if o != nil && o.RingSize != nil {
		return true
	}

	return false
}

// SetRingSize gets a reference to the given int64 and assigns it to the RingSize field.
func (o *VnicEthRxQueueSettings) SetRingSize(v int64) {
	o.RingSize = &v
}

func (o VnicEthRxQueueSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.RingSize != nil {
		toSerialize["RingSize"] = o.RingSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthRxQueueSettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicEthRxQueueSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of queue resources to allocate.
		Count *int64 `json:"Count,omitempty"`
		// The number of descriptors in each queue.
		RingSize *int64 `json:"RingSize,omitempty"`
	}

	varVnicEthRxQueueSettingsWithoutEmbeddedStruct := VnicEthRxQueueSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicEthRxQueueSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthRxQueueSettings := _VnicEthRxQueueSettings{}
		varVnicEthRxQueueSettings.ClassId = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.ClassId
		varVnicEthRxQueueSettings.ObjectType = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.ObjectType
		varVnicEthRxQueueSettings.Count = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.Count
		varVnicEthRxQueueSettings.RingSize = varVnicEthRxQueueSettingsWithoutEmbeddedStruct.RingSize
		*o = VnicEthRxQueueSettings(varVnicEthRxQueueSettings)
	} else {
		return err
	}

	varVnicEthRxQueueSettings := _VnicEthRxQueueSettings{}

	err = json.Unmarshal(bytes, &varVnicEthRxQueueSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicEthRxQueueSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "RingSize")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicEthRxQueueSettings struct {
	value *VnicEthRxQueueSettings
	isSet bool
}

func (v NullableVnicEthRxQueueSettings) Get() *VnicEthRxQueueSettings {
	return v.value
}

func (v *NullableVnicEthRxQueueSettings) Set(val *VnicEthRxQueueSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthRxQueueSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthRxQueueSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthRxQueueSettings(val *VnicEthRxQueueSettings) *NullableVnicEthRxQueueSettings {
	return &NullableVnicEthRxQueueSettings{value: val, isSet: true}
}

func (v NullableVnicEthRxQueueSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthRxQueueSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
