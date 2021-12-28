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

// BulkMoClonerAllOf Definition of the list of properties defined in 'bulk.MoCloner', excluding properties defined in parent classes.
type BulkMoClonerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	Responses            []BulkRestResult                      `json:"Responses,omitempty"`
	Sources              []MoBaseMo                            `json:"Sources,omitempty"`
	Targets              []MoBaseMo                            `json:"Targets,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoClonerAllOf BulkMoClonerAllOf

// NewBulkMoClonerAllOf instantiates a new BulkMoClonerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoClonerAllOf(classId string, objectType string) *BulkMoClonerAllOf {
	this := BulkMoClonerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkMoClonerAllOfWithDefaults instantiates a new BulkMoClonerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoClonerAllOfWithDefaults() *BulkMoClonerAllOf {
	this := BulkMoClonerAllOf{}
	var classId string = "bulk.MoCloner"
	this.ClassId = classId
	var objectType string = "bulk.MoCloner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoClonerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoClonerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoClonerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoClonerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoClonerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoClonerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetResponses returns the Responses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoClonerAllOf) GetResponses() []BulkRestResult {
	if o == nil {
		var ret []BulkRestResult
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoClonerAllOf) GetResponsesOk() (*[]BulkRestResult, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return &o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *BulkMoClonerAllOf) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BulkRestResult and assigns it to the Responses field.
func (o *BulkMoClonerAllOf) SetResponses(v []BulkRestResult) {
	o.Responses = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoClonerAllOf) GetSources() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoClonerAllOf) GetSourcesOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *BulkMoClonerAllOf) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []MoBaseMo and assigns it to the Sources field.
func (o *BulkMoClonerAllOf) SetSources(v []MoBaseMo) {
	o.Sources = v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoClonerAllOf) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoClonerAllOf) GetTargetsOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoClonerAllOf) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoClonerAllOf) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkMoClonerAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoClonerAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoClonerAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoClonerAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkMoClonerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Responses != nil {
		toSerialize["Responses"] = o.Responses
	}
	if o.Sources != nil {
		toSerialize["Sources"] = o.Sources
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkMoClonerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkMoClonerAllOf := _BulkMoClonerAllOf{}

	if err = json.Unmarshal(bytes, &varBulkMoClonerAllOf); err == nil {
		*o = BulkMoClonerAllOf(varBulkMoClonerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Responses")
		delete(additionalProperties, "Sources")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkMoClonerAllOf struct {
	value *BulkMoClonerAllOf
	isSet bool
}

func (v NullableBulkMoClonerAllOf) Get() *BulkMoClonerAllOf {
	return v.value
}

func (v *NullableBulkMoClonerAllOf) Set(val *BulkMoClonerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoClonerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoClonerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoClonerAllOf(val *BulkMoClonerAllOf) *NullableBulkMoClonerAllOf {
	return &NullableBulkMoClonerAllOf{value: val, isSet: true}
}

func (v NullableBulkMoClonerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoClonerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
