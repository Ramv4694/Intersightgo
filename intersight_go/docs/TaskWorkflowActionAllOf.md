# TaskWorkflowActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "task.WorkflowAction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "task.WorkflowAction"]
**Action** | Pointer to **string** | Action for test workflow. * &#x60;start&#x60; - Start action for the workflow. * &#x60;stop&#x60; - Stop action for the workflow. | [optional] [default to "start"]
**InputParams** | Pointer to **string** | Json formatted string input parameters to workflow. | [optional] 
**IsDynamic** | Pointer to **bool** | When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow. | [optional] 
**WaitOnDuplicate** | Pointer to **bool** | When true, the workflow will wait for previous one to complete before starting a new one. | [optional] 
**WorkflowContext** | Pointer to **string** | Json formatted string that has the contents of the workflow context used when starting a workflow. | [optional] 
**WorkflowFile** | Pointer to [**NullableTaskFileDownloadInfo**](task.FileDownloadInfo.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTaskWorkflowActionAllOf

`func NewTaskWorkflowActionAllOf(classId string, objectType string, ) *TaskWorkflowActionAllOf`

NewTaskWorkflowActionAllOf instantiates a new TaskWorkflowActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWorkflowActionAllOfWithDefaults

`func NewTaskWorkflowActionAllOfWithDefaults() *TaskWorkflowActionAllOf`

NewTaskWorkflowActionAllOfWithDefaults instantiates a new TaskWorkflowActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TaskWorkflowActionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TaskWorkflowActionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TaskWorkflowActionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TaskWorkflowActionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TaskWorkflowActionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TaskWorkflowActionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *TaskWorkflowActionAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TaskWorkflowActionAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TaskWorkflowActionAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TaskWorkflowActionAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInputParams

`func (o *TaskWorkflowActionAllOf) GetInputParams() string`

GetInputParams returns the InputParams field if non-nil, zero value otherwise.

### GetInputParamsOk

`func (o *TaskWorkflowActionAllOf) GetInputParamsOk() (*string, bool)`

GetInputParamsOk returns a tuple with the InputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParams

`func (o *TaskWorkflowActionAllOf) SetInputParams(v string)`

SetInputParams sets InputParams field to given value.

### HasInputParams

`func (o *TaskWorkflowActionAllOf) HasInputParams() bool`

HasInputParams returns a boolean if a field has been set.

### GetIsDynamic

`func (o *TaskWorkflowActionAllOf) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *TaskWorkflowActionAllOf) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *TaskWorkflowActionAllOf) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.

### HasIsDynamic

`func (o *TaskWorkflowActionAllOf) HasIsDynamic() bool`

HasIsDynamic returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *TaskWorkflowActionAllOf) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *TaskWorkflowActionAllOf) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *TaskWorkflowActionAllOf) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *TaskWorkflowActionAllOf) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetWorkflowContext

`func (o *TaskWorkflowActionAllOf) GetWorkflowContext() string`

GetWorkflowContext returns the WorkflowContext field if non-nil, zero value otherwise.

### GetWorkflowContextOk

`func (o *TaskWorkflowActionAllOf) GetWorkflowContextOk() (*string, bool)`

GetWorkflowContextOk returns a tuple with the WorkflowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowContext

`func (o *TaskWorkflowActionAllOf) SetWorkflowContext(v string)`

SetWorkflowContext sets WorkflowContext field to given value.

### HasWorkflowContext

`func (o *TaskWorkflowActionAllOf) HasWorkflowContext() bool`

HasWorkflowContext returns a boolean if a field has been set.

### GetWorkflowFile

`func (o *TaskWorkflowActionAllOf) GetWorkflowFile() TaskFileDownloadInfo`

GetWorkflowFile returns the WorkflowFile field if non-nil, zero value otherwise.

### GetWorkflowFileOk

`func (o *TaskWorkflowActionAllOf) GetWorkflowFileOk() (*TaskFileDownloadInfo, bool)`

GetWorkflowFileOk returns a tuple with the WorkflowFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFile

`func (o *TaskWorkflowActionAllOf) SetWorkflowFile(v TaskFileDownloadInfo)`

SetWorkflowFile sets WorkflowFile field to given value.

### HasWorkflowFile

`func (o *TaskWorkflowActionAllOf) HasWorkflowFile() bool`

HasWorkflowFile returns a boolean if a field has been set.

### SetWorkflowFileNil

`func (o *TaskWorkflowActionAllOf) SetWorkflowFileNil(b bool)`

 SetWorkflowFileNil sets the value for WorkflowFile to be an explicit nil

### UnsetWorkflowFile
`func (o *TaskWorkflowActionAllOf) UnsetWorkflowFile()`

UnsetWorkflowFile ensures that no value is present for WorkflowFile, not even an explicit nil
### GetAccount

`func (o *TaskWorkflowActionAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TaskWorkflowActionAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TaskWorkflowActionAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TaskWorkflowActionAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


