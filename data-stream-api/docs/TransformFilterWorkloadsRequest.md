# TransformFilterWorkloadsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_workloads&#x60; - Filter Workloads | 
**Attributes** | [**FilterWorkloadsAttributesRequest**](FilterWorkloadsAttributesRequest.md) |  | 

## Methods

### NewTransformFilterWorkloadsRequest

`func NewTransformFilterWorkloadsRequest(type_ string, attributes FilterWorkloadsAttributesRequest, ) *TransformFilterWorkloadsRequest`

NewTransformFilterWorkloadsRequest instantiates a new TransformFilterWorkloadsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformFilterWorkloadsRequestWithDefaults

`func NewTransformFilterWorkloadsRequestWithDefaults() *TransformFilterWorkloadsRequest`

NewTransformFilterWorkloadsRequestWithDefaults instantiates a new TransformFilterWorkloadsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransformFilterWorkloadsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransformFilterWorkloadsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransformFilterWorkloadsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TransformFilterWorkloadsRequest) GetAttributes() FilterWorkloadsAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransformFilterWorkloadsRequest) GetAttributesOk() (*FilterWorkloadsAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransformFilterWorkloadsRequest) SetAttributes(v FilterWorkloadsAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


