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
)

// StorageBaseArrayControllerAllOf Definition of the list of properties defined in 'storage.BaseArrayController', excluding properties defined in parent classes.
type StorageBaseArrayControllerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Storage array controller name.
	Name *string `json:"Name,omitempty"`
	// Controller running mode, Primary or Secondary. * `Unknown` - Component operational state is unknown. * `Primary` - Component operates in primary mode and accepts workloads. * `Secondary` - Component is running as a secondary or standby mode. * `Maintenance` - Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload.
	OperationalMode *string `json:"OperationalMode,omitempty"`
	// Status of the storage controller. * `Unknown` - Component status is not available. * `Ok` - Component is healthy and no issues found. * `Degraded` - Functioning, but not at full capability due to a non-fatal failure. * `Critical` - Not functioning or requiring immediate attention. * `Offline` - Component is installed, but powered off. * `Identifying` - Component is in initialization process. * `NotAvailable` - Component is not installed or not available. * `Updating` - Software update is in progress. * `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported.
	Status *string `json:"Status,omitempty"`
	// Software version running on a storage controller.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseArrayControllerAllOf StorageBaseArrayControllerAllOf

// NewStorageBaseArrayControllerAllOf instantiates a new StorageBaseArrayControllerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseArrayControllerAllOf(classId string, objectType string) *StorageBaseArrayControllerAllOf {
	this := StorageBaseArrayControllerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationalMode string = "Unknown"
	this.OperationalMode = &operationalMode
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewStorageBaseArrayControllerAllOfWithDefaults instantiates a new StorageBaseArrayControllerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseArrayControllerAllOfWithDefaults() *StorageBaseArrayControllerAllOf {
	this := StorageBaseArrayControllerAllOf{}
	var operationalMode string = "Unknown"
	this.OperationalMode = &operationalMode
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseArrayControllerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseArrayControllerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseArrayControllerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseArrayControllerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseArrayControllerAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseArrayControllerAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseArrayControllerAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperationalMode returns the OperationalMode field value if set, zero value otherwise.
func (o *StorageBaseArrayControllerAllOf) GetOperationalMode() string {
	if o == nil || o.OperationalMode == nil {
		var ret string
		return ret
	}
	return *o.OperationalMode
}

// GetOperationalModeOk returns a tuple with the OperationalMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetOperationalModeOk() (*string, bool) {
	if o == nil || o.OperationalMode == nil {
		return nil, false
	}
	return o.OperationalMode, true
}

// HasOperationalMode returns a boolean if a field has been set.
func (o *StorageBaseArrayControllerAllOf) HasOperationalMode() bool {
	if o != nil && o.OperationalMode != nil {
		return true
	}

	return false
}

// SetOperationalMode gets a reference to the given string and assigns it to the OperationalMode field.
func (o *StorageBaseArrayControllerAllOf) SetOperationalMode(v string) {
	o.OperationalMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageBaseArrayControllerAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageBaseArrayControllerAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageBaseArrayControllerAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StorageBaseArrayControllerAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseArrayControllerAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StorageBaseArrayControllerAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StorageBaseArrayControllerAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o StorageBaseArrayControllerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperationalMode != nil {
		toSerialize["OperationalMode"] = o.OperationalMode
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseArrayControllerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseArrayControllerAllOf := _StorageBaseArrayControllerAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseArrayControllerAllOf); err == nil {
		*o = StorageBaseArrayControllerAllOf(varStorageBaseArrayControllerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperationalMode")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseArrayControllerAllOf struct {
	value *StorageBaseArrayControllerAllOf
	isSet bool
}

func (v NullableStorageBaseArrayControllerAllOf) Get() *StorageBaseArrayControllerAllOf {
	return v.value
}

func (v *NullableStorageBaseArrayControllerAllOf) Set(val *StorageBaseArrayControllerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseArrayControllerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseArrayControllerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseArrayControllerAllOf(val *StorageBaseArrayControllerAllOf) *NullableStorageBaseArrayControllerAllOf {
	return &NullableStorageBaseArrayControllerAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseArrayControllerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseArrayControllerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
