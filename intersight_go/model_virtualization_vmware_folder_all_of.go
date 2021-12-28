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

// VirtualizationVmwareFolderAllOf Definition of the list of properties defined in 'virtualization.VmwareFolder', excluding properties defined in parent classes.
type VirtualizationVmwareFolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If a folder is internal, it will be set to true.
	Internal *bool `json:"Internal,omitempty"`
	// Inventory path to the folder. Example - /DC/myFolder.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Determines the type of folder. e.g. vCenter folder, VM and Templete Folder, StorageFolder, NetworkFolder, Host and Cluster Folder. * `Unknown` - The type of the folder is unknown. It may not represent that the folder does not exist but indicates that something might be wrong. * `VMTemplateFolder` - The folder contains VMs and VM templates. * `StorageFolder` - The folder contains storage devices. * `HostClusterFolder` - The folder contains hosts and clusters. * `NetworkFolder` - The folder contains network items. * `VcenterFolder` - The folder created under a vCenter or vCenter folder.
	TypeofFolder         *string                                     `json:"TypeofFolder,omitempty"`
	Datacenter           *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	HypervisorManager    *VirtualizationVmwareVcenterRelationship    `json:"HypervisorManager,omitempty"`
	VmwareFolder         *VirtualizationVmwareFolderRelationship     `json:"VmwareFolder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareFolderAllOf VirtualizationVmwareFolderAllOf

// NewVirtualizationVmwareFolderAllOf instantiates a new VirtualizationVmwareFolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareFolderAllOf(classId string, objectType string) *VirtualizationVmwareFolderAllOf {
	this := VirtualizationVmwareFolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var typeofFolder string = "Unknown"
	this.TypeofFolder = &typeofFolder
	return &this
}

// NewVirtualizationVmwareFolderAllOfWithDefaults instantiates a new VirtualizationVmwareFolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareFolderAllOfWithDefaults() *VirtualizationVmwareFolderAllOf {
	this := VirtualizationVmwareFolderAllOf{}
	var classId string = "virtualization.VmwareFolder"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareFolder"
	this.ObjectType = objectType
	var typeofFolder string = "Unknown"
	this.TypeofFolder = &typeofFolder
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareFolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareFolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareFolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareFolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *VirtualizationVmwareFolderAllOf) SetInternal(v bool) {
	o.Internal = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareFolderAllOf) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetTypeofFolder returns the TypeofFolder field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetTypeofFolder() string {
	if o == nil || o.TypeofFolder == nil {
		var ret string
		return ret
	}
	return *o.TypeofFolder
}

// GetTypeofFolderOk returns a tuple with the TypeofFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetTypeofFolderOk() (*string, bool) {
	if o == nil || o.TypeofFolder == nil {
		return nil, false
	}
	return o.TypeofFolder, true
}

// HasTypeofFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasTypeofFolder() bool {
	if o != nil && o.TypeofFolder != nil {
		return true
	}

	return false
}

// SetTypeofFolder gets a reference to the given string and assigns it to the TypeofFolder field.
func (o *VirtualizationVmwareFolderAllOf) SetTypeofFolder(v string) {
	o.TypeofFolder = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareFolderAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetHypervisorManager() VirtualizationVmwareVcenterRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given VirtualizationVmwareVcenterRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationVmwareFolderAllOf) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship) {
	o.HypervisorManager = &v
}

// GetVmwareFolder returns the VmwareFolder field value if set, zero value otherwise.
func (o *VirtualizationVmwareFolderAllOf) GetVmwareFolder() VirtualizationVmwareFolderRelationship {
	if o == nil || o.VmwareFolder == nil {
		var ret VirtualizationVmwareFolderRelationship
		return ret
	}
	return *o.VmwareFolder
}

// GetVmwareFolderOk returns a tuple with the VmwareFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareFolderAllOf) GetVmwareFolderOk() (*VirtualizationVmwareFolderRelationship, bool) {
	if o == nil || o.VmwareFolder == nil {
		return nil, false
	}
	return o.VmwareFolder, true
}

// HasVmwareFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareFolderAllOf) HasVmwareFolder() bool {
	if o != nil && o.VmwareFolder != nil {
		return true
	}

	return false
}

// SetVmwareFolder gets a reference to the given VirtualizationVmwareFolderRelationship and assigns it to the VmwareFolder field.
func (o *VirtualizationVmwareFolderAllOf) SetVmwareFolder(v VirtualizationVmwareFolderRelationship) {
	o.VmwareFolder = &v
}

func (o VirtualizationVmwareFolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.TypeofFolder != nil {
		toSerialize["TypeofFolder"] = o.TypeofFolder
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}
	if o.VmwareFolder != nil {
		toSerialize["VmwareFolder"] = o.VmwareFolder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareFolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareFolderAllOf := _VirtualizationVmwareFolderAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareFolderAllOf); err == nil {
		*o = VirtualizationVmwareFolderAllOf(varVirtualizationVmwareFolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "TypeofFolder")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "HypervisorManager")
		delete(additionalProperties, "VmwareFolder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareFolderAllOf struct {
	value *VirtualizationVmwareFolderAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareFolderAllOf) Get() *VirtualizationVmwareFolderAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareFolderAllOf) Set(val *VirtualizationVmwareFolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareFolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareFolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareFolderAllOf(val *VirtualizationVmwareFolderAllOf) *NullableVirtualizationVmwareFolderAllOf {
	return &NullableVirtualizationVmwareFolderAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareFolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareFolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
