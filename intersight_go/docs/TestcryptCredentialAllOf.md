# TestcryptCredentialAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "testcrypt.Credential"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "testcrypt.Credential"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Password associated with username, to be stored in the database encrypted. | [optional] 
**Username** | Pointer to **string** | The name of the user whose credentials are being set. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptCredentialAllOf

`func NewTestcryptCredentialAllOf(classId string, objectType string, ) *TestcryptCredentialAllOf`

NewTestcryptCredentialAllOf instantiates a new TestcryptCredentialAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptCredentialAllOfWithDefaults

`func NewTestcryptCredentialAllOfWithDefaults() *TestcryptCredentialAllOf`

NewTestcryptCredentialAllOfWithDefaults instantiates a new TestcryptCredentialAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TestcryptCredentialAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestcryptCredentialAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestcryptCredentialAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TestcryptCredentialAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TestcryptCredentialAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TestcryptCredentialAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *TestcryptCredentialAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *TestcryptCredentialAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *TestcryptCredentialAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *TestcryptCredentialAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *TestcryptCredentialAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TestcryptCredentialAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TestcryptCredentialAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TestcryptCredentialAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *TestcryptCredentialAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TestcryptCredentialAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TestcryptCredentialAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TestcryptCredentialAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccount

`func (o *TestcryptCredentialAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptCredentialAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptCredentialAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptCredentialAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


