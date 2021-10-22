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

// HyperflexSummary The storage summary of the HyperFlex cluster.
type HyperflexSummary struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of nodes currently participating in the storage cluster.
	ActiveNodes *string `json:"ActiveNodes,omitempty"`
	// The data IP address of the HyperFlex cluster.
	Address *string `json:"Address,omitempty"`
	// The time taken during last cluster startup in seconds.
	Boottime *int64 `json:"Boottime,omitempty"`
	// The cluster access policy for the HyperFlex cluster. An access policy of 'STRICT' means that the cluster becomes readonly once any fragment of data is reduced to one copy. 'LENIENT' means that the cluster stays in read-write mode even if any fragment of data is reduced to one copy.
	ClusterAccessPolicy *string `json:"ClusterAccessPolicy,omitempty"`
	// The percentage of storage space saved using data compression.
	CompressionSavings *float64 `json:"CompressionSavings,omitempty"`
	// The compliance with the data replication factor set for the HyperFlex cluster.
	DataReplicationCompliance *string `json:"DataReplicationCompliance,omitempty"`
	// The number of data copies retained by the HyperFlex cluster.
	DataReplicationFactor *string `json:"DataReplicationFactor,omitempty"`
	// The percentage of storage space saved using data deduplication.
	DeduplicationSavings *float64 `json:"DeduplicationSavings,omitempty"`
	// The amount of time the HyperFlex cluster has been offline.
	Downtime *string `json:"Downtime,omitempty"`
	// The amount of storage capacity currently not in use, represented in bytes.
	FreeCapacity *int64                                        `json:"FreeCapacity,omitempty"`
	HealingInfo  NullableHyperflexStPlatformClusterHealingInfo `json:"HealingInfo,omitempty"`
	// The name of the HyperFlex cluster.
	Name *string `json:"Name,omitempty"`
	// The details about the resiliency health of the cluster. Includes information about the cluster healing status and the storage cluster health.
	ResiliencyDetails interface{} `json:"ResiliencyDetails,omitempty"`
	// The number of elements in the resiliency details property.
	ResiliencyDetailsSize *int64                                           `json:"ResiliencyDetailsSize,omitempty"`
	ResiliencyInfo        NullableHyperflexStPlatformClusterResiliencyInfo `json:"ResiliencyInfo,omitempty"`
	// The space utilization status of the HyperFlex cluster.
	SpaceStatus *string `json:"SpaceStatus,omitempty"`
	// The operational state of the HyperFlex cluster.
	State *string `json:"State,omitempty"`
	// The total amount of storage capacity available for the HyperFlex cluster, represented in bytes.
	TotalCapacity *int64 `json:"TotalCapacity,omitempty"`
	// The percentage of storage space saved in total.
	TotalSavings *float64 `json:"TotalSavings,omitempty"`
	// The amount of time the HyperFlex cluster has been running since last startup.
	Uptime *string `json:"Uptime,omitempty"`
	// The amount of storage capacity in use, represented in bytes.
	UsedCapacity         *int64 `json:"UsedCapacity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSummary HyperflexSummary

// NewHyperflexSummary instantiates a new HyperflexSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSummary(classId string, objectType string) *HyperflexSummary {
	this := HyperflexSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSummaryWithDefaults instantiates a new HyperflexSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSummaryWithDefaults() *HyperflexSummary {
	this := HyperflexSummary{}
	var classId string = "hyperflex.Summary"
	this.ClassId = classId
	var objectType string = "hyperflex.Summary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveNodes returns the ActiveNodes field value if set, zero value otherwise.
func (o *HyperflexSummary) GetActiveNodes() string {
	if o == nil || o.ActiveNodes == nil {
		var ret string
		return ret
	}
	return *o.ActiveNodes
}

// GetActiveNodesOk returns a tuple with the ActiveNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetActiveNodesOk() (*string, bool) {
	if o == nil || o.ActiveNodes == nil {
		return nil, false
	}
	return o.ActiveNodes, true
}

// HasActiveNodes returns a boolean if a field has been set.
func (o *HyperflexSummary) HasActiveNodes() bool {
	if o != nil && o.ActiveNodes != nil {
		return true
	}

	return false
}

// SetActiveNodes gets a reference to the given string and assigns it to the ActiveNodes field.
func (o *HyperflexSummary) SetActiveNodes(v string) {
	o.ActiveNodes = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HyperflexSummary) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HyperflexSummary) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HyperflexSummary) SetAddress(v string) {
	o.Address = &v
}

// GetBoottime returns the Boottime field value if set, zero value otherwise.
func (o *HyperflexSummary) GetBoottime() int64 {
	if o == nil || o.Boottime == nil {
		var ret int64
		return ret
	}
	return *o.Boottime
}

// GetBoottimeOk returns a tuple with the Boottime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetBoottimeOk() (*int64, bool) {
	if o == nil || o.Boottime == nil {
		return nil, false
	}
	return o.Boottime, true
}

// HasBoottime returns a boolean if a field has been set.
func (o *HyperflexSummary) HasBoottime() bool {
	if o != nil && o.Boottime != nil {
		return true
	}

	return false
}

// SetBoottime gets a reference to the given int64 and assigns it to the Boottime field.
func (o *HyperflexSummary) SetBoottime(v int64) {
	o.Boottime = &v
}

// GetClusterAccessPolicy returns the ClusterAccessPolicy field value if set, zero value otherwise.
func (o *HyperflexSummary) GetClusterAccessPolicy() string {
	if o == nil || o.ClusterAccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.ClusterAccessPolicy
}

// GetClusterAccessPolicyOk returns a tuple with the ClusterAccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetClusterAccessPolicyOk() (*string, bool) {
	if o == nil || o.ClusterAccessPolicy == nil {
		return nil, false
	}
	return o.ClusterAccessPolicy, true
}

// HasClusterAccessPolicy returns a boolean if a field has been set.
func (o *HyperflexSummary) HasClusterAccessPolicy() bool {
	if o != nil && o.ClusterAccessPolicy != nil {
		return true
	}

	return false
}

// SetClusterAccessPolicy gets a reference to the given string and assigns it to the ClusterAccessPolicy field.
func (o *HyperflexSummary) SetClusterAccessPolicy(v string) {
	o.ClusterAccessPolicy = &v
}

// GetCompressionSavings returns the CompressionSavings field value if set, zero value otherwise.
func (o *HyperflexSummary) GetCompressionSavings() float64 {
	if o == nil || o.CompressionSavings == nil {
		var ret float64
		return ret
	}
	return *o.CompressionSavings
}

// GetCompressionSavingsOk returns a tuple with the CompressionSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetCompressionSavingsOk() (*float64, bool) {
	if o == nil || o.CompressionSavings == nil {
		return nil, false
	}
	return o.CompressionSavings, true
}

// HasCompressionSavings returns a boolean if a field has been set.
func (o *HyperflexSummary) HasCompressionSavings() bool {
	if o != nil && o.CompressionSavings != nil {
		return true
	}

	return false
}

// SetCompressionSavings gets a reference to the given float64 and assigns it to the CompressionSavings field.
func (o *HyperflexSummary) SetCompressionSavings(v float64) {
	o.CompressionSavings = &v
}

// GetDataReplicationCompliance returns the DataReplicationCompliance field value if set, zero value otherwise.
func (o *HyperflexSummary) GetDataReplicationCompliance() string {
	if o == nil || o.DataReplicationCompliance == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationCompliance
}

// GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetDataReplicationComplianceOk() (*string, bool) {
	if o == nil || o.DataReplicationCompliance == nil {
		return nil, false
	}
	return o.DataReplicationCompliance, true
}

// HasDataReplicationCompliance returns a boolean if a field has been set.
func (o *HyperflexSummary) HasDataReplicationCompliance() bool {
	if o != nil && o.DataReplicationCompliance != nil {
		return true
	}

	return false
}

// SetDataReplicationCompliance gets a reference to the given string and assigns it to the DataReplicationCompliance field.
func (o *HyperflexSummary) SetDataReplicationCompliance(v string) {
	o.DataReplicationCompliance = &v
}

// GetDataReplicationFactor returns the DataReplicationFactor field value if set, zero value otherwise.
func (o *HyperflexSummary) GetDataReplicationFactor() string {
	if o == nil || o.DataReplicationFactor == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationFactor
}

// GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetDataReplicationFactorOk() (*string, bool) {
	if o == nil || o.DataReplicationFactor == nil {
		return nil, false
	}
	return o.DataReplicationFactor, true
}

// HasDataReplicationFactor returns a boolean if a field has been set.
func (o *HyperflexSummary) HasDataReplicationFactor() bool {
	if o != nil && o.DataReplicationFactor != nil {
		return true
	}

	return false
}

// SetDataReplicationFactor gets a reference to the given string and assigns it to the DataReplicationFactor field.
func (o *HyperflexSummary) SetDataReplicationFactor(v string) {
	o.DataReplicationFactor = &v
}

// GetDeduplicationSavings returns the DeduplicationSavings field value if set, zero value otherwise.
func (o *HyperflexSummary) GetDeduplicationSavings() float64 {
	if o == nil || o.DeduplicationSavings == nil {
		var ret float64
		return ret
	}
	return *o.DeduplicationSavings
}

// GetDeduplicationSavingsOk returns a tuple with the DeduplicationSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetDeduplicationSavingsOk() (*float64, bool) {
	if o == nil || o.DeduplicationSavings == nil {
		return nil, false
	}
	return o.DeduplicationSavings, true
}

// HasDeduplicationSavings returns a boolean if a field has been set.
func (o *HyperflexSummary) HasDeduplicationSavings() bool {
	if o != nil && o.DeduplicationSavings != nil {
		return true
	}

	return false
}

// SetDeduplicationSavings gets a reference to the given float64 and assigns it to the DeduplicationSavings field.
func (o *HyperflexSummary) SetDeduplicationSavings(v float64) {
	o.DeduplicationSavings = &v
}

// GetDowntime returns the Downtime field value if set, zero value otherwise.
func (o *HyperflexSummary) GetDowntime() string {
	if o == nil || o.Downtime == nil {
		var ret string
		return ret
	}
	return *o.Downtime
}

// GetDowntimeOk returns a tuple with the Downtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetDowntimeOk() (*string, bool) {
	if o == nil || o.Downtime == nil {
		return nil, false
	}
	return o.Downtime, true
}

// HasDowntime returns a boolean if a field has been set.
func (o *HyperflexSummary) HasDowntime() bool {
	if o != nil && o.Downtime != nil {
		return true
	}

	return false
}

// SetDowntime gets a reference to the given string and assigns it to the Downtime field.
func (o *HyperflexSummary) SetDowntime(v string) {
	o.Downtime = &v
}

// GetFreeCapacity returns the FreeCapacity field value if set, zero value otherwise.
func (o *HyperflexSummary) GetFreeCapacity() int64 {
	if o == nil || o.FreeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.FreeCapacity
}

// GetFreeCapacityOk returns a tuple with the FreeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetFreeCapacityOk() (*int64, bool) {
	if o == nil || o.FreeCapacity == nil {
		return nil, false
	}
	return o.FreeCapacity, true
}

// HasFreeCapacity returns a boolean if a field has been set.
func (o *HyperflexSummary) HasFreeCapacity() bool {
	if o != nil && o.FreeCapacity != nil {
		return true
	}

	return false
}

// SetFreeCapacity gets a reference to the given int64 and assigns it to the FreeCapacity field.
func (o *HyperflexSummary) SetFreeCapacity(v int64) {
	o.FreeCapacity = &v
}

// GetHealingInfo returns the HealingInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummary) GetHealingInfo() HyperflexStPlatformClusterHealingInfo {
	if o == nil || o.HealingInfo.Get() == nil {
		var ret HyperflexStPlatformClusterHealingInfo
		return ret
	}
	return *o.HealingInfo.Get()
}

// GetHealingInfoOk returns a tuple with the HealingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummary) GetHealingInfoOk() (*HyperflexStPlatformClusterHealingInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealingInfo.Get(), o.HealingInfo.IsSet()
}

// HasHealingInfo returns a boolean if a field has been set.
func (o *HyperflexSummary) HasHealingInfo() bool {
	if o != nil && o.HealingInfo.IsSet() {
		return true
	}

	return false
}

// SetHealingInfo gets a reference to the given NullableHyperflexStPlatformClusterHealingInfo and assigns it to the HealingInfo field.
func (o *HyperflexSummary) SetHealingInfo(v HyperflexStPlatformClusterHealingInfo) {
	o.HealingInfo.Set(&v)
}

// SetHealingInfoNil sets the value for HealingInfo to be an explicit nil
func (o *HyperflexSummary) SetHealingInfoNil() {
	o.HealingInfo.Set(nil)
}

// UnsetHealingInfo ensures that no value is present for HealingInfo, not even an explicit nil
func (o *HyperflexSummary) UnsetHealingInfo() {
	o.HealingInfo.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexSummary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexSummary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexSummary) SetName(v string) {
	o.Name = &v
}

// GetResiliencyDetails returns the ResiliencyDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummary) GetResiliencyDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResiliencyDetails
}

// GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummary) GetResiliencyDetailsOk() (*interface{}, bool) {
	if o == nil || o.ResiliencyDetails == nil {
		return nil, false
	}
	return &o.ResiliencyDetails, true
}

// HasResiliencyDetails returns a boolean if a field has been set.
func (o *HyperflexSummary) HasResiliencyDetails() bool {
	if o != nil && o.ResiliencyDetails != nil {
		return true
	}

	return false
}

// SetResiliencyDetails gets a reference to the given interface{} and assigns it to the ResiliencyDetails field.
func (o *HyperflexSummary) SetResiliencyDetails(v interface{}) {
	o.ResiliencyDetails = v
}

// GetResiliencyDetailsSize returns the ResiliencyDetailsSize field value if set, zero value otherwise.
func (o *HyperflexSummary) GetResiliencyDetailsSize() int64 {
	if o == nil || o.ResiliencyDetailsSize == nil {
		var ret int64
		return ret
	}
	return *o.ResiliencyDetailsSize
}

// GetResiliencyDetailsSizeOk returns a tuple with the ResiliencyDetailsSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetResiliencyDetailsSizeOk() (*int64, bool) {
	if o == nil || o.ResiliencyDetailsSize == nil {
		return nil, false
	}
	return o.ResiliencyDetailsSize, true
}

// HasResiliencyDetailsSize returns a boolean if a field has been set.
func (o *HyperflexSummary) HasResiliencyDetailsSize() bool {
	if o != nil && o.ResiliencyDetailsSize != nil {
		return true
	}

	return false
}

// SetResiliencyDetailsSize gets a reference to the given int64 and assigns it to the ResiliencyDetailsSize field.
func (o *HyperflexSummary) SetResiliencyDetailsSize(v int64) {
	o.ResiliencyDetailsSize = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummary) GetResiliencyInfo() HyperflexStPlatformClusterResiliencyInfo {
	if o == nil || o.ResiliencyInfo.Get() == nil {
		var ret HyperflexStPlatformClusterResiliencyInfo
		return ret
	}
	return *o.ResiliencyInfo.Get()
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummary) GetResiliencyInfoOk() (*HyperflexStPlatformClusterResiliencyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResiliencyInfo.Get(), o.ResiliencyInfo.IsSet()
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexSummary) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo.IsSet() {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given NullableHyperflexStPlatformClusterResiliencyInfo and assigns it to the ResiliencyInfo field.
func (o *HyperflexSummary) SetResiliencyInfo(v HyperflexStPlatformClusterResiliencyInfo) {
	o.ResiliencyInfo.Set(&v)
}

// SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil
func (o *HyperflexSummary) SetResiliencyInfoNil() {
	o.ResiliencyInfo.Set(nil)
}

// UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil
func (o *HyperflexSummary) UnsetResiliencyInfo() {
	o.ResiliencyInfo.Unset()
}

// GetSpaceStatus returns the SpaceStatus field value if set, zero value otherwise.
func (o *HyperflexSummary) GetSpaceStatus() string {
	if o == nil || o.SpaceStatus == nil {
		var ret string
		return ret
	}
	return *o.SpaceStatus
}

// GetSpaceStatusOk returns a tuple with the SpaceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetSpaceStatusOk() (*string, bool) {
	if o == nil || o.SpaceStatus == nil {
		return nil, false
	}
	return o.SpaceStatus, true
}

// HasSpaceStatus returns a boolean if a field has been set.
func (o *HyperflexSummary) HasSpaceStatus() bool {
	if o != nil && o.SpaceStatus != nil {
		return true
	}

	return false
}

// SetSpaceStatus gets a reference to the given string and assigns it to the SpaceStatus field.
func (o *HyperflexSummary) SetSpaceStatus(v string) {
	o.SpaceStatus = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexSummary) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexSummary) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexSummary) SetState(v string) {
	o.State = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *HyperflexSummary) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetTotalCapacityOk() (*int64, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *HyperflexSummary) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given int64 and assigns it to the TotalCapacity field.
func (o *HyperflexSummary) SetTotalCapacity(v int64) {
	o.TotalCapacity = &v
}

// GetTotalSavings returns the TotalSavings field value if set, zero value otherwise.
func (o *HyperflexSummary) GetTotalSavings() float64 {
	if o == nil || o.TotalSavings == nil {
		var ret float64
		return ret
	}
	return *o.TotalSavings
}

// GetTotalSavingsOk returns a tuple with the TotalSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetTotalSavingsOk() (*float64, bool) {
	if o == nil || o.TotalSavings == nil {
		return nil, false
	}
	return o.TotalSavings, true
}

// HasTotalSavings returns a boolean if a field has been set.
func (o *HyperflexSummary) HasTotalSavings() bool {
	if o != nil && o.TotalSavings != nil {
		return true
	}

	return false
}

// SetTotalSavings gets a reference to the given float64 and assigns it to the TotalSavings field.
func (o *HyperflexSummary) SetTotalSavings(v float64) {
	o.TotalSavings = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *HyperflexSummary) GetUptime() string {
	if o == nil || o.Uptime == nil {
		var ret string
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetUptimeOk() (*string, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *HyperflexSummary) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given string and assigns it to the Uptime field.
func (o *HyperflexSummary) SetUptime(v string) {
	o.Uptime = &v
}

// GetUsedCapacity returns the UsedCapacity field value if set, zero value otherwise.
func (o *HyperflexSummary) GetUsedCapacity() int64 {
	if o == nil || o.UsedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.UsedCapacity
}

// GetUsedCapacityOk returns a tuple with the UsedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummary) GetUsedCapacityOk() (*int64, bool) {
	if o == nil || o.UsedCapacity == nil {
		return nil, false
	}
	return o.UsedCapacity, true
}

// HasUsedCapacity returns a boolean if a field has been set.
func (o *HyperflexSummary) HasUsedCapacity() bool {
	if o != nil && o.UsedCapacity != nil {
		return true
	}

	return false
}

// SetUsedCapacity gets a reference to the given int64 and assigns it to the UsedCapacity field.
func (o *HyperflexSummary) SetUsedCapacity(v int64) {
	o.UsedCapacity = &v
}

func (o HyperflexSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActiveNodes != nil {
		toSerialize["ActiveNodes"] = o.ActiveNodes
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Boottime != nil {
		toSerialize["Boottime"] = o.Boottime
	}
	if o.ClusterAccessPolicy != nil {
		toSerialize["ClusterAccessPolicy"] = o.ClusterAccessPolicy
	}
	if o.CompressionSavings != nil {
		toSerialize["CompressionSavings"] = o.CompressionSavings
	}
	if o.DataReplicationCompliance != nil {
		toSerialize["DataReplicationCompliance"] = o.DataReplicationCompliance
	}
	if o.DataReplicationFactor != nil {
		toSerialize["DataReplicationFactor"] = o.DataReplicationFactor
	}
	if o.DeduplicationSavings != nil {
		toSerialize["DeduplicationSavings"] = o.DeduplicationSavings
	}
	if o.Downtime != nil {
		toSerialize["Downtime"] = o.Downtime
	}
	if o.FreeCapacity != nil {
		toSerialize["FreeCapacity"] = o.FreeCapacity
	}
	if o.HealingInfo.IsSet() {
		toSerialize["HealingInfo"] = o.HealingInfo.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ResiliencyDetails != nil {
		toSerialize["ResiliencyDetails"] = o.ResiliencyDetails
	}
	if o.ResiliencyDetailsSize != nil {
		toSerialize["ResiliencyDetailsSize"] = o.ResiliencyDetailsSize
	}
	if o.ResiliencyInfo.IsSet() {
		toSerialize["ResiliencyInfo"] = o.ResiliencyInfo.Get()
	}
	if o.SpaceStatus != nil {
		toSerialize["SpaceStatus"] = o.SpaceStatus
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.TotalSavings != nil {
		toSerialize["TotalSavings"] = o.TotalSavings
	}
	if o.Uptime != nil {
		toSerialize["Uptime"] = o.Uptime
	}
	if o.UsedCapacity != nil {
		toSerialize["UsedCapacity"] = o.UsedCapacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSummary) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of nodes currently participating in the storage cluster.
		ActiveNodes *string `json:"ActiveNodes,omitempty"`
		// The data IP address of the HyperFlex cluster.
		Address *string `json:"Address,omitempty"`
		// The time taken during last cluster startup in seconds.
		Boottime *int64 `json:"Boottime,omitempty"`
		// The cluster access policy for the HyperFlex cluster. An access policy of 'STRICT' means that the cluster becomes readonly once any fragment of data is reduced to one copy. 'LENIENT' means that the cluster stays in read-write mode even if any fragment of data is reduced to one copy.
		ClusterAccessPolicy *string `json:"ClusterAccessPolicy,omitempty"`
		// The percentage of storage space saved using data compression.
		CompressionSavings *float64 `json:"CompressionSavings,omitempty"`
		// The compliance with the data replication factor set for the HyperFlex cluster.
		DataReplicationCompliance *string `json:"DataReplicationCompliance,omitempty"`
		// The number of data copies retained by the HyperFlex cluster.
		DataReplicationFactor *string `json:"DataReplicationFactor,omitempty"`
		// The percentage of storage space saved using data deduplication.
		DeduplicationSavings *float64 `json:"DeduplicationSavings,omitempty"`
		// The amount of time the HyperFlex cluster has been offline.
		Downtime *string `json:"Downtime,omitempty"`
		// The amount of storage capacity currently not in use, represented in bytes.
		FreeCapacity *int64                                        `json:"FreeCapacity,omitempty"`
		HealingInfo  NullableHyperflexStPlatformClusterHealingInfo `json:"HealingInfo,omitempty"`
		// The name of the HyperFlex cluster.
		Name *string `json:"Name,omitempty"`
		// The details about the resiliency health of the cluster. Includes information about the cluster healing status and the storage cluster health.
		ResiliencyDetails interface{} `json:"ResiliencyDetails,omitempty"`
		// The number of elements in the resiliency details property.
		ResiliencyDetailsSize *int64                                           `json:"ResiliencyDetailsSize,omitempty"`
		ResiliencyInfo        NullableHyperflexStPlatformClusterResiliencyInfo `json:"ResiliencyInfo,omitempty"`
		// The space utilization status of the HyperFlex cluster.
		SpaceStatus *string `json:"SpaceStatus,omitempty"`
		// The operational state of the HyperFlex cluster.
		State *string `json:"State,omitempty"`
		// The total amount of storage capacity available for the HyperFlex cluster, represented in bytes.
		TotalCapacity *int64 `json:"TotalCapacity,omitempty"`
		// The percentage of storage space saved in total.
		TotalSavings *float64 `json:"TotalSavings,omitempty"`
		// The amount of time the HyperFlex cluster has been running since last startup.
		Uptime *string `json:"Uptime,omitempty"`
		// The amount of storage capacity in use, represented in bytes.
		UsedCapacity *int64 `json:"UsedCapacity,omitempty"`
	}

	varHyperflexSummaryWithoutEmbeddedStruct := HyperflexSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSummaryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSummary := _HyperflexSummary{}
		varHyperflexSummary.ClassId = varHyperflexSummaryWithoutEmbeddedStruct.ClassId
		varHyperflexSummary.ObjectType = varHyperflexSummaryWithoutEmbeddedStruct.ObjectType
		varHyperflexSummary.ActiveNodes = varHyperflexSummaryWithoutEmbeddedStruct.ActiveNodes
		varHyperflexSummary.Address = varHyperflexSummaryWithoutEmbeddedStruct.Address
		varHyperflexSummary.Boottime = varHyperflexSummaryWithoutEmbeddedStruct.Boottime
		varHyperflexSummary.ClusterAccessPolicy = varHyperflexSummaryWithoutEmbeddedStruct.ClusterAccessPolicy
		varHyperflexSummary.CompressionSavings = varHyperflexSummaryWithoutEmbeddedStruct.CompressionSavings
		varHyperflexSummary.DataReplicationCompliance = varHyperflexSummaryWithoutEmbeddedStruct.DataReplicationCompliance
		varHyperflexSummary.DataReplicationFactor = varHyperflexSummaryWithoutEmbeddedStruct.DataReplicationFactor
		varHyperflexSummary.DeduplicationSavings = varHyperflexSummaryWithoutEmbeddedStruct.DeduplicationSavings
		varHyperflexSummary.Downtime = varHyperflexSummaryWithoutEmbeddedStruct.Downtime
		varHyperflexSummary.FreeCapacity = varHyperflexSummaryWithoutEmbeddedStruct.FreeCapacity
		varHyperflexSummary.HealingInfo = varHyperflexSummaryWithoutEmbeddedStruct.HealingInfo
		varHyperflexSummary.Name = varHyperflexSummaryWithoutEmbeddedStruct.Name
		varHyperflexSummary.ResiliencyDetails = varHyperflexSummaryWithoutEmbeddedStruct.ResiliencyDetails
		varHyperflexSummary.ResiliencyDetailsSize = varHyperflexSummaryWithoutEmbeddedStruct.ResiliencyDetailsSize
		varHyperflexSummary.ResiliencyInfo = varHyperflexSummaryWithoutEmbeddedStruct.ResiliencyInfo
		varHyperflexSummary.SpaceStatus = varHyperflexSummaryWithoutEmbeddedStruct.SpaceStatus
		varHyperflexSummary.State = varHyperflexSummaryWithoutEmbeddedStruct.State
		varHyperflexSummary.TotalCapacity = varHyperflexSummaryWithoutEmbeddedStruct.TotalCapacity
		varHyperflexSummary.TotalSavings = varHyperflexSummaryWithoutEmbeddedStruct.TotalSavings
		varHyperflexSummary.Uptime = varHyperflexSummaryWithoutEmbeddedStruct.Uptime
		varHyperflexSummary.UsedCapacity = varHyperflexSummaryWithoutEmbeddedStruct.UsedCapacity
		*o = HyperflexSummary(varHyperflexSummary)
	} else {
		return err
	}

	varHyperflexSummary := _HyperflexSummary{}

	err = json.Unmarshal(bytes, &varHyperflexSummary)
	if err == nil {
		o.MoBaseComplexType = varHyperflexSummary.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveNodes")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Boottime")
		delete(additionalProperties, "ClusterAccessPolicy")
		delete(additionalProperties, "CompressionSavings")
		delete(additionalProperties, "DataReplicationCompliance")
		delete(additionalProperties, "DataReplicationFactor")
		delete(additionalProperties, "DeduplicationSavings")
		delete(additionalProperties, "Downtime")
		delete(additionalProperties, "FreeCapacity")
		delete(additionalProperties, "HealingInfo")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ResiliencyDetails")
		delete(additionalProperties, "ResiliencyDetailsSize")
		delete(additionalProperties, "ResiliencyInfo")
		delete(additionalProperties, "SpaceStatus")
		delete(additionalProperties, "State")
		delete(additionalProperties, "TotalCapacity")
		delete(additionalProperties, "TotalSavings")
		delete(additionalProperties, "Uptime")
		delete(additionalProperties, "UsedCapacity")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHyperflexSummary struct {
	value *HyperflexSummary
	isSet bool
}

func (v NullableHyperflexSummary) Get() *HyperflexSummary {
	return v.value
}

func (v *NullableHyperflexSummary) Set(val *HyperflexSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSummary(val *HyperflexSummary) *NullableHyperflexSummary {
	return &NullableHyperflexSummary{value: val, isSet: true}
}

func (v NullableHyperflexSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
