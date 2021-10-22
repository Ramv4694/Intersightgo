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

// StorageR0Drive This models a Single Drive RAID Configuration for creation of on RAID0 virtual drive per physical drive and occupying the full disk space on the same.
type StorageR0Drive struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The set of drive slots where RAID0 virtual drives must be created.
	DriveSlots *string `json:"DriveSlots,omitempty"`
	// The list of drive slots where RAID0 virtual drives must be created (comma seperated).
	DriveSlotsList *string `json:"DriveSlotsList,omitempty"`
	// If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk.
	Enable               *bool                             `json:"Enable,omitempty"`
	VirtualDrivePolicy   NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageR0Drive StorageR0Drive

// NewStorageR0Drive instantiates a new StorageR0Drive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageR0Drive(classId string, objectType string) *StorageR0Drive {
	this := StorageR0Drive{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewStorageR0DriveWithDefaults instantiates a new StorageR0Drive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageR0DriveWithDefaults() *StorageR0Drive {
	this := StorageR0Drive{}
	var classId string = "storage.R0Drive"
	this.ClassId = classId
	var objectType string = "storage.R0Drive"
	this.ObjectType = objectType
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageR0Drive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageR0Drive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageR0Drive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageR0Drive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageR0Drive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageR0Drive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveSlots returns the DriveSlots field value if set, zero value otherwise.
func (o *StorageR0Drive) GetDriveSlots() string {
	if o == nil || o.DriveSlots == nil {
		var ret string
		return ret
	}
	return *o.DriveSlots
}

// GetDriveSlotsOk returns a tuple with the DriveSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0Drive) GetDriveSlotsOk() (*string, bool) {
	if o == nil || o.DriveSlots == nil {
		return nil, false
	}
	return o.DriveSlots, true
}

// HasDriveSlots returns a boolean if a field has been set.
func (o *StorageR0Drive) HasDriveSlots() bool {
	if o != nil && o.DriveSlots != nil {
		return true
	}

	return false
}

// SetDriveSlots gets a reference to the given string and assigns it to the DriveSlots field.
func (o *StorageR0Drive) SetDriveSlots(v string) {
	o.DriveSlots = &v
}

// GetDriveSlotsList returns the DriveSlotsList field value if set, zero value otherwise.
func (o *StorageR0Drive) GetDriveSlotsList() string {
	if o == nil || o.DriveSlotsList == nil {
		var ret string
		return ret
	}
	return *o.DriveSlotsList
}

// GetDriveSlotsListOk returns a tuple with the DriveSlotsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0Drive) GetDriveSlotsListOk() (*string, bool) {
	if o == nil || o.DriveSlotsList == nil {
		return nil, false
	}
	return o.DriveSlotsList, true
}

// HasDriveSlotsList returns a boolean if a field has been set.
func (o *StorageR0Drive) HasDriveSlotsList() bool {
	if o != nil && o.DriveSlotsList != nil {
		return true
	}

	return false
}

// SetDriveSlotsList gets a reference to the given string and assigns it to the DriveSlotsList field.
func (o *StorageR0Drive) SetDriveSlotsList(v string) {
	o.DriveSlotsList = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *StorageR0Drive) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0Drive) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *StorageR0Drive) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *StorageR0Drive) SetEnable(v bool) {
	o.Enable = &v
}

// GetVirtualDrivePolicy returns the VirtualDrivePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageR0Drive) GetVirtualDrivePolicy() StorageVirtualDrivePolicy {
	if o == nil || o.VirtualDrivePolicy.Get() == nil {
		var ret StorageVirtualDrivePolicy
		return ret
	}
	return *o.VirtualDrivePolicy.Get()
}

// GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageR0Drive) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDrivePolicy.Get(), o.VirtualDrivePolicy.IsSet()
}

// HasVirtualDrivePolicy returns a boolean if a field has been set.
func (o *StorageR0Drive) HasVirtualDrivePolicy() bool {
	if o != nil && o.VirtualDrivePolicy.IsSet() {
		return true
	}

	return false
}

// SetVirtualDrivePolicy gets a reference to the given NullableStorageVirtualDrivePolicy and assigns it to the VirtualDrivePolicy field.
func (o *StorageR0Drive) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy) {
	o.VirtualDrivePolicy.Set(&v)
}

// SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil
func (o *StorageR0Drive) SetVirtualDrivePolicyNil() {
	o.VirtualDrivePolicy.Set(nil)
}

// UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil
func (o *StorageR0Drive) UnsetVirtualDrivePolicy() {
	o.VirtualDrivePolicy.Unset()
}

func (o StorageR0Drive) MarshalJSON() ([]byte, error) {
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
	if o.DriveSlots != nil {
		toSerialize["DriveSlots"] = o.DriveSlots
	}
	if o.DriveSlotsList != nil {
		toSerialize["DriveSlotsList"] = o.DriveSlotsList
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.VirtualDrivePolicy.IsSet() {
		toSerialize["VirtualDrivePolicy"] = o.VirtualDrivePolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageR0Drive) UnmarshalJSON(bytes []byte) (err error) {
	type StorageR0DriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The set of drive slots where RAID0 virtual drives must be created.
		DriveSlots *string `json:"DriveSlots,omitempty"`
		// The list of drive slots where RAID0 virtual drives must be created (comma seperated).
		DriveSlotsList *string `json:"DriveSlotsList,omitempty"`
		// If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk.
		Enable             *bool                             `json:"Enable,omitempty"`
		VirtualDrivePolicy NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	}

	varStorageR0DriveWithoutEmbeddedStruct := StorageR0DriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageR0DriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageR0Drive := _StorageR0Drive{}
		varStorageR0Drive.ClassId = varStorageR0DriveWithoutEmbeddedStruct.ClassId
		varStorageR0Drive.ObjectType = varStorageR0DriveWithoutEmbeddedStruct.ObjectType
		varStorageR0Drive.DriveSlots = varStorageR0DriveWithoutEmbeddedStruct.DriveSlots
		varStorageR0Drive.DriveSlotsList = varStorageR0DriveWithoutEmbeddedStruct.DriveSlotsList
		varStorageR0Drive.Enable = varStorageR0DriveWithoutEmbeddedStruct.Enable
		varStorageR0Drive.VirtualDrivePolicy = varStorageR0DriveWithoutEmbeddedStruct.VirtualDrivePolicy
		*o = StorageR0Drive(varStorageR0Drive)
	} else {
		return err
	}

	varStorageR0Drive := _StorageR0Drive{}

	err = json.Unmarshal(bytes, &varStorageR0Drive)
	if err == nil {
		o.MoBaseComplexType = varStorageR0Drive.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveSlots")
		delete(additionalProperties, "DriveSlotsList")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "VirtualDrivePolicy")

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

type NullableStorageR0Drive struct {
	value *StorageR0Drive
	isSet bool
}

func (v NullableStorageR0Drive) Get() *StorageR0Drive {
	return v.value
}

func (v *NullableStorageR0Drive) Set(val *StorageR0Drive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageR0Drive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageR0Drive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageR0Drive(val *StorageR0Drive) *NullableStorageR0Drive {
	return &NullableStorageR0Drive{value: val, isSet: true}
}

func (v NullableStorageR0Drive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageR0Drive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
