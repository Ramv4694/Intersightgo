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

// HyperflexUcsmConfigPolicyAllOf Definition of the list of properties defined in 'hyperflex.UcsmConfigPolicy', excluding properties defined in parent classes.
type HyperflexUcsmConfigPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string                              `json:"ObjectType"`
	KvmIpRange     NullableHyperflexIpAddrRange        `json:"KvmIpRange,omitempty"`
	MacPrefixRange NullableHyperflexMacAddrPrefixRange `json:"MacPrefixRange,omitempty"`
	// The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc.
	ServerFirmwareVersion *string `json:"ServerFirmwareVersion,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexUcsmConfigPolicyAllOf HyperflexUcsmConfigPolicyAllOf

// NewHyperflexUcsmConfigPolicyAllOf instantiates a new HyperflexUcsmConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexUcsmConfigPolicyAllOf(classId string, objectType string) *HyperflexUcsmConfigPolicyAllOf {
	this := HyperflexUcsmConfigPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexUcsmConfigPolicyAllOfWithDefaults instantiates a new HyperflexUcsmConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexUcsmConfigPolicyAllOfWithDefaults() *HyperflexUcsmConfigPolicyAllOf {
	this := HyperflexUcsmConfigPolicyAllOf{}
	var classId string = "hyperflex.UcsmConfigPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.UcsmConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexUcsmConfigPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexUcsmConfigPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexUcsmConfigPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexUcsmConfigPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKvmIpRange returns the KvmIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicyAllOf) GetKvmIpRange() HyperflexIpAddrRange {
	if o == nil || o.KvmIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.KvmIpRange.Get()
}

// GetKvmIpRangeOk returns a tuple with the KvmIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicyAllOf) GetKvmIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.KvmIpRange.Get(), o.KvmIpRange.IsSet()
}

// HasKvmIpRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasKvmIpRange() bool {
	if o != nil && o.KvmIpRange.IsSet() {
		return true
	}

	return false
}

// SetKvmIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the KvmIpRange field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetKvmIpRange(v HyperflexIpAddrRange) {
	o.KvmIpRange.Set(&v)
}

// SetKvmIpRangeNil sets the value for KvmIpRange to be an explicit nil
func (o *HyperflexUcsmConfigPolicyAllOf) SetKvmIpRangeNil() {
	o.KvmIpRange.Set(nil)
}

// UnsetKvmIpRange ensures that no value is present for KvmIpRange, not even an explicit nil
func (o *HyperflexUcsmConfigPolicyAllOf) UnsetKvmIpRange() {
	o.KvmIpRange.Unset()
}

// GetMacPrefixRange returns the MacPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicyAllOf) GetMacPrefixRange() HyperflexMacAddrPrefixRange {
	if o == nil || o.MacPrefixRange.Get() == nil {
		var ret HyperflexMacAddrPrefixRange
		return ret
	}
	return *o.MacPrefixRange.Get()
}

// GetMacPrefixRangeOk returns a tuple with the MacPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicyAllOf) GetMacPrefixRangeOk() (*HyperflexMacAddrPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacPrefixRange.Get(), o.MacPrefixRange.IsSet()
}

// HasMacPrefixRange returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasMacPrefixRange() bool {
	if o != nil && o.MacPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetMacPrefixRange gets a reference to the given NullableHyperflexMacAddrPrefixRange and assigns it to the MacPrefixRange field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetMacPrefixRange(v HyperflexMacAddrPrefixRange) {
	o.MacPrefixRange.Set(&v)
}

// SetMacPrefixRangeNil sets the value for MacPrefixRange to be an explicit nil
func (o *HyperflexUcsmConfigPolicyAllOf) SetMacPrefixRangeNil() {
	o.MacPrefixRange.Set(nil)
}

// UnsetMacPrefixRange ensures that no value is present for MacPrefixRange, not even an explicit nil
func (o *HyperflexUcsmConfigPolicyAllOf) UnsetMacPrefixRange() {
	o.MacPrefixRange.Unset()
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetServerFirmwareVersion() string {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetServerFirmwareVersionOk() (*string, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given string and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetServerFirmwareVersion(v string) {
	o.ServerFirmwareVersion = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexUcsmConfigPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexUcsmConfigPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexUcsmConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexUcsmConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexUcsmConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexUcsmConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KvmIpRange.IsSet() {
		toSerialize["KvmIpRange"] = o.KvmIpRange.Get()
	}
	if o.MacPrefixRange.IsSet() {
		toSerialize["MacPrefixRange"] = o.MacPrefixRange.Get()
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
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

func (o *HyperflexUcsmConfigPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexUcsmConfigPolicyAllOf := _HyperflexUcsmConfigPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexUcsmConfigPolicyAllOf); err == nil {
		*o = HyperflexUcsmConfigPolicyAllOf(varHyperflexUcsmConfigPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmIpRange")
		delete(additionalProperties, "MacPrefixRange")
		delete(additionalProperties, "ServerFirmwareVersion")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexUcsmConfigPolicyAllOf struct {
	value *HyperflexUcsmConfigPolicyAllOf
	isSet bool
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) Get() *HyperflexUcsmConfigPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) Set(val *HyperflexUcsmConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexUcsmConfigPolicyAllOf(val *HyperflexUcsmConfigPolicyAllOf) *NullableHyperflexUcsmConfigPolicyAllOf {
	return &NullableHyperflexUcsmConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexUcsmConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexUcsmConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
