# InvoiceDetailResponseServiceresponseInvoicedetailresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customernumber** | Pointer to **string** |  | [optional] 
**Invoicenumber** | Pointer to **string** |  | [optional] 
**Invoicedate** | Pointer to **string** |  | [optional] 
**Invoicetype** | Pointer to **string** |  | [optional] 
**Customerordernumber** | Pointer to **string** |  | [optional] 
**Customerfreightamount** | Pointer to **float64** |  | [optional] 
**Customerforeignfrightamt** | Pointer to **float32** |  | [optional] 
**Totaltaxamount** | Pointer to **float64** |  | [optional] 
**Totalamount** | Pointer to **float64** |  | [optional] 
**Shiptosuffix** | Pointer to **string** |  | [optional] 
**Billtosuffix** | Pointer to **string** |  | [optional] 
**Freightamount** | Pointer to **float64** | May not be available in all countries | [optional] 
**Paymentterms** | Pointer to **string** |  | [optional] 
**Orderdate** | Pointer to **string** |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**Carrierdescription** | Pointer to **string** |  | [optional] 
**Discountamount** | Pointer to **float64** |  | [optional] 
**Taxtype** | Pointer to **string** |  | [optional] 
**Enduserponumber** | Pointer to **string** |  | [optional] 
**Freightforwardercode** | Pointer to **string** |  | [optional] 
**Creditmemoreasoncode** | Pointer to **string** |  | [optional] 
**Fulfillmentflag** | Pointer to **string** |  | [optional] 
**Holdreason** | Pointer to **string** |  | [optional] 
**Shipcomplete** | Pointer to **string** |  | [optional] 
**Shipdate** | Pointer to **string** |  | [optional] 
**Companycurrency** | Pointer to **string** |  | [optional] 
**Currencycode** | Pointer to **string** |  | [optional] 
**Currencyrate** | Pointer to **string** |  | [optional] 
**Globalorderid** | Pointer to **string** |  | [optional] 
**Originalshipcode** | Pointer to **string** |  | [optional] 
**Ordertype** | Pointer to **string** |  | [optional] 
**Orderstatus** | Pointer to **string** |  | [optional] 
**Totalotherfees** | Pointer to **float32** |  | [optional] 
**Totalsales** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **string** |  | [optional] 
**Shippableswitch** | Pointer to **string** |  | [optional] 
**Soldto** | Pointer to [**AddressType**](AddressType.md) |  | [optional] 
**Billto** | Pointer to [**AddressType**](AddressType.md) |  | [optional] 
**Shoptoaddress** | Pointer to [**AddressType**](AddressType.md) |  | [optional] 
**Lines** | Pointer to [**[]ProductLineType**](ProductLineType.md) |  | [optional] 
**Extendedspecs** | Pointer to [**[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner**](InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner.md) |  | [optional] 
**Miscfeeline** | Pointer to [**[]InvoiceDetailResponseServiceresponseInvoicedetailresponseMiscfeelineInner**](InvoiceDetailResponseServiceresponseInvoicedetailresponseMiscfeelineInner.md) |  | [optional] 

## Methods

### NewInvoiceDetailResponseServiceresponseInvoicedetailresponse

`func NewInvoiceDetailResponseServiceresponseInvoicedetailresponse() *InvoiceDetailResponseServiceresponseInvoicedetailresponse`

NewInvoiceDetailResponseServiceresponseInvoicedetailresponse instantiates a new InvoiceDetailResponseServiceresponseInvoicedetailresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailResponseServiceresponseInvoicedetailresponseWithDefaults

`func NewInvoiceDetailResponseServiceresponseInvoicedetailresponseWithDefaults() *InvoiceDetailResponseServiceresponseInvoicedetailresponse`

NewInvoiceDetailResponseServiceresponseInvoicedetailresponseWithDefaults instantiates a new InvoiceDetailResponseServiceresponseInvoicedetailresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomernumber() string`

GetCustomernumber returns the Customernumber field if non-nil, zero value otherwise.

### GetCustomernumberOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomernumberOk() (*string, bool)`

GetCustomernumberOk returns a tuple with the Customernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCustomernumber(v string)`

SetCustomernumber sets Customernumber field to given value.

### HasCustomernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCustomernumber() bool`

HasCustomernumber returns a boolean if a field has been set.

### GetInvoicenumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicenumber() string`

GetInvoicenumber returns the Invoicenumber field if non-nil, zero value otherwise.

### GetInvoicenumberOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicenumberOk() (*string, bool)`

GetInvoicenumberOk returns a tuple with the Invoicenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicenumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetInvoicenumber(v string)`

SetInvoicenumber sets Invoicenumber field to given value.

### HasInvoicenumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasInvoicenumber() bool`

HasInvoicenumber returns a boolean if a field has been set.

### GetInvoicedate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicedate() string`

GetInvoicedate returns the Invoicedate field if non-nil, zero value otherwise.

### GetInvoicedateOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicedateOk() (*string, bool)`

GetInvoicedateOk returns a tuple with the Invoicedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetInvoicedate(v string)`

SetInvoicedate sets Invoicedate field to given value.

### HasInvoicedate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasInvoicedate() bool`

HasInvoicedate returns a boolean if a field has been set.

### GetInvoicetype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicetype() string`

GetInvoicetype returns the Invoicetype field if non-nil, zero value otherwise.

### GetInvoicetypeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetInvoicetypeOk() (*string, bool)`

GetInvoicetypeOk returns a tuple with the Invoicetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicetype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetInvoicetype(v string)`

SetInvoicetype sets Invoicetype field to given value.

### HasInvoicetype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasInvoicetype() bool`

HasInvoicetype returns a boolean if a field has been set.

### GetCustomerordernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerordernumber() string`

GetCustomerordernumber returns the Customerordernumber field if non-nil, zero value otherwise.

### GetCustomerordernumberOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerordernumberOk() (*string, bool)`

GetCustomerordernumberOk returns a tuple with the Customerordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerordernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCustomerordernumber(v string)`

SetCustomerordernumber sets Customerordernumber field to given value.

### HasCustomerordernumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCustomerordernumber() bool`

HasCustomerordernumber returns a boolean if a field has been set.

### GetCustomerfreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerfreightamount() float64`

GetCustomerfreightamount returns the Customerfreightamount field if non-nil, zero value otherwise.

### GetCustomerfreightamountOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerfreightamountOk() (*float64, bool)`

GetCustomerfreightamountOk returns a tuple with the Customerfreightamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerfreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCustomerfreightamount(v float64)`

SetCustomerfreightamount sets Customerfreightamount field to given value.

### HasCustomerfreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCustomerfreightamount() bool`

HasCustomerfreightamount returns a boolean if a field has been set.

### GetCustomerforeignfrightamt

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerforeignfrightamt() float32`

GetCustomerforeignfrightamt returns the Customerforeignfrightamt field if non-nil, zero value otherwise.

### GetCustomerforeignfrightamtOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCustomerforeignfrightamtOk() (*float32, bool)`

GetCustomerforeignfrightamtOk returns a tuple with the Customerforeignfrightamt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerforeignfrightamt

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCustomerforeignfrightamt(v float32)`

SetCustomerforeignfrightamt sets Customerforeignfrightamt field to given value.

### HasCustomerforeignfrightamt

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCustomerforeignfrightamt() bool`

HasCustomerforeignfrightamt returns a boolean if a field has been set.

### GetTotaltaxamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotaltaxamount() float64`

GetTotaltaxamount returns the Totaltaxamount field if non-nil, zero value otherwise.

### GetTotaltaxamountOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotaltaxamountOk() (*float64, bool)`

GetTotaltaxamountOk returns a tuple with the Totaltaxamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotaltaxamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetTotaltaxamount(v float64)`

SetTotaltaxamount sets Totaltaxamount field to given value.

### HasTotaltaxamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasTotaltaxamount() bool`

HasTotaltaxamount returns a boolean if a field has been set.

### GetTotalamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalamount() float64`

GetTotalamount returns the Totalamount field if non-nil, zero value otherwise.

### GetTotalamountOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalamountOk() (*float64, bool)`

GetTotalamountOk returns a tuple with the Totalamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetTotalamount(v float64)`

SetTotalamount sets Totalamount field to given value.

### HasTotalamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasTotalamount() bool`

HasTotalamount returns a boolean if a field has been set.

### GetShiptosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShiptosuffix() string`

GetShiptosuffix returns the Shiptosuffix field if non-nil, zero value otherwise.

### GetShiptosuffixOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShiptosuffixOk() (*string, bool)`

GetShiptosuffixOk returns a tuple with the Shiptosuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiptosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetShiptosuffix(v string)`

SetShiptosuffix sets Shiptosuffix field to given value.

### HasShiptosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasShiptosuffix() bool`

HasShiptosuffix returns a boolean if a field has been set.

### GetBilltosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetBilltosuffix() string`

GetBilltosuffix returns the Billtosuffix field if non-nil, zero value otherwise.

### GetBilltosuffixOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetBilltosuffixOk() (*string, bool)`

GetBilltosuffixOk returns a tuple with the Billtosuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilltosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetBilltosuffix(v string)`

SetBilltosuffix sets Billtosuffix field to given value.

### HasBilltosuffix

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasBilltosuffix() bool`

HasBilltosuffix returns a boolean if a field has been set.

### GetFreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFreightamount() float64`

GetFreightamount returns the Freightamount field if non-nil, zero value otherwise.

### GetFreightamountOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFreightamountOk() (*float64, bool)`

GetFreightamountOk returns a tuple with the Freightamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetFreightamount(v float64)`

SetFreightamount sets Freightamount field to given value.

### HasFreightamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasFreightamount() bool`

HasFreightamount returns a boolean if a field has been set.

### GetPaymentterms

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetPaymentterms() string`

GetPaymentterms returns the Paymentterms field if non-nil, zero value otherwise.

### GetPaymenttermsOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetPaymenttermsOk() (*string, bool)`

GetPaymenttermsOk returns a tuple with the Paymentterms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentterms

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetPaymentterms(v string)`

SetPaymentterms sets Paymentterms field to given value.

### HasPaymentterms

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasPaymentterms() bool`

HasPaymentterms returns a boolean if a field has been set.

### GetOrderdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrderdate() string`

GetOrderdate returns the Orderdate field if non-nil, zero value otherwise.

### GetOrderdateOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrderdateOk() (*string, bool)`

GetOrderdateOk returns a tuple with the Orderdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetOrderdate(v string)`

SetOrderdate sets Orderdate field to given value.

### HasOrderdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasOrderdate() bool`

HasOrderdate returns a boolean if a field has been set.

### GetCarrier

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierdescription

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCarrierdescription() string`

GetCarrierdescription returns the Carrierdescription field if non-nil, zero value otherwise.

### GetCarrierdescriptionOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCarrierdescriptionOk() (*string, bool)`

GetCarrierdescriptionOk returns a tuple with the Carrierdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierdescription

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCarrierdescription(v string)`

SetCarrierdescription sets Carrierdescription field to given value.

### HasCarrierdescription

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCarrierdescription() bool`

HasCarrierdescription returns a boolean if a field has been set.

### GetDiscountamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetDiscountamount() float64`

GetDiscountamount returns the Discountamount field if non-nil, zero value otherwise.

### GetDiscountamountOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetDiscountamountOk() (*float64, bool)`

GetDiscountamountOk returns a tuple with the Discountamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetDiscountamount(v float64)`

SetDiscountamount sets Discountamount field to given value.

### HasDiscountamount

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasDiscountamount() bool`

HasDiscountamount returns a boolean if a field has been set.

### GetTaxtype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTaxtype() string`

GetTaxtype returns the Taxtype field if non-nil, zero value otherwise.

### GetTaxtypeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTaxtypeOk() (*string, bool)`

GetTaxtypeOk returns a tuple with the Taxtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxtype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetTaxtype(v string)`

SetTaxtype sets Taxtype field to given value.

### HasTaxtype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasTaxtype() bool`

HasTaxtype returns a boolean if a field has been set.

### GetEnduserponumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetEnduserponumber() string`

GetEnduserponumber returns the Enduserponumber field if non-nil, zero value otherwise.

### GetEnduserponumberOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetEnduserponumberOk() (*string, bool)`

GetEnduserponumberOk returns a tuple with the Enduserponumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserponumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetEnduserponumber(v string)`

SetEnduserponumber sets Enduserponumber field to given value.

### HasEnduserponumber

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasEnduserponumber() bool`

HasEnduserponumber returns a boolean if a field has been set.

### GetFreightforwardercode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFreightforwardercode() string`

GetFreightforwardercode returns the Freightforwardercode field if non-nil, zero value otherwise.

### GetFreightforwardercodeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFreightforwardercodeOk() (*string, bool)`

GetFreightforwardercodeOk returns a tuple with the Freightforwardercode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightforwardercode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetFreightforwardercode(v string)`

SetFreightforwardercode sets Freightforwardercode field to given value.

### HasFreightforwardercode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasFreightforwardercode() bool`

HasFreightforwardercode returns a boolean if a field has been set.

### GetCreditmemoreasoncode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCreditmemoreasoncode() string`

GetCreditmemoreasoncode returns the Creditmemoreasoncode field if non-nil, zero value otherwise.

### GetCreditmemoreasoncodeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCreditmemoreasoncodeOk() (*string, bool)`

GetCreditmemoreasoncodeOk returns a tuple with the Creditmemoreasoncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditmemoreasoncode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCreditmemoreasoncode(v string)`

SetCreditmemoreasoncode sets Creditmemoreasoncode field to given value.

### HasCreditmemoreasoncode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCreditmemoreasoncode() bool`

HasCreditmemoreasoncode returns a boolean if a field has been set.

### GetFulfillmentflag

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFulfillmentflag() string`

GetFulfillmentflag returns the Fulfillmentflag field if non-nil, zero value otherwise.

### GetFulfillmentflagOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetFulfillmentflagOk() (*string, bool)`

GetFulfillmentflagOk returns a tuple with the Fulfillmentflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentflag

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetFulfillmentflag(v string)`

SetFulfillmentflag sets Fulfillmentflag field to given value.

### HasFulfillmentflag

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasFulfillmentflag() bool`

HasFulfillmentflag returns a boolean if a field has been set.

### GetHoldreason

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetHoldreason() string`

GetHoldreason returns the Holdreason field if non-nil, zero value otherwise.

### GetHoldreasonOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetHoldreasonOk() (*string, bool)`

GetHoldreasonOk returns a tuple with the Holdreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldreason

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetHoldreason(v string)`

SetHoldreason sets Holdreason field to given value.

### HasHoldreason

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasHoldreason() bool`

HasHoldreason returns a boolean if a field has been set.

### GetShipcomplete

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShipcomplete() string`

GetShipcomplete returns the Shipcomplete field if non-nil, zero value otherwise.

### GetShipcompleteOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShipcompleteOk() (*string, bool)`

GetShipcompleteOk returns a tuple with the Shipcomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipcomplete

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetShipcomplete(v string)`

SetShipcomplete sets Shipcomplete field to given value.

### HasShipcomplete

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasShipcomplete() bool`

HasShipcomplete returns a boolean if a field has been set.

### GetShipdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShipdate() string`

GetShipdate returns the Shipdate field if non-nil, zero value otherwise.

### GetShipdateOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShipdateOk() (*string, bool)`

GetShipdateOk returns a tuple with the Shipdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetShipdate(v string)`

SetShipdate sets Shipdate field to given value.

### HasShipdate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasShipdate() bool`

HasShipdate returns a boolean if a field has been set.

### GetCompanycurrency

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCompanycurrency() string`

GetCompanycurrency returns the Companycurrency field if non-nil, zero value otherwise.

### GetCompanycurrencyOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCompanycurrencyOk() (*string, bool)`

GetCompanycurrencyOk returns a tuple with the Companycurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanycurrency

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCompanycurrency(v string)`

SetCompanycurrency sets Companycurrency field to given value.

### HasCompanycurrency

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCompanycurrency() bool`

HasCompanycurrency returns a boolean if a field has been set.

### GetCurrencycode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCurrencycode() string`

GetCurrencycode returns the Currencycode field if non-nil, zero value otherwise.

### GetCurrencycodeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCurrencycodeOk() (*string, bool)`

GetCurrencycodeOk returns a tuple with the Currencycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencycode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCurrencycode(v string)`

SetCurrencycode sets Currencycode field to given value.

### HasCurrencycode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCurrencycode() bool`

HasCurrencycode returns a boolean if a field has been set.

### GetCurrencyrate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCurrencyrate() string`

GetCurrencyrate returns the Currencyrate field if non-nil, zero value otherwise.

### GetCurrencyrateOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetCurrencyrateOk() (*string, bool)`

GetCurrencyrateOk returns a tuple with the Currencyrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyrate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetCurrencyrate(v string)`

SetCurrencyrate sets Currencyrate field to given value.

### HasCurrencyrate

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasCurrencyrate() bool`

HasCurrencyrate returns a boolean if a field has been set.

### GetGlobalorderid

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetGlobalorderid() string`

GetGlobalorderid returns the Globalorderid field if non-nil, zero value otherwise.

### GetGlobalorderidOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetGlobalorderidOk() (*string, bool)`

GetGlobalorderidOk returns a tuple with the Globalorderid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalorderid

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetGlobalorderid(v string)`

SetGlobalorderid sets Globalorderid field to given value.

### HasGlobalorderid

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasGlobalorderid() bool`

HasGlobalorderid returns a boolean if a field has been set.

### GetOriginalshipcode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOriginalshipcode() string`

GetOriginalshipcode returns the Originalshipcode field if non-nil, zero value otherwise.

### GetOriginalshipcodeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOriginalshipcodeOk() (*string, bool)`

GetOriginalshipcodeOk returns a tuple with the Originalshipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalshipcode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetOriginalshipcode(v string)`

SetOriginalshipcode sets Originalshipcode field to given value.

### HasOriginalshipcode

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasOriginalshipcode() bool`

HasOriginalshipcode returns a boolean if a field has been set.

### GetOrdertype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrdertype() string`

GetOrdertype returns the Ordertype field if non-nil, zero value otherwise.

### GetOrdertypeOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrdertypeOk() (*string, bool)`

GetOrdertypeOk returns a tuple with the Ordertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetOrdertype(v string)`

SetOrdertype sets Ordertype field to given value.

### HasOrdertype

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasOrdertype() bool`

HasOrdertype returns a boolean if a field has been set.

### GetOrderstatus

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrderstatus() string`

GetOrderstatus returns the Orderstatus field if non-nil, zero value otherwise.

### GetOrderstatusOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetOrderstatusOk() (*string, bool)`

GetOrderstatusOk returns a tuple with the Orderstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderstatus

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetOrderstatus(v string)`

SetOrderstatus sets Orderstatus field to given value.

### HasOrderstatus

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasOrderstatus() bool`

HasOrderstatus returns a boolean if a field has been set.

### GetTotalotherfees

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalotherfees() float32`

GetTotalotherfees returns the Totalotherfees field if non-nil, zero value otherwise.

### GetTotalotherfeesOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalotherfeesOk() (*float32, bool)`

GetTotalotherfeesOk returns a tuple with the Totalotherfees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalotherfees

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetTotalotherfees(v float32)`

SetTotalotherfees sets Totalotherfees field to given value.

### HasTotalotherfees

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasTotalotherfees() bool`

HasTotalotherfees returns a boolean if a field has been set.

### GetTotalsales

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalsales() string`

GetTotalsales returns the Totalsales field if non-nil, zero value otherwise.

### GetTotalsalesOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetTotalsalesOk() (*string, bool)`

GetTotalsalesOk returns a tuple with the Totalsales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsales

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetTotalsales(v string)`

SetTotalsales sets Totalsales field to given value.

### HasTotalsales

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasTotalsales() bool`

HasTotalsales returns a boolean if a field has been set.

### GetWeight

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetWeight(v string)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetShippableswitch

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShippableswitch() string`

GetShippableswitch returns the Shippableswitch field if non-nil, zero value otherwise.

### GetShippableswitchOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShippableswitchOk() (*string, bool)`

GetShippableswitchOk returns a tuple with the Shippableswitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippableswitch

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetShippableswitch(v string)`

SetShippableswitch sets Shippableswitch field to given value.

### HasShippableswitch

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasShippableswitch() bool`

HasShippableswitch returns a boolean if a field has been set.

### GetSoldto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetSoldto() AddressType`

GetSoldto returns the Soldto field if non-nil, zero value otherwise.

### GetSoldtoOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetSoldtoOk() (*AddressType, bool)`

GetSoldtoOk returns a tuple with the Soldto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetSoldto(v AddressType)`

SetSoldto sets Soldto field to given value.

### HasSoldto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasSoldto() bool`

HasSoldto returns a boolean if a field has been set.

### GetBillto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetBillto() AddressType`

GetBillto returns the Billto field if non-nil, zero value otherwise.

### GetBilltoOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetBilltoOk() (*AddressType, bool)`

GetBilltoOk returns a tuple with the Billto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetBillto(v AddressType)`

SetBillto sets Billto field to given value.

### HasBillto

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasBillto() bool`

HasBillto returns a boolean if a field has been set.

### GetShoptoaddress

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShoptoaddress() AddressType`

GetShoptoaddress returns the Shoptoaddress field if non-nil, zero value otherwise.

### GetShoptoaddressOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetShoptoaddressOk() (*AddressType, bool)`

GetShoptoaddressOk returns a tuple with the Shoptoaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoptoaddress

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetShoptoaddress(v AddressType)`

SetShoptoaddress sets Shoptoaddress field to given value.

### HasShoptoaddress

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasShoptoaddress() bool`

HasShoptoaddress returns a boolean if a field has been set.

### GetLines

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetLines() []ProductLineType`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetLinesOk() (*[]ProductLineType, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetLines(v []ProductLineType)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetExtendedspecs

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetExtendedspecs() []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner`

GetExtendedspecs returns the Extendedspecs field if non-nil, zero value otherwise.

### GetExtendedspecsOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetExtendedspecsOk() (*[]InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner, bool)`

GetExtendedspecsOk returns a tuple with the Extendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedspecs

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetExtendedspecs(v []InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner)`

SetExtendedspecs sets Extendedspecs field to given value.

### HasExtendedspecs

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasExtendedspecs() bool`

HasExtendedspecs returns a boolean if a field has been set.

### GetMiscfeeline

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetMiscfeeline() []InvoiceDetailResponseServiceresponseInvoicedetailresponseMiscfeelineInner`

GetMiscfeeline returns the Miscfeeline field if non-nil, zero value otherwise.

### GetMiscfeelineOk

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) GetMiscfeelineOk() (*[]InvoiceDetailResponseServiceresponseInvoicedetailresponseMiscfeelineInner, bool)`

GetMiscfeelineOk returns a tuple with the Miscfeeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscfeeline

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) SetMiscfeeline(v []InvoiceDetailResponseServiceresponseInvoicedetailresponseMiscfeelineInner)`

SetMiscfeeline sets Miscfeeline field to given value.

### HasMiscfeeline

`func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponse) HasMiscfeeline() bool`

HasMiscfeeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


