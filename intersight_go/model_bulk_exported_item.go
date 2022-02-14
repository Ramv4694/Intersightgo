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
)

// BulkExportedItem A single managed object that is being exported.
type BulkExportedItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether tags must be exported for item MO.
	ExportTags *bool `json:"ExportTags,omitempty"`
	// Name of the file corresponding to item MO.
	FileName *string  `json:"FileName,omitempty"`
	Item     *MoMoRef `json:"Item,omitempty"`
	// MO item identity (the moref corresponding to item) expressed as a string.
	Name *string `json:"Name,omitempty"`
	// Name of the service that owns the item MO.
	ServiceName *string `json:"ServiceName,omitempty"`
	// Version of the service that owns the item MO.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Status of the item's export operation. * `` - The operation has not started. * `ValidationInProgress` - The validation operation is in progress. * `Valid` - The content to be imported is valid. * `InValid` - The content to be imported is not valid and the status message will have the reason. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out. * `OperationCancelled` - The operation has been cancelled. * `CancelInProgress` - The operation is being cancelled.
	Status *string `json:"Status,omitempty"`
	// Progress or error message for the MO's export operation.
	StatusMessage *string                       `json:"StatusMessage,omitempty"`
	Export        *BulkExportRelationship       `json:"Export,omitempty"`
	ParentItem    *BulkExportedItemRelationship `json:"ParentItem,omitempty"`
	// An array of relationships to bulkExportedItem resources.
	RelatedItems         []BulkExportedItemRelationship `json:"RelatedItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkExportedItem BulkExportedItem

// NewBulkExportedItem instantiates a new BulkExportedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkExportedItem(classId string, objectType string) *BulkExportedItem {
	this := BulkExportedItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	var exportTags bool = false
	this.ExportTags = &exportTags
	var status string = ""
	this.Status = &status
	return &this
}

// NewBulkExportedItemWithDefaults instantiates a new BulkExportedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkExportedItemWithDefaults() *BulkExportedItem {
	this := BulkExportedItem{}
	var classId string = "bulk.ExportedItem"
	this.ClassId = classId
	var objectType string = "bulk.ExportedItem"
	this.ObjectType = objectType
	var exportTags bool = false
	this.ExportTags = &exportTags
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkExportedItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkExportedItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkExportedItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkExportedItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExportTags returns the ExportTags field value if set, zero value otherwise.
func (o *BulkExportedItem) GetExportTags() bool {
	if o == nil || o.ExportTags == nil {
		var ret bool
		return ret
	}
	return *o.ExportTags
}

// GetExportTagsOk returns a tuple with the ExportTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetExportTagsOk() (*bool, bool) {
	if o == nil || o.ExportTags == nil {
		return nil, false
	}
	return o.ExportTags, true
}

// HasExportTags returns a boolean if a field has been set.
func (o *BulkExportedItem) HasExportTags() bool {
	if o != nil && o.ExportTags != nil {
		return true
	}

	return false
}

// SetExportTags gets a reference to the given bool and assigns it to the ExportTags field.
func (o *BulkExportedItem) SetExportTags(v bool) {
	o.ExportTags = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *BulkExportedItem) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *BulkExportedItem) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *BulkExportedItem) SetFileName(v string) {
	o.FileName = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *BulkExportedItem) GetItem() MoMoRef {
	if o == nil || o.Item == nil {
		var ret MoMoRef
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetItemOk() (*MoMoRef, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *BulkExportedItem) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given MoMoRef and assigns it to the Item field.
func (o *BulkExportedItem) SetItem(v MoMoRef) {
	o.Item = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BulkExportedItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BulkExportedItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BulkExportedItem) SetName(v string) {
	o.Name = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *BulkExportedItem) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *BulkExportedItem) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *BulkExportedItem) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *BulkExportedItem) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *BulkExportedItem) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *BulkExportedItem) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkExportedItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkExportedItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkExportedItem) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkExportedItem) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkExportedItem) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkExportedItem) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *BulkExportedItem) GetExport() BulkExportRelationship {
	if o == nil || o.Export == nil {
		var ret BulkExportRelationship
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetExportOk() (*BulkExportRelationship, bool) {
	if o == nil || o.Export == nil {
		return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *BulkExportedItem) HasExport() bool {
	if o != nil && o.Export != nil {
		return true
	}

	return false
}

// SetExport gets a reference to the given BulkExportRelationship and assigns it to the Export field.
func (o *BulkExportedItem) SetExport(v BulkExportRelationship) {
	o.Export = &v
}

// GetParentItem returns the ParentItem field value if set, zero value otherwise.
func (o *BulkExportedItem) GetParentItem() BulkExportedItemRelationship {
	if o == nil || o.ParentItem == nil {
		var ret BulkExportedItemRelationship
		return ret
	}
	return *o.ParentItem
}

// GetParentItemOk returns a tuple with the ParentItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExportedItem) GetParentItemOk() (*BulkExportedItemRelationship, bool) {
	if o == nil || o.ParentItem == nil {
		return nil, false
	}
	return o.ParentItem, true
}

// HasParentItem returns a boolean if a field has been set.
func (o *BulkExportedItem) HasParentItem() bool {
	if o != nil && o.ParentItem != nil {
		return true
	}

	return false
}

// SetParentItem gets a reference to the given BulkExportedItemRelationship and assigns it to the ParentItem field.
func (o *BulkExportedItem) SetParentItem(v BulkExportedItemRelationship) {
	o.ParentItem = &v
}

// GetRelatedItems returns the RelatedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExportedItem) GetRelatedItems() []BulkExportedItemRelationship {
	if o == nil {
		var ret []BulkExportedItemRelationship
		return ret
	}
	return o.RelatedItems
}

// GetRelatedItemsOk returns a tuple with the RelatedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExportedItem) GetRelatedItemsOk() (*[]BulkExportedItemRelationship, bool) {
	if o == nil || o.RelatedItems == nil {
		return nil, false
	}
	return &o.RelatedItems, true
}

// HasRelatedItems returns a boolean if a field has been set.
func (o *BulkExportedItem) HasRelatedItems() bool {
	if o != nil && o.RelatedItems != nil {
		return true
	}

	return false
}

// SetRelatedItems gets a reference to the given []BulkExportedItemRelationship and assigns it to the RelatedItems field.
func (o *BulkExportedItem) SetRelatedItems(v []BulkExportedItemRelationship) {
	o.RelatedItems = v
}

func (o BulkExportedItem) MarshalJSON() ([]byte, error) {
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
	if o.ExportTags != nil {
		toSerialize["ExportTags"] = o.ExportTags
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.Item != nil {
		toSerialize["Item"] = o.Item
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServiceName != nil {
		toSerialize["ServiceName"] = o.ServiceName
	}
	if o.ServiceVersion != nil {
		toSerialize["ServiceVersion"] = o.ServiceVersion
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Export != nil {
		toSerialize["Export"] = o.Export
	}
	if o.ParentItem != nil {
		toSerialize["ParentItem"] = o.ParentItem
	}
	if o.RelatedItems != nil {
		toSerialize["RelatedItems"] = o.RelatedItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkExportedItem) UnmarshalJSON(bytes []byte) (err error) {
	type BulkExportedItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies whether tags must be exported for item MO.
		ExportTags *bool `json:"ExportTags,omitempty"`
		// Name of the file corresponding to item MO.
		FileName *string  `json:"FileName,omitempty"`
		Item     *MoMoRef `json:"Item,omitempty"`
		// MO item identity (the moref corresponding to item) expressed as a string.
		Name *string `json:"Name,omitempty"`
		// Name of the service that owns the item MO.
		ServiceName *string `json:"ServiceName,omitempty"`
		// Version of the service that owns the item MO.
		ServiceVersion *string `json:"ServiceVersion,omitempty"`
		// Status of the item's export operation. * `` - The operation has not started. * `ValidationInProgress` - The validation operation is in progress. * `Valid` - The content to be imported is valid. * `InValid` - The content to be imported is not valid and the status message will have the reason. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out. * `OperationCancelled` - The operation has been cancelled. * `CancelInProgress` - The operation is being cancelled.
		Status *string `json:"Status,omitempty"`
		// Progress or error message for the MO's export operation.
		StatusMessage *string                       `json:"StatusMessage,omitempty"`
		Export        *BulkExportRelationship       `json:"Export,omitempty"`
		ParentItem    *BulkExportedItemRelationship `json:"ParentItem,omitempty"`
		// An array of relationships to bulkExportedItem resources.
		RelatedItems []BulkExportedItemRelationship `json:"RelatedItems,omitempty"`
	}

	varBulkExportedItemWithoutEmbeddedStruct := BulkExportedItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkExportedItemWithoutEmbeddedStruct)
	if err == nil {
		varBulkExportedItem := _BulkExportedItem{}
		varBulkExportedItem.ClassId = varBulkExportedItemWithoutEmbeddedStruct.ClassId
		varBulkExportedItem.ObjectType = varBulkExportedItemWithoutEmbeddedStruct.ObjectType
		varBulkExportedItem.ExportTags = varBulkExportedItemWithoutEmbeddedStruct.ExportTags
		varBulkExportedItem.FileName = varBulkExportedItemWithoutEmbeddedStruct.FileName
		varBulkExportedItem.Item = varBulkExportedItemWithoutEmbeddedStruct.Item
		varBulkExportedItem.Name = varBulkExportedItemWithoutEmbeddedStruct.Name
		varBulkExportedItem.ServiceName = varBulkExportedItemWithoutEmbeddedStruct.ServiceName
		varBulkExportedItem.ServiceVersion = varBulkExportedItemWithoutEmbeddedStruct.ServiceVersion
		varBulkExportedItem.Status = varBulkExportedItemWithoutEmbeddedStruct.Status
		varBulkExportedItem.StatusMessage = varBulkExportedItemWithoutEmbeddedStruct.StatusMessage
		varBulkExportedItem.Export = varBulkExportedItemWithoutEmbeddedStruct.Export
		varBulkExportedItem.ParentItem = varBulkExportedItemWithoutEmbeddedStruct.ParentItem
		varBulkExportedItem.RelatedItems = varBulkExportedItemWithoutEmbeddedStruct.RelatedItems
		*o = BulkExportedItem(varBulkExportedItem)
	} else {
		return err
	}

	varBulkExportedItem := _BulkExportedItem{}

	err = json.Unmarshal(bytes, &varBulkExportedItem)
	if err == nil {
		o.MoBaseMo = varBulkExportedItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExportTags")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "Item")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServiceName")
		delete(additionalProperties, "ServiceVersion")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Export")
		delete(additionalProperties, "ParentItem")
		delete(additionalProperties, "RelatedItems")

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

type NullableBulkExportedItem struct {
	value *BulkExportedItem
	isSet bool
}

func (v NullableBulkExportedItem) Get() *BulkExportedItem {
	return v.value
}

func (v *NullableBulkExportedItem) Set(val *BulkExportedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkExportedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkExportedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkExportedItem(val *BulkExportedItem) *NullableBulkExportedItem {
	return &NullableBulkExportedItem{value: val, isSet: true}
}

func (v NullableBulkExportedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkExportedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
