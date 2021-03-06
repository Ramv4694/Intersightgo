# AssetTerraformIntegrationTerraformCloudOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.TerraformIntegrationTerraformCloudOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.TerraformIntegrationTerraformCloudOptions"]
**DefaultManagedHosts** | Pointer to **[]string** |  | [optional] 
**DefaultTerraformOrganization** | Pointer to **string** | Default organization for Terraform Cloud platform type. | [optional] 

## Methods

### NewAssetTerraformIntegrationTerraformCloudOptionsAllOf

`func NewAssetTerraformIntegrationTerraformCloudOptionsAllOf(classId string, objectType string, ) *AssetTerraformIntegrationTerraformCloudOptionsAllOf`

NewAssetTerraformIntegrationTerraformCloudOptionsAllOf instantiates a new AssetTerraformIntegrationTerraformCloudOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTerraformIntegrationTerraformCloudOptionsAllOfWithDefaults

`func NewAssetTerraformIntegrationTerraformCloudOptionsAllOfWithDefaults() *AssetTerraformIntegrationTerraformCloudOptionsAllOf`

NewAssetTerraformIntegrationTerraformCloudOptionsAllOfWithDefaults instantiates a new AssetTerraformIntegrationTerraformCloudOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultManagedHosts

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetDefaultManagedHosts() []string`

GetDefaultManagedHosts returns the DefaultManagedHosts field if non-nil, zero value otherwise.

### GetDefaultManagedHostsOk

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetDefaultManagedHostsOk() (*[]string, bool)`

GetDefaultManagedHostsOk returns a tuple with the DefaultManagedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultManagedHosts

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) SetDefaultManagedHosts(v []string)`

SetDefaultManagedHosts sets DefaultManagedHosts field to given value.

### HasDefaultManagedHosts

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) HasDefaultManagedHosts() bool`

HasDefaultManagedHosts returns a boolean if a field has been set.

### SetDefaultManagedHostsNil

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) SetDefaultManagedHostsNil(b bool)`

 SetDefaultManagedHostsNil sets the value for DefaultManagedHosts to be an explicit nil

### UnsetDefaultManagedHosts
`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) UnsetDefaultManagedHosts()`

UnsetDefaultManagedHosts ensures that no value is present for DefaultManagedHosts, not even an explicit nil
### GetDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetDefaultTerraformOrganization() string`

GetDefaultTerraformOrganization returns the DefaultTerraformOrganization field if non-nil, zero value otherwise.

### GetDefaultTerraformOrganizationOk

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) GetDefaultTerraformOrganizationOk() (*string, bool)`

GetDefaultTerraformOrganizationOk returns a tuple with the DefaultTerraformOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) SetDefaultTerraformOrganization(v string)`

SetDefaultTerraformOrganization sets DefaultTerraformOrganization field to given value.

### HasDefaultTerraformOrganization

`func (o *AssetTerraformIntegrationTerraformCloudOptionsAllOf) HasDefaultTerraformOrganization() bool`

HasDefaultTerraformOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


