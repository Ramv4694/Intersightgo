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
	"reflect"
	"strings"
)

// UuidpoolUuidLease UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
type UuidpoolUuidLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID Prefix+Suffix numbers.
	Uuid                 *string                         `json:"Uuid,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
	Pool                 *UuidpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *UuidpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *UuidpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolUuidLease UuidpoolUuidLease

// NewUuidpoolUuidLease instantiates a new UuidpoolUuidLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolUuidLease(classId string, objectType string) *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUuidpoolUuidLeaseWithDefaults instantiates a new UuidpoolUuidLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolUuidLeaseWithDefaults() *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	var classId string = "uuidpool.UuidLease"
	this.ClassId = classId
	var objectType string = "uuidpool.UuidLease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolUuidLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolUuidLease) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolUuidLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolUuidLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolUuidLease) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolUuidLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolUuidLease) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetPoolMember() UuidpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret UuidpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given UuidpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *UuidpoolUuidLease) SetPoolMember(v UuidpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetUniverse() UuidpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret UuidpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetUniverseOk() (*UuidpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given UuidpoolUniverseRelationship and assigns it to the Universe field.
func (o *UuidpoolUuidLease) SetUniverse(v UuidpoolUniverseRelationship) {
	o.Universe = &v
}

func (o UuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return []byte{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return []byte{}, errPoolAbstractLease
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
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

func (o *UuidpoolUuidLease) UnmarshalJSON(bytes []byte) (err error) {
	type UuidpoolUuidLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UUID Prefix+Suffix numbers.
		Uuid             *string                         `json:"Uuid,omitempty"`
		AssignedToEntity *MoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
		Pool             *UuidpoolPoolRelationship       `json:"Pool,omitempty"`
		PoolMember       *UuidpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
		Universe         *UuidpoolUniverseRelationship   `json:"Universe,omitempty"`
	}

	varUuidpoolUuidLeaseWithoutEmbeddedStruct := UuidpoolUuidLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUuidpoolUuidLeaseWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolUuidLease := _UuidpoolUuidLease{}
		varUuidpoolUuidLease.ClassId = varUuidpoolUuidLeaseWithoutEmbeddedStruct.ClassId
		varUuidpoolUuidLease.ObjectType = varUuidpoolUuidLeaseWithoutEmbeddedStruct.ObjectType
		varUuidpoolUuidLease.Uuid = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Uuid
		varUuidpoolUuidLease.AssignedToEntity = varUuidpoolUuidLeaseWithoutEmbeddedStruct.AssignedToEntity
		varUuidpoolUuidLease.Pool = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Pool
		varUuidpoolUuidLease.PoolMember = varUuidpoolUuidLeaseWithoutEmbeddedStruct.PoolMember
		varUuidpoolUuidLease.Universe = varUuidpoolUuidLeaseWithoutEmbeddedStruct.Universe
		*o = UuidpoolUuidLease(varUuidpoolUuidLease)
	} else {
		return err
	}

	varUuidpoolUuidLease := _UuidpoolUuidLease{}

	err = json.Unmarshal(bytes, &varUuidpoolUuidLease)
	if err == nil {
		o.PoolAbstractLease = varUuidpoolUuidLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

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

type NullableUuidpoolUuidLease struct {
	value *UuidpoolUuidLease
	isSet bool
}

func (v NullableUuidpoolUuidLease) Get() *UuidpoolUuidLease {
	return v.value
}

func (v *NullableUuidpoolUuidLease) Set(val *UuidpoolUuidLease) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolUuidLease) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolUuidLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolUuidLease(val *UuidpoolUuidLease) *NullableUuidpoolUuidLease {
	return &NullableUuidpoolUuidLease{value: val, isSet: true}
}

func (v NullableUuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolUuidLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
