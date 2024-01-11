# OrderDetailB2BLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**IngramOrderLineNumber** | Pointer to **string** | Unique Ingram Micro line number. Starts with 001. | [optional] 
**VendorSalesOrderLineNumber** | Pointer to **string** | The vendor&#39;s sales order line number. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. | [optional] 
**LineStatus** | Pointer to **string** | The status for the line item in the order. One of- Backordered, In Progress, Shipped, Delivered, Canceled, On Hold. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**VendorName** | Pointer to **string** | The vendor&#39;s name for the part in their system. | [optional] 
**PartDescription** | Pointer to **string** | The vendor&#39;s description of the part in their system. | [optional] 
**UnitWeight** | Pointer to **float64** | The unit weight of the line item. | [optional] 
**WeightUom** | Pointer to **string** | The unit of measure for the line item. | [optional] 
**UnitPrice** | Pointer to **float64** | The unit price of the line item. | [optional] 
**UpcCode** | Pointer to **string** | The UPC code of a product. | [optional] 
**ExtendedPrice** | Pointer to **float64** | Unit price X quantity for the line item. | [optional] 
**TaxAmount** | Pointer to **float64** | The tax amount for the line item. | [optional] 
**CurrencyCode** | Pointer to **string** | The country-specific three character ISO 4217 currency code for the line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity ordered of the line item. | [optional] 
**QuantityConfirmed** | Pointer to **int32** | The quantity confirmed for the line item. | [optional] 
**QuantityBackOrdered** | Pointer to **int32** | The quantity backordered for the line item. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**RequestedDeliverydate** | Pointer to **string** | Reseller-requested delivery date. Delivery date is not guaranteed. | [optional] 
**PromisedDeliveryDate** | Pointer to **string** | The delivery date promised by IngramMicro. | [optional] 
**LineNotes** | Pointer to **string** | Line-level notes for the order. | [optional] 
**ShipmentDetails** | Pointer to [**[]OrderDetailB2BLinesInnerShipmentDetailsInner**](OrderDetailB2BLinesInnerShipmentDetailsInner.md) | Shipping details for the line item. | [optional] 
**ServiceContractInfo** | Pointer to [**OrderDetailB2BLinesInnerServiceContractInfo**](OrderDetailB2BLinesInnerServiceContractInfo.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderDetailB2BLinesInnerAdditionalAttributesInner**](OrderDetailB2BLinesInnerAdditionalAttributesInner.md) |  | [optional] 
**Links** | Pointer to [**[]OrderDetailB2BLinesInnerLinksInner**](OrderDetailB2BLinesInnerLinksInner.md) |  | [optional] 
**EstimatedDates** | Pointer to [**[]OrderDetailB2BLinesInnerEstimatedDatesInner**](OrderDetailB2BLinesInnerEstimatedDatesInner.md) |  | [optional] 
**ScheduleLines** | Pointer to [**[]OrderDetailB2BLinesInnerScheduleLinesInner**](OrderDetailB2BLinesInnerScheduleLinesInner.md) |  | [optional] 
**MultipleShipments** | Pointer to [**[]OrderDetailB2BLinesInnerMultipleShipmentsInner**](OrderDetailB2BLinesInnerMultipleShipmentsInner.md) |  | [optional] 

## Methods

### NewOrderDetailB2BLinesInner

`func NewOrderDetailB2BLinesInner() *OrderDetailB2BLinesInner`

NewOrderDetailB2BLinesInner instantiates a new OrderDetailB2BLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailB2BLinesInnerWithDefaults

`func NewOrderDetailB2BLinesInnerWithDefaults() *OrderDetailB2BLinesInner`

NewOrderDetailB2BLinesInnerWithDefaults instantiates a new OrderDetailB2BLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderDetailB2BLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderDetailB2BLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderDetailB2BLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderDetailB2BLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetIngramOrderLineNumber

`func (o *OrderDetailB2BLinesInner) GetIngramOrderLineNumber() string`

