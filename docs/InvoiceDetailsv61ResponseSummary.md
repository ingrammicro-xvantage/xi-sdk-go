# InvoiceDetailsv61ResponseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**InvoiceDetailsv61ResponseSummaryLines**](InvoiceDetailsv61ResponseSummaryLines.md) |  | [optional] 
**MiscCharges** | Pointer to [**[]InvoiceDetailsv61ResponseSummaryMiscChargesInner**](InvoiceDetailsv61ResponseSummaryMiscChargesInner.md) | Miscellaneous charges. | [optional] 
**Totals** | Pointer to [**InvoiceDetailsv61ResponseSummaryTotals**](InvoiceDetailsv61ResponseSummaryTotals.md) |  | [optional] 
**ForeignFxTotals** | Pointer to [**InvoiceDetailsv61ResponseSummaryForeignFxTotals**](InvoiceDetailsv61ResponseSummaryForeignFxTotals.md) |  | [optional] 

## Methods

### NewInvoiceDetailsv61ResponseSummary

`func NewInvoiceDetailsv61ResponseSummary() *InvoiceDetailsv61ResponseSummary`

NewInvoiceDetailsv61ResponseSummary instantiates a new InvoiceDetailsv61ResponseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailsv61ResponseSummaryWithDefaults

`func NewInvoiceDetailsv61ResponseSummaryWithDefaults() *InvoiceDetailsv61ResponseSummary`

NewInvoiceDetailsv61ResponseSummaryWithDefaults instantiates a new InvoiceDetailsv61ResponseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *InvoiceDetailsv61ResponseSummary) GetLines() InvoiceDetailsv61ResponseSummaryLines`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceDetailsv61ResponseSummary) GetLinesOk() (*InvoiceDetailsv61ResponseSummaryLines, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceDetailsv61ResponseSummary) SetLines(v InvoiceDetailsv61ResponseSummaryLines)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoiceDetailsv61ResponseSummary) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMiscCharges

`func (o *InvoiceDetailsv61ResponseSummary) GetMiscCharges() []InvoiceDetailsv61ResponseSummaryMiscChargesInner`

GetMiscCharges returns the MiscCharges field if non-nil, zero value otherwise.

### GetMiscChargesOk

`func (o *InvoiceDetailsv61ResponseSummary) GetMiscChargesOk() (*[]InvoiceDetailsv61ResponseSummaryMiscChargesInner, bool)`

GetMiscChargesOk returns a tuple with the MiscCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscCharges

`func (o *InvoiceDetailsv61ResponseSummary) SetMiscCharges(v []InvoiceDetailsv61ResponseSummaryMiscChargesInner)`

SetMiscCharges sets MiscCharges field to given value.

### HasMiscCharges

`func (o *InvoiceDetailsv61ResponseSummary) HasMiscCharges() bool`

HasMiscCharges returns a boolean if a field has been set.

### SetMiscChargesNil

`func (o *InvoiceDetailsv61ResponseSummary) SetMiscChargesNil(b bool)`

 SetMiscChargesNil sets the value for MiscCharges to be an explicit nil

### UnsetMiscCharges
`func (o *InvoiceDetailsv61ResponseSummary) UnsetMiscCharges()`

UnsetMiscCharges ensures that no value is present for MiscCharges, not even an explicit nil
### GetTotals

`func (o *InvoiceDetailsv61ResponseSummary) GetTotals() InvoiceDetailsv61ResponseSummaryTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InvoiceDetailsv61ResponseSummary) GetTotalsOk() (*InvoiceDetailsv61ResponseSummaryTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InvoiceDetailsv61ResponseSummary) SetTotals(v InvoiceDetailsv61ResponseSummaryTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *InvoiceDetailsv61ResponseSummary) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetForeignFxTotals

`func (o *InvoiceDetailsv61ResponseSummary) GetForeignFxTotals() InvoiceDetailsv61ResponseSummaryForeignFxTotals`

GetForeignFxTotals returns the ForeignFxTotals field if non-nil, zero value otherwise.

### GetForeignFxTotalsOk

`func (o *InvoiceDetailsv61ResponseSummary) GetForeignFxTotalsOk() (*InvoiceDetailsv61ResponseSummaryForeignFxTotals, bool)`

GetForeignFxTotalsOk returns a tuple with the ForeignFxTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignFxTotals

`func (o *InvoiceDetailsv61ResponseSummary) SetForeignFxTotals(v InvoiceDetailsv61ResponseSummaryForeignFxTotals)`

SetForeignFxTotals sets ForeignFxTotals field to given value.

### HasForeignFxTotals

`func (o *InvoiceDetailsv61ResponseSummary) HasForeignFxTotals() bool`

HasForeignFxTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


