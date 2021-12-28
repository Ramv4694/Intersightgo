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

// AdapterEthSettings Global Ethernet settings for this adapter.
type AdapterEthSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of LLDP protocol on the adapter interfaces.
	LldpEnabled          *bool `json:"LldpEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterEthSettings AdapterEthSettings

// NewAdapterEthSettings instantiates a new AdapterEthSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterEthSettings(classId string, objectType string) *AdapterEthSettings {
	this := AdapterEthSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lldpEnabled bool = true
	this.LldpEnabled = &lldpEnabled
	return &this
}

// NewAdapterEthSettingsWithDefaults instantiates a new AdapterEthSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterEthSettingsWithDefaults() *AdapterEthSettings {
	this := AdapterEthSettings{}
	var classId string = "adapter.EthSettings"
	this.ClassId = classId
	var objectType string = "adapter.EthSettings"
	this.ObjectType = objectType
	var lldpEnabled bool = true
	this.LldpEnabled = &lldpEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterEthSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterEthSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterEthSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterEthSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterEthSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterEthSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLldpEnabled returns the LldpEnabled field value if set, zero value otherwise.
func (o *AdapterEthSettings) GetLldpEnabled() bool {
	if o == nil || o.LldpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LldpEnabled
}

// GetLldpEnabledOk returns a tuple with the LldpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterEthSettings) GetLldpEnabledOk() (*bool, bool) {
	if o == nil || o.LldpEnabled == nil {
		return nil, false
	}
	return o.LldpEnabled, true
}

// HasLldpEnabled returns a boolean if a field has been set.
func (o *AdapterEthSettings) HasLldpEnabled() bool {
	if o != nil && o.LldpEnabled != nil {
		return true
	}

	return false
}

// SetLldpEnabled gets a reference to the given bool and assigns it to the LldpEnabled field.
func (o *AdapterEthSettings) SetLldpEnabled(v bool) {
	o.LldpEnabled = &v
}

func (o AdapterEthSettings) MarshalJSON() ([]byte, error) {
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
	if o.LldpEnabled != nil {
		toSerialize["LldpEnabled"] = o.LldpEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterEthSettings) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterEthSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of LLDP protocol on the adapter interfaces.
		LldpEnabled *bool `json:"LldpEnabled,omitempty"`
	}

	varAdapterEthSettingsWithoutEmbeddedStruct := AdapterEthSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterEthSettingsWithoutEmbeddedStruct)
	if err == nil {
		varAdapterEthSettings := _AdapterEthSettings{}
		varAdapterEthSettings.ClassId = varAdapterEthSettingsWithoutEmbeddedStruct.ClassId
		varAdapterEthSettings.ObjectType = varAdapterEthSettingsWithoutEmbeddedStruct.ObjectType
		varAdapterEthSettings.LldpEnabled = varAdapterEthSettingsWithoutEmbeddedStruct.LldpEnabled
		*o = AdapterEthSettings(varAdapterEthSettings)
	} else {
		return err
	}

	varAdapterEthSettings := _AdapterEthSettings{}

	err = json.Unmarshal(bytes, &varAdapterEthSettings)
	if err == nil {
		o.MoBaseComplexType = varAdapterEthSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LldpEnabled")

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

type NullableAdapterEthSettings struct {
	value *AdapterEthSettings
	isSet bool
}

func (v NullableAdapterEthSettings) Get() *AdapterEthSettings {
	return v.value
}

func (v *NullableAdapterEthSettings) Set(val *AdapterEthSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterEthSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterEthSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterEthSettings(val *AdapterEthSettings) *NullableAdapterEthSettings {
	return &NullableAdapterEthSettings{value: val, isSet: true}
}

func (v NullableAdapterEthSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterEthSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
