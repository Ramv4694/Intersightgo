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
	"time"
)

// TechsupportmanagementTechSupportStatus The techsupport collection status.
type TechsupportmanagementTechSupportStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the Techsupport bundle file.
	FileName *string `json:"FileName,omitempty"`
	// Reason for techsupport failure, if any.
	Reason *string `json:"Reason,omitempty"`
	// Reason for status relay failure, if any.
	RelayReason *string `json:"RelayReason,omitempty"`
	// Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed.
	RelayStatus *string `json:"RelayStatus,omitempty"`
	// The time at which the techsupport request was initiated.
	RequestTs *time.Time `json:"RequestTs,omitempty"`
	// Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection.
	Status *string `json:"Status,omitempty"`
	// The Url to download the techsupport file.
	TechsupportDownloadUrl *string                                             `json:"TechsupportDownloadUrl,omitempty"`
	ClusterMember          *AssetClusterMemberRelationship                     `json:"ClusterMember,omitempty"`
	DeviceRegistration     *AssetDeviceRegistrationRelationship                `json:"DeviceRegistration,omitempty"`
	OriginResource         *MoBaseMoRelationship                               `json:"OriginResource,omitempty"`
	TechSupportRequest     *TechsupportmanagementTechSupportBundleRelationship `json:"TechSupportRequest,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _TechsupportmanagementTechSupportStatus TechsupportmanagementTechSupportStatus

// NewTechsupportmanagementTechSupportStatus instantiates a new TechsupportmanagementTechSupportStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementTechSupportStatus(classId string, objectType string) *TechsupportmanagementTechSupportStatus {
	this := TechsupportmanagementTechSupportStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementTechSupportStatusWithDefaults instantiates a new TechsupportmanagementTechSupportStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementTechSupportStatusWithDefaults() *TechsupportmanagementTechSupportStatus {
	this := TechsupportmanagementTechSupportStatus{}
	var classId string = "techsupportmanagement.TechSupportStatus"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.TechSupportStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementTechSupportStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementTechSupportStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementTechSupportStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementTechSupportStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *TechsupportmanagementTechSupportStatus) SetFileName(v string) {
	o.FileName = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TechsupportmanagementTechSupportStatus) SetReason(v string) {
	o.Reason = &v
}

// GetRelayReason returns the RelayReason field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRelayReason() string {
	if o == nil || o.RelayReason == nil {
		var ret string
		return ret
	}
	return *o.RelayReason
}

// GetRelayReasonOk returns a tuple with the RelayReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRelayReasonOk() (*string, bool) {
	if o == nil || o.RelayReason == nil {
		return nil, false
	}
	return o.RelayReason, true
}

// HasRelayReason returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRelayReason() bool {
	if o != nil && o.RelayReason != nil {
		return true
	}

	return false
}

// SetRelayReason gets a reference to the given string and assigns it to the RelayReason field.
func (o *TechsupportmanagementTechSupportStatus) SetRelayReason(v string) {
	o.RelayReason = &v
}

// GetRelayStatus returns the RelayStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRelayStatus() string {
	if o == nil || o.RelayStatus == nil {
		var ret string
		return ret
	}
	return *o.RelayStatus
}

// GetRelayStatusOk returns a tuple with the RelayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRelayStatusOk() (*string, bool) {
	if o == nil || o.RelayStatus == nil {
		return nil, false
	}
	return o.RelayStatus, true
}

// HasRelayStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRelayStatus() bool {
	if o != nil && o.RelayStatus != nil {
		return true
	}

	return false
}

// SetRelayStatus gets a reference to the given string and assigns it to the RelayStatus field.
func (o *TechsupportmanagementTechSupportStatus) SetRelayStatus(v string) {
	o.RelayStatus = &v
}

// GetRequestTs returns the RequestTs field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetRequestTs() time.Time {
	if o == nil || o.RequestTs == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestTs
}

// GetRequestTsOk returns a tuple with the RequestTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetRequestTsOk() (*time.Time, bool) {
	if o == nil || o.RequestTs == nil {
		return nil, false
	}
	return o.RequestTs, true
}

// HasRequestTs returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasRequestTs() bool {
	if o != nil && o.RequestTs != nil {
		return true
	}

	return false
}

// SetRequestTs gets a reference to the given time.Time and assigns it to the RequestTs field.
func (o *TechsupportmanagementTechSupportStatus) SetRequestTs(v time.Time) {
	o.RequestTs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TechsupportmanagementTechSupportStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrl() string {
	if o == nil || o.TechsupportDownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.TechsupportDownloadUrl
}

// GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrlOk() (*string, bool) {
	if o == nil || o.TechsupportDownloadUrl == nil {
		return nil, false
	}
	return o.TechsupportDownloadUrl, true
}

// HasTechsupportDownloadUrl returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasTechsupportDownloadUrl() bool {
	if o != nil && o.TechsupportDownloadUrl != nil {
		return true
	}

	return false
}

// SetTechsupportDownloadUrl gets a reference to the given string and assigns it to the TechsupportDownloadUrl field.
func (o *TechsupportmanagementTechSupportStatus) SetTechsupportDownloadUrl(v string) {
	o.TechsupportDownloadUrl = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *TechsupportmanagementTechSupportStatus) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TechsupportmanagementTechSupportStatus) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetOriginResource returns the OriginResource field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetOriginResource() MoBaseMoRelationship {
	if o == nil || o.OriginResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.OriginResource
}

// GetOriginResourceOk returns a tuple with the OriginResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetOriginResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.OriginResource == nil {
		return nil, false
	}
	return o.OriginResource, true
}

// HasOriginResource returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasOriginResource() bool {
	if o != nil && o.OriginResource != nil {
		return true
	}

	return false
}

// SetOriginResource gets a reference to the given MoBaseMoRelationship and assigns it to the OriginResource field.
func (o *TechsupportmanagementTechSupportStatus) SetOriginResource(v MoBaseMoRelationship) {
	o.OriginResource = &v
}

// GetTechSupportRequest returns the TechSupportRequest field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequest() TechsupportmanagementTechSupportBundleRelationship {
	if o == nil || o.TechSupportRequest == nil {
		var ret TechsupportmanagementTechSupportBundleRelationship
		return ret
	}
	return *o.TechSupportRequest
}

// GetTechSupportRequestOk returns a tuple with the TechSupportRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequestOk() (*TechsupportmanagementTechSupportBundleRelationship, bool) {
	if o == nil || o.TechSupportRequest == nil {
		return nil, false
	}
	return o.TechSupportRequest, true
}

// HasTechSupportRequest returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportStatus) HasTechSupportRequest() bool {
	if o != nil && o.TechSupportRequest != nil {
		return true
	}

	return false
}

// SetTechSupportRequest gets a reference to the given TechsupportmanagementTechSupportBundleRelationship and assigns it to the TechSupportRequest field.
func (o *TechsupportmanagementTechSupportStatus) SetTechSupportRequest(v TechsupportmanagementTechSupportBundleRelationship) {
	o.TechSupportRequest = &v
}

func (o TechsupportmanagementTechSupportStatus) MarshalJSON() ([]byte, error) {
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
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.RelayReason != nil {
		toSerialize["RelayReason"] = o.RelayReason
	}
	if o.RelayStatus != nil {
		toSerialize["RelayStatus"] = o.RelayStatus
	}
	if o.RequestTs != nil {
		toSerialize["RequestTs"] = o.RequestTs
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TechsupportDownloadUrl != nil {
		toSerialize["TechsupportDownloadUrl"] = o.TechsupportDownloadUrl
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.OriginResource != nil {
		toSerialize["OriginResource"] = o.OriginResource
	}
	if o.TechSupportRequest != nil {
		toSerialize["TechSupportRequest"] = o.TechSupportRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementTechSupportStatus) UnmarshalJSON(bytes []byte) (err error) {
	type TechsupportmanagementTechSupportStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the Techsupport bundle file.
		FileName *string `json:"FileName,omitempty"`
		// Reason for techsupport failure, if any.
		Reason *string `json:"Reason,omitempty"`
		// Reason for status relay failure, if any.
		RelayReason *string `json:"RelayReason,omitempty"`
		// Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed.
		RelayStatus *string `json:"RelayStatus,omitempty"`
		// The time at which the techsupport request was initiated.
		RequestTs *time.Time `json:"RequestTs,omitempty"`
		// Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection.
		Status *string `json:"Status,omitempty"`
		// The Url to download the techsupport file.
		TechsupportDownloadUrl *string                                             `json:"TechsupportDownloadUrl,omitempty"`
		ClusterMember          *AssetClusterMemberRelationship                     `json:"ClusterMember,omitempty"`
		DeviceRegistration     *AssetDeviceRegistrationRelationship                `json:"DeviceRegistration,omitempty"`
		OriginResource         *MoBaseMoRelationship                               `json:"OriginResource,omitempty"`
		TechSupportRequest     *TechsupportmanagementTechSupportBundleRelationship `json:"TechSupportRequest,omitempty"`
	}

	varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct := TechsupportmanagementTechSupportStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct)
	if err == nil {
		varTechsupportmanagementTechSupportStatus := _TechsupportmanagementTechSupportStatus{}
		varTechsupportmanagementTechSupportStatus.ClassId = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ClassId
		varTechsupportmanagementTechSupportStatus.ObjectType = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ObjectType
		varTechsupportmanagementTechSupportStatus.FileName = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.FileName
		varTechsupportmanagementTechSupportStatus.Reason = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.Reason
		varTechsupportmanagementTechSupportStatus.RelayReason = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RelayReason
		varTechsupportmanagementTechSupportStatus.RelayStatus = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RelayStatus
		varTechsupportmanagementTechSupportStatus.RequestTs = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.RequestTs
		varTechsupportmanagementTechSupportStatus.Status = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.Status
		varTechsupportmanagementTechSupportStatus.TechsupportDownloadUrl = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.TechsupportDownloadUrl
		varTechsupportmanagementTechSupportStatus.ClusterMember = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.ClusterMember
		varTechsupportmanagementTechSupportStatus.DeviceRegistration = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.DeviceRegistration
		varTechsupportmanagementTechSupportStatus.OriginResource = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.OriginResource
		varTechsupportmanagementTechSupportStatus.TechSupportRequest = varTechsupportmanagementTechSupportStatusWithoutEmbeddedStruct.TechSupportRequest
		*o = TechsupportmanagementTechSupportStatus(varTechsupportmanagementTechSupportStatus)
	} else {
		return err
	}

	varTechsupportmanagementTechSupportStatus := _TechsupportmanagementTechSupportStatus{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementTechSupportStatus)
	if err == nil {
		o.MoBaseMo = varTechsupportmanagementTechSupportStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "RelayReason")
		delete(additionalProperties, "RelayStatus")
		delete(additionalProperties, "RequestTs")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TechsupportDownloadUrl")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "OriginResource")
		delete(additionalProperties, "TechSupportRequest")

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

type NullableTechsupportmanagementTechSupportStatus struct {
	value *TechsupportmanagementTechSupportStatus
	isSet bool
}

func (v NullableTechsupportmanagementTechSupportStatus) Get() *TechsupportmanagementTechSupportStatus {
	return v.value
}

func (v *NullableTechsupportmanagementTechSupportStatus) Set(val *TechsupportmanagementTechSupportStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementTechSupportStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementTechSupportStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementTechSupportStatus(val *TechsupportmanagementTechSupportStatus) *NullableTechsupportmanagementTechSupportStatus {
	return &NullableTechsupportmanagementTechSupportStatus{value: val, isSet: true}
}

func (v NullableTechsupportmanagementTechSupportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementTechSupportStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
