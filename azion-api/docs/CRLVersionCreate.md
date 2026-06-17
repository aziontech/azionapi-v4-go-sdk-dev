# CRLVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Indicates if the certificate revocation list is active. This field cannot be set to false. | [optional] 
**LastEditor** | **string** |  | 
**CreatedAt** | **NullableTime** | Timestamp of the certificate revocation list creation on the platform. | 
**LastModified** | **time.Time** | Timestamp of the last modification made to the certificate content on the platform. | 
**ProductVersion** | **string** |  | 
**Issuer** | **string** |  | 
**LastUpdate** | **time.Time** | Timestamp of the last update issued by the certification revocation list issuer. | 
**NextUpdate** | **time.Time** | Timestamp of the next scheduled update from the certification revocation list issuer. | 
**Crl** | Pointer to **string** |  | [optional] 
**VersionId** | **NullableString** | ID of the version metadata (use in /versions/{id} URLs) | 
**State** | **NullableString** | Build state of this version (queued, building, ready, error, ...) | 
**SourceVersion** | Pointer to **NullableString** | ID of the version to clone from. If omitted, clones the latest ready version. | [optional] 
**Comment** | Pointer to **string** | Description for the new version. | [optional] 

## Methods

### NewCRLVersionCreate

`func NewCRLVersionCreate(id int64, lastEditor string, createdAt NullableTime, lastModified time.Time, productVersion string, issuer string, lastUpdate time.Time, nextUpdate time.Time, versionId NullableString, state NullableString, ) *CRLVersionCreate`

NewCRLVersionCreate instantiates a new CRLVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCRLVersionCreateWithDefaults

`func NewCRLVersionCreateWithDefaults() *CRLVersionCreate`

NewCRLVersionCreateWithDefaults instantiates a new CRLVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CRLVersionCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CRLVersionCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CRLVersionCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CRLVersionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CRLVersionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CRLVersionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CRLVersionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *CRLVersionCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CRLVersionCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CRLVersionCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CRLVersionCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *CRLVersionCreate) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *CRLVersionCreate) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *CRLVersionCreate) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetCreatedAt

`func (o *CRLVersionCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CRLVersionCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CRLVersionCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CRLVersionCreate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CRLVersionCreate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *CRLVersionCreate) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CRLVersionCreate) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CRLVersionCreate) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *CRLVersionCreate) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *CRLVersionCreate) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *CRLVersionCreate) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetIssuer

`func (o *CRLVersionCreate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CRLVersionCreate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CRLVersionCreate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetLastUpdate

`func (o *CRLVersionCreate) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *CRLVersionCreate) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *CRLVersionCreate) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.


### GetNextUpdate

`func (o *CRLVersionCreate) GetNextUpdate() time.Time`

GetNextUpdate returns the NextUpdate field if non-nil, zero value otherwise.

### GetNextUpdateOk

`func (o *CRLVersionCreate) GetNextUpdateOk() (*time.Time, bool)`

GetNextUpdateOk returns a tuple with the NextUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdate

`func (o *CRLVersionCreate) SetNextUpdate(v time.Time)`

SetNextUpdate sets NextUpdate field to given value.


### GetCrl

`func (o *CRLVersionCreate) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *CRLVersionCreate) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *CRLVersionCreate) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *CRLVersionCreate) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### GetVersionId

`func (o *CRLVersionCreate) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *CRLVersionCreate) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *CRLVersionCreate) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### SetVersionIdNil

`func (o *CRLVersionCreate) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *CRLVersionCreate) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil
### GetState

`func (o *CRLVersionCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CRLVersionCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CRLVersionCreate) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *CRLVersionCreate) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CRLVersionCreate) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSourceVersion

`func (o *CRLVersionCreate) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CRLVersionCreate) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CRLVersionCreate) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *CRLVersionCreate) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

### SetSourceVersionNil

`func (o *CRLVersionCreate) SetSourceVersionNil(b bool)`

 SetSourceVersionNil sets the value for SourceVersion to be an explicit nil

### UnsetSourceVersion
`func (o *CRLVersionCreate) UnsetSourceVersion()`

UnsetSourceVersion ensures that no value is present for SourceVersion, not even an explicit nil
### GetComment

`func (o *CRLVersionCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CRLVersionCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CRLVersionCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CRLVersionCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


