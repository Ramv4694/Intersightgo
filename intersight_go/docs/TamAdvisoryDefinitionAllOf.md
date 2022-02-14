# TamAdvisoryDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.AdvisoryDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.AdvisoryDefinition"]
**Actions** | Pointer to [**[]TamAction**](TamAction.md) |  | [optional] 
**AdvisoryDetails** | Pointer to [**NullableTamBaseAdvisoryDetails**](tam.BaseAdvisoryDetails.md) |  | [optional] 
**AdvisoryId** | Pointer to **string** | Cisco generated identifier for the published security/field-notice/end-of-life advisory. | [optional] 
**ApiDataSources** | Pointer to [**[]TamApiDataSource**](TamApiDataSource.md) |  | [optional] 
**DatePublished** | Pointer to **time.Time** | Date when the security/field-notice/end-of-life advisory was first published by Cisco. | [optional] 
**DateUpdated** | Pointer to **time.Time** | Date when the security/field-notice/end-of-life advisory was last updated by Cisco. | [optional] 
**ExternalUrl** | Pointer to **string** | A link to an external URL describing security Advisory in more details. | [optional] 
**Recommendation** | Pointer to **string** | Recommended action to resolve the security advisory. | [optional] 
**S3DataSources** | Pointer to [**[]TamS3DataSource**](TamS3DataSource.md) |  | [optional] 
**Type** | Pointer to **string** | The type (field notice, security advisory, end-of-life milestone advisory etc.) of Intersight advisory. * &#x60;securityAdvisory&#x60; - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * &#x60;fieldNotice&#x60; - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). * &#x60;eolAdvisory&#x60; - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html). | [optional] [default to "securityAdvisory"]
**Version** | Pointer to **string** | Cisco assigned advisory/field-notice/end-of-life version after latest revision. | [optional] 
**Workaround** | Pointer to **string** | Workarounds available for the advisory. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryDefinitionAllOf

`func NewTamAdvisoryDefinitionAllOf(classId string, objectType string, ) *TamAdvisoryDefinitionAllOf`

NewTamAdvisoryDefinitionAllOf instantiates a new TamAdvisoryDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryDefinitionAllOfWithDefaults

`func NewTamAdvisoryDefinitionAllOfWithDefaults() *TamAdvisoryDefinitionAllOf`

NewTamAdvisoryDefinitionAllOfWithDefaults instantiates a new TamAdvisoryDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamAdvisoryDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamAdvisoryDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamAdvisoryDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamAdvisoryDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamAdvisoryDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamAdvisoryDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActions

`func (o *TamAdvisoryDefinitionAllOf) GetActions() []TamAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TamAdvisoryDefinitionAllOf) GetActionsOk() (*[]TamAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TamAdvisoryDefinitionAllOf) SetActions(v []TamAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TamAdvisoryDefinitionAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *TamAdvisoryDefinitionAllOf) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *TamAdvisoryDefinitionAllOf) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetAdvisoryDetails

`func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryDetails() TamBaseAdvisoryDetails`

GetAdvisoryDetails returns the AdvisoryDetails field if non-nil, zero value otherwise.

### GetAdvisoryDetailsOk

`func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryDetailsOk() (*TamBaseAdvisoryDetails, bool)`

GetAdvisoryDetailsOk returns a tuple with the AdvisoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryDetails

`func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryDetails(v TamBaseAdvisoryDetails)`

SetAdvisoryDetails sets AdvisoryDetails field to given value.

### HasAdvisoryDetails

`func (o *TamAdvisoryDefinitionAllOf) HasAdvisoryDetails() bool`

HasAdvisoryDetails returns a boolean if a field has been set.

### SetAdvisoryDetailsNil

`func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryDetailsNil(b bool)`

 SetAdvisoryDetailsNil sets the value for AdvisoryDetails to be an explicit nil

### UnsetAdvisoryDetails
`func (o *TamAdvisoryDefinitionAllOf) UnsetAdvisoryDetails()`

UnsetAdvisoryDetails ensures that no value is present for AdvisoryDetails, not even an explicit nil
### GetAdvisoryId

`func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *TamAdvisoryDefinitionAllOf) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetApiDataSources

`func (o *TamAdvisoryDefinitionAllOf) GetApiDataSources() []TamApiDataSource`

GetApiDataSources returns the ApiDataSources field if non-nil, zero value otherwise.

### GetApiDataSourcesOk

`func (o *TamAdvisoryDefinitionAllOf) GetApiDataSourcesOk() (*[]TamApiDataSource, bool)`

