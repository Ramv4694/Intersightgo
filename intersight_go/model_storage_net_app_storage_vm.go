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

// StorageNetAppStorageVm NetApp Storage Virtual Machines contain data volumes and one or more Logical Interfaces ( LIFs ) through which they serve data to the clients.
type StorageNetAppStorageVm struct {
	StorageBaseTenant
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                  `json:"ObjectType"`
	Aggregates            []string                                `json:"Aggregates,omitempty"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.
	CifsEnabled *bool    `json:"CifsEnabled,omitempty"`
	DnsDomains  []string `json:"DnsDomains,omitempty"`
	// Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.
	FcpEnabled *bool `json:"FcpEnabled,omitempty"`
	// IPspace name. IPspaces are distinct IP address spaces in which storage virtual machines (SVMs) reside.
	Ipspace *string `json:"Ipspace,omitempty"`
	// Status for iSCSI protocol allowed to run on Vservers.
	IscsiEnabled *bool `json:"IscsiEnabled,omitempty"`
	// Unique identifier of VServer across data center.
	Key         *string  `json:"Key,omitempty"`
	NameServers []string `json:"NameServers,omitempty"`
	// Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.
	NfsEnabled *bool `json:"NfsEnabled,omitempty"`
	// Status for Network File System Protocol ( NFSv3 ) allowed to run on  Vservers.
	NfsV3Enabled *bool `json:"NfsV3Enabled,omitempty"`
	// Status for Network File System Protocol ( NFSv4.1 ) allowed to run on  Vservers.
	NfsV41Enabled *bool `json:"NfsV41Enabled,omitempty"`
	// Status for Network File System Protocol ( NFSv4 ) allowed to run on  Vservers.
	NfsV4Enabled *bool `json:"NfsV4Enabled,omitempty"`
	// Status for NVME protocol allowed to run on Vservers.
	NvmeEnabled *bool `json:"NvmeEnabled,omitempty"`
	// SVM subtype (default, dp_destination, sync_source, or sync_destination). The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster.
	Subtype *string                           `json:"Subtype,omitempty"`
	Array   *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	// An array of relationships to storageNetAppSvmEvent resources.
	Events               []StorageNetAppSvmEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageVm StorageNetAppStorageVm

// NewStorageNetAppStorageVm instantiates a new StorageNetAppStorageVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageVm(classId string, objectType string) *StorageNetAppStorageVm {
	this := StorageNetAppStorageVm{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageVmWithDefaults instantiates a new StorageNetAppStorageVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageVmWithDefaults() *StorageNetAppStorageVm {
	this := StorageNetAppStorageVm{}
	var classId string = "storage.NetAppStorageVm"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageVm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageVm) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageVm) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageVm) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageVm) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVm) GetAggregates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVm) GetAggregatesOk() (*[]string, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return &o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasAggregates() bool {
	if o != nil && o.Aggregates != nil {
		return true
	}

	return false
}

// SetAggregates gets a reference to the given []string and assigns it to the Aggregates field.
func (o *StorageNetAppStorageVm) SetAggregates(v []string) {
	o.Aggregates = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppStorageVm) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCifsEnabled returns the CifsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetCifsEnabled() bool {
	if o == nil || o.CifsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CifsEnabled
}

// GetCifsEnabledOk returns a tuple with the CifsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetCifsEnabledOk() (*bool, bool) {
	if o == nil || o.CifsEnabled == nil {
		return nil, false
	}
	return o.CifsEnabled, true
}

// HasCifsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasCifsEnabled() bool {
	if o != nil && o.CifsEnabled != nil {
		return true
	}

	return false
}

// SetCifsEnabled gets a reference to the given bool and assigns it to the CifsEnabled field.
func (o *StorageNetAppStorageVm) SetCifsEnabled(v bool) {
	o.CifsEnabled = &v
}

// GetDnsDomains returns the DnsDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVm) GetDnsDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsDomains
}

// GetDnsDomainsOk returns a tuple with the DnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVm) GetDnsDomainsOk() (*[]string, bool) {
	if o == nil || o.DnsDomains == nil {
		return nil, false
	}
	return &o.DnsDomains, true
}

// HasDnsDomains returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasDnsDomains() bool {
	if o != nil && o.DnsDomains != nil {
		return true
	}

	return false
}

// SetDnsDomains gets a reference to the given []string and assigns it to the DnsDomains field.
func (o *StorageNetAppStorageVm) SetDnsDomains(v []string) {
	o.DnsDomains = v
}

// GetFcpEnabled returns the FcpEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetFcpEnabled() bool {
	if o == nil || o.FcpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FcpEnabled
}

// GetFcpEnabledOk returns a tuple with the FcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetFcpEnabledOk() (*bool, bool) {
	if o == nil || o.FcpEnabled == nil {
		return nil, false
	}
	return o.FcpEnabled, true
}

// HasFcpEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasFcpEnabled() bool {
	if o != nil && o.FcpEnabled != nil {
		return true
	}

	return false
}

// SetFcpEnabled gets a reference to the given bool and assigns it to the FcpEnabled field.
func (o *StorageNetAppStorageVm) SetFcpEnabled(v bool) {
	o.FcpEnabled = &v
}

// GetIpspace returns the Ipspace field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetIpspace() string {
	if o == nil || o.Ipspace == nil {
		var ret string
		return ret
	}
	return *o.Ipspace
}

// GetIpspaceOk returns a tuple with the Ipspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetIpspaceOk() (*string, bool) {
	if o == nil || o.Ipspace == nil {
		return nil, false
	}
	return o.Ipspace, true
}

// HasIpspace returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasIpspace() bool {
	if o != nil && o.Ipspace != nil {
		return true
	}

	return false
}

// SetIpspace gets a reference to the given string and assigns it to the Ipspace field.
func (o *StorageNetAppStorageVm) SetIpspace(v string) {
	o.Ipspace = &v
}

// GetIscsiEnabled returns the IscsiEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetIscsiEnabled() bool {
	if o == nil || o.IscsiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IscsiEnabled
}

// GetIscsiEnabledOk returns a tuple with the IscsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetIscsiEnabledOk() (*bool, bool) {
	if o == nil || o.IscsiEnabled == nil {
		return nil, false
	}
	return o.IscsiEnabled, true
}

// HasIscsiEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasIscsiEnabled() bool {
	if o != nil && o.IscsiEnabled != nil {
		return true
	}

	return false
}

// SetIscsiEnabled gets a reference to the given bool and assigns it to the IscsiEnabled field.
func (o *StorageNetAppStorageVm) SetIscsiEnabled(v bool) {
	o.IscsiEnabled = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppStorageVm) SetKey(v string) {
	o.Key = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVm) GetNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVm) GetNameServersOk() (*[]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return &o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *StorageNetAppStorageVm) SetNameServers(v []string) {
	o.NameServers = v
}

// GetNfsEnabled returns the NfsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetNfsEnabled() bool {
	if o == nil || o.NfsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsEnabled
}

// GetNfsEnabledOk returns a tuple with the NfsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetNfsEnabledOk() (*bool, bool) {
	if o == nil || o.NfsEnabled == nil {
		return nil, false
	}
	return o.NfsEnabled, true
}

// HasNfsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNfsEnabled() bool {
	if o != nil && o.NfsEnabled != nil {
		return true
	}

	return false
}

// SetNfsEnabled gets a reference to the given bool and assigns it to the NfsEnabled field.
func (o *StorageNetAppStorageVm) SetNfsEnabled(v bool) {
	o.NfsEnabled = &v
}

// GetNfsV3Enabled returns the NfsV3Enabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetNfsV3Enabled() bool {
	if o == nil || o.NfsV3Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV3Enabled
}

// GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetNfsV3EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV3Enabled == nil {
		return nil, false
	}
	return o.NfsV3Enabled, true
}

// HasNfsV3Enabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNfsV3Enabled() bool {
	if o != nil && o.NfsV3Enabled != nil {
		return true
	}

	return false
}

// SetNfsV3Enabled gets a reference to the given bool and assigns it to the NfsV3Enabled field.
func (o *StorageNetAppStorageVm) SetNfsV3Enabled(v bool) {
	o.NfsV3Enabled = &v
}

// GetNfsV41Enabled returns the NfsV41Enabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetNfsV41Enabled() bool {
	if o == nil || o.NfsV41Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV41Enabled
}

// GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetNfsV41EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV41Enabled == nil {
		return nil, false
	}
	return o.NfsV41Enabled, true
}

// HasNfsV41Enabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNfsV41Enabled() bool {
	if o != nil && o.NfsV41Enabled != nil {
		return true
	}

	return false
}

// SetNfsV41Enabled gets a reference to the given bool and assigns it to the NfsV41Enabled field.
func (o *StorageNetAppStorageVm) SetNfsV41Enabled(v bool) {
	o.NfsV41Enabled = &v
}

// GetNfsV4Enabled returns the NfsV4Enabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetNfsV4Enabled() bool {
	if o == nil || o.NfsV4Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV4Enabled
}

// GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetNfsV4EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV4Enabled == nil {
		return nil, false
	}
	return o.NfsV4Enabled, true
}

// HasNfsV4Enabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNfsV4Enabled() bool {
	if o != nil && o.NfsV4Enabled != nil {
		return true
	}

	return false
}

// SetNfsV4Enabled gets a reference to the given bool and assigns it to the NfsV4Enabled field.
func (o *StorageNetAppStorageVm) SetNfsV4Enabled(v bool) {
	o.NfsV4Enabled = &v
}

// GetNvmeEnabled returns the NvmeEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetNvmeEnabled() bool {
	if o == nil || o.NvmeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NvmeEnabled
}

// GetNvmeEnabledOk returns a tuple with the NvmeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetNvmeEnabledOk() (*bool, bool) {
	if o == nil || o.NvmeEnabled == nil {
		return nil, false
	}
	return o.NvmeEnabled, true
}

// HasNvmeEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasNvmeEnabled() bool {
	if o != nil && o.NvmeEnabled != nil {
		return true
	}

	return false
}

// SetNvmeEnabled gets a reference to the given bool and assigns it to the NvmeEnabled field.
func (o *StorageNetAppStorageVm) SetNvmeEnabled(v bool) {
	o.NvmeEnabled = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *StorageNetAppStorageVm) SetSubtype(v string) {
	o.Subtype = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppStorageVm) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVm) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppStorageVm) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVm) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVm) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppStorageVm) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVm) GetEvents() []StorageNetAppSvmEventRelationship {
	if o == nil {
		var ret []StorageNetAppSvmEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVm) GetEventsOk() (*[]StorageNetAppSvmEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppStorageVm) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppSvmEventRelationship and assigns it to the Events field.
func (o *StorageNetAppStorageVm) SetEvents(v []StorageNetAppSvmEventRelationship) {
	o.Events = v
}

func (o StorageNetAppStorageVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseTenant, errStorageBaseTenant := json.Marshal(o.StorageBaseTenant)
	if errStorageBaseTenant != nil {
		return []byte{}, errStorageBaseTenant
	}
	errStorageBaseTenant = json.Unmarshal([]byte(serializedStorageBaseTenant), &toSerialize)
	if errStorageBaseTenant != nil {
		return []byte{}, errStorageBaseTenant
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Aggregates != nil {
		toSerialize["Aggregates"] = o.Aggregates
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CifsEnabled != nil {
		toSerialize["CifsEnabled"] = o.CifsEnabled
	}
	if o.DnsDomains != nil {
		toSerialize["DnsDomains"] = o.DnsDomains
	}
	if o.FcpEnabled != nil {
		toSerialize["FcpEnabled"] = o.FcpEnabled
	}
	if o.Ipspace != nil {
		toSerialize["Ipspace"] = o.Ipspace
	}
	if o.IscsiEnabled != nil {
		toSerialize["IscsiEnabled"] = o.IscsiEnabled
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.NameServers != nil {
		toSerialize["NameServers"] = o.NameServers
	}
	if o.NfsEnabled != nil {
		toSerialize["NfsEnabled"] = o.NfsEnabled
	}
	if o.NfsV3Enabled != nil {
		toSerialize["NfsV3Enabled"] = o.NfsV3Enabled
	}
	if o.NfsV41Enabled != nil {
		toSerialize["NfsV41Enabled"] = o.NfsV41Enabled
	}
	if o.NfsV4Enabled != nil {
		toSerialize["NfsV4Enabled"] = o.NfsV4Enabled
	}
	if o.NvmeEnabled != nil {
		toSerialize["NvmeEnabled"] = o.NvmeEnabled
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppStorageVm) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppStorageVmWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType            string                                  `json:"ObjectType"`
		Aggregates            []string                                `json:"Aggregates,omitempty"`
		AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
		// Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.
		CifsEnabled *bool    `json:"CifsEnabled,omitempty"`
		DnsDomains  []string `json:"DnsDomains,omitempty"`
		// Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.
		FcpEnabled *bool `json:"FcpEnabled,omitempty"`
		// IPspace name. IPspaces are distinct IP address spaces in which storage virtual machines (SVMs) reside.
		Ipspace *string `json:"Ipspace,omitempty"`
		// Status for iSCSI protocol allowed to run on Vservers.
		IscsiEnabled *bool `json:"IscsiEnabled,omitempty"`
		// Unique identifier of VServer across data center.
		Key         *string  `json:"Key,omitempty"`
		NameServers []string `json:"NameServers,omitempty"`
		// Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.
		NfsEnabled *bool `json:"NfsEnabled,omitempty"`
		// Status for Network File System Protocol ( NFSv3 ) allowed to run on  Vservers.
		NfsV3Enabled *bool `json:"NfsV3Enabled,omitempty"`
		// Status for Network File System Protocol ( NFSv4.1 ) allowed to run on  Vservers.
		NfsV41Enabled *bool `json:"NfsV41Enabled,omitempty"`
		// Status for Network File System Protocol ( NFSv4 ) allowed to run on  Vservers.
		NfsV4Enabled *bool `json:"NfsV4Enabled,omitempty"`
		// Status for NVME protocol allowed to run on Vservers.
		NvmeEnabled *bool `json:"NvmeEnabled,omitempty"`
		// SVM subtype (default, dp_destination, sync_source, or sync_destination). The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster.
		Subtype *string                           `json:"Subtype,omitempty"`
		Array   *StorageNetAppClusterRelationship `json:"Array,omitempty"`
		// An array of relationships to storageNetAppAggregate resources.
		DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
		// An array of relationships to storageNetAppSvmEvent resources.
		Events []StorageNetAppSvmEventRelationship `json:"Events,omitempty"`
	}

	varStorageNetAppStorageVmWithoutEmbeddedStruct := StorageNetAppStorageVmWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppStorageVmWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppStorageVm := _StorageNetAppStorageVm{}
		varStorageNetAppStorageVm.ClassId = varStorageNetAppStorageVmWithoutEmbeddedStruct.ClassId
		varStorageNetAppStorageVm.ObjectType = varStorageNetAppStorageVmWithoutEmbeddedStruct.ObjectType
		varStorageNetAppStorageVm.Aggregates = varStorageNetAppStorageVmWithoutEmbeddedStruct.Aggregates
		varStorageNetAppStorageVm.AvgPerformanceMetrics = varStorageNetAppStorageVmWithoutEmbeddedStruct.AvgPerformanceMetrics
		varStorageNetAppStorageVm.CifsEnabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.CifsEnabled
		varStorageNetAppStorageVm.DnsDomains = varStorageNetAppStorageVmWithoutEmbeddedStruct.DnsDomains
		varStorageNetAppStorageVm.FcpEnabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.FcpEnabled
		varStorageNetAppStorageVm.Ipspace = varStorageNetAppStorageVmWithoutEmbeddedStruct.Ipspace
		varStorageNetAppStorageVm.IscsiEnabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.IscsiEnabled
		varStorageNetAppStorageVm.Key = varStorageNetAppStorageVmWithoutEmbeddedStruct.Key
		varStorageNetAppStorageVm.NameServers = varStorageNetAppStorageVmWithoutEmbeddedStruct.NameServers
		varStorageNetAppStorageVm.NfsEnabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.NfsEnabled
		varStorageNetAppStorageVm.NfsV3Enabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.NfsV3Enabled
		varStorageNetAppStorageVm.NfsV41Enabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.NfsV41Enabled
		varStorageNetAppStorageVm.NfsV4Enabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.NfsV4Enabled
		varStorageNetAppStorageVm.NvmeEnabled = varStorageNetAppStorageVmWithoutEmbeddedStruct.NvmeEnabled
		varStorageNetAppStorageVm.Subtype = varStorageNetAppStorageVmWithoutEmbeddedStruct.Subtype
		varStorageNetAppStorageVm.Array = varStorageNetAppStorageVmWithoutEmbeddedStruct.Array
		varStorageNetAppStorageVm.DiskPool = varStorageNetAppStorageVmWithoutEmbeddedStruct.DiskPool
		varStorageNetAppStorageVm.Events = varStorageNetAppStorageVmWithoutEmbeddedStruct.Events
		*o = StorageNetAppStorageVm(varStorageNetAppStorageVm)
	} else {
		return err
	}

	varStorageNetAppStorageVm := _StorageNetAppStorageVm{}

	err = json.Unmarshal(bytes, &varStorageNetAppStorageVm)
	if err == nil {
		o.StorageBaseTenant = varStorageNetAppStorageVm.StorageBaseTenant
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Aggregates")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CifsEnabled")
		delete(additionalProperties, "DnsDomains")
		delete(additionalProperties, "FcpEnabled")
		delete(additionalProperties, "Ipspace")
		delete(additionalProperties, "IscsiEnabled")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "NameServers")
		delete(additionalProperties, "NfsEnabled")
		delete(additionalProperties, "NfsV3Enabled")
		delete(additionalProperties, "NfsV41Enabled")
		delete(additionalProperties, "NfsV4Enabled")
		delete(additionalProperties, "NvmeEnabled")
		delete(additionalProperties, "Subtype")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "DiskPool")
		delete(additionalProperties, "Events")

		// remove fields from embedded structs
		reflectStorageBaseTenant := reflect.ValueOf(o.StorageBaseTenant)
		for i := 0; i < reflectStorageBaseTenant.Type().NumField(); i++ {
			t := reflectStorageBaseTenant.Type().Field(i)

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

type NullableStorageNetAppStorageVm struct {
	value *StorageNetAppStorageVm
	isSet bool
}

func (v NullableStorageNetAppStorageVm) Get() *StorageNetAppStorageVm {
	return v.value
}

func (v *NullableStorageNetAppStorageVm) Set(val *StorageNetAppStorageVm) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageVm) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageVm(val *StorageNetAppStorageVm) *NullableStorageNetAppStorageVm {
	return &NullableStorageNetAppStorageVm{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
