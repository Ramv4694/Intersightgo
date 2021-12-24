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

// ConvergedinfraBasePod Abstract base type for all pods.
type ConvergedinfraBasePod struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Description of the pod. A short note about the nature or purpose of the pod.
	Description *string `json:"Description,omitempty"`
	// Name of the pod. Concrete pod will be created with this name.
	Name *string `json:"Name,omitempty"`
	// Defines the type of the pod. * `FlexPod` - Pod type is FlexPod, an integrated infrastructure solution developed by Cisco and NetApp. * `FlashStack` - Pod type is FlashStack, an integrated infrastructure solution developed by Cisco and Pure Storage.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraBasePod ConvergedinfraBasePod

// NewConvergedinfraBasePod instantiates a new ConvergedinfraBasePod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraBasePod(classId string, objectType string) *ConvergedinfraBasePod {
	this := ConvergedinfraBasePod{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "FlexPod"
	this.Type = &type_
	return &this
}

// NewConvergedinfraBasePodWithDefaults instantiates a new ConvergedinfraBasePod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraBasePodWithDefaults() *ConvergedinfraBasePod {
	this := ConvergedinfraBasePod{}
	var classId string = "convergedinfra.Pod"
	this.ClassId = classId
	var objectType string = "convergedinfra.Pod"
	this.ObjectType = objectType
	var type_ string = "FlexPod"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraBasePod) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePod) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraBasePod) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraBasePod) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePod) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraBasePod) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConvergedinfraBasePod) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePod) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConvergedinfraBasePod) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConvergedinfraBasePod) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConvergedinfraBasePod) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePod) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConvergedinfraBasePod) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConvergedinfraBasePod) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConvergedinfraBasePod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConvergedinfraBasePod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConvergedinfraBasePod) SetType(v string) {
	o.Type = &v
}

func (o ConvergedinfraBasePod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraBasePod) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraBasePodWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Description of the pod. A short note about the nature or purpose of the pod.
		Description *string `json:"Description,omitempty"`
		// Name of the pod. Concrete pod will be created with this name.
		Name *string `json:"Name,omitempty"`
		// Defines the type of the pod. * `FlexPod` - Pod type is FlexPod, an integrated infrastructure solution developed by Cisco and NetApp. * `FlashStack` - Pod type is FlashStack, an integrated infrastructure solution developed by Cisco and Pure Storage.
		Type *string `json:"Type,omitempty"`
	}

	varConvergedinfraBasePodWithoutEmbeddedStruct := ConvergedinfraBasePodWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraBasePodWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraBasePod := _ConvergedinfraBasePod{}
		varConvergedinfraBasePod.ClassId = varConvergedinfraBasePodWithoutEmbeddedStruct.ClassId
		varConvergedinfraBasePod.ObjectType = varConvergedinfraBasePodWithoutEmbeddedStruct.ObjectType
		varConvergedinfraBasePod.Description = varConvergedinfraBasePodWithoutEmbeddedStruct.Description
		varConvergedinfraBasePod.Name = varConvergedinfraBasePodWithoutEmbeddedStruct.Name
		varConvergedinfraBasePod.Type = varConvergedinfraBasePodWithoutEmbeddedStruct.Type
		*o = ConvergedinfraBasePod(varConvergedinfraBasePod)
	} else {
		return err
	}

	varConvergedinfraBasePod := _ConvergedinfraBasePod{}

	err = json.Unmarshal(bytes, &varConvergedinfraBasePod)
	if err == nil {
		o.MoBaseMo = varConvergedinfraBasePod.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableConvergedinfraBasePod struct {
	value *ConvergedinfraBasePod
	isSet bool
}

func (v NullableConvergedinfraBasePod) Get() *ConvergedinfraBasePod {
	return v.value
}

func (v *NullableConvergedinfraBasePod) Set(val *ConvergedinfraBasePod) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraBasePod) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraBasePod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraBasePod(val *ConvergedinfraBasePod) *NullableConvergedinfraBasePod {
	return &NullableConvergedinfraBasePod{value: val, isSet: true}
}

func (v NullableConvergedinfraBasePod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraBasePod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
