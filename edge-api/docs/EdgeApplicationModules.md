# EdgeApplicationModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | [**CacheModule**](CacheModule.md) |  | 
**EdgeFunctions** | [**EdgeFunctionModule**](EdgeFunctionModule.md) |  | 
**ApplicationAccelerator** | [**ApplicationAcceleratorModule**](ApplicationAcceleratorModule.md) |  | 
**ImageProcessor** | [**ImageProcessorModule**](ImageProcessorModule.md) |  | 
**TieredCache** | [**TieredCacheModule**](TieredCacheModule.md) |  | 

## Methods

### NewEdgeApplicationModules

`func NewEdgeApplicationModules(edgeCache CacheModule, edgeFunctions EdgeFunctionModule, applicationAccelerator ApplicationAcceleratorModule, imageProcessor ImageProcessorModule, tieredCache TieredCacheModule, ) *EdgeApplicationModules`

NewEdgeApplicationModules instantiates a new EdgeApplicationModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationModulesWithDefaults

`func NewEdgeApplicationModulesWithDefaults() *EdgeApplicationModules`

NewEdgeApplicationModulesWithDefaults instantiates a new EdgeApplicationModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *EdgeApplicationModules) GetEdgeCache() CacheModule`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *EdgeApplicationModules) GetEdgeCacheOk() (*CacheModule, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *EdgeApplicationModules) SetEdgeCache(v CacheModule)`

SetEdgeCache sets EdgeCache field to given value.


### GetEdgeFunctions

`func (o *EdgeApplicationModules) GetEdgeFunctions() EdgeFunctionModule`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *EdgeApplicationModules) GetEdgeFunctionsOk() (*EdgeFunctionModule, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *EdgeApplicationModules) SetEdgeFunctions(v EdgeFunctionModule)`

SetEdgeFunctions sets EdgeFunctions field to given value.


### GetApplicationAccelerator

`func (o *EdgeApplicationModules) GetApplicationAccelerator() ApplicationAcceleratorModule`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *EdgeApplicationModules) GetApplicationAcceleratorOk() (*ApplicationAcceleratorModule, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *EdgeApplicationModules) SetApplicationAccelerator(v ApplicationAcceleratorModule)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.


### GetImageProcessor

`func (o *EdgeApplicationModules) GetImageProcessor() ImageProcessorModule`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *EdgeApplicationModules) GetImageProcessorOk() (*ImageProcessorModule, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *EdgeApplicationModules) SetImageProcessor(v ImageProcessorModule)`

SetImageProcessor sets ImageProcessor field to given value.


### GetTieredCache

`func (o *EdgeApplicationModules) GetTieredCache() TieredCacheModule`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *EdgeApplicationModules) GetTieredCacheOk() (*TieredCacheModule, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *EdgeApplicationModules) SetTieredCache(v TieredCacheModule)`

SetTieredCache sets TieredCache field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


