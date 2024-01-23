# OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attention** | Pointer to **string** | Customer contact name | [optional] 
**Addressline1** | **string** | Company Name or person to deliver. *If there isn’t an attention line please add the company name on address line 1.   UPS and FedEx will create surcharges if address line 1 contains a physical address. | 
**Addressline2** | **string** | Street address for delivery | 
**Addressline3** | Pointer to **string** | Continuation of address line 2 | [optional] 
**City** | **string** | Ship to city | 
**State** | **string** | Ship to State or Region | 
**Postalcode** | **string** | Ship to Zip code or Postal code | 
**Countrycode** | Pointer to **string** | Ship to country | [optional] 

## Methods

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress(addressline1 string, addressline2 string, city string, state string, postalcode string, ) *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddressWithDefaults

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddressWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddressWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttention

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAttention() string`

GetAttention returns the Attention field if non-nil, zero value otherwise.

### GetAttentionOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAttentionOk() (*string, bool)`

GetAttentionOk returns a tuple with the Attention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttention

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAttention(v string)`

SetAttention sets Attention field to given value.

### HasAttention

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasAttention() bool`

HasAttention returns a boolean if a field has been set.

### GetAddressline1

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline1() string`

GetAddressline1 returns the Addressline1 field if non-nil, zero value otherwise.

### GetAddressline1Ok

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline1Ok() (*string, bool)`

GetAddressline1Ok returns a tuple with the Addressline1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressline1

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline1(v string)`

SetAddressline1 sets Addressline1 field to given value.


### GetAddressline2

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline2() string`

GetAddressline2 returns the Addressline2 field if non-nil, zero value otherwise.

### GetAddressline2Ok

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline2Ok() (*string, bool)`

GetAddressline2Ok returns a tuple with the Addressline2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressline2

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline2(v string)`

SetAddressline2 sets Addressline2 field to given value.


### GetAddressline3

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline3() string`

GetAddressline3 returns the Addressline3 field if non-nil, zero value otherwise.

### GetAddressline3Ok

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline3Ok() (*string, bool)`

GetAddressline3Ok returns a tuple with the Addressline3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressline3

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline3(v string)`

SetAddressline3 sets Addressline3 field to given value.

### HasAddressline3

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasAddressline3() bool`

HasAddressline3 returns a boolean if a field has been set.

### GetCity

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetState(v string)`

SetState sets State field to given value.


### GetPostalcode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetPostalcode() string`

GetPostalcode returns the Postalcode field if non-nil, zero value otherwise.

### GetPostalcodeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetPostalcodeOk() (*string, bool)`

GetPostalcodeOk returns a tuple with the Postalcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalcode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetPostalcode(v string)`

SetPostalcode sets Postalcode field to given value.


### GetCountrycode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCountrycode() string`

GetCountrycode returns the Countrycode field if non-nil, zero value otherwise.

### GetCountrycodeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCountrycodeOk() (*string, bool)`

GetCountrycodeOk returns a tuple with the Countrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrycode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetCountrycode(v string)`

SetCountrycode sets Countrycode field to given value.

### HasCountrycode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasCountrycode() bool`

HasCountrycode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


