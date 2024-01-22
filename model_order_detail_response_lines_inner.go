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

// checks if the OrderDetailResponseLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseLinesInner{}

// OrderDetailResponseLinesInner struct for OrderDetailResponseLinesInner
type OrderDetailResponseLinesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// Unique Ingram Micro line number. Starts with 001.
	IngramOrderLineNumber *string `json:"ingramOrderLineNumber,omitempty"`
	// The vendor's sales order line number.
	VendorSalesOrderLineNumber *string `json:"vendorSalesOrderLineNumber,omitempty"`
	// The reseller's line item number for reference in their system.
	CustomerLinenumber *string `json:"customerLinenumber,omitempty"`
	// The status for the line item in the order. One of- Backordered, In Progress, Shipped, Delivered, Canceled, On Hold
	LineStatus *string `json:"lineStatus,omitempty"`
	// Unique IngramMicro part number.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The vendor's part number for the line item.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// The vendor's name for the part in their system.
	VendorName *string `json:"vendorName,omitempty"`
	// The vendor's description of the part in their system.
	PartDescription *string `json:"partDescription,omitempty"`
	// The unit weight of the line item.
	UnitWeight *float32 `json:"unitWeight,omitempty"`
	// The unit of measure for the line item.
	WeightUom *string `json:"weightUom,omitempty"`
	// The unit price of the line item.
	UnitPrice *int32 `json:"unitPrice,omitempty"`
	// The UPC code of a product.
	UpcCode *string `json:"upcCode,omitempty"`
	// Unit price X quantity for the line item.
	ExtendedPrice *float32 `json:"extendedPrice,omitempty"`
	// The tax amount for the line item.
	TaxAmount *float32 `json:"taxAmount,omitempty"`
	// The country-specific three character ISO 4217 currency code for the line item.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The quantity ordered of the line item.
	QuantityOrdered *int32 `json:"quantityOrdered,omitempty"`
	// The quantity confirmed for the line item.
	QuantityConfirmed *int32 `json:"quantityConfirmed,omitempty"`
	// The quantity backordered for the line item.
	QuantityBackOrdered *int32 `json:"quantityBackOrdered,omitempty"`
	// The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid numbers take precedence over header-level bid numbers.
	SpecialBidNumber *string `json:"specialBidNumber,omitempty"`
	// Reseller-requested delivery date. Delivery date is not guaranteed.
	RequestedDeliveryDate *string `json:"requestedDeliveryDate,omitempty"`
	// The delivery date promised by IngramMicro.
	PromisedDeliveryDate *string `json:"promisedDeliveryDate,omitempty"`
	// Line-level notes for the order.
	LineNotes *string `json:"lineNotes,omitempty"`
	ShipmentDetails []OrderDetailResponseLinesInnerShipmentDetailsInner `json:"shipmentDetails,omitempty"`
	AdditionalAttributes []OrderDetailResponseLinesInnerAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
	Links []OrderDetailResponseLinesInnerLinksInner `json:"links,omitempty"`
}

// NewOrderDetailResponseLinesInner instantiates a new OrderDetailResponseLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseLinesInner() *OrderDetailResponseLinesInner {
	this := OrderDetailResponseLinesInner{}
	return &this
}

// NewOrderDetailResponseLinesInnerWithDefaults instantiates a new OrderDetailResponseLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseLinesInnerWithDefaults() *OrderDetailResponseLinesInner {
	this := OrderDetailResponseLinesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderDetailResponseLinesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetIngramOrderLineNumber returns the IngramOrderLineNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetIngramOrderLineNumber() string {
	if o == nil || IsNil(o.IngramOrderLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderLineNumber
}

// GetIngramOrderLineNumberOk returns a tuple with the IngramOrderLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetIngramOrderLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderLineNumber) {
		return nil, false
	}
	return o.IngramOrderLineNumber, true
}

// HasIngramOrderLineNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasIngramOrderLineNumber() bool {
	if o != nil && !IsNil(o.IngramOrderLineNumber) {
		return true
	}

	return false
}

// SetIngramOrderLineNumber gets a reference to the given string and assigns it to the IngramOrderLineNumber field.
func (o *OrderDetailResponseLinesInner) SetIngramOrderLineNumber(v string) {
	o.IngramOrderLineNumber = &v
}

// GetVendorSalesOrderLineNumber returns the VendorSalesOrderLineNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetVendorSalesOrderLineNumber() string {
	if o == nil || IsNil(o.VendorSalesOrderLineNumber) {
		var ret string
		return ret
	}
	return *o.VendorSalesOrderLineNumber
}

// GetVendorSalesOrderLineNumberOk returns a tuple with the VendorSalesOrderLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetVendorSalesOrderLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorSalesOrderLineNumber) {
		return nil, false
	}
	return o.VendorSalesOrderLineNumber, true
}

// HasVendorSalesOrderLineNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasVendorSalesOrderLineNumber() bool {
	if o != nil && !IsNil(o.VendorSalesOrderLineNumber) {
		return true
	}

	return false
}

// SetVendorSalesOrderLineNumber gets a reference to the given string and assigns it to the VendorSalesOrderLineNumber field.
func (o *OrderDetailResponseLinesInner) SetVendorSalesOrderLineNumber(v string) {
	o.VendorSalesOrderLineNumber = &v
}

// GetCustomerLinenumber returns the CustomerLinenumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetCustomerLinenumber() string {
	if o == nil || IsNil(o.CustomerLinenumber) {
		var ret string
		return ret
	}
	return *o.CustomerLinenumber
}

// GetCustomerLinenumberOk returns a tuple with the CustomerLinenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetCustomerLinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLinenumber) {
		return nil, false
	}
	return o.CustomerLinenumber, true
}

// HasCustomerLinenumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasCustomerLinenumber() bool {
	if o != nil && !IsNil(o.CustomerLinenumber) {
		return true
	}

	return false
}

// SetCustomerLinenumber gets a reference to the given string and assigns it to the CustomerLinenumber field.
func (o *OrderDetailResponseLinesInner) SetCustomerLinenumber(v string) {
	o.CustomerLinenumber = &v
}

// GetLineStatus returns the LineStatus field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetLineStatus() string {
	if o == nil || IsNil(o.LineStatus) {
		var ret string
		return ret
	}
	return *o.LineStatus
}

// GetLineStatusOk returns a tuple with the LineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetLineStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LineStatus) {
		return nil, false
	}
	return o.LineStatus, true
}

// HasLineStatus returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasLineStatus() bool {
	if o != nil && !IsNil(o.LineStatus) {
		return true
	}

	return false
}

// SetLineStatus gets a reference to the given string and assigns it to the LineStatus field.
func (o *OrderDetailResponseLinesInner) SetLineStatus(v string) {
	o.LineStatus = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderDetailResponseLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *OrderDetailResponseLinesInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *OrderDetailResponseLinesInner) SetVendorName(v string) {
	o.VendorName = &v
}

// GetPartDescription returns the PartDescription field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetPartDescription() string {
	if o == nil || IsNil(o.PartDescription) {
		var ret string
		return ret
	}
	return *o.PartDescription
}

// GetPartDescriptionOk returns a tuple with the PartDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetPartDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PartDescription) {
		return nil, false
	}
	return o.PartDescription, true
}

// HasPartDescription returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasPartDescription() bool {
	if o != nil && !IsNil(o.PartDescription) {
		return true
	}

	return false
}

// SetPartDescription gets a reference to the given string and assigns it to the PartDescription field.
func (o *OrderDetailResponseLinesInner) SetPartDescription(v string) {
	o.PartDescription = &v
}

