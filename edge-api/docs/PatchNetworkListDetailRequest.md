# PatchNetworkListDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | * &#x60;asn&#x60; - ASN * &#x60;countries&#x60; - Countries * &#x60;ip_cidr&#x60; - IP/CIDR | [optional] 
**Items** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchNetworkListDetailRequest

`func NewPatchNetworkListDetailRequest() *PatchNetworkListDetailRequest`

NewPatchNetworkListDetailRequest instantiates a new PatchNetworkListDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNetworkListDetailRequestWithDefaults

`func NewPatchNetworkListDetailRequestWithDefaults() *PatchNetworkListDetailRequest`

NewPatchNetworkListDetailRequestWithDefaults instantiates a new PatchNetworkListDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchNetworkListDetailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchNetworkListDetailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchNetworkListDetailRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchNetworkListDetailRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchNetworkListDetailRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchNetworkListDetailRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchNetworkListDetailRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchNetworkListDetailRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *PatchNetworkListDetailRequest) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PatchNetworkListDetailRequest) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PatchNetworkListDetailRequest) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *PatchNetworkListDetailRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetActive

`func (o *PatchNetworkListDetailRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchNetworkListDetailRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchNetworkListDetailRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchNetworkListDetailRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


