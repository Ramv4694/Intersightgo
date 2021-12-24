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

// CloudCollectInventoryAllOf Definition of the list of properties defined in 'cloud.CollectInventory', excluding properties defined in parent classes.
type CloudCollectInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the new Terraform cloud asset which was created.
	TargetId             *string                  `json:"TargetId,omitempty"`
	Target               *AssetTargetRelationship `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudCollectInventoryAllOf CloudCollectInventoryAllOf

// NewCloudCollectInventoryAllOf instantiates a new CloudCollectInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCollectInventoryAllOf(classId string, objectType string) *CloudCollectInventoryAllOf {
	this := CloudCollectInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudCollectInventoryAllOfWithDefaults instantiates a new CloudCollectInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCollectInventoryAllOfWithDefaults() *CloudCollectInventoryAllOf {
	this := CloudCollectInventoryAllOf{}
	var classId string = "cloud.CollectInventory"
	this.ClassId = classId
	var objectType string = "cloud.CollectInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudCollectInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudCollectInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudCollectInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudCollectInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudCollectInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudCollectInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CloudCollectInventoryAllOf) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCollectInventoryAllOf) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CloudCollectInventoryAllOf) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *CloudCollectInventoryAllOf) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CloudCollectInventoryAllOf) GetTarget() AssetTargetRelationship {
	if o == nil || o.Target == nil {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCollectInventoryAllOf) GetTargetOk() (*AssetTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CloudCollectInventoryAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetTargetRelationship and assigns it to the Target field.
func (o *CloudCollectInventoryAllOf) SetTarget(v AssetTargetRelationship) {
	o.Target = &v
}

func (o CloudCollectInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudCollectInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudCollectInventoryAllOf := _CloudCollectInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varCloudCollectInventoryAllOf); err == nil {
		*o = CloudCollectInventoryAllOf(varCloudCollectInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetId")
		delete(additionalProperties, "Target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudCollectInventoryAllOf struct {
	value *CloudCollectInventoryAllOf
	isSet bool
}

func (v NullableCloudCollectInventoryAllOf) Get() *CloudCollectInventoryAllOf {
	return v.value
}

func (v *NullableCloudCollectInventoryAllOf) Set(val *CloudCollectInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCollectInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCollectInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCollectInventoryAllOf(val *CloudCollectInventoryAllOf) *NullableCloudCollectInventoryAllOf {
	return &NullableCloudCollectInventoryAllOf{value: val, isSet: true}
}

func (v NullableCloudCollectInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCollectInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
