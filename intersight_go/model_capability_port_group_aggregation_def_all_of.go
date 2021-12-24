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

// CapabilityPortGroupAggregationDefAllOf Definition of the list of properties defined in 'capability.PortGroupAggregationDef', excluding properties defined in parent classes.
type CapabilityPortGroupAggregationDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Aggregation capability for port group.
	AggregationCap *string `json:"AggregationCap,omitempty"`
	// Indicates support for 40G port group capability.
	Hw40GPortGroupCap *bool `json:"Hw40GPortGroupCap,omitempty"`
	// The type of port group for which this capability is defined.
	Pgtype               *string `json:"Pgtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityPortGroupAggregationDefAllOf CapabilityPortGroupAggregationDefAllOf

// NewCapabilityPortGroupAggregationDefAllOf instantiates a new CapabilityPortGroupAggregationDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityPortGroupAggregationDefAllOf(classId string, objectType string) *CapabilityPortGroupAggregationDefAllOf {
	this := CapabilityPortGroupAggregationDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityPortGroupAggregationDefAllOfWithDefaults instantiates a new CapabilityPortGroupAggregationDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityPortGroupAggregationDefAllOfWithDefaults() *CapabilityPortGroupAggregationDefAllOf {
	this := CapabilityPortGroupAggregationDefAllOf{}
	var classId string = "capability.PortGroupAggregationDef"
	this.ClassId = classId
	var objectType string = "capability.PortGroupAggregationDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityPortGroupAggregationDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityPortGroupAggregationDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityPortGroupAggregationDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityPortGroupAggregationDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregationCap returns the AggregationCap field value if set, zero value otherwise.
func (o *CapabilityPortGroupAggregationDefAllOf) GetAggregationCap() string {
	if o == nil || o.AggregationCap == nil {
		var ret string
		return ret
	}
	return *o.AggregationCap
}

// GetAggregationCapOk returns a tuple with the AggregationCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) GetAggregationCapOk() (*string, bool) {
	if o == nil || o.AggregationCap == nil {
		return nil, false
	}
	return o.AggregationCap, true
}

// HasAggregationCap returns a boolean if a field has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) HasAggregationCap() bool {
	if o != nil && o.AggregationCap != nil {
		return true
	}

	return false
}

// SetAggregationCap gets a reference to the given string and assigns it to the AggregationCap field.
func (o *CapabilityPortGroupAggregationDefAllOf) SetAggregationCap(v string) {
	o.AggregationCap = &v
}

// GetHw40GPortGroupCap returns the Hw40GPortGroupCap field value if set, zero value otherwise.
func (o *CapabilityPortGroupAggregationDefAllOf) GetHw40GPortGroupCap() bool {
	if o == nil || o.Hw40GPortGroupCap == nil {
		var ret bool
		return ret
	}
	return *o.Hw40GPortGroupCap
}

// GetHw40GPortGroupCapOk returns a tuple with the Hw40GPortGroupCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) GetHw40GPortGroupCapOk() (*bool, bool) {
	if o == nil || o.Hw40GPortGroupCap == nil {
		return nil, false
	}
	return o.Hw40GPortGroupCap, true
}

// HasHw40GPortGroupCap returns a boolean if a field has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) HasHw40GPortGroupCap() bool {
	if o != nil && o.Hw40GPortGroupCap != nil {
		return true
	}

	return false
}

// SetHw40GPortGroupCap gets a reference to the given bool and assigns it to the Hw40GPortGroupCap field.
func (o *CapabilityPortGroupAggregationDefAllOf) SetHw40GPortGroupCap(v bool) {
	o.Hw40GPortGroupCap = &v
}

// GetPgtype returns the Pgtype field value if set, zero value otherwise.
func (o *CapabilityPortGroupAggregationDefAllOf) GetPgtype() string {
	if o == nil || o.Pgtype == nil {
		var ret string
		return ret
	}
	return *o.Pgtype
}

// GetPgtypeOk returns a tuple with the Pgtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) GetPgtypeOk() (*string, bool) {
	if o == nil || o.Pgtype == nil {
		return nil, false
	}
	return o.Pgtype, true
}

// HasPgtype returns a boolean if a field has been set.
func (o *CapabilityPortGroupAggregationDefAllOf) HasPgtype() bool {
	if o != nil && o.Pgtype != nil {
		return true
	}

	return false
}

// SetPgtype gets a reference to the given string and assigns it to the Pgtype field.
func (o *CapabilityPortGroupAggregationDefAllOf) SetPgtype(v string) {
	o.Pgtype = &v
}

func (o CapabilityPortGroupAggregationDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregationCap != nil {
		toSerialize["AggregationCap"] = o.AggregationCap
	}
	if o.Hw40GPortGroupCap != nil {
		toSerialize["Hw40GPortGroupCap"] = o.Hw40GPortGroupCap
	}
	if o.Pgtype != nil {
		toSerialize["Pgtype"] = o.Pgtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityPortGroupAggregationDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityPortGroupAggregationDefAllOf := _CapabilityPortGroupAggregationDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityPortGroupAggregationDefAllOf); err == nil {
		*o = CapabilityPortGroupAggregationDefAllOf(varCapabilityPortGroupAggregationDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregationCap")
		delete(additionalProperties, "Hw40GPortGroupCap")
		delete(additionalProperties, "Pgtype")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityPortGroupAggregationDefAllOf struct {
	value *CapabilityPortGroupAggregationDefAllOf
	isSet bool
}

func (v NullableCapabilityPortGroupAggregationDefAllOf) Get() *CapabilityPortGroupAggregationDefAllOf {
	return v.value
}

func (v *NullableCapabilityPortGroupAggregationDefAllOf) Set(val *CapabilityPortGroupAggregationDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityPortGroupAggregationDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityPortGroupAggregationDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityPortGroupAggregationDefAllOf(val *CapabilityPortGroupAggregationDefAllOf) *NullableCapabilityPortGroupAggregationDefAllOf {
	return &NullableCapabilityPortGroupAggregationDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilityPortGroupAggregationDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityPortGroupAggregationDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
