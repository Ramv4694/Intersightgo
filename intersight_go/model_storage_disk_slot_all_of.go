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

// StorageDiskSlotAllOf Definition of the list of properties defined in 'storage.DiskSlot', excluding properties defined in parent classes.
type StorageDiskSlotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDisk         *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    *StorageControllerRelationship       `json:"StorageController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageDiskSlotAllOf StorageDiskSlotAllOf

// NewStorageDiskSlotAllOf instantiates a new StorageDiskSlotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskSlotAllOf(classId string, objectType string) *StorageDiskSlotAllOf {
	this := StorageDiskSlotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageDiskSlotAllOfWithDefaults instantiates a new StorageDiskSlotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskSlotAllOfWithDefaults() *StorageDiskSlotAllOf {
	this := StorageDiskSlotAllOf{}
	var classId string = "storage.DiskSlot"
	this.ClassId = classId
	var objectType string = "storage.DiskSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDiskSlotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDiskSlotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageDiskSlotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDiskSlotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageDiskSlotAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageDiskSlotAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageDiskSlotAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise.
func (o *StorageDiskSlotAllOf) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.PhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisk == nil {
		return nil, false
	}
	return o.PhysicalDisk, true
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StorageDiskSlotAllOf) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk != nil {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StorageDiskSlotAllOf) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageDiskSlotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageDiskSlotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageDiskSlotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StorageDiskSlotAllOf) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskSlotAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageDiskSlotAllOf) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageDiskSlotAllOf) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

func (o StorageDiskSlotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDisk != nil {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageDiskSlotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageDiskSlotAllOf := _StorageDiskSlotAllOf{}

	if err = json.Unmarshal(bytes, &varStorageDiskSlotAllOf); err == nil {
		*o = StorageDiskSlotAllOf(varStorageDiskSlotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisk")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageDiskSlotAllOf struct {
	value *StorageDiskSlotAllOf
	isSet bool
}

func (v NullableStorageDiskSlotAllOf) Get() *StorageDiskSlotAllOf {
	return v.value
}

func (v *NullableStorageDiskSlotAllOf) Set(val *StorageDiskSlotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskSlotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskSlotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskSlotAllOf(val *StorageDiskSlotAllOf) *NullableStorageDiskSlotAllOf {
	return &NullableStorageDiskSlotAllOf{value: val, isSet: true}
}

func (v NullableStorageDiskSlotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskSlotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
