# OrderDeleteRequestServicerequestRequestpreamble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isocountrycode** | **string** | Country that Order is being place in. | 
**CustomerNumber** | **string** | Account number order will be billed to. INGRAM MICRO ACCOUNT NUMBER REQUIRED | 

## Methods

### NewOrderDeleteRequestServicerequestRequestpreamble

`func NewOrderDeleteRequestServicerequestRequestpreamble(isocountrycode string, customerNumber string, ) *OrderDeleteRequestServicerequestRequestpreamble`

NewOrderDeleteRequestServicerequestRequestpreamble instantiates a new OrderDeleteRequestServicerequestRequestpreamble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDeleteRequestServicerequestRequestpreambleWithDefaults

`func NewOrderDeleteRequestServicerequestRequestpreambleWithDefaults() *OrderDeleteRequestServicerequestRequestpreamble`

NewOrderDeleteRequestServicerequestRequestpreambleWithDefaults instantiates a new OrderDeleteRequestServicerequestRequestpreamble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsocountrycode

`func (o *OrderDeleteRequestServicerequestRequestpreamble) GetIsocountrycode() string`

GetIsocountrycode returns the Isocountrycode field if non-nil, zero value otherwise.

### GetIsocountrycodeOk

`func (o *OrderDeleteRequestServicerequestRequestpreamble) GetIsocountrycodeOk() (*string, bool)`

GetIsocountrycodeOk returns a tuple with the Isocountrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsocountrycode

`func (o *OrderDeleteRequestServicerequestRequestpreamble) SetIsocountrycode(v string)`

SetIsocountrycode sets Isocountrycode field to given value.


### GetCustomerNumber

`func (o *OrderDeleteRequestServicerequestRequestpreamble) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *OrderDeleteRequestServicerequestRequestpreamble) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *OrderDeleteRequestServicerequestRequestpreamble) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


