# ReturnsDetailsResponseProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramLineNumber** | Pointer to **int32** | Unique Ingram Micro line number. | [optional] 
**Description** | Pointer to **string** | The description of the line item product. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**Upc** | Pointer to **string** | The UPC code of a product. | [optional] 
**InvoiceDate** | Pointer to **string** | The date of the invoice. | [optional] 
**InvoiceNumber** | Pointer to **string** | Ingram micro Invoice number. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**Quantity** | Pointer to **float32** | The quantity of the line item. | [optional] 
**UnitPrice** | Pointer to **float32** | The unit price of the line item. | [optional] 
**ExtendedPrice** | Pointer to **float32** | Unit price X quantity for the line item. | [optional] 
**Status** | Pointer to **string** | The status of the line item. | [optional] 
**ReturnBranch** | Pointer to **string** | The code of the return branch. | [optional] 
**ShipFromBranch** | Pointer to **string** | The code of the ship from branch. | [optional] 
**RequestDetails** | Pointer to **string** | Request details. | [optional] 
**AdditionalDetails** | Pointer to **string** |  | [optional] 

## Methods

### NewReturnsDetailsResponseProductsInner

`func NewReturnsDetailsResponseProductsInner() *ReturnsDetailsResponseProductsInner`

NewReturnsDetailsResponseProductsInner instantiates a new ReturnsDetailsResponseProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsDetailsResponseProductsInnerWithDefaults

`func NewReturnsDetailsResponseProductsInnerWithDefaults() *ReturnsDetailsResponseProductsInner`

NewReturnsDetailsResponseProductsInnerWithDefaults instantiates a new ReturnsDetailsResponseProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramLineNumber

`func (o *ReturnsDetailsResponseProductsInner) GetIngramLineNumber() int32`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *ReturnsDetailsResponseProductsInner) GetIngramLineNumberOk() (*int32, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *ReturnsDetailsResponseProductsInner) SetIngramLineNumber(v int32)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *ReturnsDetailsResponseProductsInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetDescription

`func (o *ReturnsDetailsResponseProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReturnsDetailsResponseProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReturnsDetailsResponseProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReturnsDetailsResponseProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *ReturnsDetailsResponseProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *ReturnsDetailsResponseProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *ReturnsDetailsResponseProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *ReturnsDetailsResponseProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *ReturnsDetailsResponseProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *ReturnsDetailsResponseProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *ReturnsDetailsResponseProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *ReturnsDetailsResponseProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetUpc

`func (o *ReturnsDetailsResponseProductsInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ReturnsDetailsResponseProductsInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ReturnsDetailsResponseProductsInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ReturnsDetailsResponseProductsInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *ReturnsDetailsResponseProductsInner) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ReturnsDetailsResponseProductsInner) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ReturnsDetailsResponseProductsInner) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *ReturnsDetailsResponseProductsInner) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *ReturnsDetailsResponseProductsInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ReturnsDetailsResponseProductsInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ReturnsDetailsResponseProductsInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *ReturnsDetailsResponseProductsInner) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *ReturnsDetailsResponseProductsInner) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *ReturnsDetailsResponseProductsInner) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *ReturnsDetailsResponseProductsInner) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *ReturnsDetailsResponseProductsInner) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnsDetailsResponseProductsInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnsDetailsResponseProductsInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnsDetailsResponseProductsInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnsDetailsResponseProductsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *ReturnsDetailsResponseProductsInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ReturnsDetailsResponseProductsInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ReturnsDetailsResponseProductsInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *ReturnsDetailsResponseProductsInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetExtendedPrice

`func (o *ReturnsDetailsResponseProductsInner) GetExtendedPrice() float32`

GetExtendedPrice returns the ExtendedPrice field if non-nil, zero value otherwise.

### GetExtendedPriceOk

`func (o *ReturnsDetailsResponseProductsInner) GetExtendedPriceOk() (*float32, bool)`

GetExtendedPriceOk returns a tuple with the ExtendedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPrice

`func (o *ReturnsDetailsResponseProductsInner) SetExtendedPrice(v float32)`

SetExtendedPrice sets ExtendedPrice field to given value.

### HasExtendedPrice

`func (o *ReturnsDetailsResponseProductsInner) HasExtendedPrice() bool`

HasExtendedPrice returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnsDetailsResponseProductsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsDetailsResponseProductsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsDetailsResponseProductsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsDetailsResponseProductsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnBranch

`func (o *ReturnsDetailsResponseProductsInner) GetReturnBranch() string`

GetReturnBranch returns the ReturnBranch field if non-nil, zero value otherwise.

### GetReturnBranchOk

`func (o *ReturnsDetailsResponseProductsInner) GetReturnBranchOk() (*string, bool)`

GetReturnBranchOk returns a tuple with the ReturnBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnBranch

`func (o *ReturnsDetailsResponseProductsInner) SetReturnBranch(v string)`

SetReturnBranch sets ReturnBranch field to given value.

### HasReturnBranch

`func (o *ReturnsDetailsResponseProductsInner) HasReturnBranch() bool`

HasReturnBranch returns a boolean if a field has been set.

### GetShipFromBranch

`func (o *ReturnsDetailsResponseProductsInner) GetShipFromBranch() string`

GetShipFromBranch returns the ShipFromBranch field if non-nil, zero value otherwise.

### GetShipFromBranchOk

`func (o *ReturnsDetailsResponseProductsInner) GetShipFromBranchOk() (*string, bool)`

GetShipFromBranchOk returns a tuple with the ShipFromBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromBranch

`func (o *ReturnsDetailsResponseProductsInner) SetShipFromBranch(v string)`

SetShipFromBranch sets ShipFromBranch field to given value.

### HasShipFromBranch

`func (o *ReturnsDetailsResponseProductsInner) HasShipFromBranch() bool`

HasShipFromBranch returns a boolean if a field has been set.

### GetRequestDetails

`func (o *ReturnsDetailsResponseProductsInner) GetRequestDetails() string`

GetRequestDetails returns the RequestDetails field if non-nil, zero value otherwise.

### GetRequestDetailsOk

`func (o *ReturnsDetailsResponseProductsInner) GetRequestDetailsOk() (*string, bool)`

GetRequestDetailsOk returns a tuple with the RequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetails

`func (o *ReturnsDetailsResponseProductsInner) SetRequestDetails(v string)`

SetRequestDetails sets RequestDetails field to given value.

### HasRequestDetails

`func (o *ReturnsDetailsResponseProductsInner) HasRequestDetails() bool`

HasRequestDetails returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *ReturnsDetailsResponseProductsInner) GetAdditionalDetails() string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ReturnsDetailsResponseProductsInner) GetAdditionalDetailsOk() (*string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ReturnsDetailsResponseProductsInner) SetAdditionalDetails(v string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ReturnsDetailsResponseProductsInner) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


