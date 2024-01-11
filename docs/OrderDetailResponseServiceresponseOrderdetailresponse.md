# OrderDetailResponseServiceresponseOrderdetailresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ordernumber** | Pointer to **string** |  | [optional] 
**Ordertype** | Pointer to **string** | Order Type   B - BRANCH TRANSFER C - CASH ORDER D - DIRECT ORDER F - FUTURE ORDER P - SPECIAL ORDER Q - QUOTE ORDER S - STOCK ORDER M - MEMO ORDER | [optional] 
**Customerordernumber** | Pointer to **string** | Customer PO number | [optional] 
**Enduserponumber** | Pointer to **string** | End User PO number | [optional] 
**Orderstatus** | Pointer to **string** | Status of order within Ingram system S - SALES HOLD H - TAG HOLD I - INVOICED P - PENDING E - BILLING ERROR F - FORCE BILLING V - VOIDED T - TRANSFERRED D - HOLD SHIPMENT R - RELEASED O - IM ONLINE HOLD U - BILL FOR HISTORY ONLY W - ORDER NOT PRINTED A - DROP SHIP HOLD B - INTERNET CUST ORIG HOLD 1 - PICKED 2 - INSPECTED 3 - PACKED 4 - SHIPPED C - CREDIT HOLD 9 - CISCO 3A6 Q - RMA HOLD G - CREDIT HOLD N - CREDIT HOLD | [optional] 
**Entrytimestamp** | Pointer to **string** | Time stamp of the order placed | [optional] 
**Entrymethoddescription** | Pointer to **string** | Description of the entry method  | [optional] 
**Ordertotalvalue** | Pointer to **float32** | Total order value | [optional] 
**Ordersubtotal** | Pointer to **float32** | Subtotal order value | [optional] 
**Freightamount** | Pointer to **string** | Freight charges | [optional] 
**Currencycode** | Pointer to **string** | Country specific currency code | [optional] 
**Totalweight** | Pointer to **string** | Total order weight. unit -- North america - Pounds , other countries will be KG | [optional] 
**Totaltax** | Pointer to **string** | total tax on the orders placed | [optional] 
**Billtoaddress** | Pointer to [**OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress**](OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress.md) |  | [optional] 
**Shiptoaddress** | Pointer to [**OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress**](OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress.md) |  | [optional] 
**Enduserinfo** | Pointer to [**OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo**](OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInner**](OrderDetailResponseServiceresponseOrderdetailresponseLinesInner.md) |  | [optional] 
**Commentlines** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner**](OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner.md) |  | [optional] 
**Miscfeeline** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner**](OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner.md) |  | [optional] 
**Extendedspecs** | Pointer to [**[]OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner**](OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner.md) |  | [optional] 

## Methods

### NewOrderDetailResponseServiceresponseOrderdetailresponse

`func NewOrderDetailResponseServiceresponseOrderdetailresponse() *OrderDetailResponseServiceresponseOrderdetailresponse`

NewOrderDetailResponseServiceresponseOrderdetailresponse instantiates a new OrderDetailResponseServiceresponseOrderdetailresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseServiceresponseOrderdetailresponseWithDefaults

`func NewOrderDetailResponseServiceresponseOrderdetailresponseWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponse`

NewOrderDetailResponseServiceresponseOrderdetailresponseWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdernumber() string`

GetOrdernumber returns the Ordernumber field if non-nil, zero value otherwise.

### GetOrdernumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdernumberOk() (*string, bool)`

GetOrdernumberOk returns a tuple with the Ordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdernumber(v string)`

SetOrdernumber sets Ordernumber field to given value.

### HasOrdernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdernumber() bool`

HasOrdernumber returns a boolean if a field has been set.

### GetOrdertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertype() string`

GetOrdertype returns the Ordertype field if non-nil, zero value otherwise.

### GetOrdertypeOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertypeOk() (*string, bool)`

GetOrdertypeOk returns a tuple with the Ordertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdertype(v string)`

SetOrdertype sets Ordertype field to given value.

### HasOrdertype

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdertype() bool`

HasOrdertype returns a boolean if a field has been set.

### GetCustomerordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCustomerordernumber() string`

GetCustomerordernumber returns the Customerordernumber field if non-nil, zero value otherwise.

### GetCustomerordernumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCustomerordernumberOk() (*string, bool)`

