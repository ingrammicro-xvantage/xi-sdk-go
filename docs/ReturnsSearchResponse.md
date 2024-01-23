# ReturnsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | Number of records found. | [optional] 
**PageSize** | Pointer to **int32** | Number of records in a page. | [optional] 
**PageNumber** | Pointer to **int32** | Number of page. | [optional] 
**ReturnsClaims** | Pointer to [**[]ReturnsSearchResponseReturnsClaimsInner**](ReturnsSearchResponseReturnsClaimsInner.md) |  | [optional] 
**NextPage** | Pointer to **string** | URL for the next page. | [optional] 
**PreviousPage** | Pointer to **string** | URL for the previous page. | [optional] 

## Methods

### NewReturnsSearchResponse

`func NewReturnsSearchResponse() *ReturnsSearchResponse`

NewReturnsSearchResponse instantiates a new ReturnsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsSearchResponseWithDefaults

`func NewReturnsSearchResponseWithDefaults() *ReturnsSearchResponse`

NewReturnsSearchResponseWithDefaults instantiates a new ReturnsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *ReturnsSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *ReturnsSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *ReturnsSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *ReturnsSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *ReturnsSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReturnsSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReturnsSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReturnsSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *ReturnsSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ReturnsSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ReturnsSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ReturnsSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetReturnsClaims

`func (o *ReturnsSearchResponse) GetReturnsClaims() []ReturnsSearchResponseReturnsClaimsInner`

GetReturnsClaims returns the ReturnsClaims field if non-nil, zero value otherwise.

### GetReturnsClaimsOk

`func (o *ReturnsSearchResponse) GetReturnsClaimsOk() (*[]ReturnsSearchResponseReturnsClaimsInner, bool)`

GetReturnsClaimsOk returns a tuple with the ReturnsClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsClaims

`func (o *ReturnsSearchResponse) SetReturnsClaims(v []ReturnsSearchResponseReturnsClaimsInner)`

SetReturnsClaims sets ReturnsClaims field to given value.

### HasReturnsClaims

`func (o *ReturnsSearchResponse) HasReturnsClaims() bool`

HasReturnsClaims returns a boolean if a field has been set.

### GetNextPage

`func (o *ReturnsSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ReturnsSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ReturnsSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ReturnsSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *ReturnsSearchResponse) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *ReturnsSearchResponse) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *ReturnsSearchResponse) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *ReturnsSearchResponse) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


