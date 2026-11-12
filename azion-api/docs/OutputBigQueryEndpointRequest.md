# OutputBigQueryEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**BigQueryEndpointRequest**](BigQueryEndpointRequest.md) |  | 

## Methods

### NewOutputBigQueryEndpointRequest

`func NewOutputBigQueryEndpointRequest(type_ string, attributes BigQueryEndpointRequest, ) *OutputBigQueryEndpointRequest`

NewOutputBigQueryEndpointRequest instantiates a new OutputBigQueryEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputBigQueryEndpointRequestWithDefaults

`func NewOutputBigQueryEndpointRequestWithDefaults() *OutputBigQueryEndpointRequest`

NewOutputBigQueryEndpointRequestWithDefaults instantiates a new OutputBigQueryEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputBigQueryEndpointRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputBigQueryEndpointRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputBigQueryEndpointRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OutputBigQueryEndpointRequest) GetAttributes() BigQueryEndpointRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputBigQueryEndpointRequest) GetAttributesOk() (*BigQueryEndpointRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputBigQueryEndpointRequest) SetAttributes(v BigQueryEndpointRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


