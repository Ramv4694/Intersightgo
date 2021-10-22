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
)

// StoragePureVolumeSnapshotAllOf Definition of the list of properties defined in 'storage.PureVolumeSnapshot', excluding properties defined in parent classes.
type StoragePureVolumeSnapshotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique serial number of the snapshot allocated by the storage array.
	Serial                  *string                                         `json:"Serial,omitempty"`
	Array                   *StoragePureArrayRelationship                   `json:"Array,omitempty"`
	ProtectionGroupSnapshot *StoragePureProtectionGroupSnapshotRelationship `json:"ProtectionGroupSnapshot,omitempty"`
	RegisteredDevice        *AssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	Volume                  *StoragePureVolumeRelationship                  `json:"Volume,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _StoragePureVolumeSnapshotAllOf StoragePureVolumeSnapshotAllOf

// NewStoragePureVolumeSnapshotAllOf instantiates a new StoragePureVolumeSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureVolumeSnapshotAllOf(classId string, objectType string) *StoragePureVolumeSnapshotAllOf {
	this := StoragePureVolumeSnapshotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureVolumeSnapshotAllOfWithDefaults instantiates a new StoragePureVolumeSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureVolumeSnapshotAllOfWithDefaults() *StoragePureVolumeSnapshotAllOf {
	this := StoragePureVolumeSnapshotAllOf{}
	var classId string = "storage.PureVolumeSnapshot"
	this.ClassId = classId
	var objectType string = "storage.PureVolumeSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureVolumeSnapshotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureVolumeSnapshotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureVolumeSnapshotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureVolumeSnapshotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshotAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshotAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *StoragePureVolumeSnapshotAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshotAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshotAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureVolumeSnapshotAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshotAllOf) GetProtectionGroupSnapshot() StoragePureProtectionGroupSnapshotRelationship {
	if o == nil || o.ProtectionGroupSnapshot == nil {
		var ret StoragePureProtectionGroupSnapshotRelationship
		return ret
	}
	return *o.ProtectionGroupSnapshot
}

// GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetProtectionGroupSnapshotOk() (*StoragePureProtectionGroupSnapshotRelationship, bool) {
	if o == nil || o.ProtectionGroupSnapshot == nil {
		return nil, false
	}
	return o.ProtectionGroupSnapshot, true
}

// HasProtectionGroupSnapshot returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshotAllOf) HasProtectionGroupSnapshot() bool {
	if o != nil && o.ProtectionGroupSnapshot != nil {
		return true
	}

	return false
}

// SetProtectionGroupSnapshot gets a reference to the given StoragePureProtectionGroupSnapshotRelationship and assigns it to the ProtectionGroupSnapshot field.
func (o *StoragePureVolumeSnapshotAllOf) SetProtectionGroupSnapshot(v StoragePureProtectionGroupSnapshotRelationship) {
	o.ProtectionGroupSnapshot = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureVolumeSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshotAllOf) GetVolume() StoragePureVolumeRelationship {
	if o == nil || o.Volume == nil {
		var ret StoragePureVolumeRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshotAllOf) GetVolumeOk() (*StoragePureVolumeRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshotAllOf) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StoragePureVolumeRelationship and assigns it to the Volume field.
func (o *StoragePureVolumeSnapshotAllOf) SetVolume(v StoragePureVolumeRelationship) {
	o.Volume = &v
}

func (o StoragePureVolumeSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ProtectionGroupSnapshot != nil {
		toSerialize["ProtectionGroupSnapshot"] = o.ProtectionGroupSnapshot
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureVolumeSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePureVolumeSnapshotAllOf := _StoragePureVolumeSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePureVolumeSnapshotAllOf); err == nil {
		*o = StoragePureVolumeSnapshotAllOf(varStoragePureVolumeSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroupSnapshot")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volume")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePureVolumeSnapshotAllOf struct {
	value *StoragePureVolumeSnapshotAllOf
	isSet bool
}

func (v NullableStoragePureVolumeSnapshotAllOf) Get() *StoragePureVolumeSnapshotAllOf {
	return v.value
}

func (v *NullableStoragePureVolumeSnapshotAllOf) Set(val *StoragePureVolumeSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureVolumeSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureVolumeSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureVolumeSnapshotAllOf(val *StoragePureVolumeSnapshotAllOf) *NullableStoragePureVolumeSnapshotAllOf {
	return &NullableStoragePureVolumeSnapshotAllOf{value: val, isSet: true}
}

func (v NullableStoragePureVolumeSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureVolumeSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
