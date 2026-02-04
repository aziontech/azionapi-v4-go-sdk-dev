# RequestPhaseRule2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]ApplicationCriterionFieldRequest**]([]ApplicationCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]RequestPhaseBehavior2**](RequestPhaseBehavior2.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestPhaseRule2

`func NewRequestPhaseRule2(name string, criteria [][]ApplicationCriterionFieldRequest, behaviors []RequestPhaseBehavior2, ) *RequestPhaseRule2`

NewRequestPhaseRule2 instantiates a new RequestPhaseRule2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPhaseRule2WithDefaults

`func NewRequestPhaseRule2WithDefaults() *RequestPhaseRule2`

NewRequestPhaseRule2WithDefaults instantiates a new RequestPhaseRule2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestPhaseRule2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestPhaseRule2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestPhaseRule2) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *RequestPhaseRule2) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RequestPhaseRule2) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RequestPhaseRule2) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RequestPhaseRule2) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *RequestPhaseRule2) GetCriteria() [][]ApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RequestPhaseRule2) GetCriteriaOk() (*[][]ApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RequestPhaseRule2) SetCriteria(v [][]ApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *RequestPhaseRule2) GetBehaviors() []RequestPhaseBehavior2`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *RequestPhaseRule2) GetBehaviorsOk() (*[]RequestPhaseBehavior2, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *RequestPhaseRule2) SetBehaviors(v []RequestPhaseBehavior2)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *RequestPhaseRule2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestPhaseRule2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestPhaseRule2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestPhaseRule2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


