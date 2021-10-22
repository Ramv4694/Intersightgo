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
)

// NiaapiNewReleaseDetailAllOf Definition of the list of properties defined in 'niaapi.NewReleaseDetail', excluding properties defined in parent classes.
type NiaapiNewReleaseDetailAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this new verison release post.
	Description *string `json:"Description,omitempty"`
	// Link of downloading the release file.
	Link *string `json:"Link,omitempty"`
	// The link used to provide a gateway for customer to review the release note.
	ReleaseNoteLink *string `json:"ReleaseNoteLink,omitempty"`
	// The link title used to provide a gateway for customer to review the release note.
	ReleaseNoteLinkTitle *string `json:"ReleaseNoteLinkTitle,omitempty"`
	// The link used to provide a gateway for customer to download the release.
	SoftwareDownloadLink *string `json:"SoftwareDownloadLink,omitempty"`
	// The link title used to provide a gateway for customer to download the release.
	SoftwareDownloadLinkTitle *string `json:"SoftwareDownloadLinkTitle,omitempty"`
	// Title of the new verison release post.
	Title *string `json:"Title,omitempty"`
	// Version number Associate with this Post.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNewReleaseDetailAllOf NiaapiNewReleaseDetailAllOf

// NewNiaapiNewReleaseDetailAllOf instantiates a new NiaapiNewReleaseDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNewReleaseDetailAllOf(classId string, objectType string) *NiaapiNewReleaseDetailAllOf {
	this := NiaapiNewReleaseDetailAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNewReleaseDetailAllOfWithDefaults instantiates a new NiaapiNewReleaseDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNewReleaseDetailAllOfWithDefaults() *NiaapiNewReleaseDetailAllOf {
	this := NiaapiNewReleaseDetailAllOf{}
	var classId string = "niaapi.NewReleaseDetail"
	this.ClassId = classId
	var objectType string = "niaapi.NewReleaseDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNewReleaseDetailAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNewReleaseDetailAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNewReleaseDetailAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNewReleaseDetailAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiaapiNewReleaseDetailAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *NiaapiNewReleaseDetailAllOf) SetLink(v string) {
	o.Link = &v
}

// GetReleaseNoteLink returns the ReleaseNoteLink field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetReleaseNoteLink() string {
	if o == nil || o.ReleaseNoteLink == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNoteLink
}

// GetReleaseNoteLinkOk returns a tuple with the ReleaseNoteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetReleaseNoteLinkOk() (*string, bool) {
	if o == nil || o.ReleaseNoteLink == nil {
		return nil, false
	}
	return o.ReleaseNoteLink, true
}

// HasReleaseNoteLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasReleaseNoteLink() bool {
	if o != nil && o.ReleaseNoteLink != nil {
		return true
	}

	return false
}

// SetReleaseNoteLink gets a reference to the given string and assigns it to the ReleaseNoteLink field.
func (o *NiaapiNewReleaseDetailAllOf) SetReleaseNoteLink(v string) {
	o.ReleaseNoteLink = &v
}

// GetReleaseNoteLinkTitle returns the ReleaseNoteLinkTitle field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetReleaseNoteLinkTitle() string {
	if o == nil || o.ReleaseNoteLinkTitle == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNoteLinkTitle
}

// GetReleaseNoteLinkTitleOk returns a tuple with the ReleaseNoteLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetReleaseNoteLinkTitleOk() (*string, bool) {
	if o == nil || o.ReleaseNoteLinkTitle == nil {
		return nil, false
	}
	return o.ReleaseNoteLinkTitle, true
}

// HasReleaseNoteLinkTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasReleaseNoteLinkTitle() bool {
	if o != nil && o.ReleaseNoteLinkTitle != nil {
		return true
	}

	return false
}

// SetReleaseNoteLinkTitle gets a reference to the given string and assigns it to the ReleaseNoteLinkTitle field.
func (o *NiaapiNewReleaseDetailAllOf) SetReleaseNoteLinkTitle(v string) {
	o.ReleaseNoteLinkTitle = &v
}

