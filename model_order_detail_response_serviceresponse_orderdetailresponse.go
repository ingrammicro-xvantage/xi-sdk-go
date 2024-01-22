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

// checks if the OrderDetailResponseServiceresponseOrderdetailresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseServiceresponseOrderdetailresponse{}

// OrderDetailResponseServiceresponseOrderdetailresponse struct for OrderDetailResponseServiceresponseOrderdetailresponse
type OrderDetailResponseServiceresponseOrderdetailresponse struct {
	Ordernumber *string `json:"ordernumber,omitempty"`
	// Order Type   B - BRANCH TRANSFER C - CASH ORDER D - DIRECT ORDER F - FUTURE ORDER P - SPECIAL ORDER Q - QUOTE ORDER S - STOCK ORDER M - MEMO ORDER
	Ordertype *string `json:"ordertype,omitempty"`
	// Customer PO number
	Customerordernumber *string `json:"customerordernumber,omitempty"`
	// End User PO number
	Enduserponumber *string `json:"enduserponumber,omitempty"`
	// Status of order within Ingram system S - SALES HOLD H - TAG HOLD I - INVOICED P - PENDING E - BILLING ERROR F - FORCE BILLING V - VOIDED T - TRANSFERRED D - HOLD SHIPMENT R - RELEASED O - IM ONLINE HOLD U - BILL FOR HISTORY ONLY W - ORDER NOT PRINTED A - DROP SHIP HOLD B - INTERNET CUST ORIG HOLD 1 - PICKED 2 - INSPECTED 3 - PACKED 4 - SHIPPED C - CREDIT HOLD 9 - CISCO 3A6 Q - RMA HOLD G - CREDIT HOLD N - CREDIT HOLD
	Orderstatus *string `json:"orderstatus,omitempty"`
	// Time stamp of the order placed
	Entrytimestamp *string `json:"entrytimestamp,omitempty"`
	// Description of the entry method 
	Entrymethoddescription *string `json:"entrymethoddescription,omitempty"`
	// Total order value
	Ordertotalvalue *float32 `json:"ordertotalvalue,omitempty"`
	// Subtotal order value
	Ordersubtotal *float32 `json:"ordersubtotal,omitempty"`
	// Freight charges
	Freightamount *string `json:"freightamount,omitempty"`
	// Country specific currency code
	Currencycode *string `json:"currencycode,omitempty"`
	// Total order weight. unit -- North america - Pounds , other countries will be KG
	Totalweight *string `json:"totalweight,omitempty"`
	// total tax on the orders placed
	Totaltax *string `json:"totaltax,omitempty"`
	Billtoaddress *OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress `json:"billtoaddress,omitempty"`
	Shiptoaddress *OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress `json:"shiptoaddress,omitempty"`
	Enduserinfo *OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo `json:"enduserinfo,omitempty"`
	Lines []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner `json:"lines,omitempty"`
	Commentlines []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner `json:"commentlines,omitempty"`
	Miscfeeline []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner `json:"miscfeeline,omitempty"`
	Extendedspecs []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner `json:"extendedspecs,omitempty"`
}

// NewOrderDetailResponseServiceresponseOrderdetailresponse instantiates a new OrderDetailResponseServiceresponseOrderdetailresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseServiceresponseOrderdetailresponse() *OrderDetailResponseServiceresponseOrderdetailresponse {
	this := OrderDetailResponseServiceresponseOrderdetailresponse{}
	return &this
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseServiceresponseOrderdetailresponseWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponse {
	this := OrderDetailResponseServiceresponseOrderdetailresponse{}
	return &this
}

// GetOrdernumber returns the Ordernumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdernumber() string {
	if o == nil || IsNil(o.Ordernumber) {
		var ret string
		return ret
	}
	return *o.Ordernumber
}

// GetOrdernumberOk returns a tuple with the Ordernumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdernumberOk() (*string, bool) {
	if o == nil || IsNil(o.Ordernumber) {
		return nil, false
	}
	return o.Ordernumber, true
}

// HasOrdernumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdernumber() bool {
	if o != nil && !IsNil(o.Ordernumber) {
		return true
	}

	return false
}

