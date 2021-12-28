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

// ComputeServerOpStatusAllOf Definition of the list of properties defined in 'compute.ServerOpStatus', excluding properties defined in parent classes.
type ComputeServerOpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - The state denotes that the settings are applied successfully in the target server. Applying - The state denotes that the settings are being applied in the target server. Failed - The state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// The workflow type being started. The workflow name to distinguish workflow by type.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerOpStatusAllOf ComputeServerOpStatusAllOf

// NewComputeServerOpStatusAllOf instantiates a new ComputeServerOpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerOpStatusAllOf(classId string, objectType string) *ComputeServerOpStatusAllOf {
	this := ComputeServerOpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// NewComputeServerOpStatusAllOfWithDefaults instantiates a new ComputeServerOpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerOpStatusAllOfWithDefaults() *ComputeServerOpStatusAllOf {
	this := ComputeServerOpStatusAllOf{}
	var classId string = "compute.ServerOpStatus"
	this.ClassId = classId
	var objectType string = "compute.ServerOpStatus"
	this.ObjectType = objectType
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerOpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerOpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerOpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerOpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *ComputeServerOpStatusAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatusAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *ComputeServerOpStatusAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *ComputeServerOpStatusAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *ComputeServerOpStatusAllOf) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatusAllOf) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *ComputeServerOpStatusAllOf) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *ComputeServerOpStatusAllOf) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o ComputeServerOpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerOpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeServerOpStatusAllOf := _ComputeServerOpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varComputeServerOpStatusAllOf); err == nil {
		*o = ComputeServerOpStatusAllOf(varComputeServerOpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeServerOpStatusAllOf struct {
	value *ComputeServerOpStatusAllOf
	isSet bool
}

func (v NullableComputeServerOpStatusAllOf) Get() *ComputeServerOpStatusAllOf {
	return v.value
}

func (v *NullableComputeServerOpStatusAllOf) Set(val *ComputeServerOpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerOpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerOpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerOpStatusAllOf(val *ComputeServerOpStatusAllOf) *NullableComputeServerOpStatusAllOf {
	return &NullableComputeServerOpStatusAllOf{value: val, isSet: true}
}

func (v NullableComputeServerOpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerOpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
