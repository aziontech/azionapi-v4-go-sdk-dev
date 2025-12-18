# PatchCertificateRevocationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Indicates if the certificate revocation list is active. This field cannot be set to false. | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** | Timestamp of the last modification made to the certificate content on the platform. | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** | Timestamp of the last update issued by the certification revocation list issuer. | [optional] 
**NextUpdate** | Pointer to **time.Time** | Timestamp of the next scheduled update from the certification revocation list issuer. | [optional] 
**Crl** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchCertificateRevocationList

`func NewPatchCertificateRevocationList() *PatchCertificateRevocationList`

NewPatchCertificateRevocationList instantiates a new PatchCertificateRevocationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCertificateRevocationListWithDefaults

`func NewPatchCertificateRevocationListWithDefaults() *PatchCertificateRevocationList`

NewPatchCertificateRevocationListWithDefaults instantiates a new PatchCertificateRevocationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchCertificateRevocationList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchCertificateRevocationList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchCertificateRevocationList) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PatchCertificateRevocationList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchCertificateRevocationList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchCertificateRevocationList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchCertificateRevocationList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchCertificateRevocationList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchCertificateRevocationList) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchCertificateRevocationList) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchCertificateRevocationList) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchCertificateRevocationList) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *PatchCertificateRevocationList) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *PatchCertificateRevocationList) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *PatchCertificateRevocationList) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *PatchCertificateRevocationList) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *PatchCertificateRevocationList) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PatchCertificateRevocationList) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PatchCertificateRevocationList) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *PatchCertificateRevocationList) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetProductVersion

`func (o *PatchCertificateRevocationList) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *PatchCertificateRevocationList) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *PatchCertificateRevocationList) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *PatchCertificateRevocationList) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetIssuer

`func (o *PatchCertificateRevocationList) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PatchCertificateRevocationList) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PatchCertificateRevocationList) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PatchCertificateRevocationList) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLastUpdate

`func (o *PatchCertificateRevocationList) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *PatchCertificateRevocationList) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *PatchCertificateRevocationList) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *PatchCertificateRevocationList) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetNextUpdate

`func (o *PatchCertificateRevocationList) GetNextUpdate() time.Time`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *PatchCertificateRevocationList) GetNextUpdateOk() (*time.Time, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *PatchCertificateRevocationList) SetNextUpdate(v time.Time)`

SetNextUpdate sets NextUpdate field to given value.

### HasNextUpdate

`func (o *PatchCertificateRevocationList) HasNextUpdate() bool`

HasNextUpdate returns a boolean if a field has been set.

### GetCrl

`func (o *PatchCertificateRevocationList) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *PatchCertificateRevocationList) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *PatchCertificateRevocationList) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *PatchCertificateRevocationList) HasCrl() bool`

HasCrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


