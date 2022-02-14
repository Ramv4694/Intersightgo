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
	"time"
)

// ApplianceDeviceCertificate DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
type ApplianceDeviceCertificate struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The base64 encoded certificate in PEM format.
	CaCertificate *string `json:"CaCertificate,omitempty"`
	// The expiry datetime of new ca certificate which need to be applied on device connector.
	CaCertificateExpiryTime *time.Time `json:"CaCertificateExpiryTime,omitempty"`
	// The date time allocated till cert renewal will be executed. This time used here will be based on cert renewal plan.
	CertificateRenewalExpiryTime *time.Time                  `json:"CertificateRenewalExpiryTime,omitempty"`
	CompletedPhases              []ApplianceCertRenewalPhase `json:"CompletedPhases,omitempty"`
	// The operation configuration MOId.
	ConfigurationMoId *string                           `json:"ConfigurationMoId,omitempty"`
	CurrentPhase      NullableApplianceCertRenewalPhase `json:"CurrentPhase,omitempty"`
	// End date of the certificate renewal.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle.
	LastSuccessPollTime *time.Time `json:"LastSuccessPollTime,omitempty"`
	Messages            []string   `json:"Messages,omitempty"`
	// Start date of the certificate renewal.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The status of ca certificate renewal.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDeviceCertificate ApplianceDeviceCertificate

// NewApplianceDeviceCertificate instantiates a new ApplianceDeviceCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDeviceCertificate(classId string, objectType string) *ApplianceDeviceCertificate {
	this := ApplianceDeviceCertificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceDeviceCertificateWithDefaults instantiates a new ApplianceDeviceCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDeviceCertificateWithDefaults() *ApplianceDeviceCertificate {
	this := ApplianceDeviceCertificate{}
	var classId string = "appliance.DeviceCertificate"
	this.ClassId = classId
	var objectType string = "appliance.DeviceCertificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDeviceCertificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDeviceCertificate) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDeviceCertificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDeviceCertificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCaCertificate() string {
	if o == nil || o.CaCertificate == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCaCertificateOk() (*string, bool) {
	if o == nil || o.CaCertificate == nil {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCaCertificate() bool {
	if o != nil && o.CaCertificate != nil {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *ApplianceDeviceCertificate) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCaCertificateExpiryTime returns the CaCertificateExpiryTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTime() time.Time {
	if o == nil || o.CaCertificateExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CaCertificateExpiryTime
}

// GetCaCertificateExpiryTimeOk returns a tuple with the CaCertificateExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.CaCertificateExpiryTime == nil {
		return nil, false
	}
	return o.CaCertificateExpiryTime, true
}

// HasCaCertificateExpiryTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCaCertificateExpiryTime() bool {
	if o != nil && o.CaCertificateExpiryTime != nil {
		return true
	}

	return false
}

// SetCaCertificateExpiryTime gets a reference to the given time.Time and assigns it to the CaCertificateExpiryTime field.
func (o *ApplianceDeviceCertificate) SetCaCertificateExpiryTime(v time.Time) {
	o.CaCertificateExpiryTime = &v
}

// GetCertificateRenewalExpiryTime returns the CertificateRenewalExpiryTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTime() time.Time {
	if o == nil || o.CertificateRenewalExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CertificateRenewalExpiryTime
}

// GetCertificateRenewalExpiryTimeOk returns a tuple with the CertificateRenewalExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.CertificateRenewalExpiryTime == nil {
		return nil, false
	}
	return o.CertificateRenewalExpiryTime, true
}

// HasCertificateRenewalExpiryTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCertificateRenewalExpiryTime() bool {
	if o != nil && o.CertificateRenewalExpiryTime != nil {
		return true
	}

	return false
}

// SetCertificateRenewalExpiryTime gets a reference to the given time.Time and assigns it to the CertificateRenewalExpiryTime field.
func (o *ApplianceDeviceCertificate) SetCertificateRenewalExpiryTime(v time.Time) {
	o.CertificateRenewalExpiryTime = &v
}

// GetCompletedPhases returns the CompletedPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetCompletedPhases() []ApplianceCertRenewalPhase {
	if o == nil {
		var ret []ApplianceCertRenewalPhase
		return ret
	}
	return o.CompletedPhases
}

// GetCompletedPhasesOk returns a tuple with the CompletedPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetCompletedPhasesOk() (*[]ApplianceCertRenewalPhase, bool) {
	if o == nil || o.CompletedPhases == nil {
		return nil, false
	}
	return &o.CompletedPhases, true
}

