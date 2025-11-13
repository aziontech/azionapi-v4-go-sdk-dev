# KBAskRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | **string** |  | 
**TopK** | Pointer to **int64** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKBAskRequestRequest

`func NewKBAskRequestRequest(question string, ) *KBAskRequestRequest`

NewKBAskRequestRequest instantiates a new KBAskRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKBAskRequestRequestWithDefaults

`func NewKBAskRequestRequestWithDefaults() *KBAskRequestRequest`

NewKBAskRequestRequestWithDefaults instantiates a new KBAskRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *KBAskRequestRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *KBAskRequestRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *KBAskRequestRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetTopK

`func (o *KBAskRequestRequest) GetTopK() int64`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *KBAskRequestRequest) GetTopKOk() (*int64, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *KBAskRequestRequest) SetTopK(v int64)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *KBAskRequestRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetModel

`func (o *KBAskRequestRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *KBAskRequestRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *KBAskRequestRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *KBAskRequestRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *KBAskRequestRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *KBAskRequestRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


