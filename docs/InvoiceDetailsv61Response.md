# InvoiceDetailsv61Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | Pointer to **string** | The Invoice number for the order. | [optional] 
**InvoiceStatus** | Pointer to **string** | Status of the invoice. | [optional] 
**InvoiceDate** | Pointer to **string** | Date of an Invoice. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**EndCustomerOrderNumber** | Pointer to **string** | The end customer&#39;s order number for reference in their system. | [optional] 
**OrderDate** | Pointer to **string** | The date and time in UTC format that the order was created. | [optional] 
**BillToID** | Pointer to **string** | Bill to party | [optional] 
**InvoiceType** | Pointer to **string** | Type of the Invoice | [optional] 
**InvoiceDueDate** | Pointer to **string** | Date when the invoice is due. | [optional] 
**CustomerCountryCode** | Pointer to **string** | Customer country code. | [optional] 
**CustomerNumber** | Pointer to **string** | Unique customer number in Ingram&#39;s system. | [optional] 
**IngramOrderNumber** | Pointer to **string** | The IngramMicro sales order number. | [optional] 
**Notes** | Pointer to **string** | Notes for the invoice. | [optional] 
**PaymentTermsInfo** | Pointer to [**InvoiceDetailsv61ResponsePaymentTermsInfo**](InvoiceDetailsv61ResponsePaymentTermsInfo.md) |  | [optional] 
**BillToInfo** | Pointer to [**InvoiceDetailsv61ResponseBillToInfo**](InvoiceDetailsv61ResponseBillToInfo.md) |  | [optional] 
**ShipToInfo** | Pointer to [**InvoiceDetailsv61ResponseShipToInfo**](InvoiceDetailsv61ResponseShipToInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]InvoiceDetailsv61ResponseLinesInner**](InvoiceDetailsv61ResponseLinesInner.md) |  | [optional] 
**FxRateInfo** | Pointer to [**InvoiceDetailsv61ResponseFxRateInfo**](InvoiceDetailsv61ResponseFxRateInfo.md) |  | [optional] 
**Summary** | Pointer to [**InvoiceDetailsv61ResponseSummary**](InvoiceDetailsv61ResponseSummary.md) |  | [optional] 

## Methods

### NewInvoiceDetailsv61Response

`func NewInvoiceDetailsv61Response() *InvoiceDetailsv61Response`

NewInvoiceDetailsv61Response instantiates a new InvoiceDetailsv61Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailsv61ResponseWithDefaults

`func NewInvoiceDetailsv61ResponseWithDefaults() *InvoiceDetailsv61Response`

NewInvoiceDetailsv61ResponseWithDefaults instantiates a new InvoiceDetailsv61Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *InvoiceDetailsv61Response) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceDetailsv61Response) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceDetailsv61Response) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoiceDetailsv61Response) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceDetailsv61Response) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceDetailsv61Response) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceDetailsv61Response) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceDetailsv61Response) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *InvoiceDetailsv61Response) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceDetailsv61Response) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceDetailsv61Response) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *InvoiceDetailsv61Response) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *InvoiceDetailsv61Response) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetEndCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) GetEndCustomerOrderNumber() string`

GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field if non-nil, zero value otherwise.

### GetEndCustomerOrderNumberOk

`func (o *InvoiceDetailsv61Response) GetEndCustomerOrderNumberOk() (*string, bool)`

GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) SetEndCustomerOrderNumber(v string)`

SetEndCustomerOrderNumber sets EndCustomerOrderNumber field to given value.

### HasEndCustomerOrderNumber

`func (o *InvoiceDetailsv61Response) HasEndCustomerOrderNumber() bool`

HasEndCustomerOrderNumber returns a boolean if a field has been set.

### GetOrderDate

`func (o *InvoiceDetailsv61Response) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *InvoiceDetailsv61Response) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *InvoiceDetailsv61Response) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *InvoiceDetailsv61Response) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetBillToID

`func (o *InvoiceDetailsv61Response) GetBillToID() string`

GetBillToID returns the BillToID field if non-nil, zero value otherwise.

### GetBillToIDOk

`func (o *InvoiceDetailsv61Response) GetBillToIDOk() (*string, bool)`

GetBillToIDOk returns a tuple with the BillToID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToID

`func (o *InvoiceDetailsv61Response) SetBillToID(v string)`

SetBillToID sets BillToID field to given value.

### HasBillToID

`func (o *InvoiceDetailsv61Response) HasBillToID() bool`

HasBillToID returns a boolean if a field has been set.

### GetInvoiceType

`func (o *InvoiceDetailsv61Response) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *InvoiceDetailsv61Response) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *InvoiceDetailsv61Response) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *InvoiceDetailsv61Response) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetInvoiceDueDate

`func (o *InvoiceDetailsv61Response) GetInvoiceDueDate() string`

GetInvoiceDueDate returns the InvoiceDueDate field if non-nil, zero value otherwise.

### GetInvoiceDueDateOk

`func (o *InvoiceDetailsv61Response) GetInvoiceDueDateOk() (*string, bool)`

GetInvoiceDueDateOk returns a tuple with the InvoiceDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDueDate

`func (o *InvoiceDetailsv61Response) SetInvoiceDueDate(v string)`

