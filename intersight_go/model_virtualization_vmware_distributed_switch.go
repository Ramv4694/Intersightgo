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
	"reflect"
	"strings"
)

// VirtualizationVmwareDistributedSwitch The VMware Distributed Virtual Switch object is represented here.
type VirtualizationVmwareDistributedSwitch struct {
	VirtualizationBaseDistributedSwitch
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switch description (user provided), if any.
	Description *string `json:"Description,omitempty"`
	// Maximum number of ports allowed on this distributed virtual switch.
	MaxPort *int64 `json:"MaxPort,omitempty"`
	// Maximum transmission unit configured on a distributed virtual switch.
	Mtu                   *int64                                         `json:"Mtu,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// The total number of hosts attached to the distributed virtual switch.
	NumHosts *int64 `json:"NumHosts,omitempty"`
	// The total number of distributed networks in the distributed virtual switch.
	NumNetworks *int64 `json:"NumNetworks,omitempty"`
	// Number of stand-alone ports in use.
	NumStandAlonePorts *int64 `json:"NumStandAlonePorts,omitempty"`
	// Number of uplinks configured in this distributed virtual switch.
	NumUplinks     *int64                                `json:"NumUplinks,omitempty"`
	SwitchCapacity NullableVirtualizationStorageCapacity `json:"SwitchCapacity,omitempty"`
	// Universally Unique Id of this distributed virtual switch.
	Uuid *string `json:"Uuid,omitempty"`
	// The running config's version details are represented.
	Version    *string                                     `json:"Version,omitempty"`
	Datacenter *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	// An array of relationships to virtualizationVmwareHost resources.
	Hosts                []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDistributedSwitch VirtualizationVmwareDistributedSwitch

// NewVirtualizationVmwareDistributedSwitch instantiates a new VirtualizationVmwareDistributedSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDistributedSwitch(classId string, objectType string) *VirtualizationVmwareDistributedSwitch {
	this := VirtualizationVmwareDistributedSwitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareDistributedSwitchWithDefaults instantiates a new VirtualizationVmwareDistributedSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDistributedSwitchWithDefaults() *VirtualizationVmwareDistributedSwitch {
	this := VirtualizationVmwareDistributedSwitch{}
	var classId string = "virtualization.VmwareDistributedSwitch"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDistributedSwitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDistributedSwitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDistributedSwitch) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDistributedSwitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDistributedSwitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationVmwareDistributedSwitch) SetDescription(v string) {
	o.Description = &v
}

// GetMaxPort returns the MaxPort field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetMaxPort() int64 {
	if o == nil || o.MaxPort == nil {
		var ret int64
		return ret
	}
	return *o.MaxPort
}

// GetMaxPortOk returns a tuple with the MaxPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetMaxPortOk() (*int64, bool) {
	if o == nil || o.MaxPort == nil {
		return nil, false
	}
	return o.MaxPort, true
}

// HasMaxPort returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasMaxPort() bool {
	if o != nil && o.MaxPort != nil {
		return true
	}

	return false
}

// SetMaxPort gets a reference to the given int64 and assigns it to the MaxPort field.
func (o *VirtualizationVmwareDistributedSwitch) SetMaxPort(v int64) {
	o.MaxPort = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVmwareDistributedSwitch) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitch) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || o.NicTeamingAndFailover.Get() == nil {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitch) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareDistributedSwitch) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareDistributedSwitch) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareDistributedSwitch) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetNumHosts returns the NumHosts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetNumHosts() int64 {
	if o == nil || o.NumHosts == nil {
		var ret int64
		return ret
	}
	return *o.NumHosts
}

// GetNumHostsOk returns a tuple with the NumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetNumHostsOk() (*int64, bool) {
	if o == nil || o.NumHosts == nil {
		return nil, false
	}
	return o.NumHosts, true
}

// HasNumHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasNumHosts() bool {
	if o != nil && o.NumHosts != nil {
		return true
	}

	return false
}

// SetNumHosts gets a reference to the given int64 and assigns it to the NumHosts field.
func (o *VirtualizationVmwareDistributedSwitch) SetNumHosts(v int64) {
	o.NumHosts = &v
}

// GetNumNetworks returns the NumNetworks field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetNumNetworks() int64 {
	if o == nil || o.NumNetworks == nil {
		var ret int64
		return ret
	}
	return *o.NumNetworks
}

// GetNumNetworksOk returns a tuple with the NumNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetNumNetworksOk() (*int64, bool) {
	if o == nil || o.NumNetworks == nil {
		return nil, false
	}
	return o.NumNetworks, true
}

// HasNumNetworks returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasNumNetworks() bool {
	if o != nil && o.NumNetworks != nil {
		return true
	}

	return false
}

// SetNumNetworks gets a reference to the given int64 and assigns it to the NumNetworks field.
func (o *VirtualizationVmwareDistributedSwitch) SetNumNetworks(v int64) {
	o.NumNetworks = &v
}

// GetNumStandAlonePorts returns the NumStandAlonePorts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetNumStandAlonePorts() int64 {
	if o == nil || o.NumStandAlonePorts == nil {
		var ret int64
		return ret
	}
	return *o.NumStandAlonePorts
}

// GetNumStandAlonePortsOk returns a tuple with the NumStandAlonePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetNumStandAlonePortsOk() (*int64, bool) {
	if o == nil || o.NumStandAlonePorts == nil {
		return nil, false
	}
	return o.NumStandAlonePorts, true
}

// HasNumStandAlonePorts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasNumStandAlonePorts() bool {
	if o != nil && o.NumStandAlonePorts != nil {
		return true
	}

	return false
}

// SetNumStandAlonePorts gets a reference to the given int64 and assigns it to the NumStandAlonePorts field.
func (o *VirtualizationVmwareDistributedSwitch) SetNumStandAlonePorts(v int64) {
	o.NumStandAlonePorts = &v
}

// GetNumUplinks returns the NumUplinks field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetNumUplinks() int64 {
	if o == nil || o.NumUplinks == nil {
		var ret int64
		return ret
	}
	return *o.NumUplinks
}

// GetNumUplinksOk returns a tuple with the NumUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetNumUplinksOk() (*int64, bool) {
	if o == nil || o.NumUplinks == nil {
		return nil, false
	}
	return o.NumUplinks, true
}

// HasNumUplinks returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasNumUplinks() bool {
	if o != nil && o.NumUplinks != nil {
		return true
	}

	return false
}

// SetNumUplinks gets a reference to the given int64 and assigns it to the NumUplinks field.
func (o *VirtualizationVmwareDistributedSwitch) SetNumUplinks(v int64) {
	o.NumUplinks = &v
}

// GetSwitchCapacity returns the SwitchCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitch) GetSwitchCapacity() VirtualizationStorageCapacity {
	if o == nil || o.SwitchCapacity.Get() == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.SwitchCapacity.Get()
}

// GetSwitchCapacityOk returns a tuple with the SwitchCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitch) GetSwitchCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SwitchCapacity.Get(), o.SwitchCapacity.IsSet()
}

// HasSwitchCapacity returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasSwitchCapacity() bool {
	if o != nil && o.SwitchCapacity.IsSet() {
		return true
	}

	return false
}

// SetSwitchCapacity gets a reference to the given NullableVirtualizationStorageCapacity and assigns it to the SwitchCapacity field.
func (o *VirtualizationVmwareDistributedSwitch) SetSwitchCapacity(v VirtualizationStorageCapacity) {
	o.SwitchCapacity.Set(&v)
}

// SetSwitchCapacityNil sets the value for SwitchCapacity to be an explicit nil
func (o *VirtualizationVmwareDistributedSwitch) SetSwitchCapacityNil() {
	o.SwitchCapacity.Set(nil)
}

// UnsetSwitchCapacity ensures that no value is present for SwitchCapacity, not even an explicit nil
func (o *VirtualizationVmwareDistributedSwitch) UnsetSwitchCapacity() {
	o.SwitchCapacity.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationVmwareDistributedSwitch) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationVmwareDistributedSwitch) SetVersion(v string) {
	o.Version = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitch) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitch) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareDistributedSwitch) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitch) GetHosts() []VirtualizationVmwareHostRelationship {
	if o == nil {
		var ret []VirtualizationVmwareHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitch) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitch) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VirtualizationVmwareHostRelationship and assigns it to the Hosts field.
func (o *VirtualizationVmwareDistributedSwitch) SetHosts(v []VirtualizationVmwareHostRelationship) {
	o.Hosts = v
}

func (o VirtualizationVmwareDistributedSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDistributedSwitch, errVirtualizationBaseDistributedSwitch := json.Marshal(o.VirtualizationBaseDistributedSwitch)
	if errVirtualizationBaseDistributedSwitch != nil {
		return []byte{}, errVirtualizationBaseDistributedSwitch
	}
	errVirtualizationBaseDistributedSwitch = json.Unmarshal([]byte(serializedVirtualizationBaseDistributedSwitch), &toSerialize)
	if errVirtualizationBaseDistributedSwitch != nil {
		return []byte{}, errVirtualizationBaseDistributedSwitch
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.MaxPort != nil {
		toSerialize["MaxPort"] = o.MaxPort
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if o.NumHosts != nil {
		toSerialize["NumHosts"] = o.NumHosts
	}
	if o.NumNetworks != nil {
		toSerialize["NumNetworks"] = o.NumNetworks
	}
	if o.NumStandAlonePorts != nil {
		toSerialize["NumStandAlonePorts"] = o.NumStandAlonePorts
	}
	if o.NumUplinks != nil {
		toSerialize["NumUplinks"] = o.NumUplinks
	}
	if o.SwitchCapacity.IsSet() {
		toSerialize["SwitchCapacity"] = o.SwitchCapacity.Get()
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDistributedSwitch) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Switch description (user provided), if any.
		Description *string `json:"Description,omitempty"`
		// Maximum number of ports allowed on this distributed virtual switch.
		MaxPort *int64 `json:"MaxPort,omitempty"`
		// Maximum transmission unit configured on a distributed virtual switch.
		Mtu                   *int64                                         `json:"Mtu,omitempty"`
		NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
		// The total number of hosts attached to the distributed virtual switch.
		NumHosts *int64 `json:"NumHosts,omitempty"`
		// The total number of distributed networks in the distributed virtual switch.
		NumNetworks *int64 `json:"NumNetworks,omitempty"`
		// Number of stand-alone ports in use.
		NumStandAlonePorts *int64 `json:"NumStandAlonePorts,omitempty"`
		// Number of uplinks configured in this distributed virtual switch.
		NumUplinks     *int64                                `json:"NumUplinks,omitempty"`
		SwitchCapacity NullableVirtualizationStorageCapacity `json:"SwitchCapacity,omitempty"`
		// Universally Unique Id of this distributed virtual switch.
		Uuid *string `json:"Uuid,omitempty"`
		// The running config's version details are represented.
		Version    *string                                     `json:"Version,omitempty"`
		Datacenter *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
		// An array of relationships to virtualizationVmwareHost resources.
		Hosts []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	}

	varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct := VirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareDistributedSwitch := _VirtualizationVmwareDistributedSwitch{}
		varVirtualizationVmwareDistributedSwitch.ClassId = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareDistributedSwitch.ObjectType = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareDistributedSwitch.Description = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Description
		varVirtualizationVmwareDistributedSwitch.MaxPort = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.MaxPort
		varVirtualizationVmwareDistributedSwitch.Mtu = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Mtu
		varVirtualizationVmwareDistributedSwitch.NicTeamingAndFailover = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.NicTeamingAndFailover
		varVirtualizationVmwareDistributedSwitch.NumHosts = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.NumHosts
		varVirtualizationVmwareDistributedSwitch.NumNetworks = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.NumNetworks
		varVirtualizationVmwareDistributedSwitch.NumStandAlonePorts = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.NumStandAlonePorts
		varVirtualizationVmwareDistributedSwitch.NumUplinks = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.NumUplinks
		varVirtualizationVmwareDistributedSwitch.SwitchCapacity = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.SwitchCapacity
		varVirtualizationVmwareDistributedSwitch.Uuid = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Uuid
		varVirtualizationVmwareDistributedSwitch.Version = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Version
		varVirtualizationVmwareDistributedSwitch.Datacenter = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareDistributedSwitch.Hosts = varVirtualizationVmwareDistributedSwitchWithoutEmbeddedStruct.Hosts
		*o = VirtualizationVmwareDistributedSwitch(varVirtualizationVmwareDistributedSwitch)
	} else {
		return err
	}

	varVirtualizationVmwareDistributedSwitch := _VirtualizationVmwareDistributedSwitch{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDistributedSwitch)
	if err == nil {
		o.VirtualizationBaseDistributedSwitch = varVirtualizationVmwareDistributedSwitch.VirtualizationBaseDistributedSwitch
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "MaxPort")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "NumHosts")
		delete(additionalProperties, "NumNetworks")
		delete(additionalProperties, "NumStandAlonePorts")
		delete(additionalProperties, "NumUplinks")
		delete(additionalProperties, "SwitchCapacity")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "Hosts")

		// remove fields from embedded structs
		reflectVirtualizationBaseDistributedSwitch := reflect.ValueOf(o.VirtualizationBaseDistributedSwitch)
		for i := 0; i < reflectVirtualizationBaseDistributedSwitch.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDistributedSwitch.Type().Field(i)

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

type NullableVirtualizationVmwareDistributedSwitch struct {
	value *VirtualizationVmwareDistributedSwitch
	isSet bool
}

func (v NullableVirtualizationVmwareDistributedSwitch) Get() *VirtualizationVmwareDistributedSwitch {
	return v.value
}

func (v *NullableVirtualizationVmwareDistributedSwitch) Set(val *VirtualizationVmwareDistributedSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDistributedSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDistributedSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDistributedSwitch(val *VirtualizationVmwareDistributedSwitch) *NullableVirtualizationVmwareDistributedSwitch {
	return &NullableVirtualizationVmwareDistributedSwitch{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDistributedSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDistributedSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
