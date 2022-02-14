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

// OsIpv6Configuration In case of static IPv6 configuration, IP address, netmask and gateway details are provided.
type OsIpv6Configuration struct {
	OsIpConfiguration
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                    `json:"ObjectType"`
	IpV6Config           NullableCommIpV6Interface `json:"IpV6Config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsIpv6Configuration OsIpv6Configuration

// NewOsIpv6Configuration instantiates a new OsIpv6Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsIpv6Configuration(classId string, objectType string) *OsIpv6Configuration {
	this := OsIpv6Configuration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsIpv6ConfigurationWithDefaults instantiates a new OsIpv6Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsIpv6ConfigurationWithDefaults() *OsIpv6Configuration {
	this := OsIpv6Configuration{}
	var classId string = "os.Ipv6Configuration"
	this.ClassId = classId
	var objectType string = "os.Ipv6Configuration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsIpv6Configuration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsIpv6Configuration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsIpv6Configuration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsIpv6Configuration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsIpv6Configuration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsIpv6Configuration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpV6Config returns the IpV6Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsIpv6Configuration) GetIpV6Config() CommIpV6Interface {
	if o == nil || o.IpV6Config.Get() == nil {
		var ret CommIpV6Interface
		return ret
	}
	return *o.IpV6Config.Get()
}

// GetIpV6ConfigOk returns a tuple with the IpV6Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsIpv6Configuration) GetIpV6ConfigOk() (*CommIpV6Interface, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV6Config.Get(), o.IpV6Config.IsSet()
}

// HasIpV6Config returns a boolean if a field has been set.
func (o *OsIpv6Configuration) HasIpV6Config() bool {
	if o != nil && o.IpV6Config.IsSet() {
		return true
	}

	return false
}

// SetIpV6Config gets a reference to the given NullableCommIpV6Interface and assigns it to the IpV6Config field.
func (o *OsIpv6Configuration) SetIpV6Config(v CommIpV6Interface) {
	o.IpV6Config.Set(&v)
}

// SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil
func (o *OsIpv6Configuration) SetIpV6ConfigNil() {
	o.IpV6Config.Set(nil)
}

// UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
func (o *OsIpv6Configuration) UnsetIpV6Config() {
	o.IpV6Config.Unset()
}

func (o OsIpv6Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsIpConfiguration, errOsIpConfiguration := json.Marshal(o.OsIpConfiguration)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	errOsIpConfiguration = json.Unmarshal([]byte(serializedOsIpConfiguration), &toSerialize)
	if errOsIpConfiguration != nil {
		return []byte{}, errOsIpConfiguration
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpV6Config.IsSet() {
		toSerialize["IpV6Config"] = o.IpV6Config.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsIpv6Configuration) UnmarshalJSON(bytes []byte) (err error) {
	type OsIpv6ConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                    `json:"ObjectType"`
		IpV6Config NullableCommIpV6Interface `json:"IpV6Config,omitempty"`
	}

	varOsIpv6ConfigurationWithoutEmbeddedStruct := OsIpv6ConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsIpv6ConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varOsIpv6Configuration := _OsIpv6Configuration{}
		varOsIpv6Configuration.ClassId = varOsIpv6ConfigurationWithoutEmbeddedStruct.ClassId
		varOsIpv6Configuration.ObjectType = varOsIpv6ConfigurationWithoutEmbeddedStruct.ObjectType
		varOsIpv6Configuration.IpV6Config = varOsIpv6ConfigurationWithoutEmbeddedStruct.IpV6Config
		*o = OsIpv6Configuration(varOsIpv6Configuration)
	} else {
		return err
	}

	varOsIpv6Configuration := _OsIpv6Configuration{}

	err = json.Unmarshal(bytes, &varOsIpv6Configuration)
	if err == nil {
		o.OsIpConfiguration = varOsIpv6Configuration.OsIpConfiguration
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpV6Config")

		// remove fields from embedded structs
		reflectOsIpConfiguration := reflect.ValueOf(o.OsIpConfiguration)
		for i := 0; i < reflectOsIpConfiguration.Type().NumField(); i++ {
			t := reflectOsIpConfiguration.Type().Field(i)

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

type NullableOsIpv6Configuration struct {
	value *OsIpv6Configuration
	isSet bool
}

func (v NullableOsIpv6Configuration) Get() *OsIpv6Configuration {
	return v.value
}

func (v *NullableOsIpv6Configuration) Set(val *OsIpv6Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableOsIpv6Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableOsIpv6Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsIpv6Configuration(val *OsIpv6Configuration) *NullableOsIpv6Configuration {
	return &NullableOsIpv6Configuration{value: val, isSet: true}
}

func (v NullableOsIpv6Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsIpv6Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
