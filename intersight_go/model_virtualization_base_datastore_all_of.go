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

// VirtualizationBaseDatastoreAllOf Definition of the list of properties defined in 'virtualization.BaseDatastore', excluding properties defined in parent classes.
type VirtualizationBaseDatastoreAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                                `json:"ObjectType"`
	Capacity   NullableVirtualizationStorageCapacity `json:"Capacity,omitempty"`
	// Number of hosts attached to or supported-by this datastore.
	HostCount *int64 `json:"HostCount,omitempty"`
	// The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.
	Name *string `json:"Name,omitempty"`
	// A string indicating the type of the datastore (VMFS, NFS, etc). * `Unknown` - The nature of the file system is unknown. * `VMFS` - It is a Virtual Machine Filesystem. * `NFS` - It is a Network File System. * `vSAN` - It is a virtual Storage Area Network file system. * `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server.
	Type *string `json:"Type,omitempty"`
	// Number of virtual machines relying on (using) this datastore.
	VmCount              *int64 `json:"VmCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseDatastoreAllOf VirtualizationBaseDatastoreAllOf

// NewVirtualizationBaseDatastoreAllOf instantiates a new VirtualizationBaseDatastoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseDatastoreAllOf(classId string, objectType string) *VirtualizationBaseDatastoreAllOf {
	this := VirtualizationBaseDatastoreAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// NewVirtualizationBaseDatastoreAllOfWithDefaults instantiates a new VirtualizationBaseDatastoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseDatastoreAllOfWithDefaults() *VirtualizationBaseDatastoreAllOf {
	this := VirtualizationBaseDatastoreAllOf{}
	var classId string = "virtualization.VmwareDatastore"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatastore"
	this.ObjectType = objectType
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseDatastoreAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseDatastoreAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseDatastoreAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseDatastoreAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseDatastoreAllOf) GetCapacity() VirtualizationStorageCapacity {
	if o == nil || o.Capacity.Get() == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseDatastoreAllOf) GetCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasCapacity() bool {
	if o != nil && o.Capacity.IsSet() {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given NullableVirtualizationStorageCapacity and assigns it to the Capacity field.
func (o *VirtualizationBaseDatastoreAllOf) SetCapacity(v VirtualizationStorageCapacity) {
	o.Capacity.Set(&v)
}

// SetCapacityNil sets the value for Capacity to be an explicit nil
func (o *VirtualizationBaseDatastoreAllOf) SetCapacityNil() {
	o.Capacity.Set(nil)
}

// UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
func (o *VirtualizationBaseDatastoreAllOf) UnsetCapacity() {
	o.Capacity.Unset()
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastoreAllOf) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationBaseDatastoreAllOf) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastoreAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseDatastoreAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastoreAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseDatastoreAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastoreAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationBaseDatastoreAllOf) SetType(v string) {
	o.Type = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastoreAllOf) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastoreAllOf) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastoreAllOf) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationBaseDatastoreAllOf) SetVmCount(v int64) {
	o.VmCount = &v
}

func (o VirtualizationBaseDatastoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity.IsSet() {
		toSerialize["Capacity"] = o.Capacity.Get()
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseDatastoreAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseDatastoreAllOf := _VirtualizationBaseDatastoreAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseDatastoreAllOf); err == nil {
		*o = VirtualizationBaseDatastoreAllOf(varVirtualizationBaseDatastoreAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "HostCount")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VmCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseDatastoreAllOf struct {
	value *VirtualizationBaseDatastoreAllOf
	isSet bool
}

func (v NullableVirtualizationBaseDatastoreAllOf) Get() *VirtualizationBaseDatastoreAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseDatastoreAllOf) Set(val *VirtualizationBaseDatastoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseDatastoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseDatastoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseDatastoreAllOf(val *VirtualizationBaseDatastoreAllOf) *NullableVirtualizationBaseDatastoreAllOf {
	return &NullableVirtualizationBaseDatastoreAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseDatastoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseDatastoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}