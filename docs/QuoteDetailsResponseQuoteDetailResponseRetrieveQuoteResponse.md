# QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteGuid** | Pointer to **string** |  | [optional] 
**QuoteName** | Pointer to **string** |  | [optional] 
**QuoteNumber** | Pointer to **string** |  | [optional] 
**QuoteExpiryDate** | Pointer to **string** |  | [optional] 
**RevisionNumber** | Pointer to **string** |  | [optional] 
**IntroPreamble** | Pointer to **string** |  | [optional] 
**PurchaseInstructions** | Pointer to **string** |  | [optional] 
**LegalTerms** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**PriceDeviationId** | Pointer to **string** |  | [optional] 
**PriceDeviationStartDate** | Pointer to **string** |  | [optional] 
**PriceDeviationExpiryDate** | Pointer to **string** |  | [optional] 
**CustomerNeed** | Pointer to **string** |  | [optional] 
**SolutionProposed** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**LeasingCalculations** | Pointer to **string** |  | [optional] 
**LeasingInstructions** | Pointer to **string** |  | [optional] 
**AccountInfo** | Pointer to [**QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseAccountInfo**](QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseAccountInfo.md) |  | [optional] 
**ContactInfo** | Pointer to [**QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseContactInfo**](QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseContactInfo.md) |  | [optional] 
**VendorAttributes** | Pointer to [**QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseVendorAttributes**](QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseVendorAttributes.md) |  | [optional] 
**EndUser** | Pointer to [**QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser**](QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser.md) |  | [optional] 
**QuoteProductList** | Pointer to [**[]QuoteProductList**](QuoteProductList.md) |  | [optional] 
**TotalQuoteProductCount** | Pointer to **string** |  | [optional] 
**TotalExtendedMsrp** | Pointer to **string** |  | [optional] 
**TotalQuantity** | Pointer to **string** |  | [optional] 
**TotalExtendedQuotePrice** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse

`func NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse() *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse`

NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse instantiates a new QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponseWithDefaults

`func NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponseWithDefaults() *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse`

NewQuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponseWithDefaults instantiates a new QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteGuid

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteGuid() string`

GetQuoteGuid returns the QuoteGuid field if non-nil, zero value otherwise.

### GetQuoteGuidOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteGuidOk() (*string, bool)`

GetQuoteGuidOk returns a tuple with the QuoteGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteGuid

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetQuoteGuid(v string)`

SetQuoteGuid sets QuoteGuid field to given value.

### HasQuoteGuid

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasQuoteGuid() bool`

HasQuoteGuid returns a boolean if a field has been set.

### GetQuoteName

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteName() string`

GetQuoteName returns the QuoteName field if non-nil, zero value otherwise.

### GetQuoteNameOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteNameOk() (*string, bool)`

GetQuoteNameOk returns a tuple with the QuoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteName

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetQuoteName(v string)`

SetQuoteName sets QuoteName field to given value.

### HasQuoteName

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasQuoteName() bool`

HasQuoteName returns a boolean if a field has been set.

### GetQuoteNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteNumber() string`

GetQuoteNumber returns the QuoteNumber field if non-nil, zero value otherwise.

### GetQuoteNumberOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteNumberOk() (*string, bool)`

GetQuoteNumberOk returns a tuple with the QuoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetQuoteNumber(v string)`

SetQuoteNumber sets QuoteNumber field to given value.

### HasQuoteNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasQuoteNumber() bool`

HasQuoteNumber returns a boolean if a field has been set.

### GetQuoteExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteExpiryDate() string`

GetQuoteExpiryDate returns the QuoteExpiryDate field if non-nil, zero value otherwise.

### GetQuoteExpiryDateOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteExpiryDateOk() (*string, bool)`

GetQuoteExpiryDateOk returns a tuple with the QuoteExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetQuoteExpiryDate(v string)`

SetQuoteExpiryDate sets QuoteExpiryDate field to given value.

### HasQuoteExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasQuoteExpiryDate() bool`

HasQuoteExpiryDate returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetIntroPreamble

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetIntroPreamble() string`

GetIntroPreamble returns the IntroPreamble field if non-nil, zero value otherwise.

### GetIntroPreambleOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetIntroPreambleOk() (*string, bool)`

GetIntroPreambleOk returns a tuple with the IntroPreamble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroPreamble

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetIntroPreamble(v string)`

SetIntroPreamble sets IntroPreamble field to given value.

### HasIntroPreamble

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasIntroPreamble() bool`

