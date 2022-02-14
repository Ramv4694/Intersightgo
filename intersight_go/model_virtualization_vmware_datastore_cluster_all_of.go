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

// VirtualizationVmwareDatastoreClusterAllOf Definition of the list of properties defined in 'virtualization.VmwareDatastoreCluster', excluding properties defined in parent classes.
type VirtualizationVmwareDatastoreClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The global automation level for all virtual machines in this datastore cluster.
	AutomationLevel *string `json:"AutomationLevel,omitempty"`
	// Minimum level of free space for each datastore that is the threshold for action.
	FreeSpaceThreshold *int64 `json:"FreeSpaceThreshold,omitempty"`
	// Inventory path of the Datastore Cluster.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Minimum I/O latency for each datastore below which I/O load balancing moves are not considered.
	IoLatencyThreshold *int32 `json:"IoLatencyThreshold,omitempty"`
	// Storage DRS behavior when it generates recommendations for correcting I/O load imbalance in a datastore cluster.
	IoLoadBalanceAutomationMode *string `json:"IoLoadBalanceAutomationMode,omitempty"`
	// Amount of imbalance that Storage DRS should tolerate.
	IoLoadImbalanceThreshold *int32 `json:"IoLoadImbalanceThreshold,omitempty"`
	// Is I/O Metrics enabled for this datastore cluster.
	IoMetricsEnabled *bool `json:"IoMetricsEnabled,omitempty"`
	// Specify how much of an improvement DRS should look for before making a recommendation or performing a migration.
	MinSpaceUtilizationDifference *int32 `json:"MinSpaceUtilizationDifference,omitempty"`
	// Storage DRS behavior when it generates recommendations for correcting storage and VM policy violations in a datastore cluster.
	PolicyEnforcementAutomationMode *string `json:"PolicyEnforcementAutomationMode,omitempty"`
	// Storage DRS makes storage migration recommendations if total IOPs reservation of all VMs running on a datastore is higher than the specified threshold.
	ReservablePercentThreshold *int32 `json:"ReservablePercentThreshold,omitempty"`
	// Storage DRS behavior when it generates recommendations for correcting affinity rule violations in a datastore cluster.
	RuleEnforcementAutomationMode *string `json:"RuleEnforcementAutomationMode,omitempty"`
	// Storage DRS behavior when it generates recommendations for correcting space load imbalance in a datastore cluster.
	SpaceLoadBalanceAutomationMode *string `json:"SpaceLoadBalanceAutomationMode,omitempty"`
	// Runtime thresholds govern when Storage DRS performs or recommends migrations.
	SpaceThresholdMode *string `json:"SpaceThresholdMode,omitempty"`
	// Datastore cluster health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
	Status *string `json:"Status,omitempty"`
	// Is Storage DRS enabled for this datastore cluster.
	StorageDrsEnabled *bool `json:"StorageDrsEnabled,omitempty"`
	// Minimum level of consumed space for each datastore that is the threshold for action.
	UtilizedSpaceThreshold *int32 `json:"UtilizedSpaceThreshold,omitempty"`
	// Storage DRS behavior when it generates recommendations for VM evacuations from datastores in a datastore cluster.
	VmEvacuationAutomationMode *string                                     `json:"VmEvacuationAutomationMode,omitempty"`
	Datacenter                 *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _VirtualizationVmwareDatastoreClusterAllOf VirtualizationVmwareDatastoreClusterAllOf

// NewVirtualizationVmwareDatastoreClusterAllOf instantiates a new VirtualizationVmwareDatastoreClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatastoreClusterAllOf(classId string, objectType string) *VirtualizationVmwareDatastoreClusterAllOf {
	this := VirtualizationVmwareDatastoreClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationVmwareDatastoreClusterAllOfWithDefaults instantiates a new VirtualizationVmwareDatastoreClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatastoreClusterAllOfWithDefaults() *VirtualizationVmwareDatastoreClusterAllOf {
	this := VirtualizationVmwareDatastoreClusterAllOf{}
	var classId string = "virtualization.VmwareDatastoreCluster"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatastoreCluster"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutomationLevel returns the AutomationLevel field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetAutomationLevel() string {
	if o == nil || o.AutomationLevel == nil {
		var ret string
		return ret
	}
	return *o.AutomationLevel
}

// GetAutomationLevelOk returns a tuple with the AutomationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetAutomationLevelOk() (*string, bool) {
	if o == nil || o.AutomationLevel == nil {
		return nil, false
	}
	return o.AutomationLevel, true
}

