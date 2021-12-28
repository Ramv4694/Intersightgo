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

// KubernetesBaremetalNodeProfileAllOf Definition of the list of properties defined in 'kubernetes.BaremetalNodeProfile', excluding properties defined in parent classes.
type KubernetesBaremetalNodeProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Network interface from NetworkInfo (by name) to use for kubernetes VIP.
	KubernetesNic        *string                                `json:"KubernetesNic,omitempty"`
	NetworkInfo          NullableKubernetesBaremetalNetworkInfo `json:"NetworkInfo,omitempty"`
	Server               *ComputeRackUnitRelationship           `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaremetalNodeProfileAllOf KubernetesBaremetalNodeProfileAllOf

// NewKubernetesBaremetalNodeProfileAllOf instantiates a new KubernetesBaremetalNodeProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaremetalNodeProfileAllOf(classId string, objectType string) *KubernetesBaremetalNodeProfileAllOf {
	this := KubernetesBaremetalNodeProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaremetalNodeProfileAllOfWithDefaults instantiates a new KubernetesBaremetalNodeProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaremetalNodeProfileAllOfWithDefaults() *KubernetesBaremetalNodeProfileAllOf {
	this := KubernetesBaremetalNodeProfileAllOf{}
	var classId string = "kubernetes.BaremetalNodeProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.BaremetalNodeProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaremetalNodeProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaremetalNodeProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaremetalNodeProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaremetalNodeProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKubernetesNic returns the KubernetesNic field value if set, zero value otherwise.
func (o *KubernetesBaremetalNodeProfileAllOf) GetKubernetesNic() string {
	if o == nil || o.KubernetesNic == nil {
		var ret string
		return ret
	}
	return *o.KubernetesNic
}

// GetKubernetesNicOk returns a tuple with the KubernetesNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) GetKubernetesNicOk() (*string, bool) {
	if o == nil || o.KubernetesNic == nil {
		return nil, false
	}
	return o.KubernetesNic, true
}

// HasKubernetesNic returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) HasKubernetesNic() bool {
	if o != nil && o.KubernetesNic != nil {
		return true
	}

	return false
}

// SetKubernetesNic gets a reference to the given string and assigns it to the KubernetesNic field.
func (o *KubernetesBaremetalNodeProfileAllOf) SetKubernetesNic(v string) {
	o.KubernetesNic = &v
}

// GetNetworkInfo returns the NetworkInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesBaremetalNodeProfileAllOf) GetNetworkInfo() KubernetesBaremetalNetworkInfo {
	if o == nil || o.NetworkInfo.Get() == nil {
		var ret KubernetesBaremetalNetworkInfo
		return ret
	}
	return *o.NetworkInfo.Get()
}

// GetNetworkInfoOk returns a tuple with the NetworkInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesBaremetalNodeProfileAllOf) GetNetworkInfoOk() (*KubernetesBaremetalNetworkInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkInfo.Get(), o.NetworkInfo.IsSet()
}

// HasNetworkInfo returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) HasNetworkInfo() bool {
	if o != nil && o.NetworkInfo.IsSet() {
		return true
	}

	return false
}

// SetNetworkInfo gets a reference to the given NullableKubernetesBaremetalNetworkInfo and assigns it to the NetworkInfo field.
func (o *KubernetesBaremetalNodeProfileAllOf) SetNetworkInfo(v KubernetesBaremetalNetworkInfo) {
	o.NetworkInfo.Set(&v)
}

// SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil
func (o *KubernetesBaremetalNodeProfileAllOf) SetNetworkInfoNil() {
	o.NetworkInfo.Set(nil)
}

// UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
func (o *KubernetesBaremetalNodeProfileAllOf) UnsetNetworkInfo() {
	o.NetworkInfo.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KubernetesBaremetalNodeProfileAllOf) GetServer() ComputeRackUnitRelationship {
	if o == nil || o.Server == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfileAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputeRackUnitRelationship and assigns it to the Server field.
func (o *KubernetesBaremetalNodeProfileAllOf) SetServer(v ComputeRackUnitRelationship) {
	o.Server = &v
}

func (o KubernetesBaremetalNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KubernetesNic != nil {
		toSerialize["KubernetesNic"] = o.KubernetesNic
	}
	if o.NetworkInfo.IsSet() {
		toSerialize["NetworkInfo"] = o.NetworkInfo.Get()
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaremetalNodeProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesBaremetalNodeProfileAllOf := _KubernetesBaremetalNodeProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesBaremetalNodeProfileAllOf); err == nil {
		*o = KubernetesBaremetalNodeProfileAllOf(varKubernetesBaremetalNodeProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KubernetesNic")
		delete(additionalProperties, "NetworkInfo")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesBaremetalNodeProfileAllOf struct {
	value *KubernetesBaremetalNodeProfileAllOf
	isSet bool
}

func (v NullableKubernetesBaremetalNodeProfileAllOf) Get() *KubernetesBaremetalNodeProfileAllOf {
	return v.value
}

func (v *NullableKubernetesBaremetalNodeProfileAllOf) Set(val *KubernetesBaremetalNodeProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaremetalNodeProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaremetalNodeProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaremetalNodeProfileAllOf(val *KubernetesBaremetalNodeProfileAllOf) *NullableKubernetesBaremetalNodeProfileAllOf {
	return &NullableKubernetesBaremetalNodeProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesBaremetalNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaremetalNodeProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