GetIngramOrderLineNumber returns the IngramOrderLineNumber field if non-nil, zero value otherwise.

### GetIngramOrderLineNumberOk

`func (o *OrderDetailB2BLinesInner) GetIngramOrderLineNumberOk() (*string, bool)`

GetIngramOrderLineNumberOk returns a tuple with the IngramOrderLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderLineNumber

`func (o *OrderDetailB2BLinesInner) SetIngramOrderLineNumber(v string)`

SetIngramOrderLineNumber sets IngramOrderLineNumber field to given value.

### HasIngramOrderLineNumber

`func (o *OrderDetailB2BLinesInner) HasIngramOrderLineNumber() bool`

HasIngramOrderLineNumber returns a boolean if a field has been set.

### GetVendorSalesOrderLineNumber

`func (o *OrderDetailB2BLinesInner) GetVendorSalesOrderLineNumber() string`

GetVendorSalesOrderLineNumber returns the VendorSalesOrderLineNumber field if non-nil, zero value otherwise.

### GetVendorSalesOrderLineNumberOk

`func (o *OrderDetailB2BLinesInner) GetVendorSalesOrderLineNumberOk() (*string, bool)`

GetVendorSalesOrderLineNumberOk returns a tuple with the VendorSalesOrderLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSalesOrderLineNumber

`func (o *OrderDetailB2BLinesInner) SetVendorSalesOrderLineNumber(v string)`

SetVendorSalesOrderLineNumber sets VendorSalesOrderLineNumber field to given value.

### HasVendorSalesOrderLineNumber

`func (o *OrderDetailB2BLinesInner) HasVendorSalesOrderLineNumber() bool`

HasVendorSalesOrderLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderDetailB2BLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderDetailB2BLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderDetailB2BLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderDetailB2BLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *OrderDetailB2BLinesInner) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OrderDetailB2BLinesInner) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OrderDetailB2BLinesInner) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OrderDetailB2BLinesInner) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderDetailB2BLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderDetailB2BLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderDetailB2BLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderDetailB2BLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderDetailB2BLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderDetailB2BLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderDetailB2BLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderDetailB2BLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *OrderDetailB2BLinesInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *OrderDetailB2BLinesInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *OrderDetailB2BLinesInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *OrderDetailB2BLinesInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetPartDescription

`func (o *OrderDetailB2BLinesInner) GetPartDescription() string`

GetPartDescription returns the PartDescription field if non-nil, zero value otherwise.

### GetPartDescriptionOk

`func (o *OrderDetailB2BLinesInner) GetPartDescriptionOk() (*string, bool)`

GetPartDescriptionOk returns a tuple with the PartDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartDescription

`func (o *OrderDetailB2BLinesInner) SetPartDescription(v string)`

SetPartDescription sets PartDescription field to given value.

### HasPartDescription

`func (o *OrderDetailB2BLinesInner) HasPartDescription() bool`

HasPartDescription returns a boolean if a field has been set.

### GetUnitWeight

`func (o *OrderDetailB2BLinesInner) GetUnitWeight() float64`

GetUnitWeight returns the UnitWeight field if non-nil, zero value otherwise.

### GetUnitWeightOk

`func (o *OrderDetailB2BLinesInner) GetUnitWeightOk() (*float64, bool)`

GetUnitWeightOk returns a tuple with the UnitWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitWeight

`func (o *OrderDetailB2BLinesInner) SetUnitWeight(v float64)`

SetUnitWeight sets UnitWeight field to given value.

### HasUnitWeight

`func (o *OrderDetailB2BLinesInner) HasUnitWeight() bool`

HasUnitWeight returns a boolean if a field has been set.

### GetWeightUom

`func (o *OrderDetailB2BLinesInner) GetWeightUom() string`

GetWeightUom returns the WeightUom field if non-nil, zero value otherwise.

### GetWeightUomOk

`func (o *OrderDetailB2BLinesInner) GetWeightUomOk() (*string, bool)`

