# VirtualizationHsmTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.HsmTaskResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.HsmTaskResult"]
**Action** | Pointer to **string** | Depicts the action performed in the task. | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Progress** | Pointer to **int64** | Progress information of the task. | [optional] 
**StartTime** | Pointer to **time.Time** | Time of starting the task. | [optional] 
**Status** | Pointer to **string** | Depicts information about the status of the task. | [optional] 
**TargetVersion** | Pointer to **string** | Version availebale on Server. | [optional] 
**TargetVersionName** | Pointer to **string** | Information about of version used in task. | [optional] 
**TaskId** | Pointer to **string** | Depicts information about mapped HSM task. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationHsmTaskResult

`func NewVirtualizationHsmTaskResult(classId string, objectType string, ) *VirtualizationHsmTaskResult`

NewVirtualizationHsmTaskResult instantiates a new VirtualizationHsmTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationHsmTaskResultWithDefaults

`func NewVirtualizationHsmTaskResultWithDefaults() *VirtualizationHsmTaskResult`

NewVirtualizationHsmTaskResultWithDefaults instantiates a new VirtualizationHsmTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationHsmTaskResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationHsmTaskResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationHsmTaskResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationHsmTaskResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationHsmTaskResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationHsmTaskResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *VirtualizationHsmTaskResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VirtualizationHsmTaskResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VirtualizationHsmTaskResult) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VirtualizationHsmTaskResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHosts

`func (o *VirtualizationHsmTaskResult) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualizationHsmTaskResult) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualizationHsmTaskResult) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VirtualizationHsmTaskResult) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *VirtualizationHsmTaskResult) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *VirtualizationHsmTaskResult) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetMessages

`func (o *VirtualizationHsmTaskResult) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VirtualizationHsmTaskResult) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VirtualizationHsmTaskResult) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *VirtualizationHsmTaskResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *VirtualizationHsmTaskResult) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *VirtualizationHsmTaskResult) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetProgress

`func (o *VirtualizationHsmTaskResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VirtualizationHsmTaskResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VirtualizationHsmTaskResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VirtualizationHsmTaskResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartTime

`func (o *VirtualizationHsmTaskResult) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VirtualizationHsmTaskResult) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VirtualizationHsmTaskResult) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VirtualizationHsmTaskResult) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationHsmTaskResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationHsmTaskResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationHsmTaskResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationHsmTaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetVersion

`func (o *VirtualizationHsmTaskResult) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *VirtualizationHsmTaskResult) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *VirtualizationHsmTaskResult) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *VirtualizationHsmTaskResult) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetTargetVersionName

`func (o *VirtualizationHsmTaskResult) GetTargetVersionName() string`

GetTargetVersionName returns the TargetVersionName field if non-nil, zero value otherwise.

### GetTargetVersionNameOk

`func (o *VirtualizationHsmTaskResult) GetTargetVersionNameOk() (*string, bool)`

GetTargetVersionNameOk returns a tuple with the TargetVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersionName

`func (o *VirtualizationHsmTaskResult) SetTargetVersionName(v string)`

SetTargetVersionName sets TargetVersionName field to given value.

### HasTargetVersionName

`func (o *VirtualizationHsmTaskResult) HasTargetVersionName() bool`

HasTargetVersionName returns a boolean if a field has been set.

### GetTaskId

`func (o *VirtualizationHsmTaskResult) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *VirtualizationHsmTaskResult) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *VirtualizationHsmTaskResult) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *VirtualizationHsmTaskResult) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationHsmTaskResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationHsmTaskResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationHsmTaskResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationHsmTaskResult) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


