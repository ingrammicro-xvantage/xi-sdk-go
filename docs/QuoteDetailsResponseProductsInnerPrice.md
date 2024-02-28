# QuoteDetailsResponseProductsInnerPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotePrice** | Pointer to **int32** | Ingram Micro quoted price specific to the reseller and quote. | [optional] 
**Msrp** | Pointer to **int32** | Manufacturer Suggested Retail Price | [optional] 
**ExtendedMsrp** | Pointer to **int32** | Extended MSRP - Manufacturer Suggested Retail Price X Quantity | [optional] 
**ExtendedQuotePrice** | Pointer to **int32** | Extended reseller quoted price (cost to reseller) X Quantity | [optional] 
**DiscountOffList** | Pointer to **string** | Discount off list percentage extended | [optional] 
**Vendorprice** | Pointer to **float32** |  | [optional] 
**Extendedvendorprice** | Pointer to **float32** |  | [optional] 
**TotalVisibleReserveQuantity** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RecurringPriceModel** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteDetailsResponseProductsInnerPrice

`func NewQuoteDetailsResponseProductsInnerPrice() *QuoteDetailsResponseProductsInnerPrice`

NewQuoteDetailsResponseProductsInnerPrice instantiates a new QuoteDetailsResponseProductsInnerPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseProductsInnerPriceWithDefaults

`func NewQuoteDetailsResponseProductsInnerPriceWithDefaults() *QuoteDetailsResponseProductsInnerPrice`

NewQuoteDetailsResponseProductsInnerPriceWithDefaults instantiates a new QuoteDetailsResponseProductsInnerPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetQuotePrice() int32`

GetQuotePrice returns the QuotePrice field if non-nil, zero value otherwise.

### GetQuotePriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetQuotePriceOk() (*int32, bool)`

GetQuotePriceOk returns a tuple with the QuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetQuotePrice(v int32)`

SetQuotePrice sets QuotePrice field to given value.

### HasQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasQuotePrice() bool`

HasQuotePrice returns a boolean if a field has been set.

### GetMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) GetMsrp() int32`

GetMsrp returns the Msrp field if non-nil, zero value otherwise.

### GetMsrpOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetMsrpOk() (*int32, bool)`

GetMsrpOk returns a tuple with the Msrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) SetMsrp(v int32)`

SetMsrp sets Msrp field to given value.

### HasMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) HasMsrp() bool`

HasMsrp returns a boolean if a field has been set.

### GetExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedMsrp() int32`

GetExtendedMsrp returns the ExtendedMsrp field if non-nil, zero value otherwise.

### GetExtendedMsrpOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedMsrpOk() (*int32, bool)`

GetExtendedMsrpOk returns a tuple with the ExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedMsrp(v int32)`

SetExtendedMsrp sets ExtendedMsrp field to given value.

### HasExtendedMsrp

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtendedMsrp() bool`

HasExtendedMsrp returns a boolean if a field has been set.

### GetExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedQuotePrice() int32`

GetExtendedQuotePrice returns the ExtendedQuotePrice field if non-nil, zero value otherwise.

### GetExtendedQuotePriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedQuotePriceOk() (*int32, bool)`

GetExtendedQuotePriceOk returns a tuple with the ExtendedQuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedQuotePrice(v int32)`

SetExtendedQuotePrice sets ExtendedQuotePrice field to given value.

### HasExtendedQuotePrice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtendedQuotePrice() bool`

HasExtendedQuotePrice returns a boolean if a field has been set.

### GetDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscountOffList() string`

GetDiscountOffList returns the DiscountOffList field if non-nil, zero value otherwise.

### GetDiscountOffListOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetDiscountOffListOk() (*string, bool)`

GetDiscountOffListOk returns a tuple with the DiscountOffList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) SetDiscountOffList(v string)`

SetDiscountOffList sets DiscountOffList field to given value.

### HasDiscountOffList

`func (o *QuoteDetailsResponseProductsInnerPrice) HasDiscountOffList() bool`

HasDiscountOffList returns a boolean if a field has been set.

### GetVendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetVendorprice() float32`

GetVendorprice returns the Vendorprice field if non-nil, zero value otherwise.

### GetVendorpriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetVendorpriceOk() (*float32, bool)`

GetVendorpriceOk returns a tuple with the Vendorprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetVendorprice(v float32)`

SetVendorprice sets Vendorprice field to given value.

### HasVendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasVendorprice() bool`

HasVendorprice returns a boolean if a field has been set.

### GetExtendedvendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedvendorprice() float32`

GetExtendedvendorprice returns the Extendedvendorprice field if non-nil, zero value otherwise.

### GetExtendedvendorpriceOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetExtendedvendorpriceOk() (*float32, bool)`

GetExtendedvendorpriceOk returns a tuple with the Extendedvendorprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedvendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) SetExtendedvendorprice(v float32)`

SetExtendedvendorprice sets Extendedvendorprice field to given value.

### HasExtendedvendorprice

`func (o *QuoteDetailsResponseProductsInnerPrice) HasExtendedvendorprice() bool`

HasExtendedvendorprice returns a boolean if a field has been set.

### GetTotalVisibleReserveQuantity

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTotalVisibleReserveQuantity() int32`

GetTotalVisibleReserveQuantity returns the TotalVisibleReserveQuantity field if non-nil, zero value otherwise.

### GetTotalVisibleReserveQuantityOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTotalVisibleReserveQuantityOk() (*int32, bool)`

GetTotalVisibleReserveQuantityOk returns a tuple with the TotalVisibleReserveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVisibleReserveQuantity

`func (o *QuoteDetailsResponseProductsInnerPrice) SetTotalVisibleReserveQuantity(v int32)`

SetTotalVisibleReserveQuantity sets TotalVisibleReserveQuantity field to given value.

### HasTotalVisibleReserveQuantity

`func (o *QuoteDetailsResponseProductsInnerPrice) HasTotalVisibleReserveQuantity() bool`

HasTotalVisibleReserveQuantity returns a boolean if a field has been set.

### GetType

`func (o *QuoteDetailsResponseProductsInnerPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuoteDetailsResponseProductsInnerPrice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QuoteDetailsResponseProductsInnerPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRecurringPriceModel() string`

GetRecurringPriceModel returns the RecurringPriceModel field if non-nil, zero value otherwise.

### GetRecurringPriceModelOk

`func (o *QuoteDetailsResponseProductsInnerPrice) GetRecurringPriceModelOk() (*string, bool)`

GetRecurringPriceModelOk returns a tuple with the RecurringPriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) SetRecurringPriceModel(v string)`

SetRecurringPriceModel sets RecurringPriceModel field to given value.

### HasRecurringPriceModel

`func (o *QuoteDetailsResponseProductsInnerPrice) HasRecurringPriceModel() bool`

HasRecurringPriceModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


