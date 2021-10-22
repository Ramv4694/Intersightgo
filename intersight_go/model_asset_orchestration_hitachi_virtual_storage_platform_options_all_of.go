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

// AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf Definition of the list of properties defined in 'asset.OrchestrationHitachiVirtualStoragePlatformOptions', excluding properties defined in parent classes.
type AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The DNS hostname or IP address of the Hitachi Ops Center API Configuration Manager used to manage the Hitachi Virtual Storage Platform.
	OpsCenterManagementAddress *string `json:"OpsCenterManagementAddress,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf

// NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf(classId string, objectType string) *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf {
	this := AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOfWithDefaults instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOfWithDefaults() *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf {
	this := AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf{}
	var classId string = "asset.OrchestrationHitachiVirtualStoragePlatformOptions"
	this.ClassId = classId
	var objectType string = "asset.OrchestrationHitachiVirtualStoragePlatformOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOpsCenterManagementAddress returns the OpsCenterManagementAddress field value if set, zero value otherwise.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetOpsCenterManagementAddress() string {
	if o == nil || o.OpsCenterManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.OpsCenterManagementAddress
}

// GetOpsCenterManagementAddressOk returns a tuple with the OpsCenterManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetOpsCenterManagementAddressOk() (*string, bool) {
	if o == nil || o.OpsCenterManagementAddress == nil {
		return nil, false
	}
	return o.OpsCenterManagementAddress, true
}

// HasOpsCenterManagementAddress returns a boolean if a field has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) HasOpsCenterManagementAddress() bool {
	if o != nil && o.OpsCenterManagementAddress != nil {
		return true
	}

	return false
}

// SetOpsCenterManagementAddress gets a reference to the given string and assigns it to the OpsCenterManagementAddress field.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetOpsCenterManagementAddress(v string) {
	o.OpsCenterManagementAddress = &v
}

func (o AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OpsCenterManagementAddress != nil {
		toSerialize["OpsCenterManagementAddress"] = o.OpsCenterManagementAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf := _AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf); err == nil {
		*o = AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf(varAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OpsCenterManagementAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf struct {
	value *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf
	isSet bool
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) Get() *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf {
	return v.value
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) Set(val *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf(val *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf {
	return &NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
