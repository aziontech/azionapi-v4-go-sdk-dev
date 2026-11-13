# SupportWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketID** | **string** | Freshdesk ticket identifier | 
**Subject** | **string** | Ticket subject | 
**Description** | **string** | Ticket description/body submitted by the requester | 
**ContactFirstName** | **string** | First name of the ticket&#39;s contact, used to personalize the generated reply | 

## Methods

### NewSupportWebhookRequest

`func NewSupportWebhookRequest(ticketID string, subject string, description string, contactFirstName string, ) *SupportWebhookRequest`

NewSupportWebhookRequest instantiates a new SupportWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportWebhookRequestWithDefaults

`func NewSupportWebhookRequestWithDefaults() *SupportWebhookRequest`

NewSupportWebhookRequestWithDefaults instantiates a new SupportWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketID

`func (o *SupportWebhookRequest) GetTicketID() string`

GetTicketID returns the TicketID field if non-nil, zero value otherwise.

### GetTicketIDOk

`func (o *SupportWebhookRequest) GetTicketIDOk() (*string, bool)`

GetTicketIDOk returns a tuple with the TicketID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketID

`func (o *SupportWebhookRequest) SetTicketID(v string)`

SetTicketID sets TicketID field to given value.


### GetSubject

`func (o *SupportWebhookRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SupportWebhookRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SupportWebhookRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *SupportWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetContactFirstName

`func (o *SupportWebhookRequest) GetContactFirstName() string`

GetContactFirstName returns the ContactFirstName field if non-nil, zero value otherwise.

### GetContactFirstNameOk

`func (o *SupportWebhookRequest) GetContactFirstNameOk() (*string, bool)`

GetContactFirstNameOk returns a tuple with the ContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFirstName

`func (o *SupportWebhookRequest) SetContactFirstName(v string)`

SetContactFirstName sets ContactFirstName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


