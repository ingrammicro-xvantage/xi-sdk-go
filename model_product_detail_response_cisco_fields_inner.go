/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
)

// checks if the ProductDetailResponseCiscoFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailResponseCiscoFieldsInner{}

// ProductDetailResponseCiscoFieldsInner struct for ProductDetailResponseCiscoFieldsInner
type ProductDetailResponseCiscoFieldsInner struct {
	// Cisco product sub-group
	ProductSubGroup *string `json:"productSubGroup,omitempty"`
	// Cisco service program name
	ServiceProgramName *string `json:"serviceProgramName,omitempty"`
	// Cisco item catalog category
	ItemCatalogCategory *string `json:"itemCatalogCategory,omitempty"`
	// Cisco configuration indicator
	ConfigurationIndicator *string `json:"configurationIndicator,omitempty"`
	// Cisco internal business entity
	InternalBusinessEntity *string `json:"internalBusinessEntity,omitempty"`
	// Cisco item type
	ItemType *string `json:"itemType,omitempty"`
	// Cisco global list price
	GlobalListPrice *string `json:"globalListPrice,omitempty"`
}

// NewProductDetailResponseCiscoFieldsInner instantiates a new ProductDetailResponseCiscoFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailResponseCiscoFieldsInner() *ProductDetailResponseCiscoFieldsInner {
	this := ProductDetailResponseCiscoFieldsInner{}
	return &this
}

// NewProductDetailResponseCiscoFieldsInnerWithDefaults instantiates a new ProductDetailResponseCiscoFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailResponseCiscoFieldsInnerWithDefaults() *ProductDetailResponseCiscoFieldsInner {
	this := ProductDetailResponseCiscoFieldsInner{}
	return &this
}

// GetProductSubGroup returns the ProductSubGroup field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetProductSubGroup() string {
	if o == nil || IsNil(o.ProductSubGroup) {
		var ret string
		return ret
	}
	return *o.ProductSubGroup
}

// GetProductSubGroupOk returns a tuple with the ProductSubGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetProductSubGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ProductSubGroup) {
		return nil, false
	}
	return o.ProductSubGroup, true
}

// HasProductSubGroup returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasProductSubGroup() bool {
	if o != nil && !IsNil(o.ProductSubGroup) {
		return true
	}

	return false
}

// SetProductSubGroup gets a reference to the given string and assigns it to the ProductSubGroup field.
func (o *ProductDetailResponseCiscoFieldsInner) SetProductSubGroup(v string) {
	o.ProductSubGroup = &v
}

// GetServiceProgramName returns the ServiceProgramName field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetServiceProgramName() string {
	if o == nil || IsNil(o.ServiceProgramName) {
		var ret string
		return ret
	}
	return *o.ServiceProgramName
}

// GetServiceProgramNameOk returns a tuple with the ServiceProgramName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetServiceProgramNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceProgramName) {
		return nil, false
	}
	return o.ServiceProgramName, true
}

// HasServiceProgramName returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasServiceProgramName() bool {
	if o != nil && !IsNil(o.ServiceProgramName) {
		return true
	}

	return false
}

// SetServiceProgramName gets a reference to the given string and assigns it to the ServiceProgramName field.
func (o *ProductDetailResponseCiscoFieldsInner) SetServiceProgramName(v string) {
	o.ServiceProgramName = &v
}

// GetItemCatalogCategory returns the ItemCatalogCategory field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetItemCatalogCategory() string {
	if o == nil || IsNil(o.ItemCatalogCategory) {
		var ret string
		return ret
	}
	return *o.ItemCatalogCategory
}

// GetItemCatalogCategoryOk returns a tuple with the ItemCatalogCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetItemCatalogCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ItemCatalogCategory) {
		return nil, false
	}
	return o.ItemCatalogCategory, true
}

// HasItemCatalogCategory returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasItemCatalogCategory() bool {
	if o != nil && !IsNil(o.ItemCatalogCategory) {
		return true
	}

	return false
}

// SetItemCatalogCategory gets a reference to the given string and assigns it to the ItemCatalogCategory field.
func (o *ProductDetailResponseCiscoFieldsInner) SetItemCatalogCategory(v string) {
	o.ItemCatalogCategory = &v
}

// GetConfigurationIndicator returns the ConfigurationIndicator field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetConfigurationIndicator() string {
	if o == nil || IsNil(o.ConfigurationIndicator) {
		var ret string
		return ret
	}
	return *o.ConfigurationIndicator
}

