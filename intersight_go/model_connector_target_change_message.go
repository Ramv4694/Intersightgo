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
	"reflect"
	"strings"
)

// ConnectorTargetChangeMessage The message sent to the Appliance device connector when a Target is created, modified or deleted. Appliance device connector is expected to maintain a durable cache of relevent Target information such that it can perform inventory collection and change operations against the target without soliciting basic connectivity details from Intersight.
type ConnectorTargetChangeMessage struct {
	ConnectorBaseMessage
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

type _ConnectorTargetChangeMessage ConnectorTargetChangeMessage

// NewConnectorTargetChangeMessage instantiates a new ConnectorTargetChangeMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTargetChangeMessage(classId string, objectType string) *ConnectorTargetChangeMessage {
	this := ConnectorTargetChangeMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// NewConnectorTargetChangeMessageWithDefaults instantiates a new ConnectorTargetChangeMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTargetChangeMessageWithDefaults() *ConnectorTargetChangeMessage {
	this := ConnectorTargetChangeMessage{}
	var classId string = "connector.TargetChangeMessage"
	this.ClassId = classId
	var objectType string = "connector.TargetChangeMessage"
	this.ObjectType = objectType
	var modStatus string = "None"
	this.ModStatus = &modStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorTargetChangeMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorTargetChangeMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorTargetChangeMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorTargetChangeMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModStatus returns the ModStatus field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessage) GetModStatus() string {
	if o == nil || o.ModStatus == nil {
		var ret string
		return ret
	}
	return *o.ModStatus
}

// GetModStatusOk returns a tuple with the ModStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessage) GetModStatusOk() (*string, bool) {
	if o == nil || o.ModStatus == nil {
		return nil, false
	}
	return o.ModStatus, true
}

// HasModStatus returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessage) HasModStatus() bool {
	if o != nil && o.ModStatus != nil {
		return true
	}

	return false
}

// SetModStatus gets a reference to the given string and assigns it to the ModStatus field.
func (o *ConnectorTargetChangeMessage) SetModStatus(v string) {
	o.ModStatus = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessage) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessage) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessage) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *ConnectorTargetChangeMessage) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetTargetDetails returns the TargetDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorTargetChangeMessage) GetTargetDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorTargetChangeMessage) GetTargetDetailsOk() (*interface{}, bool) {
	if o == nil || o.TargetDetails == nil {
		return nil, false
	}
	return &o.TargetDetails, true
}

// HasTargetDetails returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessage) HasTargetDetails() bool {
	if o != nil && o.TargetDetails != nil {
		return true
	}

	return false
}

// SetTargetDetails gets a reference to the given interface{} and assigns it to the TargetDetails field.
func (o *ConnectorTargetChangeMessage) SetTargetDetails(v interface{}) {
	o.TargetDetails = v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *ConnectorTargetChangeMessage) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTargetChangeMessage) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessage) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *ConnectorTargetChangeMessage) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetSpecification returns the TargetSpecification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorTargetChangeMessage) GetTargetSpecification() ConnectorTargetSpecification {
	if o == nil || o.TargetSpecification.Get() == nil {
		var ret ConnectorTargetSpecification
		return ret
	}
	return *o.TargetSpecification.Get()
}

// GetTargetSpecificationOk returns a tuple with the TargetSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorTargetChangeMessage) GetTargetSpecificationOk() (*ConnectorTargetSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSpecification.Get(), o.TargetSpecification.IsSet()
}

// HasTargetSpecification returns a boolean if a field has been set.
func (o *ConnectorTargetChangeMessage) HasTargetSpecification() bool {
	if o != nil && o.TargetSpecification.IsSet() {
		return true
	}

	return false
}

// SetTargetSpecification gets a reference to the given NullableConnectorTargetSpecification and assigns it to the TargetSpecification field.
func (o *ConnectorTargetChangeMessage) SetTargetSpecification(v ConnectorTargetSpecification) {
	o.TargetSpecification.Set(&v)
}

// SetTargetSpecificationNil sets the value for TargetSpecification to be an explicit nil
func (o *ConnectorTargetChangeMessage) SetTargetSpecificationNil() {
	o.TargetSpecification.Set(nil)
}

// UnsetTargetSpecification ensures that no value is present for TargetSpecification, not even an explicit nil
func (o *ConnectorTargetChangeMessage) UnsetTargetSpecification() {
	o.TargetSpecification.Unset()
}

func (o ConnectorTargetChangeMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
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

func (o *ConnectorTargetChangeMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorTargetChangeMessageWithoutEmbeddedStruct struct {
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
		TargetMoid          *string                              `json:"TargetMoid,omitempty"`
		TargetSpecification NullableConnectorTargetSpecification `json:"TargetSpecification,omitempty"`
	}

	varConnectorTargetChangeMessageWithoutEmbeddedStruct := ConnectorTargetChangeMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorTargetChangeMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorTargetChangeMessage := _ConnectorTargetChangeMessage{}
		varConnectorTargetChangeMessage.ClassId = varConnectorTargetChangeMessageWithoutEmbeddedStruct.ClassId
		varConnectorTargetChangeMessage.ObjectType = varConnectorTargetChangeMessageWithoutEmbeddedStruct.ObjectType
		varConnectorTargetChangeMessage.ModStatus = varConnectorTargetChangeMessageWithoutEmbeddedStruct.ModStatus
		varConnectorTargetChangeMessage.ServiceType = varConnectorTargetChangeMessageWithoutEmbeddedStruct.ServiceType
		varConnectorTargetChangeMessage.TargetDetails = varConnectorTargetChangeMessageWithoutEmbeddedStruct.TargetDetails
		varConnectorTargetChangeMessage.TargetMoid = varConnectorTargetChangeMessageWithoutEmbeddedStruct.TargetMoid
		varConnectorTargetChangeMessage.TargetSpecification = varConnectorTargetChangeMessageWithoutEmbeddedStruct.TargetSpecification
		*o = ConnectorTargetChangeMessage(varConnectorTargetChangeMessage)
	} else {
		return err
	}

	varConnectorTargetChangeMessage := _ConnectorTargetChangeMessage{}

	err = json.Unmarshal(bytes, &varConnectorTargetChangeMessage)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorTargetChangeMessage.ConnectorBaseMessage
	} else {
		return err
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

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorTargetChangeMessage struct {
	value *ConnectorTargetChangeMessage
	isSet bool
}

func (v NullableConnectorTargetChangeMessage) Get() *ConnectorTargetChangeMessage {
	return v.value
}

func (v *NullableConnectorTargetChangeMessage) Set(val *ConnectorTargetChangeMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTargetChangeMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTargetChangeMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTargetChangeMessage(val *ConnectorTargetChangeMessage) *NullableConnectorTargetChangeMessage {
	return &NullableConnectorTargetChangeMessage{value: val, isSet: true}
}

func (v NullableConnectorTargetChangeMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTargetChangeMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
