/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-10-09T21:18:32Z.
 *
 * API version: 1.0.9-4809
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowTaskConstraintsAllOf Definition of the list of properties defined in 'workflow.TaskConstraints', excluding properties defined in parent classes.
type WorkflowTaskConstraintsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// List of property constraints that helps to narrow down task implementations based on target device input.
	TargetDataType       interface{} `json:"TargetDataType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskConstraintsAllOf WorkflowTaskConstraintsAllOf

// NewWorkflowTaskConstraintsAllOf instantiates a new WorkflowTaskConstraintsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskConstraintsAllOf(classId string, objectType string) *WorkflowTaskConstraintsAllOf {
	this := WorkflowTaskConstraintsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskConstraintsAllOfWithDefaults instantiates a new WorkflowTaskConstraintsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskConstraintsAllOfWithDefaults() *WorkflowTaskConstraintsAllOf {
	this := WorkflowTaskConstraintsAllOf{}
	var classId string = "workflow.TaskConstraints"
	this.ClassId = classId
	var objectType string = "workflow.TaskConstraints"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskConstraintsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskConstraintsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskConstraintsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskConstraintsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskConstraintsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskConstraintsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetDataType returns the TargetDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskConstraintsAllOf) GetTargetDataType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetDataType
}

// GetTargetDataTypeOk returns a tuple with the TargetDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskConstraintsAllOf) GetTargetDataTypeOk() (*interface{}, bool) {
	if o == nil || o.TargetDataType == nil {
		return nil, false
	}
	return &o.TargetDataType, true
}

// HasTargetDataType returns a boolean if a field has been set.
func (o *WorkflowTaskConstraintsAllOf) HasTargetDataType() bool {
	if o != nil && o.TargetDataType != nil {
		return true
	}

	return false
}

// SetTargetDataType gets a reference to the given interface{} and assigns it to the TargetDataType field.
func (o *WorkflowTaskConstraintsAllOf) SetTargetDataType(v interface{}) {
	o.TargetDataType = v
}

func (o WorkflowTaskConstraintsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetDataType != nil {
		toSerialize["TargetDataType"] = o.TargetDataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskConstraintsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTaskConstraintsAllOf := _WorkflowTaskConstraintsAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTaskConstraintsAllOf); err == nil {
		*o = WorkflowTaskConstraintsAllOf(varWorkflowTaskConstraintsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetDataType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskConstraintsAllOf struct {
	value *WorkflowTaskConstraintsAllOf
	isSet bool
}

func (v NullableWorkflowTaskConstraintsAllOf) Get() *WorkflowTaskConstraintsAllOf {
	return v.value
}

func (v *NullableWorkflowTaskConstraintsAllOf) Set(val *WorkflowTaskConstraintsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskConstraintsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskConstraintsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskConstraintsAllOf(val *WorkflowTaskConstraintsAllOf) *NullableWorkflowTaskConstraintsAllOf {
	return &NullableWorkflowTaskConstraintsAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskConstraintsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskConstraintsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
