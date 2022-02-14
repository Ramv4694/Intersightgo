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
	"reflect"
	"strings"
)

// VirtualizationVmwareVlanRange Details of the VLAN range for the trunk port.
type VirtualizationVmwareVlanRange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// End value of VLAN range for the trunk port. The valid range is from 0 to 4094.
	VlanRangeEnd *int64 `json:"VlanRangeEnd,omitempty"`
	// Start value of VLAN range for the trunk port. The valid range is from 0 to 4094.
	VlanRangeStart       *int64 `json:"VlanRangeStart,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVlanRange VirtualizationVmwareVlanRange

// NewVirtualizationVmwareVlanRange instantiates a new VirtualizationVmwareVlanRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVlanRange(classId string, objectType string) *VirtualizationVmwareVlanRange {
	this := VirtualizationVmwareVlanRange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVlanRangeWithDefaults instantiates a new VirtualizationVmwareVlanRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVlanRangeWithDefaults() *VirtualizationVmwareVlanRange {
	this := VirtualizationVmwareVlanRange{}
	var classId string = "virtualization.VmwareVlanRange"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVlanRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVlanRange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVlanRange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVlanRange) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVlanRange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVlanRange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVlanRange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVlanRangeEnd returns the VlanRangeEnd field value if set, zero value otherwise.
func (o *VirtualizationVmwareVlanRange) GetVlanRangeEnd() int64 {
	if o == nil || o.VlanRangeEnd == nil {
		var ret int64
		return ret
	}
	return *o.VlanRangeEnd
}

// GetVlanRangeEndOk returns a tuple with the VlanRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVlanRange) GetVlanRangeEndOk() (*int64, bool) {
	if o == nil || o.VlanRangeEnd == nil {
		return nil, false
	}
	return o.VlanRangeEnd, true
}

// HasVlanRangeEnd returns a boolean if a field has been set.
func (o *VirtualizationVmwareVlanRange) HasVlanRangeEnd() bool {
	if o != nil && o.VlanRangeEnd != nil {
		return true
	}

	return false
}

// SetVlanRangeEnd gets a reference to the given int64 and assigns it to the VlanRangeEnd field.
func (o *VirtualizationVmwareVlanRange) SetVlanRangeEnd(v int64) {
	o.VlanRangeEnd = &v
}

// GetVlanRangeStart returns the VlanRangeStart field value if set, zero value otherwise.
func (o *VirtualizationVmwareVlanRange) GetVlanRangeStart() int64 {
	if o == nil || o.VlanRangeStart == nil {
		var ret int64
		return ret
	}
	return *o.VlanRangeStart
}

// GetVlanRangeStartOk returns a tuple with the VlanRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVlanRange) GetVlanRangeStartOk() (*int64, bool) {
	if o == nil || o.VlanRangeStart == nil {
		return nil, false
	}
	return o.VlanRangeStart, true
}

// HasVlanRangeStart returns a boolean if a field has been set.
func (o *VirtualizationVmwareVlanRange) HasVlanRangeStart() bool {
	if o != nil && o.VlanRangeStart != nil {
		return true
	}

	return false
}

// SetVlanRangeStart gets a reference to the given int64 and assigns it to the VlanRangeStart field.
func (o *VirtualizationVmwareVlanRange) SetVlanRangeStart(v int64) {
	o.VlanRangeStart = &v
}

func (o VirtualizationVmwareVlanRange) MarshalJSON() ([]byte, error) {
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
	if o.VlanRangeEnd != nil {
		toSerialize["VlanRangeEnd"] = o.VlanRangeEnd
	}
	if o.VlanRangeStart != nil {
		toSerialize["VlanRangeStart"] = o.VlanRangeStart
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVlanRange) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVlanRangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// End value of VLAN range for the trunk port. The valid range is from 0 to 4094.
		VlanRangeEnd *int64 `json:"VlanRangeEnd,omitempty"`
		// Start value of VLAN range for the trunk port. The valid range is from 0 to 4094.
		VlanRangeStart *int64 `json:"VlanRangeStart,omitempty"`
	}

	varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct := VirtualizationVmwareVlanRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVlanRange := _VirtualizationVmwareVlanRange{}
		varVirtualizationVmwareVlanRange.ClassId = varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVlanRange.ObjectType = varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVlanRange.VlanRangeEnd = varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct.VlanRangeEnd
		varVirtualizationVmwareVlanRange.VlanRangeStart = varVirtualizationVmwareVlanRangeWithoutEmbeddedStruct.VlanRangeStart
		*o = VirtualizationVmwareVlanRange(varVirtualizationVmwareVlanRange)
	} else {
		return err
	}

	varVirtualizationVmwareVlanRange := _VirtualizationVmwareVlanRange{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVlanRange)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareVlanRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VlanRangeEnd")
		delete(additionalProperties, "VlanRangeStart")

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

type NullableVirtualizationVmwareVlanRange struct {
	value *VirtualizationVmwareVlanRange
	isSet bool
}

func (v NullableVirtualizationVmwareVlanRange) Get() *VirtualizationVmwareVlanRange {
	return v.value
}

func (v *NullableVirtualizationVmwareVlanRange) Set(val *VirtualizationVmwareVlanRange) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVlanRange) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVlanRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVlanRange(val *VirtualizationVmwareVlanRange) *NullableVirtualizationVmwareVlanRange {
	return &NullableVirtualizationVmwareVlanRange{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVlanRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVlanRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}