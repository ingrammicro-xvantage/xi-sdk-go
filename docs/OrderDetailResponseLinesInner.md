# OrderDetailResponseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**IngramOrderLineNumber** | Pointer to **string** | Unique Ingram Micro line number. Starts with 001. | [optional] 
**VendorSalesOrderLineNumber** | Pointer to **string** | The vendor&#39;s sales order line number. | [optional] 
**CustomerLinenumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. | [optional] 
**LineStatus** | Pointer to **string** | The status for the line item in the order. One of- Backordered, In Progress, Shipped, Delivered, Canceled, On Hold | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**VendorName** | Pointer to **string** | The vendor&#39;s name for the part in their system. | [optional] 
**PartDescription** | Pointer to **string** | The vendor&#39;s description of the part in their system. | [optional] 
**UnitWeight** | Pointer to **float32** | The unit weight of the line item. | [optional] 
**WeightUom** | Pointer to **string** | The unit of measure for the line item. | [optional] 
**UnitPrice** | Pointer to **int32** | The unit price of the line item. | [optional] 
**UpcCode** | Pointer to **string** | The UPC code of a product. | [optional] 
**ExtendedPrice** | Pointer to **float32** | Unit price X quantity for the line item. | [optional] 
**TaxAmount** | Pointer to **float32** | The tax amount for the line item. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three character ISO 4217 currency code for the line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity ordered of the line item. | [optional] 
**QuantityConfirmed** | Pointer to **int32** | The quantity confirmed for the line item. | [optional] 
**QuantityBackOrdered** | Pointer to **int32** | The quantity backordered for the line item. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**RequestedDeliveryDate** | Pointer to **string** | Reseller-requested delivery date. Delivery date is not guaranteed. | [optional] 
**PromisedDeliveryDate** | Pointer to **string** | The delivery date promised by IngramMicro. | [optional] 
**LineNotes** | Pointer to **string** | Line-level notes for the order. | [optional] 
**ShipmentDetails** | Pointer to [**[]OrderDetailResponseLinesInnerShipmentDetailsInner**](OrderDetailResponseLinesInnerShipmentDetailsInner.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderDetailResponseLinesInnerAdditionalAttributesInner**](OrderDetailResponseLinesInnerAdditionalAttributesInner.md) |  | [optional] 
**Links** | Pointer to [**[]OrderDetailResponseLinesInnerLinksInner**](OrderDetailResponseLinesInnerLinksInner.md) |  | [optional] 

## Methods

### NewOrderDetailResponseLinesInner

`func NewOrderDetailResponseLinesInner() *OrderDetailResponseLinesInner`

NewOrderDetailResponseLinesInner instantiates a new OrderDetailResponseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseLinesInnerWithDefaults

`func NewOrderDetailResponseLinesInnerWithDefaults() *OrderDetailResponseLinesInner`

NewOrderDetailResponseLinesInnerWithDefaults instantiates a new OrderDetailResponseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderDetailResponseLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderDetailResponseLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderDetailResponseLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderDetailResponseLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetIngramOrderLineNumber

`func (o *OrderDetailResponseLinesInner) GetIngramOrderLineNumber() string`

GetIngramOrderLineNumber returns the IngramOrderLineNumber field if non-nil, zero value otherwise.

### GetIngramOrderLineNumberOk

`func (o *OrderDetailResponseLinesInner) GetIngramOrderLineNumberOk() (*string, bool)`

GetIngramOrderLineNumberOk returns a tuple with the IngramOrderLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderLineNumber

`func (o *OrderDetailResponseLinesInner) SetIngramOrderLineNumber(v string)`

SetIngramOrderLineNumber sets IngramOrderLineNumber field to given value.

### HasIngramOrderLineNumber

`func (o *OrderDetailResponseLinesInner) HasIngramOrderLineNumber() bool`

HasIngramOrderLineNumber returns a boolean if a field has been set.

### GetVendorSalesOrderLineNumber

`func (o *OrderDetailResponseLinesInner) GetVendorSalesOrderLineNumber() string`

GetVendorSalesOrderLineNumber returns the VendorSalesOrderLineNumber field if non-nil, zero value otherwise.

### GetVendorSalesOrderLineNumberOk

`func (o *OrderDetailResponseLinesInner) GetVendorSalesOrderLineNumberOk() (*string, bool)`

GetVendorSalesOrderLineNumberOk returns a tuple with the VendorSalesOrderLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSalesOrderLineNumber

`func (o *OrderDetailResponseLinesInner) SetVendorSalesOrderLineNumber(v string)`

SetVendorSalesOrderLineNumber sets VendorSalesOrderLineNumber field to given value.

### HasVendorSalesOrderLineNumber

`func (o *OrderDetailResponseLinesInner) HasVendorSalesOrderLineNumber() bool`

HasVendorSalesOrderLineNumber returns a boolean if a field has been set.

### GetCustomerLinenumber

`func (o *OrderDetailResponseLinesInner) GetCustomerLinenumber() string`