// SetOrdernumber gets a reference to the given string and assigns it to the Ordernumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdernumber(v string) {
	o.Ordernumber = &v
}

// GetOrdertype returns the Ordertype field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertype() string {
	if o == nil || IsNil(o.Ordertype) {
		var ret string
		return ret
	}
	return *o.Ordertype
}

// GetOrdertypeOk returns a tuple with the Ordertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ordertype) {
		return nil, false
	}
	return o.Ordertype, true
}

// HasOrdertype returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdertype() bool {
	if o != nil && !IsNil(o.Ordertype) {
		return true
	}

	return false
}

// SetOrdertype gets a reference to the given string and assigns it to the Ordertype field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdertype(v string) {
	o.Ordertype = &v
}

// GetCustomerordernumber returns the Customerordernumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCustomerordernumber() string {
	if o == nil || IsNil(o.Customerordernumber) {
		var ret string
		return ret
	}
	return *o.Customerordernumber
}

// GetCustomerordernumberOk returns a tuple with the Customerordernumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCustomerordernumberOk() (*string, bool) {
	if o == nil || IsNil(o.Customerordernumber) {
		return nil, false
	}
	return o.Customerordernumber, true
}

// HasCustomerordernumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCustomerordernumber() bool {
	if o != nil && !IsNil(o.Customerordernumber) {
		return true
	}

	return false
}

// SetCustomerordernumber gets a reference to the given string and assigns it to the Customerordernumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCustomerordernumber(v string) {
	o.Customerordernumber = &v
}

// GetEnduserponumber returns the Enduserponumber field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserponumber() string {
	if o == nil || IsNil(o.Enduserponumber) {
		var ret string
		return ret
	}
	return *o.Enduserponumber
}

// GetEnduserponumberOk returns a tuple with the Enduserponumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserponumberOk() (*string, bool) {
	if o == nil || IsNil(o.Enduserponumber) {
		return nil, false
	}
	return o.Enduserponumber, true
}

// HasEnduserponumber returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEnduserponumber() bool {
	if o != nil && !IsNil(o.Enduserponumber) {
		return true
	}

	return false
}

// SetEnduserponumber gets a reference to the given string and assigns it to the Enduserponumber field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEnduserponumber(v string) {
	o.Enduserponumber = &v
}

// GetOrderstatus returns the Orderstatus field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrderstatus() string {
	if o == nil || IsNil(o.Orderstatus) {
		var ret string
		return ret
	}
	return *o.Orderstatus
}

// GetOrderstatusOk returns a tuple with the Orderstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrderstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Orderstatus) {
		return nil, false
	}
	return o.Orderstatus, true
}

// HasOrderstatus returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrderstatus() bool {
	if o != nil && !IsNil(o.Orderstatus) {
		return true
	}

	return false
}

// SetOrderstatus gets a reference to the given string and assigns it to the Orderstatus field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrderstatus(v string) {
	o.Orderstatus = &v
}

// GetEntrytimestamp returns the Entrytimestamp field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrytimestamp() string {
	if o == nil || IsNil(o.Entrytimestamp) {
		var ret string
		return ret
	}
	return *o.Entrytimestamp
}

// GetEntrytimestampOk returns a tuple with the Entrytimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrytimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Entrytimestamp) {
		return nil, false
	}
	return o.Entrytimestamp, true
}

// HasEntrytimestamp returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEntrytimestamp() bool {
	if o != nil && !IsNil(o.Entrytimestamp) {
		return true
	}

	return false
}

// SetEntrytimestamp gets a reference to the given string and assigns it to the Entrytimestamp field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEntrytimestamp(v string) {
	o.Entrytimestamp = &v
}

// GetEntrymethoddescription returns the Entrymethoddescription field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrymethoddescription() string {
	if o == nil || IsNil(o.Entrymethoddescription) {
		var ret string
		return ret
	}
	return *o.Entrymethoddescription
}

// GetEntrymethoddescriptionOk returns a tuple with the Entrymethoddescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrymethoddescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Entrymethoddescription) {
		return nil, false
	}
	return o.Entrymethoddescription, true
}

