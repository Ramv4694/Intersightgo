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

// BulkRestSubRequestAllOf Definition of the list of properties defined in 'bulk.RestSubRequest', excluding properties defined in parent classes.
type BulkRestSubRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string    `json:"ObjectType"`
	Body       *MoBaseMo `json:"Body,omitempty"`
	// Used with PATCH & DELETE actions. The moid of an existing object instance.
	TargetMoid           *string `json:"TargetMoid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRestSubRequestAllOf BulkRestSubRequestAllOf

// NewBulkRestSubRequestAllOf instantiates a new BulkRestSubRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRestSubRequestAllOf(classId string, objectType string) *BulkRestSubRequestAllOf {
	this := BulkRestSubRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkRestSubRequestAllOfWithDefaults instantiates a new BulkRestSubRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRestSubRequestAllOfWithDefaults() *BulkRestSubRequestAllOf {
	this := BulkRestSubRequestAllOf{}
	var classId string = "bulk.RestSubRequest"
	this.ClassId = classId
	var objectType string = "bulk.RestSubRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRestSubRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRestSubRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkRestSubRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRestSubRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BulkRestSubRequestAllOf) GetBody() MoBaseMo {
	if o == nil || o.Body == nil {
		var ret MoBaseMo
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequestAllOf) GetBodyOk() (*MoBaseMo, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BulkRestSubRequestAllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given MoBaseMo and assigns it to the Body field.
func (o *BulkRestSubRequestAllOf) SetBody(v MoBaseMo) {
	o.Body = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *BulkRestSubRequestAllOf) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestSubRequestAllOf) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *BulkRestSubRequestAllOf) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *BulkRestSubRequestAllOf) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

func (o BulkRestSubRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkRestSubRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkRestSubRequestAllOf := _BulkRestSubRequestAllOf{}

	if err = json.Unmarshal(bytes, &varBulkRestSubRequestAllOf); err == nil {
		*o = BulkRestSubRequestAllOf(varBulkRestSubRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "TargetMoid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkRestSubRequestAllOf struct {
	value *BulkRestSubRequestAllOf
	isSet bool
}

func (v NullableBulkRestSubRequestAllOf) Get() *BulkRestSubRequestAllOf {
	return v.value
}

func (v *NullableBulkRestSubRequestAllOf) Set(val *BulkRestSubRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRestSubRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRestSubRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRestSubRequestAllOf(val *BulkRestSubRequestAllOf) *NullableBulkRestSubRequestAllOf {
	return &NullableBulkRestSubRequestAllOf{value: val, isSet: true}
}

func (v NullableBulkRestSubRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRestSubRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}