# FwBehaviorWAttrsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WafId** | **int64** |  | 
**Mode** | **string** | * &#x60;logging&#x60; - logging * &#x60;blocking&#x60; - blocking | 

## Methods

### NewFwBehaviorWAttrsReq

`func NewFwBehaviorWAttrsReq(wafId int64, mode string, ) *FwBehaviorWAttrsReq`

NewFwBehaviorWAttrsReq instantiates a new FwBehaviorWAttrsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwBehaviorWAttrsReqWithDefaults

`func NewFwBehaviorWAttrsReqWithDefaults() *FwBehaviorWAttrsReq`

NewFwBehaviorWAttrsReqWithDefaults instantiates a new FwBehaviorWAttrsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWafId

`func (o *FwBehaviorWAttrsReq) GetWafId() int64`

GetWafId returns the WafId field if non-nil, zero value otherwise.

### GetWafIdOk

`func (o *FwBehaviorWAttrsReq) GetWafIdOk() (*int64, bool)`

GetWafIdOk returns a tuple with the WafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafId

`func (o *FwBehaviorWAttrsReq) SetWafId(v int64)`

SetWafId sets WafId field to given value.


### GetMode

`func (o *FwBehaviorWAttrsReq) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FwBehaviorWAttrsReq) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FwBehaviorWAttrsReq) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


