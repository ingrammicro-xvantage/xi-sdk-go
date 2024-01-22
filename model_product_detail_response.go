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

// checks if the ProductDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailResponse{}

// ProductDetailResponse struct for ProductDetailResponse
type ProductDetailResponse struct {
	// Ingram Micro unique part number for the product.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// Vendor’s part number for the product.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// Boolean that indicates whether a product is authorized.
	ProductAuthorized *bool `json:"productAuthorized,omitempty"`
	// The description given for the product.
	Description *string `json:"description,omitempty"`
	// The detailed description given for the product.
	ProductDetailDescription *string `json:"productDetailDescription,omitempty"`
	// The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item.
	Upc *string `json:"upc,omitempty"`
	// The category of the product.
	ProductCategory *string `json:"productCategory,omitempty"`
	// The sub-category of the product.
	ProductSubcategory *string `json:"productSubcategory,omitempty"`
	// Vendor name for the order.
	VendorName *string `json:"vendorName,omitempty"`
	// Vendor number that identifies the product.
	VendorNumber *string `json:"vendorNumber,omitempty"`
	// Status code of the product.
	ProductStatusCode *string `json:"productStatusCode,omitempty"`
	// Indicates whether the product is directly shipped from the vendor’s warehouse or if the product ships from Ingram Micro’s warehouse. Class Codes are Ingram classifications on how skus are stocked A = Product that is stocked usually in all IM warehouses and replenished on a regular basis. B = Product that is stocked in limited IM warehouses and replenished on a regular basis C = Product that is stocked in fewer IM warehouses and replenished on a regular basis. D = Product that Ingram Micro has elected to discontinue. E = Product that will be phased out later, according to the vendor. You may not want to replenish this product, but instead sell down what is in stock. F = Product that we carry for a specific customer or supplier under a contractual agreement. N = New Sku. Classification before first receipt O = Discontinued product to be liquidated S= Order for Specialized Demand (Order to backorder) X= direct ship from Vendor V = product that vendor has elected to discontinue.
	ProductClass *string `json:"productClass,omitempty"`
	// Reseller / end-user’s part number for the product.
	CustomerPartNumber *string `json:"customerPartNumber,omitempty"`
	// Indicators of the Product
	Indicators []ProductDetailResponseIndicatorsInner `json:"indicators,omitempty"`
	// Cisco product related information.
	CiscoFields []ProductDetailResponseCiscoFieldsInner `json:"ciscoFields,omitempty"`
	// Technical specifications of the product.
	TechnicalSpecifications []ProductDetailResponseTechnicalSpecificationsInner `json:"technicalSpecifications,omitempty"`
	// Warranty information related to the product.
	WarrantyInformation []map[string]interface{} `json:"warrantyInformation,omitempty"`
	AdditionalInformation *ProductDetailResponseAdditionalInformation `json:"additionalInformation,omitempty"`
}

// NewProductDetailResponse instantiates a new ProductDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailResponse() *ProductDetailResponse {
	this := ProductDetailResponse{}
	return &this
}

// NewProductDetailResponseWithDefaults instantiates a new ProductDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailResponseWithDefaults() *ProductDetailResponse {
	this := ProductDetailResponse{}
	return &this
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *ProductDetailResponse) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *ProductDetailResponse) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetProductAuthorized returns the ProductAuthorized field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductAuthorized() bool {
	if o == nil || IsNil(o.ProductAuthorized) {
		var ret bool
		return ret
	}
	return *o.ProductAuthorized
}

// GetProductAuthorizedOk returns a tuple with the ProductAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductAuthorized) {
		return nil, false
	}
	return o.ProductAuthorized, true
}

// HasProductAuthorized returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductAuthorized() bool {
	if o != nil && !IsNil(o.ProductAuthorized) {
		return true
	}

	return false
}

// SetProductAuthorized gets a reference to the given bool and assigns it to the ProductAuthorized field.
func (o *ProductDetailResponse) SetProductAuthorized(v bool) {
	o.ProductAuthorized = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductDetailResponse) SetDescription(v string) {
	o.Description = &v
}

