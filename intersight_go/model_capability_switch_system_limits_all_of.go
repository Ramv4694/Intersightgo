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

// CapabilitySwitchSystemLimitsAllOf Definition of the list of properties defined in 'capability.SwitchSystemLimits', excluding properties defined in parent classes.
type CapabilitySwitchSystemLimitsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.
	MaximumChassisCount *int64 `json:"MaximumChassisCount,omitempty"`
	// Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.
	MaximumFexPerDomain *int64 `json:"MaximumFexPerDomain,omitempty"`
	// Maximum UCS servers per Switch/Fabric-Interconnect.
	MaximumServersPerDomain *int64 `json:"MaximumServersPerDomain,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _CapabilitySwitchSystemLimitsAllOf CapabilitySwitchSystemLimitsAllOf

// NewCapabilitySwitchSystemLimitsAllOf instantiates a new CapabilitySwitchSystemLimitsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchSystemLimitsAllOf(classId string, objectType string) *CapabilitySwitchSystemLimitsAllOf {
	this := CapabilitySwitchSystemLimitsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySwitchSystemLimitsAllOfWithDefaults instantiates a new CapabilitySwitchSystemLimitsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchSystemLimitsAllOfWithDefaults() *CapabilitySwitchSystemLimitsAllOf {
	this := CapabilitySwitchSystemLimitsAllOf{}
	var classId string = "capability.SwitchSystemLimits"
	this.ClassId = classId
	var objectType string = "capability.SwitchSystemLimits"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchSystemLimitsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchSystemLimitsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchSystemLimitsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchSystemLimitsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaximumChassisCount returns the MaximumChassisCount field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumChassisCount() int64 {
	if o == nil || o.MaximumChassisCount == nil {
		var ret int64
		return ret
	}
	return *o.MaximumChassisCount
}

// GetMaximumChassisCountOk returns a tuple with the MaximumChassisCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumChassisCountOk() (*int64, bool) {
	if o == nil || o.MaximumChassisCount == nil {
		return nil, false
	}
	return o.MaximumChassisCount, true
}

// HasMaximumChassisCount returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) HasMaximumChassisCount() bool {
	if o != nil && o.MaximumChassisCount != nil {
		return true
	}

	return false
}

// SetMaximumChassisCount gets a reference to the given int64 and assigns it to the MaximumChassisCount field.
func (o *CapabilitySwitchSystemLimitsAllOf) SetMaximumChassisCount(v int64) {
	o.MaximumChassisCount = &v
}

// GetMaximumFexPerDomain returns the MaximumFexPerDomain field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumFexPerDomain() int64 {
	if o == nil || o.MaximumFexPerDomain == nil {
		var ret int64
		return ret
	}
	return *o.MaximumFexPerDomain
}

// GetMaximumFexPerDomainOk returns a tuple with the MaximumFexPerDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumFexPerDomainOk() (*int64, bool) {
	if o == nil || o.MaximumFexPerDomain == nil {
		return nil, false
	}
	return o.MaximumFexPerDomain, true
}

// HasMaximumFexPerDomain returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) HasMaximumFexPerDomain() bool {
	if o != nil && o.MaximumFexPerDomain != nil {
		return true
	}

	return false
}

// SetMaximumFexPerDomain gets a reference to the given int64 and assigns it to the MaximumFexPerDomain field.
func (o *CapabilitySwitchSystemLimitsAllOf) SetMaximumFexPerDomain(v int64) {
	o.MaximumFexPerDomain = &v
}

// GetMaximumServersPerDomain returns the MaximumServersPerDomain field value if set, zero value otherwise.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumServersPerDomain() int64 {
	if o == nil || o.MaximumServersPerDomain == nil {
		var ret int64
		return ret
	}
	return *o.MaximumServersPerDomain
}

// GetMaximumServersPerDomainOk returns a tuple with the MaximumServersPerDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) GetMaximumServersPerDomainOk() (*int64, bool) {
	if o == nil || o.MaximumServersPerDomain == nil {
		return nil, false
	}
	return o.MaximumServersPerDomain, true
}

// HasMaximumServersPerDomain returns a boolean if a field has been set.
func (o *CapabilitySwitchSystemLimitsAllOf) HasMaximumServersPerDomain() bool {
	if o != nil && o.MaximumServersPerDomain != nil {
		return true
	}

	return false
}

// SetMaximumServersPerDomain gets a reference to the given int64 and assigns it to the MaximumServersPerDomain field.
func (o *CapabilitySwitchSystemLimitsAllOf) SetMaximumServersPerDomain(v int64) {
	o.MaximumServersPerDomain = &v
}

func (o CapabilitySwitchSystemLimitsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaximumChassisCount != nil {
		toSerialize["MaximumChassisCount"] = o.MaximumChassisCount
	}
	if o.MaximumFexPerDomain != nil {
		toSerialize["MaximumFexPerDomain"] = o.MaximumFexPerDomain
	}
	if o.MaximumServersPerDomain != nil {
		toSerialize["MaximumServersPerDomain"] = o.MaximumServersPerDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchSystemLimitsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySwitchSystemLimitsAllOf := _CapabilitySwitchSystemLimitsAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySwitchSystemLimitsAllOf); err == nil {
		*o = CapabilitySwitchSystemLimitsAllOf(varCapabilitySwitchSystemLimitsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaximumChassisCount")
		delete(additionalProperties, "MaximumFexPerDomain")
		delete(additionalProperties, "MaximumServersPerDomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchSystemLimitsAllOf struct {
	value *CapabilitySwitchSystemLimitsAllOf
	isSet bool
}

func (v NullableCapabilitySwitchSystemLimitsAllOf) Get() *CapabilitySwitchSystemLimitsAllOf {
	return v.value
}

func (v *NullableCapabilitySwitchSystemLimitsAllOf) Set(val *CapabilitySwitchSystemLimitsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchSystemLimitsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchSystemLimitsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchSystemLimitsAllOf(val *CapabilitySwitchSystemLimitsAllOf) *NullableCapabilitySwitchSystemLimitsAllOf {
	return &NullableCapabilitySwitchSystemLimitsAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySwitchSystemLimitsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchSystemLimitsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