// HasCompletedPhases returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCompletedPhases() bool {
	if o != nil && o.CompletedPhases != nil {
		return true
	}

	return false
}

// SetCompletedPhases gets a reference to the given []ApplianceCertRenewalPhase and assigns it to the CompletedPhases field.
func (o *ApplianceDeviceCertificate) SetCompletedPhases(v []ApplianceCertRenewalPhase) {
	o.CompletedPhases = v
}

// GetConfigurationMoId returns the ConfigurationMoId field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetConfigurationMoId() string {
	if o == nil || o.ConfigurationMoId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationMoId
}

// GetConfigurationMoIdOk returns a tuple with the ConfigurationMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetConfigurationMoIdOk() (*string, bool) {
	if o == nil || o.ConfigurationMoId == nil {
		return nil, false
	}
	return o.ConfigurationMoId, true
}

// HasConfigurationMoId returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasConfigurationMoId() bool {
	if o != nil && o.ConfigurationMoId != nil {
		return true
	}

	return false
}

// SetConfigurationMoId gets a reference to the given string and assigns it to the ConfigurationMoId field.
func (o *ApplianceDeviceCertificate) SetConfigurationMoId(v string) {
	o.ConfigurationMoId = &v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetCurrentPhase() ApplianceCertRenewalPhase {
	if o == nil || o.CurrentPhase.Get() == nil {
		var ret ApplianceCertRenewalPhase
		return ret
	}
	return *o.CurrentPhase.Get()
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetCurrentPhaseOk() (*ApplianceCertRenewalPhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPhase.Get(), o.CurrentPhase.IsSet()
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasCurrentPhase() bool {
	if o != nil && o.CurrentPhase.IsSet() {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given NullableApplianceCertRenewalPhase and assigns it to the CurrentPhase field.
func (o *ApplianceDeviceCertificate) SetCurrentPhase(v ApplianceCertRenewalPhase) {
	o.CurrentPhase.Set(&v)
}

// SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil
func (o *ApplianceDeviceCertificate) SetCurrentPhaseNil() {
	o.CurrentPhase.Set(nil)
}

// UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
func (o *ApplianceDeviceCertificate) UnsetCurrentPhase() {
	o.CurrentPhase.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceDeviceCertificate) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetLastSuccessPollTime returns the LastSuccessPollTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetLastSuccessPollTime() time.Time {
	if o == nil || o.LastSuccessPollTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessPollTime
}

// GetLastSuccessPollTimeOk returns a tuple with the LastSuccessPollTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetLastSuccessPollTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSuccessPollTime == nil {
		return nil, false
	}
	return o.LastSuccessPollTime, true
}

// HasLastSuccessPollTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasLastSuccessPollTime() bool {
	if o != nil && o.LastSuccessPollTime != nil {
		return true
	}

	return false
}

// SetLastSuccessPollTime gets a reference to the given time.Time and assigns it to the LastSuccessPollTime field.
func (o *ApplianceDeviceCertificate) SetLastSuccessPollTime(v time.Time) {
	o.LastSuccessPollTime = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceCertificate) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceCertificate) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceDeviceCertificate) SetMessages(v []string) {
	o.Messages = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceDeviceCertificate) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceDeviceCertificate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceCertificate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceDeviceCertificate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceDeviceCertificate) SetStatus(v string) {
	o.Status = &v
}

