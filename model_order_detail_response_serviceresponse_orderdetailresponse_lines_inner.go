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

// checks if the OrderDetailResponseServiceresponseOrderdetailresponseLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseServiceresponseOrderdetailresponseLinesInner{}

// OrderDetailResponseServiceresponseOrderdetailresponseLinesInner struct for OrderDetailResponseServiceresponseOrderdetailresponseLinesInner
type OrderDetailResponseServiceresponseOrderdetailresponseLinesInner struct {
	// Impulse line number
	Linenumber *string `json:"linenumber,omitempty"`
	// Line of the Globel Sku / Customer Line Number
	Globallinenumber *string `json:"globallinenumber,omitempty"`
	// Order Suffix
	Ordersuffix *string `json:"ordersuffix,omitempty"`
	// Sales order number
	Erpordernumber *string `json:"erpordernumber,omitempty"`
	// Status of the line
	Linestatus *string `json:"linestatus,omitempty"`
	// Ingram part number
	Partnumber *string `json:"partnumber,omitempty"`
	// manufacture number of the product
	Manufacturerpartnumber *string `json:"manufacturerpartnumber,omitempty"`
	// name of the vendor
	Vendorname *string `json:"vendorname,omitempty"`
	// Ingram Micro assigned code for the vendor
	Vendorcode *string `json:"vendorcode,omitempty"`
	Partdescription1 *string `json:"partdescription1,omitempty"`
	Partdescription2 *string `json:"partdescription2,omitempty"`
	// weight of the product unit
	Unitweight *string `json:"unitweight,omitempty"`
	// Customer price of the unit
	Unitprice *float32 `json:"unitprice,omitempty"`
	// extended price of the order
	Extendedprice *float32 `json:"extendedprice,omitempty"`
	// tax amount for the order
	Taxamount *float32 `json:"taxamount,omitempty"`
	// no. of units requested
	Requestedquantity *string `json:"requestedquantity,omitempty"`
	// no. of units confirmed available
	Confirmedquantity *string `json:"confirmedquantity,omitempty"`
	// quantity of back order
	Backorderquantity *string `json:"backorderquantity,omitempty"`
	Serialnumberdetails []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner `json:"serialnumberdetails,omitempty"`
	Trackingnumber []string `json:"trackingnumber,omitempty"`
	Shipmentdetails []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner `json:"shipmentdetails,omitempty"`
	Productextendedspecs []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner `json:"productextendedspecs,omitempty"`
	// estimated date of back order
	Backorderetadate *string `json:"backorderetadate,omitempty"`
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInner instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInner() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseLinesInner{}
	return &this
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseServiceresponseOrderdetailresponseLinesInnerWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseLinesInner{}
	return &this
}

// GetLinenumber returns the Linenumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinenumber() string {
	if o == nil || IsNil(o.Linenumber) {
		var ret string
		return ret
	}
	return *o.Linenumber
}

// GetLinenumberOk returns a tuple with the Linenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Linenumber) {
		return nil, false
	}
	return o.Linenumber, true
}

// HasLinenumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasLinenumber() bool {
	if o != nil && !IsNil(o.Linenumber) {
		return true
	}

	return false
}

// SetLinenumber gets a reference to the given string and assigns it to the Linenumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetLinenumber(v string) {
	o.Linenumber = &v
}

// GetGloballinenumber returns the Globallinenumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetGloballinenumber() string {
	if o == nil || IsNil(o.Globallinenumber) {
		var ret string
		return ret
	}
	return *o.Globallinenumber
}

// GetGloballinenumberOk returns a tuple with the Globallinenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetGloballinenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Globallinenumber) {
		return nil, false
	}
	return o.Globallinenumber, true
}

// HasGloballinenumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasGloballinenumber() bool {
	if o != nil && !IsNil(o.Globallinenumber) {
		return true
	}

	return false
}

// SetGloballinenumber gets a reference to the given string and assigns it to the Globallinenumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetGloballinenumber(v string) {
	o.Globallinenumber = &v
}

// GetOrdersuffix returns the Ordersuffix field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetOrdersuffix() string {
	if o == nil || IsNil(o.Ordersuffix) {
		var ret string
		return ret
	}
	return *o.Ordersuffix
}

// GetOrdersuffixOk returns a tuple with the Ordersuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetOrdersuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Ordersuffix) {
		return nil, false
	}
	return o.Ordersuffix, true
}

