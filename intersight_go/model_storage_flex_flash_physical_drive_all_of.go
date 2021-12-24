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

// StorageFlexFlashPhysicalDriveAllOf Definition of the list of properties defined in 'storage.FlexFlashPhysicalDrive', excluding properties defined in parent classes.
type StorageFlexFlashPhysicalDriveAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The status of the flex flash physical drive.
	CardStatus *string `json:"CardStatus,omitempty"`
	// The card type of the flex flash physical drive.
	CardType *string `json:"CardType,omitempty"`
	// The OEM Identifier of the flex flash physical drive.
	OemId *string `json:"OemId,omitempty"`
	// The drive status of the flex flash physical drive.
	PdStatus                   *string                                 `json:"PdStatus,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashPhysicalDriveAllOf StorageFlexFlashPhysicalDriveAllOf

// NewStorageFlexFlashPhysicalDriveAllOf instantiates a new StorageFlexFlashPhysicalDriveAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashPhysicalDriveAllOf(classId string, objectType string) *StorageFlexFlashPhysicalDriveAllOf {
	this := StorageFlexFlashPhysicalDriveAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashPhysicalDriveAllOfWithDefaults instantiates a new StorageFlexFlashPhysicalDriveAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashPhysicalDriveAllOfWithDefaults() *StorageFlexFlashPhysicalDriveAllOf {
	this := StorageFlexFlashPhysicalDriveAllOf{}
	var classId string = "storage.FlexFlashPhysicalDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashPhysicalDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashPhysicalDriveAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashPhysicalDriveAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashPhysicalDriveAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashPhysicalDriveAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCardStatus returns the CardStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardStatus() string {
	if o == nil || o.CardStatus == nil {
		var ret string
		return ret
	}
	return *o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardStatusOk() (*string, bool) {
	if o == nil || o.CardStatus == nil {
		return nil, false
	}
	return o.CardStatus, true
}

// HasCardStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasCardStatus() bool {
	if o != nil && o.CardStatus != nil {
		return true
	}

	return false
}

// SetCardStatus gets a reference to the given string and assigns it to the CardStatus field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetCardStatus(v string) {
	o.CardStatus = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardType() string {
	if o == nil || o.CardType == nil {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardTypeOk() (*string, bool) {
	if o == nil || o.CardType == nil {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasCardType() bool {
	if o != nil && o.CardType != nil {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetCardType(v string) {
	o.CardType = &v
}

// GetOemId returns the OemId field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetOemId() string {
	if o == nil || o.OemId == nil {
		var ret string
		return ret
	}
	return *o.OemId
}

// GetOemIdOk returns a tuple with the OemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetOemIdOk() (*string, bool) {
	if o == nil || o.OemId == nil {
		return nil, false
	}
	return o.OemId, true
}

// HasOemId returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasOemId() bool {
	if o != nil && o.OemId != nil {
		return true
	}

	return false
}

// SetOemId gets a reference to the given string and assigns it to the OemId field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetOemId(v string) {
	o.OemId = &v
}

// GetPdStatus returns the PdStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetPdStatus() string {
	if o == nil || o.PdStatus == nil {
		var ret string
		return ret
	}
	return *o.PdStatus
}

// GetPdStatusOk returns a tuple with the PdStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetPdStatusOk() (*string, bool) {
	if o == nil || o.PdStatus == nil {
		return nil, false
	}
	return o.PdStatus, true
}

// HasPdStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasPdStatus() bool {
	if o != nil && o.PdStatus != nil {
		return true
	}

	return false
}

// SetPdStatus gets a reference to the given string and assigns it to the PdStatus field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetPdStatus(v string) {
	o.PdStatus = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashPhysicalDriveAllOf) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashPhysicalDriveAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

func (o StorageFlexFlashPhysicalDriveAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CardStatus != nil {
		toSerialize["CardStatus"] = o.CardStatus
	}
	if o.CardType != nil {
		toSerialize["CardType"] = o.CardType
	}
	if o.OemId != nil {
		toSerialize["OemId"] = o.OemId
	}
	if o.PdStatus != nil {
		toSerialize["PdStatus"] = o.PdStatus
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashPhysicalDriveAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageFlexFlashPhysicalDriveAllOf := _StorageFlexFlashPhysicalDriveAllOf{}

	if err = json.Unmarshal(bytes, &varStorageFlexFlashPhysicalDriveAllOf); err == nil {
		*o = StorageFlexFlashPhysicalDriveAllOf(varStorageFlexFlashPhysicalDriveAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CardStatus")
		delete(additionalProperties, "CardType")
		delete(additionalProperties, "OemId")
		delete(additionalProperties, "PdStatus")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageFlexFlashPhysicalDriveAllOf struct {
	value *StorageFlexFlashPhysicalDriveAllOf
	isSet bool
}

func (v NullableStorageFlexFlashPhysicalDriveAllOf) Get() *StorageFlexFlashPhysicalDriveAllOf {
	return v.value
}

func (v *NullableStorageFlexFlashPhysicalDriveAllOf) Set(val *StorageFlexFlashPhysicalDriveAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashPhysicalDriveAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashPhysicalDriveAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashPhysicalDriveAllOf(val *StorageFlexFlashPhysicalDriveAllOf) *NullableStorageFlexFlashPhysicalDriveAllOf {
	return &NullableStorageFlexFlashPhysicalDriveAllOf{value: val, isSet: true}
}

func (v NullableStorageFlexFlashPhysicalDriveAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashPhysicalDriveAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