GetWeightUomOk returns a tuple with the WeightUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUom

`func (o *OrderDetailB2BLinesInner) SetWeightUom(v string)`

SetWeightUom sets WeightUom field to given value.

### HasWeightUom

`func (o *OrderDetailB2BLinesInner) HasWeightUom() bool`

HasWeightUom returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderDetailB2BLinesInner) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderDetailB2BLinesInner) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderDetailB2BLinesInner) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderDetailB2BLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUpcCode

`func (o *OrderDetailB2BLinesInner) GetUpcCode() string`

GetUpcCode returns the UpcCode field if non-nil, zero value otherwise.

### GetUpcCodeOk

`func (o *OrderDetailB2BLinesInner) GetUpcCodeOk() (*string, bool)`

GetUpcCodeOk returns a tuple with the UpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcCode

`func (o *OrderDetailB2BLinesInner) SetUpcCode(v string)`

SetUpcCode sets UpcCode field to given value.

### HasUpcCode

`func (o *OrderDetailB2BLinesInner) HasUpcCode() bool`

HasUpcCode returns a boolean if a field has been set.

### GetExtendedPrice

`func (o *OrderDetailB2BLinesInner) GetExtendedPrice() float64`

GetExtendedPrice returns the ExtendedPrice field if non-nil, zero value otherwise.

### GetExtendedPriceOk

`func (o *OrderDetailB2BLinesInner) GetExtendedPriceOk() (*float64, bool)`

GetExtendedPriceOk returns a tuple with the ExtendedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPrice

`func (o *OrderDetailB2BLinesInner) SetExtendedPrice(v float64)`

SetExtendedPrice sets ExtendedPrice field to given value.

### HasExtendedPrice

`func (o *OrderDetailB2BLinesInner) HasExtendedPrice() bool`

HasExtendedPrice returns a boolean if a field has been set.

### GetTaxAmount

