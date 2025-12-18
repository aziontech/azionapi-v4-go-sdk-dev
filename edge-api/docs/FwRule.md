# FwRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]FwCriterionField**]([]FwCriterionField.md) |  | 
**Behaviors** | [**[]FwBehaviors**](FwBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 

## Methods

### NewFwRule

`func NewFwRule(id int64, name string, lastEditor string, lastModified time.Time, criteria [][]FwCriterionField, behaviors []FwBehaviors, order int64, ) *FwRule`

NewFwRule instantiates a new FwRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwRuleWithDefaults

`func NewFwRuleWithDefaults() *FwRule`

NewFwRuleWithDefaults instantiates a new FwRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FwRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FwRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FwRule) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *FwRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FwRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FwRule) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *FwRule) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *FwRule) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *FwRule) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *FwRule) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FwRule) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FwRule) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *FwRule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FwRule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FwRule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FwRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *FwRule) GetCriteria() [][]FwCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *FwRule) GetCriteriaOk() (*[][]FwCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *FwRule) SetCriteria(v [][]FwCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *FwRule) GetBehaviors() []FwBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *FwRule) GetBehaviorsOk() (*[]FwBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *FwRule) SetBehaviors(v []FwBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *FwRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FwRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FwRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FwRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *FwRule) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FwRule) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FwRule) SetOrder(v int64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


