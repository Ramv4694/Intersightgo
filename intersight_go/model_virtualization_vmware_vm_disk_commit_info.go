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

// VirtualizationVmwareVmDiskCommitInfo Information about the virtual machine's disk commits, sharing and limits. For more information, see VMware documentation.
type VirtualizationVmwareVmDiskCommitInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk committed in bytes on this virtual machine (disk space used up).
	CommittedDisk *int64 `json:"CommittedDisk,omitempty"`
	// Total uncommitted disk space that is available for use (in bytes).
	UnCommittedDisk *int64 `json:"UnCommittedDisk,omitempty"`
	// Total unshared disk space (in bytes).
	UnsharedDisk         *int64 `json:"UnsharedDisk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmDiskCommitInfo VirtualizationVmwareVmDiskCommitInfo

// NewVirtualizationVmwareVmDiskCommitInfo instantiates a new VirtualizationVmwareVmDiskCommitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmDiskCommitInfo(classId string, objectType string) *VirtualizationVmwareVmDiskCommitInfo {
	this := VirtualizationVmwareVmDiskCommitInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmDiskCommitInfoWithDefaults instantiates a new VirtualizationVmwareVmDiskCommitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmDiskCommitInfoWithDefaults() *VirtualizationVmwareVmDiskCommitInfo {
	this := VirtualizationVmwareVmDiskCommitInfo{}
	var classId string = "virtualization.VmwareVmDiskCommitInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmDiskCommitInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmDiskCommitInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmDiskCommitInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmDiskCommitInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmDiskCommitInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommittedDisk returns the CommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetCommittedDisk() int64 {
	if o == nil || o.CommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.CommittedDisk
}

// GetCommittedDiskOk returns a tuple with the CommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetCommittedDiskOk() (*int64, bool) {
	if o == nil || o.CommittedDisk == nil {
		return nil, false
	}
	return o.CommittedDisk, true
}

// HasCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) HasCommittedDisk() bool {
	if o != nil && o.CommittedDisk != nil {
		return true
	}

	return false
}

// SetCommittedDisk gets a reference to the given int64 and assigns it to the CommittedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfo) SetCommittedDisk(v int64) {
	o.CommittedDisk = &v
}

// GetUnCommittedDisk returns the UnCommittedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnCommittedDisk() int64 {
	if o == nil || o.UnCommittedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnCommittedDisk
}

// GetUnCommittedDiskOk returns a tuple with the UnCommittedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnCommittedDiskOk() (*int64, bool) {
	if o == nil || o.UnCommittedDisk == nil {
		return nil, false
	}
	return o.UnCommittedDisk, true
}

// HasUnCommittedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) HasUnCommittedDisk() bool {
	if o != nil && o.UnCommittedDisk != nil {
		return true
	}

	return false
}

// SetUnCommittedDisk gets a reference to the given int64 and assigns it to the UnCommittedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfo) SetUnCommittedDisk(v int64) {
	o.UnCommittedDisk = &v
}

// GetUnsharedDisk returns the UnsharedDisk field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnsharedDisk() int64 {
	if o == nil || o.UnsharedDisk == nil {
		var ret int64
		return ret
	}
	return *o.UnsharedDisk
}

// GetUnsharedDiskOk returns a tuple with the UnsharedDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) GetUnsharedDiskOk() (*int64, bool) {
	if o == nil || o.UnsharedDisk == nil {
		return nil, false
	}
	return o.UnsharedDisk, true
}

// HasUnsharedDisk returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmDiskCommitInfo) HasUnsharedDisk() bool {
	if o != nil && o.UnsharedDisk != nil {
		return true
	}

	return false
}

// SetUnsharedDisk gets a reference to the given int64 and assigns it to the UnsharedDisk field.
func (o *VirtualizationVmwareVmDiskCommitInfo) SetUnsharedDisk(v int64) {
	o.UnsharedDisk = &v
}

func (o VirtualizationVmwareVmDiskCommitInfo) MarshalJSON() ([]byte, error) {
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
	if o.CommittedDisk != nil {
		toSerialize["CommittedDisk"] = o.CommittedDisk
	}
	if o.UnCommittedDisk != nil {
		toSerialize["UnCommittedDisk"] = o.UnCommittedDisk
	}
	if o.UnsharedDisk != nil {
		toSerialize["UnsharedDisk"] = o.UnsharedDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmDiskCommitInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Disk committed in bytes on this virtual machine (disk space used up).
		CommittedDisk *int64 `json:"CommittedDisk,omitempty"`
		// Total uncommitted disk space that is available for use (in bytes).
		UnCommittedDisk *int64 `json:"UnCommittedDisk,omitempty"`
		// Total unshared disk space (in bytes).
		UnsharedDisk *int64 `json:"UnsharedDisk,omitempty"`
	}

	varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct := VirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVmDiskCommitInfo := _VirtualizationVmwareVmDiskCommitInfo{}
		varVirtualizationVmwareVmDiskCommitInfo.ClassId = varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVmDiskCommitInfo.ObjectType = varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVmDiskCommitInfo.CommittedDisk = varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct.CommittedDisk
		varVirtualizationVmwareVmDiskCommitInfo.UnCommittedDisk = varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct.UnCommittedDisk
		varVirtualizationVmwareVmDiskCommitInfo.UnsharedDisk = varVirtualizationVmwareVmDiskCommitInfoWithoutEmbeddedStruct.UnsharedDisk
		*o = VirtualizationVmwareVmDiskCommitInfo(varVirtualizationVmwareVmDiskCommitInfo)
	} else {
		return err
	}

	varVirtualizationVmwareVmDiskCommitInfo := _VirtualizationVmwareVmDiskCommitInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmDiskCommitInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareVmDiskCommitInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommittedDisk")
		delete(additionalProperties, "UnCommittedDisk")
		delete(additionalProperties, "UnsharedDisk")

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

type NullableVirtualizationVmwareVmDiskCommitInfo struct {
	value *VirtualizationVmwareVmDiskCommitInfo
	isSet bool
}

func (v NullableVirtualizationVmwareVmDiskCommitInfo) Get() *VirtualizationVmwareVmDiskCommitInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfo) Set(val *VirtualizationVmwareVmDiskCommitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmDiskCommitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmDiskCommitInfo(val *VirtualizationVmwareVmDiskCommitInfo) *NullableVirtualizationVmwareVmDiskCommitInfo {
	return &NullableVirtualizationVmwareVmDiskCommitInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmDiskCommitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmDiskCommitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}