SetInvoiceDueDate sets InvoiceDueDate field to given value.

### HasInvoiceDueDate

`func (o *InvoiceDetailsv61Response) HasInvoiceDueDate() bool`

HasInvoiceDueDate returns a boolean if a field has been set.

### GetCustomerCountryCode

`func (o *InvoiceDetailsv61Response) GetCustomerCountryCode() string`

GetCustomerCountryCode returns the CustomerCountryCode field if non-nil, zero value otherwise.

### GetCustomerCountryCodeOk

`func (o *InvoiceDetailsv61Response) GetCustomerCountryCodeOk() (*string, bool)`

GetCustomerCountryCodeOk returns a tuple with the CustomerCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountryCode

`func (o *InvoiceDetailsv61Response) SetCustomerCountryCode(v string)`

SetCustomerCountryCode sets CustomerCountryCode field to given value.

### HasCustomerCountryCode

`func (o *InvoiceDetailsv61Response) HasCustomerCountryCode() bool`

HasCustomerCountryCode returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *InvoiceDetailsv61Response) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *InvoiceDetailsv61Response) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *InvoiceDetailsv61Response) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *InvoiceDetailsv61Response) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetIngramOrderNumber

`func (o *InvoiceDetailsv61Response) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *InvoiceDetailsv61Response) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *InvoiceDetailsv61Response) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *InvoiceDetailsv61Response) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceDetailsv61Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceDetailsv61Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceDetailsv61Response) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceDetailsv61Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPaymentTermsInfo

`func (o *InvoiceDetailsv61Response) GetPaymentTermsInfo() InvoiceDetailsv61ResponsePaymentTermsInfo`

GetPaymentTermsInfo returns the PaymentTermsInfo field if non-nil, zero value otherwise.

### GetPaymentTermsInfoOk

`func (o *InvoiceDetailsv61Response) GetPaymentTermsInfoOk() (*InvoiceDetailsv61ResponsePaymentTermsInfo, bool)`

GetPaymentTermsInfoOk returns a tuple with the PaymentTermsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermsInfo

`func (o *InvoiceDetailsv61Response) SetPaymentTermsInfo(v InvoiceDetailsv61ResponsePaymentTermsInfo)`

SetPaymentTermsInfo sets PaymentTermsInfo field to given value.

### HasPaymentTermsInfo

`func (o *InvoiceDetailsv61Response) HasPaymentTermsInfo() bool`

HasPaymentTermsInfo returns a boolean if a field has been set.

### GetBillToInfo

`func (o *InvoiceDetailsv61Response) GetBillToInfo() InvoiceDetailsv61ResponseBillToInfo`

GetBillToInfo returns the BillToInfo field if non-nil, zero value otherwise.

### GetBillToInfoOk

`func (o *InvoiceDetailsv61Response) GetBillToInfoOk() (*InvoiceDetailsv61ResponseBillToInfo, bool)`

GetBillToInfoOk returns a tuple with the BillToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToInfo

`func (o *InvoiceDetailsv61Response) SetBillToInfo(v InvoiceDetailsv61ResponseBillToInfo)`

SetBillToInfo sets BillToInfo field to given value.

### HasBillToInfo

`func (o *InvoiceDetailsv61Response) HasBillToInfo() bool`

HasBillToInfo returns a boolean if a field has been set.

### GetShipToInfo

`func (o *InvoiceDetailsv61Response) GetShipToInfo() InvoiceDetailsv61ResponseShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *InvoiceDetailsv61Response) GetShipToInfoOk() (*InvoiceDetailsv61ResponseShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *InvoiceDetailsv61Response) SetShipToInfo(v InvoiceDetailsv61ResponseShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *InvoiceDetailsv61Response) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetLines

`func (o *InvoiceDetailsv61Response) GetLines() []InvoiceDetailsv61ResponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceDetailsv61Response) GetLinesOk() (*[]InvoiceDetailsv61ResponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceDetailsv61Response) SetLines(v []InvoiceDetailsv61ResponseLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoiceDetailsv61Response) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetFxRateInfo

`func (o *InvoiceDetailsv61Response) GetFxRateInfo() InvoiceDetailsv61ResponseFxRateInfo`

GetFxRateInfo returns the FxRateInfo field if non-nil, zero value otherwise.

### GetFxRateInfoOk

`func (o *InvoiceDetailsv61Response) GetFxRateInfoOk() (*InvoiceDetailsv61ResponseFxRateInfo, bool)`

GetFxRateInfoOk returns a tuple with the FxRateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateInfo

`func (o *InvoiceDetailsv61Response) SetFxRateInfo(v InvoiceDetailsv61ResponseFxRateInfo)`

SetFxRateInfo sets FxRateInfo field to given value.

### HasFxRateInfo

`func (o *InvoiceDetailsv61Response) HasFxRateInfo() bool`

HasFxRateInfo returns a boolean if a field has been set.

### GetSummary

`func (o *InvoiceDetailsv61Response) GetSummary() InvoiceDetailsv61ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InvoiceDetailsv61Response) GetSummaryOk() (*InvoiceDetailsv61ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InvoiceDetailsv61Response) SetSummary(v InvoiceDetailsv61ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *InvoiceDetailsv61Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


