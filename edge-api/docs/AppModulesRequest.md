# AppModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**CacheModuleRequest**](CacheModuleRequest.md) |  | [optional] 
**Functions** | Pointer to [**EdgeFunctionModuleRequest**](EdgeFunctionModuleRequest.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**AppAccelModuleRequest**](AppAccelModuleRequest.md) |  | [optional] 
**ImageProcessor** | Pointer to [**ImageProcessorModuleRequest**](ImageProcessorModuleRequest.md) |  | [optional] 

## Methods

### NewAppModulesRequest

`func NewAppModulesRequest() *AppModulesRequest`

NewAppModulesRequest instantiates a new AppModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppModulesRequestWithDefaults

`func NewAppModulesRequestWithDefaults() *AppModulesRequest`

NewAppModulesRequestWithDefaults instantiates a new AppModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *AppModulesRequest) GetCache() CacheModuleRequest`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *AppModulesRequest) GetCacheOk() (*CacheModuleRequest, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *AppModulesRequest) SetCache(v CacheModuleRequest)`

SetCache sets Cache field to given value.

### HasCache

`func (o *AppModulesRequest) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetFunctions

`func (o *AppModulesRequest) GetFunctions() EdgeFunctionModuleRequest`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *AppModulesRequest) GetFunctionsOk() (*EdgeFunctionModuleRequest, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *AppModulesRequest) SetFunctions(v EdgeFunctionModuleRequest)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *AppModulesRequest) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetApplicationAccelerator

`func (o *AppModulesRequest) GetApplicationAccelerator() AppAccelModuleRequest`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *AppModulesRequest) GetApplicationAcceleratorOk() (*AppAccelModuleRequest, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *AppModulesRequest) SetApplicationAccelerator(v AppAccelModuleRequest)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *AppModulesRequest) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.

### GetImageProcessor

`func (o *AppModulesRequest) GetImageProcessor() ImageProcessorModuleRequest`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *AppModulesRequest) GetImageProcessorOk() (*ImageProcessorModuleRequest, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *AppModulesRequest) SetImageProcessor(v ImageProcessorModuleRequest)`

SetImageProcessor sets ImageProcessor field to given value.

### HasImageProcessor

`func (o *AppModulesRequest) HasImageProcessor() bool`

HasImageProcessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


