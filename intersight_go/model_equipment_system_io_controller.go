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
	"reflect"
	"strings"
)

// EquipmentSystemIoController I/O Controller on a chassis which provides the data path to S-series server.
type EquipmentSystemIoController struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The assigned identifier for a chassis.
	ChassisId *string `json:"ChassisId,omitempty"`
	// Connection Path identifies the data path available between IOModule and FI.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// Connection status identifies the status of data path.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field gives a brief information on systemIOController.
	Description *string `json:"Description,omitempty"`
	// This field identifies the CIMC that manages the controller.
	ManagingInstance *string `json:"ManagingInstance,omitempty"`
	// This field identifies the SIOC operational state.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the IO module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for systemIOController.
	Pid *string `json:"Pid,omitempty"`
	// This represents system I/O Controller identifier.
	SystemIoControllerId *int64                               `json:"SystemIoControllerId,omitempty"`
	Cmc                  *ManagementControllerRelationship    `json:"Cmc,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	SharedIoModule       *EquipmentSharedIoModuleRelationship `json:"SharedIoModule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSystemIoController EquipmentSystemIoController

// NewEquipmentSystemIoController instantiates a new EquipmentSystemIoController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSystemIoController(classId string, objectType string) *EquipmentSystemIoController {
	this := EquipmentSystemIoController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSystemIoControllerWithDefaults instantiates a new EquipmentSystemIoController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSystemIoControllerWithDefaults() *EquipmentSystemIoController {
	this := EquipmentSystemIoController{}
	var classId string = "equipment.SystemIoController"
	this.ClassId = classId
	var objectType string = "equipment.SystemIoController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSystemIoController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSystemIoController) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSystemIoController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSystemIoController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetChassisId() string {
	if o == nil || o.ChassisId == nil {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetChassisIdOk() (*string, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *EquipmentSystemIoController) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetConnectionPath() string {
	if o == nil || o.ConnectionPath == nil {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetConnectionPathOk() (*string, bool) {
	if o == nil || o.ConnectionPath == nil {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasConnectionPath() bool {
	if o != nil && o.ConnectionPath != nil {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentSystemIoController) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentSystemIoController) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentSystemIoController) SetDescription(v string) {
	o.Description = &v
}

// GetManagingInstance returns the ManagingInstance field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetManagingInstance() string {
	if o == nil || o.ManagingInstance == nil {
		var ret string
		return ret
	}
	return *o.ManagingInstance
}

// GetManagingInstanceOk returns a tuple with the ManagingInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetManagingInstanceOk() (*string, bool) {
	if o == nil || o.ManagingInstance == nil {
		return nil, false
	}
	return o.ManagingInstance, true
}

// HasManagingInstance returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasManagingInstance() bool {
	if o != nil && o.ManagingInstance != nil {
		return true
	}

	return false
}

// SetManagingInstance gets a reference to the given string and assigns it to the ManagingInstance field.
func (o *EquipmentSystemIoController) SetManagingInstance(v string) {
	o.ManagingInstance = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSystemIoController) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSystemIoController) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentSystemIoController) SetPid(v string) {
	o.Pid = &v
}

// GetSystemIoControllerId returns the SystemIoControllerId field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetSystemIoControllerId() int64 {
	if o == nil || o.SystemIoControllerId == nil {
		var ret int64
		return ret
	}
	return *o.SystemIoControllerId
}

// GetSystemIoControllerIdOk returns a tuple with the SystemIoControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetSystemIoControllerIdOk() (*int64, bool) {
	if o == nil || o.SystemIoControllerId == nil {
		return nil, false
	}
	return o.SystemIoControllerId, true
}

// HasSystemIoControllerId returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasSystemIoControllerId() bool {
	if o != nil && o.SystemIoControllerId != nil {
		return true
	}

	return false
}

// SetSystemIoControllerId gets a reference to the given int64 and assigns it to the SystemIoControllerId field.
func (o *EquipmentSystemIoController) SetSystemIoControllerId(v int64) {
	o.SystemIoControllerId = &v
}

// GetCmc returns the Cmc field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetCmc() ManagementControllerRelationship {
	if o == nil || o.Cmc == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Cmc
}

// GetCmcOk returns a tuple with the Cmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetCmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Cmc == nil {
		return nil, false
	}
	return o.Cmc, true
}

// HasCmc returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasCmc() bool {
	if o != nil && o.Cmc != nil {
		return true
	}

	return false
}

// SetCmc gets a reference to the given ManagementControllerRelationship and assigns it to the Cmc field.
func (o *EquipmentSystemIoController) SetCmc(v ManagementControllerRelationship) {
	o.Cmc = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentSystemIoController) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSystemIoController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSystemIoController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSharedIoModule returns the SharedIoModule field value if set, zero value otherwise.
func (o *EquipmentSystemIoController) GetSharedIoModule() EquipmentSharedIoModuleRelationship {
	if o == nil || o.SharedIoModule == nil {
		var ret EquipmentSharedIoModuleRelationship
		return ret
	}
	return *o.SharedIoModule
}

