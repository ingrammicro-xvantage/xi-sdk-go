# FreightResponseFreightEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The country-specific three-character ISO 4217 currency code used for the order. | [optional] 
**TotalFreightAmount** | Pointer to **float32** | Total freight amount. | [optional] 
**TotalTaxAmount** | Pointer to **float32** | Total tax amount. | [optional] 
**TotalFees** | Pointer to **float32** | Total fees. | [optional] 
**TotalNetAmount** | Pointer to **float32** | Total net amount. | [optional] 
**GrossAmount** | Pointer to **float32** | Gross amount. | [optional] 
**Distribution** | Pointer to [**[]FreightResponseFreightEstimateResponseDistributionInner**](FreightResponseFreightEstimateResponseDistributionInner.md) |  | [optional] 
**Lines** | Pointer to [**[]FreightResponseFreightEstimateResponseLinesInner**](FreightResponseFreightEstimateResponseLinesInner.md) |  | [optional] 

## Methods

### NewFreightResponseFreightEstimateResponse

`func NewFreightResponseFreightEstimateResponse() *FreightResponseFreightEstimateResponse`

NewFreightResponseFreightEstimateResponse instantiates a new FreightResponseFreightEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreightResponseFreightEstimateResponseWithDefaults

`func NewFreightResponseFreightEstimateResponseWithDefaults() *FreightResponseFreightEstimateResponse`

NewFreightResponseFreightEstimateResponseWithDefaults instantiates a new FreightResponseFreightEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *FreightResponseFreightEstimateResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FreightResponseFreightEstimateResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FreightResponseFreightEstimateResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FreightResponseFreightEstimateResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTotalFreightAmount

`func (o *FreightResponseFreightEstimateResponse) GetTotalFreightAmount() float32`

GetTotalFreightAmount returns the TotalFreightAmount field if non-nil, zero value otherwise.

### GetTotalFreightAmountOk

`func (o *FreightResponseFreightEstimateResponse) GetTotalFreightAmountOk() (*float32, bool)`

GetTotalFreightAmountOk returns a tuple with the TotalFreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFreightAmount

`func (o *FreightResponseFreightEstimateResponse) SetTotalFreightAmount(v float32)`

SetTotalFreightAmount sets TotalFreightAmount field to given value.

### HasTotalFreightAmount

`func (o *FreightResponseFreightEstimateResponse) HasTotalFreightAmount() bool`

HasTotalFreightAmount returns a boolean if a field has been set.

### GetTotalTaxAmount

`func (o *FreightResponseFreightEstimateResponse) GetTotalTaxAmount() float32`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *FreightResponseFreightEstimateResponse) GetTotalTaxAmountOk() (*float32, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *FreightResponseFreightEstimateResponse) SetTotalTaxAmount(v float32)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *FreightResponseFreightEstimateResponse) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### GetTotalFees

`func (o *FreightResponseFreightEstimateResponse) GetTotalFees() float32`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *FreightResponseFreightEstimateResponse) GetTotalFeesOk() (*float32, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *FreightResponseFreightEstimateResponse) SetTotalFees(v float32)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *FreightResponseFreightEstimateResponse) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.

### GetTotalNetAmount

`func (o *FreightResponseFreightEstimateResponse) GetTotalNetAmount() float32`

GetTotalNetAmount returns the TotalNetAmount field if non-nil, zero value otherwise.

### GetTotalNetAmountOk

`func (o *FreightResponseFreightEstimateResponse) GetTotalNetAmountOk() (*float32, bool)`

GetTotalNetAmountOk returns a tuple with the TotalNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAmount

`func (o *FreightResponseFreightEstimateResponse) SetTotalNetAmount(v float32)`

SetTotalNetAmount sets TotalNetAmount field to given value.

### HasTotalNetAmount

`func (o *FreightResponseFreightEstimateResponse) HasTotalNetAmount() bool`

HasTotalNetAmount returns a boolean if a field has been set.

### GetGrossAmount

`func (o *FreightResponseFreightEstimateResponse) GetGrossAmount() float32`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *FreightResponseFreightEstimateResponse) GetGrossAmountOk() (*float32, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *FreightResponseFreightEstimateResponse) SetGrossAmount(v float32)`

SetGrossAmount sets GrossAmount field to given value.

### HasGrossAmount

`func (o *FreightResponseFreightEstimateResponse) HasGrossAmount() bool`

HasGrossAmount returns a boolean if a field has been set.

### GetDistribution

`func (o *FreightResponseFreightEstimateResponse) GetDistribution() []FreightResponseFreightEstimateResponseDistributionInner`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *FreightResponseFreightEstimateResponse) GetDistributionOk() (*[]FreightResponseFreightEstimateResponseDistributionInner, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *FreightResponseFreightEstimateResponse) SetDistribution(v []FreightResponseFreightEstimateResponseDistributionInner)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *FreightResponseFreightEstimateResponse) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetLines

`func (o *FreightResponseFreightEstimateResponse) GetLines() []FreightResponseFreightEstimateResponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *FreightResponseFreightEstimateResponse) GetLinesOk() (*[]FreightResponseFreightEstimateResponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *FreightResponseFreightEstimateResponse) SetLines(v []FreightResponseFreightEstimateResponseLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *FreightResponseFreightEstimateResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


