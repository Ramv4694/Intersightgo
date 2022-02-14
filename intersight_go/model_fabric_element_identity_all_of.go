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

// FabricElementIdentityAllOf Definition of the list of properties defined in 'fabric.ElementIdentity', excluding properties defined in parent classes.
type FabricElementIdentityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Fabric Interconnect domain.
	Domain *string `json:"Domain,omitempty"`
	// Replacement type specifies whether it is single FI or domain replacement. * `None` - The default action is none. * `Individual` - Replacement of single network element. * `Domain` - Domain indicates the replacement of Fabric Interconnect domain.
	ReplacementType *string `json:"ReplacementType,omitempty"`
	// Switch Identifier that uniquely represents the fabric object. * `A` - Switch Identifier of Fabric Interconnect A. * `B` - Switch Identifier of Fabric Interconnect B.
	SwitchId             *string                     `json:"SwitchId,omitempty"`
	NetworkElement       *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	ReplacementTarget    *NetworkElementRelationship `json:"ReplacementTarget,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricElementIdentityAllOf FabricElementIdentityAllOf

// NewFabricElementIdentityAllOf instantiates a new FabricElementIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricElementIdentityAllOf(classId string, objectType string) *FabricElementIdentityAllOf {
	this := FabricElementIdentityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var replacementType string = "None"
	this.ReplacementType = &replacementType
	var switchId string = "A"
	this.SwitchId = &switchId
	return &this
}

// NewFabricElementIdentityAllOfWithDefaults instantiates a new FabricElementIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricElementIdentityAllOfWithDefaults() *FabricElementIdentityAllOf {
	this := FabricElementIdentityAllOf{}
	var classId string = "fabric.ElementIdentity"
	this.ClassId = classId
	var objectType string = "fabric.ElementIdentity"
	this.ObjectType = objectType
	var replacementType string = "None"
	this.ReplacementType = &replacementType
	var switchId string = "A"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricElementIdentityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricElementIdentityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricElementIdentityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricElementIdentityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *FabricElementIdentityAllOf) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *FabricElementIdentityAllOf) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *FabricElementIdentityAllOf) SetDomain(v string) {
	o.Domain = &v
}

// GetReplacementType returns the ReplacementType field value if set, zero value otherwise.
func (o *FabricElementIdentityAllOf) GetReplacementType() string {
	if o == nil || o.ReplacementType == nil {
		var ret string
		return ret
	}
	return *o.ReplacementType
}

// GetReplacementTypeOk returns a tuple with the ReplacementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetReplacementTypeOk() (*string, bool) {
	if o == nil || o.ReplacementType == nil {
		return nil, false
	}
	return o.ReplacementType, true
}

// HasReplacementType returns a boolean if a field has been set.
func (o *FabricElementIdentityAllOf) HasReplacementType() bool {
	if o != nil && o.ReplacementType != nil {
		return true
	}

	return false
}

// SetReplacementType gets a reference to the given string and assigns it to the ReplacementType field.
func (o *FabricElementIdentityAllOf) SetReplacementType(v string) {
	o.ReplacementType = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *FabricElementIdentityAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *FabricElementIdentityAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *FabricElementIdentityAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricElementIdentityAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricElementIdentityAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricElementIdentityAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetReplacementTarget returns the ReplacementTarget field value if set, zero value otherwise.
func (o *FabricElementIdentityAllOf) GetReplacementTarget() NetworkElementRelationship {
	if o == nil || o.ReplacementTarget == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.ReplacementTarget
}

// GetReplacementTargetOk returns a tuple with the ReplacementTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricElementIdentityAllOf) GetReplacementTargetOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.ReplacementTarget == nil {
		return nil, false
	}
	return o.ReplacementTarget, true
}

// HasReplacementTarget returns a boolean if a field has been set.
func (o *FabricElementIdentityAllOf) HasReplacementTarget() bool {
	if o != nil && o.ReplacementTarget != nil {
		return true
	}

	return false
}

// SetReplacementTarget gets a reference to the given NetworkElementRelationship and assigns it to the ReplacementTarget field.
func (o *FabricElementIdentityAllOf) SetReplacementTarget(v NetworkElementRelationship) {
	o.ReplacementTarget = &v
}

func (o FabricElementIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Domain != nil {
		toSerialize["Domain"] = o.Domain
	}
	if o.ReplacementType != nil {
		toSerialize["ReplacementType"] = o.ReplacementType
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.ReplacementTarget != nil {
		toSerialize["ReplacementTarget"] = o.ReplacementTarget
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricElementIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricElementIdentityAllOf := _FabricElementIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varFabricElementIdentityAllOf); err == nil {
		*o = FabricElementIdentityAllOf(varFabricElementIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "ReplacementType")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "ReplacementTarget")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricElementIdentityAllOf struct {
	value *FabricElementIdentityAllOf
	isSet bool
}

func (v NullableFabricElementIdentityAllOf) Get() *FabricElementIdentityAllOf {
	return v.value
}

func (v *NullableFabricElementIdentityAllOf) Set(val *FabricElementIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricElementIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricElementIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricElementIdentityAllOf(val *FabricElementIdentityAllOf) *NullableFabricElementIdentityAllOf {
	return &NullableFabricElementIdentityAllOf{value: val, isSet: true}
}

func (v NullableFabricElementIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricElementIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