// HasOrdersuffix returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasOrdersuffix() bool {
	if o != nil && !IsNil(o.Ordersuffix) {
		return true
	}

	return false
}

// SetOrdersuffix gets a reference to the given string and assigns it to the Ordersuffix field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetOrdersuffix(v string) {
	o.Ordersuffix = &v
}

// GetErpordernumber returns the Erpordernumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetErpordernumber() string {
	if o == nil || IsNil(o.Erpordernumber) {
		var ret string
		return ret
	}
	return *o.Erpordernumber
}

// GetErpordernumberOk returns a tuple with the Erpordernumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetErpordernumberOk() (*string, bool) {
	if o == nil || IsNil(o.Erpordernumber) {
		return nil, false
	}
	return o.Erpordernumber, true
}

// HasErpordernumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasErpordernumber() bool {
	if o != nil && !IsNil(o.Erpordernumber) {
		return true
	}

	return false
}

// SetErpordernumber gets a reference to the given string and assigns it to the Erpordernumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetErpordernumber(v string) {
	o.Erpordernumber = &v
}

// GetLinestatus returns the Linestatus field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinestatus() string {
	if o == nil || IsNil(o.Linestatus) {
		var ret string
		return ret
	}
	return *o.Linestatus
}

// GetLinestatusOk returns a tuple with the Linestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetLinestatusOk() (*string, bool) {
	if o == nil || IsNil(o.Linestatus) {
		return nil, false
	}
	return o.Linestatus, true
}

// HasLinestatus returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasLinestatus() bool {
	if o != nil && !IsNil(o.Linestatus) {
		return true
	}

	return false
}

// SetLinestatus gets a reference to the given string and assigns it to the Linestatus field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetLinestatus(v string) {
	o.Linestatus = &v
}

// GetPartnumber returns the Partnumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartnumber() string {
	if o == nil || IsNil(o.Partnumber) {
		var ret string
		return ret
	}
	return *o.Partnumber
}

// GetPartnumberOk returns a tuple with the Partnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Partnumber) {
		return nil, false
	}
	return o.Partnumber, true
}

// HasPartnumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartnumber() bool {
	if o != nil && !IsNil(o.Partnumber) {
		return true
	}

	return false
}

// SetPartnumber gets a reference to the given string and assigns it to the Partnumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartnumber(v string) {
	o.Partnumber = &v
}

// GetManufacturerpartnumber returns the Manufacturerpartnumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetManufacturerpartnumber() string {
	if o == nil || IsNil(o.Manufacturerpartnumber) {
		var ret string
		return ret
	}
	return *o.Manufacturerpartnumber
}

// GetManufacturerpartnumberOk returns a tuple with the Manufacturerpartnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetManufacturerpartnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturerpartnumber) {
		return nil, false
	}
	return o.Manufacturerpartnumber, true
}

// HasManufacturerpartnumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasManufacturerpartnumber() bool {
	if o != nil && !IsNil(o.Manufacturerpartnumber) {
		return true
	}

	return false
}

// SetManufacturerpartnumber gets a reference to the given string and assigns it to the Manufacturerpartnumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetManufacturerpartnumber(v string) {
	o.Manufacturerpartnumber = &v
}

// GetVendorname returns the Vendorname field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorname() string {
	if o == nil || IsNil(o.Vendorname) {
		var ret string
		return ret
	}
	return *o.Vendorname
}

// GetVendornameOk returns a tuple with the Vendorname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendornameOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorname) {
		return nil, false
	}
	return o.Vendorname, true
}

// HasVendorname returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasVendorname() bool {
	if o != nil && !IsNil(o.Vendorname) {
		return true
	}

	return false
}

// SetVendorname gets a reference to the given string and assigns it to the Vendorname field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetVendorname(v string) {
	o.Vendorname = &v
}

// GetVendorcode returns the Vendorcode field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorcode() string {
	if o == nil || IsNil(o.Vendorcode) {
		var ret string
		return ret
	}
	return *o.Vendorcode
}

// GetVendorcodeOk returns a tuple with the Vendorcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetVendorcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorcode) {
		return nil, false
	}
	return o.Vendorcode, true
}

// HasVendorcode returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasVendorcode() bool {
	if o != nil && !IsNil(o.Vendorcode) {
		return true
	}

	return false
}

// SetVendorcode gets a reference to the given string and assigns it to the Vendorcode field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetVendorcode(v string) {
	o.Vendorcode = &v
}

