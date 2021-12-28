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
	"reflect"
	"strings"
)

// HyperflexConfigResult Profile configuration (deploy, validation) results with the overall state and detailed result messages.
type HyperflexConfigResult struct {
	PolicyAbstractConfigResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The progress percentage of the running configuration or workflow.
	ConfigProgress *string `json:"ConfigProgress,omitempty"`
	// The duration of the running configuration or workflow.
	Duration *string `json:"Duration,omitempty"`
	// The start time of the configuration or workflow.
	StartTime      *string                              `json:"StartTime,omitempty"`
	ClusterProfile *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
	// An array of relationships to hyperflexConfigResultEntry resources.
	ResultEntries        []HyperflexConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexConfigResult HyperflexConfigResult

// NewHyperflexConfigResult instantiates a new HyperflexConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexConfigResult(classId string, objectType string) *HyperflexConfigResult {
	this := HyperflexConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexConfigResultWithDefaults instantiates a new HyperflexConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexConfigResultWithDefaults() *HyperflexConfigResult {
	this := HyperflexConfigResult{}
	var classId string = "hyperflex.ConfigResult"
	this.ClassId = classId
	var objectType string = "hyperflex.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigProgress returns the ConfigProgress field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetConfigProgress() string {
	if o == nil || o.ConfigProgress == nil {
		var ret string
		return ret
	}
	return *o.ConfigProgress
}

// GetConfigProgressOk returns a tuple with the ConfigProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetConfigProgressOk() (*string, bool) {
	if o == nil || o.ConfigProgress == nil {
		return nil, false
	}
	return o.ConfigProgress, true
}

// HasConfigProgress returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasConfigProgress() bool {
	if o != nil && o.ConfigProgress != nil {
		return true
	}

	return false
}

// SetConfigProgress gets a reference to the given string and assigns it to the ConfigProgress field.
func (o *HyperflexConfigResult) SetConfigProgress(v string) {
	o.ConfigProgress = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *HyperflexConfigResult) SetDuration(v string) {
	o.Duration = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *HyperflexConfigResult) SetStartTime(v string) {
	o.StartTime = &v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *HyperflexConfigResult) GetClusterProfile() HyperflexClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret HyperflexClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexConfigResult) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given HyperflexClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *HyperflexConfigResult) SetClusterProfile(v HyperflexClusterProfileRelationship) {
	o.ClusterProfile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexConfigResult) GetResultEntries() []HyperflexConfigResultEntryRelationship {
	if o == nil {
		var ret []HyperflexConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexConfigResult) GetResultEntriesOk() (*[]HyperflexConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return &o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *HyperflexConfigResult) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []HyperflexConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *HyperflexConfigResult) SetResultEntries(v []HyperflexConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o HyperflexConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigProgress != nil {
		toSerialize["ConfigProgress"] = o.ConfigProgress
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The progress percentage of the running configuration or workflow.
		ConfigProgress *string `json:"ConfigProgress,omitempty"`
		// The duration of the running configuration or workflow.
		Duration *string `json:"Duration,omitempty"`
		// The start time of the configuration or workflow.
		StartTime      *string                              `json:"StartTime,omitempty"`
		ClusterProfile *HyperflexClusterProfileRelationship `json:"ClusterProfile,omitempty"`
		// An array of relationships to hyperflexConfigResultEntry resources.
		ResultEntries []HyperflexConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	}

	varHyperflexConfigResultWithoutEmbeddedStruct := HyperflexConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexConfigResult := _HyperflexConfigResult{}
		varHyperflexConfigResult.ClassId = varHyperflexConfigResultWithoutEmbeddedStruct.ClassId
		varHyperflexConfigResult.ObjectType = varHyperflexConfigResultWithoutEmbeddedStruct.ObjectType
		varHyperflexConfigResult.ConfigProgress = varHyperflexConfigResultWithoutEmbeddedStruct.ConfigProgress
		varHyperflexConfigResult.Duration = varHyperflexConfigResultWithoutEmbeddedStruct.Duration
		varHyperflexConfigResult.StartTime = varHyperflexConfigResultWithoutEmbeddedStruct.StartTime
		varHyperflexConfigResult.ClusterProfile = varHyperflexConfigResultWithoutEmbeddedStruct.ClusterProfile
		varHyperflexConfigResult.ResultEntries = varHyperflexConfigResultWithoutEmbeddedStruct.ResultEntries
		*o = HyperflexConfigResult(varHyperflexConfigResult)
	} else {
		return err
	}

	varHyperflexConfigResult := _HyperflexConfigResult{}

	err = json.Unmarshal(bytes, &varHyperflexConfigResult)
	if err == nil {
		o.PolicyAbstractConfigResult = varHyperflexConfigResult.PolicyAbstractConfigResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigProgress")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "ResultEntries")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResult := reflect.ValueOf(o.PolicyAbstractConfigResult)
		for i := 0; i < reflectPolicyAbstractConfigResult.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResult.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexConfigResult struct {
	value *HyperflexConfigResult
	isSet bool
}

func (v NullableHyperflexConfigResult) Get() *HyperflexConfigResult {
	return v.value
}

func (v *NullableHyperflexConfigResult) Set(val *HyperflexConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexConfigResult(val *HyperflexConfigResult) *NullableHyperflexConfigResult {
	return &NullableHyperflexConfigResult{value: val, isSet: true}
}

func (v NullableHyperflexConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
