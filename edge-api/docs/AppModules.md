# AppModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**CacheModule**](CacheModule.md) |  | [optional] 
**Functions** | Pointer to [**EdgeFunctionModule**](EdgeFunctionModule.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**AppAccelModule**](AppAccelModule.md) |  | [optional] 
**ImageProcessor** | Pointer to [**ImageProcessorModule**](ImageProcessorModule.md) |  | [optional] 

## Methods

### NewAppModules

`func NewAppModules() *AppModules`

NewAppModules instantiates a new AppModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppModulesWithDefaults

`func NewAppModulesWithDefaults() *AppModules`

NewAppModulesWithDefaults instantiates a new AppModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *AppModules) GetCache() CacheModule`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *AppModules) GetCacheOk() (*CacheModule, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *AppModules) SetCache(v CacheModule)`

SetCache sets Cache field to given value.

### HasCache

`func (o *AppModules) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetFunctions

`func (o *AppModules) GetFunctions() EdgeFunctionModule`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *AppModules) GetFunctionsOk() (*EdgeFunctionModule, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *AppModules) SetFunctions(v EdgeFunctionModule)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *AppModules) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetApplicationAccelerator

`func (o *AppModules) GetApplicationAccelerator() AppAccelModule`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *AppModules) GetApplicationAcceleratorOk() (*AppAccelModule, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *AppModules) SetApplicationAccelerator(v AppAccelModule)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *AppModules) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.

### GetImageProcessor

`func (o *AppModules) GetImageProcessor() ImageProcessorModule`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *AppModules) GetImageProcessorOk() (*ImageProcessorModule, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *AppModules) SetImageProcessor(v ImageProcessorModule)`

SetImageProcessor sets ImageProcessor field to given value.

### HasImageProcessor

`func (o *AppModules) HasImageProcessor() bool`

HasImageProcessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


