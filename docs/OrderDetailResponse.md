# OrderDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramOrderNumber** | Pointer to **string** | The IngramMicro sales order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The date and time in UTC format that the order was created. | [optional] 
**OrderType** | Pointer to **string** | The order type. One of B &#x3D; Branch Transfer, C &#x3D; COD, D &#x3D; Direct Ship, F &#x3D; Future Order, P &#x3D; Special Order, M &#x3D; Memo, Q &#x3D; Quote, S &#x3D; Sales Order. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end user/customer&#39;s order number for reference in their system. | [optional] 
**VendorSalesOrderNumber** | Pointer to **string** | The vendor&#39;s order number for reference in their system. | [optional] 
**OrderStatus** | Pointer to **string** | The header-level status of the order. One of- Shipped, Canceled, Backordered, Processing, On Hold, Delivered. | [optional] 
**OrderTotal** | Pointer to **float32** | The total cost for the order, includes subtotal, freight charges, and tax. | [optional] 
**OrderSubTotal** | Pointer to **float32** | The sub total cost for the order, not including tax and freight. | [optional] 
**FreightCharges** | Pointer to **float32** | The freight charges for the order. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three digit ISO 4217 currency code for the order. | [optional] 
**TotalWeight** | Pointer to **float32** | The total weight of the order. Pounds in North America, KG in all other countries. | [optional] 
**TotalTax** | Pointer to **float32** | The total tax for the order. | [optional] 
**PaymentTerms** | Pointer to **string** | The payment terms of the order. (Ex- Net 30 days). | [optional] 
**Notes** | Pointer to **string** | The header-level notes for the order. | [optional] 
**BillToInfo** | Pointer to [**OrderDetailResponseBillToInfo**](OrderDetailResponseBillToInfo.md) |  | [optional] 
**ShipToInfo** | Pointer to [**OrderDetailResponseShipToInfo**](OrderDetailResponseShipToInfo.md) |  | [optional] 
**EndUserInfo** | Pointer to [**OrderDetailResponseEndUserInfo**](OrderDetailResponseEndUserInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderDetailResponseLinesInner**](OrderDetailResponseLinesInner.md) |  | [optional] 
**MiscellaneousCharges** | Pointer to [**[]OrderDetailResponseMiscellaneousChargesInner**](OrderDetailResponseMiscellaneousChargesInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderDetailResponseLinesInnerAdditionalAttributesInner**](OrderDetailResponseLinesInnerAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewOrderDetailResponse

`func NewOrderDetailResponse() *OrderDetailResponse`

NewOrderDetailResponse instantiates a new OrderDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseWithDefaults

`func NewOrderDetailResponseWithDefaults() *OrderDetailResponse`

NewOrderDetailResponseWithDefaults instantiates a new OrderDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramOrderNumber

`func (o *OrderDetailResponse) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderDetailResponse) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderDetailResponse) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderDetailResponse) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *OrderDetailResponse) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *OrderDetailResponse) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *OrderDetailResponse) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *OrderDetailResponse) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderDetailResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderDetailResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderDetailResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderDetailResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *OrderDetailResponse) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderDetailResponse) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderDetailResponse) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderDetailResponse) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *OrderDetailResponse) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *OrderDetailResponse) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *OrderDetailResponse) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *OrderDetailResponse) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetVendorSalesOrderNumber

`func (o *OrderDetailResponse) GetVendorSalesOrderNumber() string`

GetVendorSalesOrderNumber returns the VendorSalesOrderNumber field if non-nil, zero value otherwise.

### GetVendorSalesOrderNumberOk

`func (o *OrderDetailResponse) GetVendorSalesOrderNumberOk() (*string, bool)`

GetVendorSalesOrderNumberOk returns a tuple with the VendorSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSalesOrderNumber

`func (o *OrderDetailResponse) SetVendorSalesOrderNumber(v string)`

SetVendorSalesOrderNumber sets VendorSalesOrderNumber field to given value.

### HasVendorSalesOrderNumber

`func (o *OrderDetailResponse) HasVendorSalesOrderNumber() bool`

HasVendorSalesOrderNumber returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderDetailResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderDetailResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderDetailResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderDetailResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderDetailResponse) GetOrderTotal() float32`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderDetailResponse) GetOrderTotalOk() (*float32, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderDetailResponse) SetOrderTotal(v float32)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderDetailResponse) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetOrderSubTotal

`func (o *OrderDetailResponse) GetOrderSubTotal() float32`

