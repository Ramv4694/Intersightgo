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

// KubernetesBaremetalNetworkInfo Baremetal network information used to config the operating system's network interfaces.
type KubernetesBaremetalNetworkInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string               `json:"ObjectType"`
	Ethernets            []KubernetesEthernet `json:"Ethernets,omitempty"`
	Ovsbonds             []KubernetesOvsBond  `json:"Ovsbonds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaremetalNetworkInfo KubernetesBaremetalNetworkInfo

// NewKubernetesBaremetalNetworkInfo instantiates a new KubernetesBaremetalNetworkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaremetalNetworkInfo(classId string, objectType string) *KubernetesBaremetalNetworkInfo {
	this := KubernetesBaremetalNetworkInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaremetalNetworkInfoWithDefaults instantiates a new KubernetesBaremetalNetworkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaremetalNetworkInfoWithDefaults() *KubernetesBaremetalNetworkInfo {
	this := KubernetesBaremetalNetworkInfo{}
	var classId string = "kubernetes.BaremetalNetworkInfo"
	this.ClassId = classId
	var objectType string = "kubernetes.BaremetalNetworkInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaremetalNetworkInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNetworkInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaremetalNetworkInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaremetalNetworkInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNetworkInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaremetalNetworkInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEthernets returns the Ethernets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesBaremetalNetworkInfo) GetEthernets() []KubernetesEthernet {
	if o == nil {
		var ret []KubernetesEthernet
		return ret
	}
	return o.Ethernets
}

// GetEthernetsOk returns a tuple with the Ethernets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesBaremetalNetworkInfo) GetEthernetsOk() (*[]KubernetesEthernet, bool) {
	if o == nil || o.Ethernets == nil {
		return nil, false
	}
	return &o.Ethernets, true
}

// HasEthernets returns a boolean if a field has been set.
func (o *KubernetesBaremetalNetworkInfo) HasEthernets() bool {
	if o != nil && o.Ethernets != nil {
		return true
	}

	return false
}

// SetEthernets gets a reference to the given []KubernetesEthernet and assigns it to the Ethernets field.
func (o *KubernetesBaremetalNetworkInfo) SetEthernets(v []KubernetesEthernet) {
	o.Ethernets = v
}

// GetOvsbonds returns the Ovsbonds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesBaremetalNetworkInfo) GetOvsbonds() []KubernetesOvsBond {
	if o == nil {
		var ret []KubernetesOvsBond
		return ret
	}
	return o.Ovsbonds
}

// GetOvsbondsOk returns a tuple with the Ovsbonds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesBaremetalNetworkInfo) GetOvsbondsOk() (*[]KubernetesOvsBond, bool) {
	if o == nil || o.Ovsbonds == nil {
		return nil, false
	}
	return &o.Ovsbonds, true
}

// HasOvsbonds returns a boolean if a field has been set.
func (o *KubernetesBaremetalNetworkInfo) HasOvsbonds() bool {
	if o != nil && o.Ovsbonds != nil {
		return true
	}

	return false
}

// SetOvsbonds gets a reference to the given []KubernetesOvsBond and assigns it to the Ovsbonds field.
func (o *KubernetesBaremetalNetworkInfo) SetOvsbonds(v []KubernetesOvsBond) {
	o.Ovsbonds = v
}

func (o KubernetesBaremetalNetworkInfo) MarshalJSON() ([]byte, error) {
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
	if o.Ethernets != nil {
		toSerialize["Ethernets"] = o.Ethernets
	}
	if o.Ovsbonds != nil {
		toSerialize["Ovsbonds"] = o.Ovsbonds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaremetalNetworkInfo) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesBaremetalNetworkInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string               `json:"ObjectType"`
		Ethernets  []KubernetesEthernet `json:"Ethernets,omitempty"`
		Ovsbonds   []KubernetesOvsBond  `json:"Ovsbonds,omitempty"`
	}

	varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct := KubernetesBaremetalNetworkInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesBaremetalNetworkInfo := _KubernetesBaremetalNetworkInfo{}
		varKubernetesBaremetalNetworkInfo.ClassId = varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct.ClassId
		varKubernetesBaremetalNetworkInfo.ObjectType = varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct.ObjectType
		varKubernetesBaremetalNetworkInfo.Ethernets = varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct.Ethernets
		varKubernetesBaremetalNetworkInfo.Ovsbonds = varKubernetesBaremetalNetworkInfoWithoutEmbeddedStruct.Ovsbonds
		*o = KubernetesBaremetalNetworkInfo(varKubernetesBaremetalNetworkInfo)
	} else {
		return err
	}

	varKubernetesBaremetalNetworkInfo := _KubernetesBaremetalNetworkInfo{}

	err = json.Unmarshal(bytes, &varKubernetesBaremetalNetworkInfo)
	if err == nil {
		o.MoBaseComplexType = varKubernetesBaremetalNetworkInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ethernets")
		delete(additionalProperties, "Ovsbonds")

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

type NullableKubernetesBaremetalNetworkInfo struct {
	value *KubernetesBaremetalNetworkInfo
	isSet bool
}

func (v NullableKubernetesBaremetalNetworkInfo) Get() *KubernetesBaremetalNetworkInfo {
	return v.value
}

func (v *NullableKubernetesBaremetalNetworkInfo) Set(val *KubernetesBaremetalNetworkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaremetalNetworkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaremetalNetworkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaremetalNetworkInfo(val *KubernetesBaremetalNetworkInfo) *NullableKubernetesBaremetalNetworkInfo {
	return &NullableKubernetesBaremetalNetworkInfo{value: val, isSet: true}
}

func (v NullableKubernetesBaremetalNetworkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaremetalNetworkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
