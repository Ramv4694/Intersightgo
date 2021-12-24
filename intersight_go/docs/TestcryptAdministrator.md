# TestcryptAdministrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "testcrypt.Administrator"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "testcrypt.Administrator"]
**Admin** | Pointer to [**NullableTestcryptUser**](testcrypt.User.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptAdministrator

`func NewTestcryptAdministrator(classId string, objectType string, ) *TestcryptAdministrator`

NewTestcryptAdministrator instantiates a new TestcryptAdministrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptAdministratorWithDefaults

`func NewTestcryptAdministratorWithDefaults() *TestcryptAdministrator`

NewTestcryptAdministratorWithDefaults instantiates a new TestcryptAdministrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TestcryptAdministrator) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestcryptAdministrator) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestcryptAdministrator) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TestcryptAdministrator) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TestcryptAdministrator) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TestcryptAdministrator) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdmin

`func (o *TestcryptAdministrator) GetAdmin() TestcryptUser`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *TestcryptAdministrator) GetAdminOk() (*TestcryptUser, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *TestcryptAdministrator) SetAdmin(v TestcryptUser)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *TestcryptAdministrator) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### SetAdminNil

`func (o *TestcryptAdministrator) SetAdminNil(b bool)`

 SetAdminNil sets the value for Admin to be an explicit nil

### UnsetAdmin
`func (o *TestcryptAdministrator) UnsetAdmin()`

UnsetAdmin ensures that no value is present for Admin, not even an explicit nil
### GetAccount

`func (o *TestcryptAdministrator) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptAdministrator) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptAdministrator) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptAdministrator) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