GetOrderSubTotal returns the OrderSubTotal field if non-nil, zero value otherwise.

### GetOrderSubTotalOk

`func (o *OrderDetailResponse) GetOrderSubTotalOk() (*float32, bool)`

GetOrderSubTotalOk returns a tuple with the OrderSubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubTotal

`func (o *OrderDetailResponse) SetOrderSubTotal(v float32)`

SetOrderSubTotal sets OrderSubTotal field to given value.

### HasOrderSubTotal

`func (o *OrderDetailResponse) HasOrderSubTotal() bool`

HasOrderSubTotal returns a boolean if a field has been set.

### GetFreightCharges

`func (o *OrderDetailResponse) GetFreightCharges() float32`

GetFreightCharges returns the FreightCharges field if non-nil, zero value otherwise.

### GetFreightChargesOk

`func (o *OrderDetailResponse) GetFreightChargesOk() (*float32, bool)`

GetFreightChargesOk returns a tuple with the FreightCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCharges

`func (o *OrderDetailResponse) SetFreightCharges(v float32)`

SetFreightCharges sets FreightCharges field to given value.

### HasFreightCharges

`func (o *OrderDetailResponse) HasFreightCharges() bool`

HasFreightCharges returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderDetailResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderDetailResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderDetailResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderDetailResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTotalWeight

`func (o *OrderDetailResponse) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *OrderDetailResponse) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *OrderDetailResponse) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *OrderDetailResponse) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderDetailResponse) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderDetailResponse) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderDetailResponse) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderDetailResponse) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *OrderDetailResponse) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *OrderDetailResponse) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *OrderDetailResponse) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *OrderDetailResponse) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetNotes

`func (o *OrderDetailResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderDetailResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderDetailResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderDetailResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetBillToInfo

`func (o *OrderDetailResponse) GetBillToInfo() OrderDetailResponseBillToInfo`

GetBillToInfo returns the BillToInfo field if non-nil, zero value otherwise.

### GetBillToInfoOk

`func (o *OrderDetailResponse) GetBillToInfoOk() (*OrderDetailResponseBillToInfo, bool)`

GetBillToInfoOk returns a tuple with the BillToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToInfo

`func (o *OrderDetailResponse) SetBillToInfo(v OrderDetailResponseBillToInfo)`

SetBillToInfo sets BillToInfo field to given value.

### HasBillToInfo

`func (o *OrderDetailResponse) HasBillToInfo() bool`

HasBillToInfo returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderDetailResponse) GetShipToInfo() OrderDetailResponseShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderDetailResponse) GetShipToInfoOk() (*OrderDetailResponseShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderDetailResponse) SetShipToInfo(v OrderDetailResponseShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderDetailResponse) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *OrderDetailResponse) GetEndUserInfo() OrderDetailResponseEndUserInfo`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderDetailResponse) GetEndUserInfoOk() (*OrderDetailResponseEndUserInfo, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderDetailResponse) SetEndUserInfo(v OrderDetailResponseEndUserInfo)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderDetailResponse) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderDetailResponse) GetLines() []OrderDetailResponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderDetailResponse) GetLinesOk() (*[]OrderDetailResponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderDetailResponse) SetLines(v []OrderDetailResponseLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderDetailResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMiscellaneousCharges

`func (o *OrderDetailResponse) GetMiscellaneousCharges() []OrderDetailResponseMiscellaneousChargesInner`

GetMiscellaneousCharges returns the MiscellaneousCharges field if non-nil, zero value otherwise.

### GetMiscellaneousChargesOk

`func (o *OrderDetailResponse) GetMiscellaneousChargesOk() (*[]OrderDetailResponseMiscellaneousChargesInner, bool)`

GetMiscellaneousChargesOk returns a tuple with the MiscellaneousCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscellaneousCharges

`func (o *OrderDetailResponse) SetMiscellaneousCharges(v []OrderDetailResponseMiscellaneousChargesInner)`

SetMiscellaneousCharges sets MiscellaneousCharges field to given value.

### HasMiscellaneousCharges

`func (o *OrderDetailResponse) HasMiscellaneousCharges() bool`

HasMiscellaneousCharges returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderDetailResponse) GetAdditionalAttributes() []OrderDetailResponseLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderDetailResponse) GetAdditionalAttributesOk() (*[]OrderDetailResponseLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderDetailResponse) SetAdditionalAttributes(v []OrderDetailResponseLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderDetailResponse) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


