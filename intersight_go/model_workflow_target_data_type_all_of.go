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
)

// WorkflowTargetDataTypeAllOf Definition of the list of properties defined in 'workflow.TargetDataType', excluding properties defined in parent classes.
type WorkflowTargetDataTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType               string                             `json:"ObjectType"`
	CustomDataTypeProperties NullableWorkflowCustomDataProperty `json:"CustomDataTypeProperties,omitempty"`
	// When this property is true then an array of targets can be passed as input.
	IsArray *bool `json:"IsArray,omitempty"`
	// Specify the maximum value of the array.
	Max *int64 `json:"Max,omitempty"`
	// Specify the minimum value of the array.
	Min                  *int64                   `json:"Min,omitempty"`
	Properties           []WorkflowTargetProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetDataTypeAllOf WorkflowTargetDataTypeAllOf

// NewWorkflowTargetDataTypeAllOf instantiates a new WorkflowTargetDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetDataTypeAllOf(classId string, objectType string) *WorkflowTargetDataTypeAllOf {
	this := WorkflowTargetDataTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isArray bool = false
	this.IsArray = &isArray
	return &this
}

// NewWorkflowTargetDataTypeAllOfWithDefaults instantiates a new WorkflowTargetDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetDataTypeAllOfWithDefaults() *WorkflowTargetDataTypeAllOf {
	this := WorkflowTargetDataTypeAllOf{}
	var classId string = "workflow.TargetDataType"
	this.ClassId = classId
	var objectType string = "workflow.TargetDataType"
	this.ObjectType = objectType
	var isArray bool = false
	this.IsArray = &isArray
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTargetDataTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTargetDataTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTargetDataTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTargetDataTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCustomDataTypeProperties returns the CustomDataTypeProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetDataTypeAllOf) GetCustomDataTypeProperties() WorkflowCustomDataProperty {
	if o == nil || o.CustomDataTypeProperties.Get() == nil {
		var ret WorkflowCustomDataProperty
		return ret
	}
	return *o.CustomDataTypeProperties.Get()
}

// GetCustomDataTypePropertiesOk returns a tuple with the CustomDataTypeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetDataTypeAllOf) GetCustomDataTypePropertiesOk() (*WorkflowCustomDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomDataTypeProperties.Get(), o.CustomDataTypeProperties.IsSet()
}

// HasCustomDataTypeProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasCustomDataTypeProperties() bool {
	if o != nil && o.CustomDataTypeProperties.IsSet() {
		return true
	}

	return false
}

// SetCustomDataTypeProperties gets a reference to the given NullableWorkflowCustomDataProperty and assigns it to the CustomDataTypeProperties field.
func (o *WorkflowTargetDataTypeAllOf) SetCustomDataTypeProperties(v WorkflowCustomDataProperty) {
	o.CustomDataTypeProperties.Set(&v)
}

// SetCustomDataTypePropertiesNil sets the value for CustomDataTypeProperties to be an explicit nil
func (o *WorkflowTargetDataTypeAllOf) SetCustomDataTypePropertiesNil() {
	o.CustomDataTypeProperties.Set(nil)
}

// UnsetCustomDataTypeProperties ensures that no value is present for CustomDataTypeProperties, not even an explicit nil
func (o *WorkflowTargetDataTypeAllOf) UnsetCustomDataTypeProperties() {
	o.CustomDataTypeProperties.Unset()
}

// GetIsArray returns the IsArray field value if set, zero value otherwise.
func (o *WorkflowTargetDataTypeAllOf) GetIsArray() bool {
	if o == nil || o.IsArray == nil {
		var ret bool
		return ret
	}
	return *o.IsArray
}

// GetIsArrayOk returns a tuple with the IsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetIsArrayOk() (*bool, bool) {
	if o == nil || o.IsArray == nil {
		return nil, false
	}
	return o.IsArray, true
}

// HasIsArray returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasIsArray() bool {
	if o != nil && o.IsArray != nil {
		return true
	}

	return false
}

// SetIsArray gets a reference to the given bool and assigns it to the IsArray field.
func (o *WorkflowTargetDataTypeAllOf) SetIsArray(v bool) {
	o.IsArray = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowTargetDataTypeAllOf) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *WorkflowTargetDataTypeAllOf) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowTargetDataTypeAllOf) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetDataTypeAllOf) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *WorkflowTargetDataTypeAllOf) SetMin(v int64) {
	o.Min = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetDataTypeAllOf) GetProperties() []WorkflowTargetProperty {
	if o == nil {
		var ret []WorkflowTargetProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetDataTypeAllOf) GetPropertiesOk() (*[]WorkflowTargetProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowTargetDataTypeAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowTargetProperty and assigns it to the Properties field.
func (o *WorkflowTargetDataTypeAllOf) SetProperties(v []WorkflowTargetProperty) {
	o.Properties = v
}

func (o WorkflowTargetDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CustomDataTypeProperties.IsSet() {
		toSerialize["CustomDataTypeProperties"] = o.CustomDataTypeProperties.Get()
	}
	if o.IsArray != nil {
		toSerialize["IsArray"] = o.IsArray
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetDataTypeAllOf := _WorkflowTargetDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetDataTypeAllOf); err == nil {
		*o = WorkflowTargetDataTypeAllOf(varWorkflowTargetDataTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CustomDataTypeProperties")
		delete(additionalProperties, "IsArray")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")
		delete(additionalProperties, "Properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTargetDataTypeAllOf struct {
	value *WorkflowTargetDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowTargetDataTypeAllOf) Get() *WorkflowTargetDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowTargetDataTypeAllOf) Set(val *WorkflowTargetDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetDataTypeAllOf(val *WorkflowTargetDataTypeAllOf) *NullableWorkflowTargetDataTypeAllOf {
	return &NullableWorkflowTargetDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTargetDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
