# InputDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;raw_logs&#x60; - Raw Logs | 
**Attributes** | [**DataSourceAttributesRequest**](DataSourceAttributesRequest.md) |  | 

## Methods

### NewInputDataSourceRequest

`func NewInputDataSourceRequest(type_ string, attributes DataSourceAttributesRequest, ) *InputDataSourceRequest`

NewInputDataSourceRequest instantiates a new InputDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDataSourceRequestWithDefaults

`func NewInputDataSourceRequestWithDefaults() *InputDataSourceRequest`

NewInputDataSourceRequestWithDefaults instantiates a new InputDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputDataSourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputDataSourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputDataSourceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InputDataSourceRequest) GetAttributes() DataSourceAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InputDataSourceRequest) GetAttributesOk() (*DataSourceAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InputDataSourceRequest) SetAttributes(v DataSourceAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