`func (o *OrderDetailB2BLinesInner) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *OrderDetailB2BLinesInner) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *OrderDetailB2BLinesInner) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *OrderDetailB2BLinesInner) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrderDetailB2BLinesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrderDetailB2BLinesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrderDetailB2BLinesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrderDetailB2BLinesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderDetailB2BLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderDetailB2BLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderDetailB2BLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderDetailB2BLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityConfirmed

`func (o *OrderDetailB2BLinesInner) GetQuantityConfirmed() int32`

GetQuantityConfirmed returns the QuantityConfirmed field if non-nil, zero value otherwise.

### GetQuantityConfirmedOk

`func (o *OrderDetailB2BLinesInner) GetQuantityConfirmedOk() (*int32, bool)`

GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfirmed

`func (o *OrderDetailB2BLinesInner) SetQuantityConfirmed(v int32)`

SetQuantityConfirmed sets QuantityConfirmed field to given value.

### HasQuantityConfirmed

`func (o *OrderDetailB2BLinesInner) HasQuantityConfirmed() bool`

HasQuantityConfirmed returns a boolean if a field has been set.

### GetQuantityBackOrdered

`func (o *OrderDetailB2BLinesInner) GetQuantityBackOrdered() int32`

GetQuantityBackOrdered returns the QuantityBackOrdered field if non-nil, zero value otherwise.

### GetQuantityBackOrderedOk

`func (o *OrderDetailB2BLinesInner) GetQuantityBackOrderedOk() (*int32, bool)`

GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackOrdered

`func (o *OrderDetailB2BLinesInner) SetQuantityBackOrdered(v int32)`

SetQuantityBackOrdered sets QuantityBackOrdered field to given value.

### HasQuantityBackOrdered

`func (o *OrderDetailB2BLinesInner) HasQuantityBackOrdered() bool`

HasQuantityBackOrdered returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderDetailB2BLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderDetailB2BLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderDetailB2BLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderDetailB2BLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetRequestedDeliverydate

`func (o *OrderDetailB2BLinesInner) GetRequestedDeliverydate() string`

GetRequestedDeliverydate returns the RequestedDeliverydate field if non-nil, zero value otherwise.

### GetRequestedDeliverydateOk

`func (o *OrderDetailB2BLinesInner) GetRequestedDeliverydateOk() (*string, bool)`

GetRequestedDeliverydateOk returns a tuple with the RequestedDeliverydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDeliverydate

`func (o *OrderDetailB2BLinesInner) SetRequestedDeliverydate(v string)`

SetRequestedDeliverydate sets RequestedDeliverydate field to given value.

### HasRequestedDeliverydate

`func (o *OrderDetailB2BLinesInner) HasRequestedDeliverydate() bool`

HasRequestedDeliverydate returns a boolean if a field has been set.

### GetPromisedDeliveryDate

`func (o *OrderDetailB2BLinesInner) GetPromisedDeliveryDate() string`

GetPromisedDeliveryDate returns the PromisedDeliveryDate field if non-nil, zero value otherwise.

### GetPromisedDeliveryDateOk

`func (o *OrderDetailB2BLinesInner) GetPromisedDeliveryDateOk() (*string, bool)`

GetPromisedDeliveryDateOk returns a tuple with the PromisedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDeliveryDate

`func (o *OrderDetailB2BLinesInner) SetPromisedDeliveryDate(v string)`

SetPromisedDeliveryDate sets PromisedDeliveryDate field to given value.

### HasPromisedDeliveryDate

`func (o *OrderDetailB2BLinesInner) HasPromisedDeliveryDate() bool`

HasPromisedDeliveryDate returns a boolean if a field has been set.

### GetLineNotes

`func (o *OrderDetailB2BLinesInner) GetLineNotes() string`

GetLineNotes returns the LineNotes field if non-nil, zero value otherwise.

### GetLineNotesOk

`func (o *OrderDetailB2BLinesInner) GetLineNotesOk() (*string, bool)`

GetLineNotesOk returns a tuple with the LineNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNotes

`func (o *OrderDetailB2BLinesInner) SetLineNotes(v string)`

SetLineNotes sets LineNotes field to given value.

### HasLineNotes

`func (o *OrderDetailB2BLinesInner) HasLineNotes() bool`

HasLineNotes returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderDetailB2BLinesInner) GetShipmentDetails() []OrderDetailB2BLinesInnerShipmentDetailsInner`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderDetailB2BLinesInner) GetShipmentDetailsOk() (*[]OrderDetailB2BLinesInnerShipmentDetailsInner, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderDetailB2BLinesInner) SetShipmentDetails(v []OrderDetailB2BLinesInnerShipmentDetailsInner)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderDetailB2BLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### SetShipmentDetailsNil

`func (o *OrderDetailB2BLinesInner) SetShipmentDetailsNil(b bool)`

 SetShipmentDetailsNil sets the value for ShipmentDetails to be an explicit nil

### UnsetShipmentDetails
`func (o *OrderDetailB2BLinesInner) UnsetShipmentDetails()`

UnsetShipmentDetails ensures that no value is present for ShipmentDetails, not even an explicit nil
### GetServiceContractInfo

`func (o *OrderDetailB2BLinesInner) GetServiceContractInfo() OrderDetailB2BLinesInnerServiceContractInfo`

GetServiceContractInfo returns the ServiceContractInfo field if non-nil, zero value otherwise.

### GetServiceContractInfoOk

`func (o *OrderDetailB2BLinesInner) GetServiceContractInfoOk() (*OrderDetailB2BLinesInnerServiceContractInfo, bool)`

GetServiceContractInfoOk returns a tuple with the ServiceContractInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContractInfo

`func (o *OrderDetailB2BLinesInner) SetServiceContractInfo(v OrderDetailB2BLinesInnerServiceContractInfo)`

SetServiceContractInfo sets ServiceContractInfo field to given value.

### HasServiceContractInfo

`func (o *OrderDetailB2BLinesInner) HasServiceContractInfo() bool`

