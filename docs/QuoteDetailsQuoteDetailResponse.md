# QuoteDetailsQuoteDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponsePreamble** | Pointer to [**QuoteDetailsQuoteDetailResponseResponsePreamble**](QuoteDetailsQuoteDetailResponseResponsePreamble.md) |  | [optional] 
**RetrieveQuoteResponse** | Pointer to [**QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse**](QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse.md) |  | [optional] 
**QuoteProductList** | Pointer to [**[]QuoteProductList**](QuoteProductList.md) |  | [optional] 
**TotalQuoteProductCount** | Pointer to **string** |  | [optional] 
**TotalExtendedMsrp** | Pointer to **string** |  | [optional] 
**TotalQuantity** | Pointer to **int32** |  | [optional] 
**TotalExtendedQuotePrice** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteDetailsQuoteDetailResponse

`func NewQuoteDetailsQuoteDetailResponse() *QuoteDetailsQuoteDetailResponse`

NewQuoteDetailsQuoteDetailResponse instantiates a new QuoteDetailsQuoteDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsQuoteDetailResponseWithDefaults

`func NewQuoteDetailsQuoteDetailResponseWithDefaults() *QuoteDetailsQuoteDetailResponse`

NewQuoteDetailsQuoteDetailResponseWithDefaults instantiates a new QuoteDetailsQuoteDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponsePreamble

`func (o *QuoteDetailsQuoteDetailResponse) GetResponsePreamble() QuoteDetailsQuoteDetailResponseResponsePreamble`

GetResponsePreamble returns the ResponsePreamble field if non-nil, zero value otherwise.

### GetResponsePreambleOk

`func (o *QuoteDetailsQuoteDetailResponse) GetResponsePreambleOk() (*QuoteDetailsQuoteDetailResponseResponsePreamble, bool)`

GetResponsePreambleOk returns a tuple with the ResponsePreamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePreamble

`func (o *QuoteDetailsQuoteDetailResponse) SetResponsePreamble(v QuoteDetailsQuoteDetailResponseResponsePreamble)`

SetResponsePreamble sets ResponsePreamble field to given value.

### HasResponsePreamble

`func (o *QuoteDetailsQuoteDetailResponse) HasResponsePreamble() bool`

HasResponsePreamble returns a boolean if a field has been set.

### GetRetrieveQuoteResponse

`func (o *QuoteDetailsQuoteDetailResponse) GetRetrieveQuoteResponse() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse`

GetRetrieveQuoteResponse returns the RetrieveQuoteResponse field if non-nil, zero value otherwise.

### GetRetrieveQuoteResponseOk

`func (o *QuoteDetailsQuoteDetailResponse) GetRetrieveQuoteResponseOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse, bool)`

GetRetrieveQuoteResponseOk returns a tuple with the RetrieveQuoteResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveQuoteResponse

`func (o *QuoteDetailsQuoteDetailResponse) SetRetrieveQuoteResponse(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse)`

SetRetrieveQuoteResponse sets RetrieveQuoteResponse field to given value.

### HasRetrieveQuoteResponse

`func (o *QuoteDetailsQuoteDetailResponse) HasRetrieveQuoteResponse() bool`

HasRetrieveQuoteResponse returns a boolean if a field has been set.

### GetQuoteProductList

`func (o *QuoteDetailsQuoteDetailResponse) GetQuoteProductList() []QuoteProductList`

GetQuoteProductList returns the QuoteProductList field if non-nil, zero value otherwise.

### GetQuoteProductListOk

`func (o *QuoteDetailsQuoteDetailResponse) GetQuoteProductListOk() (*[]QuoteProductList, bool)`

GetQuoteProductListOk returns a tuple with the QuoteProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductList

`func (o *QuoteDetailsQuoteDetailResponse) SetQuoteProductList(v []QuoteProductList)`

SetQuoteProductList sets QuoteProductList field to given value.

### HasQuoteProductList

`func (o *QuoteDetailsQuoteDetailResponse) HasQuoteProductList() bool`

HasQuoteProductList returns a boolean if a field has been set.

### GetTotalQuoteProductCount

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuoteProductCount() string`

GetTotalQuoteProductCount returns the TotalQuoteProductCount field if non-nil, zero value otherwise.

### GetTotalQuoteProductCountOk

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuoteProductCountOk() (*string, bool)`

GetTotalQuoteProductCountOk returns a tuple with the TotalQuoteProductCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuoteProductCount

`func (o *QuoteDetailsQuoteDetailResponse) SetTotalQuoteProductCount(v string)`

SetTotalQuoteProductCount sets TotalQuoteProductCount field to given value.

### HasTotalQuoteProductCount

`func (o *QuoteDetailsQuoteDetailResponse) HasTotalQuoteProductCount() bool`

HasTotalQuoteProductCount returns a boolean if a field has been set.

### GetTotalExtendedMsrp

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedMsrp() string`

GetTotalExtendedMsrp returns the TotalExtendedMsrp field if non-nil, zero value otherwise.

### GetTotalExtendedMsrpOk

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedMsrpOk() (*string, bool)`

GetTotalExtendedMsrpOk returns a tuple with the TotalExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExtendedMsrp

`func (o *QuoteDetailsQuoteDetailResponse) SetTotalExtendedMsrp(v string)`

SetTotalExtendedMsrp sets TotalExtendedMsrp field to given value.

### HasTotalExtendedMsrp

`func (o *QuoteDetailsQuoteDetailResponse) HasTotalExtendedMsrp() bool`

HasTotalExtendedMsrp returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuantity() int32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuantityOk() (*int32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *QuoteDetailsQuoteDetailResponse) SetTotalQuantity(v int32)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *QuoteDetailsQuoteDetailResponse) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetTotalExtendedQuotePrice

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedQuotePrice() string`

GetTotalExtendedQuotePrice returns the TotalExtendedQuotePrice field if non-nil, zero value otherwise.

### GetTotalExtendedQuotePriceOk

`func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedQuotePriceOk() (*string, bool)`

GetTotalExtendedQuotePriceOk returns a tuple with the TotalExtendedQuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExtendedQuotePrice

`func (o *QuoteDetailsQuoteDetailResponse) SetTotalExtendedQuotePrice(v string)`

SetTotalExtendedQuotePrice sets TotalExtendedQuotePrice field to given value.

### HasTotalExtendedQuotePrice

`func (o *QuoteDetailsQuoteDetailResponse) HasTotalExtendedQuotePrice() bool`

HasTotalExtendedQuotePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


