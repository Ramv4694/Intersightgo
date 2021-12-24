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
)

// RecommendationPhysicalItemAllOf Definition of the list of properties defined in 'recommendation.PhysicalItem', excluding properties defined in parent classes.
type RecommendationPhysicalItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of the physical entity added.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation.
	Count *int64 `json:"Count,omitempty"`
	// If the PhysicalItem is new, this is set to true, else false.
	IsNew *bool `json:"IsNew,omitempty"`
	// Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation.
	MaxCount *int64 `json:"MaxCount,omitempty"`
	// Model of the recommended physical device which is externally identifiable.
	Model *string `json:"Model,omitempty"`
	// Moid of the managed object which represents the existing physical entity.
	SourceMoid *string `json:"SourceMoid,omitempty"`
	// Unit of the new capacity. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes.
	Unit *string `json:"Unit,omitempty"`
	// Uuid of the recommended physical device.
	Uuid           *string                                   `json:"Uuid,omitempty"`
	CapacityRunway *RecommendationCapacityRunwayRelationship `json:"CapacityRunway,omitempty"`
	// An array of relationships to recommendationPhysicalItem resources.
	PhysicalItem         []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationPhysicalItemAllOf RecommendationPhysicalItemAllOf

// NewRecommendationPhysicalItemAllOf instantiates a new RecommendationPhysicalItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationPhysicalItemAllOf(classId string, objectType string) *RecommendationPhysicalItemAllOf {
	this := RecommendationPhysicalItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var unit string = "TB"
	this.Unit = &unit
	return &this
}

// NewRecommendationPhysicalItemAllOfWithDefaults instantiates a new RecommendationPhysicalItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationPhysicalItemAllOfWithDefaults() *RecommendationPhysicalItemAllOf {
	this := RecommendationPhysicalItemAllOf{}
	var classId string = "recommendation.PhysicalItem"
	this.ClassId = classId
	var objectType string = "recommendation.PhysicalItem"
	this.ObjectType = objectType
	var unit string = "TB"
	this.Unit = &unit
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationPhysicalItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationPhysicalItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationPhysicalItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationPhysicalItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *RecommendationPhysicalItemAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RecommendationPhysicalItemAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetIsNew returns the IsNew field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetIsNew() bool {
	if o == nil || o.IsNew == nil {
		var ret bool
		return ret
	}
	return *o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetIsNewOk() (*bool, bool) {
	if o == nil || o.IsNew == nil {
		return nil, false
	}
	return o.IsNew, true
}

// HasIsNew returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasIsNew() bool {
	if o != nil && o.IsNew != nil {
		return true
	}

	return false
}

// SetIsNew gets a reference to the given bool and assigns it to the IsNew field.
func (o *RecommendationPhysicalItemAllOf) SetIsNew(v bool) {
	o.IsNew = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetMaxCount() int64 {
	if o == nil || o.MaxCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetMaxCountOk() (*int64, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int64 and assigns it to the MaxCount field.
func (o *RecommendationPhysicalItemAllOf) SetMaxCount(v int64) {
	o.MaxCount = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *RecommendationPhysicalItemAllOf) SetModel(v string) {
	o.Model = &v
}

// GetSourceMoid returns the SourceMoid field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetSourceMoid() string {
	if o == nil || o.SourceMoid == nil {
		var ret string
		return ret
	}
	return *o.SourceMoid
}

// GetSourceMoidOk returns a tuple with the SourceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetSourceMoidOk() (*string, bool) {
	if o == nil || o.SourceMoid == nil {
		return nil, false
	}
	return o.SourceMoid, true
}

// HasSourceMoid returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasSourceMoid() bool {
	if o != nil && o.SourceMoid != nil {
		return true
	}

	return false
}

// SetSourceMoid gets a reference to the given string and assigns it to the SourceMoid field.
func (o *RecommendationPhysicalItemAllOf) SetSourceMoid(v string) {
	o.SourceMoid = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *RecommendationPhysicalItemAllOf) SetUnit(v string) {
	o.Unit = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RecommendationPhysicalItemAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *RecommendationPhysicalItemAllOf) GetCapacityRunway() RecommendationCapacityRunwayRelationship {
	if o == nil || o.CapacityRunway == nil {
		var ret RecommendationCapacityRunwayRelationship
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItemAllOf) GetCapacityRunwayOk() (*RecommendationCapacityRunwayRelationship, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given RecommendationCapacityRunwayRelationship and assigns it to the CapacityRunway field.
func (o *RecommendationPhysicalItemAllOf) SetCapacityRunway(v RecommendationCapacityRunwayRelationship) {
	o.CapacityRunway = &v
}

// GetPhysicalItem returns the PhysicalItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationPhysicalItemAllOf) GetPhysicalItem() []RecommendationPhysicalItemRelationship {
	if o == nil {
		var ret []RecommendationPhysicalItemRelationship
		return ret
	}
	return o.PhysicalItem
}

// GetPhysicalItemOk returns a tuple with the PhysicalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationPhysicalItemAllOf) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool) {
	if o == nil || o.PhysicalItem == nil {
		return nil, false
	}
	return &o.PhysicalItem, true
}

// HasPhysicalItem returns a boolean if a field has been set.
func (o *RecommendationPhysicalItemAllOf) HasPhysicalItem() bool {
	if o != nil && o.PhysicalItem != nil {
		return true
	}

	return false
}

// SetPhysicalItem gets a reference to the given []RecommendationPhysicalItemRelationship and assigns it to the PhysicalItem field.
func (o *RecommendationPhysicalItemAllOf) SetPhysicalItem(v []RecommendationPhysicalItemRelationship) {
	o.PhysicalItem = v
}

func (o RecommendationPhysicalItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.IsNew != nil {
		toSerialize["IsNew"] = o.IsNew
	}
	if o.MaxCount != nil {
		toSerialize["MaxCount"] = o.MaxCount
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.SourceMoid != nil {
		toSerialize["SourceMoid"] = o.SourceMoid
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.PhysicalItem != nil {
		toSerialize["PhysicalItem"] = o.PhysicalItem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationPhysicalItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationPhysicalItemAllOf := _RecommendationPhysicalItemAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationPhysicalItemAllOf); err == nil {
		*o = RecommendationPhysicalItemAllOf(varRecommendationPhysicalItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "IsNew")
		delete(additionalProperties, "MaxCount")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "SourceMoid")
		delete(additionalProperties, "Unit")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "CapacityRunway")
		delete(additionalProperties, "PhysicalItem")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationPhysicalItemAllOf struct {
	value *RecommendationPhysicalItemAllOf
	isSet bool
}

func (v NullableRecommendationPhysicalItemAllOf) Get() *RecommendationPhysicalItemAllOf {
	return v.value
}

func (v *NullableRecommendationPhysicalItemAllOf) Set(val *RecommendationPhysicalItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationPhysicalItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationPhysicalItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationPhysicalItemAllOf(val *RecommendationPhysicalItemAllOf) *NullableRecommendationPhysicalItemAllOf {
	return &NullableRecommendationPhysicalItemAllOf{value: val, isSet: true}
}

func (v NullableRecommendationPhysicalItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationPhysicalItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
