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

// TaskWorkflowActionAllOf Definition of the list of properties defined in 'task.WorkflowAction', excluding properties defined in parent classes.
type TaskWorkflowActionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action for test workflow. * `start` - Start action for the workflow. * `stop` - Stop action for the workflow.
	Action *string `json:"Action,omitempty"`
	// Json formatted string input parameters to workflow.
	InputParams *string `json:"InputParams,omitempty"`
	// When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow.
	IsDynamic *bool `json:"IsDynamic,omitempty"`
	// When true, the workflow will wait for previous one to complete before starting a new one.
	WaitOnDuplicate *bool `json:"WaitOnDuplicate,omitempty"`
	// Json formatted string that has the contents of the workflow context used when starting a workflow.
	WorkflowContext      *string                      `json:"WorkflowContext,omitempty"`
	WorkflowFile         NullableTaskFileDownloadInfo `json:"WorkflowFile,omitempty"`
	Account              *IamAccountRelationship      `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskWorkflowActionAllOf TaskWorkflowActionAllOf

// NewTaskWorkflowActionAllOf instantiates a new TaskWorkflowActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskWorkflowActionAllOf(classId string, objectType string) *TaskWorkflowActionAllOf {
	this := TaskWorkflowActionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "start"
	this.Action = &action
	return &this
}

// NewTaskWorkflowActionAllOfWithDefaults instantiates a new TaskWorkflowActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWorkflowActionAllOfWithDefaults() *TaskWorkflowActionAllOf {
	this := TaskWorkflowActionAllOf{}
	var classId string = "task.WorkflowAction"
	this.ClassId = classId
	var objectType string = "task.WorkflowAction"
	this.ObjectType = objectType
	var action string = "start"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *TaskWorkflowActionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TaskWorkflowActionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TaskWorkflowActionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TaskWorkflowActionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TaskWorkflowActionAllOf) SetAction(v string) {
	o.Action = &v
}

// GetInputParams returns the InputParams field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetInputParams() string {
	if o == nil || o.InputParams == nil {
		var ret string
		return ret
	}
	return *o.InputParams
}

// GetInputParamsOk returns a tuple with the InputParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetInputParamsOk() (*string, bool) {
	if o == nil || o.InputParams == nil {
		return nil, false
	}
	return o.InputParams, true
}

// HasInputParams returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasInputParams() bool {
	if o != nil && o.InputParams != nil {
		return true
	}

	return false
}

// SetInputParams gets a reference to the given string and assigns it to the InputParams field.
func (o *TaskWorkflowActionAllOf) SetInputParams(v string) {
	o.InputParams = &v
}

// GetIsDynamic returns the IsDynamic field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetIsDynamic() bool {
	if o == nil || o.IsDynamic == nil {
		var ret bool
		return ret
	}
	return *o.IsDynamic
}

// GetIsDynamicOk returns a tuple with the IsDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetIsDynamicOk() (*bool, bool) {
	if o == nil || o.IsDynamic == nil {
		return nil, false
	}
	return o.IsDynamic, true
}

// HasIsDynamic returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasIsDynamic() bool {
	if o != nil && o.IsDynamic != nil {
		return true
	}

	return false
}

// SetIsDynamic gets a reference to the given bool and assigns it to the IsDynamic field.
func (o *TaskWorkflowActionAllOf) SetIsDynamic(v bool) {
	o.IsDynamic = &v
}

// GetWaitOnDuplicate returns the WaitOnDuplicate field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetWaitOnDuplicate() bool {
	if o == nil || o.WaitOnDuplicate == nil {
		var ret bool
		return ret
	}
	return *o.WaitOnDuplicate
}

// GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetWaitOnDuplicateOk() (*bool, bool) {
	if o == nil || o.WaitOnDuplicate == nil {
		return nil, false
	}
	return o.WaitOnDuplicate, true
}

// HasWaitOnDuplicate returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasWaitOnDuplicate() bool {
	if o != nil && o.WaitOnDuplicate != nil {
		return true
	}

	return false
}

// SetWaitOnDuplicate gets a reference to the given bool and assigns it to the WaitOnDuplicate field.
func (o *TaskWorkflowActionAllOf) SetWaitOnDuplicate(v bool) {
	o.WaitOnDuplicate = &v
}

// GetWorkflowContext returns the WorkflowContext field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetWorkflowContext() string {
	if o == nil || o.WorkflowContext == nil {
		var ret string
		return ret
	}
	return *o.WorkflowContext
}

// GetWorkflowContextOk returns a tuple with the WorkflowContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetWorkflowContextOk() (*string, bool) {
	if o == nil || o.WorkflowContext == nil {
		return nil, false
	}
	return o.WorkflowContext, true
}

// HasWorkflowContext returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasWorkflowContext() bool {
	if o != nil && o.WorkflowContext != nil {
		return true
	}

	return false
}

// SetWorkflowContext gets a reference to the given string and assigns it to the WorkflowContext field.
func (o *TaskWorkflowActionAllOf) SetWorkflowContext(v string) {
	o.WorkflowContext = &v
}

// GetWorkflowFile returns the WorkflowFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskWorkflowActionAllOf) GetWorkflowFile() TaskFileDownloadInfo {
	if o == nil || o.WorkflowFile.Get() == nil {
		var ret TaskFileDownloadInfo
		return ret
	}
	return *o.WorkflowFile.Get()
}

// GetWorkflowFileOk returns a tuple with the WorkflowFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskWorkflowActionAllOf) GetWorkflowFileOk() (*TaskFileDownloadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowFile.Get(), o.WorkflowFile.IsSet()
}

// HasWorkflowFile returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasWorkflowFile() bool {
	if o != nil && o.WorkflowFile.IsSet() {
		return true
	}

	return false
}

// SetWorkflowFile gets a reference to the given NullableTaskFileDownloadInfo and assigns it to the WorkflowFile field.
func (o *TaskWorkflowActionAllOf) SetWorkflowFile(v TaskFileDownloadInfo) {
	o.WorkflowFile.Set(&v)
}

// SetWorkflowFileNil sets the value for WorkflowFile to be an explicit nil
func (o *TaskWorkflowActionAllOf) SetWorkflowFileNil() {
	o.WorkflowFile.Set(nil)
}

// UnsetWorkflowFile ensures that no value is present for WorkflowFile, not even an explicit nil
func (o *TaskWorkflowActionAllOf) UnsetWorkflowFile() {
	o.WorkflowFile.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TaskWorkflowActionAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWorkflowActionAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TaskWorkflowActionAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TaskWorkflowActionAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TaskWorkflowActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.InputParams != nil {
		toSerialize["InputParams"] = o.InputParams
	}
	if o.IsDynamic != nil {
		toSerialize["IsDynamic"] = o.IsDynamic
	}
	if o.WaitOnDuplicate != nil {
		toSerialize["WaitOnDuplicate"] = o.WaitOnDuplicate
	}
	if o.WorkflowContext != nil {
		toSerialize["WorkflowContext"] = o.WorkflowContext
	}
	if o.WorkflowFile.IsSet() {
		toSerialize["WorkflowFile"] = o.WorkflowFile.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TaskWorkflowActionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTaskWorkflowActionAllOf := _TaskWorkflowActionAllOf{}

	if err = json.Unmarshal(bytes, &varTaskWorkflowActionAllOf); err == nil {
		*o = TaskWorkflowActionAllOf(varTaskWorkflowActionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "InputParams")
		delete(additionalProperties, "IsDynamic")
		delete(additionalProperties, "WaitOnDuplicate")
		delete(additionalProperties, "WorkflowContext")
		delete(additionalProperties, "WorkflowFile")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskWorkflowActionAllOf struct {
	value *TaskWorkflowActionAllOf
	isSet bool
}

func (v NullableTaskWorkflowActionAllOf) Get() *TaskWorkflowActionAllOf {
	return v.value
}

func (v *NullableTaskWorkflowActionAllOf) Set(val *TaskWorkflowActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskWorkflowActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskWorkflowActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskWorkflowActionAllOf(val *TaskWorkflowActionAllOf) *NullableTaskWorkflowActionAllOf {
	return &NullableTaskWorkflowActionAllOf{value: val, isSet: true}
}

func (v NullableTaskWorkflowActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskWorkflowActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
