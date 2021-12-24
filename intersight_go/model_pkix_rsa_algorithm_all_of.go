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
)

// PkixRsaAlgorithmAllOf Definition of the list of properties defined in 'pkix.RsaAlgorithm', excluding properties defined in parent classes.
type PkixRsaAlgorithmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The length of the RSA key, expressed in bits, for both public and private keys. * `2048` - A key length of 2048 bits. * `2560` - A key length of 2560 bits. * `3072` - A key length of 3072 bits. * `3584` - A key length of 3584 bits. * `4096` - A key length of 4096 bits.
	Modulus              *int32 `json:"Modulus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixRsaAlgorithmAllOf PkixRsaAlgorithmAllOf

// NewPkixRsaAlgorithmAllOf instantiates a new PkixRsaAlgorithmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixRsaAlgorithmAllOf(classId string, objectType string) *PkixRsaAlgorithmAllOf {
	this := PkixRsaAlgorithmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// NewPkixRsaAlgorithmAllOfWithDefaults instantiates a new PkixRsaAlgorithmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixRsaAlgorithmAllOfWithDefaults() *PkixRsaAlgorithmAllOf {
	this := PkixRsaAlgorithmAllOf{}
	var classId string = "pkix.RsaAlgorithm"
	this.ClassId = classId
	var objectType string = "pkix.RsaAlgorithm"
	this.ObjectType = objectType
	var modulus int32 = 2048
	this.Modulus = &modulus
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixRsaAlgorithmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixRsaAlgorithmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixRsaAlgorithmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixRsaAlgorithmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *PkixRsaAlgorithmAllOf) GetModulus() int32 {
	if o == nil || o.Modulus == nil {
		var ret int32
		return ret
	}
	return *o.Modulus
}

// GetModulusOk returns a tuple with the Modulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixRsaAlgorithmAllOf) GetModulusOk() (*int32, bool) {
	if o == nil || o.Modulus == nil {
		return nil, false
	}
	return o.Modulus, true
}

// HasModulus returns a boolean if a field has been set.
func (o *PkixRsaAlgorithmAllOf) HasModulus() bool {
	if o != nil && o.Modulus != nil {
		return true
	}

	return false
}

// SetModulus gets a reference to the given int32 and assigns it to the Modulus field.
func (o *PkixRsaAlgorithmAllOf) SetModulus(v int32) {
	o.Modulus = &v
}

func (o PkixRsaAlgorithmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Modulus != nil {
		toSerialize["Modulus"] = o.Modulus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixRsaAlgorithmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPkixRsaAlgorithmAllOf := _PkixRsaAlgorithmAllOf{}

	if err = json.Unmarshal(bytes, &varPkixRsaAlgorithmAllOf); err == nil {
		*o = PkixRsaAlgorithmAllOf(varPkixRsaAlgorithmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Modulus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePkixRsaAlgorithmAllOf struct {
	value *PkixRsaAlgorithmAllOf
	isSet bool
}

func (v NullablePkixRsaAlgorithmAllOf) Get() *PkixRsaAlgorithmAllOf {
	return v.value
}

func (v *NullablePkixRsaAlgorithmAllOf) Set(val *PkixRsaAlgorithmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixRsaAlgorithmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixRsaAlgorithmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixRsaAlgorithmAllOf(val *PkixRsaAlgorithmAllOf) *NullablePkixRsaAlgorithmAllOf {
	return &NullablePkixRsaAlgorithmAllOf{value: val, isSet: true}
}

func (v NullablePkixRsaAlgorithmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixRsaAlgorithmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
