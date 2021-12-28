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

// PolicyConfigChangeContext Initial configuration state snapshot and configuration state.
type PolicyConfigChangeContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates reason for failure state of configChangeState.
	ConfigChangeError *string `json:"ConfigChangeError,omitempty"`
	// Indicates a profile's configuration change state. Used for tracking pending-changes and out-of-synch states. * `Ok` - Config change state represents Validation for change/drift is successful or is not applicable. * `Validating-change` - Config change state represents policy changes are being validated. This state starts when policy is changed and becomes different from deployed changes (Pending-changes). * `Validating-drift` - Config change state represents endpoint configuration changes are being validated. This state starts when endpoint is changed and endpoint configuration becomes different from policy configured changes (Out-of-sync). * `Change-failed` - Config change state represents there is internal error in calculating policy change. * `Drift-failed` - Config change state represents there is internal error in calculating endpoint configuraion drift.
	ConfigChangeState    *string                     `json:"ConfigChangeState,omitempty"`
	InitialConfigContext NullablePolicyConfigContext `json:"InitialConfigContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigChangeContext PolicyConfigChangeContext

// NewPolicyConfigChangeContext instantiates a new PolicyConfigChangeContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigChangeContext(classId string, objectType string) *PolicyConfigChangeContext {
	this := PolicyConfigChangeContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configChangeState string = "Ok"
	this.ConfigChangeState = &configChangeState
	return &this
}

// NewPolicyConfigChangeContextWithDefaults instantiates a new PolicyConfigChangeContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigChangeContextWithDefaults() *PolicyConfigChangeContext {
	this := PolicyConfigChangeContext{}
	var classId string = "policy.ConfigChangeContext"
	this.ClassId = classId
	var objectType string = "policy.ConfigChangeContext"
	this.ObjectType = objectType
	var configChangeState string = "Ok"
	this.ConfigChangeState = &configChangeState
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigChangeContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigChangeContext) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigChangeContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigChangeContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChangeError returns the ConfigChangeError field value if set, zero value otherwise.
func (o *PolicyConfigChangeContext) GetConfigChangeError() string {
	if o == nil || o.ConfigChangeError == nil {
		var ret string
		return ret
	}
	return *o.ConfigChangeError
}

// GetConfigChangeErrorOk returns a tuple with the ConfigChangeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeContext) GetConfigChangeErrorOk() (*string, bool) {
	if o == nil || o.ConfigChangeError == nil {
		return nil, false
	}
	return o.ConfigChangeError, true
}

// HasConfigChangeError returns a boolean if a field has been set.
func (o *PolicyConfigChangeContext) HasConfigChangeError() bool {
	if o != nil && o.ConfigChangeError != nil {
		return true
	}

	return false
}

// SetConfigChangeError gets a reference to the given string and assigns it to the ConfigChangeError field.
func (o *PolicyConfigChangeContext) SetConfigChangeError(v string) {
	o.ConfigChangeError = &v
}

// GetConfigChangeState returns the ConfigChangeState field value if set, zero value otherwise.
func (o *PolicyConfigChangeContext) GetConfigChangeState() string {
	if o == nil || o.ConfigChangeState == nil {
		var ret string
		return ret
	}
	return *o.ConfigChangeState
}

// GetConfigChangeStateOk returns a tuple with the ConfigChangeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeContext) GetConfigChangeStateOk() (*string, bool) {
	if o == nil || o.ConfigChangeState == nil {
		return nil, false
	}
	return o.ConfigChangeState, true
}

// HasConfigChangeState returns a boolean if a field has been set.
func (o *PolicyConfigChangeContext) HasConfigChangeState() bool {
	if o != nil && o.ConfigChangeState != nil {
		return true
	}

	return false
}

// SetConfigChangeState gets a reference to the given string and assigns it to the ConfigChangeState field.
func (o *PolicyConfigChangeContext) SetConfigChangeState(v string) {
	o.ConfigChangeState = &v
}

// GetInitialConfigContext returns the InitialConfigContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChangeContext) GetInitialConfigContext() PolicyConfigContext {
	if o == nil || o.InitialConfigContext.Get() == nil {
		var ret PolicyConfigContext
		return ret
	}
	return *o.InitialConfigContext.Get()
}

// GetInitialConfigContextOk returns a tuple with the InitialConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChangeContext) GetInitialConfigContextOk() (*PolicyConfigContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitialConfigContext.Get(), o.InitialConfigContext.IsSet()
}

// HasInitialConfigContext returns a boolean if a field has been set.
func (o *PolicyConfigChangeContext) HasInitialConfigContext() bool {
	if o != nil && o.InitialConfigContext.IsSet() {
		return true
	}

	return false
}

// SetInitialConfigContext gets a reference to the given NullablePolicyConfigContext and assigns it to the InitialConfigContext field.
func (o *PolicyConfigChangeContext) SetInitialConfigContext(v PolicyConfigContext) {
	o.InitialConfigContext.Set(&v)
}

// SetInitialConfigContextNil sets the value for InitialConfigContext to be an explicit nil
func (o *PolicyConfigChangeContext) SetInitialConfigContextNil() {
	o.InitialConfigContext.Set(nil)
}

// UnsetInitialConfigContext ensures that no value is present for InitialConfigContext, not even an explicit nil
func (o *PolicyConfigChangeContext) UnsetInitialConfigContext() {
	o.InitialConfigContext.Unset()
}

func (o PolicyConfigChangeContext) MarshalJSON() ([]byte, error) {
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
	if o.ConfigChangeError != nil {
		toSerialize["ConfigChangeError"] = o.ConfigChangeError
	}
	if o.ConfigChangeState != nil {
		toSerialize["ConfigChangeState"] = o.ConfigChangeState
	}
	if o.InitialConfigContext.IsSet() {
		toSerialize["InitialConfigContext"] = o.InitialConfigContext.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyConfigChangeContext) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyConfigChangeContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates reason for failure state of configChangeState.
		ConfigChangeError *string `json:"ConfigChangeError,omitempty"`
		// Indicates a profile's configuration change state. Used for tracking pending-changes and out-of-synch states. * `Ok` - Config change state represents Validation for change/drift is successful or is not applicable. * `Validating-change` - Config change state represents policy changes are being validated. This state starts when policy is changed and becomes different from deployed changes (Pending-changes). * `Validating-drift` - Config change state represents endpoint configuration changes are being validated. This state starts when endpoint is changed and endpoint configuration becomes different from policy configured changes (Out-of-sync). * `Change-failed` - Config change state represents there is internal error in calculating policy change. * `Drift-failed` - Config change state represents there is internal error in calculating endpoint configuraion drift.
		ConfigChangeState    *string                     `json:"ConfigChangeState,omitempty"`
		InitialConfigContext NullablePolicyConfigContext `json:"InitialConfigContext,omitempty"`
	}

	varPolicyConfigChangeContextWithoutEmbeddedStruct := PolicyConfigChangeContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyConfigChangeContextWithoutEmbeddedStruct)
	if err == nil {
		varPolicyConfigChangeContext := _PolicyConfigChangeContext{}
		varPolicyConfigChangeContext.ClassId = varPolicyConfigChangeContextWithoutEmbeddedStruct.ClassId
		varPolicyConfigChangeContext.ObjectType = varPolicyConfigChangeContextWithoutEmbeddedStruct.ObjectType
		varPolicyConfigChangeContext.ConfigChangeError = varPolicyConfigChangeContextWithoutEmbeddedStruct.ConfigChangeError
		varPolicyConfigChangeContext.ConfigChangeState = varPolicyConfigChangeContextWithoutEmbeddedStruct.ConfigChangeState
		varPolicyConfigChangeContext.InitialConfigContext = varPolicyConfigChangeContextWithoutEmbeddedStruct.InitialConfigContext
		*o = PolicyConfigChangeContext(varPolicyConfigChangeContext)
	} else {
		return err
	}

	varPolicyConfigChangeContext := _PolicyConfigChangeContext{}

	err = json.Unmarshal(bytes, &varPolicyConfigChangeContext)
	if err == nil {
		o.MoBaseComplexType = varPolicyConfigChangeContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChangeError")
		delete(additionalProperties, "ConfigChangeState")
		delete(additionalProperties, "InitialConfigContext")

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

type NullablePolicyConfigChangeContext struct {
	value *PolicyConfigChangeContext
	isSet bool
}

func (v NullablePolicyConfigChangeContext) Get() *PolicyConfigChangeContext {
	return v.value
}

func (v *NullablePolicyConfigChangeContext) Set(val *PolicyConfigChangeContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigChangeContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigChangeContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigChangeContext(val *PolicyConfigChangeContext) *NullablePolicyConfigChangeContext {
	return &NullablePolicyConfigChangeContext{value: val, isSet: true}
}

func (v NullablePolicyConfigChangeContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigChangeContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
