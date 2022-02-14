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
)

// KubernetesNodeAllOf Definition of the list of properties defined in 'kubernetes.Node', excluding properties defined in parent classes.
type KubernetesNodeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	NodeAddresses        []KubernetesNodeAddress              `json:"NodeAddresses,omitempty"`
	NodeInfo             NullableKubernetesNodeInfo           `json:"NodeInfo,omitempty"`
	NodeSpec             NullableKubernetesNodeSpec           `json:"NodeSpec,omitempty"`
	NodeStatuses         []KubernetesNodeStatus               `json:"NodeStatuses,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeAllOf KubernetesNodeAllOf

// NewKubernetesNodeAllOf instantiates a new KubernetesNodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeAllOf(classId string, objectType string) *KubernetesNodeAllOf {
	this := KubernetesNodeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNodeAllOfWithDefaults instantiates a new KubernetesNodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeAllOfWithDefaults() *KubernetesNodeAllOf {
	this := KubernetesNodeAllOf{}
	var classId string = "kubernetes.Node"
	this.ClassId = classId
	var objectType string = "kubernetes.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNodeAddresses returns the NodeAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeAllOf) GetNodeAddresses() []KubernetesNodeAddress {
	if o == nil {
		var ret []KubernetesNodeAddress
		return ret
	}
	return o.NodeAddresses
}

// GetNodeAddressesOk returns a tuple with the NodeAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeAllOf) GetNodeAddressesOk() (*[]KubernetesNodeAddress, bool) {
	if o == nil || o.NodeAddresses == nil {
		return nil, false
	}
	return &o.NodeAddresses, true
}

// HasNodeAddresses returns a boolean if a field has been set.
func (o *KubernetesNodeAllOf) HasNodeAddresses() bool {
	if o != nil && o.NodeAddresses != nil {
		return true
	}

	return false
}

// SetNodeAddresses gets a reference to the given []KubernetesNodeAddress and assigns it to the NodeAddresses field.
func (o *KubernetesNodeAllOf) SetNodeAddresses(v []KubernetesNodeAddress) {
	o.NodeAddresses = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeAllOf) GetNodeInfo() KubernetesNodeInfo {
	if o == nil || o.NodeInfo.Get() == nil {
		var ret KubernetesNodeInfo
		return ret
	}
	return *o.NodeInfo.Get()
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeAllOf) GetNodeInfoOk() (*KubernetesNodeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeInfo.Get(), o.NodeInfo.IsSet()
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *KubernetesNodeAllOf) HasNodeInfo() bool {
	if o != nil && o.NodeInfo.IsSet() {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given NullableKubernetesNodeInfo and assigns it to the NodeInfo field.
func (o *KubernetesNodeAllOf) SetNodeInfo(v KubernetesNodeInfo) {
	o.NodeInfo.Set(&v)
}

// SetNodeInfoNil sets the value for NodeInfo to be an explicit nil
func (o *KubernetesNodeAllOf) SetNodeInfoNil() {
	o.NodeInfo.Set(nil)
}

// UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
func (o *KubernetesNodeAllOf) UnsetNodeInfo() {
	o.NodeInfo.Unset()
}

// GetNodeSpec returns the NodeSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeAllOf) GetNodeSpec() KubernetesNodeSpec {
	if o == nil || o.NodeSpec.Get() == nil {
		var ret KubernetesNodeSpec
		return ret
	}
	return *o.NodeSpec.Get()
}

// GetNodeSpecOk returns a tuple with the NodeSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeAllOf) GetNodeSpecOk() (*KubernetesNodeSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeSpec.Get(), o.NodeSpec.IsSet()
}

// HasNodeSpec returns a boolean if a field has been set.
func (o *KubernetesNodeAllOf) HasNodeSpec() bool {
	if o != nil && o.NodeSpec.IsSet() {
		return true
	}

	return false
}

// SetNodeSpec gets a reference to the given NullableKubernetesNodeSpec and assigns it to the NodeSpec field.
func (o *KubernetesNodeAllOf) SetNodeSpec(v KubernetesNodeSpec) {
	o.NodeSpec.Set(&v)
}

// SetNodeSpecNil sets the value for NodeSpec to be an explicit nil
func (o *KubernetesNodeAllOf) SetNodeSpecNil() {
	o.NodeSpec.Set(nil)
}

// UnsetNodeSpec ensures that no value is present for NodeSpec, not even an explicit nil
func (o *KubernetesNodeAllOf) UnsetNodeSpec() {
	o.NodeSpec.Unset()
}

// GetNodeStatuses returns the NodeStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeAllOf) GetNodeStatuses() []KubernetesNodeStatus {
	if o == nil {
		var ret []KubernetesNodeStatus
		return ret
	}
	return o.NodeStatuses
}

// GetNodeStatusesOk returns a tuple with the NodeStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeAllOf) GetNodeStatusesOk() (*[]KubernetesNodeStatus, bool) {
	if o == nil || o.NodeStatuses == nil {
		return nil, false
	}
	return &o.NodeStatuses, true
}

// HasNodeStatuses returns a boolean if a field has been set.
func (o *KubernetesNodeAllOf) HasNodeStatuses() bool {
	if o != nil && o.NodeStatuses != nil {
		return true
	}

	return false
}

// SetNodeStatuses gets a reference to the given []KubernetesNodeStatus and assigns it to the NodeStatuses field.
func (o *KubernetesNodeAllOf) SetNodeStatuses(v []KubernetesNodeStatus) {
	o.NodeStatuses = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesNodeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesNodeAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesNodeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesNodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NodeAddresses != nil {
		toSerialize["NodeAddresses"] = o.NodeAddresses
	}
	if o.NodeInfo.IsSet() {
		toSerialize["NodeInfo"] = o.NodeInfo.Get()
	}
	if o.NodeSpec.IsSet() {
		toSerialize["NodeSpec"] = o.NodeSpec.Get()
	}
	if o.NodeStatuses != nil {
		toSerialize["NodeStatuses"] = o.NodeStatuses
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesNodeAllOf := _KubernetesNodeAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesNodeAllOf); err == nil {
		*o = KubernetesNodeAllOf(varKubernetesNodeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeAddresses")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "NodeSpec")
		delete(additionalProperties, "NodeStatuses")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesNodeAllOf struct {
	value *KubernetesNodeAllOf
	isSet bool
}

func (v NullableKubernetesNodeAllOf) Get() *KubernetesNodeAllOf {
	return v.value
}

func (v *NullableKubernetesNodeAllOf) Set(val *KubernetesNodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeAllOf(val *KubernetesNodeAllOf) *NullableKubernetesNodeAllOf {
	return &NullableKubernetesNodeAllOf{value: val, isSet: true}
}

func (v NullableKubernetesNodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
