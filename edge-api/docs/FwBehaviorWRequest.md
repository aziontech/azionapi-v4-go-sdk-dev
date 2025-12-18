# FwBehaviorWRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_waf&#x60; - set_waf | 
**Attributes** | [**FwBehaviorWAttrsReq**](FwBehaviorWAttrsReq.md) |  | 

## Methods

### NewFwBehaviorWRequest

`func NewFwBehaviorWRequest(type_ string, attributes FwBehaviorWAttrsReq, ) *FwBehaviorWRequest`

NewFwBehaviorWRequest instantiates a new FwBehaviorWRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorWRequestWithDefaults

`func NewFwBehaviorWRequestWithDefaults() *FwBehaviorWRequest`

NewFwBehaviorWRequestWithDefaults instantiates a new FwBehaviorWRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorWRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorWRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorWRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorWRequest) GetAttributes() FwBehaviorWAttrsReq`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorWRequest) GetAttributesOk() (*FwBehaviorWAttrsReq, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorWRequest) SetAttributes(v FwBehaviorWAttrsReq)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


