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
	"time"
)

// VirtualizationHsmTaskResultAllOf Definition of the list of properties defined in 'virtualization.HsmTaskResult', excluding properties defined in parent classes.
type VirtualizationHsmTaskResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Depicts the action performed in the task.
	Action *string `json:"Action,omitempty"`
	// Information about of the task.
	Description *string  `json:"Description,omitempty"`
	Hosts       []string `json:"Hosts,omitempty"`
	Messages    []string `json:"Messages,omitempty"`
	// Status code information about the task.
	OperationStatusCode *int64 `json:"OperationStatusCode,omitempty"`
	// Progress information of the task.
	Progress *int64 `json:"Progress,omitempty"`
	// Time of starting the task.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Depicts information about the status of the task.
	Status *string `json:"Status,omitempty"`
	// Depicts information about mapped HSM task.
	TaskId               *string                              `json:"TaskId,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationHsmTaskResultAllOf VirtualizationHsmTaskResultAllOf

// NewVirtualizationHsmTaskResultAllOf instantiates a new VirtualizationHsmTaskResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationHsmTaskResultAllOf(classId string, objectType string) *VirtualizationHsmTaskResultAllOf {
	this := VirtualizationHsmTaskResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationHsmTaskResultAllOfWithDefaults instantiates a new VirtualizationHsmTaskResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationHsmTaskResultAllOfWithDefaults() *VirtualizationHsmTaskResultAllOf {
	this := VirtualizationHsmTaskResultAllOf{}
	var classId string = "virtualization.HsmTaskResult"
	this.ClassId = classId
	var objectType string = "virtualization.HsmTaskResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationHsmTaskResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationHsmTaskResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationHsmTaskResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationHsmTaskResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *VirtualizationHsmTaskResultAllOf) SetAction(v string) {
	o.Action = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationHsmTaskResultAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationHsmTaskResultAllOf) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationHsmTaskResultAllOf) GetHostsOk() (*[]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *VirtualizationHsmTaskResultAllOf) SetHosts(v []string) {
	o.Hosts = v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationHsmTaskResultAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationHsmTaskResultAllOf) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *VirtualizationHsmTaskResultAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetOperationStatusCode returns the OperationStatusCode field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetOperationStatusCode() int64 {
	if o == nil || o.OperationStatusCode == nil {
		var ret int64
		return ret
	}
	return *o.OperationStatusCode
}

// GetOperationStatusCodeOk returns a tuple with the OperationStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetOperationStatusCodeOk() (*int64, bool) {
	if o == nil || o.OperationStatusCode == nil {
		return nil, false
	}
	return o.OperationStatusCode, true
}

// HasOperationStatusCode returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasOperationStatusCode() bool {
	if o != nil && o.OperationStatusCode != nil {
		return true
	}

	return false
}

// SetOperationStatusCode gets a reference to the given int64 and assigns it to the OperationStatusCode field.
func (o *VirtualizationHsmTaskResultAllOf) SetOperationStatusCode(v int64) {
	o.OperationStatusCode = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetProgress() int64 {
	if o == nil || o.Progress == nil {
		var ret int64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetProgressOk() (*int64, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int64 and assigns it to the Progress field.
func (o *VirtualizationHsmTaskResultAllOf) SetProgress(v int64) {
	o.Progress = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *VirtualizationHsmTaskResultAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationHsmTaskResultAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *VirtualizationHsmTaskResultAllOf) SetTaskId(v string) {
	o.TaskId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationHsmTaskResultAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHsmTaskResultAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationHsmTaskResultAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationHsmTaskResultAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o VirtualizationHsmTaskResultAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.OperationStatusCode != nil {
		toSerialize["OperationStatusCode"] = o.OperationStatusCode
	}
	if o.Progress != nil {
		toSerialize["Progress"] = o.Progress
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskId != nil {
		toSerialize["TaskId"] = o.TaskId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationHsmTaskResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationHsmTaskResultAllOf := _VirtualizationHsmTaskResultAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationHsmTaskResultAllOf); err == nil {
		*o = VirtualizationHsmTaskResultAllOf(varVirtualizationHsmTaskResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Hosts")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "OperationStatusCode")
		delete(additionalProperties, "Progress")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskId")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationHsmTaskResultAllOf struct {
	value *VirtualizationHsmTaskResultAllOf
	isSet bool
}

func (v NullableVirtualizationHsmTaskResultAllOf) Get() *VirtualizationHsmTaskResultAllOf {
	return v.value
}

func (v *NullableVirtualizationHsmTaskResultAllOf) Set(val *VirtualizationHsmTaskResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationHsmTaskResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationHsmTaskResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationHsmTaskResultAllOf(val *VirtualizationHsmTaskResultAllOf) *NullableVirtualizationHsmTaskResultAllOf {
	return &NullableVirtualizationHsmTaskResultAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationHsmTaskResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationHsmTaskResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}