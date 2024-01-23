# RenewalsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | Number of records found. | [optional] 
**PageSize** | Pointer to **int32** | Number of records in a page. | [optional] 
**PageNumber** | Pointer to **int32** | Number of page. | [optional] 
**Renewals** | Pointer to [**[]RenewalsSearchResponseRenewalsInner**](RenewalsSearchResponseRenewalsInner.md) |  | [optional] 
**NextPage** | Pointer to **string** | URL for the next page. | [optional] 
**PreviousPage** | Pointer to **string** | URL for the previous page. | [optional] 

## Methods

### NewRenewalsSearchResponse

`func NewRenewalsSearchResponse() *RenewalsSearchResponse`

NewRenewalsSearchResponse instantiates a new RenewalsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsSearchResponseWithDefaults

`func NewRenewalsSearchResponseWithDefaults() *RenewalsSearchResponse`

NewRenewalsSearchResponseWithDefaults instantiates a new RenewalsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *RenewalsSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *RenewalsSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *RenewalsSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *RenewalsSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *RenewalsSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RenewalsSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RenewalsSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *RenewalsSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *RenewalsSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *RenewalsSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *RenewalsSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *RenewalsSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetRenewals

`func (o *RenewalsSearchResponse) GetRenewals() []RenewalsSearchResponseRenewalsInner`

GetRenewals returns the Renewals field if non-nil, zero value otherwise.

### GetRenewalsOk

`func (o *RenewalsSearchResponse) GetRenewalsOk() (*[]RenewalsSearchResponseRenewalsInner, bool)`

GetRenewalsOk returns a tuple with the Renewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewals

`func (o *RenewalsSearchResponse) SetRenewals(v []RenewalsSearchResponseRenewalsInner)`

SetRenewals sets Renewals field to given value.

### HasRenewals

`func (o *RenewalsSearchResponse) HasRenewals() bool`

HasRenewals returns a boolean if a field has been set.

### GetNextPage

`func (o *RenewalsSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *RenewalsSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *RenewalsSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *RenewalsSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *RenewalsSearchResponse) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *RenewalsSearchResponse) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *RenewalsSearchResponse) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *RenewalsSearchResponse) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


