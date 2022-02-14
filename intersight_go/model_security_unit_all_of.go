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

// SecurityUnitAllOf Definition of the list of properties defined in 'security.Unit', excluding properties defined in parent classes.
type SecurityUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operational state of the security unit.
	OperState *string `json:"OperState,omitempty"`
	// Operability state of the security unit.
	Operability *string `json:"Operability,omitempty"`
	// The part number of the security unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// PCIe slot of the security unit in the server.
	PciSlot *string `json:"PciSlot,omitempty"`
	// Power state of the security unit.
	Power *string `json:"Power,omitempty"`
	// Thermal state of the security unit.
	Thermal *string `json:"Thermal,omitempty"`
	// The unique identifier assigned to the security unit within the server.
	UnitId *int64 `json:"UnitId,omitempty"`
	// The vendor identifier of the security unit.
	Vid *string `json:"Vid,omitempty"`
	// The voltage state of the security unit.
	Voltage              *string                              `json:"Voltage,omitempty"`
	ComputeBoard         *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityUnitAllOf SecurityUnitAllOf

// NewSecurityUnitAllOf instantiates a new SecurityUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityUnitAllOf(classId string, objectType string) *SecurityUnitAllOf {
	this := SecurityUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSecurityUnitAllOfWithDefaults instantiates a new SecurityUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityUnitAllOfWithDefaults() *SecurityUnitAllOf {
	this := SecurityUnitAllOf{}
	var classId string = "security.Unit"
	this.ClassId = classId
	var objectType string = "security.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SecurityUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SecurityUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SecurityUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SecurityUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *SecurityUnitAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *SecurityUnitAllOf) SetOperability(v string) {
	o.Operability = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *SecurityUnitAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *SecurityUnitAllOf) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetPower() string {
	if o == nil || o.Power == nil {
		var ret string
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetPowerOk() (*string, bool) {
	if o == nil || o.Power == nil {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasPower() bool {
	if o != nil && o.Power != nil {
		return true
	}

	return false
}

// SetPower gets a reference to the given string and assigns it to the Power field.
func (o *SecurityUnitAllOf) SetPower(v string) {
	o.Power = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *SecurityUnitAllOf) SetThermal(v string) {
	o.Thermal = &v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetUnitId() int64 {
	if o == nil || o.UnitId == nil {
		var ret int64
		return ret
	}
	return *o.UnitId
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetUnitIdOk() (*int64, bool) {
	if o == nil || o.UnitId == nil {
		return nil, false
	}
	return o.UnitId, true
}

// HasUnitId returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasUnitId() bool {
	if o != nil && o.UnitId != nil {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given int64 and assigns it to the UnitId field.
func (o *SecurityUnitAllOf) SetUnitId(v int64) {
	o.UnitId = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *SecurityUnitAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetVoltage() string {
	if o == nil || o.Voltage == nil {
		var ret string
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetVoltageOk() (*string, bool) {
	if o == nil || o.Voltage == nil {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasVoltage() bool {
	if o != nil && o.Voltage != nil {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given string and assigns it to the Voltage field.
func (o *SecurityUnitAllOf) SetVoltage(v string) {
	o.Voltage = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *SecurityUnitAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *SecurityUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *SecurityUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *SecurityUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *SecurityUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o SecurityUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.Power != nil {
		toSerialize["Power"] = o.Power
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.UnitId != nil {
		toSerialize["UnitId"] = o.UnitId
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Voltage != nil {
		toSerialize["Voltage"] = o.Voltage
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityUnitAllOf := _SecurityUnitAllOf{}

	if err = json.Unmarshal(bytes, &varSecurityUnitAllOf); err == nil {
		*o = SecurityUnitAllOf(varSecurityUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "Power")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "UnitId")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Voltage")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityUnitAllOf struct {
	value *SecurityUnitAllOf
	isSet bool
}

func (v NullableSecurityUnitAllOf) Get() *SecurityUnitAllOf {
	return v.value
}

func (v *NullableSecurityUnitAllOf) Set(val *SecurityUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityUnitAllOf(val *SecurityUnitAllOf) *NullableSecurityUnitAllOf {
	return &NullableSecurityUnitAllOf{value: val, isSet: true}
}

func (v NullableSecurityUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}