// GetPartdescription1 returns the Partdescription1 field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription1() string {
	if o == nil || IsNil(o.Partdescription1) {
		var ret string
		return ret
	}
	return *o.Partdescription1
}

// GetPartdescription1Ok returns a tuple with the Partdescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.Partdescription1) {
		return nil, false
	}
	return o.Partdescription1, true
}

// HasPartdescription1 returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartdescription1() bool {
	if o != nil && !IsNil(o.Partdescription1) {
		return true
	}

	return false
}

// SetPartdescription1 gets a reference to the given string and assigns it to the Partdescription1 field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartdescription1(v string) {
	o.Partdescription1 = &v
}

// GetPartdescription2 returns the Partdescription2 field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription2() string {
	if o == nil || IsNil(o.Partdescription2) {
		var ret string
		return ret
	}
	return *o.Partdescription2
}

// GetPartdescription2Ok returns a tuple with the Partdescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetPartdescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.Partdescription2) {
		return nil, false
	}
	return o.Partdescription2, true
}

// HasPartdescription2 returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasPartdescription2() bool {
	if o != nil && !IsNil(o.Partdescription2) {
		return true
	}

	return false
}

// SetPartdescription2 gets a reference to the given string and assigns it to the Partdescription2 field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetPartdescription2(v string) {
	o.Partdescription2 = &v
}

// GetUnitweight returns the Unitweight field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitweight() string {
	if o == nil || IsNil(o.Unitweight) {
		var ret string
		return ret
	}
	return *o.Unitweight
}

// GetUnitweightOk returns a tuple with the Unitweight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitweightOk() (*string, bool) {
	if o == nil || IsNil(o.Unitweight) {
		return nil, false
	}
	return o.Unitweight, true
}

// HasUnitweight returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasUnitweight() bool {
	if o != nil && !IsNil(o.Unitweight) {
		return true
	}

	return false
}

// SetUnitweight gets a reference to the given string and assigns it to the Unitweight field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetUnitweight(v string) {
	o.Unitweight = &v
}

// GetUnitprice returns the Unitprice field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitprice() float32 {
	if o == nil || IsNil(o.Unitprice) {
		var ret float32
		return ret
	}
	return *o.Unitprice
}

// GetUnitpriceOk returns a tuple with the Unitprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetUnitpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Unitprice) {
		return nil, false
	}
	return o.Unitprice, true
}

// HasUnitprice returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasUnitprice() bool {
	if o != nil && !IsNil(o.Unitprice) {
		return true
	}

	return false
}

// SetUnitprice gets a reference to the given float32 and assigns it to the Unitprice field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetUnitprice(v float32) {
	o.Unitprice = &v
}

// GetExtendedprice returns the Extendedprice field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetExtendedprice() float32 {
	if o == nil || IsNil(o.Extendedprice) {
		var ret float32
		return ret
	}
	return *o.Extendedprice
}

// GetExtendedpriceOk returns a tuple with the Extendedprice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetExtendedpriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Extendedprice) {
		return nil, false
	}
	return o.Extendedprice, true
}

// HasExtendedprice returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasExtendedprice() bool {
	if o != nil && !IsNil(o.Extendedprice) {
		return true
	}

	return false
}

// SetExtendedprice gets a reference to the given float32 and assigns it to the Extendedprice field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetExtendedprice(v float32) {
	o.Extendedprice = &v
}

// GetTaxamount returns the Taxamount field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTaxamount() float32 {
	if o == nil || IsNil(o.Taxamount) {
		var ret float32
		return ret
	}
	return *o.Taxamount
}

// GetTaxamountOk returns a tuple with the Taxamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTaxamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Taxamount) {
		return nil, false
	}
	return o.Taxamount, true
}

// HasTaxamount returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasTaxamount() bool {
	if o != nil && !IsNil(o.Taxamount) {
		return true
	}

	return false
}

// SetTaxamount gets a reference to the given float32 and assigns it to the Taxamount field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetTaxamount(v float32) {
	o.Taxamount = &v
}

// GetRequestedquantity returns the Requestedquantity field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetRequestedquantity() string {
	if o == nil || IsNil(o.Requestedquantity) {
		var ret string
		return ret
	}
	return *o.Requestedquantity
}

// GetRequestedquantityOk returns a tuple with the Requestedquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetRequestedquantityOk() (*string, bool) {
	if o == nil || IsNil(o.Requestedquantity) {
		return nil, false
	}
	return o.Requestedquantity, true
}

