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

// TelemetryDruidTimeSeriesRequest These types of queries take a timeseries query object and return an array of JSON objects where each object represents a value asked for by the timeseries query.
type TelemetryDruidTimeSeriesRequest struct {
	// null
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// Whether to make descending ordered result. Default is false(ascending).
	Descending *bool `json:"descending,omitempty"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals   []string                  `json:"intervals"`
	Granularity TelemetryDruidGranularity `json:"granularity"`
	Filter      *TelemetryDruidFilter     `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations *[]TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires.
	PostAggregations *[]TelemetryDruidPostAggregator `json:"postAggregations,omitempty"`
	// An integer that limits the number of results. The default is unlimited.
	Limit                *int32                      `json:"limit,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTimeSeriesRequest TelemetryDruidTimeSeriesRequest

// NewTelemetryDruidTimeSeriesRequest instantiates a new TelemetryDruidTimeSeriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTimeSeriesRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity) *TelemetryDruidTimeSeriesRequest {
	this := TelemetryDruidTimeSeriesRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Intervals = intervals
	this.Granularity = granularity
	return &this
}

// NewTelemetryDruidTimeSeriesRequestWithDefaults instantiates a new TelemetryDruidTimeSeriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTimeSeriesRequestWithDefaults() *TelemetryDruidTimeSeriesRequest {
	this := TelemetryDruidTimeSeriesRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidTimeSeriesRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidTimeSeriesRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidTimeSeriesRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidTimeSeriesRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetDescending returns the Descending field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetDescending() bool {
	if o == nil || o.Descending == nil {
		var ret bool
		return ret
	}
	return *o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetDescendingOk() (*bool, bool) {
	if o == nil || o.Descending == nil {
		return nil, false
	}
	return o.Descending, true
}

// HasDescending returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasDescending() bool {
	if o != nil && o.Descending != nil {
		return true
	}

	return false
}

// SetDescending gets a reference to the given bool and assigns it to the Descending field.
func (o *TelemetryDruidTimeSeriesRequest) SetDescending(v bool) {
	o.Descending = &v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidTimeSeriesRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidTimeSeriesRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidTimeSeriesRequest) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidTimeSeriesRequest) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidTimeSeriesRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidTimeSeriesRequest) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = &v
}

// GetPostAggregations returns the PostAggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetPostAggregations() []TelemetryDruidPostAggregator {
	if o == nil || o.PostAggregations == nil {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return *o.PostAggregations
}

// GetPostAggregationsOk returns a tuple with the PostAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool) {
	if o == nil || o.PostAggregations == nil {
		return nil, false
	}
	return o.PostAggregations, true
}

// HasPostAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasPostAggregations() bool {
	if o != nil && o.PostAggregations != nil {
		return true
	}

	return false
}

// SetPostAggregations gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the PostAggregations field.
func (o *TelemetryDruidTimeSeriesRequest) SetPostAggregations(v []TelemetryDruidPostAggregator) {
	o.PostAggregations = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *TelemetryDruidTimeSeriesRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidTimeSeriesRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeSeriesRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidTimeSeriesRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidTimeSeriesRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidTimeSeriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.Descending != nil {
		toSerialize["descending"] = o.Descending
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.PostAggregations != nil {
		toSerialize["postAggregations"] = o.PostAggregations
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidTimeSeriesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidTimeSeriesRequest := _TelemetryDruidTimeSeriesRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidTimeSeriesRequest); err == nil {
		*o = TelemetryDruidTimeSeriesRequest(varTelemetryDruidTimeSeriesRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "descending")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "postAggregations")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTimeSeriesRequest struct {
	value *TelemetryDruidTimeSeriesRequest
	isSet bool
}

func (v NullableTelemetryDruidTimeSeriesRequest) Get() *TelemetryDruidTimeSeriesRequest {
	return v.value
}

func (v *NullableTelemetryDruidTimeSeriesRequest) Set(val *TelemetryDruidTimeSeriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTimeSeriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTimeSeriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTimeSeriesRequest(val *TelemetryDruidTimeSeriesRequest) *NullableTelemetryDruidTimeSeriesRequest {
	return &NullableTelemetryDruidTimeSeriesRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidTimeSeriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTimeSeriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