// GetSharedIoModuleOk returns a tuple with the SharedIoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoController) GetSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool) {
	if o == nil || o.SharedIoModule == nil {
		return nil, false
	}
	return o.SharedIoModule, true
}

// HasSharedIoModule returns a boolean if a field has been set.
func (o *EquipmentSystemIoController) HasSharedIoModule() bool {
	if o != nil && o.SharedIoModule != nil {
		return true
	}

	return false
}

// SetSharedIoModule gets a reference to the given EquipmentSharedIoModuleRelationship and assigns it to the SharedIoModule field.
func (o *EquipmentSystemIoController) SetSharedIoModule(v EquipmentSharedIoModuleRelationship) {
	o.SharedIoModule = &v
}

func (o EquipmentSystemIoController) MarshalJSON() ([]byte, error) {
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
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.ConnectionPath != nil {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ManagingInstance != nil {
		toSerialize["ManagingInstance"] = o.ManagingInstance
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SystemIoControllerId != nil {
		toSerialize["SystemIoControllerId"] = o.SystemIoControllerId
	}
	if o.Cmc != nil {
		toSerialize["Cmc"] = o.Cmc
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SharedIoModule != nil {
		toSerialize["SharedIoModule"] = o.SharedIoModule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSystemIoController) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentSystemIoControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The assigned identifier for a chassis.
		ChassisId *string `json:"ChassisId,omitempty"`
		// Connection Path identifies the data path available between IOModule and FI.
		ConnectionPath *string `json:"ConnectionPath,omitempty"`
		// Connection status identifies the status of data path.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// This field gives a brief information on systemIOController.
		Description *string `json:"Description,omitempty"`
		// This field identifies the CIMC that manages the controller.
		ManagingInstance *string `json:"ManagingInstance,omitempty"`
		// This field identifies the SIOC operational state.
		OperState *string `json:"OperState,omitempty"`
		// Part Number identifier for the IO module.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for systemIOController.
		Pid *string `json:"Pid,omitempty"`
		// This represents system I/O Controller identifier.
		SystemIoControllerId *int64                               `json:"SystemIoControllerId,omitempty"`
		Cmc                  *ManagementControllerRelationship    `json:"Cmc,omitempty"`
		EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		SharedIoModule       *EquipmentSharedIoModuleRelationship `json:"SharedIoModule,omitempty"`
	}

	varEquipmentSystemIoControllerWithoutEmbeddedStruct := EquipmentSystemIoControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentSystemIoControllerWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSystemIoController := _EquipmentSystemIoController{}
		varEquipmentSystemIoController.ClassId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ClassId
		varEquipmentSystemIoController.ObjectType = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ObjectType
		varEquipmentSystemIoController.ChassisId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ChassisId
		varEquipmentSystemIoController.ConnectionPath = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ConnectionPath
		varEquipmentSystemIoController.ConnectionStatus = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ConnectionStatus
		varEquipmentSystemIoController.Description = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Description
		varEquipmentSystemIoController.ManagingInstance = varEquipmentSystemIoControllerWithoutEmbeddedStruct.ManagingInstance
		varEquipmentSystemIoController.OperState = varEquipmentSystemIoControllerWithoutEmbeddedStruct.OperState
		varEquipmentSystemIoController.PartNumber = varEquipmentSystemIoControllerWithoutEmbeddedStruct.PartNumber
		varEquipmentSystemIoController.Pid = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Pid
		varEquipmentSystemIoController.SystemIoControllerId = varEquipmentSystemIoControllerWithoutEmbeddedStruct.SystemIoControllerId
		varEquipmentSystemIoController.Cmc = varEquipmentSystemIoControllerWithoutEmbeddedStruct.Cmc
		varEquipmentSystemIoController.EquipmentChassis = varEquipmentSystemIoControllerWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentSystemIoController.InventoryDeviceInfo = varEquipmentSystemIoControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentSystemIoController.RegisteredDevice = varEquipmentSystemIoControllerWithoutEmbeddedStruct.RegisteredDevice
		varEquipmentSystemIoController.SharedIoModule = varEquipmentSystemIoControllerWithoutEmbeddedStruct.SharedIoModule
		*o = EquipmentSystemIoController(varEquipmentSystemIoController)
	} else {
		return err
	}

	varEquipmentSystemIoController := _EquipmentSystemIoController{}

	err = json.Unmarshal(bytes, &varEquipmentSystemIoController)
	if err == nil {
		o.EquipmentBase = varEquipmentSystemIoController.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ManagingInstance")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SystemIoControllerId")
		delete(additionalProperties, "Cmc")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SharedIoModule")

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

type NullableEquipmentSystemIoController struct {
	value *EquipmentSystemIoController
	isSet bool
}

func (v NullableEquipmentSystemIoController) Get() *EquipmentSystemIoController {
	return v.value
}

func (v *NullableEquipmentSystemIoController) Set(val *EquipmentSystemIoController) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSystemIoController) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSystemIoController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSystemIoController(val *EquipmentSystemIoController) *NullableEquipmentSystemIoController {
	return &NullableEquipmentSystemIoController{value: val, isSet: true}
}

func (v NullableEquipmentSystemIoController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSystemIoController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
