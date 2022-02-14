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
)

// NiaapiFileDownloaderAllOf Definition of the list of properties defined in 'niaapi.FileDownloader', excluding properties defined in parent classes.
type NiaapiFileDownloaderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Filename of this Metadata package file, folder will be handled by api.
	FileName *string `json:"FileName,omitempty"`
	// The presigned URL from server to download this file.
	PresignedUrl         *string `json:"PresignedUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiFileDownloaderAllOf NiaapiFileDownloaderAllOf

// NewNiaapiFileDownloaderAllOf instantiates a new NiaapiFileDownloaderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiFileDownloaderAllOf(classId string, objectType string) *NiaapiFileDownloaderAllOf {
	this := NiaapiFileDownloaderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiFileDownloaderAllOfWithDefaults instantiates a new NiaapiFileDownloaderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiFileDownloaderAllOfWithDefaults() *NiaapiFileDownloaderAllOf {
	this := NiaapiFileDownloaderAllOf{}
	var classId string = "niaapi.FileDownloader"
	this.ClassId = classId
	var objectType string = "niaapi.FileDownloader"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiFileDownloaderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloaderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiFileDownloaderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiFileDownloaderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloaderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiFileDownloaderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NiaapiFileDownloaderAllOf) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloaderAllOf) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NiaapiFileDownloaderAllOf) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NiaapiFileDownloaderAllOf) SetFileName(v string) {
	o.FileName = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *NiaapiFileDownloaderAllOf) GetPresignedUrl() string {
	if o == nil || o.PresignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloaderAllOf) GetPresignedUrlOk() (*string, bool) {
	if o == nil || o.PresignedUrl == nil {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *NiaapiFileDownloaderAllOf) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl != nil {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *NiaapiFileDownloaderAllOf) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

func (o NiaapiFileDownloaderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.PresignedUrl != nil {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiFileDownloaderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiFileDownloaderAllOf := _NiaapiFileDownloaderAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiFileDownloaderAllOf); err == nil {
		*o = NiaapiFileDownloaderAllOf(varNiaapiFileDownloaderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "PresignedUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiFileDownloaderAllOf struct {
	value *NiaapiFileDownloaderAllOf
	isSet bool
}

func (v NullableNiaapiFileDownloaderAllOf) Get() *NiaapiFileDownloaderAllOf {
	return v.value
}

func (v *NullableNiaapiFileDownloaderAllOf) Set(val *NiaapiFileDownloaderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiFileDownloaderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiFileDownloaderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiFileDownloaderAllOf(val *NiaapiFileDownloaderAllOf) *NullableNiaapiFileDownloaderAllOf {
	return &NullableNiaapiFileDownloaderAllOf{value: val, isSet: true}
}

func (v NullableNiaapiFileDownloaderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiFileDownloaderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
