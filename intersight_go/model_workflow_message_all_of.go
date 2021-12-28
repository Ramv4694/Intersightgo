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

// WorkflowMessageAllOf Definition of the list of properties defined in 'workflow.Message', excluding properties defined in parent classes.
type WorkflowMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An i18n message that can be translated in multiple languages to support internationalization.
	Message *string `json:"Message,omitempty"`
	// The severity of the Task or Workflow message warning/error/info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
	Severity             *string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMessageAllOf WorkflowMessageAllOf

// NewWorkflowMessageAllOf instantiates a new WorkflowMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMessageAllOf(classId string, objectType string) *WorkflowMessageAllOf {
	this := WorkflowMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var severity string = "Info"
	this.Severity = &severity
	return &this
}

// NewWorkflowMessageAllOfWithDefaults instantiates a new WorkflowMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMessageAllOfWithDefaults() *WorkflowMessageAllOf {
	this := WorkflowMessageAllOf{}
	var classId string = "workflow.Message"
	this.ClassId = classId
	var objectType string = "workflow.Message"
	this.ObjectType = objectType
	var severity string = "Info"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowMessageAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessageAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowMessageAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WorkflowMessageAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *WorkflowMessageAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessageAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *WorkflowMessageAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *WorkflowMessageAllOf) SetSeverity(v string) {
	o.Severity = &v
}

func (o WorkflowMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMessageAllOf := _WorkflowMessageAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowMessageAllOf); err == nil {
		*o = WorkflowMessageAllOf(varWorkflowMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Severity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowMessageAllOf struct {
	value *WorkflowMessageAllOf
	isSet bool
}

func (v NullableWorkflowMessageAllOf) Get() *WorkflowMessageAllOf {
	return v.value
}

func (v *NullableWorkflowMessageAllOf) Set(val *WorkflowMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMessageAllOf(val *WorkflowMessageAllOf) *NullableWorkflowMessageAllOf {
	return &NullableWorkflowMessageAllOf{value: val, isSet: true}
}

func (v NullableWorkflowMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
