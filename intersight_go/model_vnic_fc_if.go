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

// VnicFcIf Virtual Fibre Channel Interface.
type VnicFcIf struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the virtual fibre channel interface.
	Name *string `json:"Name,omitempty"`
	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
	Order *int64 `json:"Order,omitempty"`
	// Enables retention of LUN ID associations in memory until they are manually cleared.
	PersistentBindings *bool                         `json:"PersistentBindings,omitempty"`
	Placement          NullableVnicPlacementSettings `json:"Placement,omitempty"`
	// The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
	StaticWwpnAddress *string `json:"StaticWwpnAddress,omitempty"`
	// VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * `fc-initiator` - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * `fc-nvme-initiator` - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * `fc-nvme-target` - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * `fc-target` - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver.
	Type *string `json:"Type,omitempty"`
	// This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA.
	VifId *int64 `json:"VifId,omitempty"`
	// The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy.
	Wwpn *string `json:"Wwpn,omitempty"`
	// Type of allocation selected to assign a WWPN address to the vhba. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
	WwpnAddressType       *string                                  `json:"WwpnAddressType,omitempty"`
	FcAdapterPolicy       *VnicFcAdapterPolicyRelationship         `json:"FcAdapterPolicy,omitempty"`
	FcNetworkPolicy       *VnicFcNetworkPolicyRelationship         `json:"FcNetworkPolicy,omitempty"`
	FcQosPolicy           *VnicFcQosPolicyRelationship             `json:"FcQosPolicy,omitempty"`
	Profile               *PolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	SanConnectivityPolicy *VnicSanConnectivityPolicyRelationship   `json:"SanConnectivityPolicy,omitempty"`
	ScpVhba               *VnicFcIfRelationship                    `json:"ScpVhba,omitempty"`
	// An array of relationships to vnicFcIf resources.
	SpVhbas              []VnicFcIfRelationship   `json:"SpVhbas,omitempty"`
	WwpnLease            *FcpoolLeaseRelationship `json:"WwpnLease,omitempty"`
	WwpnPool             *FcpoolPoolRelationship  `json:"WwpnPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcIf VnicFcIf

// NewVnicFcIf instantiates a new VnicFcIf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcIf(classId string, objectType string) *VnicFcIf {
	this := VnicFcIf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "fc-initiator"
	this.Type = &type_
	var wwpnAddressType string = "POOL"
	this.WwpnAddressType = &wwpnAddressType
	return &this
}

// NewVnicFcIfWithDefaults instantiates a new VnicFcIf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcIfWithDefaults() *VnicFcIf {
	this := VnicFcIf{}
	var classId string = "vnic.FcIf"
	this.ClassId = classId
	var objectType string = "vnic.FcIf"
	this.ObjectType = objectType
	var type_ string = "fc-initiator"
	this.Type = &type_
	var wwpnAddressType string = "POOL"
	this.WwpnAddressType = &wwpnAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcIf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcIf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcIf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcIf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicFcIf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicFcIf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicFcIf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VnicFcIf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VnicFcIf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VnicFcIf) SetOrder(v int64) {
	o.Order = &v
}

// GetPersistentBindings returns the PersistentBindings field value if set, zero value otherwise.
func (o *VnicFcIf) GetPersistentBindings() bool {
	if o == nil || o.PersistentBindings == nil {
		var ret bool
		return ret
	}
	return *o.PersistentBindings
}

// GetPersistentBindingsOk returns a tuple with the PersistentBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetPersistentBindingsOk() (*bool, bool) {
	if o == nil || o.PersistentBindings == nil {
		return nil, false
	}
	return o.PersistentBindings, true
}

// HasPersistentBindings returns a boolean if a field has been set.
func (o *VnicFcIf) HasPersistentBindings() bool {
	if o != nil && o.PersistentBindings != nil {
		return true
	}

	return false
}

// SetPersistentBindings gets a reference to the given bool and assigns it to the PersistentBindings field.
func (o *VnicFcIf) SetPersistentBindings(v bool) {
	o.PersistentBindings = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetPlacement() VnicPlacementSettings {
	if o == nil || o.Placement.Get() == nil {
		var ret VnicPlacementSettings
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetPlacementOk() (*VnicPlacementSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *VnicFcIf) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableVnicPlacementSettings and assigns it to the Placement field.
func (o *VnicFcIf) SetPlacement(v VnicPlacementSettings) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *VnicFcIf) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *VnicFcIf) UnsetPlacement() {
	o.Placement.Unset()
}

// GetStaticWwpnAddress returns the StaticWwpnAddress field value if set, zero value otherwise.
func (o *VnicFcIf) GetStaticWwpnAddress() string {
	if o == nil || o.StaticWwpnAddress == nil {
		var ret string
		return ret
	}
	return *o.StaticWwpnAddress
}

// GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetStaticWwpnAddressOk() (*string, bool) {
	if o == nil || o.StaticWwpnAddress == nil {
		return nil, false
	}
	return o.StaticWwpnAddress, true
}

// HasStaticWwpnAddress returns a boolean if a field has been set.
func (o *VnicFcIf) HasStaticWwpnAddress() bool {
	if o != nil && o.StaticWwpnAddress != nil {
		return true
	}

	return false
}

// SetStaticWwpnAddress gets a reference to the given string and assigns it to the StaticWwpnAddress field.
func (o *VnicFcIf) SetStaticWwpnAddress(v string) {
	o.StaticWwpnAddress = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VnicFcIf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VnicFcIf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VnicFcIf) SetType(v string) {
	o.Type = &v
}

// GetVifId returns the VifId field value if set, zero value otherwise.
func (o *VnicFcIf) GetVifId() int64 {
	if o == nil || o.VifId == nil {
		var ret int64
		return ret
	}
	return *o.VifId
}

// GetVifIdOk returns a tuple with the VifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetVifIdOk() (*int64, bool) {
	if o == nil || o.VifId == nil {
		return nil, false
	}
	return o.VifId, true
}

// HasVifId returns a boolean if a field has been set.
func (o *VnicFcIf) HasVifId() bool {
	if o != nil && o.VifId != nil {
		return true
	}

	return false
}

// SetVifId gets a reference to the given int64 and assigns it to the VifId field.
func (o *VnicFcIf) SetVifId(v int64) {
	o.VifId = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *VnicFcIf) SetWwpn(v string) {
	o.Wwpn = &v
}

// GetWwpnAddressType returns the WwpnAddressType field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpnAddressType() string {
	if o == nil || o.WwpnAddressType == nil {
		var ret string
		return ret
	}
	return *o.WwpnAddressType
}

// GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnAddressTypeOk() (*string, bool) {
	if o == nil || o.WwpnAddressType == nil {
		return nil, false
	}
	return o.WwpnAddressType, true
}

// HasWwpnAddressType returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpnAddressType() bool {
	if o != nil && o.WwpnAddressType != nil {
		return true
	}

	return false
}

// SetWwpnAddressType gets a reference to the given string and assigns it to the WwpnAddressType field.
func (o *VnicFcIf) SetWwpnAddressType(v string) {
	o.WwpnAddressType = &v
}

// GetFcAdapterPolicy returns the FcAdapterPolicy field value if set, zero value otherwise.
func (o *VnicFcIf) GetFcAdapterPolicy() VnicFcAdapterPolicyRelationship {
	if o == nil || o.FcAdapterPolicy == nil {
		var ret VnicFcAdapterPolicyRelationship
		return ret
	}
	return *o.FcAdapterPolicy
}

// GetFcAdapterPolicyOk returns a tuple with the FcAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetFcAdapterPolicyOk() (*VnicFcAdapterPolicyRelationship, bool) {
	if o == nil || o.FcAdapterPolicy == nil {
		return nil, false
	}
	return o.FcAdapterPolicy, true
}

// HasFcAdapterPolicy returns a boolean if a field has been set.
func (o *VnicFcIf) HasFcAdapterPolicy() bool {
	if o != nil && o.FcAdapterPolicy != nil {
		return true
	}

	return false
}

// SetFcAdapterPolicy gets a reference to the given VnicFcAdapterPolicyRelationship and assigns it to the FcAdapterPolicy field.
func (o *VnicFcIf) SetFcAdapterPolicy(v VnicFcAdapterPolicyRelationship) {
	o.FcAdapterPolicy = &v
}

// GetFcNetworkPolicy returns the FcNetworkPolicy field value if set, zero value otherwise.
func (o *VnicFcIf) GetFcNetworkPolicy() VnicFcNetworkPolicyRelationship {
	if o == nil || o.FcNetworkPolicy == nil {
		var ret VnicFcNetworkPolicyRelationship
		return ret
	}
	return *o.FcNetworkPolicy
}

// GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetFcNetworkPolicyOk() (*VnicFcNetworkPolicyRelationship, bool) {
	if o == nil || o.FcNetworkPolicy == nil {
		return nil, false
	}
	return o.FcNetworkPolicy, true
}

// HasFcNetworkPolicy returns a boolean if a field has been set.
func (o *VnicFcIf) HasFcNetworkPolicy() bool {
	if o != nil && o.FcNetworkPolicy != nil {
		return true
	}

	return false
}

// SetFcNetworkPolicy gets a reference to the given VnicFcNetworkPolicyRelationship and assigns it to the FcNetworkPolicy field.
func (o *VnicFcIf) SetFcNetworkPolicy(v VnicFcNetworkPolicyRelationship) {
	o.FcNetworkPolicy = &v
}

// GetFcQosPolicy returns the FcQosPolicy field value if set, zero value otherwise.
func (o *VnicFcIf) GetFcQosPolicy() VnicFcQosPolicyRelationship {
	if o == nil || o.FcQosPolicy == nil {
		var ret VnicFcQosPolicyRelationship
		return ret
	}
	return *o.FcQosPolicy
}

// GetFcQosPolicyOk returns a tuple with the FcQosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetFcQosPolicyOk() (*VnicFcQosPolicyRelationship, bool) {
	if o == nil || o.FcQosPolicy == nil {
		return nil, false
	}
	return o.FcQosPolicy, true
}

// HasFcQosPolicy returns a boolean if a field has been set.
func (o *VnicFcIf) HasFcQosPolicy() bool {
	if o != nil && o.FcQosPolicy != nil {
		return true
	}

	return false
}

// SetFcQosPolicy gets a reference to the given VnicFcQosPolicyRelationship and assigns it to the FcQosPolicy field.
func (o *VnicFcIf) SetFcQosPolicy(v VnicFcQosPolicyRelationship) {
	o.FcQosPolicy = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VnicFcIf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicFcIf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicFcIf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile = &v
}

// GetSanConnectivityPolicy returns the SanConnectivityPolicy field value if set, zero value otherwise.
func (o *VnicFcIf) GetSanConnectivityPolicy() VnicSanConnectivityPolicyRelationship {
	if o == nil || o.SanConnectivityPolicy == nil {
		var ret VnicSanConnectivityPolicyRelationship
		return ret
	}
	return *o.SanConnectivityPolicy
}

// GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyRelationship, bool) {
	if o == nil || o.SanConnectivityPolicy == nil {
		return nil, false
	}
	return o.SanConnectivityPolicy, true
}

// HasSanConnectivityPolicy returns a boolean if a field has been set.
func (o *VnicFcIf) HasSanConnectivityPolicy() bool {
	if o != nil && o.SanConnectivityPolicy != nil {
		return true
	}

	return false
}

// SetSanConnectivityPolicy gets a reference to the given VnicSanConnectivityPolicyRelationship and assigns it to the SanConnectivityPolicy field.
func (o *VnicFcIf) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyRelationship) {
	o.SanConnectivityPolicy = &v
}

// GetScpVhba returns the ScpVhba field value if set, zero value otherwise.
func (o *VnicFcIf) GetScpVhba() VnicFcIfRelationship {
	if o == nil || o.ScpVhba == nil {
		var ret VnicFcIfRelationship
		return ret
	}
	return *o.ScpVhba
}

// GetScpVhbaOk returns a tuple with the ScpVhba field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetScpVhbaOk() (*VnicFcIfRelationship, bool) {
	if o == nil || o.ScpVhba == nil {
		return nil, false
	}
	return o.ScpVhba, true
}

// HasScpVhba returns a boolean if a field has been set.
func (o *VnicFcIf) HasScpVhba() bool {
	if o != nil && o.ScpVhba != nil {
		return true
	}

	return false
}

// SetScpVhba gets a reference to the given VnicFcIfRelationship and assigns it to the ScpVhba field.
func (o *VnicFcIf) SetScpVhba(v VnicFcIfRelationship) {
	o.ScpVhba = &v
}

// GetSpVhbas returns the SpVhbas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetSpVhbas() []VnicFcIfRelationship {
	if o == nil {
		var ret []VnicFcIfRelationship
		return ret
	}
	return o.SpVhbas
}

// GetSpVhbasOk returns a tuple with the SpVhbas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetSpVhbasOk() (*[]VnicFcIfRelationship, bool) {
	if o == nil || o.SpVhbas == nil {
		return nil, false
	}
	return &o.SpVhbas, true
}

// HasSpVhbas returns a boolean if a field has been set.
func (o *VnicFcIf) HasSpVhbas() bool {
	if o != nil && o.SpVhbas != nil {
		return true
	}

	return false
}

// SetSpVhbas gets a reference to the given []VnicFcIfRelationship and assigns it to the SpVhbas field.
func (o *VnicFcIf) SetSpVhbas(v []VnicFcIfRelationship) {
	o.SpVhbas = v
}

// GetWwpnLease returns the WwpnLease field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpnLease() FcpoolLeaseRelationship {
	if o == nil || o.WwpnLease == nil {
		var ret FcpoolLeaseRelationship
		return ret
	}
	return *o.WwpnLease
}

// GetWwpnLeaseOk returns a tuple with the WwpnLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool) {
	if o == nil || o.WwpnLease == nil {
		return nil, false
	}
	return o.WwpnLease, true
}

// HasWwpnLease returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpnLease() bool {
	if o != nil && o.WwpnLease != nil {
		return true
	}

	return false
}

// SetWwpnLease gets a reference to the given FcpoolLeaseRelationship and assigns it to the WwpnLease field.
func (o *VnicFcIf) SetWwpnLease(v FcpoolLeaseRelationship) {
	o.WwpnLease = &v
}

// GetWwpnPool returns the WwpnPool field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpnPool() FcpoolPoolRelationship {
	if o == nil || o.WwpnPool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.WwpnPool
}

// GetWwpnPoolOk returns a tuple with the WwpnPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.WwpnPool == nil {
		return nil, false
	}
	return o.WwpnPool, true
}

// HasWwpnPool returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpnPool() bool {
	if o != nil && o.WwpnPool != nil {
		return true
	}

	return false
}

// SetWwpnPool gets a reference to the given FcpoolPoolRelationship and assigns it to the WwpnPool field.
func (o *VnicFcIf) SetWwpnPool(v FcpoolPoolRelationship) {
	o.WwpnPool = &v
}

func (o VnicFcIf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.PersistentBindings != nil {
		toSerialize["PersistentBindings"] = o.PersistentBindings
	}
	if o.Placement.IsSet() {
		toSerialize["Placement"] = o.Placement.Get()
	}
	if o.StaticWwpnAddress != nil {
		toSerialize["StaticWwpnAddress"] = o.StaticWwpnAddress
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VifId != nil {
		toSerialize["VifId"] = o.VifId
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}
	if o.WwpnAddressType != nil {
		toSerialize["WwpnAddressType"] = o.WwpnAddressType
	}
	if o.FcAdapterPolicy != nil {
		toSerialize["FcAdapterPolicy"] = o.FcAdapterPolicy
	}
	if o.FcNetworkPolicy != nil {
		toSerialize["FcNetworkPolicy"] = o.FcNetworkPolicy
	}
	if o.FcQosPolicy != nil {
		toSerialize["FcQosPolicy"] = o.FcQosPolicy
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.SanConnectivityPolicy != nil {
		toSerialize["SanConnectivityPolicy"] = o.SanConnectivityPolicy
	}
	if o.ScpVhba != nil {
		toSerialize["ScpVhba"] = o.ScpVhba
	}
	if o.SpVhbas != nil {
		toSerialize["SpVhbas"] = o.SpVhbas
	}
	if o.WwpnLease != nil {
		toSerialize["WwpnLease"] = o.WwpnLease
	}
	if o.WwpnPool != nil {
		toSerialize["WwpnPool"] = o.WwpnPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcIf) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcIfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the virtual fibre channel interface.
		Name *string `json:"Name,omitempty"`
		// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
		Order *int64 `json:"Order,omitempty"`
		// Enables retention of LUN ID associations in memory until they are manually cleared.
		PersistentBindings *bool                         `json:"PersistentBindings,omitempty"`
		Placement          NullableVnicPlacementSettings `json:"Placement,omitempty"`
		// The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
		StaticWwpnAddress *string `json:"StaticWwpnAddress,omitempty"`
		// VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. * `fc-initiator` - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems. * `fc-nvme-initiator` - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. * `fc-nvme-target` - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver. * `fc-target` - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver.
		Type *string `json:"Type,omitempty"`
		// This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA.
		VifId *int64 `json:"VifId,omitempty"`
		// The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy.
		Wwpn *string `json:"Wwpn,omitempty"`
		// Type of allocation selected to assign a WWPN address to the vhba. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
		WwpnAddressType       *string                                  `json:"WwpnAddressType,omitempty"`
		FcAdapterPolicy       *VnicFcAdapterPolicyRelationship         `json:"FcAdapterPolicy,omitempty"`
		FcNetworkPolicy       *VnicFcNetworkPolicyRelationship         `json:"FcNetworkPolicy,omitempty"`
		FcQosPolicy           *VnicFcQosPolicyRelationship             `json:"FcQosPolicy,omitempty"`
		Profile               *PolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
		SanConnectivityPolicy *VnicSanConnectivityPolicyRelationship   `json:"SanConnectivityPolicy,omitempty"`
		ScpVhba               *VnicFcIfRelationship                    `json:"ScpVhba,omitempty"`
		// An array of relationships to vnicFcIf resources.
		SpVhbas   []VnicFcIfRelationship   `json:"SpVhbas,omitempty"`
		WwpnLease *FcpoolLeaseRelationship `json:"WwpnLease,omitempty"`
		WwpnPool  *FcpoolPoolRelationship  `json:"WwpnPool,omitempty"`
	}

	varVnicFcIfWithoutEmbeddedStruct := VnicFcIfWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcIfWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcIf := _VnicFcIf{}
		varVnicFcIf.ClassId = varVnicFcIfWithoutEmbeddedStruct.ClassId
		varVnicFcIf.ObjectType = varVnicFcIfWithoutEmbeddedStruct.ObjectType
		varVnicFcIf.Name = varVnicFcIfWithoutEmbeddedStruct.Name
		varVnicFcIf.Order = varVnicFcIfWithoutEmbeddedStruct.Order
		varVnicFcIf.PersistentBindings = varVnicFcIfWithoutEmbeddedStruct.PersistentBindings
		varVnicFcIf.Placement = varVnicFcIfWithoutEmbeddedStruct.Placement
		varVnicFcIf.StaticWwpnAddress = varVnicFcIfWithoutEmbeddedStruct.StaticWwpnAddress
		varVnicFcIf.Type = varVnicFcIfWithoutEmbeddedStruct.Type
		varVnicFcIf.VifId = varVnicFcIfWithoutEmbeddedStruct.VifId
		varVnicFcIf.Wwpn = varVnicFcIfWithoutEmbeddedStruct.Wwpn
		varVnicFcIf.WwpnAddressType = varVnicFcIfWithoutEmbeddedStruct.WwpnAddressType
		varVnicFcIf.FcAdapterPolicy = varVnicFcIfWithoutEmbeddedStruct.FcAdapterPolicy
		varVnicFcIf.FcNetworkPolicy = varVnicFcIfWithoutEmbeddedStruct.FcNetworkPolicy
		varVnicFcIf.FcQosPolicy = varVnicFcIfWithoutEmbeddedStruct.FcQosPolicy
		varVnicFcIf.Profile = varVnicFcIfWithoutEmbeddedStruct.Profile
		varVnicFcIf.SanConnectivityPolicy = varVnicFcIfWithoutEmbeddedStruct.SanConnectivityPolicy
		varVnicFcIf.ScpVhba = varVnicFcIfWithoutEmbeddedStruct.ScpVhba
		varVnicFcIf.SpVhbas = varVnicFcIfWithoutEmbeddedStruct.SpVhbas
		varVnicFcIf.WwpnLease = varVnicFcIfWithoutEmbeddedStruct.WwpnLease
		varVnicFcIf.WwpnPool = varVnicFcIfWithoutEmbeddedStruct.WwpnPool
		*o = VnicFcIf(varVnicFcIf)
	} else {
		return err
	}

	varVnicFcIf := _VnicFcIf{}

	err = json.Unmarshal(bytes, &varVnicFcIf)
	if err == nil {
		o.MoBaseMo = varVnicFcIf.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "PersistentBindings")
		delete(additionalProperties, "Placement")
		delete(additionalProperties, "StaticWwpnAddress")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VifId")
		delete(additionalProperties, "Wwpn")
		delete(additionalProperties, "WwpnAddressType")
		delete(additionalProperties, "FcAdapterPolicy")
		delete(additionalProperties, "FcNetworkPolicy")
		delete(additionalProperties, "FcQosPolicy")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "SanConnectivityPolicy")
		delete(additionalProperties, "ScpVhba")
		delete(additionalProperties, "SpVhbas")
		delete(additionalProperties, "WwpnLease")
		delete(additionalProperties, "WwpnPool")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableVnicFcIf struct {
	value *VnicFcIf
	isSet bool
}

func (v NullableVnicFcIf) Get() *VnicFcIf {
	return v.value
}

func (v *NullableVnicFcIf) Set(val *VnicFcIf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcIf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcIf(val *VnicFcIf) *NullableVnicFcIf {
	return &NullableVnicFcIf{value: val, isSet: true}
}

func (v NullableVnicFcIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
