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
)

// HyperflexSummaryAllOf Definition of the list of properties defined in 'hyperflex.Summary', excluding properties defined in parent classes.
type HyperflexSummaryAllOf struct {
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

type _HyperflexSummaryAllOf HyperflexSummaryAllOf

// NewHyperflexSummaryAllOf instantiates a new HyperflexSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSummaryAllOf(classId string, objectType string) *HyperflexSummaryAllOf {
	this := HyperflexSummaryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSummaryAllOfWithDefaults instantiates a new HyperflexSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSummaryAllOfWithDefaults() *HyperflexSummaryAllOf {
	this := HyperflexSummaryAllOf{}
	var classId string = "hyperflex.Summary"
	this.ClassId = classId
	var objectType string = "hyperflex.Summary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSummaryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSummaryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSummaryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSummaryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveNodes returns the ActiveNodes field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetActiveNodes() string {
	if o == nil || o.ActiveNodes == nil {
		var ret string
		return ret
	}
	return *o.ActiveNodes
}

// GetActiveNodesOk returns a tuple with the ActiveNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetActiveNodesOk() (*string, bool) {
	if o == nil || o.ActiveNodes == nil {
		return nil, false
	}
	return o.ActiveNodes, true
}

// HasActiveNodes returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasActiveNodes() bool {
	if o != nil && o.ActiveNodes != nil {
		return true
	}

	return false
}

// SetActiveNodes gets a reference to the given string and assigns it to the ActiveNodes field.
func (o *HyperflexSummaryAllOf) SetActiveNodes(v string) {
	o.ActiveNodes = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HyperflexSummaryAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetBoottime returns the Boottime field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetBoottime() int64 {
	if o == nil || o.Boottime == nil {
		var ret int64
		return ret
	}
	return *o.Boottime
}

// GetBoottimeOk returns a tuple with the Boottime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetBoottimeOk() (*int64, bool) {
	if o == nil || o.Boottime == nil {
		return nil, false
	}
	return o.Boottime, true
}

// HasBoottime returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasBoottime() bool {
	if o != nil && o.Boottime != nil {
		return true
	}

	return false
}

// SetBoottime gets a reference to the given int64 and assigns it to the Boottime field.
func (o *HyperflexSummaryAllOf) SetBoottime(v int64) {
	o.Boottime = &v
}

// GetClusterAccessPolicy returns the ClusterAccessPolicy field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetClusterAccessPolicy() string {
	if o == nil || o.ClusterAccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.ClusterAccessPolicy
}

// GetClusterAccessPolicyOk returns a tuple with the ClusterAccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetClusterAccessPolicyOk() (*string, bool) {
	if o == nil || o.ClusterAccessPolicy == nil {
		return nil, false
	}
	return o.ClusterAccessPolicy, true
}

// HasClusterAccessPolicy returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasClusterAccessPolicy() bool {
	if o != nil && o.ClusterAccessPolicy != nil {
		return true
	}

	return false
}

// SetClusterAccessPolicy gets a reference to the given string and assigns it to the ClusterAccessPolicy field.
func (o *HyperflexSummaryAllOf) SetClusterAccessPolicy(v string) {
	o.ClusterAccessPolicy = &v
}

// GetCompressionSavings returns the CompressionSavings field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetCompressionSavings() float64 {
	if o == nil || o.CompressionSavings == nil {
		var ret float64
		return ret
	}
	return *o.CompressionSavings
}

// GetCompressionSavingsOk returns a tuple with the CompressionSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetCompressionSavingsOk() (*float64, bool) {
	if o == nil || o.CompressionSavings == nil {
		return nil, false
	}
	return o.CompressionSavings, true
}

// HasCompressionSavings returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasCompressionSavings() bool {
	if o != nil && o.CompressionSavings != nil {
		return true
	}

	return false
}

// SetCompressionSavings gets a reference to the given float64 and assigns it to the CompressionSavings field.
func (o *HyperflexSummaryAllOf) SetCompressionSavings(v float64) {
	o.CompressionSavings = &v
}

