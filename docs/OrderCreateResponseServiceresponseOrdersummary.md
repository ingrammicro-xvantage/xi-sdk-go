# OrderCreateResponseServiceresponseOrdersummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customerponumber** | Pointer to **string** |  | [optional] 
**Totalorderamount** | Pointer to **string** | Total of all the orders including taxes and fees | [optional] 
**Totalordercreated** | Pointer to **string** | Number of orders created, in some cases we may create more than one order. | [optional] 
**Shiptoaddress** | Pointer to [**OrderCreateResponseServiceresponseOrdersummaryShiptoaddress**](OrderCreateResponseServiceresponseOrdersummaryShiptoaddress.md) |  | [optional] 

## Methods

### NewOrderCreateResponseServiceresponseOrdersummary

`func NewOrderCreateResponseServiceresponseOrdersummary() *OrderCreateResponseServiceresponseOrdersummary`

NewOrderCreateResponseServiceresponseOrdersummary instantiates a new OrderCreateResponseServiceresponseOrdersummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseServiceresponseOrdersummaryWithDefaults

`func NewOrderCreateResponseServiceresponseOrdersummaryWithDefaults() *OrderCreateResponseServiceresponseOrdersummary`

NewOrderCreateResponseServiceresponseOrdersummaryWithDefaults instantiates a new OrderCreateResponseServiceresponseOrdersummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerponumber

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetCustomerponumber() string`

GetCustomerponumber returns the Customerponumber field if non-nil, zero value otherwise.

### GetCustomerponumberOk

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetCustomerponumberOk() (*string, bool)`

GetCustomerponumberOk returns a tuple with the Customerponumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerponumber

`func (o *OrderCreateResponseServiceresponseOrdersummary) SetCustomerponumber(v string)`

SetCustomerponumber sets Customerponumber field to given value.

### HasCustomerponumber

`func (o *OrderCreateResponseServiceresponseOrdersummary) HasCustomerponumber() bool`

HasCustomerponumber returns a boolean if a field has been set.

### GetTotalorderamount

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetTotalorderamount() string`

GetTotalorderamount returns the Totalorderamount field if non-nil, zero value otherwise.

### GetTotalorderamountOk

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetTotalorderamountOk() (*string, bool)`

GetTotalorderamountOk returns a tuple with the Totalorderamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalorderamount

`func (o *OrderCreateResponseServiceresponseOrdersummary) SetTotalorderamount(v string)`

SetTotalorderamount sets Totalorderamount field to given value.

### HasTotalorderamount

`func (o *OrderCreateResponseServiceresponseOrdersummary) HasTotalorderamount() bool`

HasTotalorderamount returns a boolean if a field has been set.

### GetTotalordercreated

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetTotalordercreated() string`

GetTotalordercreated returns the Totalordercreated field if non-nil, zero value otherwise.

### GetTotalordercreatedOk

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetTotalordercreatedOk() (*string, bool)`

GetTotalordercreatedOk returns a tuple with the Totalordercreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalordercreated

`func (o *OrderCreateResponseServiceresponseOrdersummary) SetTotalordercreated(v string)`

SetTotalordercreated sets Totalordercreated field to given value.

### HasTotalordercreated

`func (o *OrderCreateResponseServiceresponseOrdersummary) HasTotalordercreated() bool`

HasTotalordercreated returns a boolean if a field has been set.

### GetShiptoaddress

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetShiptoaddress() OrderCreateResponseServiceresponseOrdersummaryShiptoaddress`

GetShiptoaddress returns the Shiptoaddress field if non-nil, zero value otherwise.

### GetShiptoaddressOk

`func (o *OrderCreateResponseServiceresponseOrdersummary) GetShiptoaddressOk() (*OrderCreateResponseServiceresponseOrdersummaryShiptoaddress, bool)`

GetShiptoaddressOk returns a tuple with the Shiptoaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiptoaddress

`func (o *OrderCreateResponseServiceresponseOrdersummary) SetShiptoaddress(v OrderCreateResponseServiceresponseOrdersummaryShiptoaddress)`

SetShiptoaddress sets Shiptoaddress field to given value.

### HasShiptoaddress

`func (o *OrderCreateResponseServiceresponseOrdersummary) HasShiptoaddress() bool`

HasShiptoaddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


