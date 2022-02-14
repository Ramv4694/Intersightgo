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

// StorageM2VirtualDriveConfigAllOf Definition of the list of properties defined in 'storage.M2VirtualDriveConfig', excluding properties defined in parent classes.
type StorageM2VirtualDriveConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Select the M.2 RAID controller slot on which the virtual drive is to be created. Select 'MSTOR-RAID-1' to create virtual drive on the M.2 RAID controller in the first slot or in the MSTOR-RAID slot, 'MSTOR-RAID-2' for second slot, 'MSTOR-RAID-1, MSTOR-RAID-2' for both slots or either slot. * `MSTOR-RAID-1` - Virtual drive  will be created on the M.2 RAID controller in the first slot. * `MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available. * `MSTOR-RAID-1,MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available.
	ControllerSlot *string `json:"ControllerSlot,omitempty"`
	// If enabled, this will create a virtual drive on the M.2 RAID controller.
	Enable               *bool `json:"Enable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageM2VirtualDriveConfigAllOf StorageM2VirtualDriveConfigAllOf

// NewStorageM2VirtualDriveConfigAllOf instantiates a new StorageM2VirtualDriveConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageM2VirtualDriveConfigAllOf(classId string, objectType string) *StorageM2VirtualDriveConfigAllOf {
	this := StorageM2VirtualDriveConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var controllerSlot string = "MSTOR-RAID-1"
	this.ControllerSlot = &controllerSlot
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewStorageM2VirtualDriveConfigAllOfWithDefaults instantiates a new StorageM2VirtualDriveConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageM2VirtualDriveConfigAllOfWithDefaults() *StorageM2VirtualDriveConfigAllOf {
	this := StorageM2VirtualDriveConfigAllOf{}
	var classId string = "storage.M2VirtualDriveConfig"
	this.ClassId = classId
	var objectType string = "storage.M2VirtualDriveConfig"
	this.ObjectType = objectType
	var controllerSlot string = "MSTOR-RAID-1"
	this.ControllerSlot = &controllerSlot
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageM2VirtualDriveConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageM2VirtualDriveConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageM2VirtualDriveConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageM2VirtualDriveConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerSlot returns the ControllerSlot field value if set, zero value otherwise.
func (o *StorageM2VirtualDriveConfigAllOf) GetControllerSlot() string {
	if o == nil || o.ControllerSlot == nil {
		var ret string
		return ret
	}
	return *o.ControllerSlot
}

// GetControllerSlotOk returns a tuple with the ControllerSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfigAllOf) GetControllerSlotOk() (*string, bool) {
	if o == nil || o.ControllerSlot == nil {
		return nil, false
	}
	return o.ControllerSlot, true
}

// HasControllerSlot returns a boolean if a field has been set.
func (o *StorageM2VirtualDriveConfigAllOf) HasControllerSlot() bool {
	if o != nil && o.ControllerSlot != nil {
		return true
	}

	return false
}

// SetControllerSlot gets a reference to the given string and assigns it to the ControllerSlot field.
func (o *StorageM2VirtualDriveConfigAllOf) SetControllerSlot(v string) {
	o.ControllerSlot = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *StorageM2VirtualDriveConfigAllOf) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfigAllOf) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *StorageM2VirtualDriveConfigAllOf) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *StorageM2VirtualDriveConfigAllOf) SetEnable(v bool) {
	o.Enable = &v
}

func (o StorageM2VirtualDriveConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerSlot != nil {
		toSerialize["ControllerSlot"] = o.ControllerSlot
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageM2VirtualDriveConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageM2VirtualDriveConfigAllOf := _StorageM2VirtualDriveConfigAllOf{}

	if err = json.Unmarshal(bytes, &varStorageM2VirtualDriveConfigAllOf); err == nil {
		*o = StorageM2VirtualDriveConfigAllOf(varStorageM2VirtualDriveConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerSlot")
		delete(additionalProperties, "Enable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageM2VirtualDriveConfigAllOf struct {
	value *StorageM2VirtualDriveConfigAllOf
	isSet bool
}

func (v NullableStorageM2VirtualDriveConfigAllOf) Get() *StorageM2VirtualDriveConfigAllOf {
	return v.value
}

func (v *NullableStorageM2VirtualDriveConfigAllOf) Set(val *StorageM2VirtualDriveConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageM2VirtualDriveConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageM2VirtualDriveConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageM2VirtualDriveConfigAllOf(val *StorageM2VirtualDriveConfigAllOf) *NullableStorageM2VirtualDriveConfigAllOf {
	return &NullableStorageM2VirtualDriveConfigAllOf{value: val, isSet: true}
}

func (v NullableStorageM2VirtualDriveConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageM2VirtualDriveConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}