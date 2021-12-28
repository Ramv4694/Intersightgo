/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-12-27T12:26:28Z.
 *
 * API version: 0.0.1-37434
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FabricLinkAggregationPolicyAllOf Definition of the list of properties defined in 'fabric.LinkAggregationPolicy', excluding properties defined in parent classes.
type FabricLinkAggregationPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag used to indicate whether LACP PDUs are to be sent 'fast', i.e., every 1 second. * `normal` - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * `fast` - The standard 4th generation UCS Fabric Interconnect with 54 ports.
	LacpRate *string `json:"LacpRate,omitempty"`
	// Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU.
	SuspendIndividual    *bool                                 `json:"SuspendIndividual,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricLinkAggregationPolicyAllOf FabricLinkAggregationPolicyAllOf

// NewFabricLinkAggregationPolicyAllOf instantiates a new FabricLinkAggregationPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLinkAggregationPolicyAllOf(classId string, objectType string) *FabricLinkAggregationPolicyAllOf {
	this := FabricLinkAggregationPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lacpRate string = "normal"
	this.LacpRate = &lacpRate
	return &this
}

// NewFabricLinkAggregationPolicyAllOfWithDefaults instantiates a new FabricLinkAggregationPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLinkAggregationPolicyAllOfWithDefaults() *FabricLinkAggregationPolicyAllOf {
	this := FabricLinkAggregationPolicyAllOf{}
	var classId string = "fabric.LinkAggregationPolicy"
	this.ClassId = classId
	var objectType string = "fabric.LinkAggregationPolicy"
	this.ObjectType = objectType
	var lacpRate string = "normal"
	this.LacpRate = &lacpRate
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricLinkAggregationPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricLinkAggregationPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricLinkAggregationPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricLinkAggregationPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLacpRate returns the LacpRate field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicyAllOf) GetLacpRate() string {
	if o == nil || o.LacpRate == nil {
		var ret string
		return ret
	}
	return *o.LacpRate
}

// GetLacpRateOk returns a tuple with the LacpRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicyAllOf) GetLacpRateOk() (*string, bool) {
	if o == nil || o.LacpRate == nil {
		return nil, false
	}
	return o.LacpRate, true
}

// HasLacpRate returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicyAllOf) HasLacpRate() bool {
	if o != nil && o.LacpRate != nil {
		return true
	}

	return false
}

// SetLacpRate gets a reference to the given string and assigns it to the LacpRate field.
func (o *FabricLinkAggregationPolicyAllOf) SetLacpRate(v string) {
	o.LacpRate = &v
}

// GetSuspendIndividual returns the SuspendIndividual field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicyAllOf) GetSuspendIndividual() bool {
	if o == nil || o.SuspendIndividual == nil {
		var ret bool
		return ret
	}
	return *o.SuspendIndividual
}

// GetSuspendIndividualOk returns a tuple with the SuspendIndividual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicyAllOf) GetSuspendIndividualOk() (*bool, bool) {
	if o == nil || o.SuspendIndividual == nil {
		return nil, false
	}
	return o.SuspendIndividual, true
}

// HasSuspendIndividual returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicyAllOf) HasSuspendIndividual() bool {
	if o != nil && o.SuspendIndividual != nil {
		return true
	}

	return false
}

// SetSuspendIndividual gets a reference to the given bool and assigns it to the SuspendIndividual field.
func (o *FabricLinkAggregationPolicyAllOf) SetSuspendIndividual(v bool) {
	o.SuspendIndividual = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricLinkAggregationPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricLinkAggregationPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LacpRate != nil {
		toSerialize["LacpRate"] = o.LacpRate
	}
	if o.SuspendIndividual != nil {
		toSerialize["SuspendIndividual"] = o.SuspendIndividual
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricLinkAggregationPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricLinkAggregationPolicyAllOf := _FabricLinkAggregationPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricLinkAggregationPolicyAllOf); err == nil {
		*o = FabricLinkAggregationPolicyAllOf(varFabricLinkAggregationPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LacpRate")
		delete(additionalProperties, "SuspendIndividual")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricLinkAggregationPolicyAllOf struct {
	value *FabricLinkAggregationPolicyAllOf
	isSet bool
}

func (v NullableFabricLinkAggregationPolicyAllOf) Get() *FabricLinkAggregationPolicyAllOf {
	return v.value
}

func (v *NullableFabricLinkAggregationPolicyAllOf) Set(val *FabricLinkAggregationPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLinkAggregationPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLinkAggregationPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLinkAggregationPolicyAllOf(val *FabricLinkAggregationPolicyAllOf) *NullableFabricLinkAggregationPolicyAllOf {
	return &NullableFabricLinkAggregationPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricLinkAggregationPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLinkAggregationPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
