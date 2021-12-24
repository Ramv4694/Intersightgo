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
	"time"
)

// HyperflexStorageContainerAllOf Definition of the list of properties defined in 'hyperflex.StorageContainer', excluding properties defined in parent classes.
type HyperflexStorageContainerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Storage container accessibility summary. * `NOT_APPLICABLE` - The HyperFlex storage container accessibility summary is not applicable. * `ACCESSIBLE` - The HyperFlex storage container is accessible. * `NOT_ACCESSIBLE` - The HyperFlex storage container is not accessible. * `PARTIALLY_ACCESSIBLE` - The HyperFlex storage container is partially accessible.
	AccessibilitySummary *string `json:"AccessibilitySummary,omitempty"`
	// Storage container data block size in bytes.
	DataBlockSize   *int64                                   `json:"DataBlockSize,omitempty"`
	HostMountStatus []StorageStorageContainerHostMountStatus `json:"HostMountStatus,omitempty"`
	// Indicates whether the storage container has volumes.
	InUse *bool `json:"InUse,omitempty"`
	// Indicates whether the storage container was user-created, or system-created. * `UNKNOWN` - The storage container creator is unknown. * `USER_CREATED` - The storage container was created by a user action. * `INTERNAL` - The storage container was created by the system.
	Kind *string `json:"Kind,omitempty"`
	// Storage container's last access time.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// Storage container's last modified time.
	LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
	// Storage container mount status. Applicable only for NFS type. * `NOT_APPLICABLE` - The HyperFlex storage container mount status is not applicable. * `NORMAL` - The HyperFlex storage container mount status is normal. * `ALERT` - The HyperFlex storage container mount status is alert. * `FAILED` - The HyperFlex storage container mount status is failed.
	MountStatus *string `json:"MountStatus,omitempty"`
	// Storage container mount summary. Applicable only for NFS type. * `NOT_APPLICABLE` - The mount summary is not applicable for this HyperFlex storage container. * `MOUNTED` - The HyperFlex storage container is mounted. * `UNMOUNTED` - The HyperFlex storage container is unmounted. * `MOUNT_FAILURE` - The HyperFlex storage container mount summary is failure. * `UNMOUNT_FAILURE` - The HyperFlex storage container unmount summary is failure.
	MountSummary *string `json:"MountSummary,omitempty"`
	// Provisioned capacity of the storage container in bytes.
	ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
	// Provisioned capacity utilization of all volumes associated with the storage container.
	ProvisionedVolumeCapacityUtilization *float32 `json:"ProvisionedVolumeCapacityUtilization,omitempty"`
	// Storage container type (SMB/NFS/iSCSI). * `NFS` - Storage container created/accesed through NFS protocol. * `SMB` - Storage container created/accessed through SMB protocol. * `iSCSI` - Storage container created/accessed through iSCSI protocol.
	Type *string `json:"Type,omitempty"`
	// Uncompressed bytes on storage container.
	UnCompressedUsedBytes *int64 `json:"UnCompressedUsedBytes,omitempty"`
	// UUID of the datastore/storage container.
	Uuid *string `json:"Uuid,omitempty"`
	// Number of volumes associated with the storage container.
	VolumeCount *int64                        `json:"VolumeCount,omitempty"`
	Cluster     *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to hyperflexVolume resources.
	Volumes              []HyperflexVolumeRelationship `json:"Volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexStorageContainerAllOf HyperflexStorageContainerAllOf

// NewHyperflexStorageContainerAllOf instantiates a new HyperflexStorageContainerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexStorageContainerAllOf(classId string, objectType string) *HyperflexStorageContainerAllOf {
	this := HyperflexStorageContainerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessibilitySummary string = "NOT_APPLICABLE"
	this.AccessibilitySummary = &accessibilitySummary
	var kind string = "UNKNOWN"
	this.Kind = &kind
	var mountStatus string = "NOT_APPLICABLE"
	this.MountStatus = &mountStatus
	var mountSummary string = "NOT_APPLICABLE"
	this.MountSummary = &mountSummary
	var type_ string = "NFS"
	this.Type = &type_
	return &this
}

// NewHyperflexStorageContainerAllOfWithDefaults instantiates a new HyperflexStorageContainerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexStorageContainerAllOfWithDefaults() *HyperflexStorageContainerAllOf {
	this := HyperflexStorageContainerAllOf{}
	var classId string = "hyperflex.StorageContainer"
	this.ClassId = classId
	var objectType string = "hyperflex.StorageContainer"
	this.ObjectType = objectType
	var accessibilitySummary string = "NOT_APPLICABLE"
	this.AccessibilitySummary = &accessibilitySummary
	var kind string = "UNKNOWN"
	this.Kind = &kind
	var mountStatus string = "NOT_APPLICABLE"
	this.MountStatus = &mountStatus
	var mountSummary string = "NOT_APPLICABLE"
	this.MountSummary = &mountSummary
	var type_ string = "NFS"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexStorageContainerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexStorageContainerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexStorageContainerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexStorageContainerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessibilitySummary returns the AccessibilitySummary field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetAccessibilitySummary() string {
	if o == nil || o.AccessibilitySummary == nil {
		var ret string
		return ret
	}
	return *o.AccessibilitySummary
}

// GetAccessibilitySummaryOk returns a tuple with the AccessibilitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetAccessibilitySummaryOk() (*string, bool) {
	if o == nil || o.AccessibilitySummary == nil {
		return nil, false
	}
	return o.AccessibilitySummary, true
}

// HasAccessibilitySummary returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasAccessibilitySummary() bool {
	if o != nil && o.AccessibilitySummary != nil {
		return true
	}

	return false
}

// SetAccessibilitySummary gets a reference to the given string and assigns it to the AccessibilitySummary field.
func (o *HyperflexStorageContainerAllOf) SetAccessibilitySummary(v string) {
	o.AccessibilitySummary = &v
}

// GetDataBlockSize returns the DataBlockSize field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetDataBlockSize() int64 {
	if o == nil || o.DataBlockSize == nil {
		var ret int64
		return ret
	}
	return *o.DataBlockSize
}

// GetDataBlockSizeOk returns a tuple with the DataBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetDataBlockSizeOk() (*int64, bool) {
	if o == nil || o.DataBlockSize == nil {
		return nil, false
	}
	return o.DataBlockSize, true
}

// HasDataBlockSize returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasDataBlockSize() bool {
	if o != nil && o.DataBlockSize != nil {
		return true
	}

	return false
}

// SetDataBlockSize gets a reference to the given int64 and assigns it to the DataBlockSize field.
func (o *HyperflexStorageContainerAllOf) SetDataBlockSize(v int64) {
	o.DataBlockSize = &v
}

// GetHostMountStatus returns the HostMountStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStorageContainerAllOf) GetHostMountStatus() []StorageStorageContainerHostMountStatus {
	if o == nil {
		var ret []StorageStorageContainerHostMountStatus
		return ret
	}
	return o.HostMountStatus
}

// GetHostMountStatusOk returns a tuple with the HostMountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStorageContainerAllOf) GetHostMountStatusOk() (*[]StorageStorageContainerHostMountStatus, bool) {
	if o == nil || o.HostMountStatus == nil {
		return nil, false
	}
	return &o.HostMountStatus, true
}

// HasHostMountStatus returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasHostMountStatus() bool {
	if o != nil && o.HostMountStatus != nil {
		return true
	}

	return false
}

// SetHostMountStatus gets a reference to the given []StorageStorageContainerHostMountStatus and assigns it to the HostMountStatus field.
func (o *HyperflexStorageContainerAllOf) SetHostMountStatus(v []StorageStorageContainerHostMountStatus) {
	o.HostMountStatus = v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetInUse() bool {
	if o == nil || o.InUse == nil {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetInUseOk() (*bool, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *HyperflexStorageContainerAllOf) SetInUse(v bool) {
	o.InUse = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *HyperflexStorageContainerAllOf) SetKind(v string) {
	o.Kind = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetLastAccessTime() time.Time {
	if o == nil || o.LastAccessTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *HyperflexStorageContainerAllOf) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *HyperflexStorageContainerAllOf) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMountStatus returns the MountStatus field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetMountStatus() string {
	if o == nil || o.MountStatus == nil {
		var ret string
		return ret
	}
	return *o.MountStatus
}

// GetMountStatusOk returns a tuple with the MountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetMountStatusOk() (*string, bool) {
	if o == nil || o.MountStatus == nil {
		return nil, false
	}
	return o.MountStatus, true
}

// HasMountStatus returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasMountStatus() bool {
	if o != nil && o.MountStatus != nil {
		return true
	}

	return false
}

// SetMountStatus gets a reference to the given string and assigns it to the MountStatus field.
func (o *HyperflexStorageContainerAllOf) SetMountStatus(v string) {
	o.MountStatus = &v
}

// GetMountSummary returns the MountSummary field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetMountSummary() string {
	if o == nil || o.MountSummary == nil {
		var ret string
		return ret
	}
	return *o.MountSummary
}

// GetMountSummaryOk returns a tuple with the MountSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetMountSummaryOk() (*string, bool) {
	if o == nil || o.MountSummary == nil {
		return nil, false
	}
	return o.MountSummary, true
}

// HasMountSummary returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasMountSummary() bool {
	if o != nil && o.MountSummary != nil {
		return true
	}

	return false
}

// SetMountSummary gets a reference to the given string and assigns it to the MountSummary field.
func (o *HyperflexStorageContainerAllOf) SetMountSummary(v string) {
	o.MountSummary = &v
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetProvisionedCapacity() int64 {
	if o == nil || o.ProvisionedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedCapacity
}

// GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetProvisionedCapacityOk() (*int64, bool) {
	if o == nil || o.ProvisionedCapacity == nil {
		return nil, false
	}
	return o.ProvisionedCapacity, true
}

// HasProvisionedCapacity returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasProvisionedCapacity() bool {
	if o != nil && o.ProvisionedCapacity != nil {
		return true
	}

	return false
}

// SetProvisionedCapacity gets a reference to the given int64 and assigns it to the ProvisionedCapacity field.
func (o *HyperflexStorageContainerAllOf) SetProvisionedCapacity(v int64) {
	o.ProvisionedCapacity = &v
}

// GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilization() float32 {
	if o == nil || o.ProvisionedVolumeCapacityUtilization == nil {
		var ret float32
		return ret
	}
	return *o.ProvisionedVolumeCapacityUtilization
}

// GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool) {
	if o == nil || o.ProvisionedVolumeCapacityUtilization == nil {
		return nil, false
	}
	return o.ProvisionedVolumeCapacityUtilization, true
}

// HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasProvisionedVolumeCapacityUtilization() bool {
	if o != nil && o.ProvisionedVolumeCapacityUtilization != nil {
		return true
	}

	return false
}

// SetProvisionedVolumeCapacityUtilization gets a reference to the given float32 and assigns it to the ProvisionedVolumeCapacityUtilization field.
func (o *HyperflexStorageContainerAllOf) SetProvisionedVolumeCapacityUtilization(v float32) {
	o.ProvisionedVolumeCapacityUtilization = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HyperflexStorageContainerAllOf) SetType(v string) {
	o.Type = &v
}

// GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetUnCompressedUsedBytes() int64 {
	if o == nil || o.UnCompressedUsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UnCompressedUsedBytes
}

// GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetUnCompressedUsedBytesOk() (*int64, bool) {
	if o == nil || o.UnCompressedUsedBytes == nil {
		return nil, false
	}
	return o.UnCompressedUsedBytes, true
}

// HasUnCompressedUsedBytes returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasUnCompressedUsedBytes() bool {
	if o != nil && o.UnCompressedUsedBytes != nil {
		return true
	}

	return false
}

// SetUnCompressedUsedBytes gets a reference to the given int64 and assigns it to the UnCompressedUsedBytes field.
func (o *HyperflexStorageContainerAllOf) SetUnCompressedUsedBytes(v int64) {
	o.UnCompressedUsedBytes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexStorageContainerAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeCount returns the VolumeCount field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetVolumeCount() int64 {
	if o == nil || o.VolumeCount == nil {
		var ret int64
		return ret
	}
	return *o.VolumeCount
}

// GetVolumeCountOk returns a tuple with the VolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetVolumeCountOk() (*int64, bool) {
	if o == nil || o.VolumeCount == nil {
		return nil, false
	}
	return o.VolumeCount, true
}

// HasVolumeCount returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasVolumeCount() bool {
	if o != nil && o.VolumeCount != nil {
		return true
	}

	return false
}

// SetVolumeCount gets a reference to the given int64 and assigns it to the VolumeCount field.
func (o *HyperflexStorageContainerAllOf) SetVolumeCount(v int64) {
	o.VolumeCount = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexStorageContainerAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStorageContainerAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexStorageContainerAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStorageContainerAllOf) GetVolumes() []HyperflexVolumeRelationship {
	if o == nil {
		var ret []HyperflexVolumeRelationship
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStorageContainerAllOf) GetVolumesOk() (*[]HyperflexVolumeRelationship, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *HyperflexStorageContainerAllOf) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []HyperflexVolumeRelationship and assigns it to the Volumes field.
func (o *HyperflexStorageContainerAllOf) SetVolumes(v []HyperflexVolumeRelationship) {
	o.Volumes = v
}

func (o HyperflexStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessibilitySummary != nil {
		toSerialize["AccessibilitySummary"] = o.AccessibilitySummary
	}
	if o.DataBlockSize != nil {
		toSerialize["DataBlockSize"] = o.DataBlockSize
	}
	if o.HostMountStatus != nil {
		toSerialize["HostMountStatus"] = o.HostMountStatus
	}
	if o.InUse != nil {
		toSerialize["InUse"] = o.InUse
	}
	if o.Kind != nil {
		toSerialize["Kind"] = o.Kind
	}
	if o.LastAccessTime != nil {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if o.LastModifiedTime != nil {
		toSerialize["LastModifiedTime"] = o.LastModifiedTime
	}
	if o.MountStatus != nil {
		toSerialize["MountStatus"] = o.MountStatus
	}
	if o.MountSummary != nil {
		toSerialize["MountSummary"] = o.MountSummary
	}
	if o.ProvisionedCapacity != nil {
		toSerialize["ProvisionedCapacity"] = o.ProvisionedCapacity
	}
	if o.ProvisionedVolumeCapacityUtilization != nil {
		toSerialize["ProvisionedVolumeCapacityUtilization"] = o.ProvisionedVolumeCapacityUtilization
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.UnCompressedUsedBytes != nil {
		toSerialize["UnCompressedUsedBytes"] = o.UnCompressedUsedBytes
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeCount != nil {
		toSerialize["VolumeCount"] = o.VolumeCount
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexStorageContainerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexStorageContainerAllOf := _HyperflexStorageContainerAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexStorageContainerAllOf); err == nil {
		*o = HyperflexStorageContainerAllOf(varHyperflexStorageContainerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessibilitySummary")
		delete(additionalProperties, "DataBlockSize")
		delete(additionalProperties, "HostMountStatus")
		delete(additionalProperties, "InUse")
		delete(additionalProperties, "Kind")
		delete(additionalProperties, "LastAccessTime")
		delete(additionalProperties, "LastModifiedTime")
		delete(additionalProperties, "MountStatus")
		delete(additionalProperties, "MountSummary")
		delete(additionalProperties, "ProvisionedCapacity")
		delete(additionalProperties, "ProvisionedVolumeCapacityUtilization")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "UnCompressedUsedBytes")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeCount")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Volumes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexStorageContainerAllOf struct {
	value *HyperflexStorageContainerAllOf
	isSet bool
}

func (v NullableHyperflexStorageContainerAllOf) Get() *HyperflexStorageContainerAllOf {
	return v.value
}

func (v *NullableHyperflexStorageContainerAllOf) Set(val *HyperflexStorageContainerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexStorageContainerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexStorageContainerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexStorageContainerAllOf(val *HyperflexStorageContainerAllOf) *NullableHyperflexStorageContainerAllOf {
	return &NullableHyperflexStorageContainerAllOf{value: val, isSet: true}
}

func (v NullableHyperflexStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexStorageContainerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
