# ReturnsCreateRequestListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | **string** | The Invoice number of the order. | 
**InvoiceDate** | **string** | Date of an Invoice. | 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique line number from Ingram. | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor Part Number. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the product. | [optional] 
**Quantity** | **int32** | Return quantity of the product. | 
**PrimaryReason** | **string** | Primary reason to return the product. | 
**SecondaryReason** | **string** | Secondary reason to return the product. | 
**Notes** | Pointer to **string** | Return notes. | [optional] 
**ReferenceNumber** | Pointer to **string** | Reference number to return the product. | [optional] 
**BillToAddressId** | Pointer to **string** | Suffix used to identify billing address. | [optional] 
**ShipFromInfo** | [**[]ReturnsCreateRequestListInnerShipFromInfoInner**](ReturnsCreateRequestListInnerShipFromInfoInner.md) |  | 
**NumberOfBoxes** | **int32** | Number of boxes to return. | 

## Methods

### NewReturnsCreateRequestListInner

`func NewReturnsCreateRequestListInner(invoiceNumber string, invoiceDate string, quantity int32, primaryReason string, secondaryReason string, shipFromInfo []ReturnsCreateRequestListInnerShipFromInfoInner, numberOfBoxes int32, ) *ReturnsCreateRequestListInner`

NewReturnsCreateRequestListInner instantiates a new ReturnsCreateRequestListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsCreateRequestListInnerWithDefaults

`func NewReturnsCreateRequestListInnerWithDefaults() *ReturnsCreateRequestListInner`

NewReturnsCreateRequestListInnerWithDefaults instantiates a new ReturnsCreateRequestListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *ReturnsCreateRequestListInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ReturnsCreateRequestListInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ReturnsCreateRequestListInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetInvoiceDate

`func (o *ReturnsCreateRequestListInner) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ReturnsCreateRequestListInner) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ReturnsCreateRequestListInner) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetCustomerOrderNumber

`func (o *ReturnsCreateRequestListInner) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *ReturnsCreateRequestListInner) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *ReturnsCreateRequestListInner) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *ReturnsCreateRequestListInner) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *ReturnsCreateRequestListInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *ReturnsCreateRequestListInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *ReturnsCreateRequestListInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *ReturnsCreateRequestListInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *ReturnsCreateRequestListInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *ReturnsCreateRequestListInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *ReturnsCreateRequestListInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *ReturnsCreateRequestListInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ReturnsCreateRequestListInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ReturnsCreateRequestListInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ReturnsCreateRequestListInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ReturnsCreateRequestListInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnsCreateRequestListInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnsCreateRequestListInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnsCreateRequestListInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPrimaryReason

`func (o *ReturnsCreateRequestListInner) GetPrimaryReason() string`

GetPrimaryReason returns the PrimaryReason field if non-nil, zero value otherwise.

### GetPrimaryReasonOk

`func (o *ReturnsCreateRequestListInner) GetPrimaryReasonOk() (*string, bool)`

GetPrimaryReasonOk returns a tuple with the PrimaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryReason

`func (o *ReturnsCreateRequestListInner) SetPrimaryReason(v string)`

SetPrimaryReason sets PrimaryReason field to given value.


### GetSecondaryReason

`func (o *ReturnsCreateRequestListInner) GetSecondaryReason() string`

GetSecondaryReason returns the SecondaryReason field if non-nil, zero value otherwise.

### GetSecondaryReasonOk

`func (o *ReturnsCreateRequestListInner) GetSecondaryReasonOk() (*string, bool)`

GetSecondaryReasonOk returns a tuple with the SecondaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReason

`func (o *ReturnsCreateRequestListInner) SetSecondaryReason(v string)`

SetSecondaryReason sets SecondaryReason field to given value.


### GetNotes

`func (o *ReturnsCreateRequestListInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReturnsCreateRequestListInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReturnsCreateRequestListInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReturnsCreateRequestListInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *ReturnsCreateRequestListInner) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ReturnsCreateRequestListInner) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ReturnsCreateRequestListInner) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ReturnsCreateRequestListInner) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetBillToAddressId

`func (o *ReturnsCreateRequestListInner) GetBillToAddressId() string`

GetBillToAddressId returns the BillToAddressId field if non-nil, zero value otherwise.

### GetBillToAddressIdOk

`func (o *ReturnsCreateRequestListInner) GetBillToAddressIdOk() (*string, bool)`

GetBillToAddressIdOk returns a tuple with the BillToAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToAddressId

`func (o *ReturnsCreateRequestListInner) SetBillToAddressId(v string)`

SetBillToAddressId sets BillToAddressId field to given value.

### HasBillToAddressId

`func (o *ReturnsCreateRequestListInner) HasBillToAddressId() bool`

HasBillToAddressId returns a boolean if a field has been set.

### GetShipFromInfo

`func (o *ReturnsCreateRequestListInner) GetShipFromInfo() []ReturnsCreateRequestListInnerShipFromInfoInner`

GetShipFromInfo returns the ShipFromInfo field if non-nil, zero value otherwise.

### GetShipFromInfoOk

`func (o *ReturnsCreateRequestListInner) GetShipFromInfoOk() (*[]ReturnsCreateRequestListInnerShipFromInfoInner, bool)`

GetShipFromInfoOk returns a tuple with the ShipFromInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromInfo

`func (o *ReturnsCreateRequestListInner) SetShipFromInfo(v []ReturnsCreateRequestListInnerShipFromInfoInner)`

SetShipFromInfo sets ShipFromInfo field to given value.


### GetNumberOfBoxes

`func (o *ReturnsCreateRequestListInner) GetNumberOfBoxes() int32`

GetNumberOfBoxes returns the NumberOfBoxes field if non-nil, zero value otherwise.

### GetNumberOfBoxesOk

`func (o *ReturnsCreateRequestListInner) GetNumberOfBoxesOk() (*int32, bool)`

GetNumberOfBoxesOk returns a tuple with the NumberOfBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBoxes

`func (o *ReturnsCreateRequestListInner) SetNumberOfBoxes(v int32)`

SetNumberOfBoxes sets NumberOfBoxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


