# TestcryptReadOnlyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "testcrypt.ReadOnlyUser"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "testcrypt.ReadOnlyUser"]
**Rouser** | Pointer to [**[]TestcryptUser**](TestcryptUser.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptReadOnlyUser

`func NewTestcryptReadOnlyUser(classId string, objectType string, ) *TestcryptReadOnlyUser`

NewTestcryptReadOnlyUser instantiates a new TestcryptReadOnlyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptReadOnlyUserWithDefaults

`func NewTestcryptReadOnlyUserWithDefaults() *TestcryptReadOnlyUser`

NewTestcryptReadOnlyUserWithDefaults instantiates a new TestcryptReadOnlyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TestcryptReadOnlyUser) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestcryptReadOnlyUser) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestcryptReadOnlyUser) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TestcryptReadOnlyUser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TestcryptReadOnlyUser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TestcryptReadOnlyUser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRouser

`func (o *TestcryptReadOnlyUser) GetRouser() []TestcryptUser`

GetRouser returns the Rouser field if non-nil, zero value otherwise.

### GetRouserOk

`func (o *TestcryptReadOnlyUser) GetRouserOk() (*[]TestcryptUser, bool)`

GetRouserOk returns a tuple with the Rouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouser

`func (o *TestcryptReadOnlyUser) SetRouser(v []TestcryptUser)`

SetRouser sets Rouser field to given value.

### HasRouser

`func (o *TestcryptReadOnlyUser) HasRouser() bool`

HasRouser returns a boolean if a field has been set.

### SetRouserNil

`func (o *TestcryptReadOnlyUser) SetRouserNil(b bool)`

 SetRouserNil sets the value for Rouser to be an explicit nil

### UnsetRouser
`func (o *TestcryptReadOnlyUser) UnsetRouser()`

UnsetRouser ensures that no value is present for Rouser, not even an explicit nil
### GetAccount

`func (o *TestcryptReadOnlyUser) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptReadOnlyUser) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptReadOnlyUser) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptReadOnlyUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


