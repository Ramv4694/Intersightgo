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

// VirtualizationIweDvswitchAllOf Definition of the list of properties defined in 'virtualization.IweDvswitch', excluding properties defined in parent classes.
type VirtualizationIweDvswitchAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the dvUplink referenced by this dvswitch.
	DvUplink *string `json:"DvUplink,omitempty"`
	// The last host that update this object.
	LastHostname *string                               `json:"LastHostname,omitempty"`
	Cluster      *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to virtualizationIweHost resources.
	MemberHosts []VirtualizationIweHostRelationship `json:"MemberHosts,omitempty"`
	// An array of relationships to virtualizationIweHostVswitch resources.
	MemberVswitches      []VirtualizationIweHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweDvswitchAllOf VirtualizationIweDvswitchAllOf

// NewVirtualizationIweDvswitchAllOf instantiates a new VirtualizationIweDvswitchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweDvswitchAllOf(classId string, objectType string) *VirtualizationIweDvswitchAllOf {
	this := VirtualizationIweDvswitchAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIweDvswitchAllOfWithDefaults instantiates a new VirtualizationIweDvswitchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweDvswitchAllOfWithDefaults() *VirtualizationIweDvswitchAllOf {
	this := VirtualizationIweDvswitchAllOf{}
	var classId string = "virtualization.IweDvswitch"
	this.ClassId = classId
	var objectType string = "virtualization.IweDvswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweDvswitchAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitchAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweDvswitchAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweDvswitchAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitchAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweDvswitchAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitchAllOf) GetDvUplink() string {
	if o == nil || o.DvUplink == nil {
		var ret string
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitchAllOf) GetDvUplinkOk() (*string, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitchAllOf) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given string and assigns it to the DvUplink field.
func (o *VirtualizationIweDvswitchAllOf) SetDvUplink(v string) {
	o.DvUplink = &v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitchAllOf) GetLastHostname() string {
	if o == nil || o.LastHostname == nil {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitchAllOf) GetLastHostnameOk() (*string, bool) {
	if o == nil || o.LastHostname == nil {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitchAllOf) HasLastHostname() bool {
	if o != nil && o.LastHostname != nil {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *VirtualizationIweDvswitchAllOf) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitchAllOf) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitchAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitchAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweDvswitchAllOf) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetMemberHosts returns the MemberHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvswitchAllOf) GetMemberHosts() []VirtualizationIweHostRelationship {
	if o == nil {
		var ret []VirtualizationIweHostRelationship
		return ret
	}
	return o.MemberHosts
}

// GetMemberHostsOk returns a tuple with the MemberHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvswitchAllOf) GetMemberHostsOk() (*[]VirtualizationIweHostRelationship, bool) {
	if o == nil || o.MemberHosts == nil {
		return nil, false
	}
	return &o.MemberHosts, true
}

// HasMemberHosts returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitchAllOf) HasMemberHosts() bool {
	if o != nil && o.MemberHosts != nil {
		return true
	}

	return false
}

// SetMemberHosts gets a reference to the given []VirtualizationIweHostRelationship and assigns it to the MemberHosts field.
func (o *VirtualizationIweDvswitchAllOf) SetMemberHosts(v []VirtualizationIweHostRelationship) {
	o.MemberHosts = v
}

// GetMemberVswitches returns the MemberVswitches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvswitchAllOf) GetMemberVswitches() []VirtualizationIweHostVswitchRelationship {
	if o == nil {
		var ret []VirtualizationIweHostVswitchRelationship
		return ret
	}
	return o.MemberVswitches
}

// GetMemberVswitchesOk returns a tuple with the MemberVswitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvswitchAllOf) GetMemberVswitchesOk() (*[]VirtualizationIweHostVswitchRelationship, bool) {
	if o == nil || o.MemberVswitches == nil {
		return nil, false
	}
	return &o.MemberVswitches, true
}

// HasMemberVswitches returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitchAllOf) HasMemberVswitches() bool {
	if o != nil && o.MemberVswitches != nil {
		return true
	}

	return false
}

// SetMemberVswitches gets a reference to the given []VirtualizationIweHostVswitchRelationship and assigns it to the MemberVswitches field.
func (o *VirtualizationIweDvswitchAllOf) SetMemberVswitches(v []VirtualizationIweHostVswitchRelationship) {
	o.MemberVswitches = v
}

func (o VirtualizationIweDvswitchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.LastHostname != nil {
		toSerialize["LastHostname"] = o.LastHostname
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.MemberHosts != nil {
		toSerialize["MemberHosts"] = o.MemberHosts
	}
	if o.MemberVswitches != nil {
		toSerialize["MemberVswitches"] = o.MemberVswitches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweDvswitchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationIweDvswitchAllOf := _VirtualizationIweDvswitchAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationIweDvswitchAllOf); err == nil {
		*o = VirtualizationIweDvswitchAllOf(varVirtualizationIweDvswitchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "LastHostname")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "MemberHosts")
		delete(additionalProperties, "MemberVswitches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationIweDvswitchAllOf struct {
	value *VirtualizationIweDvswitchAllOf
	isSet bool
}

func (v NullableVirtualizationIweDvswitchAllOf) Get() *VirtualizationIweDvswitchAllOf {
	return v.value
}

func (v *NullableVirtualizationIweDvswitchAllOf) Set(val *VirtualizationIweDvswitchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweDvswitchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweDvswitchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweDvswitchAllOf(val *VirtualizationIweDvswitchAllOf) *NullableVirtualizationIweDvswitchAllOf {
	return &NullableVirtualizationIweDvswitchAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationIweDvswitchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweDvswitchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
