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

// StorageNetAppVolumeEvent An event where the impacted resource type is a volume.
type StorageNetAppVolumeEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	Volume               *StorageNetAppVolumeRelationship `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppVolumeEvent StorageNetAppVolumeEvent

// NewStorageNetAppVolumeEvent instantiates a new StorageNetAppVolumeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolumeEvent(classId string, objectType string) *StorageNetAppVolumeEvent {
	this := StorageNetAppVolumeEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeEventWithDefaults instantiates a new StorageNetAppVolumeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeEventWithDefaults() *StorageNetAppVolumeEvent {
	this := StorageNetAppVolumeEvent{}
	var classId string = "storage.NetAppVolumeEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolumeEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolumeEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolumeEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolumeEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolumeEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageNetAppVolumeEvent) GetVolume() StorageNetAppVolumeRelationship {
	if o == nil || o.Volume == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeEvent) GetVolumeOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageNetAppVolumeEvent) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the Volume field.
func (o *StorageNetAppVolumeEvent) SetVolume(v StorageNetAppVolumeRelationship) {
	o.Volume = &v
}

func (o StorageNetAppVolumeEvent) MarshalJSON() ([]byte, error) {
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
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppVolumeEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppVolumeEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		Volume     *StorageNetAppVolumeRelationship `json:"Volume,omitempty"`
	}

	varStorageNetAppVolumeEventWithoutEmbeddedStruct := StorageNetAppVolumeEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppVolumeEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppVolumeEvent := _StorageNetAppVolumeEvent{}
		varStorageNetAppVolumeEvent.ClassId = varStorageNetAppVolumeEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppVolumeEvent.ObjectType = varStorageNetAppVolumeEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppVolumeEvent.Volume = varStorageNetAppVolumeEventWithoutEmbeddedStruct.Volume
		*o = StorageNetAppVolumeEvent(varStorageNetAppVolumeEvent)
	} else {
		return err
	}

	varStorageNetAppVolumeEvent := _StorageNetAppVolumeEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppVolumeEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppVolumeEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Volume")

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

type NullableStorageNetAppVolumeEvent struct {
	value *StorageNetAppVolumeEvent
	isSet bool
}

func (v NullableStorageNetAppVolumeEvent) Get() *StorageNetAppVolumeEvent {
	return v.value
}

func (v *NullableStorageNetAppVolumeEvent) Set(val *StorageNetAppVolumeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeEvent(val *StorageNetAppVolumeEvent) *NullableStorageNetAppVolumeEvent {
	return &NullableStorageNetAppVolumeEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
