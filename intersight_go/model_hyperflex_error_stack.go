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
	"reflect"
	"strings"
)

// HyperflexErrorStack Reference to an error stack reported by this object.
type HyperflexErrorStack struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The error message string for this error stack.
	Message *string `json:"Message,omitempty"`
	// The error message ID for this error stack.
	MessageId            *int64 `json:"MessageId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexErrorStack HyperflexErrorStack

// NewHyperflexErrorStack instantiates a new HyperflexErrorStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexErrorStack(classId string, objectType string) *HyperflexErrorStack {
	this := HyperflexErrorStack{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexErrorStackWithDefaults instantiates a new HyperflexErrorStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexErrorStackWithDefaults() *HyperflexErrorStack {
	this := HyperflexErrorStack{}
	var classId string = "hyperflex.ErrorStack"
	this.ClassId = classId
	var objectType string = "hyperflex.ErrorStack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexErrorStack) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStack) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexErrorStack) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexErrorStack) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStack) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexErrorStack) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HyperflexErrorStack) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStack) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HyperflexErrorStack) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HyperflexErrorStack) SetMessage(v string) {
	o.Message = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *HyperflexErrorStack) GetMessageId() int64 {
	if o == nil || o.MessageId == nil {
		var ret int64
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStack) GetMessageIdOk() (*int64, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *HyperflexErrorStack) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int64 and assigns it to the MessageId field.
func (o *HyperflexErrorStack) SetMessageId(v int64) {
	o.MessageId = &v
}

func (o HyperflexErrorStack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.MessageId != nil {
		toSerialize["MessageId"] = o.MessageId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexErrorStack) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexErrorStackWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The error message string for this error stack.
		Message *string `json:"Message,omitempty"`
		// The error message ID for this error stack.
		MessageId *int64 `json:"MessageId,omitempty"`
	}

	varHyperflexErrorStackWithoutEmbeddedStruct := HyperflexErrorStackWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexErrorStackWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexErrorStack := _HyperflexErrorStack{}
		varHyperflexErrorStack.ClassId = varHyperflexErrorStackWithoutEmbeddedStruct.ClassId
		varHyperflexErrorStack.ObjectType = varHyperflexErrorStackWithoutEmbeddedStruct.ObjectType
		varHyperflexErrorStack.Message = varHyperflexErrorStackWithoutEmbeddedStruct.Message
		varHyperflexErrorStack.MessageId = varHyperflexErrorStackWithoutEmbeddedStruct.MessageId
		*o = HyperflexErrorStack(varHyperflexErrorStack)
	} else {
		return err
	}

	varHyperflexErrorStack := _HyperflexErrorStack{}

	err = json.Unmarshal(bytes, &varHyperflexErrorStack)
	if err == nil {
		o.MoBaseComplexType = varHyperflexErrorStack.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "MessageId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexErrorStack struct {
	value *HyperflexErrorStack
	isSet bool
}

func (v NullableHyperflexErrorStack) Get() *HyperflexErrorStack {
	return v.value
}

func (v *NullableHyperflexErrorStack) Set(val *HyperflexErrorStack) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexErrorStack) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexErrorStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexErrorStack(val *HyperflexErrorStack) *NullableHyperflexErrorStack {
	return &NullableHyperflexErrorStack{value: val, isSet: true}
}

func (v NullableHyperflexErrorStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexErrorStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
