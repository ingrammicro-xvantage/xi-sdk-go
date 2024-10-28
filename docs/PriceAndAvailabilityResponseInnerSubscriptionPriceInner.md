# PriceAndAvailabilityResponseInnerSubscriptionPriceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | Id of the plan. | [optional] 
**PlanName** | Pointer to **string** | Name of the plan. | [optional] 
**PlanDescription** | Pointer to **float32** | The description of the plan. | [optional] 
**Groups** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner.md) |  | [optional] 
**BillingPeriod** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriodInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriodInner.md) |  | [optional] 
**SubscriptionPeriod** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner.md) |  | [optional] 
**Options** | Pointer to [**[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner**](PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner() *PriceAndAvailabilityResponseInnerSubscriptionPriceInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInner instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults

`func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults() *PriceAndAvailabilityResponseInnerSubscriptionPriceInner`

NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanDescription() float32`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetPlanDescriptionOk() (*float32, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetPlanDescription(v float32)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetGroups() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetGroupsOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetGroups(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetBillingPeriod() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriodInner`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetBillingPeriodOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriodInner, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetBillingPeriod(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerBillingPeriodInner)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetSubscriptionPeriod() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetSubscriptionPeriodOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetSubscriptionPeriod(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerSubscriptionPeriodInner)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetOptions() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) GetOptionsOk() (*[]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) SetOptions(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


