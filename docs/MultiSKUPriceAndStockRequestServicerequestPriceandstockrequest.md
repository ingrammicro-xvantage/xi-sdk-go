# MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Showwarehouseavailability** | Pointer to **string** | True/false to show the availability of individual warehouses | [optional] 
**Extravailabilityflag** | Pointer to **string** | Y/N to show extra availability flag | [optional] 
**Includeallsystems** | Pointer to **bool** | Flag to indicate if the price and stock information is required for all Ingram Micro systems. If it is set to true, the price and stock details will be returned from all Ingram Micro systems and if false, the price and stock will have returned from the system where the reseller number is set up in. | [optional] 
**Item** | Pointer to [**MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem**](MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem.md) |  | [optional] 

## Methods

### NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest

`func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest`

NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestWithDefaults

`func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestWithDefaults() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest`

NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestWithDefaults instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowwarehouseavailability

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetShowwarehouseavailability() string`

GetShowwarehouseavailability returns the Showwarehouseavailability field if non-nil, zero value otherwise.

### GetShowwarehouseavailabilityOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetShowwarehouseavailabilityOk() (*string, bool)`

GetShowwarehouseavailabilityOk returns a tuple with the Showwarehouseavailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowwarehouseavailability

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetShowwarehouseavailability(v string)`

SetShowwarehouseavailability sets Showwarehouseavailability field to given value.

### HasShowwarehouseavailability

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasShowwarehouseavailability() bool`

HasShowwarehouseavailability returns a boolean if a field has been set.

### GetExtravailabilityflag

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetExtravailabilityflag() string`

GetExtravailabilityflag returns the Extravailabilityflag field if non-nil, zero value otherwise.

### GetExtravailabilityflagOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetExtravailabilityflagOk() (*string, bool)`

GetExtravailabilityflagOk returns a tuple with the Extravailabilityflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtravailabilityflag

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetExtravailabilityflag(v string)`

SetExtravailabilityflag sets Extravailabilityflag field to given value.

### HasExtravailabilityflag

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasExtravailabilityflag() bool`

HasExtravailabilityflag returns a boolean if a field has been set.

### GetIncludeallsystems

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetIncludeallsystems() bool`

GetIncludeallsystems returns the Includeallsystems field if non-nil, zero value otherwise.

### GetIncludeallsystemsOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetIncludeallsystemsOk() (*bool, bool)`

GetIncludeallsystemsOk returns a tuple with the Includeallsystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeallsystems

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetIncludeallsystems(v bool)`

SetIncludeallsystems sets Includeallsystems field to given value.

### HasIncludeallsystems

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasIncludeallsystems() bool`

HasIncludeallsystems returns a boolean if a field has been set.

### GetItem

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetItem() MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetItemOk() (*MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetItem(v MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