// GetProductDetailDescription returns the ProductDetailDescription field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductDetailDescription() string {
	if o == nil || IsNil(o.ProductDetailDescription) {
		var ret string
		return ret
	}
	return *o.ProductDetailDescription
}

// GetProductDetailDescriptionOk returns a tuple with the ProductDetailDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductDetailDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDetailDescription) {
		return nil, false
	}
	return o.ProductDetailDescription, true
}

// HasProductDetailDescription returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductDetailDescription() bool {
	if o != nil && !IsNil(o.ProductDetailDescription) {
		return true
	}

	return false
}

// SetProductDetailDescription gets a reference to the given string and assigns it to the ProductDetailDescription field.
func (o *ProductDetailResponse) SetProductDetailDescription(v string) {
	o.ProductDetailDescription = &v
}

// GetUpc returns the Upc field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetUpc() string {
	if o == nil || IsNil(o.Upc) {
		var ret string
		return ret
	}
	return *o.Upc
}

// GetUpcOk returns a tuple with the Upc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetUpcOk() (*string, bool) {
	if o == nil || IsNil(o.Upc) {
		return nil, false
	}
	return o.Upc, true
}

// HasUpc returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasUpc() bool {
	if o != nil && !IsNil(o.Upc) {
		return true
	}

	return false
}

// SetUpc gets a reference to the given string and assigns it to the Upc field.
func (o *ProductDetailResponse) SetUpc(v string) {
	o.Upc = &v
}

// GetProductCategory returns the ProductCategory field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductCategory() string {
	if o == nil || IsNil(o.ProductCategory) {
		var ret string
		return ret
	}
	return *o.ProductCategory
}

// GetProductCategoryOk returns a tuple with the ProductCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCategory) {
		return nil, false
	}
	return o.ProductCategory, true
}

// HasProductCategory returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductCategory() bool {
	if o != nil && !IsNil(o.ProductCategory) {
		return true
	}

	return false
}

// SetProductCategory gets a reference to the given string and assigns it to the ProductCategory field.
func (o *ProductDetailResponse) SetProductCategory(v string) {
	o.ProductCategory = &v
}

// GetProductSubcategory returns the ProductSubcategory field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductSubcategory() string {
	if o == nil || IsNil(o.ProductSubcategory) {
		var ret string
		return ret
	}
	return *o.ProductSubcategory
}

// GetProductSubcategoryOk returns a tuple with the ProductSubcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductSubcategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ProductSubcategory) {
		return nil, false
	}
	return o.ProductSubcategory, true
}

// HasProductSubcategory returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductSubcategory() bool {
	if o != nil && !IsNil(o.ProductSubcategory) {
		return true
	}

	return false
}

// SetProductSubcategory gets a reference to the given string and assigns it to the ProductSubcategory field.
func (o *ProductDetailResponse) SetProductSubcategory(v string) {
	o.ProductSubcategory = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *ProductDetailResponse) SetVendorName(v string) {
	o.VendorName = &v
}

// GetVendorNumber returns the VendorNumber field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetVendorNumber() string {
	if o == nil || IsNil(o.VendorNumber) {
		var ret string
		return ret
	}
	return *o.VendorNumber
}

// GetVendorNumberOk returns a tuple with the VendorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetVendorNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorNumber) {
		return nil, false
	}
	return o.VendorNumber, true
}

// HasVendorNumber returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasVendorNumber() bool {
	if o != nil && !IsNil(o.VendorNumber) {
		return true
	}

	return false
}

// SetVendorNumber gets a reference to the given string and assigns it to the VendorNumber field.
func (o *ProductDetailResponse) SetVendorNumber(v string) {
	o.VendorNumber = &v
}

// GetProductStatusCode returns the ProductStatusCode field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductStatusCode() string {
	if o == nil || IsNil(o.ProductStatusCode) {
		var ret string
		return ret
	}
	return *o.ProductStatusCode
}

// GetProductStatusCodeOk returns a tuple with the ProductStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductStatusCode) {
		return nil, false
	}
	return o.ProductStatusCode, true
}

