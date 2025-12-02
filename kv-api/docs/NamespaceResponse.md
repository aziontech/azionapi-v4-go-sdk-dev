# NamespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**LastModified** | **NullableTime** |  | 

## Methods

### NewNamespaceResponse

`func NewNamespaceResponse(name string, createdAt NullableTime, lastModified NullableTime, ) *NamespaceResponse`

NewNamespaceResponse instantiates a new NamespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceResponseWithDefaults

`func NewNamespaceResponseWithDefaults() *NamespaceResponse`

NewNamespaceResponseWithDefaults instantiates a new NamespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NamespaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamespaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamespaceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *NamespaceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NamespaceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NamespaceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *NamespaceResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NamespaceResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *NamespaceResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NamespaceResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NamespaceResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *NamespaceResponse) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *NamespaceResponse) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


