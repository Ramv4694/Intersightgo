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

// StorageBaseSnapshotSchedule Configuration parameters for snapshot creation at source array.
type StorageBaseSnapshotSchedule struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Snapshot frequency. It is an interval at which snapshot is set to trigger on source array. Examples:     PT2H Snapshot is generated every 2 hours.     P4D, Snapshot is scheduled for every 4 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	Frequency *string `json:"Frequency,omitempty"`
	// Name of the snapshot schedule.
	Name *string `json:"Name,omitempty"`
	// Duration to keep the snapshots on the source array. Once this period expires, system deletes the snapshot automatically from the source array. Examples: P200D,  200 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	RetentionTime *string `json:"RetentionTime,omitempty"`
	// Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more. Format: hh:mm:ss Example: 08:30:00, Snapshot is set for 08:30 AM.
	SnapshotTime         *string `json:"SnapshotTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseSnapshotSchedule StorageBaseSnapshotSchedule

// NewStorageBaseSnapshotSchedule instantiates a new StorageBaseSnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseSnapshotSchedule(classId string, objectType string) *StorageBaseSnapshotSchedule {
	this := StorageBaseSnapshotSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseSnapshotScheduleWithDefaults instantiates a new StorageBaseSnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseSnapshotScheduleWithDefaults() *StorageBaseSnapshotSchedule {
	this := StorageBaseSnapshotSchedule{}
	var classId string = "storage.PureSnapshotSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureSnapshotSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseSnapshotSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseSnapshotSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseSnapshotSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseSnapshotSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *StorageBaseSnapshotSchedule) GetFrequency() string {
	if o == nil || o.Frequency == nil {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetFrequencyOk() (*string, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *StorageBaseSnapshotSchedule) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *StorageBaseSnapshotSchedule) SetFrequency(v string) {
	o.Frequency = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseSnapshotSchedule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseSnapshotSchedule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseSnapshotSchedule) SetName(v string) {
	o.Name = &v
}

// GetRetentionTime returns the RetentionTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshotSchedule) GetRetentionTime() string {
	if o == nil || o.RetentionTime == nil {
		var ret string
		return ret
	}
	return *o.RetentionTime
}

// GetRetentionTimeOk returns a tuple with the RetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetRetentionTimeOk() (*string, bool) {
	if o == nil || o.RetentionTime == nil {
		return nil, false
	}
	return o.RetentionTime, true
}

// HasRetentionTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshotSchedule) HasRetentionTime() bool {
	if o != nil && o.RetentionTime != nil {
		return true
	}

	return false
}

// SetRetentionTime gets a reference to the given string and assigns it to the RetentionTime field.
func (o *StorageBaseSnapshotSchedule) SetRetentionTime(v string) {
	o.RetentionTime = &v
}

// GetSnapshotTime returns the SnapshotTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshotSchedule) GetSnapshotTime() string {
	if o == nil || o.SnapshotTime == nil {
		var ret string
		return ret
	}
	return *o.SnapshotTime
}

// GetSnapshotTimeOk returns a tuple with the SnapshotTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotSchedule) GetSnapshotTimeOk() (*string, bool) {
	if o == nil || o.SnapshotTime == nil {
		return nil, false
	}
	return o.SnapshotTime, true
}

// HasSnapshotTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshotSchedule) HasSnapshotTime() bool {
	if o != nil && o.SnapshotTime != nil {
		return true
	}

	return false
}

// SetSnapshotTime gets a reference to the given string and assigns it to the SnapshotTime field.
func (o *StorageBaseSnapshotSchedule) SetSnapshotTime(v string) {
	o.SnapshotTime = &v
}

func (o StorageBaseSnapshotSchedule) MarshalJSON() ([]byte, error) {
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
	if o.Frequency != nil {
		toSerialize["Frequency"] = o.Frequency
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RetentionTime != nil {
		toSerialize["RetentionTime"] = o.RetentionTime
	}
	if o.SnapshotTime != nil {
		toSerialize["SnapshotTime"] = o.SnapshotTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseSnapshotSchedule) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseSnapshotScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Snapshot frequency. It is an interval at which snapshot is set to trigger on source array. Examples:     PT2H Snapshot is generated every 2 hours.     P4D, Snapshot is scheduled for every 4 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
		Frequency *string `json:"Frequency,omitempty"`
		// Name of the snapshot schedule.
		Name *string `json:"Name,omitempty"`
		// Duration to keep the snapshots on the source array. Once this period expires, system deletes the snapshot automatically from the source array. Examples: P200D,  200 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
		RetentionTime *string `json:"RetentionTime,omitempty"`
		// Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more. Format: hh:mm:ss Example: 08:30:00, Snapshot is set for 08:30 AM.
		SnapshotTime *string `json:"SnapshotTime,omitempty"`
	}

	varStorageBaseSnapshotScheduleWithoutEmbeddedStruct := StorageBaseSnapshotScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseSnapshotScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseSnapshotSchedule := _StorageBaseSnapshotSchedule{}
		varStorageBaseSnapshotSchedule.ClassId = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.ClassId
		varStorageBaseSnapshotSchedule.ObjectType = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.ObjectType
		varStorageBaseSnapshotSchedule.Frequency = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.Frequency
		varStorageBaseSnapshotSchedule.Name = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.Name
		varStorageBaseSnapshotSchedule.RetentionTime = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.RetentionTime
		varStorageBaseSnapshotSchedule.SnapshotTime = varStorageBaseSnapshotScheduleWithoutEmbeddedStruct.SnapshotTime
		*o = StorageBaseSnapshotSchedule(varStorageBaseSnapshotSchedule)
	} else {
		return err
	}

	varStorageBaseSnapshotSchedule := _StorageBaseSnapshotSchedule{}

	err = json.Unmarshal(bytes, &varStorageBaseSnapshotSchedule)
	if err == nil {
		o.MoBaseMo = varStorageBaseSnapshotSchedule.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Frequency")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RetentionTime")
		delete(additionalProperties, "SnapshotTime")

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

type NullableStorageBaseSnapshotSchedule struct {
	value *StorageBaseSnapshotSchedule
	isSet bool
}

func (v NullableStorageBaseSnapshotSchedule) Get() *StorageBaseSnapshotSchedule {
	return v.value
}

func (v *NullableStorageBaseSnapshotSchedule) Set(val *StorageBaseSnapshotSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseSnapshotSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseSnapshotSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseSnapshotSchedule(val *StorageBaseSnapshotSchedule) *NullableStorageBaseSnapshotSchedule {
	return &NullableStorageBaseSnapshotSchedule{value: val, isSet: true}
}

func (v NullableStorageBaseSnapshotSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseSnapshotSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
