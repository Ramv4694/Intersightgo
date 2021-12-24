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
)

// BootLocalDiskAllOf Definition of the list of properties defined in 'boot.LocalDisk', excluding properties defined in parent classes.
type BootLocalDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                 `json:"ObjectType"`
	Bootloader NullableBootBootloader `json:"Bootloader,omitempty"`
	// The slot id of the local disk device. Supported values for Standalone Rack servers are (1-205, \"M\", \"HBA\", \"SAS\", \"RAID\", \"MRAID\", \"MRAID1\", \"MRAID2\", \"MSTOR-RAID\"). Supported values for FI-attached servers are (1-205, \"MRAID\", \"FMEZZ1-SAS\", \"MRAID1\", \"MRAID2\", \"MSTOR-RAID\", \"MSTOR-RAID-1\", \"MSTOR-RAID-2\").
	Slot                 *string `json:"Slot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootLocalDiskAllOf BootLocalDiskAllOf

// NewBootLocalDiskAllOf instantiates a new BootLocalDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootLocalDiskAllOf(classId string, objectType string) *BootLocalDiskAllOf {
	this := BootLocalDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootLocalDiskAllOfWithDefaults instantiates a new BootLocalDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootLocalDiskAllOfWithDefaults() *BootLocalDiskAllOf {
	this := BootLocalDiskAllOf{}
	var classId string = "boot.LocalDisk"
	this.ClassId = classId
	var objectType string = "boot.LocalDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootLocalDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootLocalDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootLocalDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootLocalDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootLocalDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootLocalDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootLocalDiskAllOf) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader.Get() == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader.Get()
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootLocalDiskAllOf) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bootloader.Get(), o.Bootloader.IsSet()
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootLocalDiskAllOf) HasBootloader() bool {
	if o != nil && o.Bootloader.IsSet() {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given NullableBootBootloader and assigns it to the Bootloader field.
func (o *BootLocalDiskAllOf) SetBootloader(v BootBootloader) {
	o.Bootloader.Set(&v)
}

// SetBootloaderNil sets the value for Bootloader to be an explicit nil
func (o *BootLocalDiskAllOf) SetBootloaderNil() {
	o.Bootloader.Set(nil)
}

// UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
func (o *BootLocalDiskAllOf) UnsetBootloader() {
	o.Bootloader.Unset()
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootLocalDiskAllOf) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootLocalDiskAllOf) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootLocalDiskAllOf) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootLocalDiskAllOf) SetSlot(v string) {
	o.Slot = &v
}

func (o BootLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootloader.IsSet() {
		toSerialize["Bootloader"] = o.Bootloader.Get()
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootLocalDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootLocalDiskAllOf := _BootLocalDiskAllOf{}

	if err = json.Unmarshal(bytes, &varBootLocalDiskAllOf); err == nil {
		*o = BootLocalDiskAllOf(varBootLocalDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootloader")
		delete(additionalProperties, "Slot")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootLocalDiskAllOf struct {
	value *BootLocalDiskAllOf
	isSet bool
}

func (v NullableBootLocalDiskAllOf) Get() *BootLocalDiskAllOf {
	return v.value
}

func (v *NullableBootLocalDiskAllOf) Set(val *BootLocalDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootLocalDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootLocalDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootLocalDiskAllOf(val *BootLocalDiskAllOf) *NullableBootLocalDiskAllOf {
	return &NullableBootLocalDiskAllOf{value: val, isSet: true}
}

func (v NullableBootLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootLocalDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