// HasProductStatusCode returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductStatusCode() bool {
	if o != nil && !IsNil(o.ProductStatusCode) {
		return true
	}

	return false
}

// SetProductStatusCode gets a reference to the given string and assigns it to the ProductStatusCode field.
func (o *ProductDetailResponse) SetProductStatusCode(v string) {
	o.ProductStatusCode = &v
}

// GetProductClass returns the ProductClass field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetProductClass() string {
	if o == nil || IsNil(o.ProductClass) {
		var ret string
		return ret
	}
	return *o.ProductClass
}

// GetProductClassOk returns a tuple with the ProductClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetProductClassOk() (*string, bool) {
	if o == nil || IsNil(o.ProductClass) {
		return nil, false
	}
	return o.ProductClass, true
}

// HasProductClass returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasProductClass() bool {
	if o != nil && !IsNil(o.ProductClass) {
		return true
	}

	return false
}

// SetProductClass gets a reference to the given string and assigns it to the ProductClass field.
func (o *ProductDetailResponse) SetProductClass(v string) {
	o.ProductClass = &v
}

// GetCustomerPartNumber returns the CustomerPartNumber field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetCustomerPartNumber() string {
	if o == nil || IsNil(o.CustomerPartNumber) {
		var ret string
		return ret
	}
	return *o.CustomerPartNumber
}

// GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetCustomerPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerPartNumber) {
		return nil, false
	}
	return o.CustomerPartNumber, true
}

// HasCustomerPartNumber returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasCustomerPartNumber() bool {
	if o != nil && !IsNil(o.CustomerPartNumber) {
		return true
	}

	return false
}

// SetCustomerPartNumber gets a reference to the given string and assigns it to the CustomerPartNumber field.
func (o *ProductDetailResponse) SetCustomerPartNumber(v string) {
	o.CustomerPartNumber = &v
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetIndicators() []ProductDetailResponseIndicatorsInner {
	if o == nil || IsNil(o.Indicators) {
		var ret []ProductDetailResponseIndicatorsInner
		return ret
	}
	return o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetIndicatorsOk() ([]ProductDetailResponseIndicatorsInner, bool) {
	if o == nil || IsNil(o.Indicators) {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasIndicators() bool {
	if o != nil && !IsNil(o.Indicators) {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given []ProductDetailResponseIndicatorsInner and assigns it to the Indicators field.
func (o *ProductDetailResponse) SetIndicators(v []ProductDetailResponseIndicatorsInner) {
	o.Indicators = v
}

// GetCiscoFields returns the CiscoFields field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetCiscoFields() []ProductDetailResponseCiscoFieldsInner {
	if o == nil || IsNil(o.CiscoFields) {
		var ret []ProductDetailResponseCiscoFieldsInner
		return ret
	}
	return o.CiscoFields
}

// GetCiscoFieldsOk returns a tuple with the CiscoFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetCiscoFieldsOk() ([]ProductDetailResponseCiscoFieldsInner, bool) {
	if o == nil || IsNil(o.CiscoFields) {
		return nil, false
	}
	return o.CiscoFields, true
}

// HasCiscoFields returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasCiscoFields() bool {
	if o != nil && !IsNil(o.CiscoFields) {
		return true
	}

	return false
}

// SetCiscoFields gets a reference to the given []ProductDetailResponseCiscoFieldsInner and assigns it to the CiscoFields field.
func (o *ProductDetailResponse) SetCiscoFields(v []ProductDetailResponseCiscoFieldsInner) {
	o.CiscoFields = v
}

// GetTechnicalSpecifications returns the TechnicalSpecifications field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetTechnicalSpecifications() []ProductDetailResponseTechnicalSpecificationsInner {
	if o == nil || IsNil(o.TechnicalSpecifications) {
		var ret []ProductDetailResponseTechnicalSpecificationsInner
		return ret
	}
	return o.TechnicalSpecifications
}

// GetTechnicalSpecificationsOk returns a tuple with the TechnicalSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetTechnicalSpecificationsOk() ([]ProductDetailResponseTechnicalSpecificationsInner, bool) {
	if o == nil || IsNil(o.TechnicalSpecifications) {
		return nil, false
	}
	return o.TechnicalSpecifications, true
}

// HasTechnicalSpecifications returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasTechnicalSpecifications() bool {
	if o != nil && !IsNil(o.TechnicalSpecifications) {
		return true
	}

	return false
}