// HasEntrymethoddescription returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEntrymethoddescription() bool {
	if o != nil && !IsNil(o.Entrymethoddescription) {
		return true
	}

	return false
}

// SetEntrymethoddescription gets a reference to the given string and assigns it to the Entrymethoddescription field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEntrymethoddescription(v string) {
	o.Entrymethoddescription = &v
}

// GetOrdertotalvalue returns the Ordertotalvalue field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertotalvalue() float32 {
	if o == nil || IsNil(o.Ordertotalvalue) {
		var ret float32
		return ret
	}
	return *o.Ordertotalvalue
}

// GetOrdertotalvalueOk returns a tuple with the Ordertotalvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertotalvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Ordertotalvalue) {
		return nil, false
	}
	return o.Ordertotalvalue, true
}

// HasOrdertotalvalue returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdertotalvalue() bool {
	if o != nil && !IsNil(o.Ordertotalvalue) {
		return true
	}

	return false
}

// SetOrdertotalvalue gets a reference to the given float32 and assigns it to the Ordertotalvalue field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdertotalvalue(v float32) {
	o.Ordertotalvalue = &v
}

// GetOrdersubtotal returns the Ordersubtotal field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdersubtotal() float32 {
	if o == nil || IsNil(o.Ordersubtotal) {
		var ret float32
		return ret
	}
	return *o.Ordersubtotal
}

// GetOrdersubtotalOk returns a tuple with the Ordersubtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdersubtotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Ordersubtotal) {
		return nil, false
	}
	return o.Ordersubtotal, true
}

// HasOrdersubtotal returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdersubtotal() bool {
	if o != nil && !IsNil(o.Ordersubtotal) {
		return true
	}

	return false
}

// SetOrdersubtotal gets a reference to the given float32 and assigns it to the Ordersubtotal field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdersubtotal(v float32) {
	o.Ordersubtotal = &v
}

// GetFreightamount returns the Freightamount field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetFreightamount() string {
	if o == nil || IsNil(o.Freightamount) {
		var ret string
		return ret
	}
	return *o.Freightamount
}

// GetFreightamountOk returns a tuple with the Freightamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetFreightamountOk() (*string, bool) {
	if o == nil || IsNil(o.Freightamount) {
		return nil, false
	}
	return o.Freightamount, true
}

// HasFreightamount returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasFreightamount() bool {
	if o != nil && !IsNil(o.Freightamount) {
		return true
	}

	return false
}

// SetFreightamount gets a reference to the given string and assigns it to the Freightamount field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetFreightamount(v string) {
	o.Freightamount = &v
}

// GetCurrencycode returns the Currencycode field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCurrencycode() string {
	if o == nil || IsNil(o.Currencycode) {
		var ret string
		return ret
	}
	return *o.Currencycode
}

// GetCurrencycodeOk returns a tuple with the Currencycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCurrencycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Currencycode) {
		return nil, false
	}
	return o.Currencycode, true
}

// HasCurrencycode returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCurrencycode() bool {
	if o != nil && !IsNil(o.Currencycode) {
		return true
	}

	return false
}

// SetCurrencycode gets a reference to the given string and assigns it to the Currencycode field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCurrencycode(v string) {
	o.Currencycode = &v
}

// GetTotalweight returns the Totalweight field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotalweight() string {
	if o == nil || IsNil(o.Totalweight) {
		var ret string
		return ret
	}
	return *o.Totalweight
}

// GetTotalweightOk returns a tuple with the Totalweight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotalweightOk() (*string, bool) {
	if o == nil || IsNil(o.Totalweight) {
		return nil, false
	}
	return o.Totalweight, true
}

// HasTotalweight returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasTotalweight() bool {
	if o != nil && !IsNil(o.Totalweight) {
		return true
	}

	return false
}

// SetTotalweight gets a reference to the given string and assigns it to the Totalweight field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetTotalweight(v string) {
	o.Totalweight = &v
}

// GetTotaltax returns the Totaltax field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotaltax() string {
	if o == nil || IsNil(o.Totaltax) {
		var ret string
		return ret
	}
	return *o.Totaltax
}

