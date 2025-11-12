# ResponseBadRequestGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to [**ResponseBadRequestGrantRoles**](ResponseBadRequestGrantRoles.md) |  | [optional] 
**Policies** | Pointer to [**ResponseBadRequestGrantRoles**](ResponseBadRequestGrantRoles.md) |  | [optional] 
**Expires** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **[]string** |  | [optional] 
**AccountId** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestGrant

`func NewResponseBadRequestGrant() *ResponseBadRequestGrant`

NewResponseBadRequestGrant instantiates a new ResponseBadRequestGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestGrantWithDefaults

`func NewResponseBadRequestGrantWithDefaults() *ResponseBadRequestGrant`

NewResponseBadRequestGrantWithDefaults instantiates a new ResponseBadRequestGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *ResponseBadRequestGrant) GetIdentity() []string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ResponseBadRequestGrant) GetIdentityOk() (*[]string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ResponseBadRequestGrant) SetIdentity(v []string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ResponseBadRequestGrant) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetRoles

`func (o *ResponseBadRequestGrant) GetRoles() ResponseBadRequestGrantRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ResponseBadRequestGrant) GetRolesOk() (*ResponseBadRequestGrantRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ResponseBadRequestGrant) SetRoles(v ResponseBadRequestGrantRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ResponseBadRequestGrant) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPolicies

`func (o *ResponseBadRequestGrant) GetPolicies() ResponseBadRequestGrantRoles`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ResponseBadRequestGrant) GetPoliciesOk() (*ResponseBadRequestGrantRoles, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ResponseBadRequestGrant) SetPolicies(v ResponseBadRequestGrantRoles)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ResponseBadRequestGrant) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetExpires

`func (o *ResponseBadRequestGrant) GetExpires() []string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ResponseBadRequestGrant) GetExpiresOk() (*[]string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ResponseBadRequestGrant) SetExpires(v []string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ResponseBadRequestGrant) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetOwner

`func (o *ResponseBadRequestGrant) GetOwner() []string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResponseBadRequestGrant) GetOwnerOk() (*[]string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResponseBadRequestGrant) SetOwner(v []string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ResponseBadRequestGrant) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *ResponseBadRequestGrant) GetAccountId() []string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResponseBadRequestGrant) GetAccountIdOk() (*[]string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResponseBadRequestGrant) SetAccountId(v []string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResponseBadRequestGrant) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestGrant) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestGrant) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestGrant) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestGrant) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


