# OrderCreateResponseServiceresponseOrdercreateresponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numberoflineswithsuccess** | Pointer to **string** | Number of line items that were successful | [optional] 
**Numberoflineswitherror** | Pointer to **string** | Number of line items with error | [optional] 
**Numberoflineswithwarning** | Pointer to **string** | Number of line items with warnings | [optional] 
**Globalorderid** | Pointer to **string** | Ingram sales order number | [optional] 
**Ordertype** | Pointer to **string** | S&#x3D;Stocked PO D&#x3D;Direct Ship PO | [optional] 
**Ordertimestamp** | Pointer to **string** | Time order received | [optional] 
**Invoicingsystemorderid** | Pointer to **string** | Ingram Micro generated order number | [optional] 
**Taxamount** | Pointer to **float32** |  | [optional] 
**Freightamount** | Pointer to **float32** | Freight amount customer pays for freight | [optional] 
**Orderamount** | Pointer to **float32** | Total amount of order with freight and taxes | [optional] 
**Lines** | Pointer to [**[]OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner**](OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner.md) | Collection of lines | [optional] 

## Methods

### NewOrderCreateResponseServiceresponseOrdercreateresponseInner

`func NewOrderCreateResponseServiceresponseOrdercreateresponseInner() *OrderCreateResponseServiceresponseOrdercreateresponseInner`

NewOrderCreateResponseServiceresponseOrdercreateresponseInner instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseServiceresponseOrdercreateresponseInnerWithDefaults

`func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerWithDefaults() *OrderCreateResponseServiceresponseOrdercreateresponseInner`

NewOrderCreateResponseServiceresponseOrdercreateresponseInnerWithDefaults instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberoflineswithsuccess

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithsuccess() string`

GetNumberoflineswithsuccess returns the Numberoflineswithsuccess field if non-nil, zero value otherwise.

### GetNumberoflineswithsuccessOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithsuccessOk() (*string, bool)`

GetNumberoflineswithsuccessOk returns a tuple with the Numberoflineswithsuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberoflineswithsuccess

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswithsuccess(v string)`

SetNumberoflineswithsuccess sets Numberoflineswithsuccess field to given value.

### HasNumberoflineswithsuccess

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswithsuccess() bool`

HasNumberoflineswithsuccess returns a boolean if a field has been set.

### GetNumberoflineswitherror

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswitherror() string`

GetNumberoflineswitherror returns the Numberoflineswitherror field if non-nil, zero value otherwise.

### GetNumberoflineswitherrorOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswitherrorOk() (*string, bool)`

GetNumberoflineswitherrorOk returns a tuple with the Numberoflineswitherror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberoflineswitherror

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswitherror(v string)`

SetNumberoflineswitherror sets Numberoflineswitherror field to given value.

### HasNumberoflineswitherror

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswitherror() bool`

HasNumberoflineswitherror returns a boolean if a field has been set.

### GetNumberoflineswithwarning

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithwarning() string`

GetNumberoflineswithwarning returns the Numberoflineswithwarning field if non-nil, zero value otherwise.

### GetNumberoflineswithwarningOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithwarningOk() (*string, bool)`

GetNumberoflineswithwarningOk returns a tuple with the Numberoflineswithwarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberoflineswithwarning

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswithwarning(v string)`

SetNumberoflineswithwarning sets Numberoflineswithwarning field to given value.

### HasNumberoflineswithwarning

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswithwarning() bool`

HasNumberoflineswithwarning returns a boolean if a field has been set.

### GetGlobalorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetGlobalorderid() string`

GetGlobalorderid returns the Globalorderid field if non-nil, zero value otherwise.

### GetGlobalorderidOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetGlobalorderidOk() (*string, bool)`

GetGlobalorderidOk returns a tuple with the Globalorderid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetGlobalorderid(v string)`

SetGlobalorderid sets Globalorderid field to given value.

### HasGlobalorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasGlobalorderid() bool`

HasGlobalorderid returns a boolean if a field has been set.

### GetOrdertype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertype() string`

GetOrdertype returns the Ordertype field if non-nil, zero value otherwise.

### GetOrdertypeOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertypeOk() (*string, bool)`

GetOrdertypeOk returns a tuple with the Ordertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrdertype(v string)`

SetOrdertype sets Ordertype field to given value.

### HasOrdertype

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrdertype() bool`

HasOrdertype returns a boolean if a field has been set.

### GetOrdertimestamp

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertimestamp() string`

GetOrdertimestamp returns the Ordertimestamp field if non-nil, zero value otherwise.

### GetOrdertimestampOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertimestampOk() (*string, bool)`

GetOrdertimestampOk returns a tuple with the Ordertimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertimestamp

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrdertimestamp(v string)`

SetOrdertimestamp sets Ordertimestamp field to given value.

### HasOrdertimestamp

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrdertimestamp() bool`

HasOrdertimestamp returns a boolean if a field has been set.

### GetInvoicingsystemorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetInvoicingsystemorderid() string`

GetInvoicingsystemorderid returns the Invoicingsystemorderid field if non-nil, zero value otherwise.

### GetInvoicingsystemorderidOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetInvoicingsystemorderidOk() (*string, bool)`

GetInvoicingsystemorderidOk returns a tuple with the Invoicingsystemorderid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingsystemorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetInvoicingsystemorderid(v string)`

SetInvoicingsystemorderid sets Invoicingsystemorderid field to given value.

### HasInvoicingsystemorderid

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasInvoicingsystemorderid() bool`

HasInvoicingsystemorderid returns a boolean if a field has been set.

### GetTaxamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetTaxamount() float32`

GetTaxamount returns the Taxamount field if non-nil, zero value otherwise.

### GetTaxamountOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetTaxamountOk() (*float32, bool)`

GetTaxamountOk returns a tuple with the Taxamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetTaxamount(v float32)`

SetTaxamount sets Taxamount field to given value.

### HasTaxamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasTaxamount() bool`

HasTaxamount returns a boolean if a field has been set.

### GetFreightamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetFreightamount() float32`

GetFreightamount returns the Freightamount field if non-nil, zero value otherwise.

### GetFreightamountOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetFreightamountOk() (*float32, bool)`

GetFreightamountOk returns a tuple with the Freightamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetFreightamount(v float32)`

SetFreightamount sets Freightamount field to given value.

### HasFreightamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasFreightamount() bool`

HasFreightamount returns a boolean if a field has been set.

### GetOrderamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrderamount() float32`

GetOrderamount returns the Orderamount field if non-nil, zero value otherwise.

### GetOrderamountOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrderamountOk() (*float32, bool)`

GetOrderamountOk returns a tuple with the Orderamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrderamount(v float32)`

SetOrderamount sets Orderamount field to given value.

### HasOrderamount

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrderamount() bool`

HasOrderamount returns a boolean if a field has been set.

### GetLines

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetLines() []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetLinesOk() (*[]OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetLines(v []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


