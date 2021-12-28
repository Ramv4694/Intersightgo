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
)

// BootUsbAllOf Definition of the list of properties defined in 'boot.Usb', excluding properties defined in parent classes.
type BootUsbAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subtype for the selected device type. * `None` - No sub type for USB boot device. * `usb-cd` - Use of Compact Disk (CD) as sub-type for the USB boot device. * `usb-fdd` - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device. * `usb-hdd` - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device.
	Subtype              *string `json:"Subtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootUsbAllOf BootUsbAllOf

// NewBootUsbAllOf instantiates a new BootUsbAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootUsbAllOf(classId string, objectType string) *BootUsbAllOf {
	this := BootUsbAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// NewBootUsbAllOfWithDefaults instantiates a new BootUsbAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootUsbAllOfWithDefaults() *BootUsbAllOf {
	this := BootUsbAllOf{}
	var classId string = "boot.Usb"
	this.ClassId = classId
	var objectType string = "boot.Usb"
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootUsbAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootUsbAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootUsbAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootUsbAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootUsbAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootUsbAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BootUsbAllOf) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootUsbAllOf) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BootUsbAllOf) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BootUsbAllOf) SetSubtype(v string) {
	o.Subtype = &v
}

func (o BootUsbAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootUsbAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootUsbAllOf := _BootUsbAllOf{}

	if err = json.Unmarshal(bytes, &varBootUsbAllOf); err == nil {
		*o = BootUsbAllOf(varBootUsbAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Subtype")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootUsbAllOf struct {
	value *BootUsbAllOf
	isSet bool
}

func (v NullableBootUsbAllOf) Get() *BootUsbAllOf {
	return v.value
}

func (v *NullableBootUsbAllOf) Set(val *BootUsbAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootUsbAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootUsbAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootUsbAllOf(val *BootUsbAllOf) *NullableBootUsbAllOf {
	return &NullableBootUsbAllOf{value: val, isSet: true}
}

func (v NullableBootUsbAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootUsbAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
