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

// StorageNetAppExportPolicy NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.
type StorageNetAppExportPolicy struct {
	StorageBaseNfsExport
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identity of the device.
	ClusterUuid            *string                         `json:"ClusterUuid,omitempty"`
	NetAppExportPolicyRule []StorageNetAppExportPolicyRule `json:"NetAppExportPolicyRule,omitempty"`
	// ID for the Export Policy.
	PolicyId             *int64                              `json:"PolicyId,omitempty"`
	Array                *StorageNetAppClusterRelationship   `json:"Array,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppExportPolicy StorageNetAppExportPolicy

// NewStorageNetAppExportPolicy instantiates a new StorageNetAppExportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppExportPolicy(classId string, objectType string) *StorageNetAppExportPolicy {
	this := StorageNetAppExportPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppExportPolicyWithDefaults instantiates a new StorageNetAppExportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppExportPolicyWithDefaults() *StorageNetAppExportPolicy {
	this := StorageNetAppExportPolicy{}
	var classId string = "storage.NetAppExportPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppExportPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppExportPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppExportPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppExportPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppExportPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppExportPolicy) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetNetAppExportPolicyRule returns the NetAppExportPolicyRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRule() []StorageNetAppExportPolicyRule {
	if o == nil {
		var ret []StorageNetAppExportPolicyRule
		return ret
	}
	return o.NetAppExportPolicyRule
}

// GetNetAppExportPolicyRuleOk returns a tuple with the NetAppExportPolicyRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRuleOk() (*[]StorageNetAppExportPolicyRule, bool) {
	if o == nil || o.NetAppExportPolicyRule == nil {
		return nil, false
	}
	return &o.NetAppExportPolicyRule, true
}

// HasNetAppExportPolicyRule returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasNetAppExportPolicyRule() bool {
	if o != nil && o.NetAppExportPolicyRule != nil {
		return true
	}

	return false
}

// SetNetAppExportPolicyRule gets a reference to the given []StorageNetAppExportPolicyRule and assigns it to the NetAppExportPolicyRule field.
func (o *StorageNetAppExportPolicy) SetNetAppExportPolicyRule(v []StorageNetAppExportPolicyRule) {
	o.NetAppExportPolicyRule = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetPolicyId() int64 {
	if o == nil || o.PolicyId == nil {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetPolicyIdOk() (*int64, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *StorageNetAppExportPolicy) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppExportPolicy) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicy) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicy) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicy) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppExportPolicy) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppExportPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseNfsExport, errStorageBaseNfsExport := json.Marshal(o.StorageBaseNfsExport)
	if errStorageBaseNfsExport != nil {
		return []byte{}, errStorageBaseNfsExport
	}
	errStorageBaseNfsExport = json.Unmarshal([]byte(serializedStorageBaseNfsExport), &toSerialize)
	if errStorageBaseNfsExport != nil {
		return []byte{}, errStorageBaseNfsExport
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.NetAppExportPolicyRule != nil {
		toSerialize["NetAppExportPolicyRule"] = o.NetAppExportPolicyRule
	}
	if o.PolicyId != nil {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppExportPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppExportPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identity of the device.
		ClusterUuid            *string                         `json:"ClusterUuid,omitempty"`
		NetAppExportPolicyRule []StorageNetAppExportPolicyRule `json:"NetAppExportPolicyRule,omitempty"`
		// ID for the Export Policy.
		PolicyId *int64                              `json:"PolicyId,omitempty"`
		Array    *StorageNetAppClusterRelationship   `json:"Array,omitempty"`
		Tenant   *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppExportPolicyWithoutEmbeddedStruct := StorageNetAppExportPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppExportPolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppExportPolicy := _StorageNetAppExportPolicy{}
		varStorageNetAppExportPolicy.ClassId = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ClassId
		varStorageNetAppExportPolicy.ObjectType = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppExportPolicy.ClusterUuid = varStorageNetAppExportPolicyWithoutEmbeddedStruct.ClusterUuid
		varStorageNetAppExportPolicy.NetAppExportPolicyRule = varStorageNetAppExportPolicyWithoutEmbeddedStruct.NetAppExportPolicyRule
		varStorageNetAppExportPolicy.PolicyId = varStorageNetAppExportPolicyWithoutEmbeddedStruct.PolicyId
		varStorageNetAppExportPolicy.Array = varStorageNetAppExportPolicyWithoutEmbeddedStruct.Array
		varStorageNetAppExportPolicy.Tenant = varStorageNetAppExportPolicyWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppExportPolicy(varStorageNetAppExportPolicy)
	} else {
		return err
	}

	varStorageNetAppExportPolicy := _StorageNetAppExportPolicy{}

	err = json.Unmarshal(bytes, &varStorageNetAppExportPolicy)
	if err == nil {
		o.StorageBaseNfsExport = varStorageNetAppExportPolicy.StorageBaseNfsExport
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "NetAppExportPolicyRule")
		delete(additionalProperties, "PolicyId")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBaseNfsExport := reflect.ValueOf(o.StorageBaseNfsExport)
		for i := 0; i < reflectStorageBaseNfsExport.Type().NumField(); i++ {
			t := reflectStorageBaseNfsExport.Type().Field(i)

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

type NullableStorageNetAppExportPolicy struct {
	value *StorageNetAppExportPolicy
	isSet bool
}

func (v NullableStorageNetAppExportPolicy) Get() *StorageNetAppExportPolicy {
	return v.value
}

func (v *NullableStorageNetAppExportPolicy) Set(val *StorageNetAppExportPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppExportPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppExportPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppExportPolicy(val *StorageNetAppExportPolicy) *NullableStorageNetAppExportPolicy {
	return &NullableStorageNetAppExportPolicy{value: val, isSet: true}
}

func (v NullableStorageNetAppExportPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppExportPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
