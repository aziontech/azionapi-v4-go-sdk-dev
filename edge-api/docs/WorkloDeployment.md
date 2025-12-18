# WorkloDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Current** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Strategy** | [**DeployStrategyDefaultDeployStrategy**](DeployStrategyDefaultDeployStrategy.md) |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 

## Methods

### NewWorkloDeployment

`func NewWorkloDeployment(id int64, name string, strategy DeployStrategyDefaultDeployStrategy, lastEditor string, lastModified time.Time, ) *WorkloDeployment`

NewWorkloDeployment instantiates a new WorkloDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloDeploymentWithDefaults

`func NewWorkloDeploymentWithDefaults() *WorkloDeployment`

NewWorkloDeploymentWithDefaults instantiates a new WorkloDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkloDeployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloDeployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloDeployment) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *WorkloDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetCurrent

`func (o *WorkloDeployment) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkloDeployment) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkloDeployment) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkloDeployment) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetActive

`func (o *WorkloDeployment) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkloDeployment) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkloDeployment) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WorkloDeployment) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrategy

`func (o *WorkloDeployment) GetStrategy() DeployStrategyDefaultDeployStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *WorkloDeployment) GetStrategyOk() (*DeployStrategyDefaultDeployStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *WorkloDeployment) SetStrategy(v DeployStrategyDefaultDeployStrategy)`

SetStrategy sets Strategy field to given value.


### GetLastEditor

`func (o *WorkloDeployment) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *WorkloDeployment) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *WorkloDeployment) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *WorkloDeployment) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WorkloDeployment) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WorkloDeployment) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


