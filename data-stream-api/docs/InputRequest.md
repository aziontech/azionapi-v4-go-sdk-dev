# InputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;raw_logs&#x60; - Raw Logs | 
**Attributes** | [**InputAttributesPolymorphicInputDataSourceRequest**](InputAttributesPolymorphicInputDataSourceRequest.md) |  | 

## Methods

### NewInputRequest

`func NewInputRequest(type_ string, attributes InputAttributesPolymorphicInputDataSourceRequest, ) *InputRequest`

NewInputRequest instantiates a new InputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRequestWithDefaults

`func NewInputRequestWithDefaults() *InputRequest`

NewInputRequestWithDefaults instantiates a new InputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InputRequest) GetAttributes() InputAttributesPolymorphicInputDataSourceRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InputRequest) GetAttributesOk() (*InputAttributesPolymorphicInputDataSourceRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InputRequest) SetAttributes(v InputAttributesPolymorphicInputDataSourceRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


