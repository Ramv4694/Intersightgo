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
)

// WorkflowWorkflowPropertiesAllOf Definition of the list of properties defined in 'workflow.WorkflowProperties', excluding properties defined in parent classes.
type WorkflowWorkflowPropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to false workflow is not cloneable. It is set to true only if Workflow is not internal and it does not have any internal tasks.
	Cloneable *bool `json:"Cloneable,omitempty"`
	// Enabling this flag will capture request and response details as debug logs for tasks that are using workflow.BatchApi for implementation. For other tasks in the workflow which are not based on workflow.BatchApi logs will not be generated.
	EnableDebug *bool `json:"EnableDebug,omitempty"`
	// When set to false the workflow is owned by the system and used for internal services. Such workflows cannot be directly used by external entities.
	ExternalMeta *bool `json:"ExternalMeta,omitempty"`
	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty"`
	// Supported status of the definition. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
	SupportStatus        *string `json:"SupportStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowPropertiesAllOf WorkflowWorkflowPropertiesAllOf

// NewWorkflowWorkflowPropertiesAllOf instantiates a new WorkflowWorkflowPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowPropertiesAllOf(classId string, objectType string) *WorkflowWorkflowPropertiesAllOf {
	this := WorkflowWorkflowPropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cloneable bool = true
	this.Cloneable = &cloneable
	var enableDebug bool = false
	this.EnableDebug = &enableDebug
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryable bool = false
	this.Retryable = &retryable
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	return &this
}

// NewWorkflowWorkflowPropertiesAllOfWithDefaults instantiates a new WorkflowWorkflowPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowPropertiesAllOfWithDefaults() *WorkflowWorkflowPropertiesAllOf {
	this := WorkflowWorkflowPropertiesAllOf{}
	var classId string = "workflow.WorkflowProperties"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowProperties"
	this.ObjectType = objectType
	var cloneable bool = true
	this.Cloneable = &cloneable
	var enableDebug bool = false
	this.EnableDebug = &enableDebug
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryable bool = false
	this.Retryable = &retryable
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowPropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowPropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowPropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowPropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloneable returns the Cloneable field value if set, zero value otherwise.
func (o *WorkflowWorkflowPropertiesAllOf) GetCloneable() bool {
	if o == nil || o.Cloneable == nil {
		var ret bool
		return ret
	}
	return *o.Cloneable
}

// GetCloneableOk returns a tuple with the Cloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetCloneableOk() (*bool, bool) {
	if o == nil || o.Cloneable == nil {
		return nil, false
	}
	return o.Cloneable, true
}

// HasCloneable returns a boolean if a field has been set.
func (o *WorkflowWorkflowPropertiesAllOf) HasCloneable() bool {
	if o != nil && o.Cloneable != nil {
		return true
	}

	return false
}

// SetCloneable gets a reference to the given bool and assigns it to the Cloneable field.
func (o *WorkflowWorkflowPropertiesAllOf) SetCloneable(v bool) {
	o.Cloneable = &v
}

// GetEnableDebug returns the EnableDebug field value if set, zero value otherwise.
func (o *WorkflowWorkflowPropertiesAllOf) GetEnableDebug() bool {
	if o == nil || o.EnableDebug == nil {
		var ret bool
		return ret
	}
	return *o.EnableDebug
}

// GetEnableDebugOk returns a tuple with the EnableDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetEnableDebugOk() (*bool, bool) {
	if o == nil || o.EnableDebug == nil {
		return nil, false
	}
	return o.EnableDebug, true
}

// HasEnableDebug returns a boolean if a field has been set.
func (o *WorkflowWorkflowPropertiesAllOf) HasEnableDebug() bool {
	if o != nil && o.EnableDebug != nil {
		return true
	}

	return false
}

// SetEnableDebug gets a reference to the given bool and assigns it to the EnableDebug field.
func (o *WorkflowWorkflowPropertiesAllOf) SetEnableDebug(v bool) {
	o.EnableDebug = &v
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowWorkflowPropertiesAllOf) GetExternalMeta() bool {
	if o == nil || o.ExternalMeta == nil {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetExternalMetaOk() (*bool, bool) {
	if o == nil || o.ExternalMeta == nil {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowWorkflowPropertiesAllOf) HasExternalMeta() bool {
	if o != nil && o.ExternalMeta != nil {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowWorkflowPropertiesAllOf) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowPropertiesAllOf) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowPropertiesAllOf) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowPropertiesAllOf) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *WorkflowWorkflowPropertiesAllOf) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowPropertiesAllOf) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *WorkflowWorkflowPropertiesAllOf) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *WorkflowWorkflowPropertiesAllOf) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

func (o WorkflowWorkflowPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cloneable != nil {
		toSerialize["Cloneable"] = o.Cloneable
	}
	if o.EnableDebug != nil {
		toSerialize["EnableDebug"] = o.EnableDebug
	}
	if o.ExternalMeta != nil {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	if o.SupportStatus != nil {
		toSerialize["SupportStatus"] = o.SupportStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowPropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWorkflowPropertiesAllOf := _WorkflowWorkflowPropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWorkflowPropertiesAllOf); err == nil {
		*o = WorkflowWorkflowPropertiesAllOf(varWorkflowWorkflowPropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cloneable")
		delete(additionalProperties, "EnableDebug")
		delete(additionalProperties, "ExternalMeta")
		delete(additionalProperties, "Retryable")
		delete(additionalProperties, "SupportStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWorkflowPropertiesAllOf struct {
	value *WorkflowWorkflowPropertiesAllOf
	isSet bool
}

func (v NullableWorkflowWorkflowPropertiesAllOf) Get() *WorkflowWorkflowPropertiesAllOf {
	return v.value
}

func (v *NullableWorkflowWorkflowPropertiesAllOf) Set(val *WorkflowWorkflowPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowPropertiesAllOf(val *WorkflowWorkflowPropertiesAllOf) *NullableWorkflowWorkflowPropertiesAllOf {
	return &NullableWorkflowWorkflowPropertiesAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
