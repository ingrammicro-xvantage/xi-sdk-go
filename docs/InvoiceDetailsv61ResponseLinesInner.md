# InvoiceDetailsv61ResponseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramLineNumber** | Pointer to **string** | Unique line number from Ingram. | [optional] 
**CustomerLineNumber** | Pointer to **string** | Line number passes by customer while creating an order. | [optional] [default to "0"]
**IngramPartNumber** | Pointer to **string** | Ingram Micro SKU (stock keeping unit). An identification, usually alphanumeric, of a particular product that allows it to be tracked for inventory purposes. | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number. | [optional] 
**CustomerPartNumber** | Pointer to **string** | Part number from customer&#39;s system. | [optional] 
**VendorName** | Pointer to **string** | Name of the vendor. | [optional] 
**ProductDescription** | Pointer to **string** | Description of the product. | [optional] 
**UnitWeight** | Pointer to **float32** | Weight of the product. | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the product. | [optional] 
**UnitPrice** | Pointer to **float64** | Unit price of the product. | [optional] 
**UnitOfMeasure** | Pointer to **string** | Unit of measure of the product. | [optional] 
**CurrencyCode** | Pointer to **string** | Currency code. | [optional] 
**ExtendedPrice** | Pointer to **float64** | Extended price of the product. | [optional] 
**TaxPercentage** | Pointer to **float64** | Tax percentage | [optional] 
**TaxRate** | Pointer to **float64** | Tax rate | [optional] 
**TaxAmount** | Pointer to **float64** | Line level tax amount. | [optional] 
**SerialNumbers** | Pointer to [**[]InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner**](InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner.md) |  | [optional] 
**QuantityOrdered** | Pointer to **int32** | Quantity ordered by the customer. | [optional] 
**QuantityShipped** | Pointer to **int32** | Quantity shipped to the customer. | [optional] 

## Methods

### NewInvoiceDetailsv61ResponseLinesInner

`func NewInvoiceDetailsv61ResponseLinesInner() *InvoiceDetailsv61ResponseLinesInner`

NewInvoiceDetailsv61ResponseLinesInner instantiates a new InvoiceDetailsv61ResponseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailsv61ResponseLinesInnerWithDefaults

`func NewInvoiceDetailsv61ResponseLinesInnerWithDefaults() *InvoiceDetailsv61ResponseLinesInner`

NewInvoiceDetailsv61ResponseLinesInnerWithDefaults instantiates a new InvoiceDetailsv61ResponseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetUpc

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *InvoiceDetailsv61ResponseLinesInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *InvoiceDetailsv61ResponseLinesInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetCustomerPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *InvoiceDetailsv61ResponseLinesInner) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *InvoiceDetailsv61ResponseLinesInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *InvoiceDetailsv61ResponseLinesInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *InvoiceDetailsv61ResponseLinesInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetProductDescription

`func (o *InvoiceDetailsv61ResponseLinesInner) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *InvoiceDetailsv61ResponseLinesInner) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *InvoiceDetailsv61ResponseLinesInner) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetUnitWeight

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitWeight() float32`

GetUnitWeight returns the UnitWeight field if non-nil, zero value otherwise.

### GetUnitWeightOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitWeightOk() (*float32, bool)`

GetUnitWeightOk returns a tuple with the UnitWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitWeight

`func (o *InvoiceDetailsv61ResponseLinesInner) SetUnitWeight(v float32)`

SetUnitWeight sets UnitWeight field to given value.

### HasUnitWeight

`func (o *InvoiceDetailsv61ResponseLinesInner) HasUnitWeight() bool`

HasUnitWeight returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceDetailsv61ResponseLinesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceDetailsv61ResponseLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *InvoiceDetailsv61ResponseLinesInner) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *InvoiceDetailsv61ResponseLinesInner) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceDetailsv61ResponseLinesInner) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InvoiceDetailsv61ResponseLinesInner) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetExtendedPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) GetExtendedPrice() float64`

GetExtendedPrice returns the ExtendedPrice field if non-nil, zero value otherwise.

### GetExtendedPriceOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetExtendedPriceOk() (*float64, bool)`

GetExtendedPriceOk returns a tuple with the ExtendedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) SetExtendedPrice(v float64)`

SetExtendedPrice sets ExtendedPrice field to given value.

### HasExtendedPrice

`func (o *InvoiceDetailsv61ResponseLinesInner) HasExtendedPrice() bool`

HasExtendedPrice returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxPercentage() float64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxPercentageOk() (*float64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *InvoiceDetailsv61ResponseLinesInner) SetTaxPercentage(v float64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *InvoiceDetailsv61ResponseLinesInner) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTaxRate

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *InvoiceDetailsv61ResponseLinesInner) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *InvoiceDetailsv61ResponseLinesInner) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxAmount

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceDetailsv61ResponseLinesInner) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *InvoiceDetailsv61ResponseLinesInner) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *InvoiceDetailsv61ResponseLinesInner) GetSerialNumbers() []InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetSerialNumbersOk() (*[]InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *InvoiceDetailsv61ResponseLinesInner) SetSerialNumbers(v []InvoiceDetailsv61ResponseLinesInnerSerialNumbersInner)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *InvoiceDetailsv61ResponseLinesInner) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### SetSerialNumbersNil

`func (o *InvoiceDetailsv61ResponseLinesInner) SetSerialNumbersNil(b bool)`

 SetSerialNumbersNil sets the value for SerialNumbers to be an explicit nil

### UnsetSerialNumbers
`func (o *InvoiceDetailsv61ResponseLinesInner) UnsetSerialNumbers()`

UnsetSerialNumbers ensures that no value is present for SerialNumbers, not even an explicit nil
### GetQuantityOrdered

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *InvoiceDetailsv61ResponseLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *InvoiceDetailsv61ResponseLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityShipped

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantityShipped() int32`

GetQuantityShipped returns the QuantityShipped field if non-nil, zero value otherwise.

### GetQuantityShippedOk

`func (o *InvoiceDetailsv61ResponseLinesInner) GetQuantityShippedOk() (*int32, bool)`

GetQuantityShippedOk returns a tuple with the QuantityShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityShipped

`func (o *InvoiceDetailsv61ResponseLinesInner) SetQuantityShipped(v int32)`

SetQuantityShipped sets QuantityShipped field to given value.

### HasQuantityShipped

`func (o *InvoiceDetailsv61ResponseLinesInner) HasQuantityShipped() bool`

HasQuantityShipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


