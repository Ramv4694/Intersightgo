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

// KubernetesInstanceTypeDetails Instance Type details of the VM.
type KubernetesInstanceTypeDetails struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of CPUs in the instancetype.
	Cpu *int64 `json:"Cpu,omitempty"`
	// Ephemeral disk capacity to be provided with units example - 10Gi.
	DiskSize *int64 `json:"DiskSize,omitempty"`
	// Virtual machine memory defined in mebibytes (MiB).
	Memory               *int64 `json:"Memory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesInstanceTypeDetails KubernetesInstanceTypeDetails

// NewKubernetesInstanceTypeDetails instantiates a new KubernetesInstanceTypeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInstanceTypeDetails(classId string, objectType string) *KubernetesInstanceTypeDetails {
	this := KubernetesInstanceTypeDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesInstanceTypeDetailsWithDefaults instantiates a new KubernetesInstanceTypeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInstanceTypeDetailsWithDefaults() *KubernetesInstanceTypeDetails {
	this := KubernetesInstanceTypeDetails{}
	var classId string = "kubernetes.InstanceTypeDetails"
	this.ClassId = classId
	var objectType string = "kubernetes.InstanceTypeDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesInstanceTypeDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesInstanceTypeDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesInstanceTypeDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesInstanceTypeDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesInstanceTypeDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesInstanceTypeDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *KubernetesInstanceTypeDetails) GetCpu() int64 {
	if o == nil || o.Cpu == nil {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInstanceTypeDetails) GetCpuOk() (*int64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *KubernetesInstanceTypeDetails) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *KubernetesInstanceTypeDetails) SetCpu(v int64) {
	o.Cpu = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *KubernetesInstanceTypeDetails) GetDiskSize() int64 {
	if o == nil || o.DiskSize == nil {
		var ret int64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInstanceTypeDetails) GetDiskSizeOk() (*int64, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *KubernetesInstanceTypeDetails) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int64 and assigns it to the DiskSize field.
func (o *KubernetesInstanceTypeDetails) SetDiskSize(v int64) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *KubernetesInstanceTypeDetails) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInstanceTypeDetails) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *KubernetesInstanceTypeDetails) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *KubernetesInstanceTypeDetails) SetMemory(v int64) {
	o.Memory = &v
}

func (o KubernetesInstanceTypeDetails) MarshalJSON() ([]byte, error) {
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
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.DiskSize != nil {
		toSerialize["DiskSize"] = o.DiskSize
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesInstanceTypeDetails) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesInstanceTypeDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of CPUs in the instancetype.
		Cpu *int64 `json:"Cpu,omitempty"`
		// Ephemeral disk capacity to be provided with units example - 10Gi.
		DiskSize *int64 `json:"DiskSize,omitempty"`
		// Virtual machine memory defined in mebibytes (MiB).
		Memory *int64 `json:"Memory,omitempty"`
	}

	varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct := KubernetesInstanceTypeDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesInstanceTypeDetails := _KubernetesInstanceTypeDetails{}
		varKubernetesInstanceTypeDetails.ClassId = varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct.ClassId
		varKubernetesInstanceTypeDetails.ObjectType = varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct.ObjectType
		varKubernetesInstanceTypeDetails.Cpu = varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct.Cpu
		varKubernetesInstanceTypeDetails.DiskSize = varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct.DiskSize
		varKubernetesInstanceTypeDetails.Memory = varKubernetesInstanceTypeDetailsWithoutEmbeddedStruct.Memory
		*o = KubernetesInstanceTypeDetails(varKubernetesInstanceTypeDetails)
	} else {
		return err
	}

	varKubernetesInstanceTypeDetails := _KubernetesInstanceTypeDetails{}

	err = json.Unmarshal(bytes, &varKubernetesInstanceTypeDetails)
	if err == nil {
		o.MoBaseComplexType = varKubernetesInstanceTypeDetails.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cpu")
		delete(additionalProperties, "DiskSize")
		delete(additionalProperties, "Memory")

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

type NullableKubernetesInstanceTypeDetails struct {
	value *KubernetesInstanceTypeDetails
	isSet bool
}

func (v NullableKubernetesInstanceTypeDetails) Get() *KubernetesInstanceTypeDetails {
	return v.value
}

func (v *NullableKubernetesInstanceTypeDetails) Set(val *KubernetesInstanceTypeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInstanceTypeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInstanceTypeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInstanceTypeDetails(val *KubernetesInstanceTypeDetails) *NullableKubernetesInstanceTypeDetails {
	return &NullableKubernetesInstanceTypeDetails{value: val, isSet: true}
}

func (v NullableKubernetesInstanceTypeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInstanceTypeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
