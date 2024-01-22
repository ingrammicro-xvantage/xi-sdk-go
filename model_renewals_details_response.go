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

// checks if the RenewalsDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsDetailsResponse{}

// RenewalsDetailsResponse struct for RenewalsDetailsResponse
type RenewalsDetailsResponse struct {
	// Unique Ingram renewal ID.
	RenewalId *string `json:"renewalId,omitempty"`
	// The IngramMicro sales order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The IngramMicro sales order date.
	IngramOrderDate *string `json:"ingramOrderDate,omitempty"`
	// Renewal expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// Ingram purchase order number.
	IngramPurchaseOrderNumber *string `json:"ingramPurchaseOrderNumber,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// The end customer's order number for reference in their system.
	EndCustomerOrderNumber *string `json:"endCustomerOrderNumber,omitempty"`
	// The value of the renewal.
	RenewalValue *float32 `json:"renewalValue,omitempty"`
	// The company name for the end user/customer.
	EndUser *string `json:"endUser,omitempty"`
	// The name of the vendor.
	Vendor *string `json:"vendor,omitempty"`
	// The status of the renewal.
	Status *string `json:"status,omitempty"`
	EndUserInfo []RenewalsDetailsResponseEndUserInfoInner `json:"endUserInfo,omitempty"`
	ReferenceNumber []RenewalsDetailsResponseReferenceNumberInner `json:"referenceNumber,omitempty"`
	Products []RenewalsDetailsResponseProductsInner `json:"products,omitempty"`
	AdditionalAttributes []RenewalsDetailsResponseAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewRenewalsDetailsResponse instantiates a new RenewalsDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsDetailsResponse() *RenewalsDetailsResponse {
	this := RenewalsDetailsResponse{}
	return &this
}

// NewRenewalsDetailsResponseWithDefaults instantiates a new RenewalsDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsDetailsResponseWithDefaults() *RenewalsDetailsResponse {
	this := RenewalsDetailsResponse{}
	return &this
}

// GetRenewalId returns the RenewalId field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetRenewalId() string {
	if o == nil || IsNil(o.RenewalId) {
		var ret string
		return ret
	}
	return *o.RenewalId
}

// GetRenewalIdOk returns a tuple with the RenewalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetRenewalIdOk() (*string, bool) {
	if o == nil || IsNil(o.RenewalId) {
		return nil, false
	}
	return o.RenewalId, true
}

// HasRenewalId returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasRenewalId() bool {
	if o != nil && !IsNil(o.RenewalId) {
		return true
	}

	return false
}

// SetRenewalId gets a reference to the given string and assigns it to the RenewalId field.
func (o *RenewalsDetailsResponse) SetRenewalId(v string) {
	o.RenewalId = &v
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *RenewalsDetailsResponse) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetIngramOrderDate returns the IngramOrderDate field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetIngramOrderDate() string {
	if o == nil || IsNil(o.IngramOrderDate) {
		var ret string
		return ret
	}
	return *o.IngramOrderDate
}

// GetIngramOrderDateOk returns a tuple with the IngramOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetIngramOrderDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderDate) {
		return nil, false
	}
	return o.IngramOrderDate, true
}

// HasIngramOrderDate returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasIngramOrderDate() bool {
	if o != nil && !IsNil(o.IngramOrderDate) {
		return true
	}

	return false
}

// SetIngramOrderDate gets a reference to the given string and assigns it to the IngramOrderDate field.
func (o *RenewalsDetailsResponse) SetIngramOrderDate(v string) {
	o.IngramOrderDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *RenewalsDetailsResponse) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetIngramPurchaseOrderNumber returns the IngramPurchaseOrderNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetIngramPurchaseOrderNumber() string {
	if o == nil || IsNil(o.IngramPurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramPurchaseOrderNumber
}

// GetIngramPurchaseOrderNumberOk returns a tuple with the IngramPurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetIngramPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPurchaseOrderNumber) {
		return nil, false
	}
	return o.IngramPurchaseOrderNumber, true
}

// HasIngramPurchaseOrderNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasIngramPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.IngramPurchaseOrderNumber) {
		return true
	}

	return false
}

// SetIngramPurchaseOrderNumber gets a reference to the given string and assigns it to the IngramPurchaseOrderNumber field.
func (o *RenewalsDetailsResponse) SetIngramPurchaseOrderNumber(v string) {
	o.IngramPurchaseOrderNumber = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *RenewalsDetailsResponse) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		return nil, false
	}
	return o.EndCustomerOrderNumber, true
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasEndCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.EndCustomerOrderNumber) {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given string and assigns it to the EndCustomerOrderNumber field.
func (o *RenewalsDetailsResponse) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber = &v
}

// GetRenewalValue returns the RenewalValue field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetRenewalValue() float32 {
	if o == nil || IsNil(o.RenewalValue) {
		var ret float32
		return ret
	}
	return *o.RenewalValue
}

// GetRenewalValueOk returns a tuple with the RenewalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetRenewalValueOk() (*float32, bool) {
	if o == nil || IsNil(o.RenewalValue) {
		return nil, false
	}
	return o.RenewalValue, true
}

// HasRenewalValue returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasRenewalValue() bool {
	if o != nil && !IsNil(o.RenewalValue) {
		return true
	}

	return false
}

// SetRenewalValue gets a reference to the given float32 and assigns it to the RenewalValue field.
func (o *RenewalsDetailsResponse) SetRenewalValue(v float32) {
	o.RenewalValue = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetEndUser() string {
	if o == nil || IsNil(o.EndUser) {
		var ret string
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given string and assigns it to the EndUser field.
func (o *RenewalsDetailsResponse) SetEndUser(v string) {
	o.EndUser = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *RenewalsDetailsResponse) SetVendor(v string) {
	o.Vendor = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RenewalsDetailsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetEndUserInfo() []RenewalsDetailsResponseEndUserInfoInner {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret []RenewalsDetailsResponseEndUserInfoInner
		return ret
	}
	return o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetEndUserInfoOk() ([]RenewalsDetailsResponseEndUserInfoInner, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given []RenewalsDetailsResponseEndUserInfoInner and assigns it to the EndUserInfo field.
func (o *RenewalsDetailsResponse) SetEndUserInfo(v []RenewalsDetailsResponseEndUserInfoInner) {
	o.EndUserInfo = v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetReferenceNumber() []RenewalsDetailsResponseReferenceNumberInner {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret []RenewalsDetailsResponseReferenceNumberInner
		return ret
	}
	return o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetReferenceNumberOk() ([]RenewalsDetailsResponseReferenceNumberInner, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given []RenewalsDetailsResponseReferenceNumberInner and assigns it to the ReferenceNumber field.
func (o *RenewalsDetailsResponse) SetReferenceNumber(v []RenewalsDetailsResponseReferenceNumberInner) {
	o.ReferenceNumber = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetProducts() []RenewalsDetailsResponseProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []RenewalsDetailsResponseProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetProductsOk() ([]RenewalsDetailsResponseProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []RenewalsDetailsResponseProductsInner and assigns it to the Products field.
func (o *RenewalsDetailsResponse) SetProducts(v []RenewalsDetailsResponseProductsInner) {
	o.Products = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *RenewalsDetailsResponse) GetAdditionalAttributes() []RenewalsDetailsResponseAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []RenewalsDetailsResponseAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponse) GetAdditionalAttributesOk() ([]RenewalsDetailsResponseAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *RenewalsDetailsResponse) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []RenewalsDetailsResponseAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *RenewalsDetailsResponse) SetAdditionalAttributes(v []RenewalsDetailsResponseAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o RenewalsDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RenewalId) {
		toSerialize["renewalId"] = o.RenewalId
	}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.IngramOrderDate) {
		toSerialize["ingramOrderDate"] = o.IngramOrderDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.IngramPurchaseOrderNumber) {
		toSerialize["ingramPurchaseOrderNumber"] = o.IngramPurchaseOrderNumber
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.EndCustomerOrderNumber) {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber
	}
	if !IsNil(o.RenewalValue) {
		toSerialize["renewalValue"] = o.RenewalValue
	}
	if !IsNil(o.EndUser) {
		toSerialize["endUser"] = o.EndUser
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if !IsNil(o.ReferenceNumber) {
		toSerialize["referenceNumber"] = o.ReferenceNumber
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableRenewalsDetailsResponse struct {
	value *RenewalsDetailsResponse
	isSet bool
}

func (v NullableRenewalsDetailsResponse) Get() *RenewalsDetailsResponse {
	return v.value
}

func (v *NullableRenewalsDetailsResponse) Set(val *RenewalsDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsDetailsResponse(val *RenewalsDetailsResponse) *NullableRenewalsDetailsResponse {
	return &NullableRenewalsDetailsResponse{value: val, isSet: true}
}

func (v NullableRenewalsDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

