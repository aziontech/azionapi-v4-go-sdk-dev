# ResponseBadRequestGroupMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**ResponseBadRequestGrantRoles**](ResponseBadRequestGrantRoles.md) |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestGroupMembers

`func NewResponseBadRequestGroupMembers() *ResponseBadRequestGroupMembers`

NewResponseBadRequestGroupMembers instantiates a new ResponseBadRequestGroupMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestGroupMembersWithDefaults

`func NewResponseBadRequestGroupMembersWithDefaults() *ResponseBadRequestGroupMembers`

NewResponseBadRequestGroupMembersWithDefaults instantiates a new ResponseBadRequestGroupMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ResponseBadRequestGroupMembers) GetMembers() ResponseBadRequestGrantRoles`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResponseBadRequestGroupMembers) GetMembersOk() (*ResponseBadRequestGrantRoles, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResponseBadRequestGroupMembers) SetMembers(v ResponseBadRequestGrantRoles)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ResponseBadRequestGroupMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestGroupMembers) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestGroupMembers) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestGroupMembers) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestGroupMembers) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


