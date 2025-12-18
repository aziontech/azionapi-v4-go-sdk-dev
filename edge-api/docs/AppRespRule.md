# AppRespRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]AppCriterionField**]([]AppCriterionField.md) |  | 
**Behaviors** | [**[]RespBehaviors**](RespBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 
**LastEditor** | **NullableString** |  | 
**LastModified** | **NullableTime** |  | 

## Methods

### NewAppRespRule

`func NewAppRespRule(id int64, name string, criteria [][]AppCriterionField, behaviors []RespBehaviors, order int64, lastEditor NullableString, lastModified NullableTime, ) *AppRespRule`

NewAppRespRule instantiates a new AppRespRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRespRuleWithDefaults

`func NewAppRespRuleWithDefaults() *AppRespRule`

NewAppRespRuleWithDefaults instantiates a new AppRespRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRespRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRespRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRespRule) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AppRespRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRespRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRespRule) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *AppRespRule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AppRespRule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AppRespRule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AppRespRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *AppRespRule) GetCriteria() [][]AppCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AppRespRule) GetCriteriaOk() (*[][]AppCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AppRespRule) SetCriteria(v [][]AppCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *AppRespRule) GetBehaviors() []RespBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *AppRespRule) GetBehaviorsOk() (*[]RespBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *AppRespRule) SetBehaviors(v []RespBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *AppRespRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRespRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRespRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRespRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *AppRespRule) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AppRespRule) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AppRespRule) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetLastEditor

`func (o *AppRespRule) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *AppRespRule) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *AppRespRule) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### SetLastEditorNil

`func (o *AppRespRule) SetLastEditorNil(b bool)`

 SetLastEditorNil sets the value for LastEditor to be an explicit nil

### UnsetLastEditor
`func (o *AppRespRule) UnsetLastEditor()`

UnsetLastEditor ensures that no value is present for LastEditor, not even an explicit nil
### GetLastModified

`func (o *AppRespRule) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AppRespRule) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AppRespRule) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *AppRespRule) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AppRespRule) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