func (o ApplianceDeviceCertificate) MarshalJSON() ([]byte, error) {
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
	if o.CaCertificate != nil {
		toSerialize["CaCertificate"] = o.CaCertificate
	}
	if o.CaCertificateExpiryTime != nil {
		toSerialize["CaCertificateExpiryTime"] = o.CaCertificateExpiryTime
	}
	if o.CertificateRenewalExpiryTime != nil {
		toSerialize["CertificateRenewalExpiryTime"] = o.CertificateRenewalExpiryTime
	}
	if o.CompletedPhases != nil {
		toSerialize["CompletedPhases"] = o.CompletedPhases
	}
	if o.ConfigurationMoId != nil {
		toSerialize["ConfigurationMoId"] = o.ConfigurationMoId
	}
	if o.CurrentPhase.IsSet() {
		toSerialize["CurrentPhase"] = o.CurrentPhase.Get()
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.LastSuccessPollTime != nil {
		toSerialize["LastSuccessPollTime"] = o.LastSuccessPollTime
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceDeviceCertificate) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceDeviceCertificateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The base64 encoded certificate in PEM format.
		CaCertificate *string `json:"CaCertificate,omitempty"`
		// The expiry datetime of new ca certificate which need to be applied on device connector.
		CaCertificateExpiryTime *time.Time `json:"CaCertificateExpiryTime,omitempty"`
		// The date time allocated till cert renewal will be executed. This time used here will be based on cert renewal plan.
		CertificateRenewalExpiryTime *time.Time                  `json:"CertificateRenewalExpiryTime,omitempty"`
		CompletedPhases              []ApplianceCertRenewalPhase `json:"CompletedPhases,omitempty"`
		// The operation configuration MOId.
		ConfigurationMoId *string                           `json:"ConfigurationMoId,omitempty"`
		CurrentPhase      NullableApplianceCertRenewalPhase `json:"CurrentPhase,omitempty"`
		// End date of the certificate renewal.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle.
		LastSuccessPollTime *time.Time `json:"LastSuccessPollTime,omitempty"`
		Messages            []string   `json:"Messages,omitempty"`
		// Start date of the certificate renewal.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// The status of ca certificate renewal.
		Status *string `json:"Status,omitempty"`
	}

	varApplianceDeviceCertificateWithoutEmbeddedStruct := ApplianceDeviceCertificateWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceDeviceCertificateWithoutEmbeddedStruct)
	if err == nil {
		varApplianceDeviceCertificate := _ApplianceDeviceCertificate{}
		varApplianceDeviceCertificate.ClassId = varApplianceDeviceCertificateWithoutEmbeddedStruct.ClassId
		varApplianceDeviceCertificate.ObjectType = varApplianceDeviceCertificateWithoutEmbeddedStruct.ObjectType
		varApplianceDeviceCertificate.CaCertificate = varApplianceDeviceCertificateWithoutEmbeddedStruct.CaCertificate
		varApplianceDeviceCertificate.CaCertificateExpiryTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.CaCertificateExpiryTime
		varApplianceDeviceCertificate.CertificateRenewalExpiryTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.CertificateRenewalExpiryTime
		varApplianceDeviceCertificate.CompletedPhases = varApplianceDeviceCertificateWithoutEmbeddedStruct.CompletedPhases
		varApplianceDeviceCertificate.ConfigurationMoId = varApplianceDeviceCertificateWithoutEmbeddedStruct.ConfigurationMoId
		varApplianceDeviceCertificate.CurrentPhase = varApplianceDeviceCertificateWithoutEmbeddedStruct.CurrentPhase
		varApplianceDeviceCertificate.EndTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.EndTime
		varApplianceDeviceCertificate.LastSuccessPollTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.LastSuccessPollTime
		varApplianceDeviceCertificate.Messages = varApplianceDeviceCertificateWithoutEmbeddedStruct.Messages
		varApplianceDeviceCertificate.StartTime = varApplianceDeviceCertificateWithoutEmbeddedStruct.StartTime
		varApplianceDeviceCertificate.Status = varApplianceDeviceCertificateWithoutEmbeddedStruct.Status
		*o = ApplianceDeviceCertificate(varApplianceDeviceCertificate)
	} else {
		return err
	}

	varApplianceDeviceCertificate := _ApplianceDeviceCertificate{}

	err = json.Unmarshal(bytes, &varApplianceDeviceCertificate)
	if err == nil {
		o.MoBaseMo = varApplianceDeviceCertificate.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCertificate")
		delete(additionalProperties, "CaCertificateExpiryTime")
		delete(additionalProperties, "CertificateRenewalExpiryTime")
		delete(additionalProperties, "CompletedPhases")
		delete(additionalProperties, "ConfigurationMoId")
		delete(additionalProperties, "CurrentPhase")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "LastSuccessPollTime")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")

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

type NullableApplianceDeviceCertificate struct {
	value *ApplianceDeviceCertificate
	isSet bool
}

func (v NullableApplianceDeviceCertificate) Get() *ApplianceDeviceCertificate {
	return v.value
}

func (v *NullableApplianceDeviceCertificate) Set(val *ApplianceDeviceCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDeviceCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDeviceCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDeviceCertificate(val *ApplianceDeviceCertificate) *NullableApplianceDeviceCertificate {
	return &NullableApplianceDeviceCertificate{value: val, isSet: true}
}

func (v NullableApplianceDeviceCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDeviceCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}