// GetUnitWeight returns the UnitWeight field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetUnitWeight() float32 {
	if o == nil || IsNil(o.UnitWeight) {
		var ret float32
		return ret
	}
	return *o.UnitWeight
}

// GetUnitWeightOk returns a tuple with the UnitWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetUnitWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitWeight) {
		return nil, false
	}
	return o.UnitWeight, true
}

// HasUnitWeight returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasUnitWeight() bool {
	if o != nil && !IsNil(o.UnitWeight) {
		return true
	}

	return false
}

// SetUnitWeight gets a reference to the given float32 and assigns it to the UnitWeight field.
func (o *OrderDetailResponseLinesInner) SetUnitWeight(v float32) {
	o.UnitWeight = &v
}

// GetWeightUom returns the WeightUom field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetWeightUom() string {
	if o == nil || IsNil(o.WeightUom) {
		var ret string
		return ret
	}
	return *o.WeightUom
}

// GetWeightUomOk returns a tuple with the WeightUom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetWeightUomOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUom) {
		return nil, false
	}
	return o.WeightUom, true
}

// HasWeightUom returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasWeightUom() bool {
	if o != nil && !IsNil(o.WeightUom) {
		return true
	}

	return false
}

// SetWeightUom gets a reference to the given string and assigns it to the WeightUom field.
func (o *OrderDetailResponseLinesInner) SetWeightUom(v string) {
	o.WeightUom = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetUnitPrice() int32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret int32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetUnitPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given int32 and assigns it to the UnitPrice field.
func (o *OrderDetailResponseLinesInner) SetUnitPrice(v int32) {
	o.UnitPrice = &v
}

// GetUpcCode returns the UpcCode field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetUpcCode() string {
	if o == nil || IsNil(o.UpcCode) {
		var ret string
		return ret
	}
	return *o.UpcCode
}

// GetUpcCodeOk returns a tuple with the UpcCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetUpcCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UpcCode) {
		return nil, false
	}
	return o.UpcCode, true
}

// HasUpcCode returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasUpcCode() bool {
	if o != nil && !IsNil(o.UpcCode) {
		return true
	}

	return false
}

// SetUpcCode gets a reference to the given string and assigns it to the UpcCode field.
func (o *OrderDetailResponseLinesInner) SetUpcCode(v string) {
	o.UpcCode = &v
}

// GetExtendedPrice returns the ExtendedPrice field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetExtendedPrice() float32 {
	if o == nil || IsNil(o.ExtendedPrice) {
		var ret float32
		return ret
	}
	return *o.ExtendedPrice
}

// GetExtendedPriceOk returns a tuple with the ExtendedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetExtendedPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ExtendedPrice) {
		return nil, false
	}
	return o.ExtendedPrice, true
}

// HasExtendedPrice returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasExtendedPrice() bool {
	if o != nil && !IsNil(o.ExtendedPrice) {
		return true
	}

	return false
}

// SetExtendedPrice gets a reference to the given float32 and assigns it to the ExtendedPrice field.
func (o *OrderDetailResponseLinesInner) SetExtendedPrice(v float32) {
	o.ExtendedPrice = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *OrderDetailResponseLinesInner) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderDetailResponseLinesInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *OrderDetailResponseLinesInner) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetQuantityConfirmed returns the QuantityConfirmed field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetQuantityConfirmed() int32 {
	if o == nil || IsNil(o.QuantityConfirmed) {
		var ret int32
		return ret
	}
	return *o.QuantityConfirmed
}

// GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetQuantityConfirmedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityConfirmed) {
		return nil, false
	}
	return o.QuantityConfirmed, true
}

// HasQuantityConfirmed returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasQuantityConfirmed() bool {
	if o != nil && !IsNil(o.QuantityConfirmed) {
		return true
	}

	return false
}

// SetQuantityConfirmed gets a reference to the given int32 and assigns it to the QuantityConfirmed field.
func (o *OrderDetailResponseLinesInner) SetQuantityConfirmed(v int32) {
	o.QuantityConfirmed = &v
}

