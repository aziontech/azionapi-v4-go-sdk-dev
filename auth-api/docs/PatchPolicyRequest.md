# PatchPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]PolicyRuleRequest**](PolicyRuleRequest.md) |  | [optional] 

## Methods

### NewPatchPolicyRequest

`func NewPatchPolicyRequest() *PatchPolicyRequest`

NewPatchPolicyRequest instantiates a new PatchPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyRequestWithDefaults

`func NewPatchPolicyRequestWithDefaults() *PatchPolicyRequest`

NewPatchPolicyRequestWithDefaults instantiates a new PatchPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchPolicyRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchPolicyRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchPolicyRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchPolicyRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRules

`func (o *PatchPolicyRequest) GetRules() []PolicyRuleRequest`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchPolicyRequest) GetRulesOk() (*[]PolicyRuleRequest, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchPolicyRequest) SetRules(v []PolicyRuleRequest)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchPolicyRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