GetApiDataSourcesOk returns a tuple with the ApiDataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDataSources

`func (o *TamAdvisoryDefinitionAllOf) SetApiDataSources(v []TamApiDataSource)`

SetApiDataSources sets ApiDataSources field to given value.

### HasApiDataSources

`func (o *TamAdvisoryDefinitionAllOf) HasApiDataSources() bool`

HasApiDataSources returns a boolean if a field has been set.

### SetApiDataSourcesNil

`func (o *TamAdvisoryDefinitionAllOf) SetApiDataSourcesNil(b bool)`

 SetApiDataSourcesNil sets the value for ApiDataSources to be an explicit nil

### UnsetApiDataSources
`func (o *TamAdvisoryDefinitionAllOf) UnsetApiDataSources()`

UnsetApiDataSources ensures that no value is present for ApiDataSources, not even an explicit nil
### GetDatePublished

`func (o *TamAdvisoryDefinitionAllOf) GetDatePublished() time.Time`

GetDatePublished returns the DatePublished field if non-nil, zero value otherwise.

### GetDatePublishedOk

`func (o *TamAdvisoryDefinitionAllOf) GetDatePublishedOk() (*time.Time, bool)`

GetDatePublishedOk returns a tuple with the DatePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublished

`func (o *TamAdvisoryDefinitionAllOf) SetDatePublished(v time.Time)`

SetDatePublished sets DatePublished field to given value.

### HasDatePublished

`func (o *TamAdvisoryDefinitionAllOf) HasDatePublished() bool`

HasDatePublished returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TamAdvisoryDefinitionAllOf) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TamAdvisoryDefinitionAllOf) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TamAdvisoryDefinitionAllOf) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TamAdvisoryDefinitionAllOf) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetExternalUrl

`func (o *TamAdvisoryDefinitionAllOf) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *TamAdvisoryDefinitionAllOf) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *TamAdvisoryDefinitionAllOf) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *TamAdvisoryDefinitionAllOf) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetRecommendation

`func (o *TamAdvisoryDefinitionAllOf) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *TamAdvisoryDefinitionAllOf) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *TamAdvisoryDefinitionAllOf) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *TamAdvisoryDefinitionAllOf) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetS3DataSources

`func (o *TamAdvisoryDefinitionAllOf) GetS3DataSources() []TamS3DataSource`

GetS3DataSources returns the S3DataSources field if non-nil, zero value otherwise.

### GetS3DataSourcesOk

`func (o *TamAdvisoryDefinitionAllOf) GetS3DataSourcesOk() (*[]TamS3DataSource, bool)`

GetS3DataSourcesOk returns a tuple with the S3DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3DataSources

`func (o *TamAdvisoryDefinitionAllOf) SetS3DataSources(v []TamS3DataSource)`

SetS3DataSources sets S3DataSources field to given value.

### HasS3DataSources

`func (o *TamAdvisoryDefinitionAllOf) HasS3DataSources() bool`

HasS3DataSources returns a boolean if a field has been set.

### SetS3DataSourcesNil

`func (o *TamAdvisoryDefinitionAllOf) SetS3DataSourcesNil(b bool)`

 SetS3DataSourcesNil sets the value for S3DataSources to be an explicit nil

### UnsetS3DataSources
`func (o *TamAdvisoryDefinitionAllOf) UnsetS3DataSources()`

UnsetS3DataSources ensures that no value is present for S3DataSources, not even an explicit nil
### GetType

`func (o *TamAdvisoryDefinitionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamAdvisoryDefinitionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamAdvisoryDefinitionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamAdvisoryDefinitionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *TamAdvisoryDefinitionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TamAdvisoryDefinitionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TamAdvisoryDefinitionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TamAdvisoryDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkaround

`func (o *TamAdvisoryDefinitionAllOf) GetWorkaround() string`

GetWorkaround returns the Workaround field if non-nil, zero value otherwise.

### GetWorkaroundOk

`func (o *TamAdvisoryDefinitionAllOf) GetWorkaroundOk() (*string, bool)`

GetWorkaroundOk returns a tuple with the Workaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaround

`func (o *TamAdvisoryDefinitionAllOf) SetWorkaround(v string)`

SetWorkaround sets Workaround field to given value.

### HasWorkaround

`func (o *TamAdvisoryDefinitionAllOf) HasWorkaround() bool`

HasWorkaround returns a boolean if a field has been set.

### GetOrganization

`func (o *TamAdvisoryDefinitionAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TamAdvisoryDefinitionAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TamAdvisoryDefinitionAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TamAdvisoryDefinitionAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