GetCustomerLinenumber returns the CustomerLinenumber field if non-nil, zero value otherwise.

### GetCustomerLinenumberOk

`func (o *OrderDetailResponseLinesInner) GetCustomerLinenumberOk() (*string, bool)`

GetCustomerLinenumberOk returns a tuple with the CustomerLinenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLinenumber

`func (o *OrderDetailResponseLinesInner) SetCustomerLinenumber(v string)`

SetCustomerLinenumber sets CustomerLinenumber field to given value.

### HasCustomerLinenumber

`func (o *OrderDetailResponseLinesInner) HasCustomerLinenumber() bool`

HasCustomerLinenumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *OrderDetailResponseLinesInner) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OrderDetailResponseLinesInner) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OrderDetailResponseLinesInner) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OrderDetailResponseLinesInner) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderDetailResponseLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderDetailResponseLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderDetailResponseLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderDetailResponseLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderDetailResponseLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderDetailResponseLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderDetailResponseLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderDetailResponseLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *OrderDetailResponseLinesInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *OrderDetailResponseLinesInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *OrderDetailResponseLinesInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *OrderDetailResponseLinesInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetPartDescription

`func (o *OrderDetailResponseLinesInner) GetPartDescription() string`

GetPartDescription returns the PartDescription field if non-nil, zero value otherwise.

### GetPartDescriptionOk

`func (o *OrderDetailResponseLinesInner) GetPartDescriptionOk() (*string, bool)`

GetPartDescriptionOk returns a tuple with the PartDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartDescription

`func (o *OrderDetailResponseLinesInner) SetPartDescription(v string)`

SetPartDescription sets PartDescription field to given value.

### HasPartDescription

`func (o *OrderDetailResponseLinesInner) HasPartDescription() bool`

HasPartDescription returns a boolean if a field has been set.

### GetUnitWeight

`func (o *OrderDetailResponseLinesInner) GetUnitWeight() float32`

GetUnitWeight returns the UnitWeight field if non-nil, zero value otherwise.

### GetUnitWeightOk

`func (o *OrderDetailResponseLinesInner) GetUnitWeightOk() (*float32, bool)`

GetUnitWeightOk returns a tuple with the UnitWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitWeight

`func (o *OrderDetailResponseLinesInner) SetUnitWeight(v float32)`

SetUnitWeight sets UnitWeight field to given value.

### HasUnitWeight

`func (o *OrderDetailResponseLinesInner) HasUnitWeight() bool`

HasUnitWeight returns a boolean if a field has been set.

### GetWeightUom

`func (o *OrderDetailResponseLinesInner) GetWeightUom() string`

GetWeightUom returns the WeightUom field if non-nil, zero value otherwise.

### GetWeightUomOk

`func (o *OrderDetailResponseLinesInner) GetWeightUomOk() (*string, bool)`

GetWeightUomOk returns a tuple with the WeightUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUom

`func (o *OrderDetailResponseLinesInner) SetWeightUom(v string)`

SetWeightUom sets WeightUom field to given value.

### HasWeightUom

`func (o *OrderDetailResponseLinesInner) HasWeightUom() bool`

HasWeightUom returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderDetailResponseLinesInner) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderDetailResponseLinesInner) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderDetailResponseLinesInner) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderDetailResponseLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUpcCode

`func (o *OrderDetailResponseLinesInner) GetUpcCode() string`

GetUpcCode returns the UpcCode field if non-nil, zero value otherwise.

### GetUpcCodeOk

`func (o *OrderDetailResponseLinesInner) GetUpcCodeOk() (*string, bool)`

GetUpcCodeOk returns a tuple with the UpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcCode

`func (o *OrderDetailResponseLinesInner) SetUpcCode(v string)`

SetUpcCode sets UpcCode field to given value.

### HasUpcCode

`func (o *OrderDetailResponseLinesInner) HasUpcCode() bool`

HasUpcCode returns a boolean if a field has been set.

### GetExtendedPrice

`func (o *OrderDetailResponseLinesInner) GetExtendedPrice() float32`

GetExtendedPrice returns the ExtendedPrice field if non-nil, zero value otherwise.

### GetExtendedPriceOk

`func (o *OrderDetailResponseLinesInner) GetExtendedPriceOk() (*float32, bool)`

GetExtendedPriceOk returns a tuple with the ExtendedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPrice

`func (o *OrderDetailResponseLinesInner) SetExtendedPrice(v float32)`

SetExtendedPrice sets ExtendedPrice field to given value.

### HasExtendedPrice

`func (o *OrderDetailResponseLinesInner) HasExtendedPrice() bool`

HasExtendedPrice returns a boolean if a field has been set.

### GetTaxAmount

