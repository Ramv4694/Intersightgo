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
)

// TelemetryDruidTopNRequest TopN queries return a sorted set of results for the values in a given dimension according to some criteria. Conceptually, they can be thought of as an approximate GroupByQuery over a single dimension with an Ordering spec. TopNs are much faster and resource efficient than GroupBys for this use case. These types of queries take a topN query object and return an array of JSON objects where each object represents a value asked for by the topN query. TopNs are approximate in that each data process will rank their top K results and only return those top K results to the Broker. K, by default in Druid, is max(1000, threshold). In practice, this means that if you ask for the top 1000 items ordered, the correctness of the first ~900 items will be 100%, and the ordering of the results after that is not guaranteed. TopNs can be made more accurate by increasing the threshold.
type TelemetryDruidTopNRequest struct {
	// null
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals   []string                  `json:"intervals"`
	Granularity TelemetryDruidGranularity `json:"granularity"`
	Filter      *TelemetryDruidFilter     `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations *[]TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires.
	PostAggregations *[]TelemetryDruidPostAggregator `json:"postAggregations,omitempty"`
	Dimension        TelemetryDruidDimensionSpec     `json:"dimension"`
	// An integer defining the N in the topN (i.e. how many results you want in the top list).
	Threshold            int32                        `json:"threshold"`
	Metric               TelemetryDruidTopNMetricSpec `json:"metric"`
	Context              *TelemetryDruidQueryContext  `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTopNRequest TelemetryDruidTopNRequest

// NewTelemetryDruidTopNRequest instantiates a new TelemetryDruidTopNRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTopNRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, dimension TelemetryDruidDimensionSpec, threshold int32, metric TelemetryDruidTopNMetricSpec) *TelemetryDruidTopNRequest {
	this := TelemetryDruidTopNRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	this.Intervals = intervals
	this.Granularity = granularity
	this.Dimension = dimension
	this.Threshold = threshold
	this.Metric = metric
	return &this
}

// NewTelemetryDruidTopNRequestWithDefaults instantiates a new TelemetryDruidTopNRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTopNRequestWithDefaults() *TelemetryDruidTopNRequest {
	this := TelemetryDruidTopNRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidTopNRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidTopNRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidTopNRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidTopNRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidTopNRequest) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidTopNRequest) SetIntervals(v []string) {
	o.Intervals = v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidTopNRequest) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidTopNRequest) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidTopNRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequest) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequest) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidTopNRequest) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = &v
}

// GetPostAggregations returns the PostAggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequest) GetPostAggregations() []TelemetryDruidPostAggregator {
	if o == nil || o.PostAggregations == nil {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return *o.PostAggregations
}

// GetPostAggregationsOk returns a tuple with the PostAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool) {
	if o == nil || o.PostAggregations == nil {
		return nil, false
	}
	return o.PostAggregations, true
}

// HasPostAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequest) HasPostAggregations() bool {
	if o != nil && o.PostAggregations != nil {
		return true
	}

	return false
}

// SetPostAggregations gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the PostAggregations field.
func (o *TelemetryDruidTopNRequest) SetPostAggregations(v []TelemetryDruidPostAggregator) {
	o.PostAggregations = &v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidTopNRequest) GetDimension() TelemetryDruidDimensionSpec {
	if o == nil {
		var ret TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetDimensionOk() (*TelemetryDruidDimensionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidTopNRequest) SetDimension(v TelemetryDruidDimensionSpec) {
	o.Dimension = v
}

// GetThreshold returns the Threshold field value
func (o *TelemetryDruidTopNRequest) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *TelemetryDruidTopNRequest) SetThreshold(v int32) {
	o.Threshold = v
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidTopNRequest) GetMetric() TelemetryDruidTopNMetricSpec {
	if o == nil {
		var ret TelemetryDruidTopNMetricSpec
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidTopNRequest) SetMetric(v TelemetryDruidTopNMetricSpec) {
	o.Metric = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidTopNRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidTopNRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}
	if true {
		toSerialize["dataSource"] = o.DataSource
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
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidTopNRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidTopNRequest := _TelemetryDruidTopNRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidTopNRequest); err == nil {
		*o = TelemetryDruidTopNRequest(varTelemetryDruidTopNRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "postAggregations")
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTopNRequest struct {
	value *TelemetryDruidTopNRequest
	isSet bool
}

func (v NullableTelemetryDruidTopNRequest) Get() *TelemetryDruidTopNRequest {
	return v.value
}

func (v *NullableTelemetryDruidTopNRequest) Set(val *TelemetryDruidTopNRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTopNRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTopNRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTopNRequest(val *TelemetryDruidTopNRequest) *NullableTelemetryDruidTopNRequest {
	return &NullableTelemetryDruidTopNRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidTopNRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTopNRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
