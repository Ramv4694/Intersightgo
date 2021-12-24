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
	"reflect"
	"strings"
)

// HyperflexClusterStoragePolicy A policy specifying HyperFlex cluster storage settings (optional).
type HyperflexClusterStoragePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, formats existing disk partitions (destroys all user data).
	DiskPartitionCleanup         *bool                                    `json:"DiskPartitionCleanup,omitempty"`
	LogicalAvalabilityZoneConfig NullableHyperflexLogicalAvailabilityZone `json:"LogicalAvalabilityZoneConfig,omitempty"`
	// Enable or disable VDI optimization (hybrid HyperFlex systems only).
	VdiOptimization *bool `json:"VdiOptimization,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterStoragePolicy HyperflexClusterStoragePolicy

// NewHyperflexClusterStoragePolicy instantiates a new HyperflexClusterStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterStoragePolicy(classId string, objectType string) *HyperflexClusterStoragePolicy {
	this := HyperflexClusterStoragePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var diskPartitionCleanup bool = true
	this.DiskPartitionCleanup = &diskPartitionCleanup
	return &this
}

// NewHyperflexClusterStoragePolicyWithDefaults instantiates a new HyperflexClusterStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterStoragePolicyWithDefaults() *HyperflexClusterStoragePolicy {
	this := HyperflexClusterStoragePolicy{}
	var classId string = "hyperflex.ClusterStoragePolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterStoragePolicy"
	this.ObjectType = objectType
	var diskPartitionCleanup bool = true
	this.DiskPartitionCleanup = &diskPartitionCleanup
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterStoragePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterStoragePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterStoragePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterStoragePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterStoragePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterStoragePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskPartitionCleanup returns the DiskPartitionCleanup field value if set, zero value otherwise.
func (o *HyperflexClusterStoragePolicy) GetDiskPartitionCleanup() bool {
	if o == nil || o.DiskPartitionCleanup == nil {
		var ret bool
		return ret
	}
	return *o.DiskPartitionCleanup
}

// GetDiskPartitionCleanupOk returns a tuple with the DiskPartitionCleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterStoragePolicy) GetDiskPartitionCleanupOk() (*bool, bool) {
	if o == nil || o.DiskPartitionCleanup == nil {
		return nil, false
	}
	return o.DiskPartitionCleanup, true
}

// HasDiskPartitionCleanup returns a boolean if a field has been set.
func (o *HyperflexClusterStoragePolicy) HasDiskPartitionCleanup() bool {
	if o != nil && o.DiskPartitionCleanup != nil {
		return true
	}

	return false
}

// SetDiskPartitionCleanup gets a reference to the given bool and assigns it to the DiskPartitionCleanup field.
func (o *HyperflexClusterStoragePolicy) SetDiskPartitionCleanup(v bool) {
	o.DiskPartitionCleanup = &v
}

// GetLogicalAvalabilityZoneConfig returns the LogicalAvalabilityZoneConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterStoragePolicy) GetLogicalAvalabilityZoneConfig() HyperflexLogicalAvailabilityZone {
	if o == nil || o.LogicalAvalabilityZoneConfig.Get() == nil {
		var ret HyperflexLogicalAvailabilityZone
		return ret
	}
	return *o.LogicalAvalabilityZoneConfig.Get()
}

// GetLogicalAvalabilityZoneConfigOk returns a tuple with the LogicalAvalabilityZoneConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterStoragePolicy) GetLogicalAvalabilityZoneConfigOk() (*HyperflexLogicalAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogicalAvalabilityZoneConfig.Get(), o.LogicalAvalabilityZoneConfig.IsSet()
}

// HasLogicalAvalabilityZoneConfig returns a boolean if a field has been set.
func (o *HyperflexClusterStoragePolicy) HasLogicalAvalabilityZoneConfig() bool {
	if o != nil && o.LogicalAvalabilityZoneConfig.IsSet() {
		return true
	}

	return false
}

// SetLogicalAvalabilityZoneConfig gets a reference to the given NullableHyperflexLogicalAvailabilityZone and assigns it to the LogicalAvalabilityZoneConfig field.
func (o *HyperflexClusterStoragePolicy) SetLogicalAvalabilityZoneConfig(v HyperflexLogicalAvailabilityZone) {
	o.LogicalAvalabilityZoneConfig.Set(&v)
}

// SetLogicalAvalabilityZoneConfigNil sets the value for LogicalAvalabilityZoneConfig to be an explicit nil
func (o *HyperflexClusterStoragePolicy) SetLogicalAvalabilityZoneConfigNil() {
	o.LogicalAvalabilityZoneConfig.Set(nil)
}

// UnsetLogicalAvalabilityZoneConfig ensures that no value is present for LogicalAvalabilityZoneConfig, not even an explicit nil
func (o *HyperflexClusterStoragePolicy) UnsetLogicalAvalabilityZoneConfig() {
	o.LogicalAvalabilityZoneConfig.Unset()
}

// GetVdiOptimization returns the VdiOptimization field value if set, zero value otherwise.
func (o *HyperflexClusterStoragePolicy) GetVdiOptimization() bool {
	if o == nil || o.VdiOptimization == nil {
		var ret bool
		return ret
	}
	return *o.VdiOptimization
}

// GetVdiOptimizationOk returns a tuple with the VdiOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterStoragePolicy) GetVdiOptimizationOk() (*bool, bool) {
	if o == nil || o.VdiOptimization == nil {
		return nil, false
	}
	return o.VdiOptimization, true
}

// HasVdiOptimization returns a boolean if a field has been set.
func (o *HyperflexClusterStoragePolicy) HasVdiOptimization() bool {
	if o != nil && o.VdiOptimization != nil {
		return true
	}

	return false
}

// SetVdiOptimization gets a reference to the given bool and assigns it to the VdiOptimization field.
func (o *HyperflexClusterStoragePolicy) SetVdiOptimization(v bool) {
	o.VdiOptimization = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterStoragePolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterStoragePolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexClusterStoragePolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexClusterStoragePolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexClusterStoragePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexClusterStoragePolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexClusterStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexClusterStoragePolicy) MarshalJSON() ([]byte, error) {
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
	if o.DiskPartitionCleanup != nil {
		toSerialize["DiskPartitionCleanup"] = o.DiskPartitionCleanup
	}
	if o.LogicalAvalabilityZoneConfig.IsSet() {
		toSerialize["LogicalAvalabilityZoneConfig"] = o.LogicalAvalabilityZoneConfig.Get()
	}
	if o.VdiOptimization != nil {
		toSerialize["VdiOptimization"] = o.VdiOptimization
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterStoragePolicy) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexClusterStoragePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If enabled, formats existing disk partitions (destroys all user data).
		DiskPartitionCleanup         *bool                                    `json:"DiskPartitionCleanup,omitempty"`
		LogicalAvalabilityZoneConfig NullableHyperflexLogicalAvailabilityZone `json:"LogicalAvalabilityZoneConfig,omitempty"`
		// Enable or disable VDI optimization (hybrid HyperFlex systems only).
		VdiOptimization *bool `json:"VdiOptimization,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexClusterStoragePolicyWithoutEmbeddedStruct := HyperflexClusterStoragePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexClusterStoragePolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterStoragePolicy := _HyperflexClusterStoragePolicy{}
		varHyperflexClusterStoragePolicy.ClassId = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.ClassId
		varHyperflexClusterStoragePolicy.ObjectType = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterStoragePolicy.DiskPartitionCleanup = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.DiskPartitionCleanup
		varHyperflexClusterStoragePolicy.LogicalAvalabilityZoneConfig = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.LogicalAvalabilityZoneConfig
		varHyperflexClusterStoragePolicy.VdiOptimization = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.VdiOptimization
		varHyperflexClusterStoragePolicy.ClusterProfiles = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexClusterStoragePolicy.Organization = varHyperflexClusterStoragePolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexClusterStoragePolicy(varHyperflexClusterStoragePolicy)
	} else {
		return err
	}

	varHyperflexClusterStoragePolicy := _HyperflexClusterStoragePolicy{}

	err = json.Unmarshal(bytes, &varHyperflexClusterStoragePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexClusterStoragePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskPartitionCleanup")
		delete(additionalProperties, "LogicalAvalabilityZoneConfig")
		delete(additionalProperties, "VdiOptimization")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")

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

type NullableHyperflexClusterStoragePolicy struct {
	value *HyperflexClusterStoragePolicy
	isSet bool
}

func (v NullableHyperflexClusterStoragePolicy) Get() *HyperflexClusterStoragePolicy {
	return v.value
}

func (v *NullableHyperflexClusterStoragePolicy) Set(val *HyperflexClusterStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterStoragePolicy(val *HyperflexClusterStoragePolicy) *NullableHyperflexClusterStoragePolicy {
	return &NullableHyperflexClusterStoragePolicy{value: val, isSet: true}
}

func (v NullableHyperflexClusterStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
