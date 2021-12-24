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
	"reflect"
	"strings"
)

// NiatelemetryMsoContractDetails Details of contract configured from the Multi-Site Orchestrator.
type NiatelemetryMsoContractDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of contract in Multi-Site Orchestrator.
	ContractName *string `json:"ContractName,omitempty"`
	// Site Ids to which this contract is deployed to.
	DeployedSites *string `json:"DeployedSites,omitempty"`
	// Is the contract local to the Multi-Site Orchestrator.
	IsLocal *string `json:"IsLocal,omitempty"`
	// Unique reference for the contract in the Multi-Site Orchestrator.
	Reference *string `json:"Reference,omitempty"`
	// Schema ID in Multi-Site Orchestrator.
	SchemaId *string `json:"SchemaId,omitempty"`
	// Schema name this contract belongs to in Multi-Site Orchestrator.
	SchemaName *string `json:"SchemaName,omitempty"`
	// Template name this contract belongs to in Multi-Site Orchestrator.
	TemplateName         *string                              `json:"TemplateName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoContractDetails NiatelemetryMsoContractDetails

// NewNiatelemetryMsoContractDetails instantiates a new NiatelemetryMsoContractDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoContractDetails(classId string, objectType string) *NiatelemetryMsoContractDetails {
	this := NiatelemetryMsoContractDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoContractDetailsWithDefaults instantiates a new NiatelemetryMsoContractDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoContractDetailsWithDefaults() *NiatelemetryMsoContractDetails {
	this := NiatelemetryMsoContractDetails{}
	var classId string = "niatelemetry.MsoContractDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoContractDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoContractDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoContractDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoContractDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoContractDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryMsoContractDetails) SetContractName(v string) {
	o.ContractName = &v
}

// GetDeployedSites returns the DeployedSites field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetDeployedSites() string {
	if o == nil || o.DeployedSites == nil {
		var ret string
		return ret
	}
	return *o.DeployedSites
}

// GetDeployedSitesOk returns a tuple with the DeployedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetDeployedSitesOk() (*string, bool) {
	if o == nil || o.DeployedSites == nil {
		return nil, false
	}
	return o.DeployedSites, true
}

// HasDeployedSites returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasDeployedSites() bool {
	if o != nil && o.DeployedSites != nil {
		return true
	}

	return false
}

// SetDeployedSites gets a reference to the given string and assigns it to the DeployedSites field.
func (o *NiatelemetryMsoContractDetails) SetDeployedSites(v string) {
	o.DeployedSites = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetIsLocal() string {
	if o == nil || o.IsLocal == nil {
		var ret string
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetIsLocalOk() (*string, bool) {
	if o == nil || o.IsLocal == nil {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasIsLocal() bool {
	if o != nil && o.IsLocal != nil {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given string and assigns it to the IsLocal field.
func (o *NiatelemetryMsoContractDetails) SetIsLocal(v string) {
	o.IsLocal = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *NiatelemetryMsoContractDetails) SetReference(v string) {
	o.Reference = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *NiatelemetryMsoContractDetails) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasSchemaName() bool {
	if o != nil && o.SchemaName != nil {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *NiatelemetryMsoContractDetails) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *NiatelemetryMsoContractDetails) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoContractDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoContractDetails) MarshalJSON() ([]byte, error) {
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
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.DeployedSites != nil {
		toSerialize["DeployedSites"] = o.DeployedSites
	}
	if o.IsLocal != nil {
		toSerialize["IsLocal"] = o.IsLocal
	}
	if o.Reference != nil {
		toSerialize["Reference"] = o.Reference
	}
	if o.SchemaId != nil {
		toSerialize["SchemaId"] = o.SchemaId
	}
	if o.SchemaName != nil {
		toSerialize["SchemaName"] = o.SchemaName
	}
	if o.TemplateName != nil {
		toSerialize["TemplateName"] = o.TemplateName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMsoContractDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryMsoContractDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of contract in Multi-Site Orchestrator.
		ContractName *string `json:"ContractName,omitempty"`
		// Site Ids to which this contract is deployed to.
		DeployedSites *string `json:"DeployedSites,omitempty"`
		// Is the contract local to the Multi-Site Orchestrator.
		IsLocal *string `json:"IsLocal,omitempty"`
		// Unique reference for the contract in the Multi-Site Orchestrator.
		Reference *string `json:"Reference,omitempty"`
		// Schema ID in Multi-Site Orchestrator.
		SchemaId *string `json:"SchemaId,omitempty"`
		// Schema name this contract belongs to in Multi-Site Orchestrator.
		SchemaName *string `json:"SchemaName,omitempty"`
		// Template name this contract belongs to in Multi-Site Orchestrator.
		TemplateName     *string                              `json:"TemplateName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct := NiatelemetryMsoContractDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMsoContractDetails := _NiatelemetryMsoContractDetails{}
		varNiatelemetryMsoContractDetails.ClassId = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMsoContractDetails.ObjectType = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMsoContractDetails.ContractName = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.ContractName
		varNiatelemetryMsoContractDetails.DeployedSites = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.DeployedSites
		varNiatelemetryMsoContractDetails.IsLocal = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.IsLocal
		varNiatelemetryMsoContractDetails.Reference = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.Reference
		varNiatelemetryMsoContractDetails.SchemaId = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.SchemaId
		varNiatelemetryMsoContractDetails.SchemaName = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.SchemaName
		varNiatelemetryMsoContractDetails.TemplateName = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.TemplateName
		varNiatelemetryMsoContractDetails.RegisteredDevice = varNiatelemetryMsoContractDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMsoContractDetails(varNiatelemetryMsoContractDetails)
	} else {
		return err
	}

	varNiatelemetryMsoContractDetails := _NiatelemetryMsoContractDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoContractDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMsoContractDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "DeployedSites")
		delete(additionalProperties, "IsLocal")
		delete(additionalProperties, "Reference")
		delete(additionalProperties, "SchemaId")
		delete(additionalProperties, "SchemaName")
		delete(additionalProperties, "TemplateName")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryMsoContractDetails struct {
	value *NiatelemetryMsoContractDetails
	isSet bool
}

func (v NullableNiatelemetryMsoContractDetails) Get() *NiatelemetryMsoContractDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoContractDetails) Set(val *NiatelemetryMsoContractDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoContractDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoContractDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoContractDetails(val *NiatelemetryMsoContractDetails) *NullableNiatelemetryMsoContractDetails {
	return &NullableNiatelemetryMsoContractDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoContractDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoContractDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
