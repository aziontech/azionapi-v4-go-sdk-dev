# BigQueryEndpointAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;big_query&#x60; - Google BigQuery | 
**Attributes** | [**BigQueryEndpoint**](BigQueryEndpoint.md) |  | 

## Methods

### NewBigQueryEndpointAttributes

`func NewBigQueryEndpointAttributes(type_ string, attributes BigQueryEndpoint, ) *BigQueryEndpointAttributes`

NewBigQueryEndpointAttributes instantiates a new BigQueryEndpointAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryEndpointAttributesWithDefaults

`func NewBigQueryEndpointAttributesWithDefaults() *BigQueryEndpointAttributes`

NewBigQueryEndpointAttributesWithDefaults instantiates a new BigQueryEndpointAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BigQueryEndpointAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BigQueryEndpointAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BigQueryEndpointAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BigQueryEndpointAttributes) GetAttributes() BigQueryEndpoint`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BigQueryEndpointAttributes) GetAttributesOk() (*BigQueryEndpoint, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BigQueryEndpointAttributes) SetAttributes(v BigQueryEndpoint)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