// GetDataReplicationCompliance returns the DataReplicationCompliance field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetDataReplicationCompliance() string {
	if o == nil || o.DataReplicationCompliance == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationCompliance
}

// GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetDataReplicationComplianceOk() (*string, bool) {
	if o == nil || o.DataReplicationCompliance == nil {
		return nil, false
	}
	return o.DataReplicationCompliance, true
}

// HasDataReplicationCompliance returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasDataReplicationCompliance() bool {
	if o != nil && o.DataReplicationCompliance != nil {
		return true
	}

	return false
}

// SetDataReplicationCompliance gets a reference to the given string and assigns it to the DataReplicationCompliance field.
func (o *HyperflexSummaryAllOf) SetDataReplicationCompliance(v string) {
	o.DataReplicationCompliance = &v
}

// GetDataReplicationFactor returns the DataReplicationFactor field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetDataReplicationFactor() string {
	if o == nil || o.DataReplicationFactor == nil {
		var ret string
		return ret
	}
	return *o.DataReplicationFactor
}

// GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetDataReplicationFactorOk() (*string, bool) {
	if o == nil || o.DataReplicationFactor == nil {
		return nil, false
	}
	return o.DataReplicationFactor, true
}

// HasDataReplicationFactor returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasDataReplicationFactor() bool {
	if o != nil && o.DataReplicationFactor != nil {
		return true
	}

	return false
}

// SetDataReplicationFactor gets a reference to the given string and assigns it to the DataReplicationFactor field.
func (o *HyperflexSummaryAllOf) SetDataReplicationFactor(v string) {
	o.DataReplicationFactor = &v
}

// GetDeduplicationSavings returns the DeduplicationSavings field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetDeduplicationSavings() float64 {
	if o == nil || o.DeduplicationSavings == nil {
		var ret float64
		return ret
	}
	return *o.DeduplicationSavings
}

// GetDeduplicationSavingsOk returns a tuple with the DeduplicationSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetDeduplicationSavingsOk() (*float64, bool) {
	if o == nil || o.DeduplicationSavings == nil {
		return nil, false
	}
	return o.DeduplicationSavings, true
}

// HasDeduplicationSavings returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasDeduplicationSavings() bool {
	if o != nil && o.DeduplicationSavings != nil {
		return true
	}

	return false
}

// SetDeduplicationSavings gets a reference to the given float64 and assigns it to the DeduplicationSavings field.
func (o *HyperflexSummaryAllOf) SetDeduplicationSavings(v float64) {
	o.DeduplicationSavings = &v
}

// GetDowntime returns the Downtime field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetDowntime() string {
	if o == nil || o.Downtime == nil {
		var ret string
		return ret
	}
	return *o.Downtime
}

// GetDowntimeOk returns a tuple with the Downtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetDowntimeOk() (*string, bool) {
	if o == nil || o.Downtime == nil {
		return nil, false
	}
	return o.Downtime, true
}

// HasDowntime returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasDowntime() bool {
	if o != nil && o.Downtime != nil {
		return true
	}

	return false
}

// SetDowntime gets a reference to the given string and assigns it to the Downtime field.
func (o *HyperflexSummaryAllOf) SetDowntime(v string) {
	o.Downtime = &v
}

// GetFreeCapacity returns the FreeCapacity field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetFreeCapacity() int64 {
	if o == nil || o.FreeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.FreeCapacity
}

// GetFreeCapacityOk returns a tuple with the FreeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetFreeCapacityOk() (*int64, bool) {
	if o == nil || o.FreeCapacity == nil {
		return nil, false
	}
	return o.FreeCapacity, true
}

// HasFreeCapacity returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasFreeCapacity() bool {
	if o != nil && o.FreeCapacity != nil {
		return true
	}

	return false
}

// SetFreeCapacity gets a reference to the given int64 and assigns it to the FreeCapacity field.
func (o *HyperflexSummaryAllOf) SetFreeCapacity(v int64) {
	o.FreeCapacity = &v
}

