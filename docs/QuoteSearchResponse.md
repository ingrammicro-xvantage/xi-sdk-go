# QuoteSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | Total count of quotes retrieved in the request response. | [optional] 
**PageSize** | Pointer to **int32** | Number of records (quotes) displayed per page in the quote list. | [optional] 
**PageNumber** | Pointer to **int32** | Page index or page number for the list of quotes being returned. | [optional] 
**Quotes** | Pointer to [**[]QuoteSearchResponseQuotesInner**](QuoteSearchResponseQuotesInner.md) | The quote details for the requested criteria. | [optional] 

## Methods

### NewQuoteSearchResponse

`func NewQuoteSearchResponse() *QuoteSearchResponse`

NewQuoteSearchResponse instantiates a new QuoteSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteSearchResponseWithDefaults

`func NewQuoteSearchResponseWithDefaults() *QuoteSearchResponse`

NewQuoteSearchResponseWithDefaults instantiates a new QuoteSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *QuoteSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *QuoteSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *QuoteSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *QuoteSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *QuoteSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *QuoteSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *QuoteSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *QuoteSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *QuoteSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *QuoteSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *QuoteSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *QuoteSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetQuotes

`func (o *QuoteSearchResponse) GetQuotes() []QuoteSearchResponseQuotesInner`

GetQuotes returns the Quotes field if non-nil, zero value otherwise.

### GetQuotesOk

`func (o *QuoteSearchResponse) GetQuotesOk() (*[]QuoteSearchResponseQuotesInner, bool)`

GetQuotesOk returns a tuple with the Quotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotes

`func (o *QuoteSearchResponse) SetQuotes(v []QuoteSearchResponseQuotesInner)`

SetQuotes sets Quotes field to given value.

### HasQuotes

`func (o *QuoteSearchResponse) HasQuotes() bool`

HasQuotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