// HasRequestedquantity returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasRequestedquantity() bool {
	if o != nil && !IsNil(o.Requestedquantity) {
		return true
	}

	return false
}

// SetRequestedquantity gets a reference to the given string and assigns it to the Requestedquantity field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetRequestedquantity(v string) {
	o.Requestedquantity = &v
}

// GetConfirmedquantity returns the Confirmedquantity field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetConfirmedquantity() string {
	if o == nil || IsNil(o.Confirmedquantity) {
		var ret string
		return ret
	}
	return *o.Confirmedquantity
}

// GetConfirmedquantityOk returns a tuple with the Confirmedquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetConfirmedquantityOk() (*string, bool) {
	if o == nil || IsNil(o.Confirmedquantity) {
		return nil, false
	}
	return o.Confirmedquantity, true
}

// HasConfirmedquantity returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasConfirmedquantity() bool {
	if o != nil && !IsNil(o.Confirmedquantity) {
		return true
	}

	return false
}

// SetConfirmedquantity gets a reference to the given string and assigns it to the Confirmedquantity field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetConfirmedquantity(v string) {
	o.Confirmedquantity = &v
}

// GetBackorderquantity returns the Backorderquantity field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderquantity() string {
	if o == nil || IsNil(o.Backorderquantity) {
		var ret string
		return ret
	}
	return *o.Backorderquantity
}

// GetBackorderquantityOk returns a tuple with the Backorderquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderquantityOk() (*string, bool) {
	if o == nil || IsNil(o.Backorderquantity) {
		return nil, false
	}
	return o.Backorderquantity, true
}

// HasBackorderquantity returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasBackorderquantity() bool {
	if o != nil && !IsNil(o.Backorderquantity) {
		return true
	}

	return false
}

// SetBackorderquantity gets a reference to the given string and assigns it to the Backorderquantity field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetBackorderquantity(v string) {
	o.Backorderquantity = &v
}

// GetSerialnumberdetails returns the Serialnumberdetails field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetSerialnumberdetails() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner {
	if o == nil || IsNil(o.Serialnumberdetails) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner
		return ret
	}
	return o.Serialnumberdetails
}

// GetSerialnumberdetailsOk returns a tuple with the Serialnumberdetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetSerialnumberdetailsOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner, bool) {
	if o == nil || IsNil(o.Serialnumberdetails) {
		return nil, false
	}
	return o.Serialnumberdetails, true
}

// HasSerialnumberdetails returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasSerialnumberdetails() bool {
	if o != nil && !IsNil(o.Serialnumberdetails) {
		return true
	}

	return false
}

// SetSerialnumberdetails gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner and assigns it to the Serialnumberdetails field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetSerialnumberdetails(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerSerialnumberdetailsInner) {
	o.Serialnumberdetails = v
}

// GetTrackingnumber returns the Trackingnumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTrackingnumber() []string {
	if o == nil || IsNil(o.Trackingnumber) {
		var ret []string
		return ret
	}
	return o.Trackingnumber
}

// GetTrackingnumberOk returns a tuple with the Trackingnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetTrackingnumberOk() ([]string, bool) {
	if o == nil || IsNil(o.Trackingnumber) {
		return nil, false
	}
	return o.Trackingnumber, true
}

// HasTrackingnumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasTrackingnumber() bool {
	if o != nil && !IsNil(o.Trackingnumber) {
		return true
	}

	return false
}

// SetTrackingnumber gets a reference to the given []string and assigns it to the Trackingnumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetTrackingnumber(v []string) {
	o.Trackingnumber = v
}

// GetShipmentdetails returns the Shipmentdetails field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetShipmentdetails() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner {
	if o == nil || IsNil(o.Shipmentdetails) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner
		return ret
	}
	return o.Shipmentdetails
}

// GetShipmentdetailsOk returns a tuple with the Shipmentdetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetShipmentdetailsOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner, bool) {
	if o == nil || IsNil(o.Shipmentdetails) {
		return nil, false
	}
	return o.Shipmentdetails, true
}

// HasShipmentdetails returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasShipmentdetails() bool {
	if o != nil && !IsNil(o.Shipmentdetails) {
		return true
	}

	return false
}

// SetShipmentdetails gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner and assigns it to the Shipmentdetails field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetShipmentdetails(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInnerShipmentdetailsInner) {
	o.Shipmentdetails = v
}

