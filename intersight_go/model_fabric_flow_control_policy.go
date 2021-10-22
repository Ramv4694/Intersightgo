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

// FabricFlowControlPolicy Priority Flow Control setting for each port.
type FabricFlowControlPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Configure PFC on a per-port basis to enable the no-drop behavior for the CoS as defined by the active network qos policy. * `auto` - Enables the no-drop CoS values to be advertised by the DCBXP and negotiated with the peer.A successful negotiation enables PFC on the no-drop CoS.Any failures because of a mismatch in the capability of peers causes the PFC not to be enabled. * `on` - Enables PFC on the local port regardless of the capability of the peers.
	PriorityFlowControlMode *string `json:"PriorityFlowControlMode,omitempty"`
	// Link-level Flow Control configured in the receive direction. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	ReceiveDirection *string `json:"ReceiveDirection,omitempty"`
	// Link-level Flow Control configured in the send direction. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	SendDirection        *string                               `json:"SendDirection,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFlowControlPolicy FabricFlowControlPolicy

// NewFabricFlowControlPolicy instantiates a new FabricFlowControlPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFlowControlPolicy(classId string, objectType string) *FabricFlowControlPolicy {
	this := FabricFlowControlPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var priorityFlowControlMode string = "auto"
	this.PriorityFlowControlMode = &priorityFlowControlMode
	var receiveDirection string = "Disabled"
	this.ReceiveDirection = &receiveDirection
	var sendDirection string = "Disabled"
	this.SendDirection = &sendDirection
	return &this
}

// NewFabricFlowControlPolicyWithDefaults instantiates a new FabricFlowControlPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFlowControlPolicyWithDefaults() *FabricFlowControlPolicy {
	this := FabricFlowControlPolicy{}
	var classId string = "fabric.FlowControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.FlowControlPolicy"
	this.ObjectType = objectType
	var priorityFlowControlMode string = "auto"
	this.PriorityFlowControlMode = &priorityFlowControlMode
	var receiveDirection string = "Disabled"
	this.ReceiveDirection = &receiveDirection
	var sendDirection string = "Disabled"
	this.SendDirection = &sendDirection
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFlowControlPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFlowControlPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFlowControlPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFlowControlPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPriorityFlowControlMode returns the PriorityFlowControlMode field value if set, zero value otherwise.
func (o *FabricFlowControlPolicy) GetPriorityFlowControlMode() string {
	if o == nil || o.PriorityFlowControlMode == nil {
		var ret string
		return ret
	}
	return *o.PriorityFlowControlMode
}

// GetPriorityFlowControlModeOk returns a tuple with the PriorityFlowControlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetPriorityFlowControlModeOk() (*string, bool) {
	if o == nil || o.PriorityFlowControlMode == nil {
		return nil, false
	}
	return o.PriorityFlowControlMode, true
}

// HasPriorityFlowControlMode returns a boolean if a field has been set.
func (o *FabricFlowControlPolicy) HasPriorityFlowControlMode() bool {
	if o != nil && o.PriorityFlowControlMode != nil {
		return true
	}

	return false
}

// SetPriorityFlowControlMode gets a reference to the given string and assigns it to the PriorityFlowControlMode field.
func (o *FabricFlowControlPolicy) SetPriorityFlowControlMode(v string) {
	o.PriorityFlowControlMode = &v
}

// GetReceiveDirection returns the ReceiveDirection field value if set, zero value otherwise.
func (o *FabricFlowControlPolicy) GetReceiveDirection() string {
	if o == nil || o.ReceiveDirection == nil {
		var ret string
		return ret
	}
	return *o.ReceiveDirection
}

// GetReceiveDirectionOk returns a tuple with the ReceiveDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetReceiveDirectionOk() (*string, bool) {
	if o == nil || o.ReceiveDirection == nil {
		return nil, false
	}
	return o.ReceiveDirection, true
}

// HasReceiveDirection returns a boolean if a field has been set.
func (o *FabricFlowControlPolicy) HasReceiveDirection() bool {
	if o != nil && o.ReceiveDirection != nil {
		return true
	}

	return false
}

// SetReceiveDirection gets a reference to the given string and assigns it to the ReceiveDirection field.
func (o *FabricFlowControlPolicy) SetReceiveDirection(v string) {
	o.ReceiveDirection = &v
}

