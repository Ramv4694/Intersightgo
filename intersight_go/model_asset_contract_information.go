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

// AssetContractInformation Type for saving contract information.
type AssetContractInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                          `json:"ObjectType"`
	BillTo               NullableAssetAddressInformation `json:"BillTo,omitempty"`
	BillToGlobalUltimate NullableAssetGlobalUltimate     `json:"BillToGlobalUltimate,omitempty"`
	// Contract number for the Cisco support contract purchased for the Cisco device.
	ContractNumber *string `json:"ContractNumber,omitempty"`
	// Contract status as per the Cisco Contract APIx.
	LineStatus           *string `json:"LineStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetContractInformation AssetContractInformation

// NewAssetContractInformation instantiates a new AssetContractInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetContractInformation(classId string, objectType string) *AssetContractInformation {
	this := AssetContractInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetContractInformationWithDefaults instantiates a new AssetContractInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetContractInformationWithDefaults() *AssetContractInformation {
	this := AssetContractInformation{}
	var classId string = "asset.ContractInformation"
	this.ClassId = classId
	var objectType string = "asset.ContractInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetContractInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetContractInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetContractInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetContractInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetContractInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetContractInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillTo returns the BillTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetContractInformation) GetBillTo() AssetAddressInformation {
	if o == nil || o.BillTo.Get() == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.BillTo.Get()
}

// GetBillToOk returns a tuple with the BillTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetContractInformation) GetBillToOk() (*AssetAddressInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillTo.Get(), o.BillTo.IsSet()
}

// HasBillTo returns a boolean if a field has been set.
func (o *AssetContractInformation) HasBillTo() bool {
	if o != nil && o.BillTo.IsSet() {
		return true
	}

	return false
}

// SetBillTo gets a reference to the given NullableAssetAddressInformation and assigns it to the BillTo field.
func (o *AssetContractInformation) SetBillTo(v AssetAddressInformation) {
	o.BillTo.Set(&v)
}

// SetBillToNil sets the value for BillTo to be an explicit nil
func (o *AssetContractInformation) SetBillToNil() {
	o.BillTo.Set(nil)
}

// UnsetBillTo ensures that no value is present for BillTo, not even an explicit nil
func (o *AssetContractInformation) UnsetBillTo() {
	o.BillTo.Unset()
}

// GetBillToGlobalUltimate returns the BillToGlobalUltimate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetContractInformation) GetBillToGlobalUltimate() AssetGlobalUltimate {
	if o == nil || o.BillToGlobalUltimate.Get() == nil {
		var ret AssetGlobalUltimate
		return ret
	}
	return *o.BillToGlobalUltimate.Get()
}

// GetBillToGlobalUltimateOk returns a tuple with the BillToGlobalUltimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetContractInformation) GetBillToGlobalUltimateOk() (*AssetGlobalUltimate, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillToGlobalUltimate.Get(), o.BillToGlobalUltimate.IsSet()
}

// HasBillToGlobalUltimate returns a boolean if a field has been set.
func (o *AssetContractInformation) HasBillToGlobalUltimate() bool {
	if o != nil && o.BillToGlobalUltimate.IsSet() {
		return true
	}

	return false
}

// SetBillToGlobalUltimate gets a reference to the given NullableAssetGlobalUltimate and assigns it to the BillToGlobalUltimate field.
func (o *AssetContractInformation) SetBillToGlobalUltimate(v AssetGlobalUltimate) {
	o.BillToGlobalUltimate.Set(&v)
}

// SetBillToGlobalUltimateNil sets the value for BillToGlobalUltimate to be an explicit nil
func (o *AssetContractInformation) SetBillToGlobalUltimateNil() {
	o.BillToGlobalUltimate.Set(nil)
}

// UnsetBillToGlobalUltimate ensures that no value is present for BillToGlobalUltimate, not even an explicit nil
func (o *AssetContractInformation) UnsetBillToGlobalUltimate() {
	o.BillToGlobalUltimate.Unset()
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *AssetContractInformation) GetContractNumber() string {
	if o == nil || o.ContractNumber == nil {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformation) GetContractNumberOk() (*string, bool) {
	if o == nil || o.ContractNumber == nil {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *AssetContractInformation) HasContractNumber() bool {
	if o != nil && o.ContractNumber != nil {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *AssetContractInformation) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetLineStatus returns the LineStatus field value if set, zero value otherwise.
func (o *AssetContractInformation) GetLineStatus() string {
	if o == nil || o.LineStatus == nil {
		var ret string
		return ret
	}
	return *o.LineStatus
}

// GetLineStatusOk returns a tuple with the LineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetContractInformation) GetLineStatusOk() (*string, bool) {
	if o == nil || o.LineStatus == nil {
		return nil, false
	}
	return o.LineStatus, true
}

// HasLineStatus returns a boolean if a field has been set.
func (o *AssetContractInformation) HasLineStatus() bool {
	if o != nil && o.LineStatus != nil {
		return true
	}

	return false
}

// SetLineStatus gets a reference to the given string and assigns it to the LineStatus field.
func (o *AssetContractInformation) SetLineStatus(v string) {
	o.LineStatus = &v
}

func (o AssetContractInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillTo.IsSet() {
		toSerialize["BillTo"] = o.BillTo.Get()
	}
	if o.BillToGlobalUltimate.IsSet() {
		toSerialize["BillToGlobalUltimate"] = o.BillToGlobalUltimate.Get()
	}
	if o.ContractNumber != nil {
		toSerialize["ContractNumber"] = o.ContractNumber
	}
	if o.LineStatus != nil {
		toSerialize["LineStatus"] = o.LineStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetContractInformation) UnmarshalJSON(bytes []byte) (err error) {
	type AssetContractInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                          `json:"ObjectType"`
		BillTo               NullableAssetAddressInformation `json:"BillTo,omitempty"`
		BillToGlobalUltimate NullableAssetGlobalUltimate     `json:"BillToGlobalUltimate,omitempty"`
		// Contract number for the Cisco support contract purchased for the Cisco device.
		ContractNumber *string `json:"ContractNumber,omitempty"`
		// Contract status as per the Cisco Contract APIx.
		LineStatus *string `json:"LineStatus,omitempty"`
	}

	varAssetContractInformationWithoutEmbeddedStruct := AssetContractInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetContractInformationWithoutEmbeddedStruct)
	if err == nil {
		varAssetContractInformation := _AssetContractInformation{}
		varAssetContractInformation.ClassId = varAssetContractInformationWithoutEmbeddedStruct.ClassId
		varAssetContractInformation.ObjectType = varAssetContractInformationWithoutEmbeddedStruct.ObjectType
		varAssetContractInformation.BillTo = varAssetContractInformationWithoutEmbeddedStruct.BillTo
		varAssetContractInformation.BillToGlobalUltimate = varAssetContractInformationWithoutEmbeddedStruct.BillToGlobalUltimate
		varAssetContractInformation.ContractNumber = varAssetContractInformationWithoutEmbeddedStruct.ContractNumber
		varAssetContractInformation.LineStatus = varAssetContractInformationWithoutEmbeddedStruct.LineStatus
		*o = AssetContractInformation(varAssetContractInformation)
	} else {
		return err
	}

	varAssetContractInformation := _AssetContractInformation{}

	err = json.Unmarshal(bytes, &varAssetContractInformation)
	if err == nil {
		o.MoBaseComplexType = varAssetContractInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillTo")
		delete(additionalProperties, "BillToGlobalUltimate")
		delete(additionalProperties, "ContractNumber")
		delete(additionalProperties, "LineStatus")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableAssetContractInformation struct {
	value *AssetContractInformation
	isSet bool
}

func (v NullableAssetContractInformation) Get() *AssetContractInformation {
	return v.value
}

func (v *NullableAssetContractInformation) Set(val *AssetContractInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetContractInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetContractInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetContractInformation(val *AssetContractInformation) *NullableAssetContractInformation {
	return &NullableAssetContractInformation{value: val, isSet: true}
}

func (v NullableAssetContractInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetContractInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
