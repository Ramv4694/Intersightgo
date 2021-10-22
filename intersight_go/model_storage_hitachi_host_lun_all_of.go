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

// StorageHitachiHostLunAllOf Definition of the list of properties defined in 'storage.HitachiHostLun', excluding properties defined in parent classes.
type StorageHitachiHostLunAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Port ID of the Hitachi host lun.
	PortId               *string                              `json:"PortId,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	Host                 *StorageHitachiHostRelationship      `json:"Host,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Volume               *StorageHitachiVolumeRelationship    `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiHostLunAllOf StorageHitachiHostLunAllOf

// NewStorageHitachiHostLunAllOf instantiates a new StorageHitachiHostLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiHostLunAllOf(classId string, objectType string) *StorageHitachiHostLunAllOf {
	this := StorageHitachiHostLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiHostLunAllOfWithDefaults instantiates a new StorageHitachiHostLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiHostLunAllOfWithDefaults() *StorageHitachiHostLunAllOf {
	this := StorageHitachiHostLunAllOf{}
	var classId string = "storage.HitachiHostLun"
	this.ClassId = classId
	var objectType string = "storage.HitachiHostLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiHostLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiHostLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiHostLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiHostLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageHitachiHostLunAllOf) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageHitachiHostLunAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageHitachiHostLunAllOf) SetPortId(v string) {
	o.PortId = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiHostLunAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiHostLunAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiHostLunAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *StorageHitachiHostLunAllOf) GetHost() StorageHitachiHostRelationship {
	if o == nil || o.Host == nil {
		var ret StorageHitachiHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetHostOk() (*StorageHitachiHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StorageHitachiHostLunAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given StorageHitachiHostRelationship and assigns it to the Host field.
func (o *StorageHitachiHostLunAllOf) SetHost(v StorageHitachiHostRelationship) {
	o.Host = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiHostLunAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiHostLunAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiHostLunAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageHitachiHostLunAllOf) GetVolume() StorageHitachiVolumeRelationship {
	if o == nil || o.Volume == nil {
		var ret StorageHitachiVolumeRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiHostLunAllOf) GetVolumeOk() (*StorageHitachiVolumeRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageHitachiHostLunAllOf) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StorageHitachiVolumeRelationship and assigns it to the Volume field.
func (o *StorageHitachiHostLunAllOf) SetVolume(v StorageHitachiVolumeRelationship) {
	o.Volume = &v
}

func (o StorageHitachiHostLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
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

func (o *StorageHitachiHostLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiHostLunAllOf := _StorageHitachiHostLunAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiHostLunAllOf); err == nil {
		*o = StorageHitachiHostLunAllOf(varStorageHitachiHostLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volume")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiHostLunAllOf struct {
	value *StorageHitachiHostLunAllOf
	isSet bool
}

func (v NullableStorageHitachiHostLunAllOf) Get() *StorageHitachiHostLunAllOf {
	return v.value
}

func (v *NullableStorageHitachiHostLunAllOf) Set(val *StorageHitachiHostLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiHostLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiHostLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiHostLunAllOf(val *StorageHitachiHostLunAllOf) *NullableStorageHitachiHostLunAllOf {
	return &NullableStorageHitachiHostLunAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiHostLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiHostLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
