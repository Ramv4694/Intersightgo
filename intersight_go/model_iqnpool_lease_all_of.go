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

// IqnpoolLeaseAllOf Definition of the list of properties defined in 'iqnpool.Lease', excluding properties defined in parent classes.
type IqnpoolLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IQN address allocated for pool-based allocation \"prefix+suffix+number\".
	IqnAddress           *string                        `json:"IqnAddress,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	Pool                 *IqnpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *IqnpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *IqnpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolLeaseAllOf IqnpoolLeaseAllOf

// NewIqnpoolLeaseAllOf instantiates a new IqnpoolLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolLeaseAllOf(classId string, objectType string) *IqnpoolLeaseAllOf {
	this := IqnpoolLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIqnpoolLeaseAllOfWithDefaults instantiates a new IqnpoolLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolLeaseAllOfWithDefaults() *IqnpoolLeaseAllOf {
	this := IqnpoolLeaseAllOf{}
	var classId string = "iqnpool.Lease"
	this.ClassId = classId
	var objectType string = "iqnpool.Lease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIqnAddress returns the IqnAddress field value if set, zero value otherwise.
func (o *IqnpoolLeaseAllOf) GetIqnAddress() string {
	if o == nil || o.IqnAddress == nil {
		var ret string
		return ret
	}
	return *o.IqnAddress
}

// GetIqnAddressOk returns a tuple with the IqnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetIqnAddressOk() (*string, bool) {
	if o == nil || o.IqnAddress == nil {
		return nil, false
	}
	return o.IqnAddress, true
}

// HasIqnAddress returns a boolean if a field has been set.
func (o *IqnpoolLeaseAllOf) HasIqnAddress() bool {
	if o != nil && o.IqnAddress != nil {
		return true
	}

	return false
}

// SetIqnAddress gets a reference to the given string and assigns it to the IqnAddress field.
func (o *IqnpoolLeaseAllOf) SetIqnAddress(v string) {
	o.IqnAddress = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IqnpoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IqnpoolLeaseAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IqnpoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IqnpoolLeaseAllOf) GetPool() IqnpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IqnpoolLeaseAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IqnpoolPoolRelationship and assigns it to the Pool field.
func (o *IqnpoolLeaseAllOf) SetPool(v IqnpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *IqnpoolLeaseAllOf) GetPoolMember() IqnpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret IqnpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IqnpoolLeaseAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given IqnpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IqnpoolLeaseAllOf) SetPoolMember(v IqnpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IqnpoolLeaseAllOf) GetUniverse() IqnpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IqnpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLeaseAllOf) GetUniverseOk() (*IqnpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IqnpoolLeaseAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IqnpoolUniverseRelationship and assigns it to the Universe field.
func (o *IqnpoolLeaseAllOf) SetUniverse(v IqnpoolUniverseRelationship) {
	o.Universe = &v
}

func (o IqnpoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IqnAddress != nil {
		toSerialize["IqnAddress"] = o.IqnAddress
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIqnpoolLeaseAllOf := _IqnpoolLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varIqnpoolLeaseAllOf); err == nil {
		*o = IqnpoolLeaseAllOf(varIqnpoolLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IqnAddress")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIqnpoolLeaseAllOf struct {
	value *IqnpoolLeaseAllOf
	isSet bool
}

func (v NullableIqnpoolLeaseAllOf) Get() *IqnpoolLeaseAllOf {
	return v.value
}

func (v *NullableIqnpoolLeaseAllOf) Set(val *IqnpoolLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolLeaseAllOf(val *IqnpoolLeaseAllOf) *NullableIqnpoolLeaseAllOf {
	return &NullableIqnpoolLeaseAllOf{value: val, isSet: true}
}

func (v NullableIqnpoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
