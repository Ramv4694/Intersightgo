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

// FabricSwitchControlPolicy A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
type FabricSwitchControlPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or Disable Ethernet End Host Switching Mode. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	EthernetSwitchingMode *string `json:"EthernetSwitchingMode,omitempty"`
	// Enable or Disable FC End Host Switching Mode. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	FcSwitchingMode  *string                          `json:"FcSwitchingMode,omitempty"`
	MacAgingSettings NullableFabricMacAgingSettings   `json:"MacAgingSettings,omitempty"`
	UdldSettings     NullableFabricUdldGlobalSettings `json:"UdldSettings,omitempty"`
	// To enable or disable the VLAN port count optimization.
	VlanPortOptimizationEnabled *bool                                 `json:"VlanPortOptimizationEnabled,omitempty"`
	Organization                *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSwitchControlPolicy FabricSwitchControlPolicy

// NewFabricSwitchControlPolicy instantiates a new FabricSwitchControlPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchControlPolicy(classId string, objectType string) *FabricSwitchControlPolicy {
	this := FabricSwitchControlPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ethernetSwitchingMode string = "end-host"
	this.EthernetSwitchingMode = &ethernetSwitchingMode
	var fcSwitchingMode string = "end-host"
	this.FcSwitchingMode = &fcSwitchingMode
	var vlanPortOptimizationEnabled bool = false
	this.VlanPortOptimizationEnabled = &vlanPortOptimizationEnabled
	return &this
}

// NewFabricSwitchControlPolicyWithDefaults instantiates a new FabricSwitchControlPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchControlPolicyWithDefaults() *FabricSwitchControlPolicy {
	this := FabricSwitchControlPolicy{}
	var classId string = "fabric.SwitchControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.SwitchControlPolicy"
	this.ObjectType = objectType
	var ethernetSwitchingMode string = "end-host"
	this.EthernetSwitchingMode = &ethernetSwitchingMode
	var fcSwitchingMode string = "end-host"
	this.FcSwitchingMode = &fcSwitchingMode
	var vlanPortOptimizationEnabled bool = false
	this.VlanPortOptimizationEnabled = &vlanPortOptimizationEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchControlPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchControlPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchControlPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchControlPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEthernetSwitchingMode returns the EthernetSwitchingMode field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicy) GetEthernetSwitchingMode() string {
	if o == nil || o.EthernetSwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.EthernetSwitchingMode
}

// GetEthernetSwitchingModeOk returns a tuple with the EthernetSwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetEthernetSwitchingModeOk() (*string, bool) {
	if o == nil || o.EthernetSwitchingMode == nil {
		return nil, false
	}
	return o.EthernetSwitchingMode, true
}

// HasEthernetSwitchingMode returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasEthernetSwitchingMode() bool {
	if o != nil && o.EthernetSwitchingMode != nil {
		return true
	}

	return false
}

// SetEthernetSwitchingMode gets a reference to the given string and assigns it to the EthernetSwitchingMode field.
func (o *FabricSwitchControlPolicy) SetEthernetSwitchingMode(v string) {
	o.EthernetSwitchingMode = &v
}

// GetFcSwitchingMode returns the FcSwitchingMode field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicy) GetFcSwitchingMode() string {
	if o == nil || o.FcSwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.FcSwitchingMode
}

// GetFcSwitchingModeOk returns a tuple with the FcSwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetFcSwitchingModeOk() (*string, bool) {
	if o == nil || o.FcSwitchingMode == nil {
		return nil, false
	}
	return o.FcSwitchingMode, true
}

// HasFcSwitchingMode returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasFcSwitchingMode() bool {
	if o != nil && o.FcSwitchingMode != nil {
		return true
	}

	return false
}

// SetFcSwitchingMode gets a reference to the given string and assigns it to the FcSwitchingMode field.
func (o *FabricSwitchControlPolicy) SetFcSwitchingMode(v string) {
	o.FcSwitchingMode = &v
}

