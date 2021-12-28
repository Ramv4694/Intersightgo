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

// StorageFlexUtilPhysicalDrive Storage Flex Util Physical Drive.
type StorageFlexUtilPhysicalDrive struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Block size of the FlexUtil Physical drive.
	BlockSize *string `json:"BlockSize,omitempty"`
	// Capacity of the FlexUtil Physical drive.
	Capacity *string `json:"Capacity,omitempty"`
	// Type of the Physical Drive Controller.
	Controller *string `json:"Controller,omitempty"`
	// The number of drives enabled in the FlexUtil Physical Drive.
	DrivesEnabled *string `json:"DrivesEnabled,omitempty"`
	// Health of the FlexUtil Physical drive.
	Health *string `json:"Health,omitempty"`
	// Manufacturing date of the FlexUtil Physical Drive.
	ManufacturerDate *string `json:"ManufacturerDate,omitempty"`
	// Manufacturer identity of the FlexUtil Physical Drive.
	ManufacturerId *string `json:"ManufacturerId,omitempty"`
	// The OEM Identifier of the FlexUtil physical drive.
	OemId *string `json:"OemId,omitempty"`
	// The number of partitions present on the FlexUtil Physical Drive.
	PartitionCount *string `json:"PartitionCount,omitempty"`
	// Status of the FlexUtil Physical Drive.
	PdStatus *string `json:"PdStatus,omitempty"`
	// The type of physical drive. Example - microSD.
	PhysicalDrive *string `json:"PhysicalDrive,omitempty"`
	// Product name of the FlexUtil Physical Drive.
	ProductName *string `json:"ProductName,omitempty"`
	// Product revision of the FlexUtil Physical Drive.
	ProductRevision *string `json:"ProductRevision,omitempty"`
	// Read error count of the FlexUtil Physical Drive.
	ReadErrorCount *string `json:"ReadErrorCount,omitempty"`
	// Read error threshold for FlexUtil Physical Drive.
	ReadErrorThreshold *string `json:"ReadErrorThreshold,omitempty"`
	// Write access state of the FlexUtil Physical Drive.
	WriteEnabled *string `json:"WriteEnabled,omitempty"`
	// Write error count of the FlexUtil Physical Drive.
	WriteErrorCount *string `json:"WriteErrorCount,omitempty"`
	// Write error threshold for FlexUtil Physical Drive.
	WriteErrorThreshold       *string                                `json:"WriteErrorThreshold,omitempty"`
	InventoryDeviceInfo       *InventoryDeviceInfoRelationship       `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	StorageFlexUtilController *StorageFlexUtilControllerRelationship `json:"StorageFlexUtilController,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _StorageFlexUtilPhysicalDrive StorageFlexUtilPhysicalDrive

