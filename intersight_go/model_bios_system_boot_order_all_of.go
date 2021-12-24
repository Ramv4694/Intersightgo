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
)

// BiosSystemBootOrderAllOf Definition of the list of properties defined in 'bios.SystemBootOrder', excluding properties defined in parent classes.
type BiosSystemBootOrderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme. * `Legacy` - Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader. * `Uefi` - UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.
	BootMode *string `json:"BootMode,omitempty"`
	// The Distinguished Name for this object, used to uniquely identify this object.
	Dn *string `json:"Dn,omitempty"`
	// Secure boot if set to enabled, enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM). * `NotAvailable` - Set the state of Secure Boot to Not Available. * `Disabled` - Set the state of Secure Boot to Disabled. * `Enabled` - Set the state of Secure Boot to Enabled.
	SecureBoot *string               `json:"SecureBoot,omitempty"`
	BiosUnit   *BiosUnitRelationship `json:"BiosUnit,omitempty"`
	// An array of relationships to biosBootDevice resources.
	BootDevices          []BiosBootDeviceRelationship         `json:"BootDevices,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosSystemBootOrderAllOf BiosSystemBootOrderAllOf

// NewBiosSystemBootOrderAllOf instantiates a new BiosSystemBootOrderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosSystemBootOrderAllOf(classId string, objectType string) *BiosSystemBootOrderAllOf {
	this := BiosSystemBootOrderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var bootMode string = "Legacy"
	this.BootMode = &bootMode
	var secureBoot string = "NotAvailable"
	this.SecureBoot = &secureBoot
	return &this
}

// NewBiosSystemBootOrderAllOfWithDefaults instantiates a new BiosSystemBootOrderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosSystemBootOrderAllOfWithDefaults() *BiosSystemBootOrderAllOf {
	this := BiosSystemBootOrderAllOf{}
	var classId string = "bios.SystemBootOrder"
	this.ClassId = classId
	var objectType string = "bios.SystemBootOrder"
	this.ObjectType = objectType
	var bootMode string = "Legacy"
	this.BootMode = &bootMode
	var secureBoot string = "NotAvailable"
	this.SecureBoot = &secureBoot
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosSystemBootOrderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosSystemBootOrderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosSystemBootOrderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosSystemBootOrderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootMode returns the BootMode field value if set, zero value otherwise.
func (o *BiosSystemBootOrderAllOf) GetBootMode() string {
	if o == nil || o.BootMode == nil {
		var ret string
		return ret
	}
	return *o.BootMode
}

// GetBootModeOk returns a tuple with the BootMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetBootModeOk() (*string, bool) {
	if o == nil || o.BootMode == nil {
		return nil, false
	}
	return o.BootMode, true
}

// HasBootMode returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasBootMode() bool {
	if o != nil && o.BootMode != nil {
		return true
	}

	return false
}

// SetBootMode gets a reference to the given string and assigns it to the BootMode field.
func (o *BiosSystemBootOrderAllOf) SetBootMode(v string) {
	o.BootMode = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *BiosSystemBootOrderAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *BiosSystemBootOrderAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetSecureBoot returns the SecureBoot field value if set, zero value otherwise.
func (o *BiosSystemBootOrderAllOf) GetSecureBoot() string {
	if o == nil || o.SecureBoot == nil {
		var ret string
		return ret
	}
	return *o.SecureBoot
}

// GetSecureBootOk returns a tuple with the SecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetSecureBootOk() (*string, bool) {
	if o == nil || o.SecureBoot == nil {
		return nil, false
	}
	return o.SecureBoot, true
}

// HasSecureBoot returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasSecureBoot() bool {
	if o != nil && o.SecureBoot != nil {
		return true
	}

	return false
}

// SetSecureBoot gets a reference to the given string and assigns it to the SecureBoot field.
func (o *BiosSystemBootOrderAllOf) SetSecureBoot(v string) {
	o.SecureBoot = &v
}

// GetBiosUnit returns the BiosUnit field value if set, zero value otherwise.
func (o *BiosSystemBootOrderAllOf) GetBiosUnit() BiosUnitRelationship {
	if o == nil || o.BiosUnit == nil {
		var ret BiosUnitRelationship
		return ret
	}
	return *o.BiosUnit
}

// GetBiosUnitOk returns a tuple with the BiosUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetBiosUnitOk() (*BiosUnitRelationship, bool) {
	if o == nil || o.BiosUnit == nil {
		return nil, false
	}
	return o.BiosUnit, true
}

// HasBiosUnit returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasBiosUnit() bool {
	if o != nil && o.BiosUnit != nil {
		return true
	}

	return false
}

// SetBiosUnit gets a reference to the given BiosUnitRelationship and assigns it to the BiosUnit field.
func (o *BiosSystemBootOrderAllOf) SetBiosUnit(v BiosUnitRelationship) {
	o.BiosUnit = &v
}

// GetBootDevices returns the BootDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosSystemBootOrderAllOf) GetBootDevices() []BiosBootDeviceRelationship {
	if o == nil {
		var ret []BiosBootDeviceRelationship
		return ret
	}
	return o.BootDevices
}

// GetBootDevicesOk returns a tuple with the BootDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosSystemBootOrderAllOf) GetBootDevicesOk() (*[]BiosBootDeviceRelationship, bool) {
	if o == nil || o.BootDevices == nil {
		return nil, false
	}
	return &o.BootDevices, true
}

// HasBootDevices returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasBootDevices() bool {
	if o != nil && o.BootDevices != nil {
		return true
	}

	return false
}

// SetBootDevices gets a reference to the given []BiosBootDeviceRelationship and assigns it to the BootDevices field.
func (o *BiosSystemBootOrderAllOf) SetBootDevices(v []BiosBootDeviceRelationship) {
	o.BootDevices = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosSystemBootOrderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosSystemBootOrderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosSystemBootOrderAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosSystemBootOrderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosSystemBootOrderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootMode != nil {
		toSerialize["BootMode"] = o.BootMode
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.SecureBoot != nil {
		toSerialize["SecureBoot"] = o.SecureBoot
	}
	if o.BiosUnit != nil {
		toSerialize["BiosUnit"] = o.BiosUnit
	}
	if o.BootDevices != nil {
		toSerialize["BootDevices"] = o.BootDevices
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosSystemBootOrderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBiosSystemBootOrderAllOf := _BiosSystemBootOrderAllOf{}

	if err = json.Unmarshal(bytes, &varBiosSystemBootOrderAllOf); err == nil {
		*o = BiosSystemBootOrderAllOf(varBiosSystemBootOrderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootMode")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "SecureBoot")
		delete(additionalProperties, "BiosUnit")
		delete(additionalProperties, "BootDevices")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBiosSystemBootOrderAllOf struct {
	value *BiosSystemBootOrderAllOf
	isSet bool
}

func (v NullableBiosSystemBootOrderAllOf) Get() *BiosSystemBootOrderAllOf {
	return v.value
}

func (v *NullableBiosSystemBootOrderAllOf) Set(val *BiosSystemBootOrderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosSystemBootOrderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosSystemBootOrderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosSystemBootOrderAllOf(val *BiosSystemBootOrderAllOf) *NullableBiosSystemBootOrderAllOf {
	return &NullableBiosSystemBootOrderAllOf{value: val, isSet: true}
}

func (v NullableBiosSystemBootOrderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosSystemBootOrderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
