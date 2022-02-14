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

// StorageNetAppFcInterface NetApp FC Interface is a logical interface.
type StorageNetAppFcInterface struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// FC interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// The state of the FC interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Uuid of  NetApp FC Interface.
	Uuid            *string                        `json:"Uuid,omitempty"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppFcInterfaceEvent resources.
	Events               []StorageNetAppFcInterfaceEventRelationship `json:"Events,omitempty"`
	PhysicalPort         *StorageNetAppFcPortRelationship            `json:"PhysicalPort,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship         `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcInterface StorageNetAppFcInterface

// NewStorageNetAppFcInterface instantiates a new StorageNetAppFcInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcInterface(classId string, objectType string) *StorageNetAppFcInterface {
	this := StorageNetAppFcInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	return &this
}

// NewStorageNetAppFcInterfaceWithDefaults instantiates a new StorageNetAppFcInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcInterfaceWithDefaults() *StorageNetAppFcInterface {
	this := StorageNetAppFcInterface{}
	var classId string = "storage.NetAppFcInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcInterface"
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppFcInterface) SetEnabled(v string) {
	o.Enabled = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppFcInterface) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppFcInterface) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppFcInterface) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcInterface) GetEvents() []StorageNetAppFcInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppFcInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcInterface) GetEventsOk() (*[]StorageNetAppFcInterfaceEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppFcInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppFcInterface) SetEvents(v []StorageNetAppFcInterfaceEventRelationship) {
	o.Events = v
}

// GetPhysicalPort returns the PhysicalPort field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetPhysicalPort() StorageNetAppFcPortRelationship {
	if o == nil || o.PhysicalPort == nil {
		var ret StorageNetAppFcPortRelationship
		return ret
	}
	return *o.PhysicalPort
}

// GetPhysicalPortOk returns a tuple with the PhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetPhysicalPortOk() (*StorageNetAppFcPortRelationship, bool) {
	if o == nil || o.PhysicalPort == nil {
		return nil, false
	}
	return o.PhysicalPort, true
}

// HasPhysicalPort returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasPhysicalPort() bool {
	if o != nil && o.PhysicalPort != nil {
		return true
	}

	return false
}

// SetPhysicalPort gets a reference to the given StorageNetAppFcPortRelationship and assigns it to the PhysicalPort field.
func (o *StorageNetAppFcInterface) SetPhysicalPort(v StorageNetAppFcPortRelationship) {
	o.PhysicalPort = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppFcInterface) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppFcInterface) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppFcInterface) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppFcInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBasePhysicalPort, errStorageBasePhysicalPort := json.Marshal(o.StorageBasePhysicalPort)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	errStorageBasePhysicalPort = json.Unmarshal([]byte(serializedStorageBasePhysicalPort), &toSerialize)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.PhysicalPort != nil {
		toSerialize["PhysicalPort"] = o.PhysicalPort
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppFcInterface) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppFcInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// FC interface is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// The state of the FC interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		State *string `json:"State,omitempty"`
		// Uuid of  NetApp FC Interface.
		Uuid            *string                        `json:"Uuid,omitempty"`
		ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppFcInterfaceEvent resources.
		Events       []StorageNetAppFcInterfaceEventRelationship `json:"Events,omitempty"`
		PhysicalPort *StorageNetAppFcPortRelationship            `json:"PhysicalPort,omitempty"`
		Tenant       *StorageNetAppStorageVmRelationship         `json:"Tenant,omitempty"`
	}

	varStorageNetAppFcInterfaceWithoutEmbeddedStruct := StorageNetAppFcInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppFcInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppFcInterface := _StorageNetAppFcInterface{}
		varStorageNetAppFcInterface.ClassId = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppFcInterface.ObjectType = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppFcInterface.Enabled = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Enabled
		varStorageNetAppFcInterface.State = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.State
		varStorageNetAppFcInterface.Uuid = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Uuid
		varStorageNetAppFcInterface.ArrayController = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.ArrayController
		varStorageNetAppFcInterface.Events = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Events
		varStorageNetAppFcInterface.PhysicalPort = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.PhysicalPort
		varStorageNetAppFcInterface.Tenant = varStorageNetAppFcInterfaceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppFcInterface(varStorageNetAppFcInterface)
	} else {
		return err
	}

	varStorageNetAppFcInterface := _StorageNetAppFcInterface{}

	err = json.Unmarshal(bytes, &varStorageNetAppFcInterface)
	if err == nil {
		o.StorageBasePhysicalPort = varStorageNetAppFcInterface.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "PhysicalPort")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBasePhysicalPort := reflect.ValueOf(o.StorageBasePhysicalPort)
		for i := 0; i < reflectStorageBasePhysicalPort.Type().NumField(); i++ {
			t := reflectStorageBasePhysicalPort.Type().Field(i)

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

type NullableStorageNetAppFcInterface struct {
	value *StorageNetAppFcInterface
	isSet bool
}

func (v NullableStorageNetAppFcInterface) Get() *StorageNetAppFcInterface {
	return v.value
}

func (v *NullableStorageNetAppFcInterface) Set(val *StorageNetAppFcInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcInterface(val *StorageNetAppFcInterface) *NullableStorageNetAppFcInterface {
	return &NullableStorageNetAppFcInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppFcInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
