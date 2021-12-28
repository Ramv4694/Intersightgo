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

// VirtualizationBondState Bond or teaming information about an interface.
type VirtualizationBondState struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The current active slave. For active-active mode, this field is empty.
	ActiveSlave *string `json:"ActiveSlave,omitempty"`
	// Bond mode, such as \"active-backup\", \"balance-slb\", \"balance-tcp\".
	Mode                 *string  `json:"Mode,omitempty"`
	Slaves               []string `json:"Slaves,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBondState VirtualizationBondState

// NewVirtualizationBondState instantiates a new VirtualizationBondState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBondState(classId string, objectType string) *VirtualizationBondState {
	this := VirtualizationBondState{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBondStateWithDefaults instantiates a new VirtualizationBondState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBondStateWithDefaults() *VirtualizationBondState {
	this := VirtualizationBondState{}
	var classId string = "virtualization.BondState"
	this.ClassId = classId
	var objectType string = "virtualization.BondState"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBondState) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBondState) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBondState) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBondState) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBondState) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBondState) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveSlave returns the ActiveSlave field value if set, zero value otherwise.
func (o *VirtualizationBondState) GetActiveSlave() string {
	if o == nil || o.ActiveSlave == nil {
		var ret string
		return ret
	}
	return *o.ActiveSlave
}

// GetActiveSlaveOk returns a tuple with the ActiveSlave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBondState) GetActiveSlaveOk() (*string, bool) {
	if o == nil || o.ActiveSlave == nil {
		return nil, false
	}
	return o.ActiveSlave, true
}

// HasActiveSlave returns a boolean if a field has been set.
func (o *VirtualizationBondState) HasActiveSlave() bool {
	if o != nil && o.ActiveSlave != nil {
		return true
	}

	return false
}

// SetActiveSlave gets a reference to the given string and assigns it to the ActiveSlave field.
func (o *VirtualizationBondState) SetActiveSlave(v string) {
	o.ActiveSlave = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VirtualizationBondState) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBondState) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VirtualizationBondState) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VirtualizationBondState) SetMode(v string) {
	o.Mode = &v
}

// GetSlaves returns the Slaves field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBondState) GetSlaves() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Slaves
}

// GetSlavesOk returns a tuple with the Slaves field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBondState) GetSlavesOk() (*[]string, bool) {
	if o == nil || o.Slaves == nil {
		return nil, false
	}
	return &o.Slaves, true
}

// HasSlaves returns a boolean if a field has been set.
func (o *VirtualizationBondState) HasSlaves() bool {
	if o != nil && o.Slaves != nil {
		return true
	}

	return false
}

// SetSlaves gets a reference to the given []string and assigns it to the Slaves field.
func (o *VirtualizationBondState) SetSlaves(v []string) {
	o.Slaves = v
}

func (o VirtualizationBondState) MarshalJSON() ([]byte, error) {
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
	if o.ActiveSlave != nil {
		toSerialize["ActiveSlave"] = o.ActiveSlave
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Slaves != nil {
		toSerialize["Slaves"] = o.Slaves
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBondState) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBondStateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The current active slave. For active-active mode, this field is empty.
		ActiveSlave *string `json:"ActiveSlave,omitempty"`
		// Bond mode, such as \"active-backup\", \"balance-slb\", \"balance-tcp\".
		Mode   *string  `json:"Mode,omitempty"`
		Slaves []string `json:"Slaves,omitempty"`
	}

	varVirtualizationBondStateWithoutEmbeddedStruct := VirtualizationBondStateWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBondStateWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBondState := _VirtualizationBondState{}
		varVirtualizationBondState.ClassId = varVirtualizationBondStateWithoutEmbeddedStruct.ClassId
		varVirtualizationBondState.ObjectType = varVirtualizationBondStateWithoutEmbeddedStruct.ObjectType
		varVirtualizationBondState.ActiveSlave = varVirtualizationBondStateWithoutEmbeddedStruct.ActiveSlave
		varVirtualizationBondState.Mode = varVirtualizationBondStateWithoutEmbeddedStruct.Mode
		varVirtualizationBondState.Slaves = varVirtualizationBondStateWithoutEmbeddedStruct.Slaves
		*o = VirtualizationBondState(varVirtualizationBondState)
	} else {
		return err
	}

	varVirtualizationBondState := _VirtualizationBondState{}

	err = json.Unmarshal(bytes, &varVirtualizationBondState)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationBondState.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveSlave")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Slaves")

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

type NullableVirtualizationBondState struct {
	value *VirtualizationBondState
	isSet bool
}

func (v NullableVirtualizationBondState) Get() *VirtualizationBondState {
	return v.value
}

func (v *NullableVirtualizationBondState) Set(val *VirtualizationBondState) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBondState) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBondState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBondState(val *VirtualizationBondState) *NullableVirtualizationBondState {
	return &NullableVirtualizationBondState{value: val, isSet: true}
}

func (v NullableVirtualizationBondState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBondState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
