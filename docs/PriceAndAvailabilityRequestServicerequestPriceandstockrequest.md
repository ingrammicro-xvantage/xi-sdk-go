# PriceAndAvailabilityRequestServicerequestPriceandstockrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Showwarehouseavailability** | Pointer to **string** | True/false to show the availability of individual warehouses | [optional] 
**Extravailabilityflag** | Pointer to **string** | Y/N to show extra availability flag | [optional] 
**Includeallsystems** | Pointer to **bool** | Flag to indicate if the price and stock information is required for all Ingram Micro systems. | [optional] 
**Item** | Pointer to [**[]PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner**](PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityRequestServicerequestPriceandstockrequest

`func NewPriceAndAvailabilityRequestServicerequestPriceandstockrequest() *PriceAndAvailabilityRequestServicerequestPriceandstockrequest`

NewPriceAndAvailabilityRequestServicerequestPriceandstockrequest instantiates a new PriceAndAvailabilityRequestServicerequestPriceandstockrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestWithDefaults

`func NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestWithDefaults() *PriceAndAvailabilityRequestServicerequestPriceandstockrequest`

NewPriceAndAvailabilityRequestServicerequestPriceandstockrequestWithDefaults instantiates a new PriceAndAvailabilityRequestServicerequestPriceandstockrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowwarehouseavailability

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetShowwarehouseavailability() string`

GetShowwarehouseavailability returns the Showwarehouseavailability field if non-nil, zero value otherwise.

### GetShowwarehouseavailabilityOk

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetShowwarehouseavailabilityOk() (*string, bool)`

GetShowwarehouseavailabilityOk returns a tuple with the Showwarehouseavailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowwarehouseavailability

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) SetShowwarehouseavailability(v string)`

SetShowwarehouseavailability sets Showwarehouseavailability field to given value.

### HasShowwarehouseavailability

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) HasShowwarehouseavailability() bool`

HasShowwarehouseavailability returns a boolean if a field has been set.

### GetExtravailabilityflag

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetExtravailabilityflag() string`

GetExtravailabilityflag returns the Extravailabilityflag field if non-nil, zero value otherwise.

### GetExtravailabilityflagOk

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetExtravailabilityflagOk() (*string, bool)`

GetExtravailabilityflagOk returns a tuple with the Extravailabilityflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtravailabilityflag

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) SetExtravailabilityflag(v string)`

SetExtravailabilityflag sets Extravailabilityflag field to given value.

### HasExtravailabilityflag

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) HasExtravailabilityflag() bool`

HasExtravailabilityflag returns a boolean if a field has been set.

### GetIncludeallsystems

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetIncludeallsystems() bool`

GetIncludeallsystems returns the Includeallsystems field if non-nil, zero value otherwise.

### GetIncludeallsystemsOk

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetIncludeallsystemsOk() (*bool, bool)`

GetIncludeallsystemsOk returns a tuple with the Includeallsystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeallsystems

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) SetIncludeallsystems(v bool)`

SetIncludeallsystems sets Includeallsystems field to given value.

### HasIncludeallsystems

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) HasIncludeallsystems() bool`

HasIncludeallsystems returns a boolean if a field has been set.

### GetItem

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetItem() []PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) GetItemOk() (*[]PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) SetItem(v []PriceAndAvailabilityRequestServicerequestPriceandstockrequestItemInner)`

SetItem sets Item field to given value.

### HasItem

`func (o *PriceAndAvailabilityRequestServicerequestPriceandstockrequest) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


