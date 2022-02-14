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
	"time"
)

// HyperflexClusterHealthCheckExecutionSnapshotAllOf Definition of the list of properties defined in 'hyperflex.ClusterHealthCheckExecutionSnapshot', excluding properties defined in parent classes.
type HyperflexClusterHealthCheckExecutionSnapshotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Timestamp of the last health check execution on the HyperFlex cluster.
	Timestamp            *time.Time                           `json:"Timestamp,omitempty"`
	HxCluster            *HyperflexClusterRelationship        `json:"HxCluster,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship    `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterHealthCheckExecutionSnapshotAllOf HyperflexClusterHealthCheckExecutionSnapshotAllOf

// NewHyperflexClusterHealthCheckExecutionSnapshotAllOf instantiates a new HyperflexClusterHealthCheckExecutionSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterHealthCheckExecutionSnapshotAllOf(classId string, objectType string) *HyperflexClusterHealthCheckExecutionSnapshotAllOf {
	this := HyperflexClusterHealthCheckExecutionSnapshotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexClusterHealthCheckExecutionSnapshotAllOfWithDefaults instantiates a new HyperflexClusterHealthCheckExecutionSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterHealthCheckExecutionSnapshotAllOfWithDefaults() *HyperflexClusterHealthCheckExecutionSnapshotAllOf {
	this := HyperflexClusterHealthCheckExecutionSnapshotAllOf{}
	var classId string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o HyperflexClusterHealthCheckExecutionSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexClusterHealthCheckExecutionSnapshotAllOf := _HyperflexClusterHealthCheckExecutionSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexClusterHealthCheckExecutionSnapshotAllOf); err == nil {
		*o = HyperflexClusterHealthCheckExecutionSnapshotAllOf(varHyperflexClusterHealthCheckExecutionSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf struct {
	value *HyperflexClusterHealthCheckExecutionSnapshotAllOf
	isSet bool
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) Get() *HyperflexClusterHealthCheckExecutionSnapshotAllOf {
	return v.value
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) Set(val *HyperflexClusterHealthCheckExecutionSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterHealthCheckExecutionSnapshotAllOf(val *HyperflexClusterHealthCheckExecutionSnapshotAllOf) *NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf {
	return &NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf{value: val, isSet: true}
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
