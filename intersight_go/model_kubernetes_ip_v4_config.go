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

// KubernetesIpV4Config Network interface configuration data for IPv4 interfaces.
type KubernetesIpV4Config struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IPv4 Address in CIDR format.
	Ip                   *string  `json:"Ip,omitempty"`
	Lease                *MoMoRef `json:"Lease,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesIpV4Config KubernetesIpV4Config

// NewKubernetesIpV4Config instantiates a new KubernetesIpV4Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIpV4Config(classId string, objectType string) *KubernetesIpV4Config {
	this := KubernetesIpV4Config{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesIpV4ConfigWithDefaults instantiates a new KubernetesIpV4Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIpV4ConfigWithDefaults() *KubernetesIpV4Config {
	this := KubernetesIpV4Config{}
	var classId string = "kubernetes.IpV4Config"
	this.ClassId = classId
	var objectType string = "kubernetes.IpV4Config"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesIpV4Config) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesIpV4Config) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesIpV4Config) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesIpV4Config) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *KubernetesIpV4Config) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *KubernetesIpV4Config) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *KubernetesIpV4Config) SetIp(v string) {
	o.Ip = &v
}

// GetLease returns the Lease field value if set, zero value otherwise.
func (o *KubernetesIpV4Config) GetLease() MoMoRef {
	if o == nil || o.Lease == nil {
		var ret MoMoRef
		return ret
	}
	return *o.Lease
}

// GetLeaseOk returns a tuple with the Lease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIpV4Config) GetLeaseOk() (*MoMoRef, bool) {
	if o == nil || o.Lease == nil {
		return nil, false
	}
	return o.Lease, true
}

// HasLease returns a boolean if a field has been set.
func (o *KubernetesIpV4Config) HasLease() bool {
	if o != nil && o.Lease != nil {
		return true
	}

	return false
}

// SetLease gets a reference to the given MoMoRef and assigns it to the Lease field.
func (o *KubernetesIpV4Config) SetLease(v MoMoRef) {
	o.Lease = &v
}

func (o KubernetesIpV4Config) MarshalJSON() ([]byte, error) {
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
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.Lease != nil {
		toSerialize["Lease"] = o.Lease
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesIpV4Config) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesIpV4ConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IPv4 Address in CIDR format.
		Ip    *string  `json:"Ip,omitempty"`
		Lease *MoMoRef `json:"Lease,omitempty"`
	}

	varKubernetesIpV4ConfigWithoutEmbeddedStruct := KubernetesIpV4ConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesIpV4ConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesIpV4Config := _KubernetesIpV4Config{}
		varKubernetesIpV4Config.ClassId = varKubernetesIpV4ConfigWithoutEmbeddedStruct.ClassId
		varKubernetesIpV4Config.ObjectType = varKubernetesIpV4ConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesIpV4Config.Ip = varKubernetesIpV4ConfigWithoutEmbeddedStruct.Ip
		varKubernetesIpV4Config.Lease = varKubernetesIpV4ConfigWithoutEmbeddedStruct.Lease
		*o = KubernetesIpV4Config(varKubernetesIpV4Config)
	} else {
		return err
	}

	varKubernetesIpV4Config := _KubernetesIpV4Config{}

	err = json.Unmarshal(bytes, &varKubernetesIpV4Config)
	if err == nil {
		o.MoBaseComplexType = varKubernetesIpV4Config.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Lease")

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

type NullableKubernetesIpV4Config struct {
	value *KubernetesIpV4Config
	isSet bool
}

func (v NullableKubernetesIpV4Config) Get() *KubernetesIpV4Config {
	return v.value
}

func (v *NullableKubernetesIpV4Config) Set(val *KubernetesIpV4Config) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIpV4Config) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIpV4Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIpV4Config(val *KubernetesIpV4Config) *NullableKubernetesIpV4Config {
	return &NullableKubernetesIpV4Config{value: val, isSet: true}
}

func (v NullableKubernetesIpV4Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIpV4Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
