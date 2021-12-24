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
	"reflect"
	"strings"
)

// OsPhysicalDiskResponse Physical Disk target entry for the UI.
type OsPhysicalDiskResponse struct {
	OsInstallTargetResponse
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bootable field of the Physical Drive target.
	Bootable *string `json:"Bootable,omitempty"`
	// The Physical Disk Name to be used as Install Target.
	Name *string `json:"Name,omitempty"`
	// Serial Number of the Physical Disk Target.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The Storage Controller associated to the physical disk.
	StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OsPhysicalDiskResponse OsPhysicalDiskResponse

// NewOsPhysicalDiskResponse instantiates a new OsPhysicalDiskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsPhysicalDiskResponse(classId string, objectType string) *OsPhysicalDiskResponse {
	this := OsPhysicalDiskResponse{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsPhysicalDiskResponseWithDefaults instantiates a new OsPhysicalDiskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsPhysicalDiskResponseWithDefaults() *OsPhysicalDiskResponse {
	this := OsPhysicalDiskResponse{}
	var classId string = "os.PhysicalDiskResponse"
	this.ClassId = classId
	var objectType string = "os.PhysicalDiskResponse"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsPhysicalDiskResponse) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsPhysicalDiskResponse) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsPhysicalDiskResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsPhysicalDiskResponse) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *OsPhysicalDiskResponse) SetBootable(v string) {
	o.Bootable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsPhysicalDiskResponse) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OsPhysicalDiskResponse) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStorageControllerSlotId returns the StorageControllerSlotId field value if set, zero value otherwise.
func (o *OsPhysicalDiskResponse) GetStorageControllerSlotId() string {
	if o == nil || o.StorageControllerSlotId == nil {
		var ret string
		return ret
	}
	return *o.StorageControllerSlotId
}

// GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsPhysicalDiskResponse) GetStorageControllerSlotIdOk() (*string, bool) {
	if o == nil || o.StorageControllerSlotId == nil {
		return nil, false
	}
	return o.StorageControllerSlotId, true
}

// HasStorageControllerSlotId returns a boolean if a field has been set.
func (o *OsPhysicalDiskResponse) HasStorageControllerSlotId() bool {
	if o != nil && o.StorageControllerSlotId != nil {
		return true
	}

	return false
}

// SetStorageControllerSlotId gets a reference to the given string and assigns it to the StorageControllerSlotId field.
func (o *OsPhysicalDiskResponse) SetStorageControllerSlotId(v string) {
	o.StorageControllerSlotId = &v
}

func (o OsPhysicalDiskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsInstallTargetResponse, errOsInstallTargetResponse := json.Marshal(o.OsInstallTargetResponse)
	if errOsInstallTargetResponse != nil {
		return []byte{}, errOsInstallTargetResponse
	}
	errOsInstallTargetResponse = json.Unmarshal([]byte(serializedOsInstallTargetResponse), &toSerialize)
	if errOsInstallTargetResponse != nil {
		return []byte{}, errOsInstallTargetResponse
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.StorageControllerSlotId != nil {
		toSerialize["StorageControllerSlotId"] = o.StorageControllerSlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsPhysicalDiskResponse) UnmarshalJSON(bytes []byte) (err error) {
	type OsPhysicalDiskResponseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bootable field of the Physical Drive target.
		Bootable *string `json:"Bootable,omitempty"`
		// The Physical Disk Name to be used as Install Target.
		Name *string `json:"Name,omitempty"`
		// Serial Number of the Physical Disk Target.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The Storage Controller associated to the physical disk.
		StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	}

	varOsPhysicalDiskResponseWithoutEmbeddedStruct := OsPhysicalDiskResponseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsPhysicalDiskResponseWithoutEmbeddedStruct)
	if err == nil {
		varOsPhysicalDiskResponse := _OsPhysicalDiskResponse{}
		varOsPhysicalDiskResponse.ClassId = varOsPhysicalDiskResponseWithoutEmbeddedStruct.ClassId
		varOsPhysicalDiskResponse.ObjectType = varOsPhysicalDiskResponseWithoutEmbeddedStruct.ObjectType
		varOsPhysicalDiskResponse.Bootable = varOsPhysicalDiskResponseWithoutEmbeddedStruct.Bootable
		varOsPhysicalDiskResponse.Name = varOsPhysicalDiskResponseWithoutEmbeddedStruct.Name
		varOsPhysicalDiskResponse.SerialNumber = varOsPhysicalDiskResponseWithoutEmbeddedStruct.SerialNumber
		varOsPhysicalDiskResponse.StorageControllerSlotId = varOsPhysicalDiskResponseWithoutEmbeddedStruct.StorageControllerSlotId
		*o = OsPhysicalDiskResponse(varOsPhysicalDiskResponse)
	} else {
		return err
	}

	varOsPhysicalDiskResponse := _OsPhysicalDiskResponse{}

	err = json.Unmarshal(bytes, &varOsPhysicalDiskResponse)
	if err == nil {
		o.OsInstallTargetResponse = varOsPhysicalDiskResponse.OsInstallTargetResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "StorageControllerSlotId")

		// remove fields from embedded structs
		reflectOsInstallTargetResponse := reflect.ValueOf(o.OsInstallTargetResponse)
		for i := 0; i < reflectOsInstallTargetResponse.Type().NumField(); i++ {
			t := reflectOsInstallTargetResponse.Type().Field(i)

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

type NullableOsPhysicalDiskResponse struct {
	value *OsPhysicalDiskResponse
	isSet bool
}

func (v NullableOsPhysicalDiskResponse) Get() *OsPhysicalDiskResponse {
	return v.value
}

func (v *NullableOsPhysicalDiskResponse) Set(val *OsPhysicalDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOsPhysicalDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOsPhysicalDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsPhysicalDiskResponse(val *OsPhysicalDiskResponse) *NullableOsPhysicalDiskResponse {
	return &NullableOsPhysicalDiskResponse{value: val, isSet: true}
}

func (v NullableOsPhysicalDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsPhysicalDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