// GetProductextendedspecs returns the Productextendedspecs field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetProductextendedspecs() []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner {
	if o == nil || IsNil(o.Productextendedspecs) {
		var ret []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner
		return ret
	}
	return o.Productextendedspecs
}

// GetProductextendedspecsOk returns a tuple with the Productextendedspecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetProductextendedspecsOk() ([]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner, bool) {
	if o == nil || IsNil(o.Productextendedspecs) {
		return nil, false
	}
	return o.Productextendedspecs, true
}

// HasProductextendedspecs returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasProductextendedspecs() bool {
	if o != nil && !IsNil(o.Productextendedspecs) {
		return true
	}

	return false
}

// SetProductextendedspecs gets a reference to the given []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner and assigns it to the Productextendedspecs field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetProductextendedspecs(v []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) {
	o.Productextendedspecs = v
}

// GetBackorderetadate returns the Backorderetadate field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderetadate() string {
	if o == nil || IsNil(o.Backorderetadate) {
		var ret string
		return ret
	}
	return *o.Backorderetadate
}

// GetBackorderetadateOk returns a tuple with the Backorderetadate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) GetBackorderetadateOk() (*string, bool) {
	if o == nil || IsNil(o.Backorderetadate) {
		return nil, false
	}
	return o.Backorderetadate, true
}

// HasBackorderetadate returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) HasBackorderetadate() bool {
	if o != nil && !IsNil(o.Backorderetadate) {
		return true
	}

	return false
}

// SetBackorderetadate gets a reference to the given string and assigns it to the Backorderetadate field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) SetBackorderetadate(v string) {
	o.Backorderetadate = &v
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Linenumber) {
		toSerialize["linenumber"] = o.Linenumber
	}
	if !IsNil(o.Globallinenumber) {
		toSerialize["globallinenumber"] = o.Globallinenumber
	}
	if !IsNil(o.Ordersuffix) {
		toSerialize["ordersuffix"] = o.Ordersuffix
	}
	if !IsNil(o.Erpordernumber) {
		toSerialize["erpordernumber"] = o.Erpordernumber
	}
	if !IsNil(o.Linestatus) {
		toSerialize["linestatus"] = o.Linestatus
	}
	if !IsNil(o.Partnumber) {
		toSerialize["partnumber"] = o.Partnumber
	}
	if !IsNil(o.Manufacturerpartnumber) {
		toSerialize["manufacturerpartnumber"] = o.Manufacturerpartnumber
	}
	if !IsNil(o.Vendorname) {
		toSerialize["vendorname"] = o.Vendorname
	}
	if !IsNil(o.Vendorcode) {
		toSerialize["vendorcode"] = o.Vendorcode
	}
	if !IsNil(o.Partdescription1) {
		toSerialize["partdescription1"] = o.Partdescription1
	}
	if !IsNil(o.Partdescription2) {
		toSerialize["partdescription2"] = o.Partdescription2
	}
	if !IsNil(o.Unitweight) {
		toSerialize["unitweight"] = o.Unitweight
	}
	if !IsNil(o.Unitprice) {
		toSerialize["unitprice"] = o.Unitprice
	}
	if !IsNil(o.Extendedprice) {
		toSerialize["extendedprice"] = o.Extendedprice
	}
	if !IsNil(o.Taxamount) {
		toSerialize["taxamount"] = o.Taxamount
	}
	if !IsNil(o.Requestedquantity) {
		toSerialize["requestedquantity"] = o.Requestedquantity
	}
	if !IsNil(o.Confirmedquantity) {
		toSerialize["confirmedquantity"] = o.Confirmedquantity
	}
	if !IsNil(o.Backorderquantity) {
		toSerialize["backorderquantity"] = o.Backorderquantity
	}
	if !IsNil(o.Serialnumberdetails) {
		toSerialize["serialnumberdetails"] = o.Serialnumberdetails
	}
	if !IsNil(o.Trackingnumber) {
		toSerialize["trackingnumber"] = o.Trackingnumber
	}
	if !IsNil(o.Shipmentdetails) {
		toSerialize["shipmentdetails"] = o.Shipmentdetails
	}
	if !IsNil(o.Productextendedspecs) {
		toSerialize["productextendedspecs"] = o.Productextendedspecs
	}
	if !IsNil(o.Backorderetadate) {
		toSerialize["backorderetadate"] = o.Backorderetadate
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner struct {
	value *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner
	isSet bool
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) Get() *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner {
	return v.value
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) Set(val *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner(val *OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) *NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner {
	return &NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

