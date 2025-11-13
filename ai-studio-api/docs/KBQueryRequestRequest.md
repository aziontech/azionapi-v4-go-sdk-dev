# KBQueryRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**TopK** | Pointer to **int64** |  | [optional] 

## Methods

### NewKBQueryRequestRequest

`func NewKBQueryRequestRequest(query string, ) *KBQueryRequestRequest`

NewKBQueryRequestRequest instantiates a new KBQueryRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKBQueryRequestRequestWithDefaults

`func NewKBQueryRequestRequestWithDefaults() *KBQueryRequestRequest`

NewKBQueryRequestRequestWithDefaults instantiates a new KBQueryRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *KBQueryRequestRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KBQueryRequestRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KBQueryRequestRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTopK

`func (o *KBQueryRequestRequest) GetTopK() int64`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *KBQueryRequestRequest) GetTopKOk() (*int64, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *KBQueryRequestRequest) SetTopK(v int64)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *KBQueryRequestRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


