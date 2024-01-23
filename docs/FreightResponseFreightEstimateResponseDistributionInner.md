# FreightResponseFreightEstimateResponseDistributionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipFromBranchNumber** | Pointer to **string** | The ID of the warehouse the line item will ship from. | [optional] 
**CarrierCode** | Pointer to **string** | The code for the shipping carrier for the line item. | [optional] 
**ShipVia** | Pointer to **string** | The name of the shipping carrier. | [optional] 
**FreightRate** | Pointer to **float32** | Estimated freight charge. | [optional] 
**TotalWeight** | Pointer to **float32** | Total weight. | [optional] 
**TransitDays** | Pointer to **int32** | Number of transit days. | [optional] 
**CarrierList** | Pointer to [**[]FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner**](FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner.md) |  | [optional] 

## Methods

### NewFreightResponseFreightEstimateResponseDistributionInner

`func NewFreightResponseFreightEstimateResponseDistributionInner() *FreightResponseFreightEstimateResponseDistributionInner`

NewFreightResponseFreightEstimateResponseDistributionInner instantiates a new FreightResponseFreightEstimateResponseDistributionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreightResponseFreightEstimateResponseDistributionInnerWithDefaults

`func NewFreightResponseFreightEstimateResponseDistributionInnerWithDefaults() *FreightResponseFreightEstimateResponseDistributionInner`

NewFreightResponseFreightEstimateResponseDistributionInnerWithDefaults instantiates a new FreightResponseFreightEstimateResponseDistributionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipFromBranchNumber

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipFromBranchNumber() string`

GetShipFromBranchNumber returns the ShipFromBranchNumber field if non-nil, zero value otherwise.

### GetShipFromBranchNumberOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipFromBranchNumberOk() (*string, bool)`

GetShipFromBranchNumberOk returns a tuple with the ShipFromBranchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromBranchNumber

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetShipFromBranchNumber(v string)`

SetShipFromBranchNumber sets ShipFromBranchNumber field to given value.

### HasShipFromBranchNumber

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasShipFromBranchNumber() bool`

HasShipFromBranchNumber returns a boolean if a field has been set.

### GetCarrierCode

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetShipVia

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipVia() string`

GetShipVia returns the ShipVia field if non-nil, zero value otherwise.

### GetShipViaOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipViaOk() (*string, bool)`

GetShipViaOk returns a tuple with the ShipVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipVia

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetShipVia(v string)`

SetShipVia sets ShipVia field to given value.

### HasShipVia

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasShipVia() bool`

HasShipVia returns a boolean if a field has been set.

### GetFreightRate

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetFreightRate() float32`

GetFreightRate returns the FreightRate field if non-nil, zero value otherwise.

### GetFreightRateOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetFreightRateOk() (*float32, bool)`

GetFreightRateOk returns a tuple with the FreightRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightRate

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetFreightRate(v float32)`

SetFreightRate sets FreightRate field to given value.

### HasFreightRate

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasFreightRate() bool`

HasFreightRate returns a boolean if a field has been set.

### GetTotalWeight

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTotalWeight() float32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTotalWeightOk() (*float32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetTotalWeight(v float32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetTransitDays

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTransitDays() int32`

GetTransitDays returns the TransitDays field if non-nil, zero value otherwise.

### GetTransitDaysOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTransitDaysOk() (*int32, bool)`

GetTransitDaysOk returns a tuple with the TransitDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitDays

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetTransitDays(v int32)`

SetTransitDays sets TransitDays field to given value.

### HasTransitDays

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasTransitDays() bool`

HasTransitDays returns a boolean if a field has been set.

### GetCarrierList

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierList() []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner`

GetCarrierList returns the CarrierList field if non-nil, zero value otherwise.

### GetCarrierListOk

`func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierListOk() (*[]FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner, bool)`

GetCarrierListOk returns a tuple with the CarrierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierList

`func (o *FreightResponseFreightEstimateResponseDistributionInner) SetCarrierList(v []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner)`

SetCarrierList sets CarrierList field to given value.

### HasCarrierList

`func (o *FreightResponseFreightEstimateResponseDistributionInner) HasCarrierList() bool`

HasCarrierList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