// GetQuantityBackOrdered returns the QuantityBackOrdered field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetQuantityBackOrdered() int32 {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityBackOrdered
}

// GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetQuantityBackOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		return nil, false
	}
	return o.QuantityBackOrdered, true
}

// HasQuantityBackOrdered returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasQuantityBackOrdered() bool {
	if o != nil && !IsNil(o.QuantityBackOrdered) {
		return true
	}

	return false
}

// SetQuantityBackOrdered gets a reference to the given int32 and assigns it to the QuantityBackOrdered field.
func (o *OrderDetailResponseLinesInner) SetQuantityBackOrdered(v int32) {
	o.QuantityBackOrdered = &v
}

// GetSpecialBidNumber returns the SpecialBidNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetSpecialBidNumber() string {
	if o == nil || IsNil(o.SpecialBidNumber) {
		var ret string
		return ret
	}
	return *o.SpecialBidNumber
}

// GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetSpecialBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidNumber) {
		return nil, false
	}
	return o.SpecialBidNumber, true
}

// HasSpecialBidNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasSpecialBidNumber() bool {
	if o != nil && !IsNil(o.SpecialBidNumber) {
		return true
	}

	return false
}

// SetSpecialBidNumber gets a reference to the given string and assigns it to the SpecialBidNumber field.
func (o *OrderDetailResponseLinesInner) SetSpecialBidNumber(v string) {
	o.SpecialBidNumber = &v
}

// GetRequestedDeliveryDate returns the RequestedDeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetRequestedDeliveryDate() string {
	if o == nil || IsNil(o.RequestedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.RequestedDeliveryDate
}

// GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetRequestedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedDeliveryDate) {
		return nil, false
	}
	return o.RequestedDeliveryDate, true
}

// HasRequestedDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasRequestedDeliveryDate() bool {
	if o != nil && !IsNil(o.RequestedDeliveryDate) {
		return true
	}

	return false
}

// SetRequestedDeliveryDate gets a reference to the given string and assigns it to the RequestedDeliveryDate field.
func (o *OrderDetailResponseLinesInner) SetRequestedDeliveryDate(v string) {
	o.RequestedDeliveryDate = &v
}

// GetPromisedDeliveryDate returns the PromisedDeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetPromisedDeliveryDate() string {
	if o == nil || IsNil(o.PromisedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.PromisedDeliveryDate
}

// GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetPromisedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.PromisedDeliveryDate) {
		return nil, false
	}
	return o.PromisedDeliveryDate, true
}

// HasPromisedDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasPromisedDeliveryDate() bool {
	if o != nil && !IsNil(o.PromisedDeliveryDate) {
		return true
	}

	return false
}

// SetPromisedDeliveryDate gets a reference to the given string and assigns it to the PromisedDeliveryDate field.
func (o *OrderDetailResponseLinesInner) SetPromisedDeliveryDate(v string) {
	o.PromisedDeliveryDate = &v
}

// GetLineNotes returns the LineNotes field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetLineNotes() string {
	if o == nil || IsNil(o.LineNotes) {
		var ret string
		return ret
	}
	return *o.LineNotes
}

// GetLineNotesOk returns a tuple with the LineNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetLineNotesOk() (*string, bool) {
	if o == nil || IsNil(o.LineNotes) {
		return nil, false
	}
	return o.LineNotes, true
}

// HasLineNotes returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasLineNotes() bool {
	if o != nil && !IsNil(o.LineNotes) {
		return true
	}

	return false
}

