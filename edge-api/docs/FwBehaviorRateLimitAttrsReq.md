# FwBehaviorRateLimitAttrsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | * &#x60;second&#x60; - second * &#x60;minute&#x60; - minute | [optional] 
**LimitBy** | **string** | * &#x60;client_ip&#x60; - client_ip * &#x60;global&#x60; - global | 
**AverageRateLimit** | **int64** |  | 
**MaximumBurstSize** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewFwBehaviorRateLimitAttrsReq

`func NewFwBehaviorRateLimitAttrsReq(limitBy string, averageRateLimit int64, ) *FwBehaviorRateLimitAttrsReq`

NewFwBehaviorRateLimitAttrsReq instantiates a new FwBehaviorRateLimitAttrsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorRateLimitAttrsReqWithDefaults

`func NewFwBehaviorRateLimitAttrsReqWithDefaults() *FwBehaviorRateLimitAttrsReq`

NewFwBehaviorRateLimitAttrsReqWithDefaults instantiates a new FwBehaviorRateLimitAttrsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FwBehaviorRateLimitAttrsReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FwBehaviorRateLimitAttrsReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FwBehaviorRateLimitAttrsReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FwBehaviorRateLimitAttrsReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *FwBehaviorRateLimitAttrsReq) GetLimitBy() string`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *FwBehaviorRateLimitAttrsReq) GetLimitByOk() (*string, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *FwBehaviorRateLimitAttrsReq) SetLimitBy(v string)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *FwBehaviorRateLimitAttrsReq) GetAverageRateLimit() int64`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *FwBehaviorRateLimitAttrsReq) GetAverageRateLimitOk() (*int64, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *FwBehaviorRateLimitAttrsReq) SetAverageRateLimit(v int64)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *FwBehaviorRateLimitAttrsReq) GetMaximumBurstSize() int64`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *FwBehaviorRateLimitAttrsReq) GetMaximumBurstSizeOk() (*int64, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *FwBehaviorRateLimitAttrsReq) SetMaximumBurstSize(v int64)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *FwBehaviorRateLimitAttrsReq) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.

### SetMaximumBurstSizeNil

`func (o *FwBehaviorRateLimitAttrsReq) SetMaximumBurstSizeNil(b bool)`

 SetMaximumBurstSizeNil sets the value for MaximumBurstSize to be an explicit nil

### UnsetMaximumBurstSize
`func (o *FwBehaviorRateLimitAttrsReq) UnsetMaximumBurstSize()`

UnsetMaximumBurstSize ensures that no value is present for MaximumBurstSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