// NewStorageFlexUtilPhysicalDrive instantiates a new StorageFlexUtilPhysicalDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexUtilPhysicalDrive(classId string, objectType string) *StorageFlexUtilPhysicalDrive {
	this := StorageFlexUtilPhysicalDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexUtilPhysicalDriveWithDefaults instantiates a new StorageFlexUtilPhysicalDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexUtilPhysicalDriveWithDefaults() *StorageFlexUtilPhysicalDrive {
	this := StorageFlexUtilPhysicalDrive{}
	var classId string = "storage.FlexUtilPhysicalDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexUtilPhysicalDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexUtilPhysicalDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexUtilPhysicalDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexUtilPhysicalDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexUtilPhysicalDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetBlockSize() string {
	if o == nil || o.BlockSize == nil {
		var ret string
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetBlockSizeOk() (*string, bool) {
	if o == nil || o.BlockSize == nil {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasBlockSize() bool {
	if o != nil && o.BlockSize != nil {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given string and assigns it to the BlockSize field.
func (o *StorageFlexUtilPhysicalDrive) SetBlockSize(v string) {
	o.BlockSize = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *StorageFlexUtilPhysicalDrive) SetCapacity(v string) {
	o.Capacity = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetController() string {
	if o == nil || o.Controller == nil {
		var ret string
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetControllerOk() (*string, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given string and assigns it to the Controller field.
func (o *StorageFlexUtilPhysicalDrive) SetController(v string) {
	o.Controller = &v
}

// GetDrivesEnabled returns the DrivesEnabled field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetDrivesEnabled() string {
	if o == nil || o.DrivesEnabled == nil {
		var ret string
		return ret
	}
	return *o.DrivesEnabled
}

// GetDrivesEnabledOk returns a tuple with the DrivesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetDrivesEnabledOk() (*string, bool) {
	if o == nil || o.DrivesEnabled == nil {
		return nil, false
	}
	return o.DrivesEnabled, true
}

// HasDrivesEnabled returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasDrivesEnabled() bool {
	if o != nil && o.DrivesEnabled != nil {
		return true
	}

	return false
}

// SetDrivesEnabled gets a reference to the given string and assigns it to the DrivesEnabled field.
func (o *StorageFlexUtilPhysicalDrive) SetDrivesEnabled(v string) {
	o.DrivesEnabled = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StorageFlexUtilPhysicalDrive) SetHealth(v string) {
	o.Health = &v
}

// GetManufacturerDate returns the ManufacturerDate field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetManufacturerDate() string {
	if o == nil || o.ManufacturerDate == nil {
		var ret string
		return ret
	}
	return *o.ManufacturerDate
}

// GetManufacturerDateOk returns a tuple with the ManufacturerDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetManufacturerDateOk() (*string, bool) {
	if o == nil || o.ManufacturerDate == nil {
		return nil, false
	}
	return o.ManufacturerDate, true
}

// HasManufacturerDate returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasManufacturerDate() bool {
	if o != nil && o.ManufacturerDate != nil {
		return true
	}

	return false
}

// SetManufacturerDate gets a reference to the given string and assigns it to the ManufacturerDate field.
func (o *StorageFlexUtilPhysicalDrive) SetManufacturerDate(v string) {
	o.ManufacturerDate = &v
}

// GetManufacturerId returns the ManufacturerId field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetManufacturerId() string {
	if o == nil || o.ManufacturerId == nil {
		var ret string
		return ret
	}
	return *o.ManufacturerId
}

// GetManufacturerIdOk returns a tuple with the ManufacturerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetManufacturerIdOk() (*string, bool) {
	if o == nil || o.ManufacturerId == nil {
		return nil, false
	}
	return o.ManufacturerId, true
}

// HasManufacturerId returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasManufacturerId() bool {
	if o != nil && o.ManufacturerId != nil {
		return true
	}

	return false
}

// SetManufacturerId gets a reference to the given string and assigns it to the ManufacturerId field.
func (o *StorageFlexUtilPhysicalDrive) SetManufacturerId(v string) {
	o.ManufacturerId = &v
}

// GetOemId returns the OemId field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetOemId() string {
	if o == nil || o.OemId == nil {
		var ret string
		return ret
	}
	return *o.OemId
}

// GetOemIdOk returns a tuple with the OemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetOemIdOk() (*string, bool) {
	if o == nil || o.OemId == nil {
		return nil, false
	}
	return o.OemId, true
}

// HasOemId returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasOemId() bool {
	if o != nil && o.OemId != nil {
		return true
	}

	return false
}

// SetOemId gets a reference to the given string and assigns it to the OemId field.
func (o *StorageFlexUtilPhysicalDrive) SetOemId(v string) {
	o.OemId = &v
}

// GetPartitionCount returns the PartitionCount field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetPartitionCount() string {
	if o == nil || o.PartitionCount == nil {
		var ret string
		return ret
	}
	return *o.PartitionCount
}

// GetPartitionCountOk returns a tuple with the PartitionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetPartitionCountOk() (*string, bool) {
	if o == nil || o.PartitionCount == nil {
		return nil, false
	}
	return o.PartitionCount, true
}

// HasPartitionCount returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasPartitionCount() bool {
	if o != nil && o.PartitionCount != nil {
		return true
	}

	return false
}