// HasAutomationLevel returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasAutomationLevel() bool {
	if o != nil && o.AutomationLevel != nil {
		return true
	}

	return false
}

// SetAutomationLevel gets a reference to the given string and assigns it to the AutomationLevel field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetAutomationLevel(v string) {
	o.AutomationLevel = &v
}

// GetFreeSpaceThreshold returns the FreeSpaceThreshold field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetFreeSpaceThreshold() int64 {
	if o == nil || o.FreeSpaceThreshold == nil {
		var ret int64
		return ret
	}
	return *o.FreeSpaceThreshold
}

// GetFreeSpaceThresholdOk returns a tuple with the FreeSpaceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetFreeSpaceThresholdOk() (*int64, bool) {
	if o == nil || o.FreeSpaceThreshold == nil {
		return nil, false
	}
	return o.FreeSpaceThreshold, true
}

// HasFreeSpaceThreshold returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasFreeSpaceThreshold() bool {
	if o != nil && o.FreeSpaceThreshold != nil {
		return true
	}

	return false
}

// SetFreeSpaceThreshold gets a reference to the given int64 and assigns it to the FreeSpaceThreshold field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetFreeSpaceThreshold(v int64) {
	o.FreeSpaceThreshold = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetIoLatencyThreshold returns the IoLatencyThreshold field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLatencyThreshold() int32 {
	if o == nil || o.IoLatencyThreshold == nil {
		var ret int32
		return ret
	}
	return *o.IoLatencyThreshold
}

// GetIoLatencyThresholdOk returns a tuple with the IoLatencyThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLatencyThresholdOk() (*int32, bool) {
	if o == nil || o.IoLatencyThreshold == nil {
		return nil, false
	}
	return o.IoLatencyThreshold, true
}

// HasIoLatencyThreshold returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasIoLatencyThreshold() bool {
	if o != nil && o.IoLatencyThreshold != nil {
		return true
	}

	return false
}

// SetIoLatencyThreshold gets a reference to the given int32 and assigns it to the IoLatencyThreshold field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetIoLatencyThreshold(v int32) {
	o.IoLatencyThreshold = &v
}

// GetIoLoadBalanceAutomationMode returns the IoLoadBalanceAutomationMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLoadBalanceAutomationMode() string {
	if o == nil || o.IoLoadBalanceAutomationMode == nil {
		var ret string
		return ret
	}
	return *o.IoLoadBalanceAutomationMode
}

// GetIoLoadBalanceAutomationModeOk returns a tuple with the IoLoadBalanceAutomationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLoadBalanceAutomationModeOk() (*string, bool) {
	if o == nil || o.IoLoadBalanceAutomationMode == nil {
		return nil, false
	}
	return o.IoLoadBalanceAutomationMode, true
}

// HasIoLoadBalanceAutomationMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasIoLoadBalanceAutomationMode() bool {
	if o != nil && o.IoLoadBalanceAutomationMode != nil {
		return true
	}

	return false
}

// SetIoLoadBalanceAutomationMode gets a reference to the given string and assigns it to the IoLoadBalanceAutomationMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetIoLoadBalanceAutomationMode(v string) {
	o.IoLoadBalanceAutomationMode = &v
}

// GetIoLoadImbalanceThreshold returns the IoLoadImbalanceThreshold field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLoadImbalanceThreshold() int32 {
	if o == nil || o.IoLoadImbalanceThreshold == nil {
		var ret int32
		return ret
	}
	return *o.IoLoadImbalanceThreshold
}

// GetIoLoadImbalanceThresholdOk returns a tuple with the IoLoadImbalanceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoLoadImbalanceThresholdOk() (*int32, bool) {
	if o == nil || o.IoLoadImbalanceThreshold == nil {
		return nil, false
	}
	return o.IoLoadImbalanceThreshold, true
}

// HasIoLoadImbalanceThreshold returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasIoLoadImbalanceThreshold() bool {
	if o != nil && o.IoLoadImbalanceThreshold != nil {
		return true
	}

	return false
}

// SetIoLoadImbalanceThreshold gets a reference to the given int32 and assigns it to the IoLoadImbalanceThreshold field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetIoLoadImbalanceThreshold(v int32) {
	o.IoLoadImbalanceThreshold = &v
}

// GetIoMetricsEnabled returns the IoMetricsEnabled field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoMetricsEnabled() bool {
	if o == nil || o.IoMetricsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IoMetricsEnabled
}

// GetIoMetricsEnabledOk returns a tuple with the IoMetricsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetIoMetricsEnabledOk() (*bool, bool) {
	if o == nil || o.IoMetricsEnabled == nil {
		return nil, false
	}
	return o.IoMetricsEnabled, true
}

