# Azion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | A unique session identifier | [optional] 
**Url** | Pointer to **string** | URL related to the request | [optional] 
**App** | Pointer to **string** | The application making the request | [optional] 
**UserName** | Pointer to **string** | The name of the user | [optional] 
**ClientId** | Pointer to **string** | A client identifier | [optional] 
**SystemPrompt** | Pointer to **string** | A system-level prompt | [optional] 
**UserPrompt** | Pointer to **string** | A user-level prompt | [optional] 
**Id** | Pointer to **int64** | Numeric identifier of the user | [optional] 
**FirstName** | Pointer to **string** | The first name of the user | [optional] 
**LastName** | Pointer to **string** | The last name of the user | [optional] 
**Email** | Pointer to **string** | The email of the user | [optional] 
**SupportPlan** | Pointer to **string** | The user&#39;s support plan | [optional] 

## Methods

### NewAzion

`func NewAzion() *Azion`

NewAzion instantiates a new Azion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzionWithDefaults

`func NewAzionWithDefaults() *Azion`

NewAzionWithDefaults instantiates a new Azion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *Azion) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Azion) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Azion) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Azion) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUrl

`func (o *Azion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Azion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Azion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Azion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApp

`func (o *Azion) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Azion) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Azion) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *Azion) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetUserName

`func (o *Azion) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Azion) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Azion) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Azion) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetClientId

`func (o *Azion) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Azion) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Azion) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Azion) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *Azion) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *Azion) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *Azion) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *Azion) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetUserPrompt

`func (o *Azion) GetUserPrompt() string`

GetUserPrompt returns the UserPrompt field if non-nil, zero value otherwise.

### GetUserPromptOk

`func (o *Azion) GetUserPromptOk() (*string, bool)`

GetUserPromptOk returns a tuple with the UserPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrompt

`func (o *Azion) SetUserPrompt(v string)`

SetUserPrompt sets UserPrompt field to given value.

### HasUserPrompt

`func (o *Azion) HasUserPrompt() bool`

HasUserPrompt returns a boolean if a field has been set.

### GetId

`func (o *Azion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Azion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Azion) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Azion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *Azion) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Azion) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Azion) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Azion) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Azion) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Azion) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Azion) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Azion) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Azion) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Azion) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Azion) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Azion) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSupportPlan

`func (o *Azion) GetSupportPlan() string`

GetSupportPlan returns the SupportPlan field if non-nil, zero value otherwise.

### GetSupportPlanOk

`func (o *Azion) GetSupportPlanOk() (*string, bool)`

GetSupportPlanOk returns a tuple with the SupportPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPlan

`func (o *Azion) SetSupportPlan(v string)`

SetSupportPlan sets SupportPlan field to given value.

### HasSupportPlan

`func (o *Azion) HasSupportPlan() bool`

HasSupportPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


