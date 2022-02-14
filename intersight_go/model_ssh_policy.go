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

// SshPolicy Secure shell policy on the endpoint.
type SshPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of SSH service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// Port used for secure shell access.
	Port *int64 `json:"Port,omitempty"`
	// Number of seconds to wait before the system considers a SSH request to have timed out.
	Timeout      *int64                                `json:"Timeout,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SshPolicy SshPolicy

// NewSshPolicy instantiates a new SshPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshPolicy(classId string, objectType string) *SshPolicy {
	this := SshPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var port int64 = 22
	this.Port = &port
	var timeout int64 = 1800
	this.Timeout = &timeout
	return &this
}

// NewSshPolicyWithDefaults instantiates a new SshPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshPolicyWithDefaults() *SshPolicy {
	this := SshPolicy{}
	var classId string = "ssh.Policy"
	this.ClassId = classId
	var objectType string = "ssh.Policy"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var port int64 = 22
	this.Port = &port
	var timeout int64 = 1800
	this.Timeout = &timeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *SshPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SshPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SshPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SshPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SshPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SshPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SshPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SshPolicy) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SshPolicy) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SshPolicy) SetPort(v int64) {
	o.Port = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SshPolicy) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SshPolicy) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SshPolicy) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SshPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SshPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SshPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SshPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SshPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SshPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SshPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SshPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SshPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type SshPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of SSH service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// Port used for secure shell access.
		Port *int64 `json:"Port,omitempty"`
		// Number of seconds to wait before the system considers a SSH request to have timed out.
		Timeout      *int64                                `json:"Timeout,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varSshPolicyWithoutEmbeddedStruct := SshPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSshPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSshPolicy := _SshPolicy{}
		varSshPolicy.ClassId = varSshPolicyWithoutEmbeddedStruct.ClassId
		varSshPolicy.ObjectType = varSshPolicyWithoutEmbeddedStruct.ObjectType
		varSshPolicy.Enabled = varSshPolicyWithoutEmbeddedStruct.Enabled
		varSshPolicy.Port = varSshPolicyWithoutEmbeddedStruct.Port
		varSshPolicy.Timeout = varSshPolicyWithoutEmbeddedStruct.Timeout
		varSshPolicy.Organization = varSshPolicyWithoutEmbeddedStruct.Organization
		varSshPolicy.Profiles = varSshPolicyWithoutEmbeddedStruct.Profiles
		*o = SshPolicy(varSshPolicy)
	} else {
		return err
	}

	varSshPolicy := _SshPolicy{}

	err = json.Unmarshal(bytes, &varSshPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSshPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableSshPolicy struct {
	value *SshPolicy
	isSet bool
}

func (v NullableSshPolicy) Get() *SshPolicy {
	return v.value
}

func (v *NullableSshPolicy) Set(val *SshPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSshPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSshPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshPolicy(val *SshPolicy) *NullableSshPolicy {
	return &NullableSshPolicy{value: val, isSet: true}
}

func (v NullableSshPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
