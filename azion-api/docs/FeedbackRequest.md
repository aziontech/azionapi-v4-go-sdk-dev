# FeedbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feedback** | **bool** | Whether the feedback is positive (true) or negative (false) | 
**Comments** | Pointer to **string** | Optional free-text comment providing more detail | [optional] 
**RunId** | **string** | The LangSmith run ID the feedback should be attached to | 

## Methods

### NewFeedbackRequest

`func NewFeedbackRequest(feedback bool, runId string, ) *FeedbackRequest`

NewFeedbackRequest instantiates a new FeedbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackRequestWithDefaults

`func NewFeedbackRequestWithDefaults() *FeedbackRequest`

NewFeedbackRequestWithDefaults instantiates a new FeedbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedback

`func (o *FeedbackRequest) GetFeedback() bool`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *FeedbackRequest) GetFeedbackOk() (*bool, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *FeedbackRequest) SetFeedback(v bool)`

SetFeedback sets Feedback field to given value.


### GetComments

`func (o *FeedbackRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *FeedbackRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *FeedbackRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *FeedbackRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetRunId

`func (o *FeedbackRequest) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *FeedbackRequest) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *FeedbackRequest) SetRunId(v string)`

SetRunId sets RunId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


