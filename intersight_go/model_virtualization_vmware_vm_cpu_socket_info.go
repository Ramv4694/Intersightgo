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

// VirtualizationVmwareVmCpuSocketInfo Information about the virtual machine's hardware platform (CPU, memory).
type VirtualizationVmwareVmCpuSocketInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of core per CPU socket (value may be empty).
	CoresPerSocket *int64 `json:"CoresPerSocket,omitempty"`
	// Number of CPUs allocated to this VM.
	NumCpus *int64 `json:"NumCpus,omitempty"`
	// The number of CPU sockets allocated.
	NumSockets           *int64 `json:"NumSockets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmCpuSocketInfo VirtualizationVmwareVmCpuSocketInfo

// NewVirtualizationVmwareVmCpuSocketInfo instantiates a new VirtualizationVmwareVmCpuSocketInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmCpuSocketInfo(classId string, objectType string) *VirtualizationVmwareVmCpuSocketInfo {
	this := VirtualizationVmwareVmCpuSocketInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmCpuSocketInfoWithDefaults instantiates a new VirtualizationVmwareVmCpuSocketInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmCpuSocketInfoWithDefaults() *VirtualizationVmwareVmCpuSocketInfo {
	this := VirtualizationVmwareVmCpuSocketInfo{}
	var classId string = "virtualization.VmwareVmCpuSocketInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmCpuSocketInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmCpuSocketInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmCpuSocketInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmCpuSocketInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmCpuSocketInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCoresPerSocket returns the CoresPerSocket field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetCoresPerSocket() int64 {
	if o == nil || o.CoresPerSocket == nil {
		var ret int64
		return ret
	}
	return *o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil || o.CoresPerSocket == nil {
		return nil, false
	}
	return o.CoresPerSocket, true
}

// HasCoresPerSocket returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) HasCoresPerSocket() bool {
	if o != nil && o.CoresPerSocket != nil {
		return true
	}

	return false
}

// SetCoresPerSocket gets a reference to the given int64 and assigns it to the CoresPerSocket field.
func (o *VirtualizationVmwareVmCpuSocketInfo) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = &v
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetNumCpus() int64 {
	if o == nil || o.NumCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetNumCpusOk() (*int64, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int64 and assigns it to the NumCpus field.
func (o *VirtualizationVmwareVmCpuSocketInfo) SetNumCpus(v int64) {
	o.NumCpus = &v
}

// GetNumSockets returns the NumSockets field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetNumSockets() int64 {
	if o == nil || o.NumSockets == nil {
		var ret int64
		return ret
	}
	return *o.NumSockets
}

// GetNumSocketsOk returns a tuple with the NumSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) GetNumSocketsOk() (*int64, bool) {
	if o == nil || o.NumSockets == nil {
		return nil, false
	}
	return o.NumSockets, true
}

// HasNumSockets returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuSocketInfo) HasNumSockets() bool {
	if o != nil && o.NumSockets != nil {
		return true
	}

	return false
}

// SetNumSockets gets a reference to the given int64 and assigns it to the NumSockets field.
func (o *VirtualizationVmwareVmCpuSocketInfo) SetNumSockets(v int64) {
	o.NumSockets = &v
}

func (o VirtualizationVmwareVmCpuSocketInfo) MarshalJSON() ([]byte, error) {
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
	if o.CoresPerSocket != nil {
		toSerialize["CoresPerSocket"] = o.CoresPerSocket
	}
	if o.NumCpus != nil {
		toSerialize["NumCpus"] = o.NumCpus
	}
	if o.NumSockets != nil {
		toSerialize["NumSockets"] = o.NumSockets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmCpuSocketInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of core per CPU socket (value may be empty).
		CoresPerSocket *int64 `json:"CoresPerSocket,omitempty"`
		// Number of CPUs allocated to this VM.
		NumCpus *int64 `json:"NumCpus,omitempty"`
		// The number of CPU sockets allocated.
		NumSockets *int64 `json:"NumSockets,omitempty"`
	}

	varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct := VirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVmCpuSocketInfo := _VirtualizationVmwareVmCpuSocketInfo{}
		varVirtualizationVmwareVmCpuSocketInfo.ClassId = varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVmCpuSocketInfo.ObjectType = varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVmCpuSocketInfo.CoresPerSocket = varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct.CoresPerSocket
		varVirtualizationVmwareVmCpuSocketInfo.NumCpus = varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct.NumCpus
		varVirtualizationVmwareVmCpuSocketInfo.NumSockets = varVirtualizationVmwareVmCpuSocketInfoWithoutEmbeddedStruct.NumSockets
		*o = VirtualizationVmwareVmCpuSocketInfo(varVirtualizationVmwareVmCpuSocketInfo)
	} else {
		return err
	}

	varVirtualizationVmwareVmCpuSocketInfo := _VirtualizationVmwareVmCpuSocketInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmCpuSocketInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareVmCpuSocketInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CoresPerSocket")
		delete(additionalProperties, "NumCpus")
		delete(additionalProperties, "NumSockets")

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

type NullableVirtualizationVmwareVmCpuSocketInfo struct {
	value *VirtualizationVmwareVmCpuSocketInfo
	isSet bool
}

func (v NullableVirtualizationVmwareVmCpuSocketInfo) Get() *VirtualizationVmwareVmCpuSocketInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfo) Set(val *VirtualizationVmwareVmCpuSocketInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmCpuSocketInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmCpuSocketInfo(val *VirtualizationVmwareVmCpuSocketInfo) *NullableVirtualizationVmwareVmCpuSocketInfo {
	return &NullableVirtualizationVmwareVmCpuSocketInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmCpuSocketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmCpuSocketInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
