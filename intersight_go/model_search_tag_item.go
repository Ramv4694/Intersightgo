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

// SearchTagItem The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax. See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
type SearchTagItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of times this tag key has been set across all resources.
	Count *int64 `json:"Count,omitempty"`
	// Key of the Tag from all the resources in Intersight.
	Key                  *string  `json:"Key,omitempty"`
	Values               []string `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchTagItem SearchTagItem

// NewSearchTagItem instantiates a new SearchTagItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTagItem(classId string, objectType string) *SearchTagItem {
	this := SearchTagItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSearchTagItemWithDefaults instantiates a new SearchTagItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTagItemWithDefaults() *SearchTagItem {
	this := SearchTagItem{}
	var classId string = "search.TagItem"
	this.ClassId = classId
	var objectType string = "search.TagItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SearchTagItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SearchTagItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SearchTagItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SearchTagItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SearchTagItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SearchTagItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchTagItem) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTagItem) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchTagItem) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SearchTagItem) SetCount(v int64) {
	o.Count = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SearchTagItem) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTagItem) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SearchTagItem) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SearchTagItem) SetKey(v string) {
	o.Key = &v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchTagItem) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchTagItem) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SearchTagItem) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SearchTagItem) SetValues(v []string) {
	o.Values = v
}

func (o SearchTagItem) MarshalJSON() ([]byte, error) {
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
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchTagItem) UnmarshalJSON(bytes []byte) (err error) {
	type SearchTagItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of times this tag key has been set across all resources.
		Count *int64 `json:"Count,omitempty"`
		// Key of the Tag from all the resources in Intersight.
		Key    *string  `json:"Key,omitempty"`
		Values []string `json:"Values,omitempty"`
	}

	varSearchTagItemWithoutEmbeddedStruct := SearchTagItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSearchTagItemWithoutEmbeddedStruct)
	if err == nil {
		varSearchTagItem := _SearchTagItem{}
		varSearchTagItem.ClassId = varSearchTagItemWithoutEmbeddedStruct.ClassId
		varSearchTagItem.ObjectType = varSearchTagItemWithoutEmbeddedStruct.ObjectType
		varSearchTagItem.Count = varSearchTagItemWithoutEmbeddedStruct.Count
		varSearchTagItem.Key = varSearchTagItemWithoutEmbeddedStruct.Key
		varSearchTagItem.Values = varSearchTagItemWithoutEmbeddedStruct.Values
		*o = SearchTagItem(varSearchTagItem)
	} else {
		return err
	}

	varSearchTagItem := _SearchTagItem{}

	err = json.Unmarshal(bytes, &varSearchTagItem)
	if err == nil {
		o.MoBaseMo = varSearchTagItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Values")

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

type NullableSearchTagItem struct {
	value *SearchTagItem
	isSet bool
}

func (v NullableSearchTagItem) Get() *SearchTagItem {
	return v.value
}

func (v *NullableSearchTagItem) Set(val *SearchTagItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTagItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTagItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTagItem(val *SearchTagItem) *NullableSearchTagItem {
	return &NullableSearchTagItem{value: val, isSet: true}
}

func (v NullableSearchTagItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTagItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
