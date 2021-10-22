/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-10-09T21:18:32Z.
 *
 * API version: 1.0.9-4809
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesVirtualMachineInfraConfigPolicyAllOf Definition of the list of properties defined in 'kubernetes.VirtualMachineInfraConfigPolicy', excluding properties defined in parent classes.
type KubernetesVirtualMachineInfraConfigPolicyAllOf struct {
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

type _KubernetesVirtualMachineInfraConfigPolicyAllOf KubernetesVirtualMachineInfraConfigPolicyAllOf

// NewKubernetesVirtualMachineInfraConfigPolicyAllOf instantiates a new KubernetesVirtualMachineInfraConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineInfraConfigPolicyAllOf(classId string, objectType string) *KubernetesVirtualMachineInfraConfigPolicyAllOf {
	this := KubernetesVirtualMachineInfraConfigPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVirtualMachineInfraConfigPolicyAllOfWithDefaults instantiates a new KubernetesVirtualMachineInfraConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineInfraConfigPolicyAllOfWithDefaults() *KubernetesVirtualMachineInfraConfigPolicyAllOf {
	this := KubernetesVirtualMachineInfraConfigPolicyAllOf{}
	var classId string = "kubernetes.VirtualMachineInfraConfigPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInfraConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVmConfig returns the VmConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetVmConfig() KubernetesBaseVirtualMachineInfraConfig {
	if o == nil || o.VmConfig.Get() == nil {
		var ret KubernetesBaseVirtualMachineInfraConfig
		return ret
	}
	return *o.VmConfig.Get()
}

// GetVmConfigOk returns a tuple with the VmConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetVmConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmConfig.Get(), o.VmConfig.IsSet()
}

// HasVmConfig returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) HasVmConfig() bool {
	if o != nil && o.VmConfig.IsSet() {
		return true
	}

	return false
}

// SetVmConfig gets a reference to the given NullableKubernetesBaseVirtualMachineInfraConfig and assigns it to the VmConfig field.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetVmConfig(v KubernetesBaseVirtualMachineInfraConfig) {
	o.VmConfig.Set(&v)
}

// SetVmConfigNil sets the value for VmConfig to be an explicit nil
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetVmConfigNil() {
	o.VmConfig.Set(nil)
}

// UnsetVmConfig ensures that no value is present for VmConfig, not even an explicit nil
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) UnsetVmConfig() {
	o.VmConfig.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship {
	if o == nil {
		var ret []KubernetesVirtualMachineInfrastructureProviderRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesVirtualMachineInfrastructureProviderRelationship and assigns it to the Profiles field.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship) {
	o.Profiles = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || o.Target == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target = &v
}

func (o KubernetesVirtualMachineInfraConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *KubernetesVirtualMachineInfraConfigPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesVirtualMachineInfraConfigPolicyAllOf := _KubernetesVirtualMachineInfraConfigPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInfraConfigPolicyAllOf); err == nil {
		*o = KubernetesVirtualMachineInfraConfigPolicyAllOf(varKubernetesVirtualMachineInfraConfigPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VmConfig")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "Target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesVirtualMachineInfraConfigPolicyAllOf struct {
	value *KubernetesVirtualMachineInfraConfigPolicyAllOf
	isSet bool
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) Get() *KubernetesVirtualMachineInfraConfigPolicyAllOf {
	return v.value
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) Set(val *KubernetesVirtualMachineInfraConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineInfraConfigPolicyAllOf(val *KubernetesVirtualMachineInfraConfigPolicyAllOf) *NullableKubernetesVirtualMachineInfraConfigPolicyAllOf {
	return &NullableKubernetesVirtualMachineInfraConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineInfraConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
