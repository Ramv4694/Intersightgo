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

// PoolAbstractLeaseAllOf Definition of the list of properties defined in 'pool.AbstractLease', excluding properties defined in parent classes.
type PoolAbstractLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Type of the lease allocation either static or dynamic (i.e via pool). * `dynamic` - Identifiers to be allocated by system. * `static` - Identifiers are assigned by the user.
	AllocationType       *string `json:"AllocationType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractLeaseAllOf PoolAbstractLeaseAllOf

// NewPoolAbstractLeaseAllOf instantiates a new PoolAbstractLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractLeaseAllOf(classId string, objectType string) *PoolAbstractLeaseAllOf {
	this := PoolAbstractLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// NewPoolAbstractLeaseAllOfWithDefaults instantiates a new PoolAbstractLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractLeaseAllOfWithDefaults() *PoolAbstractLeaseAllOf {
	this := PoolAbstractLeaseAllOf{}
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocationType returns the AllocationType field value if set, zero value otherwise.
func (o *PoolAbstractLeaseAllOf) GetAllocationType() string {
	if o == nil || o.AllocationType == nil {
		var ret string
		return ret
	}
	return *o.AllocationType
}

// GetAllocationTypeOk returns a tuple with the AllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractLeaseAllOf) GetAllocationTypeOk() (*string, bool) {
	if o == nil || o.AllocationType == nil {
		return nil, false
	}
	return o.AllocationType, true
}

// HasAllocationType returns a boolean if a field has been set.
func (o *PoolAbstractLeaseAllOf) HasAllocationType() bool {
	if o != nil && o.AllocationType != nil {
		return true
	}

	return false
}

// SetAllocationType gets a reference to the given string and assigns it to the AllocationType field.
func (o *PoolAbstractLeaseAllOf) SetAllocationType(v string) {
	o.AllocationType = &v
}

func (o PoolAbstractLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocationType != nil {
		toSerialize["AllocationType"] = o.AllocationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractLeaseAllOf := _PoolAbstractLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractLeaseAllOf); err == nil {
		*o = PoolAbstractLeaseAllOf(varPoolAbstractLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractLeaseAllOf struct {
	value *PoolAbstractLeaseAllOf
	isSet bool
}

func (v NullablePoolAbstractLeaseAllOf) Get() *PoolAbstractLeaseAllOf {
	return v.value
}

func (v *NullablePoolAbstractLeaseAllOf) Set(val *PoolAbstractLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractLeaseAllOf(val *PoolAbstractLeaseAllOf) *NullablePoolAbstractLeaseAllOf {
	return &NullablePoolAbstractLeaseAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