// GetTotaltaxOk returns a tuple with the Totaltax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotaltaxOk() (*string, bool) {
	if o == nil || IsNil(o.Totaltax) {
		return nil, false
	}
	return o.Totaltax, true
}

// HasTotaltax returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasTotaltax() bool {
	if o != nil && !IsNil(o.Totaltax) {
		return true
	}

	return false
}

// SetTotaltax gets a reference to the given string and assigns it to the Totaltax field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetTotaltax(v string) {
	o.Totaltax = &v
}

// GetBilltoaddress returns the Billtoaddress field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetBilltoaddress() OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress {
	if o == nil || IsNil(o.Billtoaddress) {
		var ret OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress
		return ret
	}
	return *o.Billtoaddress
}

// GetBilltoaddressOk returns a tuple with the Billtoaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetBilltoaddressOk() (*OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress, bool) {
	if o == nil || IsNil(o.Billtoaddress) {
		return nil, false
	}
	return o.Billtoaddress, true
}

// HasBilltoaddress returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasBilltoaddress() bool {
	if o != nil && !IsNil(o.Billtoaddress) {
		return true
	}

	return false
}

// SetBilltoaddress gets a reference to the given OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress and assigns it to the Billtoaddress field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetBilltoaddress(v OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress) {
	o.Billtoaddress = &v
}

// GetShiptoaddress returns the Shiptoaddress field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetShiptoaddress() OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress {
	if o == nil || IsNil(o.Shiptoaddress) {
		var ret OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress
		return ret
	}
	return *o.Shiptoaddress
}

// GetShiptoaddressOk returns a tuple with the Shiptoaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetShiptoaddressOk() (*OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress, bool) {
	if o == nil || IsNil(o.Shiptoaddress) {
		return nil, false
	}
	return o.Shiptoaddress, true
}

// HasShiptoaddress returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasShiptoaddress() bool {
	if o != nil && !IsNil(o.Shiptoaddress) {
		return true
	}

	return false
}

// SetShiptoaddress gets a reference to the given OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress and assigns it to the Shiptoaddress field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetShiptoaddress(v OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress) {
	o.Shiptoaddress = &v
}

// GetEnduserinfo returns the Enduserinfo field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserinfo() OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo {
	if o == nil || IsNil(o.Enduserinfo) {
		var ret OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo
		return ret
	}
	return *o.Enduserinfo
}

// GetEnduserinfoOk returns a tuple with the Enduserinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserinfoOk() (*OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo, bool) {
	if o == nil || IsNil(o.Enduserinfo) {
		return nil, false
	}
	return o.Enduserinfo, true
}

// HasEnduserinfo returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEnduserinfo() bool {
	if o != nil && !IsNil(o.Enduserinfo) {
		return true
	}

	return false
}

// SetEnduserinfo gets a reference to the given OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo and assigns it to the Enduserinfo field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEnduserinfo(v OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo) {
	o.Enduserinfo = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetLines() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetLinesOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner and assigns it to the Lines field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetLines(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner) {
	o.Lines = v
}

// GetCommentlines returns the Commentlines field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCommentlines() []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner {
	if o == nil || IsNil(o.Commentlines) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner
		return ret
	}
	return o.Commentlines
}

// GetCommentlinesOk returns a tuple with the Commentlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCommentlinesOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner, bool) {
	if o == nil || IsNil(o.Commentlines) {
		return nil, false
	}
	return o.Commentlines, true
}

// HasCommentlines returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCommentlines() bool {
	if o != nil && !IsNil(o.Commentlines) {
		return true
	}

	return false
}

// SetCommentlines gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner and assigns it to the Commentlines field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCommentlines(v []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) {
	o.Commentlines = v
}

// GetMiscfeeline returns the Miscfeeline field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetMiscfeeline() []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner {
	if o == nil || IsNil(o.Miscfeeline) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner
		return ret
	}
	return o.Miscfeeline
}

// GetMiscfeelineOk returns a tuple with the Miscfeeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetMiscfeelineOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner, bool) {
	if o == nil || IsNil(o.Miscfeeline) {
		return nil, false
	}
	return o.Miscfeeline, true
}

