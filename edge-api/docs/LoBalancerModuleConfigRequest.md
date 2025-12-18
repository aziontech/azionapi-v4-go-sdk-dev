# LoBalancerModuleConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Load balancing method to use (e.g., round-robin, least_conn).  * &#x60;round_robin&#x60; - Round Robin * &#x60;least_conn&#x60; - Least Connections * &#x60;ip_hash&#x60; - IP Hash | [optional] 
**MaxRetries** | Pointer to **int64** | Maximum number of retry attempts on connection failure. | [optional] 
**ConnectionTimeout** | Pointer to **int64** | Maximum time (in seconds) to wait for a connection to be established. | [optional] 
**ReadWriteTimeout** | Pointer to **int64** | Maximum time (in seconds) to wait for data read/write after connection. | [optional] 

## Methods

### NewLoBalancerModuleConfigRequest

`func NewLoBalancerModuleConfigRequest() *LoBalancerModuleConfigRequest`

NewLoBalancerModuleConfigRequest instantiates a new LoBalancerModuleConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoBalancerModuleConfigRequestWithDefaults

`func NewLoBalancerModuleConfigRequestWithDefaults() *LoBalancerModuleConfigRequest`

NewLoBalancerModuleConfigRequestWithDefaults instantiates a new LoBalancerModuleConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LoBalancerModuleConfigRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoBalancerModuleConfigRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoBalancerModuleConfigRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoBalancerModuleConfigRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMaxRetries

`func (o *LoBalancerModuleConfigRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *LoBalancerModuleConfigRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *LoBalancerModuleConfigRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *LoBalancerModuleConfigRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *LoBalancerModuleConfigRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *LoBalancerModuleConfigRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *LoBalancerModuleConfigRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *LoBalancerModuleConfigRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *LoBalancerModuleConfigRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *LoBalancerModuleConfigRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *LoBalancerModuleConfigRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *LoBalancerModuleConfigRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