// SetPartitionCount gets a reference to the given string and assigns it to the PartitionCount field.
func (o *StorageFlexUtilPhysicalDrive) SetPartitionCount(v string) {
	o.PartitionCount = &v
}

// GetPdStatus returns the PdStatus field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetPdStatus() string {
	if o == nil || o.PdStatus == nil {
		var ret string
		return ret
	}
	return *o.PdStatus
}

// GetPdStatusOk returns a tuple with the PdStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetPdStatusOk() (*string, bool) {
	if o == nil || o.PdStatus == nil {
		return nil, false
	}
	return o.PdStatus, true
}

// HasPdStatus returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasPdStatus() bool {
	if o != nil && o.PdStatus != nil {
		return true
	}

	return false
}

// SetPdStatus gets a reference to the given string and assigns it to the PdStatus field.
func (o *StorageFlexUtilPhysicalDrive) SetPdStatus(v string) {
	o.PdStatus = &v
}

// GetPhysicalDrive returns the PhysicalDrive field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetPhysicalDrive() string {
	if o == nil || o.PhysicalDrive == nil {
		var ret string
		return ret
	}
	return *o.PhysicalDrive
}

// GetPhysicalDriveOk returns a tuple with the PhysicalDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetPhysicalDriveOk() (*string, bool) {
	if o == nil || o.PhysicalDrive == nil {
		return nil, false
	}
	return o.PhysicalDrive, true
}

// HasPhysicalDrive returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasPhysicalDrive() bool {
	if o != nil && o.PhysicalDrive != nil {
		return true
	}

	return false
}

// SetPhysicalDrive gets a reference to the given string and assigns it to the PhysicalDrive field.
func (o *StorageFlexUtilPhysicalDrive) SetPhysicalDrive(v string) {
	o.PhysicalDrive = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *StorageFlexUtilPhysicalDrive) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductRevision returns the ProductRevision field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetProductRevision() string {
	if o == nil || o.ProductRevision == nil {
		var ret string
		return ret
	}
	return *o.ProductRevision
}

// GetProductRevisionOk returns a tuple with the ProductRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetProductRevisionOk() (*string, bool) {
	if o == nil || o.ProductRevision == nil {
		return nil, false
	}
	return o.ProductRevision, true
}

// HasProductRevision returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasProductRevision() bool {
	if o != nil && o.ProductRevision != nil {
		return true
	}

	return false
}

// SetProductRevision gets a reference to the given string and assigns it to the ProductRevision field.
func (o *StorageFlexUtilPhysicalDrive) SetProductRevision(v string) {
	o.ProductRevision = &v
}

// GetReadErrorCount returns the ReadErrorCount field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetReadErrorCount() string {
	if o == nil || o.ReadErrorCount == nil {
		var ret string
		return ret
	}
	return *o.ReadErrorCount
}

// GetReadErrorCountOk returns a tuple with the ReadErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetReadErrorCountOk() (*string, bool) {
	if o == nil || o.ReadErrorCount == nil {
		return nil, false
	}
	return o.ReadErrorCount, true
}

// HasReadErrorCount returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasReadErrorCount() bool {
	if o != nil && o.ReadErrorCount != nil {
		return true
	}

	return false
}

// SetReadErrorCount gets a reference to the given string and assigns it to the ReadErrorCount field.
func (o *StorageFlexUtilPhysicalDrive) SetReadErrorCount(v string) {
	o.ReadErrorCount = &v
}

// GetReadErrorThreshold returns the ReadErrorThreshold field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetReadErrorThreshold() string {
	if o == nil || o.ReadErrorThreshold == nil {
		var ret string
		return ret
	}
	return *o.ReadErrorThreshold
}

// GetReadErrorThresholdOk returns a tuple with the ReadErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetReadErrorThresholdOk() (*string, bool) {
	if o == nil || o.ReadErrorThreshold == nil {
		return nil, false
	}
	return o.ReadErrorThreshold, true
}

