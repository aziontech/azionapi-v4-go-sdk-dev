# CertificateVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**Issuer** | **NullableString** |  | 
**SubjectName** | **[]string** |  | 
**Validity** | **NullableString** |  | 
**Type** | Pointer to **string** | The value can&#39;t be changed after the certificate creation.  * &#x60;certificate&#x60; - Certificate * &#x60;trusted_ca_certificate&#x60; - Trusted CA Certificate | [optional] 
**Managed** | **bool** |  | 
**Status** | **string** | * &#x60;pending&#x60; - Pending * &#x60;challenge_verification&#x60; - Challenge Verification * &#x60;active&#x60; - Active * &#x60;inactive&#x60; - Inactive * &#x60;expired&#x60; - Expired * &#x60;failed&#x60; - Failed | 
**StatusDetail** | **string** |  | 
**Csr** | **NullableString** |  | 
**Challenge** | **string** | * &#x60;dns&#x60; - Uses DNS to solve the ACME challenge. * &#x60;http&#x60; - Uses HTTP to solve the ACME challenge. | 
**Authority** | **string** | * &#x60;lets_encrypt&#x60; - lets_encrypt | 
**KeyAlgorithm** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | 
**LastEditor** | **string** |  | 
**CreatedAt** | **NullableTime** | Timestamp of the certificate creation on the platform. | 
**LastModified** | **time.Time** | Timestamp of the last modification made to the certificate content on the platform. | 
**RenewedAt** | **NullableTime** | Timestamp indicating when the managed certificate was renewed on our platform. | 
**SourceVersion** | Pointer to **NullableString** | ID of the version to clone from. If omitted, clones the latest ready version. | [optional] 
**Comment** | Pointer to **string** | Description for the new version. | [optional] 

## Methods

### NewCertificateVersionCreate

`func NewCertificateVersionCreate(id int64, issuer NullableString, subjectName []string, validity NullableString, managed bool, status string, statusDetail string, csr NullableString, challenge string, authority string, keyAlgorithm string, productVersion string, lastEditor string, createdAt NullableTime, lastModified time.Time, renewedAt NullableTime, ) *CertificateVersionCreate`

NewCertificateVersionCreate instantiates a new CertificateVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateVersionCreateWithDefaults

`func NewCertificateVersionCreateWithDefaults() *CertificateVersionCreate`

NewCertificateVersionCreateWithDefaults instantiates a new CertificateVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateVersionCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateVersionCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateVersionCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateVersionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateVersionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateVersionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateVersionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateVersionCreate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateVersionCreate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateVersionCreate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateVersionCreate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CertificateVersionCreate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CertificateVersionCreate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *CertificateVersionCreate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateVersionCreate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateVersionCreate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CertificateVersionCreate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CertificateVersionCreate) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CertificateVersionCreate) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetIssuer

`func (o *CertificateVersionCreate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateVersionCreate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateVersionCreate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### SetIssuerNil

`func (o *CertificateVersionCreate) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *CertificateVersionCreate) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetSubjectName

`func (o *CertificateVersionCreate) GetSubjectName() []string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *CertificateVersionCreate) GetSubjectNameOk() (*[]string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *CertificateVersionCreate) SetSubjectName(v []string)`

SetSubjectName sets SubjectName field to given value.


### GetValidity

`func (o *CertificateVersionCreate) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CertificateVersionCreate) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CertificateVersionCreate) SetValidity(v string)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *CertificateVersionCreate) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *CertificateVersionCreate) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetType

`func (o *CertificateVersionCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateVersionCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateVersionCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateVersionCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetManaged

`func (o *CertificateVersionCreate) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *CertificateVersionCreate) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *CertificateVersionCreate) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetStatus

`func (o *CertificateVersionCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateVersionCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateVersionCreate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetail

`func (o *CertificateVersionCreate) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *CertificateVersionCreate) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *CertificateVersionCreate) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.


### GetCsr

`func (o *CertificateVersionCreate) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *CertificateVersionCreate) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *CertificateVersionCreate) SetCsr(v string)`

SetCsr sets Csr field to given value.


### SetCsrNil

`func (o *CertificateVersionCreate) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *CertificateVersionCreate) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetChallenge

`func (o *CertificateVersionCreate) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CertificateVersionCreate) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CertificateVersionCreate) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetAuthority

`func (o *CertificateVersionCreate) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *CertificateVersionCreate) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *CertificateVersionCreate) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetKeyAlgorithm

`func (o *CertificateVersionCreate) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertificateVersionCreate) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertificateVersionCreate) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetActive

`func (o *CertificateVersionCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CertificateVersionCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CertificateVersionCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CertificateVersionCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *CertificateVersionCreate) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *CertificateVersionCreate) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *CertificateVersionCreate) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetLastEditor

`func (o *CertificateVersionCreate) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *CertificateVersionCreate) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *CertificateVersionCreate) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetCreatedAt

`func (o *CertificateVersionCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CertificateVersionCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CertificateVersionCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CertificateVersionCreate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CertificateVersionCreate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *CertificateVersionCreate) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CertificateVersionCreate) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CertificateVersionCreate) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetRenewedAt

`func (o *CertificateVersionCreate) GetRenewedAt() time.Time`

GetRenewedAt returns the RenewedAt field if non-nil, zero value otherwise.

### GetRenewedAtOk

`func (o *CertificateVersionCreate) GetRenewedAtOk() (*time.Time, bool)`

GetRenewedAtOk returns a tuple with the RenewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewedAt

`func (o *CertificateVersionCreate) SetRenewedAt(v time.Time)`

SetRenewedAt sets RenewedAt field to given value.


### SetRenewedAtNil

`func (o *CertificateVersionCreate) SetRenewedAtNil(b bool)`

 SetRenewedAtNil sets the value for RenewedAt to be an explicit nil

### UnsetRenewedAt
`func (o *CertificateVersionCreate) UnsetRenewedAt()`

UnsetRenewedAt ensures that no value is present for RenewedAt, not even an explicit nil
### GetSourceVersion

`func (o *CertificateVersionCreate) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CertificateVersionCreate) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CertificateVersionCreate) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *CertificateVersionCreate) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

### SetSourceVersionNil

`func (o *CertificateVersionCreate) SetSourceVersionNil(b bool)`

 SetSourceVersionNil sets the value for SourceVersion to be an explicit nil

### UnsetSourceVersion
`func (o *CertificateVersionCreate) UnsetSourceVersion()`

UnsetSourceVersion ensures that no value is present for SourceVersion, not even an explicit nil
### GetComment

`func (o *CertificateVersionCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CertificateVersionCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CertificateVersionCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CertificateVersionCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


