# OrderDetailB2B

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramOrderNumber** | Pointer to **string** | The IngramMicro sales order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The IngramMicro sales order date. | [optional] 
**OrderType** | Pointer to **string** | The IngramMicro sales order type. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end customer&#39;s order number for reference in their system. | [optional] 
**WebOrderId** | Pointer to **string** | The web order id of the order. | [optional] 
**VendorSalesOrderNumber** | Pointer to **string** | The vendor&#39;s order number for reference in their system | [optional] 
**IngramPurchaseOrderNumber** | Pointer to **string** | Ingram purchase order number. | [optional] 
**OrderStatus** | Pointer to **string** | The header-level status of the order. One of- Shipped, Canceled, Backordered, Processing, On Hold, Delivered. | [optional] 
**OrderTotal** | Pointer to **float64** | The total cost for the order, includes subtotal, freight charges, and tax. | [optional] 
**OrderSubTotal** | Pointer to **float64** | The sub total cost for the order, not including tax and freight. | [optional] 
**FreightCharges** | Pointer to **float64** | The freight charges for the order. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three digit ISO 4217 currency code for the order. | [optional] 
**TotalWeight** | Pointer to **float64** | Total order weight. unit -- North america - Pounds , other countries will be KG. | [optional] 
**TotalTax** | Pointer to **float64** | Total tax on the orders placed. | [optional] 
**TotalFees** | Pointer to **float64** | Total fees on the orders placed. | [optional] 
**PaymentTerms** | Pointer to **string** | The payment terms of the order. (Ex- Net 30 days). | [optional] 
**Notes** | Pointer to **string** | The header-level notes for the order. | [optional] 
**BillToInfo** | Pointer to [**OrderDetailB2BBillToInfo**](OrderDetailB2BBillToInfo.md) |  | [optional] 
**ShipToInfo** | Pointer to [**OrderDetailB2BShipToInfo**](OrderDetailB2BShipToInfo.md) |  | [optional] 
**EndUserInfo** | Pointer to [**OrderDetailB2BEndUserInfo**](OrderDetailB2BEndUserInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderDetailB2BLinesInner**](OrderDetailB2BLinesInner.md) |  | [optional] 
**MiscellaneousCharges** | Pointer to [**[]OrderDetailB2BMiscellaneousChargesInner**](OrderDetailB2BMiscellaneousChargesInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderDetailB2BAdditionalAttributesInner**](OrderDetailB2BAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewOrderDetailB2B

`func NewOrderDetailB2B() *OrderDetailB2B`

NewOrderDetailB2B instantiates a new OrderDetailB2B object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailB2BWithDefaults

`func NewOrderDetailB2BWithDefaults() *OrderDetailB2B`

NewOrderDetailB2BWithDefaults instantiates a new OrderDetailB2B object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramOrderNumber

`func (o *OrderDetailB2B) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderDetailB2B) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderDetailB2B) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderDetailB2B) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *OrderDetailB2B) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *OrderDetailB2B) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *OrderDetailB2B) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *OrderDetailB2B) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderDetailB2B) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderDetailB2B) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderDetailB2B) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderDetailB2B) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *OrderDetailB2B) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderDetailB2B) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderDetailB2B) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderDetailB2B) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *OrderDetailB2B) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *OrderDetailB2B) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *OrderDetailB2B) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *OrderDetailB2B) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetWebOrderId

`func (o *OrderDetailB2B) GetWebOrderId() string`

GetWebOrderId returns the WebOrderId field if non-nil, zero value otherwise.

### GetWebOrderIdOk

`func (o *OrderDetailB2B) GetWebOrderIdOk() (*string, bool)`

GetWebOrderIdOk returns a tuple with the WebOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrderId

`func (o *OrderDetailB2B) SetWebOrderId(v string)`

SetWebOrderId sets WebOrderId field to given value.

### HasWebOrderId

`func (o *OrderDetailB2B) HasWebOrderId() bool`

HasWebOrderId returns a boolean if a field has been set.

### GetVendorSalesOrderNumber

`func (o *OrderDetailB2B) GetVendorSalesOrderNumber() string`

GetVendorSalesOrderNumber returns the VendorSalesOrderNumber field if non-nil, zero value otherwise.

### GetVendorSalesOrderNumberOk

`func (o *OrderDetailB2B) GetVendorSalesOrderNumberOk() (*string, bool)`

GetVendorSalesOrderNumberOk returns a tuple with the VendorSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSalesOrderNumber

`func (o *OrderDetailB2B) SetVendorSalesOrderNumber(v string)`

SetVendorSalesOrderNumber sets VendorSalesOrderNumber field to given value.

### HasVendorSalesOrderNumber

`func (o *OrderDetailB2B) HasVendorSalesOrderNumber() bool`

HasVendorSalesOrderNumber returns a boolean if a field has been set.

### GetIngramPurchaseOrderNumber

`func (o *OrderDetailB2B) GetIngramPurchaseOrderNumber() string`

GetIngramPurchaseOrderNumber returns the IngramPurchaseOrderNumber field if non-nil, zero value otherwise.

### GetIngramPurchaseOrderNumberOk

`func (o *OrderDetailB2B) GetIngramPurchaseOrderNumberOk() (*string, bool)`

GetIngramPurchaseOrderNumberOk returns a tuple with the IngramPurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPurchaseOrderNumber

`func (o *OrderDetailB2B) SetIngramPurchaseOrderNumber(v string)`

SetIngramPurchaseOrderNumber sets IngramPurchaseOrderNumber field to given value.

### HasIngramPurchaseOrderNumber

`func (o *OrderDetailB2B) HasIngramPurchaseOrderNumber() bool`

HasIngramPurchaseOrderNumber returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderDetailB2B) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderDetailB2B) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderDetailB2B) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderDetailB2B) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderDetailB2B) GetOrderTotal() float64`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderDetailB2B) GetOrderTotalOk() (*float64, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderDetailB2B) SetOrderTotal(v float64)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderDetailB2B) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetOrderSubTotal

