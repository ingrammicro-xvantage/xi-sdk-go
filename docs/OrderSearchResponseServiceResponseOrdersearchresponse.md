# OrderSearchResponseServiceResponseOrdersearchresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ordersfound** | **string** | Number of records found in the search result | 
**Pagesize** | Pointer to **string** | The submitted pagesize, default is 25 | [optional] 
**Pagenumber** | Pointer to **string** | The submitted pager number, default is 1 | [optional] 
**Orders** | Pointer to [**[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner**](OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner.md) | An array of orders in the search result | [optional] 

## Methods

### NewOrderSearchResponseServiceResponseOrdersearchresponse

`func NewOrderSearchResponseServiceResponseOrdersearchresponse(ordersfound string, ) *OrderSearchResponseServiceResponseOrdersearchresponse`

NewOrderSearchResponseServiceResponseOrdersearchresponse instantiates a new OrderSearchResponseServiceResponseOrdersearchresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseServiceResponseOrdersearchresponseWithDefaults

`func NewOrderSearchResponseServiceResponseOrdersearchresponseWithDefaults() *OrderSearchResponseServiceResponseOrdersearchresponse`

NewOrderSearchResponseServiceResponseOrdersearchresponseWithDefaults instantiates a new OrderSearchResponseServiceResponseOrdersearchresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdersfound

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetOrdersfound() string`

GetOrdersfound returns the Ordersfound field if non-nil, zero value otherwise.

### GetOrdersfoundOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetOrdersfoundOk() (*string, bool)`

GetOrdersfoundOk returns a tuple with the Ordersfound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersfound

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) SetOrdersfound(v string)`

SetOrdersfound sets Ordersfound field to given value.


### GetPagesize

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetPagesize() string`

GetPagesize returns the Pagesize field if non-nil, zero value otherwise.

### GetPagesizeOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetPagesizeOk() (*string, bool)`

GetPagesizeOk returns a tuple with the Pagesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesize

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) SetPagesize(v string)`

SetPagesize sets Pagesize field to given value.

### HasPagesize

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) HasPagesize() bool`

HasPagesize returns a boolean if a field has been set.

### GetPagenumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetPagenumber() string`

GetPagenumber returns the Pagenumber field if non-nil, zero value otherwise.

### GetPagenumberOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetPagenumberOk() (*string, bool)`

GetPagenumberOk returns a tuple with the Pagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagenumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) SetPagenumber(v string)`

SetPagenumber sets Pagenumber field to given value.

### HasPagenumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) HasPagenumber() bool`

HasPagenumber returns a boolean if a field has been set.

### GetOrders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetOrders() []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) GetOrdersOk() (*[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) SetOrders(v []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


