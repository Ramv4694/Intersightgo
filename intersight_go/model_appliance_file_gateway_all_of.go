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
	"time"
)

// ApplianceFileGatewayAllOf Definition of the list of properties defined in 'appliance.FileGateway', excluding properties defined in parent classes.
type ApplianceFileGatewayAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bucket name in the cloud storage service where the file is stored.
	BucketName *string `json:"BucketName,omitempty"`
	// Size of the file in bytes. FileSize maybe zero if the storage service does not report file size.
	FileSize *int64 `json:"FileSize,omitempty"`
	// File timestamp as reported by the cloud storage service.
	FileTime *time.Time `json:"FileTime,omitempty"`
	// User defined file type supplied by the caller.
	FileType *string `json:"FileType,omitempty"`
	// Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path.
	Filename *string `json:"Filename,omitempty"`
	// Pre-signed URL obtained from the cloud storage service.
	PresignedUrl *string `json:"PresignedUrl,omitempty"`
	// SSL certificate of the cloud storage service.
	ServerCert *string `json:"ServerCert,omitempty"`
	// Signed URL's validity period in minutes.
	ValidityPeriod *int64 `json:"ValidityPeriod,omitempty"`
	// File version as reported by the cloud storage service.
	Version              *string                 `json:"Version,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceFileGatewayAllOf ApplianceFileGatewayAllOf

// NewApplianceFileGatewayAllOf instantiates a new ApplianceFileGatewayAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceFileGatewayAllOf(classId string, objectType string) *ApplianceFileGatewayAllOf {
	this := ApplianceFileGatewayAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceFileGatewayAllOfWithDefaults instantiates a new ApplianceFileGatewayAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceFileGatewayAllOfWithDefaults() *ApplianceFileGatewayAllOf {
	this := ApplianceFileGatewayAllOf{}
	var classId string = "appliance.FileGateway"
	this.ClassId = classId
	var objectType string = "appliance.FileGateway"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceFileGatewayAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceFileGatewayAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceFileGatewayAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceFileGatewayAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *ApplianceFileGatewayAllOf) SetBucketName(v string) {
	o.BucketName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetFileSize() int64 {
	if o == nil || o.FileSize == nil {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetFileSizeOk() (*int64, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *ApplianceFileGatewayAllOf) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetFileTime returns the FileTime field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetFileTime() time.Time {
	if o == nil || o.FileTime == nil {
		var ret time.Time
		return ret
	}
	return *o.FileTime
}

// GetFileTimeOk returns a tuple with the FileTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetFileTimeOk() (*time.Time, bool) {
	if o == nil || o.FileTime == nil {
		return nil, false
	}
	return o.FileTime, true
}

// HasFileTime returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasFileTime() bool {
	if o != nil && o.FileTime != nil {
		return true
	}

	return false
}

// SetFileTime gets a reference to the given time.Time and assigns it to the FileTime field.
func (o *ApplianceFileGatewayAllOf) SetFileTime(v time.Time) {
	o.FileTime = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *ApplianceFileGatewayAllOf) SetFileType(v string) {
	o.FileType = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceFileGatewayAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetPresignedUrl() string {
	if o == nil || o.PresignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetPresignedUrlOk() (*string, bool) {
	if o == nil || o.PresignedUrl == nil {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl != nil {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *ApplianceFileGatewayAllOf) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

// GetServerCert returns the ServerCert field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetServerCert() string {
	if o == nil || o.ServerCert == nil {
		var ret string
		return ret
	}
	return *o.ServerCert
}

// GetServerCertOk returns a tuple with the ServerCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetServerCertOk() (*string, bool) {
	if o == nil || o.ServerCert == nil {
		return nil, false
	}
	return o.ServerCert, true
}

// HasServerCert returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasServerCert() bool {
	if o != nil && o.ServerCert != nil {
		return true
	}

	return false
}

// SetServerCert gets a reference to the given string and assigns it to the ServerCert field.
func (o *ApplianceFileGatewayAllOf) SetServerCert(v string) {
	o.ServerCert = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetValidityPeriod() int64 {
	if o == nil || o.ValidityPeriod == nil {
		var ret int64
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetValidityPeriodOk() (*int64, bool) {
	if o == nil || o.ValidityPeriod == nil {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasValidityPeriod() bool {
	if o != nil && o.ValidityPeriod != nil {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given int64 and assigns it to the ValidityPeriod field.
func (o *ApplianceFileGatewayAllOf) SetValidityPeriod(v int64) {
	o.ValidityPeriod = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceFileGatewayAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceFileGatewayAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGatewayAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceFileGatewayAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceFileGatewayAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceFileGatewayAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BucketName != nil {
		toSerialize["BucketName"] = o.BucketName
	}
	if o.FileSize != nil {
		toSerialize["FileSize"] = o.FileSize
	}
	if o.FileTime != nil {
		toSerialize["FileTime"] = o.FileTime
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.PresignedUrl != nil {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}
	if o.ServerCert != nil {
		toSerialize["ServerCert"] = o.ServerCert
	}
	if o.ValidityPeriod != nil {
		toSerialize["ValidityPeriod"] = o.ValidityPeriod
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceFileGatewayAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceFileGatewayAllOf := _ApplianceFileGatewayAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceFileGatewayAllOf); err == nil {
		*o = ApplianceFileGatewayAllOf(varApplianceFileGatewayAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "FileTime")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "PresignedUrl")
		delete(additionalProperties, "ServerCert")
		delete(additionalProperties, "ValidityPeriod")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceFileGatewayAllOf struct {
	value *ApplianceFileGatewayAllOf
	isSet bool
}

func (v NullableApplianceFileGatewayAllOf) Get() *ApplianceFileGatewayAllOf {
	return v.value
}

func (v *NullableApplianceFileGatewayAllOf) Set(val *ApplianceFileGatewayAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFileGatewayAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFileGatewayAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFileGatewayAllOf(val *ApplianceFileGatewayAllOf) *NullableApplianceFileGatewayAllOf {
	return &NullableApplianceFileGatewayAllOf{value: val, isSet: true}
}

func (v NullableApplianceFileGatewayAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFileGatewayAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