HasIntroPreamble returns a boolean if a field has been set.

### GetPurchaseInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPurchaseInstructions() string`

GetPurchaseInstructions returns the PurchaseInstructions field if non-nil, zero value otherwise.

### GetPurchaseInstructionsOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPurchaseInstructionsOk() (*string, bool)`

GetPurchaseInstructionsOk returns a tuple with the PurchaseInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetPurchaseInstructions(v string)`

SetPurchaseInstructions sets PurchaseInstructions field to given value.

### HasPurchaseInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasPurchaseInstructions() bool`

HasPurchaseInstructions returns a boolean if a field has been set.

### GetLegalTerms

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLegalTerms() string`

GetLegalTerms returns the LegalTerms field if non-nil, zero value otherwise.

### GetLegalTermsOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLegalTermsOk() (*string, bool)`

GetLegalTermsOk returns a tuple with the LegalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalTerms

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetLegalTerms(v string)`

SetLegalTerms sets LegalTerms field to given value.

### HasLegalTerms

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasLegalTerms() bool`

HasLegalTerms returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPriceDeviationId

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationId() string`

GetPriceDeviationId returns the PriceDeviationId field if non-nil, zero value otherwise.

### GetPriceDeviationIdOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationIdOk() (*string, bool)`

GetPriceDeviationIdOk returns a tuple with the PriceDeviationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDeviationId

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetPriceDeviationId(v string)`

SetPriceDeviationId sets PriceDeviationId field to given value.

### HasPriceDeviationId

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasPriceDeviationId() bool`

HasPriceDeviationId returns a boolean if a field has been set.

### GetPriceDeviationStartDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationStartDate() string`

GetPriceDeviationStartDate returns the PriceDeviationStartDate field if non-nil, zero value otherwise.

### GetPriceDeviationStartDateOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationStartDateOk() (*string, bool)`

GetPriceDeviationStartDateOk returns a tuple with the PriceDeviationStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDeviationStartDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetPriceDeviationStartDate(v string)`

SetPriceDeviationStartDate sets PriceDeviationStartDate field to given value.

### HasPriceDeviationStartDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasPriceDeviationStartDate() bool`

HasPriceDeviationStartDate returns a boolean if a field has been set.

### GetPriceDeviationExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationExpiryDate() string`

GetPriceDeviationExpiryDate returns the PriceDeviationExpiryDate field if non-nil, zero value otherwise.

### GetPriceDeviationExpiryDateOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetPriceDeviationExpiryDateOk() (*string, bool)`

GetPriceDeviationExpiryDateOk returns a tuple with the PriceDeviationExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDeviationExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetPriceDeviationExpiryDate(v string)`

SetPriceDeviationExpiryDate sets PriceDeviationExpiryDate field to given value.

### HasPriceDeviationExpiryDate

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasPriceDeviationExpiryDate() bool`

HasPriceDeviationExpiryDate returns a boolean if a field has been set.

### GetCustomerNeed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCustomerNeed() string`

GetCustomerNeed returns the CustomerNeed field if non-nil, zero value otherwise.

### GetCustomerNeedOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCustomerNeedOk() (*string, bool)`

GetCustomerNeedOk returns a tuple with the CustomerNeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNeed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetCustomerNeed(v string)`

SetCustomerNeed sets CustomerNeed field to given value.

### HasCustomerNeed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasCustomerNeed() bool`

HasCustomerNeed returns a boolean if a field has been set.

### GetSolutionProposed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetSolutionProposed() string`

GetSolutionProposed returns the SolutionProposed field if non-nil, zero value otherwise.

### GetSolutionProposedOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetSolutionProposedOk() (*string, bool)`

GetSolutionProposedOk returns a tuple with the SolutionProposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionProposed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetSolutionProposed(v string)`

SetSolutionProposed sets SolutionProposed field to given value.

### HasSolutionProposed

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasSolutionProposed() bool`

HasSolutionProposed returns a boolean if a field has been set.

### GetStatus

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetLeasingCalculations

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLeasingCalculations() string`

GetLeasingCalculations returns the LeasingCalculations field if non-nil, zero value otherwise.

### GetLeasingCalculationsOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLeasingCalculationsOk() (*string, bool)`

