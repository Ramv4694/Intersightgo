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
)

// VirtualizationEsxiConsoleAllOf Definition of the list of properties defined in 'virtualization.EsxiConsole', excluding properties defined in parent classes.
type VirtualizationEsxiConsoleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The stream ID of the host console session opened.
	StreamId             *string                               `json:"StreamId,omitempty"`
	DeviceRegistration   *AssetDeviceRegistrationRelationship  `json:"DeviceRegistration,omitempty"`
	Host                 *VirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiConsoleAllOf VirtualizationEsxiConsoleAllOf

// NewVirtualizationEsxiConsoleAllOf instantiates a new VirtualizationEsxiConsoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiConsoleAllOf(classId string, objectType string) *VirtualizationEsxiConsoleAllOf {
	this := VirtualizationEsxiConsoleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiConsoleAllOfWithDefaults instantiates a new VirtualizationEsxiConsoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiConsoleAllOfWithDefaults() *VirtualizationEsxiConsoleAllOf {
	this := VirtualizationEsxiConsoleAllOf{}
	var classId string = "virtualization.EsxiConsole"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiConsole"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiConsoleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiConsoleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiConsoleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiConsoleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiConsoleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiConsoleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStreamId returns the StreamId field value if set, zero value otherwise.
func (o *VirtualizationEsxiConsoleAllOf) GetStreamId() string {
	if o == nil || o.StreamId == nil {
		var ret string
		return ret
	}
	return *o.StreamId
}

// GetStreamIdOk returns a tuple with the StreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiConsoleAllOf) GetStreamIdOk() (*string, bool) {
	if o == nil || o.StreamId == nil {
		return nil, false
	}
	return o.StreamId, true
}

// HasStreamId returns a boolean if a field has been set.
func (o *VirtualizationEsxiConsoleAllOf) HasStreamId() bool {
	if o != nil && o.StreamId != nil {
		return true
	}

	return false
}

// SetStreamId gets a reference to the given string and assigns it to the StreamId field.
func (o *VirtualizationEsxiConsoleAllOf) SetStreamId(v string) {
	o.StreamId = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *VirtualizationEsxiConsoleAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiConsoleAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *VirtualizationEsxiConsoleAllOf) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *VirtualizationEsxiConsoleAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationEsxiConsoleAllOf) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiConsoleAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationEsxiConsoleAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationEsxiConsoleAllOf) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

func (o VirtualizationEsxiConsoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.StreamId != nil {
		toSerialize["StreamId"] = o.StreamId
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiConsoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationEsxiConsoleAllOf := _VirtualizationEsxiConsoleAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationEsxiConsoleAllOf); err == nil {
		*o = VirtualizationEsxiConsoleAllOf(varVirtualizationEsxiConsoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StreamId")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationEsxiConsoleAllOf struct {
	value *VirtualizationEsxiConsoleAllOf
	isSet bool
}

func (v NullableVirtualizationEsxiConsoleAllOf) Get() *VirtualizationEsxiConsoleAllOf {
	return v.value
}

func (v *NullableVirtualizationEsxiConsoleAllOf) Set(val *VirtualizationEsxiConsoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiConsoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiConsoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiConsoleAllOf(val *VirtualizationEsxiConsoleAllOf) *NullableVirtualizationEsxiConsoleAllOf {
	return &NullableVirtualizationEsxiConsoleAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiConsoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiConsoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