// GetConfigurationIndicatorOk returns a tuple with the ConfigurationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetConfigurationIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationIndicator) {
		return nil, false
	}
	return o.ConfigurationIndicator, true
}

// HasConfigurationIndicator returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasConfigurationIndicator() bool {
	if o != nil && !IsNil(o.ConfigurationIndicator) {
		return true
	}

	return false
}

// SetConfigurationIndicator gets a reference to the given string and assigns it to the ConfigurationIndicator field.
func (o *ProductDetailResponseCiscoFieldsInner) SetConfigurationIndicator(v string) {
	o.ConfigurationIndicator = &v
}

// GetInternalBusinessEntity returns the InternalBusinessEntity field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetInternalBusinessEntity() string {
	if o == nil || IsNil(o.InternalBusinessEntity) {
		var ret string
		return ret
	}
	return *o.InternalBusinessEntity
}

// GetInternalBusinessEntityOk returns a tuple with the InternalBusinessEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetInternalBusinessEntityOk() (*string, bool) {
	if o == nil || IsNil(o.InternalBusinessEntity) {
		return nil, false
	}
	return o.InternalBusinessEntity, true
}

// HasInternalBusinessEntity returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasInternalBusinessEntity() bool {
	if o != nil && !IsNil(o.InternalBusinessEntity) {
		return true
	}

	return false
}

// SetInternalBusinessEntity gets a reference to the given string and assigns it to the InternalBusinessEntity field.
func (o *ProductDetailResponseCiscoFieldsInner) SetInternalBusinessEntity(v string) {
	o.InternalBusinessEntity = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *ProductDetailResponseCiscoFieldsInner) SetItemType(v string) {
	o.ItemType = &v
}

// GetGlobalListPrice returns the GlobalListPrice field value if set, zero value otherwise.
func (o *ProductDetailResponseCiscoFieldsInner) GetGlobalListPrice() string {
	if o == nil || IsNil(o.GlobalListPrice) {
		var ret string
		return ret
	}
	return *o.GlobalListPrice
}

// GetGlobalListPriceOk returns a tuple with the GlobalListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseCiscoFieldsInner) GetGlobalListPriceOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalListPrice) {
		return nil, false
	}
	return o.GlobalListPrice, true
}

// HasGlobalListPrice returns a boolean if a field has been set.
func (o *ProductDetailResponseCiscoFieldsInner) HasGlobalListPrice() bool {
	if o != nil && !IsNil(o.GlobalListPrice) {
		return true
	}

	return false
}

// SetGlobalListPrice gets a reference to the given string and assigns it to the GlobalListPrice field.
func (o *ProductDetailResponseCiscoFieldsInner) SetGlobalListPrice(v string) {
	o.GlobalListPrice = &v
}

func (o ProductDetailResponseCiscoFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDetailResponseCiscoFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductSubGroup) {
		toSerialize["productSubGroup"] = o.ProductSubGroup
	}
	if !IsNil(o.ServiceProgramName) {
		toSerialize["serviceProgramName"] = o.ServiceProgramName
	}
	if !IsNil(o.ItemCatalogCategory) {
		toSerialize["itemCatalogCategory"] = o.ItemCatalogCategory
	}
	if !IsNil(o.ConfigurationIndicator) {
		toSerialize["configurationIndicator"] = o.ConfigurationIndicator
	}
	if !IsNil(o.InternalBusinessEntity) {
		toSerialize["internalBusinessEntity"] = o.InternalBusinessEntity
	}
	if !IsNil(o.ItemType) {
		toSerialize["itemType"] = o.ItemType
	}
	if !IsNil(o.GlobalListPrice) {
		toSerialize["globalListPrice"] = o.GlobalListPrice
	}
	return toSerialize, nil
}

type NullableProductDetailResponseCiscoFieldsInner struct {
	value *ProductDetailResponseCiscoFieldsInner
	isSet bool
}

func (v NullableProductDetailResponseCiscoFieldsInner) Get() *ProductDetailResponseCiscoFieldsInner {
	return v.value
}

func (v *NullableProductDetailResponseCiscoFieldsInner) Set(val *ProductDetailResponseCiscoFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailResponseCiscoFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailResponseCiscoFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailResponseCiscoFieldsInner(val *ProductDetailResponseCiscoFieldsInner) *NullableProductDetailResponseCiscoFieldsInner {
	return &NullableProductDetailResponseCiscoFieldsInner{value: val, isSet: true}
}

func (v NullableProductDetailResponseCiscoFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailResponseCiscoFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