GetCustomerordernumberOk returns a tuple with the Customerordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCustomerordernumber(v string)`

SetCustomerordernumber sets Customerordernumber field to given value.

### HasCustomerordernumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCustomerordernumber() bool`

HasCustomerordernumber returns a boolean if a field has been set.

### GetEnduserponumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserponumber() string`

GetEnduserponumber returns the Enduserponumber field if non-nil, zero value otherwise.

### GetEnduserponumberOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserponumberOk() (*string, bool)`

GetEnduserponumberOk returns a tuple with the Enduserponumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserponumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEnduserponumber(v string)`

SetEnduserponumber sets Enduserponumber field to given value.

### HasEnduserponumber

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEnduserponumber() bool`

HasEnduserponumber returns a boolean if a field has been set.

### GetOrderstatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrderstatus() string`

GetOrderstatus returns the Orderstatus field if non-nil, zero value otherwise.

### GetOrderstatusOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrderstatusOk() (*string, bool)`

GetOrderstatusOk returns a tuple with the Orderstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderstatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrderstatus(v string)`

SetOrderstatus sets Orderstatus field to given value.

### HasOrderstatus

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrderstatus() bool`

HasOrderstatus returns a boolean if a field has been set.

### GetEntrytimestamp

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrytimestamp() string`

GetEntrytimestamp returns the Entrytimestamp field if non-nil, zero value otherwise.

### GetEntrytimestampOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrytimestampOk() (*string, bool)`

GetEntrytimestampOk returns a tuple with the Entrytimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrytimestamp

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEntrytimestamp(v string)`

SetEntrytimestamp sets Entrytimestamp field to given value.

### HasEntrytimestamp

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEntrytimestamp() bool`

HasEntrytimestamp returns a boolean if a field has been set.

### GetEntrymethoddescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrymethoddescription() string`

GetEntrymethoddescription returns the Entrymethoddescription field if non-nil, zero value otherwise.

### GetEntrymethoddescriptionOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEntrymethoddescriptionOk() (*string, bool)`

GetEntrymethoddescriptionOk returns a tuple with the Entrymethoddescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrymethoddescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEntrymethoddescription(v string)`

SetEntrymethoddescription sets Entrymethoddescription field to given value.

### HasEntrymethoddescription

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEntrymethoddescription() bool`

HasEntrymethoddescription returns a boolean if a field has been set.

### GetOrdertotalvalue

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertotalvalue() float32`

GetOrdertotalvalue returns the Ordertotalvalue field if non-nil, zero value otherwise.

### GetOrdertotalvalueOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdertotalvalueOk() (*float32, bool)`

GetOrdertotalvalueOk returns a tuple with the Ordertotalvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdertotalvalue

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdertotalvalue(v float32)`

SetOrdertotalvalue sets Ordertotalvalue field to given value.

### HasOrdertotalvalue

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdertotalvalue() bool`

HasOrdertotalvalue returns a boolean if a field has been set.

### GetOrdersubtotal

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdersubtotal() float32`

GetOrdersubtotal returns the Ordersubtotal field if non-nil, zero value otherwise.

### GetOrdersubtotalOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetOrdersubtotalOk() (*float32, bool)`

GetOrdersubtotalOk returns a tuple with the Ordersubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersubtotal

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetOrdersubtotal(v float32)`

SetOrdersubtotal sets Ordersubtotal field to given value.

### HasOrdersubtotal

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasOrdersubtotal() bool`

HasOrdersubtotal returns a boolean if a field has been set.

### GetFreightamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetFreightamount() string`

GetFreightamount returns the Freightamount field if non-nil, zero value otherwise.

### GetFreightamountOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetFreightamountOk() (*string, bool)`

GetFreightamountOk returns a tuple with the Freightamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetFreightamount(v string)`

SetFreightamount sets Freightamount field to given value.

### HasFreightamount

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasFreightamount() bool`

HasFreightamount returns a boolean if a field has been set.

### GetCurrencycode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCurrencycode() string`

GetCurrencycode returns the Currencycode field if non-nil, zero value otherwise.

### GetCurrencycodeOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCurrencycodeOk() (*string, bool)`

GetCurrencycodeOk returns a tuple with the Currencycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencycode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCurrencycode(v string)`

SetCurrencycode sets Currencycode field to given value.

### HasCurrencycode

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCurrencycode() bool`

HasCurrencycode returns a boolean if a field has been set.

### GetTotalweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotalweight() string`

