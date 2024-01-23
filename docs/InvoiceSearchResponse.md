# InvoiceSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsFound** | Pointer to **int32** | Total count of quotes retrieved in the request response. | [optional] 
**PageSize** | Pointer to **int32** | Number of records (quotes) displayed per page in the quote list. | [optional] 
**PageNumber** | Pointer to **int32** | Page index or page number for the list of quotes being returned. | [optional] 
**Invoices** | Pointer to [**[]InvoiceSearchResponseInvoicesInner**](InvoiceSearchResponseInvoicesInner.md) | The Invoices details for the requested criteria. | [optional] 
**NextPage** | Pointer to **string** | Next page of the pagination. | [optional] 

## Methods

### NewInvoiceSearchResponse

`func NewInvoiceSearchResponse() *InvoiceSearchResponse`

NewInvoiceSearchResponse instantiates a new InvoiceSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSearchResponseWithDefaults

`func NewInvoiceSearchResponseWithDefaults() *InvoiceSearchResponse`

NewInvoiceSearchResponseWithDefaults instantiates a new InvoiceSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsFound

`func (o *InvoiceSearchResponse) GetRecordsFound() int32`

GetRecordsFound returns the RecordsFound field if non-nil, zero value otherwise.

### GetRecordsFoundOk

`func (o *InvoiceSearchResponse) GetRecordsFoundOk() (*int32, bool)`

GetRecordsFoundOk returns a tuple with the RecordsFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsFound

`func (o *InvoiceSearchResponse) SetRecordsFound(v int32)`

SetRecordsFound sets RecordsFound field to given value.

### HasRecordsFound

`func (o *InvoiceSearchResponse) HasRecordsFound() bool`

HasRecordsFound returns a boolean if a field has been set.

### GetPageSize

`func (o *InvoiceSearchResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InvoiceSearchResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InvoiceSearchResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *InvoiceSearchResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *InvoiceSearchResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *InvoiceSearchResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *InvoiceSearchResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *InvoiceSearchResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetInvoices

`func (o *InvoiceSearchResponse) GetInvoices() []InvoiceSearchResponseInvoicesInner`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *InvoiceSearchResponse) GetInvoicesOk() (*[]InvoiceSearchResponseInvoicesInner, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *InvoiceSearchResponse) SetInvoices(v []InvoiceSearchResponseInvoicesInner)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *InvoiceSearchResponse) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetNextPage

`func (o *InvoiceSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *InvoiceSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *InvoiceSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *InvoiceSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