// GetMacAgingSettings returns the MacAgingSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicy) GetMacAgingSettings() FabricMacAgingSettings {
	if o == nil || o.MacAgingSettings.Get() == nil {
		var ret FabricMacAgingSettings
		return ret
	}
	return *o.MacAgingSettings.Get()
}

// GetMacAgingSettingsOk returns a tuple with the MacAgingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicy) GetMacAgingSettingsOk() (*FabricMacAgingSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAgingSettings.Get(), o.MacAgingSettings.IsSet()
}

// HasMacAgingSettings returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasMacAgingSettings() bool {
	if o != nil && o.MacAgingSettings.IsSet() {
		return true
	}

	return false
}

// SetMacAgingSettings gets a reference to the given NullableFabricMacAgingSettings and assigns it to the MacAgingSettings field.
func (o *FabricSwitchControlPolicy) SetMacAgingSettings(v FabricMacAgingSettings) {
	o.MacAgingSettings.Set(&v)
}

// SetMacAgingSettingsNil sets the value for MacAgingSettings to be an explicit nil
func (o *FabricSwitchControlPolicy) SetMacAgingSettingsNil() {
	o.MacAgingSettings.Set(nil)
}

// UnsetMacAgingSettings ensures that no value is present for MacAgingSettings, not even an explicit nil
func (o *FabricSwitchControlPolicy) UnsetMacAgingSettings() {
	o.MacAgingSettings.Unset()
}

// GetUdldSettings returns the UdldSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicy) GetUdldSettings() FabricUdldGlobalSettings {
	if o == nil || o.UdldSettings.Get() == nil {
		var ret FabricUdldGlobalSettings
		return ret
	}
	return *o.UdldSettings.Get()
}

// GetUdldSettingsOk returns a tuple with the UdldSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicy) GetUdldSettingsOk() (*FabricUdldGlobalSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.UdldSettings.Get(), o.UdldSettings.IsSet()
}

// HasUdldSettings returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasUdldSettings() bool {
	if o != nil && o.UdldSettings.IsSet() {
		return true
	}

	return false
}

// SetUdldSettings gets a reference to the given NullableFabricUdldGlobalSettings and assigns it to the UdldSettings field.
func (o *FabricSwitchControlPolicy) SetUdldSettings(v FabricUdldGlobalSettings) {
	o.UdldSettings.Set(&v)
}

// SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil
func (o *FabricSwitchControlPolicy) SetUdldSettingsNil() {
	o.UdldSettings.Set(nil)
}

// UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
func (o *FabricSwitchControlPolicy) UnsetUdldSettings() {
	o.UdldSettings.Unset()
}

// GetVlanPortOptimizationEnabled returns the VlanPortOptimizationEnabled field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicy) GetVlanPortOptimizationEnabled() bool {
	if o == nil || o.VlanPortOptimizationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.VlanPortOptimizationEnabled
}

// GetVlanPortOptimizationEnabledOk returns a tuple with the VlanPortOptimizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetVlanPortOptimizationEnabledOk() (*bool, bool) {
	if o == nil || o.VlanPortOptimizationEnabled == nil {
		return nil, false
	}
	return o.VlanPortOptimizationEnabled, true
}

// HasVlanPortOptimizationEnabled returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasVlanPortOptimizationEnabled() bool {
	if o != nil && o.VlanPortOptimizationEnabled != nil {
		return true
	}

	return false
}

