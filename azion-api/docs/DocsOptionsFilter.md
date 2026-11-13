# DocsOptionsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Comparison operator applied to the filter | 
**Column** | **string** | Column the filter is applied to | 
**Value** | **string** | Value compared against the column | 

## Methods

### NewDocsOptionsFilter

`func NewDocsOptionsFilter(operator string, column string, value string, ) *DocsOptionsFilter`

NewDocsOptionsFilter instantiates a new DocsOptionsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsOptionsFilterWithDefaults

`func NewDocsOptionsFilterWithDefaults() *DocsOptionsFilter`

NewDocsOptionsFilterWithDefaults instantiates a new DocsOptionsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *DocsOptionsFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DocsOptionsFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DocsOptionsFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetColumn

`func (o *DocsOptionsFilter) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *DocsOptionsFilter) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *DocsOptionsFilter) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetValue

`func (o *DocsOptionsFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DocsOptionsFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DocsOptionsFilter) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


