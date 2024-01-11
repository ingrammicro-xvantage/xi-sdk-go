# QuoteListResponseQuoteSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponsePreamble** | Pointer to [**QuoteListResponseQuoteSearchResponseResponsePreamble**](QuoteListResponseQuoteSearchResponseResponsePreamble.md) |  | [optional] 
**QuoteList** | Pointer to [**[]QuoteListResponseQuoteSearchResponseQuoteListInner**](QuoteListResponseQuoteSearchResponseQuoteListInner.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | Total count of quotes retrieved in the request response. | [optional] 

## Methods

### NewQuoteListResponseQuoteSearchResponse

`func NewQuoteListResponseQuoteSearchResponse() *QuoteListResponseQuoteSearchResponse`

NewQuoteListResponseQuoteSearchResponse instantiates a new QuoteListResponseQuoteSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteListResponseQuoteSearchResponseWithDefaults

`func NewQuoteListResponseQuoteSearchResponseWithDefaults() *QuoteListResponseQuoteSearchResponse`

NewQuoteListResponseQuoteSearchResponseWithDefaults instantiates a new QuoteListResponseQuoteSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponsePreamble

`func (o *QuoteListResponseQuoteSearchResponse) GetResponsePreamble() QuoteListResponseQuoteSearchResponseResponsePreamble`

GetResponsePreamble returns the ResponsePreamble field if non-nil, zero value otherwise.

### GetResponsePreambleOk

`func (o *QuoteListResponseQuoteSearchResponse) GetResponsePreambleOk() (*QuoteListResponseQuoteSearchResponseResponsePreamble, bool)`

GetResponsePreambleOk returns a tuple with the ResponsePreamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePreamble

`func (o *QuoteListResponseQuoteSearchResponse) SetResponsePreamble(v QuoteListResponseQuoteSearchResponseResponsePreamble)`

SetResponsePreamble sets ResponsePreamble field to given value.

### HasResponsePreamble

`func (o *QuoteListResponseQuoteSearchResponse) HasResponsePreamble() bool`

HasResponsePreamble returns a boolean if a field has been set.

### GetQuoteList

`func (o *QuoteListResponseQuoteSearchResponse) GetQuoteList() []QuoteListResponseQuoteSearchResponseQuoteListInner`

GetQuoteList returns the QuoteList field if non-nil, zero value otherwise.

### GetQuoteListOk

`func (o *QuoteListResponseQuoteSearchResponse) GetQuoteListOk() (*[]QuoteListResponseQuoteSearchResponseQuoteListInner, bool)`

GetQuoteListOk returns a tuple with the QuoteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteList

`func (o *QuoteListResponseQuoteSearchResponse) SetQuoteList(v []QuoteListResponseQuoteSearchResponseQuoteListInner)`

SetQuoteList sets QuoteList field to given value.

### HasQuoteList

`func (o *QuoteListResponseQuoteSearchResponse) HasQuoteList() bool`

HasQuoteList returns a boolean if a field has been set.

### GetTotalCount

`func (o *QuoteListResponseQuoteSearchResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QuoteListResponseQuoteSearchResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QuoteListResponseQuoteSearchResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QuoteListResponseQuoteSearchResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


