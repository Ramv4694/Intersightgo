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

// StorageNetAppClusterEvent An event where the impacted resource type is a cluster.
type StorageNetAppClusterEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                            `json:"ObjectType"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterEvent StorageNetAppClusterEvent

// NewStorageNetAppClusterEvent instantiates a new StorageNetAppClusterEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterEvent(classId string, objectType string) *StorageNetAppClusterEvent {
	this := StorageNetAppClusterEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterEventWithDefaults instantiates a new StorageNetAppClusterEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterEventWithDefaults() *StorageNetAppClusterEvent {
	this := StorageNetAppClusterEvent{}
	var classId string = "storage.NetAppClusterEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppClusterEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppClusterEvent) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEvent) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppClusterEvent) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppClusterEvent) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppClusterEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppClusterEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppClusterEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                            `json:"ObjectType"`
		Array      *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppClusterEventWithoutEmbeddedStruct := StorageNetAppClusterEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppClusterEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppClusterEvent := _StorageNetAppClusterEvent{}
		varStorageNetAppClusterEvent.ClassId = varStorageNetAppClusterEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppClusterEvent.ObjectType = varStorageNetAppClusterEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppClusterEvent.Array = varStorageNetAppClusterEventWithoutEmbeddedStruct.Array
		*o = StorageNetAppClusterEvent(varStorageNetAppClusterEvent)
	} else {
		return err
	}

	varStorageNetAppClusterEvent := _StorageNetAppClusterEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppClusterEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppClusterEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppClusterEvent struct {
	value *StorageNetAppClusterEvent
	isSet bool
}

func (v NullableStorageNetAppClusterEvent) Get() *StorageNetAppClusterEvent {
	return v.value
}

func (v *NullableStorageNetAppClusterEvent) Set(val *StorageNetAppClusterEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterEvent(val *StorageNetAppClusterEvent) *NullableStorageNetAppClusterEvent {
	return &NullableStorageNetAppClusterEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