// GetHealingInfo returns the HealingInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummaryAllOf) GetHealingInfo() HyperflexStPlatformClusterHealingInfo {
	if o == nil || o.HealingInfo.Get() == nil {
		var ret HyperflexStPlatformClusterHealingInfo
		return ret
	}
	return *o.HealingInfo.Get()
}

// GetHealingInfoOk returns a tuple with the HealingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummaryAllOf) GetHealingInfoOk() (*HyperflexStPlatformClusterHealingInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealingInfo.Get(), o.HealingInfo.IsSet()
}

// HasHealingInfo returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasHealingInfo() bool {
	if o != nil && o.HealingInfo.IsSet() {
		return true
	}

	return false
}

// SetHealingInfo gets a reference to the given NullableHyperflexStPlatformClusterHealingInfo and assigns it to the HealingInfo field.
func (o *HyperflexSummaryAllOf) SetHealingInfo(v HyperflexStPlatformClusterHealingInfo) {
	o.HealingInfo.Set(&v)
}

// SetHealingInfoNil sets the value for HealingInfo to be an explicit nil
func (o *HyperflexSummaryAllOf) SetHealingInfoNil() {
	o.HealingInfo.Set(nil)
}

// UnsetHealingInfo ensures that no value is present for HealingInfo, not even an explicit nil
func (o *HyperflexSummaryAllOf) UnsetHealingInfo() {
	o.HealingInfo.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexSummaryAllOf) SetName(v string) {
	o.Name = &v
}

// GetResiliencyDetails returns the ResiliencyDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummaryAllOf) GetResiliencyDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResiliencyDetails
}

// GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummaryAllOf) GetResiliencyDetailsOk() (*interface{}, bool) {
	if o == nil || o.ResiliencyDetails == nil {
		return nil, false
	}
	return &o.ResiliencyDetails, true
}

// HasResiliencyDetails returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasResiliencyDetails() bool {
	if o != nil && o.ResiliencyDetails != nil {
		return true
	}

	return false
}

// SetResiliencyDetails gets a reference to the given interface{} and assigns it to the ResiliencyDetails field.
func (o *HyperflexSummaryAllOf) SetResiliencyDetails(v interface{}) {
	o.ResiliencyDetails = v
}

// GetResiliencyDetailsSize returns the ResiliencyDetailsSize field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetResiliencyDetailsSize() int64 {
	if o == nil || o.ResiliencyDetailsSize == nil {
		var ret int64
		return ret
	}
	return *o.ResiliencyDetailsSize
}

// GetResiliencyDetailsSizeOk returns a tuple with the ResiliencyDetailsSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetResiliencyDetailsSizeOk() (*int64, bool) {
	if o == nil || o.ResiliencyDetailsSize == nil {
		return nil, false
	}
	return o.ResiliencyDetailsSize, true
}

// HasResiliencyDetailsSize returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasResiliencyDetailsSize() bool {
	if o != nil && o.ResiliencyDetailsSize != nil {
		return true
	}

	return false
}

// SetResiliencyDetailsSize gets a reference to the given int64 and assigns it to the ResiliencyDetailsSize field.
func (o *HyperflexSummaryAllOf) SetResiliencyDetailsSize(v int64) {
	o.ResiliencyDetailsSize = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSummaryAllOf) GetResiliencyInfo() HyperflexStPlatformClusterResiliencyInfo {
	if o == nil || o.ResiliencyInfo.Get() == nil {
		var ret HyperflexStPlatformClusterResiliencyInfo
		return ret
	}
	return *o.ResiliencyInfo.Get()
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSummaryAllOf) GetResiliencyInfoOk() (*HyperflexStPlatformClusterResiliencyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResiliencyInfo.Get(), o.ResiliencyInfo.IsSet()
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo.IsSet() {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given NullableHyperflexStPlatformClusterResiliencyInfo and assigns it to the ResiliencyInfo field.
func (o *HyperflexSummaryAllOf) SetResiliencyInfo(v HyperflexStPlatformClusterResiliencyInfo) {
	o.ResiliencyInfo.Set(&v)
}

// SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil
func (o *HyperflexSummaryAllOf) SetResiliencyInfoNil() {
	o.ResiliencyInfo.Set(nil)
}

// UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil
func (o *HyperflexSummaryAllOf) UnsetResiliencyInfo() {
	o.ResiliencyInfo.Unset()
}

// GetSpaceStatus returns the SpaceStatus field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetSpaceStatus() string {
	if o == nil || o.SpaceStatus == nil {
		var ret string
		return ret
	}
	return *o.SpaceStatus
}

// GetSpaceStatusOk returns a tuple with the SpaceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetSpaceStatusOk() (*string, bool) {
	if o == nil || o.SpaceStatus == nil {
		return nil, false
	}
	return o.SpaceStatus, true
}

