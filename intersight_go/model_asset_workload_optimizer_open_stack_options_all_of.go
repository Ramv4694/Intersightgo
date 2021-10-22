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

// AssetWorkloadOptimizerOpenStackOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerOpenStackOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerOpenStackOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// OpenStack Identity Service (keystone) domain name. Domain is an additional namespaces you can create in keystone to partition users, groups, and projects. Default domain name value is \"Default\".
	Domain *string `json:"Domain,omitempty"`
	// The name of tenant which has assigned administrator role this OpenStack target user is in. A tenant or project is referred to as a group of users of OpenStack. Each tenant can be assigned a role which gives all its member users their rights and privileges to perform a specific set of operations.
	Tenant               *string `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerOpenStackOptionsAllOf AssetWorkloadOptimizerOpenStackOptionsAllOf

// NewAssetWorkloadOptimizerOpenStackOptionsAllOf instantiates a new AssetWorkloadOptimizerOpenStackOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerOpenStackOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerOpenStackOptionsAllOf {
	this := AssetWorkloadOptimizerOpenStackOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerOpenStackOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerOpenStackOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerOpenStackOptionsAllOfWithDefaults() *AssetWorkloadOptimizerOpenStackOptionsAllOf {
	this := AssetWorkloadOptimizerOpenStackOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerOpenStackOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerOpenStackOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) SetDomain(v string) {
	o.Domain = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) SetTenant(v string) {
	o.Tenant = &v
}

func (o AssetWorkloadOptimizerOpenStackOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Domain != nil {
		toSerialize["Domain"] = o.Domain
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerOpenStackOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerOpenStackOptionsAllOf := _AssetWorkloadOptimizerOpenStackOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerOpenStackOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerOpenStackOptionsAllOf(varAssetWorkloadOptimizerOpenStackOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerOpenStackOptionsAllOf struct {
	value *AssetWorkloadOptimizerOpenStackOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) Get() *AssetWorkloadOptimizerOpenStackOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) Set(val *AssetWorkloadOptimizerOpenStackOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerOpenStackOptionsAllOf(val *AssetWorkloadOptimizerOpenStackOptionsAllOf) *NullableAssetWorkloadOptimizerOpenStackOptionsAllOf {
	return &NullableAssetWorkloadOptimizerOpenStackOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerOpenStackOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
