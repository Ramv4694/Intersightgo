/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-24T09:42:08Z.
 *
 * API version: 0.0.1-37430
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ConvergedinfraHealthCheckExecutionAllOf Definition of the list of properties defined in 'convergedinfra.HealthCheckExecution', excluding properties defined in parent classes.
type ConvergedinfraHealthCheckExecutionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error details of a health check execution failure.
	Error *string `json:"Error,omitempty"`
	// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the pod.
	Result *string `json:"Result,omitempty"`
	// Status of the health check execution. * `Unknown` - Indicates that the health heck execution status is unknown. This mostly happens in case where health check could not be performed due to connectivity issues. * `Succeeded` - Indicates that the health check execution has succeeded. * `Failed` - Indicates that the health check execution has failed. * `Timedout` - Indicates that the health check execution timed out before completion.
	Status *string `json:"Status,omitempty"`
	// A brief summary of health check results.
	Summary               *string                                          `json:"Summary,omitempty"`
	HealthCheckDefinition *ConvergedinfraHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ConvergedinfraHealthCheckExecutionAllOf ConvergedinfraHealthCheckExecutionAllOf

// NewConvergedinfraHealthCheckExecutionAllOf instantiates a new ConvergedinfraHealthCheckExecutionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraHealthCheckExecutionAllOf(classId string, objectType string) *ConvergedinfraHealthCheckExecutionAllOf {
	this := ConvergedinfraHealthCheckExecutionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var result string = "Unknown"
	this.Result = &result
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewConvergedinfraHealthCheckExecutionAllOfWithDefaults instantiates a new ConvergedinfraHealthCheckExecutionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraHealthCheckExecutionAllOfWithDefaults() *ConvergedinfraHealthCheckExecutionAllOf {
	this := ConvergedinfraHealthCheckExecutionAllOf{}
	var classId string = "convergedinfra.HealthCheckExecution"
	this.ClassId = classId
	var objectType string = "convergedinfra.HealthCheckExecution"
	this.ObjectType = objectType
	var result string = "Unknown"
	this.Result = &result
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetError(v string) {
	o.Error = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetResult(v string) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetSummary(v string) {
	o.Summary = &v
}

// GetHealthCheckDefinition returns the HealthCheckDefinition field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetHealthCheckDefinition() ConvergedinfraHealthCheckDefinitionRelationship {
	if o == nil || o.HealthCheckDefinition == nil {
		var ret ConvergedinfraHealthCheckDefinitionRelationship
		return ret
	}
	return *o.HealthCheckDefinition
}

// GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) GetHealthCheckDefinitionOk() (*ConvergedinfraHealthCheckDefinitionRelationship, bool) {
	if o == nil || o.HealthCheckDefinition == nil {
		return nil, false
	}
	return o.HealthCheckDefinition, true
}

// HasHealthCheckDefinition returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecutionAllOf) HasHealthCheckDefinition() bool {
	if o != nil && o.HealthCheckDefinition != nil {
		return true
	}

	return false
}

// SetHealthCheckDefinition gets a reference to the given ConvergedinfraHealthCheckDefinitionRelationship and assigns it to the HealthCheckDefinition field.
func (o *ConvergedinfraHealthCheckExecutionAllOf) SetHealthCheckDefinition(v ConvergedinfraHealthCheckDefinitionRelationship) {
	o.HealthCheckDefinition = &v
}

func (o ConvergedinfraHealthCheckExecutionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Error != nil {
		toSerialize["Error"] = o.Error
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}
	if o.HealthCheckDefinition != nil {
		toSerialize["HealthCheckDefinition"] = o.HealthCheckDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraHealthCheckExecutionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraHealthCheckExecutionAllOf := _ConvergedinfraHealthCheckExecutionAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraHealthCheckExecutionAllOf); err == nil {
		*o = ConvergedinfraHealthCheckExecutionAllOf(varConvergedinfraHealthCheckExecutionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "HealthCheckDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraHealthCheckExecutionAllOf struct {
	value *ConvergedinfraHealthCheckExecutionAllOf
	isSet bool
}

func (v NullableConvergedinfraHealthCheckExecutionAllOf) Get() *ConvergedinfraHealthCheckExecutionAllOf {
	return v.value
}

func (v *NullableConvergedinfraHealthCheckExecutionAllOf) Set(val *ConvergedinfraHealthCheckExecutionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraHealthCheckExecutionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraHealthCheckExecutionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraHealthCheckExecutionAllOf(val *ConvergedinfraHealthCheckExecutionAllOf) *NullableConvergedinfraHealthCheckExecutionAllOf {
	return &NullableConvergedinfraHealthCheckExecutionAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraHealthCheckExecutionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraHealthCheckExecutionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