// HasSpaceStatus returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasSpaceStatus() bool {
	if o != nil && o.SpaceStatus != nil {
		return true
	}

	return false
}

// SetSpaceStatus gets a reference to the given string and assigns it to the SpaceStatus field.
func (o *HyperflexSummaryAllOf) SetSpaceStatus(v string) {
	o.SpaceStatus = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexSummaryAllOf) SetState(v string) {
	o.State = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetTotalCapacityOk() (*int64, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given int64 and assigns it to the TotalCapacity field.
func (o *HyperflexSummaryAllOf) SetTotalCapacity(v int64) {
	o.TotalCapacity = &v
}

// GetTotalSavings returns the TotalSavings field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetTotalSavings() float64 {
	if o == nil || o.TotalSavings == nil {
		var ret float64
		return ret
	}
	return *o.TotalSavings
}

// GetTotalSavingsOk returns a tuple with the TotalSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetTotalSavingsOk() (*float64, bool) {
	if o == nil || o.TotalSavings == nil {
		return nil, false
	}
	return o.TotalSavings, true
}

// HasTotalSavings returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasTotalSavings() bool {
	if o != nil && o.TotalSavings != nil {
		return true
	}

	return false
}

// SetTotalSavings gets a reference to the given float64 and assigns it to the TotalSavings field.
func (o *HyperflexSummaryAllOf) SetTotalSavings(v float64) {
	o.TotalSavings = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetUptime() string {
	if o == nil || o.Uptime == nil {
		var ret string
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetUptimeOk() (*string, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given string and assigns it to the Uptime field.
func (o *HyperflexSummaryAllOf) SetUptime(v string) {
	o.Uptime = &v
}

// GetUsedCapacity returns the UsedCapacity field value if set, zero value otherwise.
func (o *HyperflexSummaryAllOf) GetUsedCapacity() int64 {
	if o == nil || o.UsedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.UsedCapacity
}

// GetUsedCapacityOk returns a tuple with the UsedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSummaryAllOf) GetUsedCapacityOk() (*int64, bool) {
	if o == nil || o.UsedCapacity == nil {
		return nil, false
	}
	return o.UsedCapacity, true
}

// HasUsedCapacity returns a boolean if a field has been set.
func (o *HyperflexSummaryAllOf) HasUsedCapacity() bool {
	if o != nil && o.UsedCapacity != nil {
		return true
	}

	return false
}

// SetUsedCapacity gets a reference to the given int64 and assigns it to the UsedCapacity field.
func (o *HyperflexSummaryAllOf) SetUsedCapacity(v int64) {
	o.UsedCapacity = &v
}

func (o HyperflexSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *HyperflexSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSummaryAllOf := _HyperflexSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSummaryAllOf); err == nil {
		*o = HyperflexSummaryAllOf(varHyperflexSummaryAllOf)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSummaryAllOf struct {
	value *HyperflexSummaryAllOf
	isSet bool
}

func (v NullableHyperflexSummaryAllOf) Get() *HyperflexSummaryAllOf {
	return v.value
}

func (v *NullableHyperflexSummaryAllOf) Set(val *HyperflexSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSummaryAllOf(val *HyperflexSummaryAllOf) *NullableHyperflexSummaryAllOf {
	return &NullableHyperflexSummaryAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
