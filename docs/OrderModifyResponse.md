# OrderModifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramOrderNumber** | Pointer to **string** | The IngramMicro order number. | [optional] 
**ChangeDescription** | Pointer to **string** | The description of the change. | [optional] 
**OrderModifiedDate** | Pointer to **string** | The date the order was modified. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end user/customer&#39;s order number for reference in their system. | [optional] 
**OrderTotal** | Pointer to **float32** | The total for the order. | [optional] 
**Notes** | Pointer to **string** | Order-level notes. | [optional] 
**OrderSubTotal** | Pointer to **float32** | The sub total for the order. | [optional] 
**FreightCharges** | Pointer to **float32** | The freight charges for the order. | [optional] 
**TotalTax** | Pointer to **float32** | The total tax for the order. | [optional] 
**OrderStatus** | Pointer to **string** | The status of the order. One of the following. Backordered, In Progress, Shipped, Delivered, Canceled, On Hold | [optional] 
**BillToAddressId** | Pointer to **string** | Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit. | [optional] 
**ShipToInfo** | Pointer to [**OrderModifyResponseShipToInfo**](OrderModifyResponseShipToInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderModifyResponseLinesInner**](OrderModifyResponseLinesInner.md) | The line-level details for the order. | [optional] 
**RejectedLineItems** | Pointer to [**[]OrderModifyResponseRejectedLineItemsInner**](OrderModifyResponseRejectedLineItemsInner.md) | Details for failed lines in the order. | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderModifyResponseLinesInnerAdditionalAttributesInner**](OrderModifyResponseLinesInnerAdditionalAttributesInner.md) | Header-level additional attributes. | [optional] 

## Methods

### NewOrderModifyResponse

`func NewOrderModifyResponse() *OrderModifyResponse`

NewOrderModifyResponse instantiates a new OrderModifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyResponseWithDefaults

`func NewOrderModifyResponseWithDefaults() *OrderModifyResponse`

NewOrderModifyResponseWithDefaults instantiates a new OrderModifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramOrderNumber

`func (o *OrderModifyResponse) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderModifyResponse) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderModifyResponse) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderModifyResponse) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetChangeDescription

`func (o *OrderModifyResponse) GetChangeDescription() string`

GetChangeDescription returns the ChangeDescription field if non-nil, zero value otherwise.

### GetChangeDescriptionOk

`func (o *OrderModifyResponse) GetChangeDescriptionOk() (*string, bool)`

GetChangeDescriptionOk returns a tuple with the ChangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDescription

`func (o *OrderModifyResponse) SetChangeDescription(v string)`

SetChangeDescription sets ChangeDescription field to given value.

### HasChangeDescription

`func (o *OrderModifyResponse) HasChangeDescription() bool`

HasChangeDescription returns a boolean if a field has been set.

### GetOrderModifiedDate

`func (o *OrderModifyResponse) GetOrderModifiedDate() string`

GetOrderModifiedDate returns the OrderModifiedDate field if non-nil, zero value otherwise.

### GetOrderModifiedDateOk

`func (o *OrderModifyResponse) GetOrderModifiedDateOk() (*string, bool)`

GetOrderModifiedDateOk returns a tuple with the OrderModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderModifiedDate

`func (o *OrderModifyResponse) SetOrderModifiedDate(v string)`

SetOrderModifiedDate sets OrderModifiedDate field to given value.

### HasOrderModifiedDate

`func (o *OrderModifyResponse) HasOrderModifiedDate() bool`

HasOrderModifiedDate returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *OrderModifyResponse) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderModifyResponse) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderModifyResponse) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderModifyResponse) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *OrderModifyResponse) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *OrderModifyResponse) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *OrderModifyResponse) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *OrderModifyResponse) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderModifyResponse) GetOrderTotal() float32`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderModifyResponse) GetOrderTotalOk() (*float32, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderModifyResponse) SetOrderTotal(v float32)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderModifyResponse) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetNotes

`func (o *OrderModifyResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderModifyResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderModifyResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderModifyResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrderSubTotal

`func (o *OrderModifyResponse) GetOrderSubTotal() float32`

GetOrderSubTotal returns the OrderSubTotal field if non-nil, zero value otherwise.

### GetOrderSubTotalOk

`func (o *OrderModifyResponse) GetOrderSubTotalOk() (*float32, bool)`

GetOrderSubTotalOk returns a tuple with the OrderSubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubTotal

`func (o *OrderModifyResponse) SetOrderSubTotal(v float32)`

SetOrderSubTotal sets OrderSubTotal field to given value.

### HasOrderSubTotal

`func (o *OrderModifyResponse) HasOrderSubTotal() bool`

HasOrderSubTotal returns a boolean if a field has been set.

### GetFreightCharges

`func (o *OrderModifyResponse) GetFreightCharges() float32`

GetFreightCharges returns the FreightCharges field if non-nil, zero value otherwise.

### GetFreightChargesOk

`func (o *OrderModifyResponse) GetFreightChargesOk() (*float32, bool)`

GetFreightChargesOk returns a tuple with the FreightCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCharges

`func (o *OrderModifyResponse) SetFreightCharges(v float32)`

SetFreightCharges sets FreightCharges field to given value.

### HasFreightCharges

`func (o *OrderModifyResponse) HasFreightCharges() bool`

HasFreightCharges returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderModifyResponse) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderModifyResponse) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderModifyResponse) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderModifyResponse) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderModifyResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderModifyResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderModifyResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderModifyResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetBillToAddressId

`func (o *OrderModifyResponse) GetBillToAddressId() string`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *OrderModifyResponse) GetBillToAddressIdOk() (*string, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *OrderModifyResponse) SetBillToAddressId(v string)`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *OrderModifyResponse) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderModifyResponse) GetShipToInfo() OrderModifyResponseShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderModifyResponse) GetShipToInfoOk() (*OrderModifyResponseShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderModifyResponse) SetShipToInfo(v OrderModifyResponseShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderModifyResponse) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderModifyResponse) GetLines() []OrderModifyResponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderModifyResponse) GetLinesOk() (*[]OrderModifyResponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderModifyResponse) SetLines(v []OrderModifyResponseLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderModifyResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetRejectedLineItems

`func (o *OrderModifyResponse) GetRejectedLineItems() []OrderModifyResponseRejectedLineItemsInner`

GetRejectedLineItems returns the RejectedLineItems field if non-nil, zero value otherwise.

### GetRejectedLineItemsOk

`func (o *OrderModifyResponse) GetRejectedLineItemsOk() (*[]OrderModifyResponseRejectedLineItemsInner, bool)`

GetRejectedLineItemsOk returns a tuple with the RejectedLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedLineItems

`func (o *OrderModifyResponse) SetRejectedLineItems(v []OrderModifyResponseRejectedLineItemsInner)`

SetRejectedLineItems sets RejectedLineItems field to given value.

### HasRejectedLineItems

`func (o *OrderModifyResponse) HasRejectedLineItems() bool`

HasRejectedLineItems returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderModifyResponse) GetAdditionalAttributes() []OrderModifyResponseLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderModifyResponse) GetAdditionalAttributesOk() (*[]OrderModifyResponseLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderModifyResponse) SetAdditionalAttributes(v []OrderModifyResponseLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderModifyResponse) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


