/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-10-09T21:18:32Z.
 *
 * API version: 1.0.9-4809
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareComponentImpactAllOf Definition of the list of properties defined in 'firmware.ComponentImpact', excluding properties defined in parent classes.
type FirmwareComponentImpactAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Impact on the component after the upgrade. * `ALL` - This represents all the components. * `ALL,HDD` - This represents all the components plus the HDDs. * `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category). * `Storage` - This represents the storage controller components. * `None` - This represents none of the components. * `NXOS` - This represents NXOS components. * `IOM` - This represents IOM components. * `PSU` - This represents PSU components. * `CIMC` - This represents CIMC components. * `BIOS` - This represents BIOS components. * `PCIE` - This represents PCIE components. * `Drive` - This represents Drive components. * `DIMM` - This represents DIMM components. * `BoardController` - This represents Board Controller components. * `StorageController` - This represents Storage Controller components. * `HBA` - This represents HBA components. * `GPU` - This represents GPU components. * `SasExpander` - This represents SasExpander components. * `MSwitch` - This represents mSwitch components. * `CMC` - This represents CMC components.
	Component            *string `json:"Component,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentImpactAllOf FirmwareComponentImpactAllOf

// NewFirmwareComponentImpactAllOf instantiates a new FirmwareComponentImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentImpactAllOf(classId string, objectType string) *FirmwareComponentImpactAllOf {
	this := FirmwareComponentImpactAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var component string = "ALL"
	this.Component = &component
	return &this
}

// NewFirmwareComponentImpactAllOfWithDefaults instantiates a new FirmwareComponentImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentImpactAllOfWithDefaults() *FirmwareComponentImpactAllOf {
	this := FirmwareComponentImpactAllOf{}
	var classId string = "firmware.ComponentImpact"
	this.ClassId = classId
	var objectType string = "firmware.ComponentImpact"
	this.ObjectType = objectType
	var component string = "ALL"
	this.Component = &component
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentImpactAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpactAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentImpactAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentImpactAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpactAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentImpactAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FirmwareComponentImpactAllOf) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentImpactAllOf) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FirmwareComponentImpactAllOf) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FirmwareComponentImpactAllOf) SetComponent(v string) {
	o.Component = &v
}

func (o FirmwareComponentImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Component != nil {
		toSerialize["Component"] = o.Component
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareComponentImpactAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareComponentImpactAllOf := _FirmwareComponentImpactAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareComponentImpactAllOf); err == nil {
		*o = FirmwareComponentImpactAllOf(varFirmwareComponentImpactAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Component")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareComponentImpactAllOf struct {
	value *FirmwareComponentImpactAllOf
	isSet bool
}

func (v NullableFirmwareComponentImpactAllOf) Get() *FirmwareComponentImpactAllOf {
	return v.value
}

func (v *NullableFirmwareComponentImpactAllOf) Set(val *FirmwareComponentImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentImpactAllOf(val *FirmwareComponentImpactAllOf) *NullableFirmwareComponentImpactAllOf {
	return &NullableFirmwareComponentImpactAllOf{value: val, isSet: true}
}

func (v NullableFirmwareComponentImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
