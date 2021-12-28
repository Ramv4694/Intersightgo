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
	"time"
)

// VirtualizationVmwareHost The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.
type VirtualizationVmwareHost struct {
	VirtualizationBaseHost
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time when this host booted up.
	BootTime *time.Time `json:"BootTime,omitempty"`
	// Indicates if the host is connected to the vCenter. Values are connected, not connected.
	ConnectionState *string `json:"ConnectionState,omitempty"`
	// This field stores the inventory path of a datacenter. Used in host maintenance action.
	DcInvPath  *string  `json:"DcInvPath,omitempty"`
	DnsServers []string `json:"DnsServers,omitempty"`
	// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
	HwPowerState *string `json:"HwPowerState,omitempty"`
	// True if SSH is enabled in the host, false otherwise.
	IsSshEnabled *bool `json:"IsSshEnabled,omitempty"`
	// The count of all network adapters attached to this host.
	NetworkAdapterCount *int64                                          `json:"NetworkAdapterCount,omitempty"`
	NtpServers          []string                                        `json:"NtpServers,omitempty"`
	ResourceConsumed    NullableVirtualizationVmwareResourceConsumption `json:"ResourceConsumed,omitempty"`
	// The count of all storage adapters attached to this host.
	StorageAdapterCount *int64 `json:"StorageAdapterCount,omitempty"`
	// Time zone this host is in.
	TimeZone *string `json:"TimeZone,omitempty"`
	// The identity of this host within vCenter (optional).
	VcenterHostId *string                                     `json:"VcenterHostId,omitempty"`
	Cluster       *VirtualizationVmwareClusterRelationship    `json:"Cluster,omitempty"`
	Datacenter    *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	// An array of relationships to virtualizationVmwareDatastore resources.
	Datastores []VirtualizationVmwareDatastoreRelationship `json:"Datastores,omitempty"`
	// An array of relationships to virtualizationVmwareDistributedNetwork resources.
	DistributedNetworks []VirtualizationVmwareDistributedNetworkRelationship `json:"DistributedNetworks,omitempty"`
	// An array of relationships to virtualizationVmwareDistributedSwitch resources.
	DistributedSwitches  []VirtualizationVmwareDistributedSwitchRelationship `json:"DistributedSwitches,omitempty"`
	HyperFlexNode        *HyperflexNodeRelationship                          `json:"HyperFlexNode,omitempty"`
	Server               *ComputePhysicalSummaryRelationship                 `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareHost VirtualizationVmwareHost

// NewVirtualizationVmwareHost instantiates a new VirtualizationVmwareHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareHost(classId string, objectType string) *VirtualizationVmwareHost {
	this := VirtualizationVmwareHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// NewVirtualizationVmwareHostWithDefaults instantiates a new VirtualizationVmwareHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareHostWithDefaults() *VirtualizationVmwareHost {
	this := VirtualizationVmwareHost{}
	var classId string = "virtualization.VmwareHost"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareHost"
	this.ObjectType = objectType
	var hwPowerState string = "Unknown"
	this.HwPowerState = &hwPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootTime returns the BootTime field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetBootTime() time.Time {
	if o == nil || o.BootTime == nil {
		var ret time.Time
		return ret
	}
	return *o.BootTime
}

// GetBootTimeOk returns a tuple with the BootTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetBootTimeOk() (*time.Time, bool) {
	if o == nil || o.BootTime == nil {
		return nil, false
	}
	return o.BootTime, true
}

// HasBootTime returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasBootTime() bool {
	if o != nil && o.BootTime != nil {
		return true
	}

	return false
}

// SetBootTime gets a reference to the given time.Time and assigns it to the BootTime field.
func (o *VirtualizationVmwareHost) SetBootTime(v time.Time) {
	o.BootTime = &v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetConnectionState() string {
	if o == nil || o.ConnectionState == nil {
		var ret string
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetConnectionStateOk() (*string, bool) {
	if o == nil || o.ConnectionState == nil {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasConnectionState() bool {
	if o != nil && o.ConnectionState != nil {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given string and assigns it to the ConnectionState field.
func (o *VirtualizationVmwareHost) SetConnectionState(v string) {
	o.ConnectionState = &v
}

// GetDcInvPath returns the DcInvPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetDcInvPath() string {
	if o == nil || o.DcInvPath == nil {
		var ret string
		return ret
	}
	return *o.DcInvPath
}

// GetDcInvPathOk returns a tuple with the DcInvPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetDcInvPathOk() (*string, bool) {
	if o == nil || o.DcInvPath == nil {
		return nil, false
	}
	return o.DcInvPath, true
}

// HasDcInvPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDcInvPath() bool {
	if o != nil && o.DcInvPath != nil {
		return true
	}

	return false
}

// SetDcInvPath gets a reference to the given string and assigns it to the DcInvPath field.
func (o *VirtualizationVmwareHost) SetDcInvPath(v string) {
	o.DcInvPath = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetDnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetDnsServersOk() (*[]string, bool) {
	if o == nil || o.DnsServers == nil {
		return nil, false
	}
	return &o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDnsServers() bool {
	if o != nil && o.DnsServers != nil {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *VirtualizationVmwareHost) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetHwPowerState returns the HwPowerState field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetHwPowerState() string {
	if o == nil || o.HwPowerState == nil {
		var ret string
		return ret
	}
	return *o.HwPowerState
}

// GetHwPowerStateOk returns a tuple with the HwPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetHwPowerStateOk() (*string, bool) {
	if o == nil || o.HwPowerState == nil {
		return nil, false
	}
	return o.HwPowerState, true
}

// HasHwPowerState returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasHwPowerState() bool {
	if o != nil && o.HwPowerState != nil {
		return true
	}

	return false
}

// SetHwPowerState gets a reference to the given string and assigns it to the HwPowerState field.
func (o *VirtualizationVmwareHost) SetHwPowerState(v string) {
	o.HwPowerState = &v
}

// GetIsSshEnabled returns the IsSshEnabled field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetIsSshEnabled() bool {
	if o == nil || o.IsSshEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSshEnabled
}

// GetIsSshEnabledOk returns a tuple with the IsSshEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetIsSshEnabledOk() (*bool, bool) {
	if o == nil || o.IsSshEnabled == nil {
		return nil, false
	}
	return o.IsSshEnabled, true
}

// HasIsSshEnabled returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasIsSshEnabled() bool {
	if o != nil && o.IsSshEnabled != nil {
		return true
	}

	return false
}

// SetIsSshEnabled gets a reference to the given bool and assigns it to the IsSshEnabled field.
func (o *VirtualizationVmwareHost) SetIsSshEnabled(v bool) {
	o.IsSshEnabled = &v
}

// GetNetworkAdapterCount returns the NetworkAdapterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetNetworkAdapterCount() int64 {
	if o == nil || o.NetworkAdapterCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkAdapterCount
}

// GetNetworkAdapterCountOk returns a tuple with the NetworkAdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetNetworkAdapterCountOk() (*int64, bool) {
	if o == nil || o.NetworkAdapterCount == nil {
		return nil, false
	}
	return o.NetworkAdapterCount, true
}

// HasNetworkAdapterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasNetworkAdapterCount() bool {
	if o != nil && o.NetworkAdapterCount != nil {
		return true
	}

	return false
}

// SetNetworkAdapterCount gets a reference to the given int64 and assigns it to the NetworkAdapterCount field.
func (o *VirtualizationVmwareHost) SetNetworkAdapterCount(v int64) {
	o.NetworkAdapterCount = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetNtpServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return &o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *VirtualizationVmwareHost) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetResourceConsumed returns the ResourceConsumed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetResourceConsumed() VirtualizationVmwareResourceConsumption {
	if o == nil || o.ResourceConsumed.Get() == nil {
		var ret VirtualizationVmwareResourceConsumption
		return ret
	}
	return *o.ResourceConsumed.Get()
}

// GetResourceConsumedOk returns a tuple with the ResourceConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetResourceConsumedOk() (*VirtualizationVmwareResourceConsumption, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceConsumed.Get(), o.ResourceConsumed.IsSet()
}

// HasResourceConsumed returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasResourceConsumed() bool {
	if o != nil && o.ResourceConsumed.IsSet() {
		return true
	}

	return false
}

// SetResourceConsumed gets a reference to the given NullableVirtualizationVmwareResourceConsumption and assigns it to the ResourceConsumed field.
func (o *VirtualizationVmwareHost) SetResourceConsumed(v VirtualizationVmwareResourceConsumption) {
	o.ResourceConsumed.Set(&v)
}

// SetResourceConsumedNil sets the value for ResourceConsumed to be an explicit nil
func (o *VirtualizationVmwareHost) SetResourceConsumedNil() {
	o.ResourceConsumed.Set(nil)
}

// UnsetResourceConsumed ensures that no value is present for ResourceConsumed, not even an explicit nil
func (o *VirtualizationVmwareHost) UnsetResourceConsumed() {
	o.ResourceConsumed.Unset()
}

// GetStorageAdapterCount returns the StorageAdapterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetStorageAdapterCount() int64 {
	if o == nil || o.StorageAdapterCount == nil {
		var ret int64
		return ret
	}
	return *o.StorageAdapterCount
}

// GetStorageAdapterCountOk returns a tuple with the StorageAdapterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetStorageAdapterCountOk() (*int64, bool) {
	if o == nil || o.StorageAdapterCount == nil {
		return nil, false
	}
	return o.StorageAdapterCount, true
}

// HasStorageAdapterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasStorageAdapterCount() bool {
	if o != nil && o.StorageAdapterCount != nil {
		return true
	}

	return false
}

// SetStorageAdapterCount gets a reference to the given int64 and assigns it to the StorageAdapterCount field.
func (o *VirtualizationVmwareHost) SetStorageAdapterCount(v int64) {
	o.StorageAdapterCount = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *VirtualizationVmwareHost) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetVcenterHostId returns the VcenterHostId field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetVcenterHostId() string {
	if o == nil || o.VcenterHostId == nil {
		var ret string
		return ret
	}
	return *o.VcenterHostId
}

// GetVcenterHostIdOk returns a tuple with the VcenterHostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetVcenterHostIdOk() (*string, bool) {
	if o == nil || o.VcenterHostId == nil {
		return nil, false
	}
	return o.VcenterHostId, true
}

// HasVcenterHostId returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasVcenterHostId() bool {
	if o != nil && o.VcenterHostId != nil {
		return true
	}

	return false
}

// SetVcenterHostId gets a reference to the given string and assigns it to the VcenterHostId field.
func (o *VirtualizationVmwareHost) SetVcenterHostId(v string) {
	o.VcenterHostId = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetCluster() VirtualizationVmwareClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationVmwareClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationVmwareClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVmwareHost) SetCluster(v VirtualizationVmwareClusterRelationship) {
	o.Cluster = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareHost) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetDatastores() []VirtualizationVmwareDatastoreRelationship {
	if o == nil {
		var ret []VirtualizationVmwareDatastoreRelationship
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return &o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []VirtualizationVmwareDatastoreRelationship and assigns it to the Datastores field.
func (o *VirtualizationVmwareHost) SetDatastores(v []VirtualizationVmwareDatastoreRelationship) {
	o.Datastores = v
}

// GetDistributedNetworks returns the DistributedNetworks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetDistributedNetworks() []VirtualizationVmwareDistributedNetworkRelationship {
	if o == nil {
		var ret []VirtualizationVmwareDistributedNetworkRelationship
		return ret
	}
	return o.DistributedNetworks
}

// GetDistributedNetworksOk returns a tuple with the DistributedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetDistributedNetworksOk() (*[]VirtualizationVmwareDistributedNetworkRelationship, bool) {
	if o == nil || o.DistributedNetworks == nil {
		return nil, false
	}
	return &o.DistributedNetworks, true
}

// HasDistributedNetworks returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDistributedNetworks() bool {
	if o != nil && o.DistributedNetworks != nil {
		return true
	}

	return false
}

// SetDistributedNetworks gets a reference to the given []VirtualizationVmwareDistributedNetworkRelationship and assigns it to the DistributedNetworks field.
func (o *VirtualizationVmwareHost) SetDistributedNetworks(v []VirtualizationVmwareDistributedNetworkRelationship) {
	o.DistributedNetworks = v
}

// GetDistributedSwitches returns the DistributedSwitches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareHost) GetDistributedSwitches() []VirtualizationVmwareDistributedSwitchRelationship {
	if o == nil {
		var ret []VirtualizationVmwareDistributedSwitchRelationship
		return ret
	}
	return o.DistributedSwitches
}

// GetDistributedSwitchesOk returns a tuple with the DistributedSwitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareHost) GetDistributedSwitchesOk() (*[]VirtualizationVmwareDistributedSwitchRelationship, bool) {
	if o == nil || o.DistributedSwitches == nil {
		return nil, false
	}
	return &o.DistributedSwitches, true
}

// HasDistributedSwitches returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasDistributedSwitches() bool {
	if o != nil && o.DistributedSwitches != nil {
		return true
	}

	return false
}

// SetDistributedSwitches gets a reference to the given []VirtualizationVmwareDistributedSwitchRelationship and assigns it to the DistributedSwitches field.
func (o *VirtualizationVmwareHost) SetDistributedSwitches(v []VirtualizationVmwareDistributedSwitchRelationship) {
	o.DistributedSwitches = v
}

// GetHyperFlexNode returns the HyperFlexNode field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetHyperFlexNode() HyperflexNodeRelationship {
	if o == nil || o.HyperFlexNode == nil {
		var ret HyperflexNodeRelationship
		return ret
	}
	return *o.HyperFlexNode
}

// GetHyperFlexNodeOk returns a tuple with the HyperFlexNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetHyperFlexNodeOk() (*HyperflexNodeRelationship, bool) {
	if o == nil || o.HyperFlexNode == nil {
		return nil, false
	}
	return o.HyperFlexNode, true
}

// HasHyperFlexNode returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasHyperFlexNode() bool {
	if o != nil && o.HyperFlexNode != nil {
		return true
	}

	return false
}

// SetHyperFlexNode gets a reference to the given HyperflexNodeRelationship and assigns it to the HyperFlexNode field.
func (o *VirtualizationVmwareHost) SetHyperFlexNode(v HyperflexNodeRelationship) {
	o.HyperFlexNode = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *VirtualizationVmwareHost) GetServer() ComputePhysicalSummaryRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalSummaryRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareHost) GetServerOk() (*ComputePhysicalSummaryRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *VirtualizationVmwareHost) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalSummaryRelationship and assigns it to the Server field.
func (o *VirtualizationVmwareHost) SetServer(v ComputePhysicalSummaryRelationship) {
	o.Server = &v
}

func (o VirtualizationVmwareHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHost, errVirtualizationBaseHost := json.Marshal(o.VirtualizationBaseHost)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	errVirtualizationBaseHost = json.Unmarshal([]byte(serializedVirtualizationBaseHost), &toSerialize)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootTime != nil {
		toSerialize["BootTime"] = o.BootTime
	}
	if o.ConnectionState != nil {
		toSerialize["ConnectionState"] = o.ConnectionState
	}
	if o.DcInvPath != nil {
		toSerialize["DcInvPath"] = o.DcInvPath
	}
	if o.DnsServers != nil {
		toSerialize["DnsServers"] = o.DnsServers
	}
	if o.HwPowerState != nil {
		toSerialize["HwPowerState"] = o.HwPowerState
	}
	if o.IsSshEnabled != nil {
		toSerialize["IsSshEnabled"] = o.IsSshEnabled
	}
	if o.NetworkAdapterCount != nil {
		toSerialize["NetworkAdapterCount"] = o.NetworkAdapterCount
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.ResourceConsumed.IsSet() {
		toSerialize["ResourceConsumed"] = o.ResourceConsumed.Get()
	}
	if o.StorageAdapterCount != nil {
		toSerialize["StorageAdapterCount"] = o.StorageAdapterCount
	}
	if o.TimeZone != nil {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.VcenterHostId != nil {
		toSerialize["VcenterHostId"] = o.VcenterHostId
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.Datastores != nil {
		toSerialize["Datastores"] = o.Datastores
	}
	if o.DistributedNetworks != nil {
		toSerialize["DistributedNetworks"] = o.DistributedNetworks
	}
	if o.DistributedSwitches != nil {
		toSerialize["DistributedSwitches"] = o.DistributedSwitches
	}
	if o.HyperFlexNode != nil {
		toSerialize["HyperFlexNode"] = o.HyperFlexNode
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareHost) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time when this host booted up.
		BootTime *time.Time `json:"BootTime,omitempty"`
		// Indicates if the host is connected to the vCenter. Values are connected, not connected.
		ConnectionState *string `json:"ConnectionState,omitempty"`
		// This field stores the inventory path of a datacenter. Used in host maintenance action.
		DcInvPath  *string  `json:"DcInvPath,omitempty"`
		DnsServers []string `json:"DnsServers,omitempty"`
		// Is the host Powered-up or Powered-down. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
		HwPowerState *string `json:"HwPowerState,omitempty"`
		// True if SSH is enabled in the host, false otherwise.
		IsSshEnabled *bool `json:"IsSshEnabled,omitempty"`
		// The count of all network adapters attached to this host.
		NetworkAdapterCount *int64                                          `json:"NetworkAdapterCount,omitempty"`
		NtpServers          []string                                        `json:"NtpServers,omitempty"`
		ResourceConsumed    NullableVirtualizationVmwareResourceConsumption `json:"ResourceConsumed,omitempty"`
		// The count of all storage adapters attached to this host.
		StorageAdapterCount *int64 `json:"StorageAdapterCount,omitempty"`
		// Time zone this host is in.
		TimeZone *string `json:"TimeZone,omitempty"`
		// The identity of this host within vCenter (optional).
		VcenterHostId *string                                     `json:"VcenterHostId,omitempty"`
		Cluster       *VirtualizationVmwareClusterRelationship    `json:"Cluster,omitempty"`
		Datacenter    *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
		// An array of relationships to virtualizationVmwareDatastore resources.
		Datastores []VirtualizationVmwareDatastoreRelationship `json:"Datastores,omitempty"`
		// An array of relationships to virtualizationVmwareDistributedNetwork resources.
		DistributedNetworks []VirtualizationVmwareDistributedNetworkRelationship `json:"DistributedNetworks,omitempty"`
		// An array of relationships to virtualizationVmwareDistributedSwitch resources.
		DistributedSwitches []VirtualizationVmwareDistributedSwitchRelationship `json:"DistributedSwitches,omitempty"`
		HyperFlexNode       *HyperflexNodeRelationship                          `json:"HyperFlexNode,omitempty"`
		Server              *ComputePhysicalSummaryRelationship                 `json:"Server,omitempty"`
	}

	varVirtualizationVmwareHostWithoutEmbeddedStruct := VirtualizationVmwareHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareHostWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareHost := _VirtualizationVmwareHost{}
		varVirtualizationVmwareHost.ClassId = varVirtualizationVmwareHostWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareHost.ObjectType = varVirtualizationVmwareHostWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareHost.BootTime = varVirtualizationVmwareHostWithoutEmbeddedStruct.BootTime
		varVirtualizationVmwareHost.ConnectionState = varVirtualizationVmwareHostWithoutEmbeddedStruct.ConnectionState
		varVirtualizationVmwareHost.DcInvPath = varVirtualizationVmwareHostWithoutEmbeddedStruct.DcInvPath
		varVirtualizationVmwareHost.DnsServers = varVirtualizationVmwareHostWithoutEmbeddedStruct.DnsServers
		varVirtualizationVmwareHost.HwPowerState = varVirtualizationVmwareHostWithoutEmbeddedStruct.HwPowerState
		varVirtualizationVmwareHost.IsSshEnabled = varVirtualizationVmwareHostWithoutEmbeddedStruct.IsSshEnabled
		varVirtualizationVmwareHost.NetworkAdapterCount = varVirtualizationVmwareHostWithoutEmbeddedStruct.NetworkAdapterCount
		varVirtualizationVmwareHost.NtpServers = varVirtualizationVmwareHostWithoutEmbeddedStruct.NtpServers
		varVirtualizationVmwareHost.ResourceConsumed = varVirtualizationVmwareHostWithoutEmbeddedStruct.ResourceConsumed
		varVirtualizationVmwareHost.StorageAdapterCount = varVirtualizationVmwareHostWithoutEmbeddedStruct.StorageAdapterCount
		varVirtualizationVmwareHost.TimeZone = varVirtualizationVmwareHostWithoutEmbeddedStruct.TimeZone
		varVirtualizationVmwareHost.VcenterHostId = varVirtualizationVmwareHostWithoutEmbeddedStruct.VcenterHostId
		varVirtualizationVmwareHost.Cluster = varVirtualizationVmwareHostWithoutEmbeddedStruct.Cluster
		varVirtualizationVmwareHost.Datacenter = varVirtualizationVmwareHostWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareHost.Datastores = varVirtualizationVmwareHostWithoutEmbeddedStruct.Datastores
		varVirtualizationVmwareHost.DistributedNetworks = varVirtualizationVmwareHostWithoutEmbeddedStruct.DistributedNetworks
		varVirtualizationVmwareHost.DistributedSwitches = varVirtualizationVmwareHostWithoutEmbeddedStruct.DistributedSwitches
		varVirtualizationVmwareHost.HyperFlexNode = varVirtualizationVmwareHostWithoutEmbeddedStruct.HyperFlexNode
		varVirtualizationVmwareHost.Server = varVirtualizationVmwareHostWithoutEmbeddedStruct.Server
		*o = VirtualizationVmwareHost(varVirtualizationVmwareHost)
	} else {
		return err
	}

	varVirtualizationVmwareHost := _VirtualizationVmwareHost{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareHost)
	if err == nil {
		o.VirtualizationBaseHost = varVirtualizationVmwareHost.VirtualizationBaseHost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootTime")
		delete(additionalProperties, "ConnectionState")
		delete(additionalProperties, "DcInvPath")
		delete(additionalProperties, "DnsServers")
		delete(additionalProperties, "HwPowerState")
		delete(additionalProperties, "IsSshEnabled")
		delete(additionalProperties, "NetworkAdapterCount")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "ResourceConsumed")
		delete(additionalProperties, "StorageAdapterCount")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "VcenterHostId")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "Datastores")
		delete(additionalProperties, "DistributedNetworks")
		delete(additionalProperties, "DistributedSwitches")
		delete(additionalProperties, "HyperFlexNode")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectVirtualizationBaseHost := reflect.ValueOf(o.VirtualizationBaseHost)
		for i := 0; i < reflectVirtualizationBaseHost.Type().NumField(); i++ {
			t := reflectVirtualizationBaseHost.Type().Field(i)

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

type NullableVirtualizationVmwareHost struct {
	value *VirtualizationVmwareHost
	isSet bool
}

func (v NullableVirtualizationVmwareHost) Get() *VirtualizationVmwareHost {
	return v.value
}

func (v *NullableVirtualizationVmwareHost) Set(val *VirtualizationVmwareHost) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareHost) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareHost(val *VirtualizationVmwareHost) *NullableVirtualizationVmwareHost {
	return &NullableVirtualizationVmwareHost{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
