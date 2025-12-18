# FwBehaviorRateLimit2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator field for FwBehaviors | 
**Attributes** | [**FwBehaviorRateLimitAttrs**](FwBehaviorRateLimitAttrs.md) |  | 

## Methods

### NewFwBehaviorRateLimit2

`func NewFwBehaviorRateLimit2(type_ string, attributes FwBehaviorRateLimitAttrs, ) *FwBehaviorRateLimit2`

NewFwBehaviorRateLimit2 instantiates a new FwBehaviorRateLimit2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorRateLimit2WithDefaults

`func NewFwBehaviorRateLimit2WithDefaults() *FwBehaviorRateLimit2`

NewFwBehaviorRateLimit2WithDefaults instantiates a new FwBehaviorRateLimit2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorRateLimit2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorRateLimit2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorRateLimit2) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FwBehaviorRateLimit2) GetAttributes() FwBehaviorRateLimitAttrs`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FwBehaviorRateLimit2) GetAttributesOk() (*FwBehaviorRateLimitAttrs, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FwBehaviorRateLimit2) SetAttributes(v FwBehaviorRateLimitAttrs)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


