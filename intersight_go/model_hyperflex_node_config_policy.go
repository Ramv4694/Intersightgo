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

// HyperflexNodeConfigPolicy A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
type HyperflexNodeConfigPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType               string                       `json:"ObjectType"`
	DataIpRange              NullableHyperflexIpAddrRange `json:"DataIpRange,omitempty"`
	HxdpIpRange              NullableHyperflexIpAddrRange `json:"HxdpIpRange,omitempty"`
	HypervisorControlIpRange NullableHyperflexIpAddrRange `json:"HypervisorControlIpRange,omitempty"`
	MgmtIpRange              NullableHyperflexIpAddrRange `json:"MgmtIpRange,omitempty"`
	// The node name prefix that is used to automatically generate the default hostname for each server. A dash (-) will be appended to the prefix followed by the node number to form a hostname. This default naming scheme can be manually overridden in the node configuration. The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must start with an alphanumeric character.
	NodeNamePrefix *string `json:"NodeNamePrefix,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNodeConfigPolicy HyperflexNodeConfigPolicy

// NewHyperflexNodeConfigPolicy instantiates a new HyperflexNodeConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNodeConfigPolicy(classId string, objectType string) *HyperflexNodeConfigPolicy {
	this := HyperflexNodeConfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexNodeConfigPolicyWithDefaults instantiates a new HyperflexNodeConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeConfigPolicyWithDefaults() *HyperflexNodeConfigPolicy {
	this := HyperflexNodeConfigPolicy{}
	var classId string = "hyperflex.NodeConfigPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.NodeConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNodeConfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNodeConfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNodeConfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNodeConfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataIpRange returns the DataIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeConfigPolicy) GetDataIpRange() HyperflexIpAddrRange {
	if o == nil || o.DataIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.DataIpRange.Get()
}

// GetDataIpRangeOk returns a tuple with the DataIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeConfigPolicy) GetDataIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataIpRange.Get(), o.DataIpRange.IsSet()
}

// HasDataIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasDataIpRange() bool {
	if o != nil && o.DataIpRange.IsSet() {
		return true
	}

	return false
}

// SetDataIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the DataIpRange field.
func (o *HyperflexNodeConfigPolicy) SetDataIpRange(v HyperflexIpAddrRange) {
	o.DataIpRange.Set(&v)
}

// SetDataIpRangeNil sets the value for DataIpRange to be an explicit nil
func (o *HyperflexNodeConfigPolicy) SetDataIpRangeNil() {
	o.DataIpRange.Set(nil)
}

// UnsetDataIpRange ensures that no value is present for DataIpRange, not even an explicit nil
func (o *HyperflexNodeConfigPolicy) UnsetDataIpRange() {
	o.DataIpRange.Unset()
}

// GetHxdpIpRange returns the HxdpIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeConfigPolicy) GetHxdpIpRange() HyperflexIpAddrRange {
	if o == nil || o.HxdpIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.HxdpIpRange.Get()
}

// GetHxdpIpRangeOk returns a tuple with the HxdpIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeConfigPolicy) GetHxdpIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.HxdpIpRange.Get(), o.HxdpIpRange.IsSet()
}

// HasHxdpIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasHxdpIpRange() bool {
	if o != nil && o.HxdpIpRange.IsSet() {
		return true
	}

	return false
}

// SetHxdpIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the HxdpIpRange field.
func (o *HyperflexNodeConfigPolicy) SetHxdpIpRange(v HyperflexIpAddrRange) {
	o.HxdpIpRange.Set(&v)
}

// SetHxdpIpRangeNil sets the value for HxdpIpRange to be an explicit nil
func (o *HyperflexNodeConfigPolicy) SetHxdpIpRangeNil() {
	o.HxdpIpRange.Set(nil)
}

// UnsetHxdpIpRange ensures that no value is present for HxdpIpRange, not even an explicit nil
func (o *HyperflexNodeConfigPolicy) UnsetHxdpIpRange() {
	o.HxdpIpRange.Unset()
}

// GetHypervisorControlIpRange returns the HypervisorControlIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeConfigPolicy) GetHypervisorControlIpRange() HyperflexIpAddrRange {
	if o == nil || o.HypervisorControlIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.HypervisorControlIpRange.Get()
}

// GetHypervisorControlIpRangeOk returns a tuple with the HypervisorControlIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeConfigPolicy) GetHypervisorControlIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.HypervisorControlIpRange.Get(), o.HypervisorControlIpRange.IsSet()
}

// HasHypervisorControlIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasHypervisorControlIpRange() bool {
	if o != nil && o.HypervisorControlIpRange.IsSet() {
		return true
	}

	return false
}

// SetHypervisorControlIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the HypervisorControlIpRange field.
func (o *HyperflexNodeConfigPolicy) SetHypervisorControlIpRange(v HyperflexIpAddrRange) {
	o.HypervisorControlIpRange.Set(&v)
}

// SetHypervisorControlIpRangeNil sets the value for HypervisorControlIpRange to be an explicit nil
func (o *HyperflexNodeConfigPolicy) SetHypervisorControlIpRangeNil() {
	o.HypervisorControlIpRange.Set(nil)
}

// UnsetHypervisorControlIpRange ensures that no value is present for HypervisorControlIpRange, not even an explicit nil
func (o *HyperflexNodeConfigPolicy) UnsetHypervisorControlIpRange() {
	o.HypervisorControlIpRange.Unset()
}

// GetMgmtIpRange returns the MgmtIpRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeConfigPolicy) GetMgmtIpRange() HyperflexIpAddrRange {
	if o == nil || o.MgmtIpRange.Get() == nil {
		var ret HyperflexIpAddrRange
		return ret
	}
	return *o.MgmtIpRange.Get()
}

// GetMgmtIpRangeOk returns a tuple with the MgmtIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeConfigPolicy) GetMgmtIpRangeOk() (*HyperflexIpAddrRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MgmtIpRange.Get(), o.MgmtIpRange.IsSet()
}

// HasMgmtIpRange returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasMgmtIpRange() bool {
	if o != nil && o.MgmtIpRange.IsSet() {
		return true
	}

	return false
}

// SetMgmtIpRange gets a reference to the given NullableHyperflexIpAddrRange and assigns it to the MgmtIpRange field.
func (o *HyperflexNodeConfigPolicy) SetMgmtIpRange(v HyperflexIpAddrRange) {
	o.MgmtIpRange.Set(&v)
}

// SetMgmtIpRangeNil sets the value for MgmtIpRange to be an explicit nil
func (o *HyperflexNodeConfigPolicy) SetMgmtIpRangeNil() {
	o.MgmtIpRange.Set(nil)
}

// UnsetMgmtIpRange ensures that no value is present for MgmtIpRange, not even an explicit nil
func (o *HyperflexNodeConfigPolicy) UnsetMgmtIpRange() {
	o.MgmtIpRange.Unset()
}

// GetNodeNamePrefix returns the NodeNamePrefix field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicy) GetNodeNamePrefix() string {
	if o == nil || o.NodeNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.NodeNamePrefix
}

// GetNodeNamePrefixOk returns a tuple with the NodeNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicy) GetNodeNamePrefixOk() (*string, bool) {
	if o == nil || o.NodeNamePrefix == nil {
		return nil, false
	}
	return o.NodeNamePrefix, true
}

// HasNodeNamePrefix returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasNodeNamePrefix() bool {
	if o != nil && o.NodeNamePrefix != nil {
		return true
	}

	return false
}

// SetNodeNamePrefix gets a reference to the given string and assigns it to the NodeNamePrefix field.
func (o *HyperflexNodeConfigPolicy) SetNodeNamePrefix(v string) {
	o.NodeNamePrefix = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeConfigPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeConfigPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexNodeConfigPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexNodeConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexNodeConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexNodeConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexNodeConfigPolicy) MarshalJSON() ([]byte, error) {
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
	if o.DataIpRange.IsSet() {
		toSerialize["DataIpRange"] = o.DataIpRange.Get()
	}
	if o.HxdpIpRange.IsSet() {
		toSerialize["HxdpIpRange"] = o.HxdpIpRange.Get()
	}
	if o.HypervisorControlIpRange.IsSet() {
		toSerialize["HypervisorControlIpRange"] = o.HypervisorControlIpRange.Get()
	}
	if o.MgmtIpRange.IsSet() {
		toSerialize["MgmtIpRange"] = o.MgmtIpRange.Get()
	}
	if o.NodeNamePrefix != nil {
		toSerialize["NodeNamePrefix"] = o.NodeNamePrefix
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

func (o *HyperflexNodeConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexNodeConfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType               string                       `json:"ObjectType"`
		DataIpRange              NullableHyperflexIpAddrRange `json:"DataIpRange,omitempty"`
		HxdpIpRange              NullableHyperflexIpAddrRange `json:"HxdpIpRange,omitempty"`
		HypervisorControlIpRange NullableHyperflexIpAddrRange `json:"HypervisorControlIpRange,omitempty"`
		MgmtIpRange              NullableHyperflexIpAddrRange `json:"MgmtIpRange,omitempty"`
		// The node name prefix that is used to automatically generate the default hostname for each server. A dash (-) will be appended to the prefix followed by the node number to form a hostname. This default naming scheme can be manually overridden in the node configuration. The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must start with an alphanumeric character.
		NodeNamePrefix *string `json:"NodeNamePrefix,omitempty"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varHyperflexNodeConfigPolicyWithoutEmbeddedStruct := HyperflexNodeConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexNodeConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNodeConfigPolicy := _HyperflexNodeConfigPolicy{}
		varHyperflexNodeConfigPolicy.ClassId = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.ClassId
		varHyperflexNodeConfigPolicy.ObjectType = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.ObjectType
		varHyperflexNodeConfigPolicy.DataIpRange = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.DataIpRange
		varHyperflexNodeConfigPolicy.HxdpIpRange = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.HxdpIpRange
		varHyperflexNodeConfigPolicy.HypervisorControlIpRange = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.HypervisorControlIpRange
		varHyperflexNodeConfigPolicy.MgmtIpRange = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.MgmtIpRange
		varHyperflexNodeConfigPolicy.NodeNamePrefix = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.NodeNamePrefix
		varHyperflexNodeConfigPolicy.ClusterProfiles = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.ClusterProfiles
		varHyperflexNodeConfigPolicy.Organization = varHyperflexNodeConfigPolicyWithoutEmbeddedStruct.Organization
		*o = HyperflexNodeConfigPolicy(varHyperflexNodeConfigPolicy)
	} else {
		return err
	}

	varHyperflexNodeConfigPolicy := _HyperflexNodeConfigPolicy{}

	err = json.Unmarshal(bytes, &varHyperflexNodeConfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varHyperflexNodeConfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataIpRange")
		delete(additionalProperties, "HxdpIpRange")
		delete(additionalProperties, "HypervisorControlIpRange")
		delete(additionalProperties, "MgmtIpRange")
		delete(additionalProperties, "NodeNamePrefix")
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

type NullableHyperflexNodeConfigPolicy struct {
	value *HyperflexNodeConfigPolicy
	isSet bool
}

func (v NullableHyperflexNodeConfigPolicy) Get() *HyperflexNodeConfigPolicy {
	return v.value
}

func (v *NullableHyperflexNodeConfigPolicy) Set(val *HyperflexNodeConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNodeConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNodeConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNodeConfigPolicy(val *HyperflexNodeConfigPolicy) *NullableHyperflexNodeConfigPolicy {
	return &NullableHyperflexNodeConfigPolicy{value: val, isSet: true}
}

func (v NullableHyperflexNodeConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNodeConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