// SetTechnicalSpecifications gets a reference to the given []ProductDetailResponseTechnicalSpecificationsInner and assigns it to the TechnicalSpecifications field.
func (o *ProductDetailResponse) SetTechnicalSpecifications(v []ProductDetailResponseTechnicalSpecificationsInner) {
	o.TechnicalSpecifications = v
}

// GetWarrantyInformation returns the WarrantyInformation field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetWarrantyInformation() []map[string]interface{} {
	if o == nil || IsNil(o.WarrantyInformation) {
		var ret []map[string]interface{}
		return ret
	}
	return o.WarrantyInformation
}

// GetWarrantyInformationOk returns a tuple with the WarrantyInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetWarrantyInformationOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.WarrantyInformation) {
		return nil, false
	}
	return o.WarrantyInformation, true
}

// HasWarrantyInformation returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasWarrantyInformation() bool {
	if o != nil && !IsNil(o.WarrantyInformation) {
		return true
	}

	return false
}

// SetWarrantyInformation gets a reference to the given []map[string]interface{} and assigns it to the WarrantyInformation field.
func (o *ProductDetailResponse) SetWarrantyInformation(v []map[string]interface{}) {
	o.WarrantyInformation = v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *ProductDetailResponse) GetAdditionalInformation() ProductDetailResponseAdditionalInformation {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret ProductDetailResponseAdditionalInformation
		return ret
	}
	return *o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponse) GetAdditionalInformationOk() (*ProductDetailResponseAdditionalInformation, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *ProductDetailResponse) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given ProductDetailResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *ProductDetailResponse) SetAdditionalInformation(v ProductDetailResponseAdditionalInformation) {
	o.AdditionalInformation = &v
}

func (o ProductDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.ProductAuthorized) {
		toSerialize["productAuthorized"] = o.ProductAuthorized
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProductDetailDescription) {
		toSerialize["productDetailDescription"] = o.ProductDetailDescription
	}
	if !IsNil(o.Upc) {
		toSerialize["upc"] = o.Upc
	}
	if !IsNil(o.ProductCategory) {
		toSerialize["productCategory"] = o.ProductCategory
	}
	if !IsNil(o.ProductSubcategory) {
		toSerialize["productSubcategory"] = o.ProductSubcategory
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.VendorNumber) {
		toSerialize["vendorNumber"] = o.VendorNumber
	}
	if !IsNil(o.ProductStatusCode) {
		toSerialize["productStatusCode"] = o.ProductStatusCode
	}
	if !IsNil(o.ProductClass) {
		toSerialize["productClass"] = o.ProductClass
	}
	if !IsNil(o.CustomerPartNumber) {
		toSerialize["customerPartNumber"] = o.CustomerPartNumber
	}
	if !IsNil(o.Indicators) {
		toSerialize["indicators"] = o.Indicators
	}
	if !IsNil(o.CiscoFields) {
		toSerialize["ciscoFields"] = o.CiscoFields
	}
	if !IsNil(o.TechnicalSpecifications) {
		toSerialize["technicalSpecifications"] = o.TechnicalSpecifications
	}
	if !IsNil(o.WarrantyInformation) {
		toSerialize["warrantyInformation"] = o.WarrantyInformation
	}
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	return toSerialize, nil
}

type NullableProductDetailResponse struct {
	value *ProductDetailResponse
	isSet bool
}

func (v NullableProductDetailResponse) Get() *ProductDetailResponse {
	return v.value
}

func (v *NullableProductDetailResponse) Set(val *ProductDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailResponse(val *ProductDetailResponse) *NullableProductDetailResponse {
	return &NullableProductDetailResponse{value: val, isSet: true}
}

func (v NullableProductDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

