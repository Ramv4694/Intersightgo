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

// CapabilityCimcFirmwareDescriptor Descriptor that identifies the server's redfish integration capability using cimc firmware info.
type CapabilityCimcFirmwareDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the server uses ep proxy to communicate with the adapter.
	AdapterEpProxyEnabled *bool `json:"AdapterEpProxyEnabled,omitempty"`
	// Revision information for the server.
	Revision *string `json:"Revision,omitempty"`
	// Minimum server firmware version for UUID feature support.
	UuidSupportedVer     *string `json:"UuidSupportedVer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityCimcFirmwareDescriptor CapabilityCimcFirmwareDescriptor

// NewCapabilityCimcFirmwareDescriptor instantiates a new CapabilityCimcFirmwareDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityCimcFirmwareDescriptor(classId string, objectType string) *CapabilityCimcFirmwareDescriptor {
	this := CapabilityCimcFirmwareDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityCimcFirmwareDescriptorWithDefaults instantiates a new CapabilityCimcFirmwareDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityCimcFirmwareDescriptorWithDefaults() *CapabilityCimcFirmwareDescriptor {
	this := CapabilityCimcFirmwareDescriptor{}
	var classId string = "capability.CimcFirmwareDescriptor"
	this.ClassId = classId
	var objectType string = "capability.CimcFirmwareDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityCimcFirmwareDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityCimcFirmwareDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityCimcFirmwareDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityCimcFirmwareDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterEpProxyEnabled returns the AdapterEpProxyEnabled field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabled() bool {
	if o == nil || o.AdapterEpProxyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AdapterEpProxyEnabled
}

// GetAdapterEpProxyEnabledOk returns a tuple with the AdapterEpProxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetAdapterEpProxyEnabledOk() (*bool, bool) {
	if o == nil || o.AdapterEpProxyEnabled == nil {
		return nil, false
	}
	return o.AdapterEpProxyEnabled, true
}

// HasAdapterEpProxyEnabled returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasAdapterEpProxyEnabled() bool {
	if o != nil && o.AdapterEpProxyEnabled != nil {
		return true
	}

	return false
}

// SetAdapterEpProxyEnabled gets a reference to the given bool and assigns it to the AdapterEpProxyEnabled field.
func (o *CapabilityCimcFirmwareDescriptor) SetAdapterEpProxyEnabled(v bool) {
	o.AdapterEpProxyEnabled = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilityCimcFirmwareDescriptor) SetRevision(v string) {
	o.Revision = &v
}

// GetUuidSupportedVer returns the UuidSupportedVer field value if set, zero value otherwise.
func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVer() string {
	if o == nil || o.UuidSupportedVer == nil {
		var ret string
		return ret
	}
	return *o.UuidSupportedVer
}

// GetUuidSupportedVerOk returns a tuple with the UuidSupportedVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCimcFirmwareDescriptor) GetUuidSupportedVerOk() (*string, bool) {
	if o == nil || o.UuidSupportedVer == nil {
		return nil, false
	}
	return o.UuidSupportedVer, true
}

// HasUuidSupportedVer returns a boolean if a field has been set.
func (o *CapabilityCimcFirmwareDescriptor) HasUuidSupportedVer() bool {
	if o != nil && o.UuidSupportedVer != nil {
		return true
	}

	return false
}

// SetUuidSupportedVer gets a reference to the given string and assigns it to the UuidSupportedVer field.
func (o *CapabilityCimcFirmwareDescriptor) SetUuidSupportedVer(v string) {
	o.UuidSupportedVer = &v
}

func (o CapabilityCimcFirmwareDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterEpProxyEnabled != nil {
		toSerialize["AdapterEpProxyEnabled"] = o.AdapterEpProxyEnabled
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	if o.UuidSupportedVer != nil {
		toSerialize["UuidSupportedVer"] = o.UuidSupportedVer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityCimcFirmwareDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the server uses ep proxy to communicate with the adapter.
		AdapterEpProxyEnabled *bool `json:"AdapterEpProxyEnabled,omitempty"`
		// Revision information for the server.
		Revision *string `json:"Revision,omitempty"`
		// Minimum server firmware version for UUID feature support.
		UuidSupportedVer *string `json:"UuidSupportedVer,omitempty"`
	}

	varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct := CapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityCimcFirmwareDescriptor := _CapabilityCimcFirmwareDescriptor{}
		varCapabilityCimcFirmwareDescriptor.ClassId = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityCimcFirmwareDescriptor.ObjectType = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityCimcFirmwareDescriptor.AdapterEpProxyEnabled = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.AdapterEpProxyEnabled
		varCapabilityCimcFirmwareDescriptor.Revision = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.Revision
		varCapabilityCimcFirmwareDescriptor.UuidSupportedVer = varCapabilityCimcFirmwareDescriptorWithoutEmbeddedStruct.UuidSupportedVer
		*o = CapabilityCimcFirmwareDescriptor(varCapabilityCimcFirmwareDescriptor)
	} else {
		return err
	}

	varCapabilityCimcFirmwareDescriptor := _CapabilityCimcFirmwareDescriptor{}

	err = json.Unmarshal(bytes, &varCapabilityCimcFirmwareDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilityCimcFirmwareDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterEpProxyEnabled")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "UuidSupportedVer")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilityCimcFirmwareDescriptor struct {
	value *CapabilityCimcFirmwareDescriptor
	isSet bool
}

func (v NullableCapabilityCimcFirmwareDescriptor) Get() *CapabilityCimcFirmwareDescriptor {
	return v.value
}

func (v *NullableCapabilityCimcFirmwareDescriptor) Set(val *CapabilityCimcFirmwareDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityCimcFirmwareDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityCimcFirmwareDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityCimcFirmwareDescriptor(val *CapabilityCimcFirmwareDescriptor) *NullableCapabilityCimcFirmwareDescriptor {
	return &NullableCapabilityCimcFirmwareDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityCimcFirmwareDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityCimcFirmwareDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
