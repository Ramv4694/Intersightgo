/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-27T12:26:28Z.
 *
 * API version: 0.0.1-37434
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// TamSecurityAdvisoryAllOf Definition of the list of properties defined in 'tam.SecurityAdvisory', excluding properties defined in parent classes.
type TamSecurityAdvisoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string      `json:"ObjectType"`
	Actions    []TamAction `json:"Actions,omitempty"`
	// Cisco generated identifier for the published security advisory.
	AdvisoryId     *string            `json:"AdvisoryId,omitempty"`
	ApiDataSources []TamApiDataSource `json:"ApiDataSources,omitempty"`
	// CVSS version 3 base score for the security Advisory.
	BaseScore *float32 `json:"BaseScore,omitempty"`
	CveIds    []string `json:"CveIds,omitempty"`
	// Date when the security advisory was first published by Cisco.
	DatePublished *time.Time `json:"DatePublished,omitempty"`
	// Date when the security advisory was last updated by Cisco.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// CVSS version 3 environmental score for the security Advisory.
	EnvironmentalScore *float32 `json:"EnvironmentalScore,omitempty"`
	// A link to an external URL describing security Advisory in more details.
	ExternalUrl *string `json:"ExternalUrl,omitempty"`
	// Recommended action to resolve the security advisory.
	Recommendation *string `json:"Recommendation,omitempty"`
	// Cisco assigned status of the published advisory based on whether the investigation is complete or on-going. * `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.
	Status *string `json:"Status,omitempty"`
	// CVSS version 3 temporal score for the security Advisory.
	TemporalScore *float32 `json:"TemporalScore,omitempty"`
	// Cisco assigned advisory version after latest revision.
	Version *string `json:"Version,omitempty"`
	// Workarounds available for the advisory.
	Workaround           *string                               `json:"Workaround,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamSecurityAdvisoryAllOf TamSecurityAdvisoryAllOf

// NewTamSecurityAdvisoryAllOf instantiates a new TamSecurityAdvisoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamSecurityAdvisoryAllOf(classId string, objectType string) *TamSecurityAdvisoryAllOf {
	this := TamSecurityAdvisoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// NewTamSecurityAdvisoryAllOfWithDefaults instantiates a new TamSecurityAdvisoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamSecurityAdvisoryAllOfWithDefaults() *TamSecurityAdvisoryAllOf {
	this := TamSecurityAdvisoryAllOf{}
	var classId string = "tam.SecurityAdvisory"
	this.ClassId = classId
	var objectType string = "tam.SecurityAdvisory"
	this.ObjectType = objectType
	var status string = "interim"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamSecurityAdvisoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamSecurityAdvisoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamSecurityAdvisoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamSecurityAdvisoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamSecurityAdvisoryAllOf) GetActions() []TamAction {
	if o == nil {
		var ret []TamAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamSecurityAdvisoryAllOf) GetActionsOk() (*[]TamAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []TamAction and assigns it to the Actions field.
func (o *TamSecurityAdvisoryAllOf) SetActions(v []TamAction) {
	o.Actions = v
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetAdvisoryId() string {
	if o == nil || o.AdvisoryId == nil {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || o.AdvisoryId == nil {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasAdvisoryId() bool {
	if o != nil && o.AdvisoryId != nil {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *TamSecurityAdvisoryAllOf) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetApiDataSources returns the ApiDataSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamSecurityAdvisoryAllOf) GetApiDataSources() []TamApiDataSource {
	if o == nil {
		var ret []TamApiDataSource
		return ret
	}
	return o.ApiDataSources
}

// GetApiDataSourcesOk returns a tuple with the ApiDataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamSecurityAdvisoryAllOf) GetApiDataSourcesOk() (*[]TamApiDataSource, bool) {
	if o == nil || o.ApiDataSources == nil {
		return nil, false
	}
	return &o.ApiDataSources, true
}

// HasApiDataSources returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasApiDataSources() bool {
	if o != nil && o.ApiDataSources != nil {
		return true
	}

	return false
}

// SetApiDataSources gets a reference to the given []TamApiDataSource and assigns it to the ApiDataSources field.
func (o *TamSecurityAdvisoryAllOf) SetApiDataSources(v []TamApiDataSource) {
	o.ApiDataSources = v
}

// GetBaseScore returns the BaseScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetBaseScore() float32 {
	if o == nil || o.BaseScore == nil {
		var ret float32
		return ret
	}
	return *o.BaseScore
}

// GetBaseScoreOk returns a tuple with the BaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetBaseScoreOk() (*float32, bool) {
	if o == nil || o.BaseScore == nil {
		return nil, false
	}
	return o.BaseScore, true
}

// HasBaseScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasBaseScore() bool {
	if o != nil && o.BaseScore != nil {
		return true
	}

	return false
}

// SetBaseScore gets a reference to the given float32 and assigns it to the BaseScore field.
func (o *TamSecurityAdvisoryAllOf) SetBaseScore(v float32) {
	o.BaseScore = &v
}

// GetCveIds returns the CveIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamSecurityAdvisoryAllOf) GetCveIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CveIds
}

// GetCveIdsOk returns a tuple with the CveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamSecurityAdvisoryAllOf) GetCveIdsOk() (*[]string, bool) {
	if o == nil || o.CveIds == nil {
		return nil, false
	}
	return &o.CveIds, true
}

// HasCveIds returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasCveIds() bool {
	if o != nil && o.CveIds != nil {
		return true
	}

	return false
}

// SetCveIds gets a reference to the given []string and assigns it to the CveIds field.
func (o *TamSecurityAdvisoryAllOf) SetCveIds(v []string) {
	o.CveIds = v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetDatePublished() time.Time {
	if o == nil || o.DatePublished == nil {
		var ret time.Time
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetDatePublishedOk() (*time.Time, bool) {
	if o == nil || o.DatePublished == nil {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasDatePublished() bool {
	if o != nil && o.DatePublished != nil {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given time.Time and assigns it to the DatePublished field.
func (o *TamSecurityAdvisoryAllOf) SetDatePublished(v time.Time) {
	o.DatePublished = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *TamSecurityAdvisoryAllOf) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

// GetEnvironmentalScore returns the EnvironmentalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetEnvironmentalScore() float32 {
	if o == nil || o.EnvironmentalScore == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentalScore
}

// GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetEnvironmentalScoreOk() (*float32, bool) {
	if o == nil || o.EnvironmentalScore == nil {
		return nil, false
	}
	return o.EnvironmentalScore, true
}

// HasEnvironmentalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasEnvironmentalScore() bool {
	if o != nil && o.EnvironmentalScore != nil {
		return true
	}

	return false
}

// SetEnvironmentalScore gets a reference to the given float32 and assigns it to the EnvironmentalScore field.
func (o *TamSecurityAdvisoryAllOf) SetEnvironmentalScore(v float32) {
	o.EnvironmentalScore = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetExternalUrl() string {
	if o == nil || o.ExternalUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetExternalUrlOk() (*string, bool) {
	if o == nil || o.ExternalUrl == nil {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasExternalUrl() bool {
	if o != nil && o.ExternalUrl != nil {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *TamSecurityAdvisoryAllOf) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetRecommendation() string {
	if o == nil || o.Recommendation == nil {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetRecommendationOk() (*string, bool) {
	if o == nil || o.Recommendation == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasRecommendation() bool {
	if o != nil && o.Recommendation != nil {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *TamSecurityAdvisoryAllOf) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TamSecurityAdvisoryAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTemporalScore returns the TemporalScore field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetTemporalScore() float32 {
	if o == nil || o.TemporalScore == nil {
		var ret float32
		return ret
	}
	return *o.TemporalScore
}

// GetTemporalScoreOk returns a tuple with the TemporalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetTemporalScoreOk() (*float32, bool) {
	if o == nil || o.TemporalScore == nil {
		return nil, false
	}
	return o.TemporalScore, true
}

// HasTemporalScore returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasTemporalScore() bool {
	if o != nil && o.TemporalScore != nil {
		return true
	}

	return false
}

// SetTemporalScore gets a reference to the given float32 and assigns it to the TemporalScore field.
func (o *TamSecurityAdvisoryAllOf) SetTemporalScore(v float32) {
	o.TemporalScore = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TamSecurityAdvisoryAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetWorkaround returns the Workaround field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetWorkaround() string {
	if o == nil || o.Workaround == nil {
		var ret string
		return ret
	}
	return *o.Workaround
}

// GetWorkaroundOk returns a tuple with the Workaround field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetWorkaroundOk() (*string, bool) {
	if o == nil || o.Workaround == nil {
		return nil, false
	}
	return o.Workaround, true
}

// HasWorkaround returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasWorkaround() bool {
	if o != nil && o.Workaround != nil {
		return true
	}

	return false
}

// SetWorkaround gets a reference to the given string and assigns it to the Workaround field.
func (o *TamSecurityAdvisoryAllOf) SetWorkaround(v string) {
	o.Workaround = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *TamSecurityAdvisoryAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamSecurityAdvisoryAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *TamSecurityAdvisoryAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *TamSecurityAdvisoryAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o TamSecurityAdvisoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.AdvisoryId != nil {
		toSerialize["AdvisoryId"] = o.AdvisoryId
	}
	if o.ApiDataSources != nil {
		toSerialize["ApiDataSources"] = o.ApiDataSources
	}
	if o.BaseScore != nil {
		toSerialize["BaseScore"] = o.BaseScore
	}
	if o.CveIds != nil {
		toSerialize["CveIds"] = o.CveIds
	}
	if o.DatePublished != nil {
		toSerialize["DatePublished"] = o.DatePublished
	}
	if o.DateUpdated != nil {
		toSerialize["DateUpdated"] = o.DateUpdated
	}
	if o.EnvironmentalScore != nil {
		toSerialize["EnvironmentalScore"] = o.EnvironmentalScore
	}
	if o.ExternalUrl != nil {
		toSerialize["ExternalUrl"] = o.ExternalUrl
	}
	if o.Recommendation != nil {
		toSerialize["Recommendation"] = o.Recommendation
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TemporalScore != nil {
		toSerialize["TemporalScore"] = o.TemporalScore
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Workaround != nil {
		toSerialize["Workaround"] = o.Workaround
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamSecurityAdvisoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamSecurityAdvisoryAllOf := _TamSecurityAdvisoryAllOf{}

	if err = json.Unmarshal(bytes, &varTamSecurityAdvisoryAllOf); err == nil {
		*o = TamSecurityAdvisoryAllOf(varTamSecurityAdvisoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "AdvisoryId")
		delete(additionalProperties, "ApiDataSources")
		delete(additionalProperties, "BaseScore")
		delete(additionalProperties, "CveIds")
		delete(additionalProperties, "DatePublished")
		delete(additionalProperties, "DateUpdated")
		delete(additionalProperties, "EnvironmentalScore")
		delete(additionalProperties, "ExternalUrl")
		delete(additionalProperties, "Recommendation")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TemporalScore")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Workaround")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamSecurityAdvisoryAllOf struct {
	value *TamSecurityAdvisoryAllOf
	isSet bool
}

func (v NullableTamSecurityAdvisoryAllOf) Get() *TamSecurityAdvisoryAllOf {
	return v.value
}

func (v *NullableTamSecurityAdvisoryAllOf) Set(val *TamSecurityAdvisoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamSecurityAdvisoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamSecurityAdvisoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamSecurityAdvisoryAllOf(val *TamSecurityAdvisoryAllOf) *NullableTamSecurityAdvisoryAllOf {
	return &NullableTamSecurityAdvisoryAllOf{value: val, isSet: true}
}

func (v NullableTamSecurityAdvisoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamSecurityAdvisoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