// HasReadErrorThreshold returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasReadErrorThreshold() bool {
	if o != nil && o.ReadErrorThreshold != nil {
		return true
	}

	return false
}

// SetReadErrorThreshold gets a reference to the given string and assigns it to the ReadErrorThreshold field.
func (o *StorageFlexUtilPhysicalDrive) SetReadErrorThreshold(v string) {
	o.ReadErrorThreshold = &v
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetWriteEnabled() string {
	if o == nil || o.WriteEnabled == nil {
		var ret string
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetWriteEnabledOk() (*string, bool) {
	if o == nil || o.WriteEnabled == nil {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasWriteEnabled() bool {
	if o != nil && o.WriteEnabled != nil {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given string and assigns it to the WriteEnabled field.
func (o *StorageFlexUtilPhysicalDrive) SetWriteEnabled(v string) {
	o.WriteEnabled = &v
}

// GetWriteErrorCount returns the WriteErrorCount field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorCount() string {
	if o == nil || o.WriteErrorCount == nil {
		var ret string
		return ret
	}
	return *o.WriteErrorCount
}

// GetWriteErrorCountOk returns a tuple with the WriteErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorCountOk() (*string, bool) {
	if o == nil || o.WriteErrorCount == nil {
		return nil, false
	}
	return o.WriteErrorCount, true
}

// HasWriteErrorCount returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasWriteErrorCount() bool {
	if o != nil && o.WriteErrorCount != nil {
		return true
	}

	return false
}

// SetWriteErrorCount gets a reference to the given string and assigns it to the WriteErrorCount field.
func (o *StorageFlexUtilPhysicalDrive) SetWriteErrorCount(v string) {
	o.WriteErrorCount = &v
}

// GetWriteErrorThreshold returns the WriteErrorThreshold field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorThreshold() string {
	if o == nil || o.WriteErrorThreshold == nil {
		var ret string
		return ret
	}
	return *o.WriteErrorThreshold
}

// GetWriteErrorThresholdOk returns a tuple with the WriteErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetWriteErrorThresholdOk() (*string, bool) {
	if o == nil || o.WriteErrorThreshold == nil {
		return nil, false
	}
	return o.WriteErrorThreshold, true
}

// HasWriteErrorThreshold returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasWriteErrorThreshold() bool {
	if o != nil && o.WriteErrorThreshold != nil {
		return true
	}

	return false
}

// SetWriteErrorThreshold gets a reference to the given string and assigns it to the WriteErrorThreshold field.
func (o *StorageFlexUtilPhysicalDrive) SetWriteErrorThreshold(v string) {
	o.WriteErrorThreshold = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexUtilPhysicalDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexUtilPhysicalDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexUtilController returns the StorageFlexUtilController field value if set, zero value otherwise.
func (o *StorageFlexUtilPhysicalDrive) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship {
	if o == nil || o.StorageFlexUtilController == nil {
		var ret StorageFlexUtilControllerRelationship
		return ret
	}
	return *o.StorageFlexUtilController
}

// GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilPhysicalDrive) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool) {
	if o == nil || o.StorageFlexUtilController == nil {
		return nil, false
	}
	return o.StorageFlexUtilController, true
}

// HasStorageFlexUtilController returns a boolean if a field has been set.
func (o *StorageFlexUtilPhysicalDrive) HasStorageFlexUtilController() bool {
	if o != nil && o.StorageFlexUtilController != nil {
		return true
	}

	return false
}

// SetStorageFlexUtilController gets a reference to the given StorageFlexUtilControllerRelationship and assigns it to the StorageFlexUtilController field.
func (o *StorageFlexUtilPhysicalDrive) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship) {
	o.StorageFlexUtilController = &v
}

func (o StorageFlexUtilPhysicalDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BlockSize != nil {
		toSerialize["BlockSize"] = o.BlockSize
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.DrivesEnabled != nil {
		toSerialize["DrivesEnabled"] = o.DrivesEnabled
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.ManufacturerDate != nil {
		toSerialize["ManufacturerDate"] = o.ManufacturerDate
	}
	if o.ManufacturerId != nil {
		toSerialize["ManufacturerId"] = o.ManufacturerId
	}
	if o.OemId != nil {
		toSerialize["OemId"] = o.OemId
	}
	if o.PartitionCount != nil {
		toSerialize["PartitionCount"] = o.PartitionCount
	}
	if o.PdStatus != nil {
		toSerialize["PdStatus"] = o.PdStatus
	}
	if o.PhysicalDrive != nil {
		toSerialize["PhysicalDrive"] = o.PhysicalDrive
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.ProductRevision != nil {
		toSerialize["ProductRevision"] = o.ProductRevision
	}
	if o.ReadErrorCount != nil {
		toSerialize["ReadErrorCount"] = o.ReadErrorCount
	}
	if o.ReadErrorThreshold != nil {
		toSerialize["ReadErrorThreshold"] = o.ReadErrorThreshold
	}
	if o.WriteEnabled != nil {
		toSerialize["WriteEnabled"] = o.WriteEnabled
	}
	if o.WriteErrorCount != nil {
		toSerialize["WriteErrorCount"] = o.WriteErrorCount
	}
	if o.WriteErrorThreshold != nil {
		toSerialize["WriteErrorThreshold"] = o.WriteErrorThreshold
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexUtilController != nil {
		toSerialize["StorageFlexUtilController"] = o.StorageFlexUtilController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexUtilPhysicalDrive) UnmarshalJSON(bytes []byte) (err error) {
	type StorageFlexUtilPhysicalDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Block size of the FlexUtil Physical drive.
		BlockSize *string `json:"BlockSize,omitempty"`
		// Capacity of the FlexUtil Physical drive.
		Capacity *string `json:"Capacity,omitempty"`
		// Type of the Physical Drive Controller.
		Controller *string `json:"Controller,omitempty"`
		// The number of drives enabled in the FlexUtil Physical Drive.
		DrivesEnabled *string `json:"DrivesEnabled,omitempty"`
		// Health of the FlexUtil Physical drive.
		Health *string `json:"Health,omitempty"`
		// Manufacturing date of the FlexUtil Physical Drive.
		ManufacturerDate *string `json:"ManufacturerDate,omitempty"`
		// Manufacturer identity of the FlexUtil Physical Drive.
		ManufacturerId *string `json:"ManufacturerId,omitempty"`
		// The OEM Identifier of the FlexUtil physical drive.
		OemId *string `json:"OemId,omitempty"`
		// The number of partitions present on the FlexUtil Physical Drive.
		PartitionCount *string `json:"PartitionCount,omitempty"`
		// Status of the FlexUtil Physical Drive.
		PdStatus *string `json:"PdStatus,omitempty"`
		// The type of physical drive. Example - microSD.
		PhysicalDrive *string `json:"PhysicalDrive,omitempty"`
		// Product name of the FlexUtil Physical Drive.
		ProductName *string `json:"ProductName,omitempty"`
		// Product revision of the FlexUtil Physical Drive.
		ProductRevision *string `json:"ProductRevision,omitempty"`
		// Read error count of the FlexUtil Physical Drive.
		ReadErrorCount *string `json:"ReadErrorCount,omitempty"`
		// Read error threshold for FlexUtil Physical Drive.
		ReadErrorThreshold *string `json:"ReadErrorThreshold,omitempty"`
		// Write access state of the FlexUtil Physical Drive.
		WriteEnabled *string `json:"WriteEnabled,omitempty"`
		// Write error count of the FlexUtil Physical Drive.
		WriteErrorCount *string `json:"WriteErrorCount,omitempty"`
		// Write error threshold for FlexUtil Physical Drive.
		WriteErrorThreshold       *string                                `json:"WriteErrorThreshold,omitempty"`
		InventoryDeviceInfo       *InventoryDeviceInfoRelationship       `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice          *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
		StorageFlexUtilController *StorageFlexUtilControllerRelationship `json:"StorageFlexUtilController,omitempty"`
	}

	varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct := StorageFlexUtilPhysicalDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexUtilPhysicalDrive := _StorageFlexUtilPhysicalDrive{}
		varStorageFlexUtilPhysicalDrive.ClassId = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ClassId
		varStorageFlexUtilPhysicalDrive.ObjectType = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ObjectType
		varStorageFlexUtilPhysicalDrive.BlockSize = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.BlockSize
		varStorageFlexUtilPhysicalDrive.Capacity = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.Capacity
		varStorageFlexUtilPhysicalDrive.Controller = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.Controller
		varStorageFlexUtilPhysicalDrive.DrivesEnabled = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.DrivesEnabled
		varStorageFlexUtilPhysicalDrive.Health = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.Health
		varStorageFlexUtilPhysicalDrive.ManufacturerDate = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ManufacturerDate
		varStorageFlexUtilPhysicalDrive.ManufacturerId = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ManufacturerId
		varStorageFlexUtilPhysicalDrive.OemId = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.OemId
		varStorageFlexUtilPhysicalDrive.PartitionCount = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.PartitionCount
		varStorageFlexUtilPhysicalDrive.PdStatus = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.PdStatus
		varStorageFlexUtilPhysicalDrive.PhysicalDrive = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.PhysicalDrive
		varStorageFlexUtilPhysicalDrive.ProductName = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ProductName
		varStorageFlexUtilPhysicalDrive.ProductRevision = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ProductRevision
		varStorageFlexUtilPhysicalDrive.ReadErrorCount = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ReadErrorCount
		varStorageFlexUtilPhysicalDrive.ReadErrorThreshold = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.ReadErrorThreshold
		varStorageFlexUtilPhysicalDrive.WriteEnabled = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.WriteEnabled
		varStorageFlexUtilPhysicalDrive.WriteErrorCount = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.WriteErrorCount
		varStorageFlexUtilPhysicalDrive.WriteErrorThreshold = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.WriteErrorThreshold
		varStorageFlexUtilPhysicalDrive.InventoryDeviceInfo = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexUtilPhysicalDrive.RegisteredDevice = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexUtilPhysicalDrive.StorageFlexUtilController = varStorageFlexUtilPhysicalDriveWithoutEmbeddedStruct.StorageFlexUtilController
		*o = StorageFlexUtilPhysicalDrive(varStorageFlexUtilPhysicalDrive)
	} else {
		return err
	}

	varStorageFlexUtilPhysicalDrive := _StorageFlexUtilPhysicalDrive{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilPhysicalDrive)
	if err == nil {
		o.EquipmentBase = varStorageFlexUtilPhysicalDrive.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockSize")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "DrivesEnabled")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "ManufacturerDate")
		delete(additionalProperties, "ManufacturerId")
		delete(additionalProperties, "OemId")
		delete(additionalProperties, "PartitionCount")
		delete(additionalProperties, "PdStatus")
		delete(additionalProperties, "PhysicalDrive")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductRevision")
		delete(additionalProperties, "ReadErrorCount")
		delete(additionalProperties, "ReadErrorThreshold")
		delete(additionalProperties, "WriteEnabled")
		delete(additionalProperties, "WriteErrorCount")
		delete(additionalProperties, "WriteErrorThreshold")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexUtilController")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageFlexUtilPhysicalDrive struct {
	value *StorageFlexUtilPhysicalDrive
	isSet bool
}

func (v NullableStorageFlexUtilPhysicalDrive) Get() *StorageFlexUtilPhysicalDrive {
	return v.value
}

func (v *NullableStorageFlexUtilPhysicalDrive) Set(val *StorageFlexUtilPhysicalDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexUtilPhysicalDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexUtilPhysicalDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexUtilPhysicalDrive(val *StorageFlexUtilPhysicalDrive) *NullableStorageFlexUtilPhysicalDrive {
	return &NullableStorageFlexUtilPhysicalDrive{value: val, isSet: true}
}

func (v NullableStorageFlexUtilPhysicalDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexUtilPhysicalDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