`func (o *OrderDetailB2B) GetOrderSubTotal() float64`

GetOrderSubTotal returns the OrderSubTotal field if non-nil, zero value otherwise.

### GetOrderSubTotalOk

`func (o *OrderDetailB2B) GetOrderSubTotalOk() (*float64, bool)`

GetOrderSubTotalOk returns a tuple with the OrderSubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubTotal

`func (o *OrderDetailB2B) SetOrderSubTotal(v float64)`

SetOrderSubTotal sets OrderSubTotal field to given value.

### HasOrderSubTotal

`func (o *OrderDetailB2B) HasOrderSubTotal() bool`

HasOrderSubTotal returns a boolean if a field has been set.

### GetFreightCharges

`func (o *OrderDetailB2B) GetFreightCharges() float64`

GetFreightCharges returns the FreightCharges field if non-nil, zero value otherwise.

### GetFreightChargesOk

`func (o *OrderDetailB2B) GetFreightChargesOk() (*float64, bool)`

GetFreightChargesOk returns a tuple with the FreightCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCharges

`func (o *OrderDetailB2B) SetFreightCharges(v float64)`

SetFreightCharges sets FreightCharges field to given value.

### HasFreightCharges

`func (o *OrderDetailB2B) HasFreightCharges() bool`

HasFreightCharges returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderDetailB2B) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderDetailB2B) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderDetailB2B) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderDetailB2B) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTotalWeight

`func (o *OrderDetailB2B) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *OrderDetailB2B) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *OrderDetailB2B) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *OrderDetailB2B) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderDetailB2B) GetTotalTax() float64`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderDetailB2B) GetTotalTaxOk() (*float64, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderDetailB2B) SetTotalTax(v float64)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderDetailB2B) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetTotalFees

`func (o *OrderDetailB2B) GetTotalFees() float64`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *OrderDetailB2B) GetTotalFeesOk() (*float64, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *OrderDetailB2B) SetTotalFees(v float64)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *OrderDetailB2B) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *OrderDetailB2B) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *OrderDetailB2B) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *OrderDetailB2B) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *OrderDetailB2B) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetNotes

