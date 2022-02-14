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

// CloudAvailabilityZoneAllOf Definition of the list of properties defined in 'cloud.AvailabilityZone', excluding properties defined in parent classes.
type CloudAvailabilityZoneAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the availability zone.
	Name *string `json:"Name,omitempty"`
	// The ID of the availability zone.
	ZoneId               *string `json:"ZoneId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAvailabilityZoneAllOf CloudAvailabilityZoneAllOf

// NewCloudAvailabilityZoneAllOf instantiates a new CloudAvailabilityZoneAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAvailabilityZoneAllOf(classId string, objectType string) *CloudAvailabilityZoneAllOf {
	this := CloudAvailabilityZoneAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAvailabilityZoneAllOfWithDefaults instantiates a new CloudAvailabilityZoneAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAvailabilityZoneAllOfWithDefaults() *CloudAvailabilityZoneAllOf {
	this := CloudAvailabilityZoneAllOf{}
	var classId string = "cloud.AvailabilityZone"
	this.ClassId = classId
	var objectType string = "cloud.AvailabilityZone"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAvailabilityZoneAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAvailabilityZoneAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAvailabilityZoneAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAvailabilityZoneAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAvailabilityZoneAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAvailabilityZoneAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudAvailabilityZoneAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAvailabilityZoneAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudAvailabilityZoneAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudAvailabilityZoneAllOf) SetName(v string) {
	o.Name = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *CloudAvailabilityZoneAllOf) GetZoneId() string {
	if o == nil || o.ZoneId == nil {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAvailabilityZoneAllOf) GetZoneIdOk() (*string, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *CloudAvailabilityZoneAllOf) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *CloudAvailabilityZoneAllOf) SetZoneId(v string) {
	o.ZoneId = &v
}

func (o CloudAvailabilityZoneAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ZoneId != nil {
		toSerialize["ZoneId"] = o.ZoneId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAvailabilityZoneAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAvailabilityZoneAllOf := _CloudAvailabilityZoneAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAvailabilityZoneAllOf); err == nil {
		*o = CloudAvailabilityZoneAllOf(varCloudAvailabilityZoneAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ZoneId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAvailabilityZoneAllOf struct {
	value *CloudAvailabilityZoneAllOf
	isSet bool
}

func (v NullableCloudAvailabilityZoneAllOf) Get() *CloudAvailabilityZoneAllOf {
	return v.value
}

func (v *NullableCloudAvailabilityZoneAllOf) Set(val *CloudAvailabilityZoneAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAvailabilityZoneAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAvailabilityZoneAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAvailabilityZoneAllOf(val *CloudAvailabilityZoneAllOf) *NullableCloudAvailabilityZoneAllOf {
	return &NullableCloudAvailabilityZoneAllOf{value: val, isSet: true}
}

func (v NullableCloudAvailabilityZoneAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAvailabilityZoneAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