GetTotalweight returns the Totalweight field if non-nil, zero value otherwise.

### GetTotalweightOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotalweightOk() (*string, bool)`

GetTotalweightOk returns a tuple with the Totalweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetTotalweight(v string)`

SetTotalweight sets Totalweight field to given value.

### HasTotalweight

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasTotalweight() bool`

HasTotalweight returns a boolean if a field has been set.

### GetTotaltax

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotaltax() string`

GetTotaltax returns the Totaltax field if non-nil, zero value otherwise.

### GetTotaltaxOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetTotaltaxOk() (*string, bool)`

GetTotaltaxOk returns a tuple with the Totaltax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotaltax

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetTotaltax(v string)`

SetTotaltax sets Totaltax field to given value.

### HasTotaltax

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasTotaltax() bool`

HasTotaltax returns a boolean if a field has been set.

### GetBilltoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetBilltoaddress() OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress`

GetBilltoaddress returns the Billtoaddress field if non-nil, zero value otherwise.

### GetBilltoaddressOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetBilltoaddressOk() (*OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress, bool)`

GetBilltoaddressOk returns a tuple with the Billtoaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilltoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetBilltoaddress(v OrderDetailResponseServiceresponseOrderdetailresponseBilltoaddress)`

SetBilltoaddress sets Billtoaddress field to given value.

### HasBilltoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasBilltoaddress() bool`

HasBilltoaddress returns a boolean if a field has been set.

### GetShiptoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetShiptoaddress() OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress`

GetShiptoaddress returns the Shiptoaddress field if non-nil, zero value otherwise.

### GetShiptoaddressOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetShiptoaddressOk() (*OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress, bool)`

GetShiptoaddressOk returns a tuple with the Shiptoaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiptoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetShiptoaddress(v OrderDetailResponseServiceresponseOrderdetailresponseShiptoaddress)`

SetShiptoaddress sets Shiptoaddress field to given value.

### HasShiptoaddress

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasShiptoaddress() bool`

HasShiptoaddress returns a boolean if a field has been set.

### GetEnduserinfo

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserinfo() OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo`

GetEnduserinfo returns the Enduserinfo field if non-nil, zero value otherwise.

### GetEnduserinfoOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetEnduserinfoOk() (*OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo, bool)`

GetEnduserinfoOk returns a tuple with the Enduserinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnduserinfo

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetEnduserinfo(v OrderDetailResponseServiceresponseOrderdetailresponseEnduserinfo)`

SetEnduserinfo sets Enduserinfo field to given value.

### HasEnduserinfo

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasEnduserinfo() bool`

HasEnduserinfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetLines() []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetLinesOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetLines(v []OrderDetailResponseServiceresponseOrderdetailresponseLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetCommentlines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCommentlines() []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner`

GetCommentlines returns the Commentlines field if non-nil, zero value otherwise.

### GetCommentlinesOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetCommentlinesOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner, bool)`

GetCommentlinesOk returns a tuple with the Commentlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentlines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetCommentlines(v []OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner)`

SetCommentlines sets Commentlines field to given value.

### HasCommentlines

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasCommentlines() bool`

HasCommentlines returns a boolean if a field has been set.

### GetMiscfeeline

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetMiscfeeline() []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner`

GetMiscfeeline returns the Miscfeeline field if non-nil, zero value otherwise.

### GetMiscfeelineOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetMiscfeelineOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner, bool)`

GetMiscfeelineOk returns a tuple with the Miscfeeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscfeeline

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetMiscfeeline(v []OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner)`

SetMiscfeeline sets Miscfeeline field to given value.

### HasMiscfeeline

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasMiscfeeline() bool`

HasMiscfeeline returns a boolean if a field has been set.

### GetExtendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetExtendedspecs() []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner`

GetExtendedspecs returns the Extendedspecs field if non-nil, zero value otherwise.

### GetExtendedspecsOk

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) GetExtendedspecsOk() (*[]OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner, bool)`

GetExtendedspecsOk returns a tuple with the Extendedspecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) SetExtendedspecs(v []OrderDetailResponseServiceresponseOrderdetailresponseExtendedspecsInner)`

SetExtendedspecs sets Extendedspecs field to given value.

### HasExtendedspecs

`func (o *OrderDetailResponseServiceresponseOrderdetailresponse) HasExtendedspecs() bool`

HasExtendedspecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


