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

// KubernetesVirtualMachineInfraConfigPolicy A policy specifying compute, storage and network infrastructure configuration for a Virtual Machine.
type KubernetesVirtualMachineInfraConfigPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                          `json:"ObjectType"`
	VmConfig     NullableKubernetesBaseVirtualMachineInfraConfig `json:"VmConfig,omitempty"`
	Organization *OrganizationOrganizationRelationship           `json:"Organization,omitempty"`
	// An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources.
	Profiles             []KubernetesVirtualMachineInfrastructureProviderRelationship `json:"Profiles,omitempty"`
	Target               *AssetDeviceRegistrationRelationship                         `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineInfraConfigPolicy KubernetesVirtualMachineInfraConfigPolicy

// NewKubernetesVirtualMachineInfraConfigPolicy instantiates a new KubernetesVirtualMachineInfraConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineInfraConfigPolicy(classId string, objectType string) *KubernetesVirtualMachineInfraConfigPolicy {
	this := KubernetesVirtualMachineInfraConfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVirtualMachineInfraConfigPolicyWithDefaults instantiates a new KubernetesVirtualMachineInfraConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineInfraConfigPolicyWithDefaults() *KubernetesVirtualMachineInfraConfigPolicy {
	this := KubernetesVirtualMachineInfraConfigPolicy{}
	var classId string = "kubernetes.VirtualMachineInfraConfigPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInfraConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVmConfig returns the VmConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetVmConfig() KubernetesBaseVirtualMachineInfraConfig {
	if o == nil || o.VmConfig.Get() == nil {
		var ret KubernetesBaseVirtualMachineInfraConfig
		return ret
	}
	return *o.VmConfig.Get()
}

// GetVmConfigOk returns a tuple with the VmConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetVmConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmConfig.Get(), o.VmConfig.IsSet()
}

// HasVmConfig returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) HasVmConfig() bool {
	if o != nil && o.VmConfig.IsSet() {
		return true
	}

	return false
}

// SetVmConfig gets a reference to the given NullableKubernetesBaseVirtualMachineInfraConfig and assigns it to the VmConfig field.
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetVmConfig(v KubernetesBaseVirtualMachineInfraConfig) {
	o.VmConfig.Set(&v)
}

// SetVmConfigNil sets the value for VmConfig to be an explicit nil
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetVmConfigNil() {
	o.VmConfig.Set(nil)
}

// UnsetVmConfig ensures that no value is present for VmConfig, not even an explicit nil
func (o *KubernetesVirtualMachineInfraConfigPolicy) UnsetVmConfig() {
	o.VmConfig.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship {
	if o == nil {
		var ret []KubernetesVirtualMachineInfrastructureProviderRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesVirtualMachineInfrastructureProviderRelationship and assigns it to the Profiles field.
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship) {
	o.Profiles = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || o.Target == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicy) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesVirtualMachineInfraConfigPolicy) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target = &v
}

func (o KubernetesVirtualMachineInfraConfigPolicy) MarshalJSON() ([]byte, error) {
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
	if o.VmConfig.IsSet() {
		toSerialize["VmConfig"] = o.VmConfig.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineInfraConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                          `json:"ObjectType"`
		VmConfig     NullableKubernetesBaseVirtualMachineInfraConfig `json:"VmConfig,omitempty"`
		Organization *OrganizationOrganizationRelationship           `json:"Organization,omitempty"`
		// An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources.
		Profiles []KubernetesVirtualMachineInfrastructureProviderRelationship `json:"Profiles,omitempty"`
		Target   *AssetDeviceRegistrationRelationship                         `json:"Target,omitempty"`
	}

	varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct := KubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesVirtualMachineInfraConfigPolicy := _KubernetesVirtualMachineInfraConfigPolicy{}
		varKubernetesVirtualMachineInfraConfigPolicy.ClassId = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesVirtualMachineInfraConfigPolicy.ObjectType = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesVirtualMachineInfraConfigPolicy.VmConfig = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.VmConfig
		varKubernetesVirtualMachineInfraConfigPolicy.Organization = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.Organization
		varKubernetesVirtualMachineInfraConfigPolicy.Profiles = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.Profiles
		varKubernetesVirtualMachineInfraConfigPolicy.Target = varKubernetesVirtualMachineInfraConfigPolicyWithoutEmbeddedStruct.Target
		*o = KubernetesVirtualMachineInfraConfigPolicy(varKubernetesVirtualMachineInfraConfigPolicy)
	} else {
		return err
	}

	varKubernetesVirtualMachineInfraConfigPolicy := _KubernetesVirtualMachineInfraConfigPolicy{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInfraConfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesVirtualMachineInfraConfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VmConfig")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "Target")

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

type NullableKubernetesVirtualMachineInfraConfigPolicy struct {
	value *KubernetesVirtualMachineInfraConfigPolicy
	isSet bool
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicy) Get() *KubernetesVirtualMachineInfraConfigPolicy {
	return v.value
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicy) Set(val *KubernetesVirtualMachineInfraConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineInfraConfigPolicy(val *KubernetesVirtualMachineInfraConfigPolicy) *NullableKubernetesVirtualMachineInfraConfigPolicy {
	return &NullableKubernetesVirtualMachineInfraConfigPolicy{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
