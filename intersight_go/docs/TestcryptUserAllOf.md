# TestcryptUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "testcrypt.User"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "testcrypt.User"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | Password associated with username, to be stored in the database encrypted. | [optional] 
**Username** | Pointer to **string** | The name of the user whose credentials are being set. | [optional] 

## Methods

### NewTestcryptUserAllOf

`func NewTestcryptUserAllOf(classId string, objectType string, ) *TestcryptUserAllOf`

NewTestcryptUserAllOf instantiates a new TestcryptUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptUserAllOfWithDefaults

`func NewTestcryptUserAllOfWithDefaults() *TestcryptUserAllOf`

NewTestcryptUserAllOfWithDefaults instantiates a new TestcryptUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TestcryptUserAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TestcryptUserAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TestcryptUserAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TestcryptUserAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TestcryptUserAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TestcryptUserAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *TestcryptUserAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *TestcryptUserAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *TestcryptUserAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *TestcryptUserAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *TestcryptUserAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TestcryptUserAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TestcryptUserAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TestcryptUserAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *TestcryptUserAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TestcryptUserAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TestcryptUserAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TestcryptUserAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


