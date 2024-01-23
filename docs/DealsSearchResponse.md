# DealsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | Number of records found. | [optional] 
**PageSize** | Pointer to **int32** | Number of records in a page. | [optional] 
**PageNumber** | Pointer to **int32** | Number of page. | [optional] 
**Deals** | Pointer to [**[]DealsSearchResponseDealsInner**](DealsSearchResponseDealsInner.md) |  | [optional] 
**NextPage** | Pointer to **string** | URL for the next page. | [optional] 
**PreviousPage** | Pointer to **string** | URL for the previous page. | [optional] 

## Methods

### NewDealsSearchResponse

`func NewDealsSearchResponse() *DealsSearchResponse`

NewDealsSearchResponse instantiates a new DealsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealsSearchResponseWithDefaults

`func NewDealsSearchResponseWithDefaults() *DealsSearchResponse`

NewDealsSearchResponseWithDefaults instantiates a new DealsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *DealsSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *DealsSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *DealsSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *DealsSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *DealsSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DealsSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DealsSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DealsSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *DealsSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DealsSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DealsSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DealsSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetDeals

`func (o *DealsSearchResponse) GetDeals() []DealsSearchResponseDealsInner`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *DealsSearchResponse) GetDealsOk() (*[]DealsSearchResponseDealsInner, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *DealsSearchResponse) SetDeals(v []DealsSearchResponseDealsInner)`

SetDeals sets Deals field to given value.

### HasDeals

`func (o *DealsSearchResponse) HasDeals() bool`

HasDeals returns a boolean if a field has been set.

### GetNextPage

`func (o *DealsSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *DealsSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *DealsSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *DealsSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *DealsSearchResponse) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *DealsSearchResponse) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *DealsSearchResponse) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *DealsSearchResponse) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


