# OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linetype** | Pointer to **string** | “P”-Line or SKU Number “C”-Comment Line | [optional] 
**Globallinenumber** | Pointer to **string** | Ingram generated line number | [optional] 
**Partnumber** | Pointer to **string** | Ingram Micro Sku Number | [optional] 
**Globalskuid** | Pointer to **string** |  | [optional] 
**Linenumber** | Pointer to **string** |  | [optional] 
**Carriercode** | Pointer to **string** | Transportation 2 digit codes | [optional] 
**Carrierdescription** | Pointer to **string** | Transportation Carrier Name | [optional] 
**Requestedunitprice** | Pointer to **float32** | Price requested by reseller. Price Variance can be set up by Ingram Micro Sales Rep | [optional] 
**Requestedquantity** | Pointer to **int32** | Quanity Requested | [optional] 
**Confirmedquantity** | Pointer to **int32** | Quanity Shipped | [optional] 
**Backorderedquantity** | Pointer to **int32** | Quanity of units that didn’t ship | [optional] 
**Unitproductprice** | Pointer to **float32** | Price Per Unit | [optional] 
**Netamount** | Pointer to **float32** | Total amount. Quantity X Unit Price | [optional] 
**Warehouseid** | Pointer to **string** |  | [optional] 
**Ordersuffix** | Pointer to **string** | Use order suffix with the globalorderid for this line item. | [optional] 

## Methods

### NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner

`func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner() *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner`

NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInnerWithDefaults

`func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInnerWithDefaults() *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner`

NewOrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInnerWithDefaults instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinetype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinetype() string`

GetLinetype returns the Linetype field if non-nil, zero value otherwise.

### GetLinetypeOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinetypeOk() (*string, bool)`

GetLinetypeOk returns a tuple with the Linetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinetype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetLinetype(v string)`

SetLinetype sets Linetype field to given value.

### HasLinetype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasLinetype() bool`

HasLinetype returns a boolean if a field has been set.

### GetGloballinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGloballinenumber() string`

GetGloballinenumber returns the Globallinenumber field if non-nil, zero value otherwise.

### GetGloballinenumberOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGloballinenumberOk() (*string, bool)`

GetGloballinenumberOk returns a tuple with the Globallinenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetGloballinenumber(v string)`

SetGloballinenumber sets Globallinenumber field to given value.

### HasGloballinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasGloballinenumber() bool`

HasGloballinenumber returns a boolean if a field has been set.

### GetPartnumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetPartnumber() string`

GetPartnumber returns the Partnumber field if non-nil, zero value otherwise.

### GetPartnumberOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetPartnumberOk() (*string, bool)`

GetPartnumberOk returns a tuple with the Partnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetPartnumber(v string)`

SetPartnumber sets Partnumber field to given value.

### HasPartnumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasPartnumber() bool`

HasPartnumber returns a boolean if a field has been set.

### GetGlobalskuid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGlobalskuid() string`

GetGlobalskuid returns the Globalskuid field if non-nil, zero value otherwise.

### GetGlobalskuidOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetGlobalskuidOk() (*string, bool)`

GetGlobalskuidOk returns a tuple with the Globalskuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalskuid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetGlobalskuid(v string)`

SetGlobalskuid sets Globalskuid field to given value.

### HasGlobalskuid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasGlobalskuid() bool`

HasGlobalskuid returns a boolean if a field has been set.

### GetLinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinenumber() string`

GetLinenumber returns the Linenumber field if non-nil, zero value otherwise.

### GetLinenumberOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetLinenumberOk() (*string, bool)`

GetLinenumberOk returns a tuple with the Linenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetLinenumber(v string)`

SetLinenumber sets Linenumber field to given value.

### HasLinenumber

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasLinenumber() bool`

HasLinenumber returns a boolean if a field has been set.

### GetCarriercode

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarriercode() string`

GetCarriercode returns the Carriercode field if non-nil, zero value otherwise.

### GetCarriercodeOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarriercodeOk() (*string, bool)`

GetCarriercodeOk returns a tuple with the Carriercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriercode

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetCarriercode(v string)`

SetCarriercode sets Carriercode field to given value.

### HasCarriercode

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasCarriercode() bool`

HasCarriercode returns a boolean if a field has been set.

### GetCarrierdescription

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarrierdescription() string`

GetCarrierdescription returns the Carrierdescription field if non-nil, zero value otherwise.

### GetCarrierdescriptionOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetCarrierdescriptionOk() (*string, bool)`