// GetSoftwareDownloadLink returns the SoftwareDownloadLink field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetSoftwareDownloadLink() string {
	if o == nil || o.SoftwareDownloadLink == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownloadLink
}

// GetSoftwareDownloadLinkOk returns a tuple with the SoftwareDownloadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetSoftwareDownloadLinkOk() (*string, bool) {
	if o == nil || o.SoftwareDownloadLink == nil {
		return nil, false
	}
	return o.SoftwareDownloadLink, true
}

// HasSoftwareDownloadLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasSoftwareDownloadLink() bool {
	if o != nil && o.SoftwareDownloadLink != nil {
		return true
	}

	return false
}

// SetSoftwareDownloadLink gets a reference to the given string and assigns it to the SoftwareDownloadLink field.
func (o *NiaapiNewReleaseDetailAllOf) SetSoftwareDownloadLink(v string) {
	o.SoftwareDownloadLink = &v
}

// GetSoftwareDownloadLinkTitle returns the SoftwareDownloadLinkTitle field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetSoftwareDownloadLinkTitle() string {
	if o == nil || o.SoftwareDownloadLinkTitle == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownloadLinkTitle
}

// GetSoftwareDownloadLinkTitleOk returns a tuple with the SoftwareDownloadLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetSoftwareDownloadLinkTitleOk() (*string, bool) {
	if o == nil || o.SoftwareDownloadLinkTitle == nil {
		return nil, false
	}
	return o.SoftwareDownloadLinkTitle, true
}

// HasSoftwareDownloadLinkTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasSoftwareDownloadLinkTitle() bool {
	if o != nil && o.SoftwareDownloadLinkTitle != nil {
		return true
	}

	return false
}

// SetSoftwareDownloadLinkTitle gets a reference to the given string and assigns it to the SoftwareDownloadLinkTitle field.
func (o *NiaapiNewReleaseDetailAllOf) SetSoftwareDownloadLinkTitle(v string) {
	o.SoftwareDownloadLinkTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NiaapiNewReleaseDetailAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetailAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetailAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetailAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiaapiNewReleaseDetailAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o NiaapiNewReleaseDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Link != nil {
		toSerialize["Link"] = o.Link
	}
	if o.ReleaseNoteLink != nil {
		toSerialize["ReleaseNoteLink"] = o.ReleaseNoteLink
	}
	if o.ReleaseNoteLinkTitle != nil {
		toSerialize["ReleaseNoteLinkTitle"] = o.ReleaseNoteLinkTitle
	}
	if o.SoftwareDownloadLink != nil {
		toSerialize["SoftwareDownloadLink"] = o.SoftwareDownloadLink
	}
	if o.SoftwareDownloadLinkTitle != nil {
		toSerialize["SoftwareDownloadLinkTitle"] = o.SoftwareDownloadLinkTitle
	}
	if o.Title != nil {
		toSerialize["Title"] = o.Title
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNewReleaseDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiNewReleaseDetailAllOf := _NiaapiNewReleaseDetailAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiNewReleaseDetailAllOf); err == nil {
		*o = NiaapiNewReleaseDetailAllOf(varNiaapiNewReleaseDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Link")
		delete(additionalProperties, "ReleaseNoteLink")
		delete(additionalProperties, "ReleaseNoteLinkTitle")
		delete(additionalProperties, "SoftwareDownloadLink")
		delete(additionalProperties, "SoftwareDownloadLinkTitle")
		delete(additionalProperties, "Title")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiNewReleaseDetailAllOf struct {
	value *NiaapiNewReleaseDetailAllOf
	isSet bool
}

func (v NullableNiaapiNewReleaseDetailAllOf) Get() *NiaapiNewReleaseDetailAllOf {
	return v.value
}

func (v *NullableNiaapiNewReleaseDetailAllOf) Set(val *NiaapiNewReleaseDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNewReleaseDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNewReleaseDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNewReleaseDetailAllOf(val *NiaapiNewReleaseDetailAllOf) *NullableNiaapiNewReleaseDetailAllOf {
	return &NullableNiaapiNewReleaseDetailAllOf{value: val, isSet: true}
}

func (v NullableNiaapiNewReleaseDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNewReleaseDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
