# WAFAttackListsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **[]string** | Signature lists turned on in the WAF. A list blocks matching requests only while the WAF mode is blocking and the Identified Attack threat is enabled. Available lists: &#x60;php2shell&#x60;: PHP2Shell session poisoning through upload progress (CVE-2026-109618, CVE-2026-109619). | 

## Methods

### NewWAFAttackListsRequest

`func NewWAFAttackListsRequest(active []string, ) *WAFAttackListsRequest`

NewWAFAttackListsRequest instantiates a new WAFAttackListsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFAttackListsRequestWithDefaults

`func NewWAFAttackListsRequestWithDefaults() *WAFAttackListsRequest`

NewWAFAttackListsRequestWithDefaults instantiates a new WAFAttackListsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WAFAttackListsRequest) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WAFAttackListsRequest) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WAFAttackListsRequest) SetActive(v []string)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


