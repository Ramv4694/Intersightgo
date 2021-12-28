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

// WorkflowJoinTaskAllOf Definition of the list of properties defined in 'workflow.JoinTask', excluding properties defined in parent classes.
type WorkflowJoinTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	JoinOnTasks []string `json:"JoinOnTasks,omitempty"`
	// Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	OnSuccess            *string `json:"OnSuccess,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowJoinTaskAllOf WorkflowJoinTaskAllOf

// NewWorkflowJoinTaskAllOf instantiates a new WorkflowJoinTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowJoinTaskAllOf(classId string, objectType string) *WorkflowJoinTaskAllOf {
	this := WorkflowJoinTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowJoinTaskAllOfWithDefaults instantiates a new WorkflowJoinTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowJoinTaskAllOfWithDefaults() *WorkflowJoinTaskAllOf {
	this := WorkflowJoinTaskAllOf{}
	var classId string = "workflow.JoinTask"
	this.ClassId = classId
	var objectType string = "workflow.JoinTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowJoinTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowJoinTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowJoinTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowJoinTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJoinOnTasks returns the JoinOnTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowJoinTaskAllOf) GetJoinOnTasks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.JoinOnTasks
}

// GetJoinOnTasksOk returns a tuple with the JoinOnTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowJoinTaskAllOf) GetJoinOnTasksOk() (*[]string, bool) {
	if o == nil || o.JoinOnTasks == nil {
		return nil, false
	}
	return &o.JoinOnTasks, true
}

// HasJoinOnTasks returns a boolean if a field has been set.
func (o *WorkflowJoinTaskAllOf) HasJoinOnTasks() bool {
	if o != nil && o.JoinOnTasks != nil {
		return true
	}

	return false
}

// SetJoinOnTasks gets a reference to the given []string and assigns it to the JoinOnTasks field.
func (o *WorkflowJoinTaskAllOf) SetJoinOnTasks(v []string) {
	o.JoinOnTasks = v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *WorkflowJoinTaskAllOf) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTaskAllOf) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *WorkflowJoinTaskAllOf) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *WorkflowJoinTaskAllOf) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

func (o WorkflowJoinTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.JoinOnTasks != nil {
		toSerialize["JoinOnTasks"] = o.JoinOnTasks
	}
	if o.OnSuccess != nil {
		toSerialize["OnSuccess"] = o.OnSuccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowJoinTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowJoinTaskAllOf := _WorkflowJoinTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowJoinTaskAllOf); err == nil {
		*o = WorkflowJoinTaskAllOf(varWorkflowJoinTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JoinOnTasks")
		delete(additionalProperties, "OnSuccess")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowJoinTaskAllOf struct {
	value *WorkflowJoinTaskAllOf
	isSet bool
}

func (v NullableWorkflowJoinTaskAllOf) Get() *WorkflowJoinTaskAllOf {
	return v.value
}

func (v *NullableWorkflowJoinTaskAllOf) Set(val *WorkflowJoinTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowJoinTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowJoinTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowJoinTaskAllOf(val *WorkflowJoinTaskAllOf) *NullableWorkflowJoinTaskAllOf {
	return &NullableWorkflowJoinTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowJoinTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowJoinTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