// SetVlanPortOptimizationEnabled gets a reference to the given bool and assigns it to the VlanPortOptimizationEnabled field.
func (o *FabricSwitchControlPolicy) SetVlanPortOptimizationEnabled(v bool) {
	o.VlanPortOptimizationEnabled = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSwitchControlPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicy) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricSwitchControlPolicy) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricSwitchControlPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EthernetSwitchingMode != nil {
		toSerialize["EthernetSwitchingMode"] = o.EthernetSwitchingMode
	}
	if o.FcSwitchingMode != nil {
		toSerialize["FcSwitchingMode"] = o.FcSwitchingMode
	}
	if o.MacAgingSettings.IsSet() {
		toSerialize["MacAgingSettings"] = o.MacAgingSettings.Get()
	}
	if o.UdldSettings.IsSet() {
		toSerialize["UdldSettings"] = o.UdldSettings.Get()
	}
	if o.VlanPortOptimizationEnabled != nil {
		toSerialize["VlanPortOptimizationEnabled"] = o.VlanPortOptimizationEnabled
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSwitchControlPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricSwitchControlPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or Disable Ethernet End Host Switching Mode. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
		EthernetSwitchingMode *string `json:"EthernetSwitchingMode,omitempty"`
		// Enable or Disable FC End Host Switching Mode. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
		FcSwitchingMode  *string                          `json:"FcSwitchingMode,omitempty"`
		MacAgingSettings NullableFabricMacAgingSettings   `json:"MacAgingSettings,omitempty"`
		UdldSettings     NullableFabricUdldGlobalSettings `json:"UdldSettings,omitempty"`
		// To enable or disable the VLAN port count optimization.
		VlanPortOptimizationEnabled *bool                                 `json:"VlanPortOptimizationEnabled,omitempty"`
		Organization                *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfile resources.
		Profiles []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	}

	varFabricSwitchControlPolicyWithoutEmbeddedStruct := FabricSwitchControlPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricSwitchControlPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricSwitchControlPolicy := _FabricSwitchControlPolicy{}
		varFabricSwitchControlPolicy.ClassId = varFabricSwitchControlPolicyWithoutEmbeddedStruct.ClassId
		varFabricSwitchControlPolicy.ObjectType = varFabricSwitchControlPolicyWithoutEmbeddedStruct.ObjectType
		varFabricSwitchControlPolicy.EthernetSwitchingMode = varFabricSwitchControlPolicyWithoutEmbeddedStruct.EthernetSwitchingMode
		varFabricSwitchControlPolicy.FcSwitchingMode = varFabricSwitchControlPolicyWithoutEmbeddedStruct.FcSwitchingMode
		varFabricSwitchControlPolicy.MacAgingSettings = varFabricSwitchControlPolicyWithoutEmbeddedStruct.MacAgingSettings
		varFabricSwitchControlPolicy.UdldSettings = varFabricSwitchControlPolicyWithoutEmbeddedStruct.UdldSettings
		varFabricSwitchControlPolicy.VlanPortOptimizationEnabled = varFabricSwitchControlPolicyWithoutEmbeddedStruct.VlanPortOptimizationEnabled
		varFabricSwitchControlPolicy.Organization = varFabricSwitchControlPolicyWithoutEmbeddedStruct.Organization
		varFabricSwitchControlPolicy.Profiles = varFabricSwitchControlPolicyWithoutEmbeddedStruct.Profiles
		*o = FabricSwitchControlPolicy(varFabricSwitchControlPolicy)
	} else {
		return err
	}

	varFabricSwitchControlPolicy := _FabricSwitchControlPolicy{}

	err = json.Unmarshal(bytes, &varFabricSwitchControlPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricSwitchControlPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EthernetSwitchingMode")
		delete(additionalProperties, "FcSwitchingMode")
		delete(additionalProperties, "MacAgingSettings")
		delete(additionalProperties, "UdldSettings")
		delete(additionalProperties, "VlanPortOptimizationEnabled")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricSwitchControlPolicy struct {
	value *FabricSwitchControlPolicy
	isSet bool
}

func (v NullableFabricSwitchControlPolicy) Get() *FabricSwitchControlPolicy {
	return v.value
}

func (v *NullableFabricSwitchControlPolicy) Set(val *FabricSwitchControlPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchControlPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchControlPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchControlPolicy(val *FabricSwitchControlPolicy) *NullableFabricSwitchControlPolicy {
	return &NullableFabricSwitchControlPolicy{value: val, isSet: true}
}

func (v NullableFabricSwitchControlPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchControlPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
