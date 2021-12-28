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

// VirtualizationVmwarePhysicalNetworkInterfaceAllOf Definition of the list of properties defined in 'virtualization.VmwarePhysicalNetworkInterface', excluding properties defined in parent classes.
type VirtualizationVmwarePhysicalNetworkInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Driver type associated with physical network interface.
	Driver *string `json:"Driver,omitempty"`
	// Link speed of the physical network interface.
	LinkSpeed *int32 `json:"LinkSpeed,omitempty"`
	// Standard MAC address assigned to physical network interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// PCI info for physical network interface.
	Pci *string `json:"Pci,omitempty"`
	// Switch associated with the physical network interface.
	SwitchName           *string                               `json:"SwitchName,omitempty"`
	Host                 *VirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwarePhysicalNetworkInterfaceAllOf VirtualizationVmwarePhysicalNetworkInterfaceAllOf

// NewVirtualizationVmwarePhysicalNetworkInterfaceAllOf instantiates a new VirtualizationVmwarePhysicalNetworkInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwarePhysicalNetworkInterfaceAllOf(classId string, objectType string) *VirtualizationVmwarePhysicalNetworkInterfaceAllOf {
	this := VirtualizationVmwarePhysicalNetworkInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwarePhysicalNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationVmwarePhysicalNetworkInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwarePhysicalNetworkInterfaceAllOfWithDefaults() *VirtualizationVmwarePhysicalNetworkInterfaceAllOf {
	this := VirtualizationVmwarePhysicalNetworkInterfaceAllOf{}
	var classId string = "virtualization.VmwarePhysicalNetworkInterface"
	this.ClassId = classId
	var objectType string = "virtualization.VmwarePhysicalNetworkInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetDriver(v string) {
	o.Driver = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetLinkSpeed() int32 {
	if o == nil || o.LinkSpeed == nil {
		var ret int32
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetLinkSpeedOk() (*int32, bool) {
	if o == nil || o.LinkSpeed == nil {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasLinkSpeed() bool {
	if o != nil && o.LinkSpeed != nil {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given int32 and assigns it to the LinkSpeed field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetLinkSpeed(v int32) {
	o.LinkSpeed = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPci returns the Pci field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetPci() string {
	if o == nil || o.Pci == nil {
		var ret string
		return ret
	}
	return *o.Pci
}

// GetPciOk returns a tuple with the Pci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetPciOk() (*string, bool) {
	if o == nil || o.Pci == nil {
		return nil, false
	}
	return o.Pci, true
}

// HasPci returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasPci() bool {
	if o != nil && o.Pci != nil {
		return true
	}

	return false
}

// SetPci gets a reference to the given string and assigns it to the Pci field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetPci(v string) {
	o.Pci = &v
}

// GetSwitchName returns the SwitchName field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetSwitchName() string {
	if o == nil || o.SwitchName == nil {
		var ret string
		return ret
	}
	return *o.SwitchName
}

// GetSwitchNameOk returns a tuple with the SwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetSwitchNameOk() (*string, bool) {
	if o == nil || o.SwitchName == nil {
		return nil, false
	}
	return o.SwitchName, true
}

// HasSwitchName returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasSwitchName() bool {
	if o != nil && o.SwitchName != nil {
		return true
	}

	return false
}

// SetSwitchName gets a reference to the given string and assigns it to the SwitchName field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetSwitchName(v string) {
	o.SwitchName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

func (o VirtualizationVmwarePhysicalNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Driver != nil {
		toSerialize["Driver"] = o.Driver
	}
	if o.LinkSpeed != nil {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Pci != nil {
		toSerialize["Pci"] = o.Pci
	}
	if o.SwitchName != nil {
		toSerialize["SwitchName"] = o.SwitchName
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwarePhysicalNetworkInterfaceAllOf := _VirtualizationVmwarePhysicalNetworkInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwarePhysicalNetworkInterfaceAllOf); err == nil {
		*o = VirtualizationVmwarePhysicalNetworkInterfaceAllOf(varVirtualizationVmwarePhysicalNetworkInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Driver")
		delete(additionalProperties, "LinkSpeed")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Pci")
		delete(additionalProperties, "SwitchName")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf struct {
	value *VirtualizationVmwarePhysicalNetworkInterfaceAllOf
	isSet bool
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) Get() *VirtualizationVmwarePhysicalNetworkInterfaceAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) Set(val *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf(val *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) *NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf {
	return &NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
