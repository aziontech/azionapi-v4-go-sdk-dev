# Transform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;render_template&#x60; - Render Template * &#x60;sampling&#x60; - Sampling * &#x60;filter_workloads&#x60; - Filter Workloads | 
**Attributes** | [**TransformAttributesPolymorphic**](TransformAttributesPolymorphic.md) |  | 

## Methods

### NewTransform

`func NewTransform(type_ string, attributes TransformAttributesPolymorphic, ) *Transform`

NewTransform instantiates a new Transform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformWithDefaults

`func NewTransformWithDefaults() *Transform`

NewTransformWithDefaults instantiates a new Transform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transform) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Transform) GetAttributes() TransformAttributesPolymorphic`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Transform) GetAttributesOk() (*TransformAttributesPolymorphic, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Transform) SetAttributes(v TransformAttributesPolymorphic)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