`func (o *OrderDetailResponseLinesInner) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *OrderDetailResponseLinesInner) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *OrderDetailResponseLinesInner) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *OrderDetailResponseLinesInner) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderDetailResponseLinesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderDetailResponseLinesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderDetailResponseLinesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderDetailResponseLinesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderDetailResponseLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderDetailResponseLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderDetailResponseLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderDetailResponseLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityConfirmed

`func (o *OrderDetailResponseLinesInner) GetQuantityConfirmed() int32`

GetQuantityConfirmed returns the QuantityConfirmed field if non-nil, zero value otherwise.

### GetQuantityConfirmedOk

`func (o *OrderDetailResponseLinesInner) GetQuantityConfirmedOk() (*int32, bool)`

GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfirmed

`func (o *OrderDetailResponseLinesInner) SetQuantityConfirmed(v int32)`

SetQuantityConfirmed sets QuantityConfirmed field to given value.

### HasQuantityConfirmed

`func (o *OrderDetailResponseLinesInner) HasQuantityConfirmed() bool`

HasQuantityConfirmed returns a boolean if a field has been set.

### GetQuantityBackOrdered

`func (o *OrderDetailResponseLinesInner) GetQuantityBackOrdered() int32`

GetQuantityBackOrdered returns the QuantityBackOrdered field if non-nil, zero value otherwise.

### GetQuantityBackOrderedOk

`func (o *OrderDetailResponseLinesInner) GetQuantityBackOrderedOk() (*int32, bool)`

GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackOrdered

`func (o *OrderDetailResponseLinesInner) SetQuantityBackOrdered(v int32)`

SetQuantityBackOrdered sets QuantityBackOrdered field to given value.

### HasQuantityBackOrdered

`func (o *OrderDetailResponseLinesInner) HasQuantityBackOrdered() bool`

HasQuantityBackOrdered returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderDetailResponseLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderDetailResponseLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderDetailResponseLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderDetailResponseLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetRequestedDeliveryDate

`func (o *OrderDetailResponseLinesInner) GetRequestedDeliveryDate() string`

GetRequestedDeliveryDate returns the RequestedDeliveryDate field if non-nil, zero value otherwise.

### GetRequestedDeliveryDateOk

`func (o *OrderDetailResponseLinesInner) GetRequestedDeliveryDateOk() (*string, bool)`

GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDeliveryDate

`func (o *OrderDetailResponseLinesInner) SetRequestedDeliveryDate(v string)`

SetRequestedDeliveryDate sets RequestedDeliveryDate field to given value.

### HasRequestedDeliveryDate

`func (o *OrderDetailResponseLinesInner) HasRequestedDeliveryDate() bool`

HasRequestedDeliveryDate returns a boolean if a field has been set.

### GetPromisedDeliveryDate

`func (o *OrderDetailResponseLinesInner) GetPromisedDeliveryDate() string`

GetPromisedDeliveryDate returns the PromisedDeliveryDate field if non-nil, zero value otherwise.

### GetPromisedDeliveryDateOk

`func (o *OrderDetailResponseLinesInner) GetPromisedDeliveryDateOk() (*string, bool)`

GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDeliveryDate

`func (o *OrderDetailResponseLinesInner) SetPromisedDeliveryDate(v string)`

SetPromisedDeliveryDate sets PromisedDeliveryDate field to given value.

### HasPromisedDeliveryDate

`func (o *OrderDetailResponseLinesInner) HasPromisedDeliveryDate() bool`

HasPromisedDeliveryDate returns a boolean if a field has been set.

### GetLineNotes

`func (o *OrderDetailResponseLinesInner) GetLineNotes() string`

GetLineNotes returns the LineNotes field if non-nil, zero value otherwise.

### GetLineNotesOk

`func (o *OrderDetailResponseLinesInner) GetLineNotesOk() (*string, bool)`

GetLineNotesOk returns a tuple with the LineNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNotes

`func (o *OrderDetailResponseLinesInner) SetLineNotes(v string)`

SetLineNotes sets LineNotes field to given value.

### HasLineNotes

`func (o *OrderDetailResponseLinesInner) HasLineNotes() bool`

HasLineNotes returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderDetailResponseLinesInner) GetShipmentDetails() []OrderDetailResponseLinesInnerShipmentDetailsInner`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderDetailResponseLinesInner) GetShipmentDetailsOk() (*[]OrderDetailResponseLinesInnerShipmentDetailsInner, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderDetailResponseLinesInner) SetShipmentDetails(v []OrderDetailResponseLinesInnerShipmentDetailsInner)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderDetailResponseLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderDetailResponseLinesInner) GetAdditionalAttributes() []OrderDetailResponseLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderDetailResponseLinesInner) GetAdditionalAttributesOk() (*[]OrderDetailResponseLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderDetailResponseLinesInner) SetAdditionalAttributes(v []OrderDetailResponseLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderDetailResponseLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *OrderDetailResponseLinesInner) GetLinks() []OrderDetailResponseLinesInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderDetailResponseLinesInner) GetLinksOk() (*[]OrderDetailResponseLinesInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderDetailResponseLinesInner) SetLinks(v []OrderDetailResponseLinesInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderDetailResponseLinesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


