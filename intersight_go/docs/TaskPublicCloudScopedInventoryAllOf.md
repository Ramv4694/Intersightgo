# TaskPublicCloudScopedInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "task.PublicCloudScopedInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "task.PublicCloudScopedInventory"]
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewTaskPublicCloudScopedInventoryAllOf

`func NewTaskPublicCloudScopedInventoryAllOf(classId string, objectType string, ) *TaskPublicCloudScopedInventoryAllOf`

NewTaskPublicCloudScopedInventoryAllOf instantiates a new TaskPublicCloudScopedInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPublicCloudScopedInventoryAllOfWithDefaults

`func NewTaskPublicCloudScopedInventoryAllOfWithDefaults() *TaskPublicCloudScopedInventoryAllOf`

NewTaskPublicCloudScopedInventoryAllOfWithDefaults instantiates a new TaskPublicCloudScopedInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TaskPublicCloudScopedInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TaskPublicCloudScopedInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TaskPublicCloudScopedInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TaskPublicCloudScopedInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TaskPublicCloudScopedInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TaskPublicCloudScopedInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRegisteredDevice

`func (o *TaskPublicCloudScopedInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *TaskPublicCloudScopedInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *TaskPublicCloudScopedInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *TaskPublicCloudScopedInventoryAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


