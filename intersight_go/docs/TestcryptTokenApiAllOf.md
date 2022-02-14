# TestcryptTokenApiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "testcrypt.TokenApi"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "testcrypt.TokenApi"]
**SessionId** | Pointer to **string** | The session identifier to be used to create/read/update or delete Encryption token. This could be a user&#39;s (web) session id or api key id, etc. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptTokenApiAllOf

`func NewTestcryptTokenApiAllOf(classId string, objectType string, ) *TestcryptTokenApiAllOf`

NewTestcryptTokenApiAllOf instantiates a new TestcryptTokenApiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptTokenApiAllOfWithDefaults

`func NewTestcryptTokenApiAllOfWithDefaults() *TestcryptTokenApiAllOf`

NewTestcryptTokenApiAllOfWithDefaults instantiates a new TestcryptTokenApiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TestcryptTokenApiAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestcryptTokenApiAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestcryptTokenApiAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TestcryptTokenApiAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TestcryptTokenApiAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TestcryptTokenApiAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSessionId

`func (o *TestcryptTokenApiAllOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *TestcryptTokenApiAllOf) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *TestcryptTokenApiAllOf) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *TestcryptTokenApiAllOf) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetAccount

`func (o *TestcryptTokenApiAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptTokenApiAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptTokenApiAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptTokenApiAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


