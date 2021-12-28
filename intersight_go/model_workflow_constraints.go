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

// WorkflowConstraints Captures the constraints for valid parameter values.
type WorkflowConstraints struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string              `json:"ObjectType"`
	EnumList   []WorkflowEnumEntry `json:"EnumList,omitempty"`
	// Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.
	Max *float64 `json:"Max,omitempty"`
	// Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.
	Min *float64 `json:"Min,omitempty"`
	// When the parameter is a string this regular expression is used to ensure the value is valid.
	Regex                *string `json:"Regex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowConstraints WorkflowConstraints

// NewWorkflowConstraints instantiates a new WorkflowConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowConstraints(classId string, objectType string) *WorkflowConstraints {
	this := WorkflowConstraints{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowConstraintsWithDefaults instantiates a new WorkflowConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowConstraintsWithDefaults() *WorkflowConstraints {
	this := WorkflowConstraints{}
	var classId string = "workflow.Constraints"
	this.ClassId = classId
	var objectType string = "workflow.Constraints"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowConstraints) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowConstraints) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowConstraints) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowConstraints) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowConstraints) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowConstraints) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnumList returns the EnumList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowConstraints) GetEnumList() []WorkflowEnumEntry {
	if o == nil {
		var ret []WorkflowEnumEntry
		return ret
	}
	return o.EnumList
}

// GetEnumListOk returns a tuple with the EnumList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowConstraints) GetEnumListOk() (*[]WorkflowEnumEntry, bool) {
	if o == nil || o.EnumList == nil {
		return nil, false
	}
	return &o.EnumList, true
}

// HasEnumList returns a boolean if a field has been set.
func (o *WorkflowConstraints) HasEnumList() bool {
	if o != nil && o.EnumList != nil {
		return true
	}

	return false
}

// SetEnumList gets a reference to the given []WorkflowEnumEntry and assigns it to the EnumList field.
func (o *WorkflowConstraints) SetEnumList(v []WorkflowEnumEntry) {
	o.EnumList = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowConstraints) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraints) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowConstraints) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *WorkflowConstraints) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowConstraints) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraints) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowConstraints) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *WorkflowConstraints) SetMin(v float64) {
	o.Min = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *WorkflowConstraints) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraints) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *WorkflowConstraints) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *WorkflowConstraints) SetRegex(v string) {
	o.Regex = &v
}

func (o WorkflowConstraints) MarshalJSON() ([]byte, error) {
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
	if o.EnumList != nil {
		toSerialize["EnumList"] = o.EnumList
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}
	if o.Regex != nil {
		toSerialize["Regex"] = o.Regex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowConstraints) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowConstraintsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string              `json:"ObjectType"`
		EnumList   []WorkflowEnumEntry `json:"EnumList,omitempty"`
		// Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.
		Max *float64 `json:"Max,omitempty"`
		// Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.
		Min *float64 `json:"Min,omitempty"`
		// When the parameter is a string this regular expression is used to ensure the value is valid.
		Regex *string `json:"Regex,omitempty"`
	}

	varWorkflowConstraintsWithoutEmbeddedStruct := WorkflowConstraintsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowConstraintsWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowConstraints := _WorkflowConstraints{}
		varWorkflowConstraints.ClassId = varWorkflowConstraintsWithoutEmbeddedStruct.ClassId
		varWorkflowConstraints.ObjectType = varWorkflowConstraintsWithoutEmbeddedStruct.ObjectType
		varWorkflowConstraints.EnumList = varWorkflowConstraintsWithoutEmbeddedStruct.EnumList
		varWorkflowConstraints.Max = varWorkflowConstraintsWithoutEmbeddedStruct.Max
		varWorkflowConstraints.Min = varWorkflowConstraintsWithoutEmbeddedStruct.Min
		varWorkflowConstraints.Regex = varWorkflowConstraintsWithoutEmbeddedStruct.Regex
		*o = WorkflowConstraints(varWorkflowConstraints)
	} else {
		return err
	}

	varWorkflowConstraints := _WorkflowConstraints{}

	err = json.Unmarshal(bytes, &varWorkflowConstraints)
	if err == nil {
		o.MoBaseComplexType = varWorkflowConstraints.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnumList")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")
		delete(additionalProperties, "Regex")

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

type NullableWorkflowConstraints struct {
	value *WorkflowConstraints
	isSet bool
}

func (v NullableWorkflowConstraints) Get() *WorkflowConstraints {
	return v.value
}

func (v *NullableWorkflowConstraints) Set(val *WorkflowConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowConstraints(val *WorkflowConstraints) *NullableWorkflowConstraints {
	return &NullableWorkflowConstraints{value: val, isSet: true}
}

func (v NullableWorkflowConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
