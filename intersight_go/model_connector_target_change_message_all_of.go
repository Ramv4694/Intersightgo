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

// ConnectorTargetChangeMessageAllOf Definition of the list of properties defined in 'connector.TargetChangeMessage', excluding properties defined in parent classes.
type ConnectorTargetChangeMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ModStatus indicates if the change message conveys a creation, modification or deletion of an Target instance. * `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully. * `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.
	ModStatus *string `json:"ModStatus,omitempty"`
	// The name of the service to be deployed for the given target. If more than one service need to be deployed for a given target, multiple target change request is sent to Intersight Assist, each consists of one service type. It is different from the target type e.g., asset.OrchestrationService, asset.TerraformIntegrationService and asset.WorkloadOptimizerService are currently supported in Intersight Assist.
	ServiceType *string `json:"ServiceType,omitempty"`
	// A Json-serialized representation of the 'configuration' portion of the Target instance. Ie the representation contains configuration properties like the target's connectivity information but not operation status. The representation include credential information, encrypted with the RSA public key of the Appliance device connector. Appliance device connector is the sole maintainer of the RSA private key and the only system component which is capable of interpreting the credential.
	TargetDetails interface{} `json:"TargetDetails,omitempty"`
	// The Moid identifying the Target instance being created, modified or deleted.
	TargetMoid           *string                              `json:"TargetMoid,omitempty"`
	TargetSpecification  NullableConnectorTargetSpecification `json:"TargetSpecification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorTargetChangeMessageAllOf ConnectorTargetChangeMessageAllOf

// NewConnectorTargetChangeMessageAllOf instantiates a new ConnectorTargetChangeMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTargetChangeMessageAllOf(classId string, objectType string) *ConnectorTargetChangeMessageAllOf {
	this := ConnectorTargetChangeMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// NewConnectorTargetChangeMessageAllOfWithDefaults instantiates a new ConnectorTargetChangeMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTargetChangeMessageAllOfWithDefaults() *ConnectorTargetChangeMessageAllOf {
	this := ConnectorTargetChangeMessageAllOf{}
	var classId string = "connector.TargetChangeMessage"
	this.ClassId = classId
	var objectType string = "connector.TargetChangeMessage"
	this.ObjectType = objectType
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorTargetChangeMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorTargetChangeMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorTargetChangeMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorTargetChangeMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModStatus returns the ModStatus field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessageAllOf) GetModStatus() string {
	if o == nil || o.ModStatus == nil {
		var ret string
		return ret
	}
	return *o.ModStatus
}

// GetModStatusOk returns a tuple with the ModStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessageAllOf) GetModStatusOk() (*string, bool) {
	if o == nil || o.ModStatus == nil {
		return nil, false
	}
	return o.ModStatus, true
}

// HasModStatus returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessageAllOf) HasModStatus() bool {
	if o != nil && o.ModStatus != nil {
		return true
	}

	return false
}

// SetModStatus gets a reference to the given string and assigns it to the ModStatus field.
func (o *ConnectorTargetChangeMessageAllOf) SetModStatus(v string) {
	o.ModStatus = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessageAllOf) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessageAllOf) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessageAllOf) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *ConnectorTargetChangeMessageAllOf) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetTargetDetails returns the TargetDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorTargetChangeMessageAllOf) GetTargetDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorTargetChangeMessageAllOf) GetTargetDetailsOk() (*interface{}, bool) {
	if o == nil || o.TargetDetails == nil {
		return nil, false
	}
	return &o.TargetDetails, true
}

// HasTargetDetails returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessageAllOf) HasTargetDetails() bool {
	if o != nil && o.TargetDetails != nil {
		return true
	}

	return false
}

// SetTargetDetails gets a reference to the given interface{} and assigns it to the TargetDetails field.
func (o *ConnectorTargetChangeMessageAllOf) SetTargetDetails(v interface{}) {
	o.TargetDetails = v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessageAllOf) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessageAllOf) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessageAllOf) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *ConnectorTargetChangeMessageAllOf) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetSpecification returns the TargetSpecification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorTargetChangeMessageAllOf) GetTargetSpecification() ConnectorTargetSpecification {
	if o == nil || o.TargetSpecification.Get() == nil {
		var ret ConnectorTargetSpecification
		return ret
	}
	return *o.TargetSpecification.Get()
}

// GetTargetSpecificationOk returns a tuple with the TargetSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorTargetChangeMessageAllOf) GetTargetSpecificationOk() (*ConnectorTargetSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSpecification.Get(), o.TargetSpecification.IsSet()
}

// HasTargetSpecification returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessageAllOf) HasTargetSpecification() bool {
	if o != nil && o.TargetSpecification.IsSet() {
		return true
	}

	return false
}

// SetTargetSpecification gets a reference to the given NullableConnectorTargetSpecification and assigns it to the TargetSpecification field.
func (o *ConnectorTargetChangeMessageAllOf) SetTargetSpecification(v ConnectorTargetSpecification) {
	o.TargetSpecification.Set(&v)
}

// SetTargetSpecificationNil sets the value for TargetSpecification to be an explicit nil
func (o *ConnectorTargetChangeMessageAllOf) SetTargetSpecificationNil() {
	o.TargetSpecification.Set(nil)
}

// UnsetTargetSpecification ensures that no value is present for TargetSpecification, not even an explicit nil
func (o *ConnectorTargetChangeMessageAllOf) UnsetTargetSpecification() {
	o.TargetSpecification.Unset()
}

func (o ConnectorTargetChangeMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ModStatus != nil {
		toSerialize["ModStatus"] = o.ModStatus
	}
	if o.ServiceType != nil {
		toSerialize["ServiceType"] = o.ServiceType
	}
	if o.TargetDetails != nil {
		toSerialize["TargetDetails"] = o.TargetDetails
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetSpecification.IsSet() {
		toSerialize["TargetSpecification"] = o.TargetSpecification.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorTargetChangeMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorTargetChangeMessageAllOf := _ConnectorTargetChangeMessageAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorTargetChangeMessageAllOf); err == nil {
		*o = ConnectorTargetChangeMessageAllOf(varConnectorTargetChangeMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ModStatus")
		delete(additionalProperties, "ServiceType")
		delete(additionalProperties, "TargetDetails")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "TargetSpecification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorTargetChangeMessageAllOf struct {
	value *ConnectorTargetChangeMessageAllOf
	isSet bool
}

func (v NullableConnectorTargetChangeMessageAllOf) Get() *ConnectorTargetChangeMessageAllOf {
	return v.value
}

func (v *NullableConnectorTargetChangeMessageAllOf) Set(val *ConnectorTargetChangeMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTargetChangeMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTargetChangeMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTargetChangeMessageAllOf(val *ConnectorTargetChangeMessageAllOf) *NullableConnectorTargetChangeMessageAllOf {
	return &NullableConnectorTargetChangeMessageAllOf{value: val, isSet: true}
}

func (v NullableConnectorTargetChangeMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTargetChangeMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
