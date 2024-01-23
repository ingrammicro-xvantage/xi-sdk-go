# OrderCreateRequestOrdercreaterequestOrdercreatedetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customerponumber** | **string** | The customers unique Purchase Order number. Keep it unique to retrieve order information | 
**Ordertype** | **string** | Order Type - Standard orders, Direct ship orders | 
**Enduserordernumber** | Pointer to **string** | Customers End-user PO number | [optional] 
**Billtosuffix** | Pointer to **string** | Designates flooring acct to be used | [optional] 
**Shiptosuffix** | Pointer to **string** | Applies to customers with multiple ship to locations (store locations) | [optional] 
**Shiptoaddress** | [**OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress**](OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress.md) |  | 
**Carriercode** | Pointer to **string** | A customer can dictate what carrier to use for their shipment (Ingram 2-digit carrier code is required). Our recommendation is leave this field blank which will allow Ingram Micro to choose the best carrier to gain the best freight rates. | [optional] 
**Thirdpartyfreightaccountnumber** | Pointer to **string** | Refers to a third-party freight account number for charging freight against. The account number should be passed within this field and the appropriate carrier code should be supplied within the carrier code tags. Prior to sending your request containing the third-party account number, it must be first entered into our system. Your Ingram Micro Sales Representative can action this for you. If submitted within an order without this preapproval the third-party account number will be ignored.  Note: USA partners- For FedEx Air only (carrier codes F1, FO, F2, FG.), please send three leading zeros before your third-party freight account number (i.e.: 000999999999.)  | [optional] 
**Specialbidnumber** | Pointer to **string** | This is the special quote number given to a customer either by a vendor for special pricing or by Ingram Micro. To receive the special pricing assigned to this number it must be included on the order. | [optional] 
**Lines** | [**[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner**](OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner.md) |  | 
**Extendedspecs** | Pointer to [**[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsExtendedspecsInner**](OrderCreateRequestOrdercreaterequestOrdercreatedetailsExtendedspecsInner.md) |  | [optional] 

## Methods

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetails

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetails(customerponumber string, ordertype string, shiptoaddress OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress, lines []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner, ) *OrderCreateRequestOrdercreaterequestOrdercreatedetails`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetails instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsWithDefaults

`func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetails`

NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerponumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetCustomerponumber() string`

GetCustomerponumber returns the Customerponumber field if non-nil, zero value otherwise.

### GetCustomerponumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetCustomerponumberOk() (*string, bool)`

GetCustomerponumberOk returns a tuple with the Customerponumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerponumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetCustomerponumber(v string)`

SetCustomerponumber sets Customerponumber field to given value.


### GetOrdertype

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetOrdertype() string`

GetOrdertype returns the Ordertype field if non-nil, zero value otherwise.

### GetOrdertypeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetOrdertypeOk() (*string, bool)`

GetOrdertypeOk returns a tuple with the Ordertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertype

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetOrdertype(v string)`

SetOrdertype sets Ordertype field to given value.


### GetEnduserordernumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetEnduserordernumber() string`

GetEnduserordernumber returns the Enduserordernumber field if non-nil, zero value otherwise.

### GetEnduserordernumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetEnduserordernumberOk() (*string, bool)`

GetEnduserordernumberOk returns a tuple with the Enduserordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserordernumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetEnduserordernumber(v string)`

SetEnduserordernumber sets Enduserordernumber field to given value.

### HasEnduserordernumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasEnduserordernumber() bool`

HasEnduserordernumber returns a boolean if a field has been set.

### GetBilltosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetBilltosuffix() string`

GetBilltosuffix returns the Billtosuffix field if non-nil, zero value otherwise.

### GetBilltosuffixOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetBilltosuffixOk() (*string, bool)`

GetBilltosuffixOk returns a tuple with the Billtosuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilltosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetBilltosuffix(v string)`

SetBilltosuffix sets Billtosuffix field to given value.

### HasBilltosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasBilltosuffix() bool`

HasBilltosuffix returns a boolean if a field has been set.

### GetShiptosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetShiptosuffix() string`

GetShiptosuffix returns the Shiptosuffix field if non-nil, zero value otherwise.

### GetShiptosuffixOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetShiptosuffixOk() (*string, bool)`

GetShiptosuffixOk returns a tuple with the Shiptosuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiptosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetShiptosuffix(v string)`

SetShiptosuffix sets Shiptosuffix field to given value.

### HasShiptosuffix

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasShiptosuffix() bool`

HasShiptosuffix returns a boolean if a field has been set.

### GetShiptoaddress

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetShiptoaddress() OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress`

GetShiptoaddress returns the Shiptoaddress field if non-nil, zero value otherwise.

### GetShiptoaddressOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetShiptoaddressOk() (*OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress, bool)`

GetShiptoaddressOk returns a tuple with the Shiptoaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiptoaddress

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetShiptoaddress(v OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress)`

SetShiptoaddress sets Shiptoaddress field to given value.


### GetCarriercode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetCarriercode() string`

GetCarriercode returns the Carriercode field if non-nil, zero value otherwise.

### GetCarriercodeOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetCarriercodeOk() (*string, bool)`

GetCarriercodeOk returns a tuple with the Carriercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriercode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetCarriercode(v string)`

SetCarriercode sets Carriercode field to given value.

### HasCarriercode

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasCarriercode() bool`

HasCarriercode returns a boolean if a field has been set.

### GetThirdpartyfreightaccountnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetThirdpartyfreightaccountnumber() string`

GetThirdpartyfreightaccountnumber returns the Thirdpartyfreightaccountnumber field if non-nil, zero value otherwise.

### GetThirdpartyfreightaccountnumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetThirdpartyfreightaccountnumberOk() (*string, bool)`

GetThirdpartyfreightaccountnumberOk returns a tuple with the Thirdpartyfreightaccountnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdpartyfreightaccountnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetThirdpartyfreightaccountnumber(v string)`

SetThirdpartyfreightaccountnumber sets Thirdpartyfreightaccountnumber field to given value.

### HasThirdpartyfreightaccountnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasThirdpartyfreightaccountnumber() bool`

HasThirdpartyfreightaccountnumber returns a boolean if a field has been set.

### GetSpecialbidnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetSpecialbidnumber() string`

GetSpecialbidnumber returns the Specialbidnumber field if non-nil, zero value otherwise.

### GetSpecialbidnumberOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetSpecialbidnumberOk() (*string, bool)`

GetSpecialbidnumberOk returns a tuple with the Specialbidnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialbidnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetSpecialbidnumber(v string)`

SetSpecialbidnumber sets Specialbidnumber field to given value.

### HasSpecialbidnumber

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasSpecialbidnumber() bool`

HasSpecialbidnumber returns a boolean if a field has been set.

### GetLines

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetLines() []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetLinesOk() (*[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetLines(v []OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInner)`

SetLines sets Lines field to given value.


### GetExtendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetExtendedspecs() []OrderCreateRequestOrdercreaterequestOrdercreatedetailsExtendedspecsInner`

GetExtendedspecs returns the Extendedspecs field if non-nil, zero value otherwise.

### GetExtendedspecsOk

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) GetExtendedspecsOk() (*[]OrderCreateRequestOrdercreaterequestOrdercreatedetailsExtendedspecsInner, bool)`

GetExtendedspecsOk returns a tuple with the Extendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) SetExtendedspecs(v []OrderCreateRequestOrdercreaterequestOrdercreatedetailsExtendedspecsInner)`

SetExtendedspecs sets Extendedspecs field to given value.

### HasExtendedspecs

`func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetails) HasExtendedspecs() bool`

HasExtendedspecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


