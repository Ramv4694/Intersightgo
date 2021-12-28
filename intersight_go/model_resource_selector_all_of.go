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

// ResourceSelectorAllOf Definition of the list of properties defined in 'resource.Selector', excluding properties defined in parent classes.
type ResourceSelectorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
	Selector             *string `json:"Selector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorAllOf ResourceSelectorAllOf

// NewResourceSelectorAllOf instantiates a new ResourceSelectorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorAllOf(classId string, objectType string) *ResourceSelectorAllOf {
	this := ResourceSelectorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceSelectorAllOfWithDefaults instantiates a new ResourceSelectorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorAllOfWithDefaults() *ResourceSelectorAllOf {
	this := ResourceSelectorAllOf{}
	var classId string = "resource.Selector"
	this.ClassId = classId
	var objectType string = "resource.Selector"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceSelectorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceSelectorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceSelectorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceSelectorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceSelectorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceSelectorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ResourceSelectorAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ResourceSelectorAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *ResourceSelectorAllOf) SetSelector(v string) {
	o.Selector = &v
}

func (o ResourceSelectorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorAllOf := _ResourceSelectorAllOf{}

	if err = json.Unmarshal(bytes, &varResourceSelectorAllOf); err == nil {
		*o = ResourceSelectorAllOf(varResourceSelectorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Selector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSelectorAllOf struct {
	value *ResourceSelectorAllOf
	isSet bool
}

func (v NullableResourceSelectorAllOf) Get() *ResourceSelectorAllOf {
	return v.value
}

func (v *NullableResourceSelectorAllOf) Set(val *ResourceSelectorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorAllOf(val *ResourceSelectorAllOf) *NullableResourceSelectorAllOf {
	return &NullableResourceSelectorAllOf{value: val, isSet: true}
}

func (v NullableResourceSelectorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