// HasIoMetricsEnabled returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasIoMetricsEnabled() bool {
	if o != nil && o.IoMetricsEnabled != nil {
		return true
	}

	return false
}

// SetIoMetricsEnabled gets a reference to the given bool and assigns it to the IoMetricsEnabled field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetIoMetricsEnabled(v bool) {
	o.IoMetricsEnabled = &v
}

// GetMinSpaceUtilizationDifference returns the MinSpaceUtilizationDifference field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetMinSpaceUtilizationDifference() int32 {
	if o == nil || o.MinSpaceUtilizationDifference == nil {
		var ret int32
		return ret
	}
	return *o.MinSpaceUtilizationDifference
}

// GetMinSpaceUtilizationDifferenceOk returns a tuple with the MinSpaceUtilizationDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetMinSpaceUtilizationDifferenceOk() (*int32, bool) {
	if o == nil || o.MinSpaceUtilizationDifference == nil {
		return nil, false
	}
	return o.MinSpaceUtilizationDifference, true
}

// HasMinSpaceUtilizationDifference returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasMinSpaceUtilizationDifference() bool {
	if o != nil && o.MinSpaceUtilizationDifference != nil {
		return true
	}

	return false
}

// SetMinSpaceUtilizationDifference gets a reference to the given int32 and assigns it to the MinSpaceUtilizationDifference field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetMinSpaceUtilizationDifference(v int32) {
	o.MinSpaceUtilizationDifference = &v
}

// GetPolicyEnforcementAutomationMode returns the PolicyEnforcementAutomationMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetPolicyEnforcementAutomationMode() string {
	if o == nil || o.PolicyEnforcementAutomationMode == nil {
		var ret string
		return ret
	}
	return *o.PolicyEnforcementAutomationMode
}

// GetPolicyEnforcementAutomationModeOk returns a tuple with the PolicyEnforcementAutomationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetPolicyEnforcementAutomationModeOk() (*string, bool) {
	if o == nil || o.PolicyEnforcementAutomationMode == nil {
		return nil, false
	}
	return o.PolicyEnforcementAutomationMode, true
}

// HasPolicyEnforcementAutomationMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasPolicyEnforcementAutomationMode() bool {
	if o != nil && o.PolicyEnforcementAutomationMode != nil {
		return true
	}

	return false
}

// SetPolicyEnforcementAutomationMode gets a reference to the given string and assigns it to the PolicyEnforcementAutomationMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetPolicyEnforcementAutomationMode(v string) {
	o.PolicyEnforcementAutomationMode = &v
}

// GetReservablePercentThreshold returns the ReservablePercentThreshold field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetReservablePercentThreshold() int32 {
	if o == nil || o.ReservablePercentThreshold == nil {
		var ret int32
		return ret
	}
	return *o.ReservablePercentThreshold
}

// GetReservablePercentThresholdOk returns a tuple with the ReservablePercentThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetReservablePercentThresholdOk() (*int32, bool) {
	if o == nil || o.ReservablePercentThreshold == nil {
		return nil, false
	}
	return o.ReservablePercentThreshold, true
}

// HasReservablePercentThreshold returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasReservablePercentThreshold() bool {
	if o != nil && o.ReservablePercentThreshold != nil {
		return true
	}

	return false
}

// SetReservablePercentThreshold gets a reference to the given int32 and assigns it to the ReservablePercentThreshold field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetReservablePercentThreshold(v int32) {
	o.ReservablePercentThreshold = &v
}

// GetRuleEnforcementAutomationMode returns the RuleEnforcementAutomationMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetRuleEnforcementAutomationMode() string {
	if o == nil || o.RuleEnforcementAutomationMode == nil {
		var ret string
		return ret
	}
	return *o.RuleEnforcementAutomationMode
}

// GetRuleEnforcementAutomationModeOk returns a tuple with the RuleEnforcementAutomationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetRuleEnforcementAutomationModeOk() (*string, bool) {
	if o == nil || o.RuleEnforcementAutomationMode == nil {
		return nil, false
	}
	return o.RuleEnforcementAutomationMode, true
}

// HasRuleEnforcementAutomationMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasRuleEnforcementAutomationMode() bool {
	if o != nil && o.RuleEnforcementAutomationMode != nil {
		return true
	}

	return false
}

// SetRuleEnforcementAutomationMode gets a reference to the given string and assigns it to the RuleEnforcementAutomationMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetRuleEnforcementAutomationMode(v string) {
	o.RuleEnforcementAutomationMode = &v
}

