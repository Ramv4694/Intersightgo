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

// NiatelemetryInterface Stores data about interfaces per switch.
type NiatelemetryInterface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of number of interafces down.
	InterfaceDownCount *int64 `json:"InterfaceDownCount,omitempty"`
	// Return value of number of interafces up.
	InterfaceUpCount     *int64 `json:"InterfaceUpCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryInterface NiatelemetryInterface

// NewNiatelemetryInterface instantiates a new NiatelemetryInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryInterface(classId string, objectType string) *NiatelemetryInterface {
	this := NiatelemetryInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryInterfaceWithDefaults instantiates a new NiatelemetryInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryInterfaceWithDefaults() *NiatelemetryInterface {
	this := NiatelemetryInterface{}
	var classId string = "niatelemetry.Interface"
	this.ClassId = classId
	var objectType string = "niatelemetry.Interface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceDownCount returns the InterfaceDownCount field value if set, zero value otherwise.
func (o *NiatelemetryInterface) GetInterfaceDownCount() int64 {
	if o == nil || o.InterfaceDownCount == nil {
		var ret int64
		return ret
	}
	return *o.InterfaceDownCount
}

// GetInterfaceDownCountOk returns a tuple with the InterfaceDownCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterface) GetInterfaceDownCountOk() (*int64, bool) {
	if o == nil || o.InterfaceDownCount == nil {
		return nil, false
	}
	return o.InterfaceDownCount, true
}

// HasInterfaceDownCount returns a boolean if a field has been set.
func (o *NiatelemetryInterface) HasInterfaceDownCount() bool {
	if o != nil && o.InterfaceDownCount != nil {
		return true
	}

	return false
}

// SetInterfaceDownCount gets a reference to the given int64 and assigns it to the InterfaceDownCount field.
func (o *NiatelemetryInterface) SetInterfaceDownCount(v int64) {
	o.InterfaceDownCount = &v
}

// GetInterfaceUpCount returns the InterfaceUpCount field value if set, zero value otherwise.
func (o *NiatelemetryInterface) GetInterfaceUpCount() int64 {
	if o == nil || o.InterfaceUpCount == nil {
		var ret int64
		return ret
	}
	return *o.InterfaceUpCount
}

// GetInterfaceUpCountOk returns a tuple with the InterfaceUpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInterface) GetInterfaceUpCountOk() (*int64, bool) {
	if o == nil || o.InterfaceUpCount == nil {
		return nil, false
	}
	return o.InterfaceUpCount, true
}

// HasInterfaceUpCount returns a boolean if a field has been set.
func (o *NiatelemetryInterface) HasInterfaceUpCount() bool {
	if o != nil && o.InterfaceUpCount != nil {
		return true
	}

	return false
}

// SetInterfaceUpCount gets a reference to the given int64 and assigns it to the InterfaceUpCount field.
func (o *NiatelemetryInterface) SetInterfaceUpCount(v int64) {
	o.InterfaceUpCount = &v
}

func (o NiatelemetryInterface) MarshalJSON() ([]byte, error) {
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
	if o.InterfaceDownCount != nil {
		toSerialize["InterfaceDownCount"] = o.InterfaceDownCount
	}
	if o.InterfaceUpCount != nil {
		toSerialize["InterfaceUpCount"] = o.InterfaceUpCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryInterface) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return value of number of interafces down.
		InterfaceDownCount *int64 `json:"InterfaceDownCount,omitempty"`
		// Return value of number of interafces up.
		InterfaceUpCount *int64 `json:"InterfaceUpCount,omitempty"`
	}

	varNiatelemetryInterfaceWithoutEmbeddedStruct := NiatelemetryInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryInterface := _NiatelemetryInterface{}
		varNiatelemetryInterface.ClassId = varNiatelemetryInterfaceWithoutEmbeddedStruct.ClassId
		varNiatelemetryInterface.ObjectType = varNiatelemetryInterfaceWithoutEmbeddedStruct.ObjectType
		varNiatelemetryInterface.InterfaceDownCount = varNiatelemetryInterfaceWithoutEmbeddedStruct.InterfaceDownCount
		varNiatelemetryInterface.InterfaceUpCount = varNiatelemetryInterfaceWithoutEmbeddedStruct.InterfaceUpCount
		*o = NiatelemetryInterface(varNiatelemetryInterface)
	} else {
		return err
	}

	varNiatelemetryInterface := _NiatelemetryInterface{}

	err = json.Unmarshal(bytes, &varNiatelemetryInterface)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryInterface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceDownCount")
		delete(additionalProperties, "InterfaceUpCount")

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

type NullableNiatelemetryInterface struct {
	value *NiatelemetryInterface
	isSet bool
}

func (v NullableNiatelemetryInterface) Get() *NiatelemetryInterface {
	return v.value
}

func (v *NullableNiatelemetryInterface) Set(val *NiatelemetryInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryInterface(val *NiatelemetryInterface) *NullableNiatelemetryInterface {
	return &NullableNiatelemetryInterface{value: val, isSet: true}
}

func (v NullableNiatelemetryInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
