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

// VirtualizationVmwareDatastoreAllOf Definition of the list of properties defined in 'virtualization.VmwareDatastore', excluding properties defined in parent classes.
type VirtualizationVmwareDatastoreAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Shows if this datastore is accessible.
	Accessible *bool `json:"Accessible,omitempty"`
	// Inventory path of the Datastore.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.
	MaintenanceMode *bool `json:"MaintenanceMode,omitempty"`
	// Indicates if this datastore is connected to multiple hosts.
	MultipleHostAccess *bool `json:"MultipleHostAccess,omitempty"`
	// Datastore health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
	Status *string `json:"Status,omitempty"`
	// Indicates if this datastore supports thin provisioning for files.
	ThinProvisioningSupported *bool `json:"ThinProvisioningSupported,omitempty"`
	// Space uncommitted in this datastore in bytes.
	UnCommitted *int64 `json:"UnCommitted,omitempty"`
	// The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').
	Url *string `json:"Url,omitempty"`
	// Number of virtual machine templates relying on (using) this datastore.
	VmTemplateCount  *int64                                            `json:"VmTemplateCount,omitempty"`
	Cluster          *VirtualizationVmwareClusterRelationship          `json:"Cluster,omitempty"`
	Datacenter       *VirtualizationVmwareDatacenterRelationship       `json:"Datacenter,omitempty"`
	DatastoreCluster *VirtualizationVmwareDatastoreClusterRelationship `json:"DatastoreCluster,omitempty"`
	// An array of relationships to virtualizationVmwareHost resources.
	Hosts                []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDatastoreAllOf VirtualizationVmwareDatastoreAllOf

// NewVirtualizationVmwareDatastoreAllOf instantiates a new VirtualizationVmwareDatastoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatastoreAllOf(classId string, objectType string) *VirtualizationVmwareDatastoreAllOf {
	this := VirtualizationVmwareDatastoreAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationVmwareDatastoreAllOfWithDefaults instantiates a new VirtualizationVmwareDatastoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatastoreAllOfWithDefaults() *VirtualizationVmwareDatastoreAllOf {
	this := VirtualizationVmwareDatastoreAllOf{}
	var classId string = "virtualization.VmwareDatastore"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatastore"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDatastoreAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDatastoreAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDatastoreAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDatastoreAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetAccessible() bool {
	if o == nil || o.Accessible == nil {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetAccessibleOk() (*bool, bool) {
	if o == nil || o.Accessible == nil {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasAccessible() bool {
	if o != nil && o.Accessible != nil {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *VirtualizationVmwareDatastoreAllOf) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareDatastoreAllOf) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetMaintenanceMode() bool {
	if o == nil || o.MaintenanceMode == nil {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || o.MaintenanceMode == nil {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasMaintenanceMode() bool {
	if o != nil && o.MaintenanceMode != nil {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *VirtualizationVmwareDatastoreAllOf) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetMultipleHostAccess returns the MultipleHostAccess field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetMultipleHostAccess() bool {
	if o == nil || o.MultipleHostAccess == nil {
		var ret bool
		return ret
	}
	return *o.MultipleHostAccess
}

// GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetMultipleHostAccessOk() (*bool, bool) {
	if o == nil || o.MultipleHostAccess == nil {
		return nil, false
	}
	return o.MultipleHostAccess, true
}

// HasMultipleHostAccess returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasMultipleHostAccess() bool {
	if o != nil && o.MultipleHostAccess != nil {
		return true
	}

	return false
}

// SetMultipleHostAccess gets a reference to the given bool and assigns it to the MultipleHostAccess field.
func (o *VirtualizationVmwareDatastoreAllOf) SetMultipleHostAccess(v bool) {
	o.MultipleHostAccess = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationVmwareDatastoreAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetThinProvisioningSupported returns the ThinProvisioningSupported field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetThinProvisioningSupported() bool {
	if o == nil || o.ThinProvisioningSupported == nil {
		var ret bool
		return ret
	}
	return *o.ThinProvisioningSupported
}

// GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetThinProvisioningSupportedOk() (*bool, bool) {
	if o == nil || o.ThinProvisioningSupported == nil {
		return nil, false
	}
	return o.ThinProvisioningSupported, true
}

// HasThinProvisioningSupported returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasThinProvisioningSupported() bool {
	if o != nil && o.ThinProvisioningSupported != nil {
		return true
	}

	return false
}

// SetThinProvisioningSupported gets a reference to the given bool and assigns it to the ThinProvisioningSupported field.
func (o *VirtualizationVmwareDatastoreAllOf) SetThinProvisioningSupported(v bool) {
	o.ThinProvisioningSupported = &v
}

// GetUnCommitted returns the UnCommitted field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetUnCommitted() int64 {
	if o == nil || o.UnCommitted == nil {
		var ret int64
		return ret
	}
	return *o.UnCommitted
}

// GetUnCommittedOk returns a tuple with the UnCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetUnCommittedOk() (*int64, bool) {
	if o == nil || o.UnCommitted == nil {
		return nil, false
	}
	return o.UnCommitted, true
}

// HasUnCommitted returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasUnCommitted() bool {
	if o != nil && o.UnCommitted != nil {
		return true
	}

	return false
}

// SetUnCommitted gets a reference to the given int64 and assigns it to the UnCommitted field.
func (o *VirtualizationVmwareDatastoreAllOf) SetUnCommitted(v int64) {
	o.UnCommitted = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VirtualizationVmwareDatastoreAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetVmTemplateCount returns the VmTemplateCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetVmTemplateCount() int64 {
	if o == nil || o.VmTemplateCount == nil {
		var ret int64
		return ret
	}
	return *o.VmTemplateCount
}

// GetVmTemplateCountOk returns a tuple with the VmTemplateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetVmTemplateCountOk() (*int64, bool) {
	if o == nil || o.VmTemplateCount == nil {
		return nil, false
	}
	return o.VmTemplateCount, true
}

// HasVmTemplateCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasVmTemplateCount() bool {
	if o != nil && o.VmTemplateCount != nil {
		return true
	}

	return false
}

// SetVmTemplateCount gets a reference to the given int64 and assigns it to the VmTemplateCount field.
func (o *VirtualizationVmwareDatastoreAllOf) SetVmTemplateCount(v int64) {
	o.VmTemplateCount = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetCluster() VirtualizationVmwareClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationVmwareClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationVmwareClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVmwareDatastoreAllOf) SetCluster(v VirtualizationVmwareClusterRelationship) {
	o.Cluster = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareDatastoreAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetDatastoreCluster returns the DatastoreCluster field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreAllOf) GetDatastoreCluster() VirtualizationVmwareDatastoreClusterRelationship {
	if o == nil || o.DatastoreCluster == nil {
		var ret VirtualizationVmwareDatastoreClusterRelationship
		return ret
	}
	return *o.DatastoreCluster
}

// GetDatastoreClusterOk returns a tuple with the DatastoreCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreAllOf) GetDatastoreClusterOk() (*VirtualizationVmwareDatastoreClusterRelationship, bool) {
	if o == nil || o.DatastoreCluster == nil {
		return nil, false
	}
	return o.DatastoreCluster, true
}

// HasDatastoreCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasDatastoreCluster() bool {
	if o != nil && o.DatastoreCluster != nil {
		return true
	}

	return false
}

// SetDatastoreCluster gets a reference to the given VirtualizationVmwareDatastoreClusterRelationship and assigns it to the DatastoreCluster field.
func (o *VirtualizationVmwareDatastoreAllOf) SetDatastoreCluster(v VirtualizationVmwareDatastoreClusterRelationship) {
	o.DatastoreCluster = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastoreAllOf) GetHosts() []VirtualizationVmwareHostRelationship {
	if o == nil {
		var ret []VirtualizationVmwareHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastoreAllOf) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreAllOf) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VirtualizationVmwareHostRelationship and assigns it to the Hosts field.
func (o *VirtualizationVmwareDatastoreAllOf) SetHosts(v []VirtualizationVmwareHostRelationship) {
	o.Hosts = v
}

func (o VirtualizationVmwareDatastoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Accessible != nil {
		toSerialize["Accessible"] = o.Accessible
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.MaintenanceMode != nil {
		toSerialize["MaintenanceMode"] = o.MaintenanceMode
	}
	if o.MultipleHostAccess != nil {
		toSerialize["MultipleHostAccess"] = o.MultipleHostAccess
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ThinProvisioningSupported != nil {
		toSerialize["ThinProvisioningSupported"] = o.ThinProvisioningSupported
	}
	if o.UnCommitted != nil {
		toSerialize["UnCommitted"] = o.UnCommitted
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.VmTemplateCount != nil {
		toSerialize["VmTemplateCount"] = o.VmTemplateCount
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.DatastoreCluster != nil {
		toSerialize["DatastoreCluster"] = o.DatastoreCluster
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDatastoreAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareDatastoreAllOf := _VirtualizationVmwareDatastoreAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareDatastoreAllOf); err == nil {
		*o = VirtualizationVmwareDatastoreAllOf(varVirtualizationVmwareDatastoreAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accessible")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "MaintenanceMode")
		delete(additionalProperties, "MultipleHostAccess")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "ThinProvisioningSupported")
		delete(additionalProperties, "UnCommitted")
		delete(additionalProperties, "Url")
		delete(additionalProperties, "VmTemplateCount")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "DatastoreCluster")
		delete(additionalProperties, "Hosts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareDatastoreAllOf struct {
	value *VirtualizationVmwareDatastoreAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareDatastoreAllOf) Get() *VirtualizationVmwareDatastoreAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareDatastoreAllOf) Set(val *VirtualizationVmwareDatastoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatastoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatastoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatastoreAllOf(val *VirtualizationVmwareDatastoreAllOf) *NullableVirtualizationVmwareDatastoreAllOf {
	return &NullableVirtualizationVmwareDatastoreAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatastoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatastoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