`func (o *OrderDetailB2B) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderDetailB2B) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderDetailB2B) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderDetailB2B) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetBillToInfo

`func (o *OrderDetailB2B) GetBillToInfo() OrderDetailB2BBillToInfo`

GetBillToInfo returns the BillToInfo field if non-nil, zero value otherwise.

### GetBillToInfoOk

`func (o *OrderDetailB2B) GetBillToInfoOk() (*OrderDetailB2BBillToInfo, bool)`

GetBillToInfoOk returns a tuple with the BillToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToInfo

`func (o *OrderDetailB2B) SetBillToInfo(v OrderDetailB2BBillToInfo)`

SetBillToInfo sets BillToInfo field to given value.

### HasBillToInfo

`func (o *OrderDetailB2B) HasBillToInfo() bool`

HasBillToInfo returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderDetailB2B) GetShipToInfo() OrderDetailB2BShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderDetailB2B) GetShipToInfoOk() (*OrderDetailB2BShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderDetailB2B) SetShipToInfo(v OrderDetailB2BShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderDetailB2B) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *OrderDetailB2B) GetEndUserInfo() OrderDetailB2BEndUserInfo`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderDetailB2B) GetEndUserInfoOk() (*OrderDetailB2BEndUserInfo, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderDetailB2B) SetEndUserInfo(v OrderDetailB2BEndUserInfo)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderDetailB2B) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderDetailB2B) GetLines() []OrderDetailB2BLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderDetailB2B) GetLinesOk() (*[]OrderDetailB2BLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderDetailB2B) SetLines(v []OrderDetailB2BLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderDetailB2B) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *OrderDetailB2B) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *OrderDetailB2B) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetMiscellaneousCharges

`func (o *OrderDetailB2B) GetMiscellaneousCharges() []OrderDetailB2BMiscellaneousChargesInner`

GetMiscellaneousCharges returns the MiscellaneousCharges field if non-nil, zero value otherwise.

### GetMiscellaneousChargesOk

`func (o *OrderDetailB2B) GetMiscellaneousChargesOk() (*[]OrderDetailB2BMiscellaneousChargesInner, bool)`

GetMiscellaneousChargesOk returns a tuple with the MiscellaneousCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscellaneousCharges

`func (o *OrderDetailB2B) SetMiscellaneousCharges(v []OrderDetailB2BMiscellaneousChargesInner)`

SetMiscellaneousCharges sets MiscellaneousCharges field to given value.

### HasMiscellaneousCharges

`func (o *OrderDetailB2B) HasMiscellaneousCharges() bool`

HasMiscellaneousCharges returns a boolean if a field has been set.

### SetMiscellaneousChargesNil

`func (o *OrderDetailB2B) SetMiscellaneousChargesNil(b bool)`

 SetMiscellaneousChargesNil sets the value for MiscellaneousCharges to be an explicit nil

### UnsetMiscellaneousCharges
`func (o *OrderDetailB2B) UnsetMiscellaneousCharges()`

UnsetMiscellaneousCharges ensures that no value is present for MiscellaneousCharges, not even an explicit nil
### GetAdditionalAttributes

`func (o *OrderDetailB2B) GetAdditionalAttributes() []OrderDetailB2BAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderDetailB2B) GetAdditionalAttributesOk() (*[]OrderDetailB2BAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderDetailB2B) SetAdditionalAttributes(v []OrderDetailB2BAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderDetailB2B) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributesNil

`func (o *OrderDetailB2B) SetAdditionalAttributesNil(b bool)`

 SetAdditionalAttributesNil sets the value for AdditionalAttributes to be an explicit nil

### UnsetAdditionalAttributes
`func (o *OrderDetailB2B) UnsetAdditionalAttributes()`

UnsetAdditionalAttributes ensures that no value is present for AdditionalAttributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


