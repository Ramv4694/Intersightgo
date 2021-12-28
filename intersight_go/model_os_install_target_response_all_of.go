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

// OsInstallTargetResponseAllOf Definition of the list of properties defined in 'os.InstallTargetResponse', excluding properties defined in parent classes.
type OsInstallTargetResponseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType           string   `json:"ObjectType"`
	SourceMo             *MoMoRef `json:"SourceMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsInstallTargetResponseAllOf OsInstallTargetResponseAllOf

// NewOsInstallTargetResponseAllOf instantiates a new OsInstallTargetResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsInstallTargetResponseAllOf(classId string, objectType string) *OsInstallTargetResponseAllOf {
	this := OsInstallTargetResponseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsInstallTargetResponseAllOfWithDefaults instantiates a new OsInstallTargetResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsInstallTargetResponseAllOfWithDefaults() *OsInstallTargetResponseAllOf {
	this := OsInstallTargetResponseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsInstallTargetResponseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsInstallTargetResponseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsInstallTargetResponseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsInstallTargetResponseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsInstallTargetResponseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsInstallTargetResponseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSourceMo returns the SourceMo field value if set, zero value otherwise.
func (o *OsInstallTargetResponseAllOf) GetSourceMo() MoMoRef {
	if o == nil || o.SourceMo == nil {
		var ret MoMoRef
		return ret
	}
	return *o.SourceMo
}

// GetSourceMoOk returns a tuple with the SourceMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallTargetResponseAllOf) GetSourceMoOk() (*MoMoRef, bool) {
	if o == nil || o.SourceMo == nil {
		return nil, false
	}
	return o.SourceMo, true
}

// HasSourceMo returns a boolean if a field has been set.
func (o *OsInstallTargetResponseAllOf) HasSourceMo() bool {
	if o != nil && o.SourceMo != nil {
		return true
	}

	return false
}

// SetSourceMo gets a reference to the given MoMoRef and assigns it to the SourceMo field.
func (o *OsInstallTargetResponseAllOf) SetSourceMo(v MoMoRef) {
	o.SourceMo = &v
}

func (o OsInstallTargetResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SourceMo != nil {
		toSerialize["SourceMo"] = o.SourceMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsInstallTargetResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsInstallTargetResponseAllOf := _OsInstallTargetResponseAllOf{}

	if err = json.Unmarshal(bytes, &varOsInstallTargetResponseAllOf); err == nil {
		*o = OsInstallTargetResponseAllOf(varOsInstallTargetResponseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SourceMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsInstallTargetResponseAllOf struct {
	value *OsInstallTargetResponseAllOf
	isSet bool
}

func (v NullableOsInstallTargetResponseAllOf) Get() *OsInstallTargetResponseAllOf {
	return v.value
}

func (v *NullableOsInstallTargetResponseAllOf) Set(val *OsInstallTargetResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsInstallTargetResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsInstallTargetResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsInstallTargetResponseAllOf(val *OsInstallTargetResponseAllOf) *NullableOsInstallTargetResponseAllOf {
	return &NullableOsInstallTargetResponseAllOf{value: val, isSet: true}
}

func (v NullableOsInstallTargetResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsInstallTargetResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
