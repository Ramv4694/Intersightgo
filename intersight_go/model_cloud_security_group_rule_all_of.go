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

// CloudSecurityGroupRuleAllOf Definition of the list of properties defined in 'cloud.SecurityGroupRule', excluding properties defined in parent classes.
type CloudSecurityGroupRuleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action against the traffic to  the virtual machine, for example deny, permit, etc.
	Action *string `json:"Action,omitempty"`
	// Description about the security group rule.
	Description *string `json:"Description,omitempty"`
	// The end of port range for the security group rule IP protocol.When all the traffic is allowed into/from network boundary of virtual machine, both start port and end port will be zero.
	EndPort *int64 `json:"EndPort,omitempty"`
	// IP version of source CIDR (Classless Inter-Domain Routing), such as IPv4 or IPv6. * `IPv4` - Internet protocol version 4. * `IPv6` - Internet protocol version 6.
	EtherType *string `json:"EtherType,omitempty"`
	// Identity of this security group rule, which aids in uniquely identifying the security group rule.
	Identity *string `json:"Identity,omitempty"`
	// Position of security group rule in a security group.
	Index *int64 `json:"Index,omitempty"`
	// User-provided name to identify the security group rule.
	Name     *string `json:"Name,omitempty"`
	PortList []int64 `json:"PortList,omitempty"`
	// The IP protocol name that's open to network traffic, such as TCP, UDP, etc. * `tcp` - The TCP (Transmission Control Protocol) protocol. * `udp` - The UDP (User Datagram Protocol) protocol. * `icmp` - The ICMP (Internet Control Message Protocol) protocol. * `esp` - The ESP (Encapsulating Security Payload) protocol. * `ah` - The AH (Authentication Header) protocol. * `sctp` - The SCTP (Stream Control Transmission Protocol) protocol. * `all` - All of TCP, UDP, ICMP, ESP, AH and SCTP. * `none` - None of TCP, UDP, ICMP, ESP, AH and SCTP.
	Protocol   *string  `json:"Protocol,omitempty"`
	SourceCidr []string `json:"SourceCidr,omitempty"`
	// Reference to the existing security group, where the security group rule is defined.
	SourceSecurityGroup *string `json:"SourceSecurityGroup,omitempty"`
	// The start of port range for the security group rule IP protocol.
	StartPort            *int64 `json:"StartPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSecurityGroupRuleAllOf CloudSecurityGroupRuleAllOf

// NewCloudSecurityGroupRuleAllOf instantiates a new CloudSecurityGroupRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSecurityGroupRuleAllOf(classId string, objectType string) *CloudSecurityGroupRuleAllOf {
	this := CloudSecurityGroupRuleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var etherType string = "IPv4"
	this.EtherType = &etherType
	var protocol string = "tcp"
	this.Protocol = &protocol
	return &this
}

// NewCloudSecurityGroupRuleAllOfWithDefaults instantiates a new CloudSecurityGroupRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSecurityGroupRuleAllOfWithDefaults() *CloudSecurityGroupRuleAllOf {
	this := CloudSecurityGroupRuleAllOf{}
	var classId string = "cloud.SecurityGroupRule"
	this.ClassId = classId
	var objectType string = "cloud.SecurityGroupRule"
	this.ObjectType = objectType
	var etherType string = "IPv4"
	this.EtherType = &etherType
	var protocol string = "tcp"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSecurityGroupRuleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSecurityGroupRuleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSecurityGroupRuleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSecurityGroupRuleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CloudSecurityGroupRuleAllOf) SetAction(v string) {
	o.Action = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudSecurityGroupRuleAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndPort returns the EndPort field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetEndPort() int64 {
	if o == nil || o.EndPort == nil {
		var ret int64
		return ret
	}
	return *o.EndPort
}

// GetEndPortOk returns a tuple with the EndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetEndPortOk() (*int64, bool) {
	if o == nil || o.EndPort == nil {
		return nil, false
	}
	return o.EndPort, true
}

// HasEndPort returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasEndPort() bool {
	if o != nil && o.EndPort != nil {
		return true
	}

	return false
}

// SetEndPort gets a reference to the given int64 and assigns it to the EndPort field.
func (o *CloudSecurityGroupRuleAllOf) SetEndPort(v int64) {
	o.EndPort = &v
}

// GetEtherType returns the EtherType field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetEtherType() string {
	if o == nil || o.EtherType == nil {
		var ret string
		return ret
	}
	return *o.EtherType
}

// GetEtherTypeOk returns a tuple with the EtherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetEtherTypeOk() (*string, bool) {
	if o == nil || o.EtherType == nil {
		return nil, false
	}
	return o.EtherType, true
}

// HasEtherType returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasEtherType() bool {
	if o != nil && o.EtherType != nil {
		return true
	}

	return false
}

// SetEtherType gets a reference to the given string and assigns it to the EtherType field.
func (o *CloudSecurityGroupRuleAllOf) SetEtherType(v string) {
	o.EtherType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudSecurityGroupRuleAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *CloudSecurityGroupRuleAllOf) SetIndex(v int64) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudSecurityGroupRuleAllOf) SetName(v string) {
	o.Name = &v
}

// GetPortList returns the PortList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSecurityGroupRuleAllOf) GetPortList() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PortList
}

// GetPortListOk returns a tuple with the PortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSecurityGroupRuleAllOf) GetPortListOk() (*[]int64, bool) {
	if o == nil || o.PortList == nil {
		return nil, false
	}
	return &o.PortList, true
}

// HasPortList returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasPortList() bool {
	if o != nil && o.PortList != nil {
		return true
	}

	return false
}

// SetPortList gets a reference to the given []int64 and assigns it to the PortList field.
func (o *CloudSecurityGroupRuleAllOf) SetPortList(v []int64) {
	o.PortList = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *CloudSecurityGroupRuleAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSourceCidr returns the SourceCidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSecurityGroupRuleAllOf) GetSourceCidr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceCidr
}

// GetSourceCidrOk returns a tuple with the SourceCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSecurityGroupRuleAllOf) GetSourceCidrOk() (*[]string, bool) {
	if o == nil || o.SourceCidr == nil {
		return nil, false
	}
	return &o.SourceCidr, true
}

// HasSourceCidr returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasSourceCidr() bool {
	if o != nil && o.SourceCidr != nil {
		return true
	}

	return false
}

// SetSourceCidr gets a reference to the given []string and assigns it to the SourceCidr field.
func (o *CloudSecurityGroupRuleAllOf) SetSourceCidr(v []string) {
	o.SourceCidr = v
}

// GetSourceSecurityGroup returns the SourceSecurityGroup field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetSourceSecurityGroup() string {
	if o == nil || o.SourceSecurityGroup == nil {
		var ret string
		return ret
	}
	return *o.SourceSecurityGroup
}

// GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetSourceSecurityGroupOk() (*string, bool) {
	if o == nil || o.SourceSecurityGroup == nil {
		return nil, false
	}
	return o.SourceSecurityGroup, true
}

// HasSourceSecurityGroup returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasSourceSecurityGroup() bool {
	if o != nil && o.SourceSecurityGroup != nil {
		return true
	}

	return false
}

// SetSourceSecurityGroup gets a reference to the given string and assigns it to the SourceSecurityGroup field.
func (o *CloudSecurityGroupRuleAllOf) SetSourceSecurityGroup(v string) {
	o.SourceSecurityGroup = &v
}

// GetStartPort returns the StartPort field value if set, zero value otherwise.
func (o *CloudSecurityGroupRuleAllOf) GetStartPort() int64 {
	if o == nil || o.StartPort == nil {
		var ret int64
		return ret
	}
	return *o.StartPort
}

// GetStartPortOk returns a tuple with the StartPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSecurityGroupRuleAllOf) GetStartPortOk() (*int64, bool) {
	if o == nil || o.StartPort == nil {
		return nil, false
	}
	return o.StartPort, true
}

// HasStartPort returns a boolean if a field has been set.
func (o *CloudSecurityGroupRuleAllOf) HasStartPort() bool {
	if o != nil && o.StartPort != nil {
		return true
	}

	return false
}

// SetStartPort gets a reference to the given int64 and assigns it to the StartPort field.
func (o *CloudSecurityGroupRuleAllOf) SetStartPort(v int64) {
	o.StartPort = &v
}

func (o CloudSecurityGroupRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndPort != nil {
		toSerialize["EndPort"] = o.EndPort
	}
	if o.EtherType != nil {
		toSerialize["EtherType"] = o.EtherType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PortList != nil {
		toSerialize["PortList"] = o.PortList
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.SourceCidr != nil {
		toSerialize["SourceCidr"] = o.SourceCidr
	}
	if o.SourceSecurityGroup != nil {
		toSerialize["SourceSecurityGroup"] = o.SourceSecurityGroup
	}
	if o.StartPort != nil {
		toSerialize["StartPort"] = o.StartPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSecurityGroupRuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudSecurityGroupRuleAllOf := _CloudSecurityGroupRuleAllOf{}

	if err = json.Unmarshal(bytes, &varCloudSecurityGroupRuleAllOf); err == nil {
		*o = CloudSecurityGroupRuleAllOf(varCloudSecurityGroupRuleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndPort")
		delete(additionalProperties, "EtherType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Index")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PortList")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "SourceCidr")
		delete(additionalProperties, "SourceSecurityGroup")
		delete(additionalProperties, "StartPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSecurityGroupRuleAllOf struct {
	value *CloudSecurityGroupRuleAllOf
	isSet bool
}

func (v NullableCloudSecurityGroupRuleAllOf) Get() *CloudSecurityGroupRuleAllOf {
	return v.value
}

func (v *NullableCloudSecurityGroupRuleAllOf) Set(val *CloudSecurityGroupRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSecurityGroupRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSecurityGroupRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSecurityGroupRuleAllOf(val *CloudSecurityGroupRuleAllOf) *NullableCloudSecurityGroupRuleAllOf {
	return &NullableCloudSecurityGroupRuleAllOf{value: val, isSet: true}
}

func (v NullableCloudSecurityGroupRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSecurityGroupRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
