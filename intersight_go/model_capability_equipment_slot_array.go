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
	"reflect"
	"strings"
)

// CapabilityEquipmentSlotArray Type to represent additional switch specific capabilities.
type CapabilityEquipmentSlotArray struct {
	CapabilitySwitchCapabilityDef
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// First Index information for a Switch/Fabric-Interconnect hardware.
	FirstIndex *float32 `json:"FirstIndex,omitempty"`
	// Height information for a Switch/Fabric-Interconnect hardware.
	Height *float32 `json:"Height,omitempty"`
	// Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware.
	HorizontalStartOffset *float32 `json:"HorizontalStartOffset,omitempty"`
	// Inline Group Separation information for a Switch/Fabric-Interconnect hardware.
	InlineGroupSeparation *float32 `json:"InlineGroupSeparation,omitempty"`
	// Inline Group Size information for a Switch/Fabric-Interconnect hardware.
	InlineGroupSize *float32 `json:"InlineGroupSize,omitempty"`
	// Inline Offset information for a Switch/Fabric-Interconnect hardware.
	InlineOffset *float32 `json:"InlineOffset,omitempty"`
	// Location information for a Switch/Fabric-Interconnect hardware.
	Location *string `json:"Location,omitempty"`
	// Number of Slots information for a Switch/Fabric-Interconnect hardware.
	NumberOfSlots *int64 `json:"NumberOfSlots,omitempty"`
	// Orientation information for a Switch/Fabric-Interconnect hardware.
	Orientation *string `json:"Orientation,omitempty"`
	// Selector information for a Switch/Fabric-Interconnect hardware.
	Selector *string `json:"Selector,omitempty"`
	// Slots per line information for a Switch/Fabric-Interconnect hardware.
	SlotsPerLine *int64 `json:"SlotsPerLine,omitempty"`
	// Transverse Group Separation information for a Switch/Fabric-Interconnect hardware.
	TransverseGroupSeparation *float32 `json:"TransverseGroupSeparation,omitempty"`
	// Transverse Group Size information for a Switch/Fabric-Interconnect hardware.
	TransverseGroupSize *float32 `json:"TransverseGroupSize,omitempty"`
	// Transverse Offset information for a Switch/Fabric-Interconnect hardware.
	TransverseOffset *float32 `json:"TransverseOffset,omitempty"`
	// Vertical Start Offset information for a Switch/Fabric-Interconnect hardware.
	VerticalStartOffset *float32 `json:"VerticalStartOffset,omitempty"`
	// Width information for a Switch/Fabric-Interconnect hardware.
	Width                *float32 `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEquipmentSlotArray CapabilityEquipmentSlotArray

// NewCapabilityEquipmentSlotArray instantiates a new CapabilityEquipmentSlotArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEquipmentSlotArray(classId string, objectType string) *CapabilityEquipmentSlotArray {
	this := CapabilityEquipmentSlotArray{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityEquipmentSlotArrayWithDefaults instantiates a new CapabilityEquipmentSlotArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEquipmentSlotArrayWithDefaults() *CapabilityEquipmentSlotArray {
	this := CapabilityEquipmentSlotArray{}
	var classId string = "capability.EquipmentSlotArray"
	this.ClassId = classId
	var objectType string = "capability.EquipmentSlotArray"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityEquipmentSlotArray) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityEquipmentSlotArray) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityEquipmentSlotArray) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityEquipmentSlotArray) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFirstIndex returns the FirstIndex field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetFirstIndex() float32 {
	if o == nil || o.FirstIndex == nil {
		var ret float32
		return ret
	}
	return *o.FirstIndex
}

// GetFirstIndexOk returns a tuple with the FirstIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetFirstIndexOk() (*float32, bool) {
	if o == nil || o.FirstIndex == nil {
		return nil, false
	}
	return o.FirstIndex, true
}

// HasFirstIndex returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasFirstIndex() bool {
	if o != nil && o.FirstIndex != nil {
		return true
	}

	return false
}

// SetFirstIndex gets a reference to the given float32 and assigns it to the FirstIndex field.
func (o *CapabilityEquipmentSlotArray) SetFirstIndex(v float32) {
	o.FirstIndex = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CapabilityEquipmentSlotArray) SetHeight(v float32) {
	o.Height = &v
}

// GetHorizontalStartOffset returns the HorizontalStartOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetHorizontalStartOffset() float32 {
	if o == nil || o.HorizontalStartOffset == nil {
		var ret float32
		return ret
	}
	return *o.HorizontalStartOffset
}

// GetHorizontalStartOffsetOk returns a tuple with the HorizontalStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetHorizontalStartOffsetOk() (*float32, bool) {
	if o == nil || o.HorizontalStartOffset == nil {
		return nil, false
	}
	return o.HorizontalStartOffset, true
}

// HasHorizontalStartOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasHorizontalStartOffset() bool {
	if o != nil && o.HorizontalStartOffset != nil {
		return true
	}

	return false
}

// SetHorizontalStartOffset gets a reference to the given float32 and assigns it to the HorizontalStartOffset field.
func (o *CapabilityEquipmentSlotArray) SetHorizontalStartOffset(v float32) {
	o.HorizontalStartOffset = &v
}

// GetInlineGroupSeparation returns the InlineGroupSeparation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetInlineGroupSeparation() float32 {
	if o == nil || o.InlineGroupSeparation == nil {
		var ret float32
		return ret
	}
	return *o.InlineGroupSeparation
}

// GetInlineGroupSeparationOk returns a tuple with the InlineGroupSeparation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetInlineGroupSeparationOk() (*float32, bool) {
	if o == nil || o.InlineGroupSeparation == nil {
		return nil, false
	}
	return o.InlineGroupSeparation, true
}

// HasInlineGroupSeparation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasInlineGroupSeparation() bool {
	if o != nil && o.InlineGroupSeparation != nil {
		return true
	}

	return false
}

// SetInlineGroupSeparation gets a reference to the given float32 and assigns it to the InlineGroupSeparation field.
func (o *CapabilityEquipmentSlotArray) SetInlineGroupSeparation(v float32) {
	o.InlineGroupSeparation = &v
}

// GetInlineGroupSize returns the InlineGroupSize field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetInlineGroupSize() float32 {
	if o == nil || o.InlineGroupSize == nil {
		var ret float32
		return ret
	}
	return *o.InlineGroupSize
}

// GetInlineGroupSizeOk returns a tuple with the InlineGroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetInlineGroupSizeOk() (*float32, bool) {
	if o == nil || o.InlineGroupSize == nil {
		return nil, false
	}
	return o.InlineGroupSize, true
}

// HasInlineGroupSize returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasInlineGroupSize() bool {
	if o != nil && o.InlineGroupSize != nil {
		return true
	}

	return false
}

// SetInlineGroupSize gets a reference to the given float32 and assigns it to the InlineGroupSize field.
func (o *CapabilityEquipmentSlotArray) SetInlineGroupSize(v float32) {
	o.InlineGroupSize = &v
}

// GetInlineOffset returns the InlineOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetInlineOffset() float32 {
	if o == nil || o.InlineOffset == nil {
		var ret float32
		return ret
	}
	return *o.InlineOffset
}

// GetInlineOffsetOk returns a tuple with the InlineOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetInlineOffsetOk() (*float32, bool) {
	if o == nil || o.InlineOffset == nil {
		return nil, false
	}
	return o.InlineOffset, true
}

// HasInlineOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasInlineOffset() bool {
	if o != nil && o.InlineOffset != nil {
		return true
	}

	return false
}

// SetInlineOffset gets a reference to the given float32 and assigns it to the InlineOffset field.
func (o *CapabilityEquipmentSlotArray) SetInlineOffset(v float32) {
	o.InlineOffset = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *CapabilityEquipmentSlotArray) SetLocation(v string) {
	o.Location = &v
}

// GetNumberOfSlots returns the NumberOfSlots field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetNumberOfSlots() int64 {
	if o == nil || o.NumberOfSlots == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSlots
}

// GetNumberOfSlotsOk returns a tuple with the NumberOfSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetNumberOfSlotsOk() (*int64, bool) {
	if o == nil || o.NumberOfSlots == nil {
		return nil, false
	}
	return o.NumberOfSlots, true
}

// HasNumberOfSlots returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasNumberOfSlots() bool {
	if o != nil && o.NumberOfSlots != nil {
		return true
	}

	return false
}

// SetNumberOfSlots gets a reference to the given int64 and assigns it to the NumberOfSlots field.
func (o *CapabilityEquipmentSlotArray) SetNumberOfSlots(v int64) {
	o.NumberOfSlots = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *CapabilityEquipmentSlotArray) SetOrientation(v string) {
	o.Orientation = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *CapabilityEquipmentSlotArray) SetSelector(v string) {
	o.Selector = &v
}

// GetSlotsPerLine returns the SlotsPerLine field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetSlotsPerLine() int64 {
	if o == nil || o.SlotsPerLine == nil {
		var ret int64
		return ret
	}
	return *o.SlotsPerLine
}

// GetSlotsPerLineOk returns a tuple with the SlotsPerLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetSlotsPerLineOk() (*int64, bool) {
	if o == nil || o.SlotsPerLine == nil {
		return nil, false
	}
	return o.SlotsPerLine, true
}

// HasSlotsPerLine returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasSlotsPerLine() bool {
	if o != nil && o.SlotsPerLine != nil {
		return true
	}

	return false
}

// SetSlotsPerLine gets a reference to the given int64 and assigns it to the SlotsPerLine field.
func (o *CapabilityEquipmentSlotArray) SetSlotsPerLine(v int64) {
	o.SlotsPerLine = &v
}

// GetTransverseGroupSeparation returns the TransverseGroupSeparation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSeparation() float32 {
	if o == nil || o.TransverseGroupSeparation == nil {
		var ret float32
		return ret
	}
	return *o.TransverseGroupSeparation
}

// GetTransverseGroupSeparationOk returns a tuple with the TransverseGroupSeparation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSeparationOk() (*float32, bool) {
	if o == nil || o.TransverseGroupSeparation == nil {
		return nil, false
	}
	return o.TransverseGroupSeparation, true
}

// HasTransverseGroupSeparation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasTransverseGroupSeparation() bool {
	if o != nil && o.TransverseGroupSeparation != nil {
		return true
	}

	return false
}

// SetTransverseGroupSeparation gets a reference to the given float32 and assigns it to the TransverseGroupSeparation field.
func (o *CapabilityEquipmentSlotArray) SetTransverseGroupSeparation(v float32) {
	o.TransverseGroupSeparation = &v
}

// GetTransverseGroupSize returns the TransverseGroupSize field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSize() float32 {
	if o == nil || o.TransverseGroupSize == nil {
		var ret float32
		return ret
	}
	return *o.TransverseGroupSize
}

// GetTransverseGroupSizeOk returns a tuple with the TransverseGroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetTransverseGroupSizeOk() (*float32, bool) {
	if o == nil || o.TransverseGroupSize == nil {
		return nil, false
	}
	return o.TransverseGroupSize, true
}

// HasTransverseGroupSize returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasTransverseGroupSize() bool {
	if o != nil && o.TransverseGroupSize != nil {
		return true
	}

	return false
}

// SetTransverseGroupSize gets a reference to the given float32 and assigns it to the TransverseGroupSize field.
func (o *CapabilityEquipmentSlotArray) SetTransverseGroupSize(v float32) {
	o.TransverseGroupSize = &v
}

// GetTransverseOffset returns the TransverseOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetTransverseOffset() float32 {
	if o == nil || o.TransverseOffset == nil {
		var ret float32
		return ret
	}
	return *o.TransverseOffset
}

// GetTransverseOffsetOk returns a tuple with the TransverseOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetTransverseOffsetOk() (*float32, bool) {
	if o == nil || o.TransverseOffset == nil {
		return nil, false
	}
	return o.TransverseOffset, true
}

// HasTransverseOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasTransverseOffset() bool {
	if o != nil && o.TransverseOffset != nil {
		return true
	}

	return false
}

// SetTransverseOffset gets a reference to the given float32 and assigns it to the TransverseOffset field.
func (o *CapabilityEquipmentSlotArray) SetTransverseOffset(v float32) {
	o.TransverseOffset = &v
}

// GetVerticalStartOffset returns the VerticalStartOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetVerticalStartOffset() float32 {
	if o == nil || o.VerticalStartOffset == nil {
		var ret float32
		return ret
	}
	return *o.VerticalStartOffset
}

// GetVerticalStartOffsetOk returns a tuple with the VerticalStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetVerticalStartOffsetOk() (*float32, bool) {
	if o == nil || o.VerticalStartOffset == nil {
		return nil, false
	}
	return o.VerticalStartOffset, true
}

// HasVerticalStartOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasVerticalStartOffset() bool {
	if o != nil && o.VerticalStartOffset != nil {
		return true
	}

	return false
}

// SetVerticalStartOffset gets a reference to the given float32 and assigns it to the VerticalStartOffset field.
func (o *CapabilityEquipmentSlotArray) SetVerticalStartOffset(v float32) {
	o.VerticalStartOffset = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArray) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArray) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArray) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CapabilityEquipmentSlotArray) SetWidth(v float32) {
	o.Width = &v
}

func (o CapabilityEquipmentSlotArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilitySwitchCapabilityDef, errCapabilitySwitchCapabilityDef := json.Marshal(o.CapabilitySwitchCapabilityDef)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	errCapabilitySwitchCapabilityDef = json.Unmarshal([]byte(serializedCapabilitySwitchCapabilityDef), &toSerialize)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FirstIndex != nil {
		toSerialize["FirstIndex"] = o.FirstIndex
	}
	if o.Height != nil {
		toSerialize["Height"] = o.Height
	}
	if o.HorizontalStartOffset != nil {
		toSerialize["HorizontalStartOffset"] = o.HorizontalStartOffset
	}
	if o.InlineGroupSeparation != nil {
		toSerialize["InlineGroupSeparation"] = o.InlineGroupSeparation
	}
	if o.InlineGroupSize != nil {
		toSerialize["InlineGroupSize"] = o.InlineGroupSize
	}
	if o.InlineOffset != nil {
		toSerialize["InlineOffset"] = o.InlineOffset
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.NumberOfSlots != nil {
		toSerialize["NumberOfSlots"] = o.NumberOfSlots
	}
	if o.Orientation != nil {
		toSerialize["Orientation"] = o.Orientation
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}
	if o.SlotsPerLine != nil {
		toSerialize["SlotsPerLine"] = o.SlotsPerLine
	}
	if o.TransverseGroupSeparation != nil {
		toSerialize["TransverseGroupSeparation"] = o.TransverseGroupSeparation
	}
	if o.TransverseGroupSize != nil {
		toSerialize["TransverseGroupSize"] = o.TransverseGroupSize
	}
	if o.TransverseOffset != nil {
		toSerialize["TransverseOffset"] = o.TransverseOffset
	}
	if o.VerticalStartOffset != nil {
		toSerialize["VerticalStartOffset"] = o.VerticalStartOffset
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEquipmentSlotArray) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityEquipmentSlotArrayWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// First Index information for a Switch/Fabric-Interconnect hardware.
		FirstIndex *float32 `json:"FirstIndex,omitempty"`
		// Height information for a Switch/Fabric-Interconnect hardware.
		Height *float32 `json:"Height,omitempty"`
		// Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware.
		HorizontalStartOffset *float32 `json:"HorizontalStartOffset,omitempty"`
		// Inline Group Separation information for a Switch/Fabric-Interconnect hardware.
		InlineGroupSeparation *float32 `json:"InlineGroupSeparation,omitempty"`
		// Inline Group Size information for a Switch/Fabric-Interconnect hardware.
		InlineGroupSize *float32 `json:"InlineGroupSize,omitempty"`
		// Inline Offset information for a Switch/Fabric-Interconnect hardware.
		InlineOffset *float32 `json:"InlineOffset,omitempty"`
		// Location information for a Switch/Fabric-Interconnect hardware.
		Location *string `json:"Location,omitempty"`
		// Number of Slots information for a Switch/Fabric-Interconnect hardware.
		NumberOfSlots *int64 `json:"NumberOfSlots,omitempty"`
		// Orientation information for a Switch/Fabric-Interconnect hardware.
		Orientation *string `json:"Orientation,omitempty"`
		// Selector information for a Switch/Fabric-Interconnect hardware.
		Selector *string `json:"Selector,omitempty"`
		// Slots per line information for a Switch/Fabric-Interconnect hardware.
		SlotsPerLine *int64 `json:"SlotsPerLine,omitempty"`
		// Transverse Group Separation information for a Switch/Fabric-Interconnect hardware.
		TransverseGroupSeparation *float32 `json:"TransverseGroupSeparation,omitempty"`
		// Transverse Group Size information for a Switch/Fabric-Interconnect hardware.
		TransverseGroupSize *float32 `json:"TransverseGroupSize,omitempty"`
		// Transverse Offset information for a Switch/Fabric-Interconnect hardware.
		TransverseOffset *float32 `json:"TransverseOffset,omitempty"`
		// Vertical Start Offset information for a Switch/Fabric-Interconnect hardware.
		VerticalStartOffset *float32 `json:"VerticalStartOffset,omitempty"`
		// Width information for a Switch/Fabric-Interconnect hardware.
		Width *float32 `json:"Width,omitempty"`
	}

	varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct := CapabilityEquipmentSlotArrayWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityEquipmentSlotArray := _CapabilityEquipmentSlotArray{}
		varCapabilityEquipmentSlotArray.ClassId = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.ClassId
		varCapabilityEquipmentSlotArray.ObjectType = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.ObjectType
		varCapabilityEquipmentSlotArray.FirstIndex = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.FirstIndex
		varCapabilityEquipmentSlotArray.Height = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.Height
		varCapabilityEquipmentSlotArray.HorizontalStartOffset = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.HorizontalStartOffset
		varCapabilityEquipmentSlotArray.InlineGroupSeparation = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.InlineGroupSeparation
		varCapabilityEquipmentSlotArray.InlineGroupSize = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.InlineGroupSize
		varCapabilityEquipmentSlotArray.InlineOffset = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.InlineOffset
		varCapabilityEquipmentSlotArray.Location = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.Location
		varCapabilityEquipmentSlotArray.NumberOfSlots = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.NumberOfSlots
		varCapabilityEquipmentSlotArray.Orientation = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.Orientation
		varCapabilityEquipmentSlotArray.Selector = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.Selector
		varCapabilityEquipmentSlotArray.SlotsPerLine = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.SlotsPerLine
		varCapabilityEquipmentSlotArray.TransverseGroupSeparation = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.TransverseGroupSeparation
		varCapabilityEquipmentSlotArray.TransverseGroupSize = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.TransverseGroupSize
		varCapabilityEquipmentSlotArray.TransverseOffset = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.TransverseOffset
		varCapabilityEquipmentSlotArray.VerticalStartOffset = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.VerticalStartOffset
		varCapabilityEquipmentSlotArray.Width = varCapabilityEquipmentSlotArrayWithoutEmbeddedStruct.Width
		*o = CapabilityEquipmentSlotArray(varCapabilityEquipmentSlotArray)
	} else {
		return err
	}

	varCapabilityEquipmentSlotArray := _CapabilityEquipmentSlotArray{}

	err = json.Unmarshal(bytes, &varCapabilityEquipmentSlotArray)
	if err == nil {
		o.CapabilitySwitchCapabilityDef = varCapabilityEquipmentSlotArray.CapabilitySwitchCapabilityDef
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FirstIndex")
		delete(additionalProperties, "Height")
		delete(additionalProperties, "HorizontalStartOffset")
		delete(additionalProperties, "InlineGroupSeparation")
		delete(additionalProperties, "InlineGroupSize")
		delete(additionalProperties, "InlineOffset")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "NumberOfSlots")
		delete(additionalProperties, "Orientation")
		delete(additionalProperties, "Selector")
		delete(additionalProperties, "SlotsPerLine")
		delete(additionalProperties, "TransverseGroupSeparation")
		delete(additionalProperties, "TransverseGroupSize")
		delete(additionalProperties, "TransverseOffset")
		delete(additionalProperties, "VerticalStartOffset")
		delete(additionalProperties, "Width")

		// remove fields from embedded structs
		reflectCapabilitySwitchCapabilityDef := reflect.ValueOf(o.CapabilitySwitchCapabilityDef)
		for i := 0; i < reflectCapabilitySwitchCapabilityDef.Type().NumField(); i++ {
			t := reflectCapabilitySwitchCapabilityDef.Type().Field(i)

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

type NullableCapabilityEquipmentSlotArray struct {
	value *CapabilityEquipmentSlotArray
	isSet bool
}

func (v NullableCapabilityEquipmentSlotArray) Get() *CapabilityEquipmentSlotArray {
	return v.value
}

func (v *NullableCapabilityEquipmentSlotArray) Set(val *CapabilityEquipmentSlotArray) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEquipmentSlotArray) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEquipmentSlotArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEquipmentSlotArray(val *CapabilityEquipmentSlotArray) *NullableCapabilityEquipmentSlotArray {
	return &NullableCapabilityEquipmentSlotArray{value: val, isSet: true}
}

func (v NullableCapabilityEquipmentSlotArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEquipmentSlotArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
