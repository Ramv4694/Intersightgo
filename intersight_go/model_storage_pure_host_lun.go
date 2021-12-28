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

// StoragePureHostLun A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
type StoragePureHostLun struct {
	StorageBaseHostLun
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the host group associated with LUN.
	HostGroupName *string `json:"HostGroupName,omitempty"`
	// Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
	Shared               *bool                                `json:"Shared,omitempty"`
	Array                *StoragePureArrayRelationship        `json:"Array,omitempty"`
	Host                 *StoragePureHostRelationship         `json:"Host,omitempty"`
	HostGroup            *StoragePureHostGroupRelationship    `json:"HostGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Volume               *StoragePureVolumeRelationship       `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureHostLun StoragePureHostLun

// NewStoragePureHostLun instantiates a new StoragePureHostLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostLun(classId string, objectType string) *StoragePureHostLun {
	this := StoragePureHostLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureHostLunWithDefaults instantiates a new StoragePureHostLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostLunWithDefaults() *StoragePureHostLun {
	this := StoragePureHostLun{}
	var classId string = "storage.PureHostLun"
	this.ClassId = classId
	var objectType string = "storage.PureHostLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureHostLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureHostLun) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureHostLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureHostLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostGroupName returns the HostGroupName field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetHostGroupName() string {
	if o == nil || o.HostGroupName == nil {
		var ret string
		return ret
	}
	return *o.HostGroupName
}

// GetHostGroupNameOk returns a tuple with the HostGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetHostGroupNameOk() (*string, bool) {
	if o == nil || o.HostGroupName == nil {
		return nil, false
	}
	return o.HostGroupName, true
}

// HasHostGroupName returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHostGroupName() bool {
	if o != nil && o.HostGroupName != nil {
		return true
	}

	return false
}

// SetHostGroupName gets a reference to the given string and assigns it to the HostGroupName field.
func (o *StoragePureHostLun) SetHostGroupName(v string) {
	o.HostGroupName = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetShared() bool {
	if o == nil || o.Shared == nil {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetSharedOk() (*bool, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *StoragePureHostLun) SetShared(v bool) {
	o.Shared = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostLun) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetHost() StoragePureHostRelationship {
	if o == nil || o.Host == nil {
		var ret StoragePureHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetHostOk() (*StoragePureHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given StoragePureHostRelationship and assigns it to the Host field.
func (o *StoragePureHostLun) SetHost(v StoragePureHostRelationship) {
	o.Host = &v
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetHostGroup() StoragePureHostGroupRelationship {
	if o == nil || o.HostGroup == nil {
		var ret StoragePureHostGroupRelationship
		return ret
	}
	return *o.HostGroup
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool) {
	if o == nil || o.HostGroup == nil {
		return nil, false
	}
	return o.HostGroup, true
}

// HasHostGroup returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasHostGroup() bool {
	if o != nil && o.HostGroup != nil {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given StoragePureHostGroupRelationship and assigns it to the HostGroup field.
func (o *StoragePureHostLun) SetHostGroup(v StoragePureHostGroupRelationship) {
	o.HostGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostLun) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StoragePureHostLun) GetVolume() StoragePureVolumeRelationship {
	if o == nil || o.Volume == nil {
		var ret StoragePureVolumeRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostLun) GetVolumeOk() (*StoragePureVolumeRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StoragePureHostLun) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StoragePureVolumeRelationship and assigns it to the Volume field.
func (o *StoragePureHostLun) SetVolume(v StoragePureVolumeRelationship) {
	o.Volume = &v
}

func (o StoragePureHostLun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostLun, errStorageBaseHostLun := json.Marshal(o.StorageBaseHostLun)
	if errStorageBaseHostLun != nil {
		return []byte{}, errStorageBaseHostLun
	}
	errStorageBaseHostLun = json.Unmarshal([]byte(serializedStorageBaseHostLun), &toSerialize)
	if errStorageBaseHostLun != nil {
		return []byte{}, errStorageBaseHostLun
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostGroupName != nil {
		toSerialize["HostGroupName"] = o.HostGroupName
	}
	if o.Shared != nil {
		toSerialize["Shared"] = o.Shared
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.HostGroup != nil {
		toSerialize["HostGroup"] = o.HostGroup
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

func (o *StoragePureHostLun) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureHostLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the host group associated with LUN.
		HostGroupName *string `json:"HostGroupName,omitempty"`
		// Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection.
		Shared           *bool                                `json:"Shared,omitempty"`
		Array            *StoragePureArrayRelationship        `json:"Array,omitempty"`
		Host             *StoragePureHostRelationship         `json:"Host,omitempty"`
		HostGroup        *StoragePureHostGroupRelationship    `json:"HostGroup,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Volume           *StoragePureVolumeRelationship       `json:"Volume,omitempty"`
	}

	varStoragePureHostLunWithoutEmbeddedStruct := StoragePureHostLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureHostLunWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureHostLun := _StoragePureHostLun{}
		varStoragePureHostLun.ClassId = varStoragePureHostLunWithoutEmbeddedStruct.ClassId
		varStoragePureHostLun.ObjectType = varStoragePureHostLunWithoutEmbeddedStruct.ObjectType
		varStoragePureHostLun.HostGroupName = varStoragePureHostLunWithoutEmbeddedStruct.HostGroupName
		varStoragePureHostLun.Shared = varStoragePureHostLunWithoutEmbeddedStruct.Shared
		varStoragePureHostLun.Array = varStoragePureHostLunWithoutEmbeddedStruct.Array
		varStoragePureHostLun.Host = varStoragePureHostLunWithoutEmbeddedStruct.Host
		varStoragePureHostLun.HostGroup = varStoragePureHostLunWithoutEmbeddedStruct.HostGroup
		varStoragePureHostLun.RegisteredDevice = varStoragePureHostLunWithoutEmbeddedStruct.RegisteredDevice
		varStoragePureHostLun.Volume = varStoragePureHostLunWithoutEmbeddedStruct.Volume
		*o = StoragePureHostLun(varStoragePureHostLun)
	} else {
		return err
	}

	varStoragePureHostLun := _StoragePureHostLun{}

	err = json.Unmarshal(bytes, &varStoragePureHostLun)
	if err == nil {
		o.StorageBaseHostLun = varStoragePureHostLun.StorageBaseHostLun
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostGroupName")
		delete(additionalProperties, "Shared")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "HostGroup")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volume")

		// remove fields from embedded structs
		reflectStorageBaseHostLun := reflect.ValueOf(o.StorageBaseHostLun)
		for i := 0; i < reflectStorageBaseHostLun.Type().NumField(); i++ {
			t := reflectStorageBaseHostLun.Type().Field(i)

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

type NullableStoragePureHostLun struct {
	value *StoragePureHostLun
	isSet bool
}

func (v NullableStoragePureHostLun) Get() *StoragePureHostLun {
	return v.value
}

func (v *NullableStoragePureHostLun) Set(val *StoragePureHostLun) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostLun) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostLun(val *StoragePureHostLun) *NullableStoragePureHostLun {
	return &NullableStoragePureHostLun{value: val, isSet: true}
}

func (v NullableStoragePureHostLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
