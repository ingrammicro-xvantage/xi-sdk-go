# OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linetype** | Pointer to **string** | Values are “P” for product or “C” for comments. This can be left blank when ordering product and a “P” will be assumed.  If you are adding a COMMENT, then this value must be “C”.  Extended spec for comments:   Attribute Name: “commenttext” Attribute Value: “thank you for the order”  To make the comment invisible to the packing slip place “///” in front of the comment in the Attribute Value field.  This will allow the Ingram sales rep to see the comment on the order but will not forward on to shipping documents. | [optional] 
**Linenumber** | Pointer to **string** | This is used when a partner wants to use their own line number. Can be left blank. | [optional] 
**Ingrampartnumber** | Pointer to **string** | This is the Ingram sku number to be used for placing an order. | [optional] 
**Quantity** | **string** | The quantity that is to be ordered. | 
**Vendorpartnumber** | Pointer to **string** | The Manufacturer part number. Can be used to place an order instead of the Ingram sku.  If there are multiple Ingram part numbers to one vendor part number.  The order will be rejected. | [optional] 
**Customerpartnumber** | Pointer to **string** | This is the Customers unique part numbers that must be crossed referenced to the Ingram Micro Sku before it can be used.  Please contact your sales rep for additional information on how to set this up. | [optional] 
**UPCCode** | Pointer to **string** |  | [optional] 
**Warehouseid** | Pointer to **string** |  | [optional] 
**Unitprice** | Pointer to **string** | This is a requested price from the customer. Pre-approval is necessary before using this feature.  A methodology called price variance to manage requested pricing needs to be setup in advance by your sales rep.  If unit price is provided without this advanced setup the unit price will be ignored and standard Ingram Micro pricing will apply. | [optional] 
**Enduser** | Pointer to [**OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser**](OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser.md) |  | [optional] 
**Productextendedspecs** | Pointer to [**[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner**](OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner.md) |  | [optional] 

## Methods

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner(quantity string, ) *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerWithDefaults

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinetype

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinetype() string`

GetLinetype returns the Linetype field if non-nil, zero value otherwise.

### GetLinetypeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinetypeOk() (*string, bool)`

GetLinetypeOk returns a tuple with the Linetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinetype

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetLinetype(v string)`

SetLinetype sets Linetype field to given value.

### HasLinetype

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasLinetype() bool`

HasLinetype returns a boolean if a field has been set.

### GetLinenumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinenumber() string`

GetLinenumber returns the Linenumber field if non-nil, zero value otherwise.

### GetLinenumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetLinenumberOk() (*string, bool)`

GetLinenumberOk returns a tuple with the Linenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinenumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetLinenumber(v string)`

SetLinenumber sets Linenumber field to given value.

### HasLinenumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasLinenumber() bool`

HasLinenumber returns a boolean if a field has been set.

### GetIngrampartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetIngrampartnumber() string`

GetIngrampartnumber returns the Ingrampartnumber field if non-nil, zero value otherwise.

### GetIngrampartnumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetIngrampartnumberOk() (*string, bool)`

GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngrampartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetIngrampartnumber(v string)`

SetIngrampartnumber sets Ingrampartnumber field to given value.

### HasIngrampartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasIngrampartnumber() bool`

HasIngrampartnumber returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetVendorpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetVendorpartnumber() string`

GetVendorpartnumber returns the Vendorpartnumber field if non-nil, zero value otherwise.

### GetVendorpartnumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetVendorpartnumberOk() (*string, bool)`

GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetVendorpartnumber(v string)`

SetVendorpartnumber sets Vendorpartnumber field to given value.

### HasVendorpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasVendorpartnumber() bool`

HasVendorpartnumber returns a boolean if a field has been set.

### GetCustomerpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetCustomerpartnumber() string`

GetCustomerpartnumber returns the Customerpartnumber field if non-nil, zero value otherwise.

### GetCustomerpartnumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetCustomerpartnumberOk() (*string, bool)`

GetCustomerpartnumberOk returns a tuple with the Customerpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetCustomerpartnumber(v string)`

SetCustomerpartnumber sets Customerpartnumber field to given value.

### HasCustomerpartnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasCustomerpartnumber() bool`

HasCustomerpartnumber returns a boolean if a field has been set.

### GetUPCCode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUPCCode() string`

GetUPCCode returns the UPCCode field if non-nil, zero value otherwise.

### GetUPCCodeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUPCCodeOk() (*string, bool)`

GetUPCCodeOk returns a tuple with the UPCCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPCCode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetUPCCode(v string)`

SetUPCCode sets UPCCode field to given value.

### HasUPCCode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasUPCCode() bool`

HasUPCCode returns a boolean if a field has been set.

### GetWarehouseid

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetWarehouseid() string`

GetWarehouseid returns the Warehouseid field if non-nil, zero value otherwise.

### GetWarehouseidOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetWarehouseidOk() (*string, bool)`

GetWarehouseidOk returns a tuple with the Warehouseid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseid

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetWarehouseid(v string)`

SetWarehouseid sets Warehouseid field to given value.

### HasWarehouseid

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasWarehouseid() bool`

HasWarehouseid returns a boolean if a field has been set.

### GetUnitprice

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUnitprice() string`

GetUnitprice returns the Unitprice field if non-nil, zero value otherwise.

### GetUnitpriceOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetUnitpriceOk() (*string, bool)`

GetUnitpriceOk returns a tuple with the Unitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitprice

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetUnitprice(v string)`

SetUnitprice sets Unitprice field to given value.

### HasUnitprice

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasUnitprice() bool`

HasUnitprice returns a boolean if a field has been set.

### GetEnduser

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetEnduser() OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser`

GetEnduser returns the Enduser field if non-nil, zero value otherwise.

### GetEnduserOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetEnduserOk() (*OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser, bool)`

GetEnduserOk returns a tuple with the Enduser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduser

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetEnduser(v OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerEnduser)`

SetEnduser sets Enduser field to given value.

### HasEnduser

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasEnduser() bool`

HasEnduser returns a boolean if a field has been set.

### GetProductextendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetProductextendedspecs() []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner`

GetProductextendedspecs returns the Productextendedspecs field if non-nil, zero value otherwise.

### GetProductextendedspecsOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) GetProductextendedspecsOk() (*[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner, bool)`

GetProductextendedspecsOk returns a tuple with the Productextendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductextendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) SetProductextendedspecs(v []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner)`

SetProductextendedspecs sets Productextendedspecs field to given value.

### HasProductextendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner) HasProductextendedspecs() bool`

HasProductextendedspecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


