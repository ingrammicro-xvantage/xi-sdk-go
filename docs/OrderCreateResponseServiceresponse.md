# OrderCreateResponseServiceresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responsepreamble** | Pointer to [**InvoiceDetailResponseServiceresponseResponsepreamble**](InvoiceDetailResponseServiceresponseResponsepreamble.md) |  | [optional] 
**Ordersummary** | Pointer to [**OrderCreateResponseServiceresponseOrdersummary**](OrderCreateResponseServiceresponseOrdersummary.md) |  | [optional] 
**Ordercreateresponse** | Pointer to [**[]OrderCreateResponseServiceresponseOrdercreateresponseInner**](OrderCreateResponseServiceresponseOrdercreateresponseInner.md) | Collection of orders | [optional] 

## Methods

### NewOrderCreateResponseServiceresponse

`func NewOrderCreateResponseServiceresponse() *OrderCreateResponseServiceresponse`

NewOrderCreateResponseServiceresponse instantiates a new OrderCreateResponseServiceresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseServiceresponseWithDefaults

`func NewOrderCreateResponseServiceresponseWithDefaults() *OrderCreateResponseServiceresponse`

NewOrderCreateResponseServiceresponseWithDefaults instantiates a new OrderCreateResponseServiceresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponsepreamble

`func (o *OrderCreateResponseServiceresponse) GetResponsepreamble() InvoiceDetailResponseServiceresponseResponsepreamble`

GetResponsepreamble returns the Responsepreamble field if non-nil, zero value otherwise.

### GetResponsepreambleOk

`func (o *OrderCreateResponseServiceresponse) GetResponsepreambleOk() (*InvoiceDetailResponseServiceresponseResponsepreamble, bool)`

GetResponsepreambleOk returns a tuple with the Responsepreamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsepreamble

`func (o *OrderCreateResponseServiceresponse) SetResponsepreamble(v InvoiceDetailResponseServiceresponseResponsepreamble)`

SetResponsepreamble sets Responsepreamble field to given value.

### HasResponsepreamble

`func (o *OrderCreateResponseServiceresponse) HasResponsepreamble() bool`

HasResponsepreamble returns a boolean if a field has been set.

### GetOrdersummary

`func (o *OrderCreateResponseServiceresponse) GetOrdersummary() OrderCreateResponseServiceresponseOrdersummary`

GetOrdersummary returns the Ordersummary field if non-nil, zero value otherwise.

### GetOrdersummaryOk

`func (o *OrderCreateResponseServiceresponse) GetOrdersummaryOk() (*OrderCreateResponseServiceresponseOrdersummary, bool)`

GetOrdersummaryOk returns a tuple with the Ordersummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersummary

`func (o *OrderCreateResponseServiceresponse) SetOrdersummary(v OrderCreateResponseServiceresponseOrdersummary)`

SetOrdersummary sets Ordersummary field to given value.

### HasOrdersummary

`func (o *OrderCreateResponseServiceresponse) HasOrdersummary() bool`

HasOrdersummary returns a boolean if a field has been set.

### GetOrdercreateresponse

`func (o *OrderCreateResponseServiceresponse) GetOrdercreateresponse() []OrderCreateResponseServiceresponseOrdercreateresponseInner`

GetOrdercreateresponse returns the Ordercreateresponse field if non-nil, zero value otherwise.

### GetOrdercreateresponseOk

`func (o *OrderCreateResponseServiceresponse) GetOrdercreateresponseOk() (*[]OrderCreateResponseServiceresponseOrdercreateresponseInner, bool)`

GetOrdercreateresponseOk returns a tuple with the Ordercreateresponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdercreateresponse

`func (o *OrderCreateResponseServiceresponse) SetOrdercreateresponse(v []OrderCreateResponseServiceresponseOrdercreateresponseInner)`

SetOrdercreateresponse sets Ordercreateresponse field to given value.

### HasOrdercreateresponse

`func (o *OrderCreateResponseServiceresponse) HasOrdercreateresponse() bool`

HasOrdercreateresponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


