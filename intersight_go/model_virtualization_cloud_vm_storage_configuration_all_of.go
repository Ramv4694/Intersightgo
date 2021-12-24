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

// VirtualizationCloudVmStorageConfigurationAllOf Definition of the list of properties defined in 'virtualization.CloudVmStorageConfiguration', excluding properties defined in parent classes.
type VirtualizationCloudVmStorageConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType           string                     `json:"ObjectType"`
	Volumes              []VirtualizationVolumeInfo `json:"Volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCloudVmStorageConfigurationAllOf VirtualizationCloudVmStorageConfigurationAllOf

// NewVirtualizationCloudVmStorageConfigurationAllOf instantiates a new VirtualizationCloudVmStorageConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCloudVmStorageConfigurationAllOf(classId string, objectType string) *VirtualizationCloudVmStorageConfigurationAllOf {
	this := VirtualizationCloudVmStorageConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCloudVmStorageConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmStorageConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCloudVmStorageConfigurationAllOfWithDefaults() *VirtualizationCloudVmStorageConfigurationAllOf {
	this := VirtualizationCloudVmStorageConfigurationAllOf{}
	var classId string = "virtualization.AwsVmStorageConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.AwsVmStorageConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetVolumes() []VirtualizationVolumeInfo {
	if o == nil {
		var ret []VirtualizationVolumeInfo
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetVolumesOk() (*[]VirtualizationVolumeInfo, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *VirtualizationCloudVmStorageConfigurationAllOf) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []VirtualizationVolumeInfo and assigns it to the Volumes field.
func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetVolumes(v []VirtualizationVolumeInfo) {
	o.Volumes = v
}

func (o VirtualizationCloudVmStorageConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCloudVmStorageConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationCloudVmStorageConfigurationAllOf := _VirtualizationCloudVmStorageConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationCloudVmStorageConfigurationAllOf); err == nil {
		*o = VirtualizationCloudVmStorageConfigurationAllOf(varVirtualizationCloudVmStorageConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Volumes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationCloudVmStorageConfigurationAllOf struct {
	value *VirtualizationCloudVmStorageConfigurationAllOf
	isSet bool
}

func (v NullableVirtualizationCloudVmStorageConfigurationAllOf) Get() *VirtualizationCloudVmStorageConfigurationAllOf {
	return v.value
}

func (v *NullableVirtualizationCloudVmStorageConfigurationAllOf) Set(val *VirtualizationCloudVmStorageConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCloudVmStorageConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCloudVmStorageConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCloudVmStorageConfigurationAllOf(val *VirtualizationCloudVmStorageConfigurationAllOf) *NullableVirtualizationCloudVmStorageConfigurationAllOf {
	return &NullableVirtualizationCloudVmStorageConfigurationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationCloudVmStorageConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCloudVmStorageConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
