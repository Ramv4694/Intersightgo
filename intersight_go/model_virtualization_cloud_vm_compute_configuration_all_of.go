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

// VirtualizationCloudVmComputeConfigurationAllOf Definition of the list of properties defined in 'virtualization.CloudVmComputeConfiguration', excluding properties defined in parent classes.
type VirtualizationCloudVmComputeConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Instance Type used by this VM.
	InstanceTypeId       *string `json:"InstanceTypeId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCloudVmComputeConfigurationAllOf VirtualizationCloudVmComputeConfigurationAllOf

// NewVirtualizationCloudVmComputeConfigurationAllOf instantiates a new VirtualizationCloudVmComputeConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCloudVmComputeConfigurationAllOf(classId string, objectType string) *VirtualizationCloudVmComputeConfigurationAllOf {
	this := VirtualizationCloudVmComputeConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCloudVmComputeConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmComputeConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCloudVmComputeConfigurationAllOfWithDefaults() *VirtualizationCloudVmComputeConfigurationAllOf {
	this := VirtualizationCloudVmComputeConfigurationAllOf{}
	var classId string = "virtualization.AwsVmComputeConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.AwsVmComputeConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInstanceTypeId returns the InstanceTypeId field value if set, zero value otherwise.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetInstanceTypeId() string {
	if o == nil || o.InstanceTypeId == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeId
}

// GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetInstanceTypeIdOk() (*string, bool) {
	if o == nil || o.InstanceTypeId == nil {
		return nil, false
	}
	return o.InstanceTypeId, true
}

// HasInstanceTypeId returns a boolean if a field has been set.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) HasInstanceTypeId() bool {
	if o != nil && o.InstanceTypeId != nil {
		return true
	}

	return false
}

// SetInstanceTypeId gets a reference to the given string and assigns it to the InstanceTypeId field.
func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetInstanceTypeId(v string) {
	o.InstanceTypeId = &v
}

func (o VirtualizationCloudVmComputeConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InstanceTypeId != nil {
		toSerialize["InstanceTypeId"] = o.InstanceTypeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCloudVmComputeConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationCloudVmComputeConfigurationAllOf := _VirtualizationCloudVmComputeConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationCloudVmComputeConfigurationAllOf); err == nil {
		*o = VirtualizationCloudVmComputeConfigurationAllOf(varVirtualizationCloudVmComputeConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InstanceTypeId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationCloudVmComputeConfigurationAllOf struct {
	value *VirtualizationCloudVmComputeConfigurationAllOf
	isSet bool
}

func (v NullableVirtualizationCloudVmComputeConfigurationAllOf) Get() *VirtualizationCloudVmComputeConfigurationAllOf {
	return v.value
}

func (v *NullableVirtualizationCloudVmComputeConfigurationAllOf) Set(val *VirtualizationCloudVmComputeConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCloudVmComputeConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCloudVmComputeConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCloudVmComputeConfigurationAllOf(val *VirtualizationCloudVmComputeConfigurationAllOf) *NullableVirtualizationCloudVmComputeConfigurationAllOf {
	return &NullableVirtualizationCloudVmComputeConfigurationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationCloudVmComputeConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCloudVmComputeConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}