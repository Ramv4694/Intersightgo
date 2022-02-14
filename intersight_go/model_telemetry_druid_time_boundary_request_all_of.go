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

// TelemetryDruidTimeBoundaryRequestAllOf struct for TelemetryDruidTimeBoundaryRequestAllOf
type TelemetryDruidTimeBoundaryRequestAllOf struct {
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// Optional, set to maxTime or minTime to return only the latest or earliest timestamp. Default to returning both if not set.
	Bound                *string                     `json:"bound,omitempty"`
	Filter               *TelemetryDruidFilter       `json:"filter,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTimeBoundaryRequestAllOf TelemetryDruidTimeBoundaryRequestAllOf

// NewTelemetryDruidTimeBoundaryRequestAllOf instantiates a new TelemetryDruidTimeBoundaryRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTimeBoundaryRequestAllOf(dataSource TelemetryDruidDataSource) *TelemetryDruidTimeBoundaryRequestAllOf {
	this := TelemetryDruidTimeBoundaryRequestAllOf{}
	this.DataSource = dataSource
	return &this
}

// NewTelemetryDruidTimeBoundaryRequestAllOfWithDefaults instantiates a new TelemetryDruidTimeBoundaryRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTimeBoundaryRequestAllOfWithDefaults() *TelemetryDruidTimeBoundaryRequestAllOf {
	this := TelemetryDruidTimeBoundaryRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidTimeBoundaryRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetBound returns the Bound field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetBound() string {
	if o == nil || o.Bound == nil {
		var ret string
		return ret
	}
	return *o.Bound
}

// GetBoundOk returns a tuple with the Bound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetBoundOk() (*string, bool) {
	if o == nil || o.Bound == nil {
		return nil, false
	}
	return o.Bound, true
}

// HasBound returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) HasBound() bool {
	if o != nil && o.Bound != nil {
		return true
	}

	return false
}

// SetBound gets a reference to the given string and assigns it to the Bound field.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) SetBound(v string) {
	o.Bound = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidTimeBoundaryRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidTimeBoundaryRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.Bound != nil {
		toSerialize["bound"] = o.Bound
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidTimeBoundaryRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidTimeBoundaryRequestAllOf := _TelemetryDruidTimeBoundaryRequestAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidTimeBoundaryRequestAllOf); err == nil {
		*o = TelemetryDruidTimeBoundaryRequestAllOf(varTelemetryDruidTimeBoundaryRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "bound")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTimeBoundaryRequestAllOf struct {
	value *TelemetryDruidTimeBoundaryRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidTimeBoundaryRequestAllOf) Get() *TelemetryDruidTimeBoundaryRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidTimeBoundaryRequestAllOf) Set(val *TelemetryDruidTimeBoundaryRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTimeBoundaryRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTimeBoundaryRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTimeBoundaryRequestAllOf(val *TelemetryDruidTimeBoundaryRequestAllOf) *NullableTelemetryDruidTimeBoundaryRequestAllOf {
	return &NullableTelemetryDruidTimeBoundaryRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidTimeBoundaryRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTimeBoundaryRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}