GetLeasingCalculationsOk returns a tuple with the LeasingCalculations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingCalculations

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetLeasingCalculations(v string)`

SetLeasingCalculations sets LeasingCalculations field to given value.

### HasLeasingCalculations

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasLeasingCalculations() bool`

HasLeasingCalculations returns a boolean if a field has been set.

### GetLeasingInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLeasingInstructions() string`

GetLeasingInstructions returns the LeasingInstructions field if non-nil, zero value otherwise.

### GetLeasingInstructionsOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetLeasingInstructionsOk() (*string, bool)`

GetLeasingInstructionsOk returns a tuple with the LeasingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetLeasingInstructions(v string)`

SetLeasingInstructions sets LeasingInstructions field to given value.

### HasLeasingInstructions

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasLeasingInstructions() bool`

HasLeasingInstructions returns a boolean if a field has been set.

### GetAccountInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetAccountInfo() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseAccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetAccountInfoOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseAccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetAccountInfo(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseAccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetContactInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetContactInfo() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetContactInfoOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetContactInfo(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetVendorAttributes

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetVendorAttributes() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseVendorAttributes`

GetVendorAttributes returns the VendorAttributes field if non-nil, zero value otherwise.

### GetVendorAttributesOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetVendorAttributesOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseVendorAttributes, bool)`

GetVendorAttributesOk returns a tuple with the VendorAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAttributes

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetVendorAttributes(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseVendorAttributes)`

SetVendorAttributes sets VendorAttributes field to given value.

### HasVendorAttributes

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasVendorAttributes() bool`

HasVendorAttributes returns a boolean if a field has been set.

### GetEndUser

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetEndUser() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetEndUserOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetEndUser(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetQuoteProductList

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteProductList() []QuoteProductList`

GetQuoteProductList returns the QuoteProductList field if non-nil, zero value otherwise.

### GetQuoteProductListOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetQuoteProductListOk() (*[]QuoteProductList, bool)`

GetQuoteProductListOk returns a tuple with the QuoteProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteProductList

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetQuoteProductList(v []QuoteProductList)`

SetQuoteProductList sets QuoteProductList field to given value.

### HasQuoteProductList

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasQuoteProductList() bool`

HasQuoteProductList returns a boolean if a field has been set.

### GetTotalQuoteProductCount

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalQuoteProductCount() string`

GetTotalQuoteProductCount returns the TotalQuoteProductCount field if non-nil, zero value otherwise.

### GetTotalQuoteProductCountOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalQuoteProductCountOk() (*string, bool)`

GetTotalQuoteProductCountOk returns a tuple with the TotalQuoteProductCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuoteProductCount

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetTotalQuoteProductCount(v string)`

SetTotalQuoteProductCount sets TotalQuoteProductCount field to given value.

### HasTotalQuoteProductCount

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasTotalQuoteProductCount() bool`

HasTotalQuoteProductCount returns a boolean if a field has been set.

### GetTotalExtendedMsrp

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalExtendedMsrp() string`

GetTotalExtendedMsrp returns the TotalExtendedMsrp field if non-nil, zero value otherwise.

### GetTotalExtendedMsrpOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalExtendedMsrpOk() (*string, bool)`

GetTotalExtendedMsrpOk returns a tuple with the TotalExtendedMsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExtendedMsrp

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetTotalExtendedMsrp(v string)`

SetTotalExtendedMsrp sets TotalExtendedMsrp field to given value.

### HasTotalExtendedMsrp

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasTotalExtendedMsrp() bool`

HasTotalExtendedMsrp returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalQuantity() string`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalQuantityOk() (*string, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetTotalQuantity(v string)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetTotalExtendedQuotePrice

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalExtendedQuotePrice() string`

GetTotalExtendedQuotePrice returns the TotalExtendedQuotePrice field if non-nil, zero value otherwise.

### GetTotalExtendedQuotePriceOk

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) GetTotalExtendedQuotePriceOk() (*string, bool)`

GetTotalExtendedQuotePriceOk returns a tuple with the TotalExtendedQuotePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExtendedQuotePrice

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) SetTotalExtendedQuotePrice(v string)`

SetTotalExtendedQuotePrice sets TotalExtendedQuotePrice field to given value.

### HasTotalExtendedQuotePrice

`func (o *QuoteDetailsResponseQuoteDetailResponseRetrieveQuoteResponse) HasTotalExtendedQuotePrice() bool`

HasTotalExtendedQuotePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