HasServiceContractInfo returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderDetailB2BLinesInner) GetAdditionalAttributes() []OrderDetailB2BLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderDetailB2BLinesInner) GetAdditionalAttributesOk() (*[]OrderDetailB2BLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderDetailB2BLinesInner) SetAdditionalAttributes(v []OrderDetailB2BLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderDetailB2BLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### SetAdditionalAttributesNil

`func (o *OrderDetailB2BLinesInner) SetAdditionalAttributesNil(b bool)`

 SetAdditionalAttributesNil sets the value for AdditionalAttributes to be an explicit nil

### UnsetAdditionalAttributes
`func (o *OrderDetailB2BLinesInner) UnsetAdditionalAttributes()`

UnsetAdditionalAttributes ensures that no value is present for AdditionalAttributes, not even an explicit nil
### GetLinks

`func (o *OrderDetailB2BLinesInner) GetLinks() []OrderDetailB2BLinesInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderDetailB2BLinesInner) GetLinksOk() (*[]OrderDetailB2BLinesInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderDetailB2BLinesInner) SetLinks(v []OrderDetailB2BLinesInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderDetailB2BLinesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *OrderDetailB2BLinesInner) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *OrderDetailB2BLinesInner) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetEstimatedDates

`func (o *OrderDetailB2BLinesInner) GetEstimatedDates() []OrderDetailB2BLinesInnerEstimatedDatesInner`

GetEstimatedDates returns the EstimatedDates field if non-nil, zero value otherwise.

### GetEstimatedDatesOk

`func (o *OrderDetailB2BLinesInner) GetEstimatedDatesOk() (*[]OrderDetailB2BLinesInnerEstimatedDatesInner, bool)`

GetEstimatedDatesOk returns a tuple with the EstimatedDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDates

`func (o *OrderDetailB2BLinesInner) SetEstimatedDates(v []OrderDetailB2BLinesInnerEstimatedDatesInner)`

SetEstimatedDates sets EstimatedDates field to given value.

### HasEstimatedDates

`func (o *OrderDetailB2BLinesInner) HasEstimatedDates() bool`

HasEstimatedDates returns a boolean if a field has been set.

### GetScheduleLines

`func (o *OrderDetailB2BLinesInner) GetScheduleLines() []OrderDetailB2BLinesInnerScheduleLinesInner`

GetScheduleLines returns the ScheduleLines field if non-nil, zero value otherwise.

### GetScheduleLinesOk

`func (o *OrderDetailB2BLinesInner) GetScheduleLinesOk() (*[]OrderDetailB2BLinesInnerScheduleLinesInner, bool)`

GetScheduleLinesOk returns a tuple with the ScheduleLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleLines

`func (o *OrderDetailB2BLinesInner) SetScheduleLines(v []OrderDetailB2BLinesInnerScheduleLinesInner)`

SetScheduleLines sets ScheduleLines field to given value.

### HasScheduleLines

`func (o *OrderDetailB2BLinesInner) HasScheduleLines() bool`

HasScheduleLines returns a boolean if a field has been set.

### GetMultipleShipments

`func (o *OrderDetailB2BLinesInner) GetMultipleShipments() []OrderDetailB2BLinesInnerMultipleShipmentsInner`

GetMultipleShipments returns the MultipleShipments field if non-nil, zero value otherwise.

### GetMultipleShipmentsOk

`func (o *OrderDetailB2BLinesInner) GetMultipleShipmentsOk() (*[]OrderDetailB2BLinesInnerMultipleShipmentsInner, bool)`

GetMultipleShipmentsOk returns a tuple with the MultipleShipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleShipments

`func (o *OrderDetailB2BLinesInner) SetMultipleShipments(v []OrderDetailB2BLinesInnerMultipleShipmentsInner)`

SetMultipleShipments sets MultipleShipments field to given value.

### HasMultipleShipments

`func (o *OrderDetailB2BLinesInner) HasMultipleShipments() bool`

HasMultipleShipments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


