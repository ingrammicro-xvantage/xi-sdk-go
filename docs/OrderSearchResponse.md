# OrderSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | No of recourds found for the search. | [optional] 
**PageSize** | Pointer to **int32** | No of results per page.(default is 25) | [optional] 
**PageNumber** | Pointer to **int32** | Current page number.(default is 1) | [optional] 
**Orders** | Pointer to [**[]OrderSearchResponseOrdersInner**](OrderSearchResponseOrdersInner.md) | The details for the order. | [optional] 
**NextPage** | Pointer to **string** | link/URL for accessing next page. | [optional] 
**PreviousPage** | Pointer to **string** | link/URL for accessing previous page. | [optional] 

## Methods

### NewOrderSearchResponse

`func NewOrderSearchResponse() *OrderSearchResponse`

NewOrderSearchResponse instantiates a new OrderSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseWithDefaults

`func NewOrderSearchResponseWithDefaults() *OrderSearchResponse`

NewOrderSearchResponseWithDefaults instantiates a new OrderSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *OrderSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *OrderSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *OrderSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *OrderSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *OrderSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrderSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrderSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OrderSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *OrderSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OrderSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OrderSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OrderSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetOrders

`func (o *OrderSearchResponse) GetOrders() []OrderSearchResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrderSearchResponse) GetOrdersOk() (*[]OrderSearchResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrderSearchResponse) SetOrders(v []OrderSearchResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrderSearchResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetNextPage

`func (o *OrderSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *OrderSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *OrderSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *OrderSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *OrderSearchResponse) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *OrderSearchResponse) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *OrderSearchResponse) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *OrderSearchResponse) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


