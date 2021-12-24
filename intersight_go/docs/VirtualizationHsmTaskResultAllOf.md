# VirtualizationHsmTaskResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.HsmTaskResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.HsmTaskResult"]
**Action** | Pointer to **string** | Depicts the action performed in the task. | [optional] 
**Description** | Pointer to **string** | Information about of the task. | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**OperationStatusCode** | Pointer to **int64** | Status code information about the task. | [optional] 
**Progress** | Pointer to **int64** | Progress information of the task. | [optional] 
**StartTime** | Pointer to **time.Time** | Time of starting the task. | [optional] 
**Status** | Pointer to **string** | Depicts information about the status of the task. | [optional] 
**TaskId** | Pointer to **string** | Depicts information about mapped HSM task. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationHsmTaskResultAllOf

`func NewVirtualizationHsmTaskResultAllOf(classId string, objectType string, ) *VirtualizationHsmTaskResultAllOf`

NewVirtualizationHsmTaskResultAllOf instantiates a new VirtualizationHsmTaskResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationHsmTaskResultAllOfWithDefaults

`func NewVirtualizationHsmTaskResultAllOfWithDefaults() *VirtualizationHsmTaskResultAllOf`

NewVirtualizationHsmTaskResultAllOfWithDefaults instantiates a new VirtualizationHsmTaskResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationHsmTaskResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationHsmTaskResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationHsmTaskResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationHsmTaskResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationHsmTaskResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationHsmTaskResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *VirtualizationHsmTaskResultAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VirtualizationHsmTaskResultAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VirtualizationHsmTaskResultAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VirtualizationHsmTaskResultAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationHsmTaskResultAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationHsmTaskResultAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationHsmTaskResultAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationHsmTaskResultAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationHsmTaskResultAllOf) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationHsmTaskResultAllOf) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationHsmTaskResultAllOf) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationHsmTaskResultAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationHsmTaskResultAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationHsmTaskResultAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetMessages

`func (o *VirtualizationHsmTaskResultAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VirtualizationHsmTaskResultAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VirtualizationHsmTaskResultAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *VirtualizationHsmTaskResultAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *VirtualizationHsmTaskResultAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *VirtualizationHsmTaskResultAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetOperationStatusCode

`func (o *VirtualizationHsmTaskResultAllOf) GetOperationStatusCode() int64`

GetOperationStatusCode returns the OperationStatusCode field if non-nil, zero value otherwise.

### GetOperationStatusCodeOk

`func (o *VirtualizationHsmTaskResultAllOf) GetOperationStatusCodeOk() (*int64, bool)`

GetOperationStatusCodeOk returns a tuple with the OperationStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatusCode

`func (o *VirtualizationHsmTaskResultAllOf) SetOperationStatusCode(v int64)`

SetOperationStatusCode sets OperationStatusCode field to given value.

### HasOperationStatusCode

`func (o *VirtualizationHsmTaskResultAllOf) HasOperationStatusCode() bool`

HasOperationStatusCode returns a boolean if a field has been set.

### GetProgress

`func (o *VirtualizationHsmTaskResultAllOf) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VirtualizationHsmTaskResultAllOf) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VirtualizationHsmTaskResultAllOf) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VirtualizationHsmTaskResultAllOf) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartTime

`func (o *VirtualizationHsmTaskResultAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VirtualizationHsmTaskResultAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VirtualizationHsmTaskResultAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VirtualizationHsmTaskResultAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationHsmTaskResultAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationHsmTaskResultAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationHsmTaskResultAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationHsmTaskResultAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *VirtualizationHsmTaskResultAllOf) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *VirtualizationHsmTaskResultAllOf) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *VirtualizationHsmTaskResultAllOf) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *VirtualizationHsmTaskResultAllOf) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationHsmTaskResultAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationHsmTaskResultAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationHsmTaskResultAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationHsmTaskResultAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

