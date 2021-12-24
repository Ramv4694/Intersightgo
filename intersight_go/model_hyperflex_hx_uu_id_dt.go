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

// HyperflexHxUuIdDt The unique identifier of an entity.
type HyperflexHxUuIdDt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string              `json:"ObjectType"`
	Links      []HyperflexHxLinkDt `json:"Links,omitempty"`
	// The unique identifier string of an entity.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxUuIdDt HyperflexHxUuIdDt

// NewHyperflexHxUuIdDt instantiates a new HyperflexHxUuIdDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxUuIdDt(classId string, objectType string) *HyperflexHxUuIdDt {
	this := HyperflexHxUuIdDt{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxUuIdDtWithDefaults instantiates a new HyperflexHxUuIdDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxUuIdDtWithDefaults() *HyperflexHxUuIdDt {
	this := HyperflexHxUuIdDt{}
	var classId string = "hyperflex.HxUuIdDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxUuIdDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxUuIdDt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDt) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxUuIdDt) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxUuIdDt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDt) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxUuIdDt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxUuIdDt) GetLinks() []HyperflexHxLinkDt {
	if o == nil {
		var ret []HyperflexHxLinkDt
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxUuIdDt) GetLinksOk() (*[]HyperflexHxLinkDt, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HyperflexHxUuIdDt) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []HyperflexHxLinkDt and assigns it to the Links field.
func (o *HyperflexHxUuIdDt) SetLinks(v []HyperflexHxLinkDt) {
	o.Links = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHxUuIdDt) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxUuIdDt) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHxUuIdDt) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHxUuIdDt) SetUuid(v string) {
	o.Uuid = &v
}

func (o HyperflexHxUuIdDt) MarshalJSON() ([]byte, error) {
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
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxUuIdDt) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxUuIdDtWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string              `json:"ObjectType"`
		Links      []HyperflexHxLinkDt `json:"Links,omitempty"`
		// The unique identifier string of an entity.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varHyperflexHxUuIdDtWithoutEmbeddedStruct := HyperflexHxUuIdDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxUuIdDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxUuIdDt := _HyperflexHxUuIdDt{}
		varHyperflexHxUuIdDt.ClassId = varHyperflexHxUuIdDtWithoutEmbeddedStruct.ClassId
		varHyperflexHxUuIdDt.ObjectType = varHyperflexHxUuIdDtWithoutEmbeddedStruct.ObjectType
		varHyperflexHxUuIdDt.Links = varHyperflexHxUuIdDtWithoutEmbeddedStruct.Links
		varHyperflexHxUuIdDt.Uuid = varHyperflexHxUuIdDtWithoutEmbeddedStruct.Uuid
		*o = HyperflexHxUuIdDt(varHyperflexHxUuIdDt)
	} else {
		return err
	}

	varHyperflexHxUuIdDt := _HyperflexHxUuIdDt{}

	err = json.Unmarshal(bytes, &varHyperflexHxUuIdDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxUuIdDt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Links")
		delete(additionalProperties, "Uuid")

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

type NullableHyperflexHxUuIdDt struct {
	value *HyperflexHxUuIdDt
	isSet bool
}

func (v NullableHyperflexHxUuIdDt) Get() *HyperflexHxUuIdDt {
	return v.value
}

func (v *NullableHyperflexHxUuIdDt) Set(val *HyperflexHxUuIdDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxUuIdDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxUuIdDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxUuIdDt(val *HyperflexHxUuIdDt) *NullableHyperflexHxUuIdDt {
	return &NullableHyperflexHxUuIdDt{value: val, isSet: true}
}

func (v NullableHyperflexHxUuIdDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxUuIdDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
