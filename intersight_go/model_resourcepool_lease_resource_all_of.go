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

// ResourcepoolLeaseResourceAllOf Definition of the list of properties defined in 'resourcepool.LeaseResource', excluding properties defined in parent classes.
type ResourcepoolLeaseResourceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Lease opertion applied for the feature.
	Feature              *string               `json:"Feature,omitempty"`
	Resource             *MoBaseMoRelationship `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolLeaseResourceAllOf ResourcepoolLeaseResourceAllOf

// NewResourcepoolLeaseResourceAllOf instantiates a new ResourcepoolLeaseResourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolLeaseResourceAllOf(classId string, objectType string) *ResourcepoolLeaseResourceAllOf {
	this := ResourcepoolLeaseResourceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourcepoolLeaseResourceAllOfWithDefaults instantiates a new ResourcepoolLeaseResourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolLeaseResourceAllOfWithDefaults() *ResourcepoolLeaseResourceAllOf {
	this := ResourcepoolLeaseResourceAllOf{}
	var classId string = "resourcepool.LeaseResource"
	this.ClassId = classId
	var objectType string = "resourcepool.LeaseResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolLeaseResourceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseResourceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolLeaseResourceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolLeaseResourceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseResourceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolLeaseResourceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ResourcepoolLeaseResourceAllOf) GetFeature() string {
	if o == nil || o.Feature == nil {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseResourceAllOf) GetFeatureOk() (*string, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *ResourcepoolLeaseResourceAllOf) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ResourcepoolLeaseResourceAllOf) SetFeature(v string) {
	o.Feature = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourcepoolLeaseResourceAllOf) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseResourceAllOf) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourcepoolLeaseResourceAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourcepoolLeaseResourceAllOf) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o ResourcepoolLeaseResourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Feature != nil {
		toSerialize["Feature"] = o.Feature
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolLeaseResourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcepoolLeaseResourceAllOf := _ResourcepoolLeaseResourceAllOf{}

	if err = json.Unmarshal(bytes, &varResourcepoolLeaseResourceAllOf); err == nil {
		*o = ResourcepoolLeaseResourceAllOf(varResourcepoolLeaseResourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Feature")
		delete(additionalProperties, "Resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcepoolLeaseResourceAllOf struct {
	value *ResourcepoolLeaseResourceAllOf
	isSet bool
}

func (v NullableResourcepoolLeaseResourceAllOf) Get() *ResourcepoolLeaseResourceAllOf {
	return v.value
}

func (v *NullableResourcepoolLeaseResourceAllOf) Set(val *ResourcepoolLeaseResourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolLeaseResourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolLeaseResourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolLeaseResourceAllOf(val *ResourcepoolLeaseResourceAllOf) *NullableResourcepoolLeaseResourceAllOf {
	return &NullableResourcepoolLeaseResourceAllOf{value: val, isSet: true}
}

func (v NullableResourcepoolLeaseResourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolLeaseResourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}