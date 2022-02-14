/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2022-02-14T08:05:27Z.
 *
 * API version: 0.0.1-38392
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ConvergedinfraComplianceSummary The summary of compliance information based on the compliance analysis (HCL, IMT, etc.) of components.
type ConvergedinfraComplianceSummary struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of elements where compliance information is not available. eg. For HCL of server, OS information is not available.
	Incomplete *int64 `json:"Incomplete,omitempty"`
	// The count of elements where compliance has not been evaluated. e.g. For HCL of server, status has not been evaluated due to lack of information.
	NotEvaluated *int64 `json:"NotEvaluated,omitempty"`
	// The count of elements where compliance has failed for one or more reason. e.g. For HCL of server, some part of the HCL validation has failed.
	NotListed *int64 `json:"NotListed,omitempty"`
	// The count of elements where compliance has passed validation for all components. e.g. For HCL of server, all of the components have passed validation.
	Validated            *int64 `json:"Validated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraComplianceSummary ConvergedinfraComplianceSummary

// NewConvergedinfraComplianceSummary instantiates a new ConvergedinfraComplianceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraComplianceSummary(classId string, objectType string) *ConvergedinfraComplianceSummary {
	this := ConvergedinfraComplianceSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraComplianceSummaryWithDefaults instantiates a new ConvergedinfraComplianceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraComplianceSummaryWithDefaults() *ConvergedinfraComplianceSummary {
	this := ConvergedinfraComplianceSummary{}
	var classId string = "convergedinfra.ComplianceSummary"
	this.ClassId = classId
	var objectType string = "convergedinfra.ComplianceSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraComplianceSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraComplianceSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraComplianceSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraComplianceSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIncomplete returns the Incomplete field value if set, zero value otherwise.
func (o *ConvergedinfraComplianceSummary) GetIncomplete() int64 {
	if o == nil || o.Incomplete == nil {
		var ret int64
		return ret
	}
	return *o.Incomplete
}

// GetIncompleteOk returns a tuple with the Incomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetIncompleteOk() (*int64, bool) {
	if o == nil || o.Incomplete == nil {
		return nil, false
	}
	return o.Incomplete, true
}

// HasIncomplete returns a boolean if a field has been set.
func (o *ConvergedinfraComplianceSummary) HasIncomplete() bool {
	if o != nil && o.Incomplete != nil {
		return true
	}

	return false
}

// SetIncomplete gets a reference to the given int64 and assigns it to the Incomplete field.
func (o *ConvergedinfraComplianceSummary) SetIncomplete(v int64) {
	o.Incomplete = &v
}

// GetNotEvaluated returns the NotEvaluated field value if set, zero value otherwise.
func (o *ConvergedinfraComplianceSummary) GetNotEvaluated() int64 {
	if o == nil || o.NotEvaluated == nil {
		var ret int64
		return ret
	}
	return *o.NotEvaluated
}

// GetNotEvaluatedOk returns a tuple with the NotEvaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetNotEvaluatedOk() (*int64, bool) {
	if o == nil || o.NotEvaluated == nil {
		return nil, false
	}
	return o.NotEvaluated, true
}

// HasNotEvaluated returns a boolean if a field has been set.
func (o *ConvergedinfraComplianceSummary) HasNotEvaluated() bool {
	if o != nil && o.NotEvaluated != nil {
		return true
	}

	return false
}

// SetNotEvaluated gets a reference to the given int64 and assigns it to the NotEvaluated field.
func (o *ConvergedinfraComplianceSummary) SetNotEvaluated(v int64) {
	o.NotEvaluated = &v
}

// GetNotListed returns the NotListed field value if set, zero value otherwise.
func (o *ConvergedinfraComplianceSummary) GetNotListed() int64 {
	if o == nil || o.NotListed == nil {
		var ret int64
		return ret
	}
	return *o.NotListed
}

// GetNotListedOk returns a tuple with the NotListed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetNotListedOk() (*int64, bool) {
	if o == nil || o.NotListed == nil {
		return nil, false
	}
	return o.NotListed, true
}

// HasNotListed returns a boolean if a field has been set.
func (o *ConvergedinfraComplianceSummary) HasNotListed() bool {
	if o != nil && o.NotListed != nil {
		return true
	}

	return false
}

// SetNotListed gets a reference to the given int64 and assigns it to the NotListed field.
func (o *ConvergedinfraComplianceSummary) SetNotListed(v int64) {
	o.NotListed = &v
}

// GetValidated returns the Validated field value if set, zero value otherwise.
func (o *ConvergedinfraComplianceSummary) GetValidated() int64 {
	if o == nil || o.Validated == nil {
		var ret int64
		return ret
	}
	return *o.Validated
}

// GetValidatedOk returns a tuple with the Validated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraComplianceSummary) GetValidatedOk() (*int64, bool) {
	if o == nil || o.Validated == nil {
		return nil, false
	}
	return o.Validated, true
}

// HasValidated returns a boolean if a field has been set.
func (o *ConvergedinfraComplianceSummary) HasValidated() bool {
	if o != nil && o.Validated != nil {
		return true
	}

	return false
}

// SetValidated gets a reference to the given int64 and assigns it to the Validated field.
func (o *ConvergedinfraComplianceSummary) SetValidated(v int64) {
	o.Validated = &v
}

func (o ConvergedinfraComplianceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Incomplete != nil {
		toSerialize["Incomplete"] = o.Incomplete
	}
	if o.NotEvaluated != nil {
		toSerialize["NotEvaluated"] = o.NotEvaluated
	}
	if o.NotListed != nil {
		toSerialize["NotListed"] = o.NotListed
	}
	if o.Validated != nil {
		toSerialize["Validated"] = o.Validated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraComplianceSummary) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraComplianceSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The count of elements where compliance information is not available. eg. For HCL of server, OS information is not available.
		Incomplete *int64 `json:"Incomplete,omitempty"`
		// The count of elements where compliance has not been evaluated. e.g. For HCL of server, status has not been evaluated due to lack of information.
		NotEvaluated *int64 `json:"NotEvaluated,omitempty"`
		// The count of elements where compliance has failed for one or more reason. e.g. For HCL of server, some part of the HCL validation has failed.
		NotListed *int64 `json:"NotListed,omitempty"`
		// The count of elements where compliance has passed validation for all components. e.g. For HCL of server, all of the components have passed validation.
		Validated *int64 `json:"Validated,omitempty"`
	}

	varConvergedinfraComplianceSummaryWithoutEmbeddedStruct := ConvergedinfraComplianceSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraComplianceSummaryWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraComplianceSummary := _ConvergedinfraComplianceSummary{}
		varConvergedinfraComplianceSummary.ClassId = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.ClassId
		varConvergedinfraComplianceSummary.ObjectType = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.ObjectType
		varConvergedinfraComplianceSummary.Incomplete = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.Incomplete
		varConvergedinfraComplianceSummary.NotEvaluated = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.NotEvaluated
		varConvergedinfraComplianceSummary.NotListed = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.NotListed
		varConvergedinfraComplianceSummary.Validated = varConvergedinfraComplianceSummaryWithoutEmbeddedStruct.Validated
		*o = ConvergedinfraComplianceSummary(varConvergedinfraComplianceSummary)
	} else {
		return err
	}

	varConvergedinfraComplianceSummary := _ConvergedinfraComplianceSummary{}

	err = json.Unmarshal(bytes, &varConvergedinfraComplianceSummary)
	if err == nil {
		o.MoBaseComplexType = varConvergedinfraComplianceSummary.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Incomplete")
		delete(additionalProperties, "NotEvaluated")
		delete(additionalProperties, "NotListed")
		delete(additionalProperties, "Validated")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableConvergedinfraComplianceSummary struct {
	value *ConvergedinfraComplianceSummary
	isSet bool
}

func (v NullableConvergedinfraComplianceSummary) Get() *ConvergedinfraComplianceSummary {
	return v.value
}

func (v *NullableConvergedinfraComplianceSummary) Set(val *ConvergedinfraComplianceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraComplianceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraComplianceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraComplianceSummary(val *ConvergedinfraComplianceSummary) *NullableConvergedinfraComplianceSummary {
	return &NullableConvergedinfraComplianceSummary{value: val, isSet: true}
}

func (v NullableConvergedinfraComplianceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraComplianceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
