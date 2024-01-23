# OrderCreateResponseOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfLinesWithSuccess** | Pointer to **int32** | The number of lines in the order that were successful. | [optional] 
**NumberOfLinesWithError** | Pointer to **int32** | The number of lines in the order that have errors. | [optional] 
**NumberOfLinesWithWarning** | Pointer to **int32** | The number of lines in the order that have a warning. | [optional] 
**IngramOrderNumber** | Pointer to **string** | The Ingram Micro order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The date in UTC format that the order was created in Ingram Micro&#39;s system. | [optional] 
**Notes** | Pointer to **string** | Order-level notes. | [optional] 
**OrderType** | Pointer to **string** | The order typer. One of: S&#x3D;Stocked PO D&#x3D;Direct Ship PO | [optional] 
**OrderTotal** | Pointer to **float32** | The total price for the order. | [optional] 
**FreightCharges** | Pointer to **float32** | The total freight charges for the order. | [optional] 
**TotalTax** | Pointer to **float32** | The total tax for the order. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three character ISO 4217 currency code used for the order. | [optional] 
**Lines** | Pointer to [**[]OrderCreateResponseOrdersInnerLinesInner**](OrderCreateResponseOrdersInnerLinesInner.md) | The line-level details for the order. | [optional] 
**MiscellaneousCharges** | Pointer to [**[]OrderCreateResponseOrdersInnerMiscellaneousChargesInner**](OrderCreateResponseOrdersInnerMiscellaneousChargesInner.md) |  | [optional] 
**Links** | Pointer to [**[]OrderCreateResponseOrdersInnerLinksInner**](OrderCreateResponseOrdersInnerLinksInner.md) | Link to Order Details for the order(s). | [optional] 
**RejectedLineItems** | Pointer to [**[]OrderCreateResponseOrdersInnerRejectedLineItemsInner**](OrderCreateResponseOrdersInnerRejectedLineItemsInner.md) | A list of rejected line items. | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateResponseOrdersInnerAdditionalAttributesInner**](OrderCreateResponseOrdersInnerAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewOrderCreateResponseOrdersInner

`func NewOrderCreateResponseOrdersInner() *OrderCreateResponseOrdersInner`

NewOrderCreateResponseOrdersInner instantiates a new OrderCreateResponseOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerWithDefaults() *OrderCreateResponseOrdersInner`

NewOrderCreateResponseOrdersInnerWithDefaults instantiates a new OrderCreateResponseOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfLinesWithSuccess

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithSuccess() int32`

GetNumberOfLinesWithSuccess returns the NumberOfLinesWithSuccess field if non-nil, zero value otherwise.

### GetNumberOfLinesWithSuccessOk

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithSuccessOk() (*int32, bool)`

GetNumberOfLinesWithSuccessOk returns a tuple with the NumberOfLinesWithSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithSuccess

`func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithSuccess(v int32)`

SetNumberOfLinesWithSuccess sets NumberOfLinesWithSuccess field to given value.

### HasNumberOfLinesWithSuccess

`func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithSuccess() bool`

HasNumberOfLinesWithSuccess returns a boolean if a field has been set.

### GetNumberOfLinesWithError

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithError() int32`

GetNumberOfLinesWithError returns the NumberOfLinesWithError field if non-nil, zero value otherwise.

### GetNumberOfLinesWithErrorOk

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithErrorOk() (*int32, bool)`

GetNumberOfLinesWithErrorOk returns a tuple with the NumberOfLinesWithError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithError

`func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithError(v int32)`

SetNumberOfLinesWithError sets NumberOfLinesWithError field to given value.

### HasNumberOfLinesWithError

`func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithError() bool`

HasNumberOfLinesWithError returns a boolean if a field has been set.

### GetNumberOfLinesWithWarning

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithWarning() int32`

GetNumberOfLinesWithWarning returns the NumberOfLinesWithWarning field if non-nil, zero value otherwise.

### GetNumberOfLinesWithWarningOk

`func (o *OrderCreateResponseOrdersInner) GetNumberOfLinesWithWarningOk() (*int32, bool)`

GetNumberOfLinesWithWarningOk returns a tuple with the NumberOfLinesWithWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinesWithWarning

`func (o *OrderCreateResponseOrdersInner) SetNumberOfLinesWithWarning(v int32)`

SetNumberOfLinesWithWarning sets NumberOfLinesWithWarning field to given value.

### HasNumberOfLinesWithWarning

`func (o *OrderCreateResponseOrdersInner) HasNumberOfLinesWithWarning() bool`

HasNumberOfLinesWithWarning returns a boolean if a field has been set.

### GetIngramOrderNumber

`func (o *OrderCreateResponseOrdersInner) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderCreateResponseOrdersInner) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderCreateResponseOrdersInner) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderCreateResponseOrdersInner) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *OrderCreateResponseOrdersInner) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *OrderCreateResponseOrdersInner) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *OrderCreateResponseOrdersInner) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *OrderCreateResponseOrdersInner) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateResponseOrdersInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateResponseOrdersInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateResponseOrdersInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateResponseOrdersInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderCreateResponseOrdersInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderCreateResponseOrdersInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderCreateResponseOrdersInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderCreateResponseOrdersInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderCreateResponseOrdersInner) GetOrderTotal() float32`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderCreateResponseOrdersInner) GetOrderTotalOk() (*float32, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderCreateResponseOrdersInner) SetOrderTotal(v float32)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderCreateResponseOrdersInner) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetFreightCharges

`func (o *OrderCreateResponseOrdersInner) GetFreightCharges() float32`

GetFreightCharges returns the FreightCharges field if non-nil, zero value otherwise.

### GetFreightChargesOk

`func (o *OrderCreateResponseOrdersInner) GetFreightChargesOk() (*float32, bool)`

GetFreightChargesOk returns a tuple with the FreightCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCharges

`func (o *OrderCreateResponseOrdersInner) SetFreightCharges(v float32)`

SetFreightCharges sets FreightCharges field to given value.

### HasFreightCharges

`func (o *OrderCreateResponseOrdersInner) HasFreightCharges() bool`

HasFreightCharges returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderCreateResponseOrdersInner) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderCreateResponseOrdersInner) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderCreateResponseOrdersInner) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderCreateResponseOrdersInner) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderCreateResponseOrdersInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderCreateResponseOrdersInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderCreateResponseOrdersInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderCreateResponseOrdersInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetLines

`func (o *OrderCreateResponseOrdersInner) GetLines() []OrderCreateResponseOrdersInnerLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderCreateResponseOrdersInner) GetLinesOk() (*[]OrderCreateResponseOrdersInnerLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderCreateResponseOrdersInner) SetLines(v []OrderCreateResponseOrdersInnerLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderCreateResponseOrdersInner) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMiscellaneousCharges

`func (o *OrderCreateResponseOrdersInner) GetMiscellaneousCharges() []OrderCreateResponseOrdersInnerMiscellaneousChargesInner`

GetMiscellaneousCharges returns the MiscellaneousCharges field if non-nil, zero value otherwise.

### GetMiscellaneousChargesOk

`func (o *OrderCreateResponseOrdersInner) GetMiscellaneousChargesOk() (*[]OrderCreateResponseOrdersInnerMiscellaneousChargesInner, bool)`

GetMiscellaneousChargesOk returns a tuple with the MiscellaneousCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscellaneousCharges

`func (o *OrderCreateResponseOrdersInner) SetMiscellaneousCharges(v []OrderCreateResponseOrdersInnerMiscellaneousChargesInner)`

SetMiscellaneousCharges sets MiscellaneousCharges field to given value.

### HasMiscellaneousCharges

`func (o *OrderCreateResponseOrdersInner) HasMiscellaneousCharges() bool`

HasMiscellaneousCharges returns a boolean if a field has been set.

### GetLinks

`func (o *OrderCreateResponseOrdersInner) GetLinks() []OrderCreateResponseOrdersInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderCreateResponseOrdersInner) GetLinksOk() (*[]OrderCreateResponseOrdersInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderCreateResponseOrdersInner) SetLinks(v []OrderCreateResponseOrdersInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderCreateResponseOrdersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRejectedLineItems

`func (o *OrderCreateResponseOrdersInner) GetRejectedLineItems() []OrderCreateResponseOrdersInnerRejectedLineItemsInner`

GetRejectedLineItems returns the RejectedLineItems field if non-nil, zero value otherwise.

### GetRejectedLineItemsOk

`func (o *OrderCreateResponseOrdersInner) GetRejectedLineItemsOk() (*[]OrderCreateResponseOrdersInnerRejectedLineItemsInner, bool)`

GetRejectedLineItemsOk returns a tuple with the RejectedLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedLineItems

`func (o *OrderCreateResponseOrdersInner) SetRejectedLineItems(v []OrderCreateResponseOrdersInnerRejectedLineItemsInner)`

SetRejectedLineItems sets RejectedLineItems field to given value.

### HasRejectedLineItems

`func (o *OrderCreateResponseOrdersInner) HasRejectedLineItems() bool`

HasRejectedLineItems returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderCreateResponseOrdersInner) GetAdditionalAttributes() []OrderCreateResponseOrdersInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderCreateResponseOrdersInner) GetAdditionalAttributesOk() (*[]OrderCreateResponseOrdersInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderCreateResponseOrdersInner) SetAdditionalAttributes(v []OrderCreateResponseOrdersInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderCreateResponseOrdersInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


