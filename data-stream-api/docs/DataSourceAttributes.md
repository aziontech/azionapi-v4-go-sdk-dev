# DataSourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | **string** | * &#x60;http&#x60; - Edge Applications * &#x60;waf&#x60; - WAF Events * &#x60;functions&#x60; - Edge Functions * &#x60;activity&#x60; - Activity History | 

## Methods

### NewDataSourceAttributes

`func NewDataSourceAttributes(dataSource string, ) *DataSourceAttributes`

NewDataSourceAttributes instantiates a new DataSourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceAttributesWithDefaults

`func NewDataSourceAttributesWithDefaults() *DataSourceAttributes`

NewDataSourceAttributesWithDefaults instantiates a new DataSourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *DataSourceAttributes) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataSourceAttributes) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataSourceAttributes) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