// HasMiscfeeline returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasMiscfeeline() bool {
	if o != nil && !IsNil(o.Miscfeeline) {
		return true
	}

	return false
}

// SetMiscfeeline gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner and assigns it to the Miscfeeline field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetMiscfeeline(v []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) {
	o.Miscfeeline = v
}

// GetExtendedspecs returns the Extendedspecs field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetExtendedspecs() []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner {
	if o == nil || IsNil(o.Extendedspecs) {
		var ret []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner
		return ret
	}
	return o.Extendedspecs
}

// GetExtendedspecsOk returns a tuple with the Extendedspecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetExtendedspecsOk() ([]OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner, bool) {
	if o == nil || IsNil(o.Extendedspecs) {
		return nil, false
	}
	return o.Extendedspecs, true
}

// HasExtendedspecs returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasExtendedspecs() bool {
	if o != nil && !IsNil(o.Extendedspecs) {
		return true
	}

	return false
}

// SetExtendedspecs gets a reference to the given []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner and assigns it to the Extendedspecs field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetExtendedspecs(v []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner) {
	o.Extendedspecs = v
}

func (o OrderDetailResponseServiceresponseOrderdetailresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseServiceresponseOrderdetailresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ordernumber) {
		toSerialize["ordernumber"] = o.Ordernumber
	}
	if !IsNil(o.Ordertype) {
		toSerialize["ordertype"] = o.Ordertype
	}
	if !IsNil(o.Customerordernumber) {
		toSerialize["customerordernumber"] = o.Customerordernumber
	}
	if !IsNil(o.Enduserponumber) {
		toSerialize["enduserponumber"] = o.Enduserponumber
	}
	if !IsNil(o.Orderstatus) {
		toSerialize["orderstatus"] = o.Orderstatus
	}
	if !IsNil(o.Entrytimestamp) {
		toSerialize["entrytimestamp"] = o.Entrytimestamp
	}
	if !IsNil(o.Entrymethoddescription) {
		toSerialize["entrymethoddescription"] = o.Entrymethoddescription
	}
	if !IsNil(o.Ordertotalvalue) {
		toSerialize["ordertotalvalue"] = o.Ordertotalvalue
	}
	if !IsNil(o.Ordersubtotal) {
		toSerialize["ordersubtotal"] = o.Ordersubtotal
	}
	if !IsNil(o.Freightamount) {
		toSerialize["freightamount"] = o.Freightamount
	}
	if !IsNil(o.Currencycode) {
		toSerialize["currencycode"] = o.Currencycode
	}
	if !IsNil(o.Totalweight) {
		toSerialize["totalweight"] = o.Totalweight
	}
	if !IsNil(o.Totaltax) {
		toSerialize["totaltax"] = o.Totaltax
	}
	if !IsNil(o.Billtoaddress) {
		toSerialize["billtoaddress"] = o.Billtoaddress
	}
	if !IsNil(o.Shiptoaddress) {
		toSerialize["shiptoaddress"] = o.Shiptoaddress
	}
	if !IsNil(o.Enduserinfo) {
		toSerialize["enduserinfo"] = o.Enduserinfo
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Commentlines) {
		toSerialize["commentlines"] = o.Commentlines
	}
	if !IsNil(o.Miscfeeline) {
		toSerialize["miscfeeline"] = o.Miscfeeline
	}
	if !IsNil(o.Extendedspecs) {
		toSerialize["extendedspecs"] = o.Extendedspecs
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseServiceresponseOrderdetailresponse struct {
	value *OrderDetailResponseServiceresponseOrderdetailresponse
	isSet bool
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponse) Get() *OrderDetailResponseServiceresponseOrderdetailresponse {
	return v.value
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponse) Set(val *OrderDetailResponseServiceresponseOrderdetailresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseServiceresponseOrderdetailresponse(val *OrderDetailResponseServiceresponseOrderdetailresponse) *NullableOrderDetailResponseServiceresponseOrderdetailresponse {
	return &NullableOrderDetailResponseServiceresponseOrderdetailresponse{value: val, isSet: true}
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