// GetSendDirection returns the SendDirection field value if set, zero value otherwise.
func (o *FabricFlowControlPolicy) GetSendDirection() string {
	if o == nil || o.SendDirection == nil {
		var ret string
		return ret
	}
	return *o.SendDirection
}

// GetSendDirectionOk returns a tuple with the SendDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetSendDirectionOk() (*string, bool) {
	if o == nil || o.SendDirection == nil {
		return nil, false
	}
	return o.SendDirection, true
}

// HasSendDirection returns a boolean if a field has been set.
func (o *FabricFlowControlPolicy) HasSendDirection() bool {
	if o != nil && o.SendDirection != nil {
		return true
	}

	return false
}

// SetSendDirection gets a reference to the given string and assigns it to the SendDirection field.
func (o *FabricFlowControlPolicy) SetSendDirection(v string) {
	o.SendDirection = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricFlowControlPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFlowControlPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricFlowControlPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricFlowControlPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricFlowControlPolicy) MarshalJSON() ([]byte, error) {
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
	if o.PriorityFlowControlMode != nil {
		toSerialize["PriorityFlowControlMode"] = o.PriorityFlowControlMode
	}
	if o.ReceiveDirection != nil {
		toSerialize["ReceiveDirection"] = o.ReceiveDirection
	}
	if o.SendDirection != nil {
		toSerialize["SendDirection"] = o.SendDirection
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFlowControlPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFlowControlPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Configure PFC on a per-port basis to enable the no-drop behavior for the CoS as defined by the active network qos policy. * `auto` - Enables the no-drop CoS values to be advertised by the DCBXP and negotiated with the peer.A successful negotiation enables PFC on the no-drop CoS.Any failures because of a mismatch in the capability of peers causes the PFC not to be enabled. * `on` - Enables PFC on the local port regardless of the capability of the peers.
		PriorityFlowControlMode *string `json:"PriorityFlowControlMode,omitempty"`
		// Link-level Flow Control configured in the receive direction. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		ReceiveDirection *string `json:"ReceiveDirection,omitempty"`
		// Link-level Flow Control configured in the send direction. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
		SendDirection *string                               `json:"SendDirection,omitempty"`
		Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varFabricFlowControlPolicyWithoutEmbeddedStruct := FabricFlowControlPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFlowControlPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricFlowControlPolicy := _FabricFlowControlPolicy{}
		varFabricFlowControlPolicy.ClassId = varFabricFlowControlPolicyWithoutEmbeddedStruct.ClassId
		varFabricFlowControlPolicy.ObjectType = varFabricFlowControlPolicyWithoutEmbeddedStruct.ObjectType
		varFabricFlowControlPolicy.PriorityFlowControlMode = varFabricFlowControlPolicyWithoutEmbeddedStruct.PriorityFlowControlMode
		varFabricFlowControlPolicy.ReceiveDirection = varFabricFlowControlPolicyWithoutEmbeddedStruct.ReceiveDirection
		varFabricFlowControlPolicy.SendDirection = varFabricFlowControlPolicyWithoutEmbeddedStruct.SendDirection
		varFabricFlowControlPolicy.Organization = varFabricFlowControlPolicyWithoutEmbeddedStruct.Organization
		*o = FabricFlowControlPolicy(varFabricFlowControlPolicy)
	} else {
		return err
	}

	varFabricFlowControlPolicy := _FabricFlowControlPolicy{}

	err = json.Unmarshal(bytes, &varFabricFlowControlPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricFlowControlPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PriorityFlowControlMode")
		delete(additionalProperties, "ReceiveDirection")
		delete(additionalProperties, "SendDirection")
		delete(additionalProperties, "Organization")

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

type NullableFabricFlowControlPolicy struct {
	value *FabricFlowControlPolicy
	isSet bool
}

func (v NullableFabricFlowControlPolicy) Get() *FabricFlowControlPolicy {
	return v.value
}

func (v *NullableFabricFlowControlPolicy) Set(val *FabricFlowControlPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFlowControlPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFlowControlPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFlowControlPolicy(val *FabricFlowControlPolicy) *NullableFabricFlowControlPolicy {
	return &NullableFabricFlowControlPolicy{value: val, isSet: true}
}

func (v NullableFabricFlowControlPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFlowControlPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