GetCarrierdescriptionOk returns a tuple with the Carrierdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierdescription

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetCarrierdescription(v string)`

SetCarrierdescription sets Carrierdescription field to given value.

### HasCarrierdescription

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasCarrierdescription() bool`

HasCarrierdescription returns a boolean if a field has been set.

### GetRequestedunitprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedunitprice() float32`

GetRequestedunitprice returns the Requestedunitprice field if non-nil, zero value otherwise.

### GetRequestedunitpriceOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedunitpriceOk() (*float32, bool)`

GetRequestedunitpriceOk returns a tuple with the Requestedunitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedunitprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetRequestedunitprice(v float32)`

SetRequestedunitprice sets Requestedunitprice field to given value.

### HasRequestedunitprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasRequestedunitprice() bool`

HasRequestedunitprice returns a boolean if a field has been set.

### GetRequestedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedquantity() int32`

GetRequestedquantity returns the Requestedquantity field if non-nil, zero value otherwise.

### GetRequestedquantityOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetRequestedquantityOk() (*int32, bool)`

GetRequestedquantityOk returns a tuple with the Requestedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetRequestedquantity(v int32)`

SetRequestedquantity sets Requestedquantity field to given value.

### HasRequestedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasRequestedquantity() bool`

HasRequestedquantity returns a boolean if a field has been set.

### GetConfirmedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetConfirmedquantity() int32`

GetConfirmedquantity returns the Confirmedquantity field if non-nil, zero value otherwise.

### GetConfirmedquantityOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetConfirmedquantityOk() (*int32, bool)`

GetConfirmedquantityOk returns a tuple with the Confirmedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetConfirmedquantity(v int32)`

SetConfirmedquantity sets Confirmedquantity field to given value.

### HasConfirmedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasConfirmedquantity() bool`

HasConfirmedquantity returns a boolean if a field has been set.

### GetBackorderedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetBackorderedquantity() int32`

GetBackorderedquantity returns the Backorderedquantity field if non-nil, zero value otherwise.

### GetBackorderedquantityOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetBackorderedquantityOk() (*int32, bool)`

GetBackorderedquantityOk returns a tuple with the Backorderedquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetBackorderedquantity(v int32)`

SetBackorderedquantity sets Backorderedquantity field to given value.

### HasBackorderedquantity

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasBackorderedquantity() bool`

HasBackorderedquantity returns a boolean if a field has been set.

### GetUnitproductprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetUnitproductprice() float32`

GetUnitproductprice returns the Unitproductprice field if non-nil, zero value otherwise.

### GetUnitproductpriceOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetUnitproductpriceOk() (*float32, bool)`

GetUnitproductpriceOk returns a tuple with the Unitproductprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitproductprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetUnitproductprice(v float32)`

SetUnitproductprice sets Unitproductprice field to given value.

### HasUnitproductprice

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasUnitproductprice() bool`

HasUnitproductprice returns a boolean if a field has been set.

### GetNetamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetNetamount() float32`

GetNetamount returns the Netamount field if non-nil, zero value otherwise.

### GetNetamountOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetNetamountOk() (*float32, bool)`

GetNetamountOk returns a tuple with the Netamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetNetamount(v float32)`

SetNetamount sets Netamount field to given value.

### HasNetamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasNetamount() bool`

HasNetamount returns a boolean if a field has been set.

### GetWarehouseid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetWarehouseid() string`

GetWarehouseid returns the Warehouseid field if non-nil, zero value otherwise.

### GetWarehouseidOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetWarehouseidOk() (*string, bool)`

GetWarehouseidOk returns a tuple with the Warehouseid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetWarehouseid(v string)`

SetWarehouseid sets Warehouseid field to given value.

### HasWarehouseid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasWarehouseid() bool`

HasWarehouseid returns a boolean if a field has been set.

### GetOrdersuffix

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetOrdersuffix() string`

GetOrdersuffix returns the Ordersuffix field if non-nil, zero value otherwise.

### GetOrdersuffixOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) GetOrdersuffixOk() (*string, bool)`

GetOrdersuffixOk returns a tuple with the Ordersuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersuffix

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) SetOrdersuffix(v string)`

SetOrdersuffix sets Ordersuffix field to given value.

### HasOrdersuffix

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) HasOrdersuffix() bool`

HasOrdersuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