// GetSpaceLoadBalanceAutomationMode returns the SpaceLoadBalanceAutomationMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetSpaceLoadBalanceAutomationMode() string {
	if o == nil || o.SpaceLoadBalanceAutomationMode == nil {
		var ret string
		return ret
	}
	return *o.SpaceLoadBalanceAutomationMode
}

// GetSpaceLoadBalanceAutomationModeOk returns a tuple with the SpaceLoadBalanceAutomationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetSpaceLoadBalanceAutomationModeOk() (*string, bool) {
	if o == nil || o.SpaceLoadBalanceAutomationMode == nil {
		return nil, false
	}
	return o.SpaceLoadBalanceAutomationMode, true
}

// HasSpaceLoadBalanceAutomationMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasSpaceLoadBalanceAutomationMode() bool {
	if o != nil && o.SpaceLoadBalanceAutomationMode != nil {
		return true
	}

	return false
}

// SetSpaceLoadBalanceAutomationMode gets a reference to the given string and assigns it to the SpaceLoadBalanceAutomationMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetSpaceLoadBalanceAutomationMode(v string) {
	o.SpaceLoadBalanceAutomationMode = &v
}

// GetSpaceThresholdMode returns the SpaceThresholdMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetSpaceThresholdMode() string {
	if o == nil || o.SpaceThresholdMode == nil {
		var ret string
		return ret
	}
	return *o.SpaceThresholdMode
}

// GetSpaceThresholdModeOk returns a tuple with the SpaceThresholdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetSpaceThresholdModeOk() (*string, bool) {
	if o == nil || o.SpaceThresholdMode == nil {
		return nil, false
	}
	return o.SpaceThresholdMode, true
}

// HasSpaceThresholdMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasSpaceThresholdMode() bool {
	if o != nil && o.SpaceThresholdMode != nil {
		return true
	}

	return false
}

// SetSpaceThresholdMode gets a reference to the given string and assigns it to the SpaceThresholdMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetSpaceThresholdMode(v string) {
	o.SpaceThresholdMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStorageDrsEnabled returns the StorageDrsEnabled field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetStorageDrsEnabled() bool {
	if o == nil || o.StorageDrsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.StorageDrsEnabled
}

// GetStorageDrsEnabledOk returns a tuple with the StorageDrsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetStorageDrsEnabledOk() (*bool, bool) {
	if o == nil || o.StorageDrsEnabled == nil {
		return nil, false
	}
	return o.StorageDrsEnabled, true
}

// HasStorageDrsEnabled returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasStorageDrsEnabled() bool {
	if o != nil && o.StorageDrsEnabled != nil {
		return true
	}

	return false
}

// SetStorageDrsEnabled gets a reference to the given bool and assigns it to the StorageDrsEnabled field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetStorageDrsEnabled(v bool) {
	o.StorageDrsEnabled = &v
}

// GetUtilizedSpaceThreshold returns the UtilizedSpaceThreshold field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetUtilizedSpaceThreshold() int32 {
	if o == nil || o.UtilizedSpaceThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UtilizedSpaceThreshold
}

// GetUtilizedSpaceThresholdOk returns a tuple with the UtilizedSpaceThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetUtilizedSpaceThresholdOk() (*int32, bool) {
	if o == nil || o.UtilizedSpaceThreshold == nil {
		return nil, false
	}
	return o.UtilizedSpaceThreshold, true
}

// HasUtilizedSpaceThreshold returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasUtilizedSpaceThreshold() bool {
	if o != nil && o.UtilizedSpaceThreshold != nil {
		return true
	}

	return false
}

// SetUtilizedSpaceThreshold gets a reference to the given int32 and assigns it to the UtilizedSpaceThreshold field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetUtilizedSpaceThreshold(v int32) {
	o.UtilizedSpaceThreshold = &v
}

// GetVmEvacuationAutomationMode returns the VmEvacuationAutomationMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetVmEvacuationAutomationMode() string {
	if o == nil || o.VmEvacuationAutomationMode == nil {
		var ret string
		return ret
	}
	return *o.VmEvacuationAutomationMode
}

// GetVmEvacuationAutomationModeOk returns a tuple with the VmEvacuationAutomationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetVmEvacuationAutomationModeOk() (*string, bool) {
	if o == nil || o.VmEvacuationAutomationMode == nil {
		return nil, false
	}
	return o.VmEvacuationAutomationMode, true
}

// HasVmEvacuationAutomationMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasVmEvacuationAutomationMode() bool {
	if o != nil && o.VmEvacuationAutomationMode != nil {
		return true
	}

	return false
}

// SetVmEvacuationAutomationMode gets a reference to the given string and assigns it to the VmEvacuationAutomationMode field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetVmEvacuationAutomationMode(v string) {
	o.VmEvacuationAutomationMode = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastoreClusterAllOf) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareDatastoreClusterAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

func (o VirtualizationVmwareDatastoreClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutomationLevel != nil {
		toSerialize["AutomationLevel"] = o.AutomationLevel
	}
	if o.FreeSpaceThreshold != nil {
		toSerialize["FreeSpaceThreshold"] = o.FreeSpaceThreshold
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.IoLatencyThreshold != nil {
		toSerialize["IoLatencyThreshold"] = o.IoLatencyThreshold
	}
	if o.IoLoadBalanceAutomationMode != nil {
		toSerialize["IoLoadBalanceAutomationMode"] = o.IoLoadBalanceAutomationMode
	}
	if o.IoLoadImbalanceThreshold != nil {
		toSerialize["IoLoadImbalanceThreshold"] = o.IoLoadImbalanceThreshold
	}
	if o.IoMetricsEnabled != nil {
		toSerialize["IoMetricsEnabled"] = o.IoMetricsEnabled
	}
	if o.MinSpaceUtilizationDifference != nil {
		toSerialize["MinSpaceUtilizationDifference"] = o.MinSpaceUtilizationDifference
	}
	if o.PolicyEnforcementAutomationMode != nil {
		toSerialize["PolicyEnforcementAutomationMode"] = o.PolicyEnforcementAutomationMode
	}
	if o.ReservablePercentThreshold != nil {
		toSerialize["ReservablePercentThreshold"] = o.ReservablePercentThreshold
	}
	if o.RuleEnforcementAutomationMode != nil {
		toSerialize["RuleEnforcementAutomationMode"] = o.RuleEnforcementAutomationMode
	}
	if o.SpaceLoadBalanceAutomationMode != nil {
		toSerialize["SpaceLoadBalanceAutomationMode"] = o.SpaceLoadBalanceAutomationMode
	}
	if o.SpaceThresholdMode != nil {
		toSerialize["SpaceThresholdMode"] = o.SpaceThresholdMode
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StorageDrsEnabled != nil {
		toSerialize["StorageDrsEnabled"] = o.StorageDrsEnabled
	}
	if o.UtilizedSpaceThreshold != nil {
		toSerialize["UtilizedSpaceThreshold"] = o.UtilizedSpaceThreshold
	}
	if o.VmEvacuationAutomationMode != nil {
		toSerialize["VmEvacuationAutomationMode"] = o.VmEvacuationAutomationMode
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDatastoreClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareDatastoreClusterAllOf := _VirtualizationVmwareDatastoreClusterAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareDatastoreClusterAllOf); err == nil {
		*o = VirtualizationVmwareDatastoreClusterAllOf(varVirtualizationVmwareDatastoreClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutomationLevel")
		delete(additionalProperties, "FreeSpaceThreshold")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "IoLatencyThreshold")
		delete(additionalProperties, "IoLoadBalanceAutomationMode")
		delete(additionalProperties, "IoLoadImbalanceThreshold")
		delete(additionalProperties, "IoMetricsEnabled")
		delete(additionalProperties, "MinSpaceUtilizationDifference")
		delete(additionalProperties, "PolicyEnforcementAutomationMode")
		delete(additionalProperties, "ReservablePercentThreshold")
		delete(additionalProperties, "RuleEnforcementAutomationMode")
		delete(additionalProperties, "SpaceLoadBalanceAutomationMode")
		delete(additionalProperties, "SpaceThresholdMode")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StorageDrsEnabled")
		delete(additionalProperties, "UtilizedSpaceThreshold")
		delete(additionalProperties, "VmEvacuationAutomationMode")
		delete(additionalProperties, "Datacenter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareDatastoreClusterAllOf struct {
	value *VirtualizationVmwareDatastoreClusterAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareDatastoreClusterAllOf) Get() *VirtualizationVmwareDatastoreClusterAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareDatastoreClusterAllOf) Set(val *VirtualizationVmwareDatastoreClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatastoreClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatastoreClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatastoreClusterAllOf(val *VirtualizationVmwareDatastoreClusterAllOf) *NullableVirtualizationVmwareDatastoreClusterAllOf {
	return &NullableVirtualizationVmwareDatastoreClusterAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatastoreClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatastoreClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