// SetLineNotes gets a reference to the given string and assigns it to the LineNotes field.
func (o *OrderDetailResponseLinesInner) SetLineNotes(v string) {
	o.LineNotes = &v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetShipmentDetails() []OrderDetailResponseLinesInnerShipmentDetailsInner {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret []OrderDetailResponseLinesInnerShipmentDetailsInner
		return ret
	}
	return o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetShipmentDetailsOk() ([]OrderDetailResponseLinesInnerShipmentDetailsInner, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given []OrderDetailResponseLinesInnerShipmentDetailsInner and assigns it to the ShipmentDetails field.
func (o *OrderDetailResponseLinesInner) SetShipmentDetails(v []OrderDetailResponseLinesInnerShipmentDetailsInner) {
	o.ShipmentDetails = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetAdditionalAttributes() []OrderDetailResponseLinesInnerAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderDetailResponseLinesInnerAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetAdditionalAttributesOk() ([]OrderDetailResponseLinesInnerAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderDetailResponseLinesInnerAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderDetailResponseLinesInner) SetAdditionalAttributes(v []OrderDetailResponseLinesInnerAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInner) GetLinks() []OrderDetailResponseLinesInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []OrderDetailResponseLinesInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInner) GetLinksOk() ([]OrderDetailResponseLinesInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []OrderDetailResponseLinesInnerLinksInner and assigns it to the Links field.
func (o *OrderDetailResponseLinesInner) SetLinks(v []OrderDetailResponseLinesInnerLinksInner) {
	o.Links = v
}

func (o OrderDetailResponseLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubOrderNumber) {
		toSerialize["subOrderNumber"] = o.SubOrderNumber
	}
	if !IsNil(o.IngramOrderLineNumber) {
		toSerialize["ingramOrderLineNumber"] = o.IngramOrderLineNumber
	}
	if !IsNil(o.VendorSalesOrderLineNumber) {
		toSerialize["vendorSalesOrderLineNumber"] = o.VendorSalesOrderLineNumber
	}
	if !IsNil(o.CustomerLinenumber) {
		toSerialize["customerLinenumber"] = o.CustomerLinenumber
	}
	if !IsNil(o.LineStatus) {
		toSerialize["lineStatus"] = o.LineStatus
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.PartDescription) {
		toSerialize["partDescription"] = o.PartDescription
	}
	if !IsNil(o.UnitWeight) {
		toSerialize["unitWeight"] = o.UnitWeight
	}
	if !IsNil(o.WeightUom) {
		toSerialize["weightUom"] = o.WeightUom
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !IsNil(o.UpcCode) {
		toSerialize["upcCode"] = o.UpcCode
	}
	if !IsNil(o.ExtendedPrice) {
		toSerialize["extendedPrice"] = o.ExtendedPrice
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.QuantityOrdered) {
		toSerialize["quantityOrdered"] = o.QuantityOrdered
	}
	if !IsNil(o.QuantityConfirmed) {
		toSerialize["quantityConfirmed"] = o.QuantityConfirmed
	}
	if !IsNil(o.QuantityBackOrdered) {
		toSerialize["quantityBackOrdered"] = o.QuantityBackOrdered
	}
	if !IsNil(o.SpecialBidNumber) {
		toSerialize["specialBidNumber"] = o.SpecialBidNumber
	}
	if !IsNil(o.RequestedDeliveryDate) {
		toSerialize["requestedDeliveryDate"] = o.RequestedDeliveryDate
	}
	if !IsNil(o.PromisedDeliveryDate) {
		toSerialize["promisedDeliveryDate"] = o.PromisedDeliveryDate
	}
	if !IsNil(o.LineNotes) {
		toSerialize["lineNotes"] = o.LineNotes
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseLinesInner struct {
	value *OrderDetailResponseLinesInner
	isSet bool
}

func (v NullableOrderDetailResponseLinesInner) Get() *OrderDetailResponseLinesInner {
	return v.value
}

func (v *NullableOrderDetailResponseLinesInner) Set(val *OrderDetailResponseLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseLinesInner(val *OrderDetailResponseLinesInner) *NullableOrderDetailResponseLinesInner {
	return &NullableOrderDetailResponseLinesInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

