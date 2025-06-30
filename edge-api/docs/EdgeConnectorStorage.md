# EdgeConnectorStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Modules** | [**EdgeConnectorModules**](EdgeConnectorModules.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 
**Type** | **string** | * &#x60;http&#x60; - HTTP * &#x60;s3&#x60; - S3 * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Tls** | Pointer to [**TLSEdgeConnector**](TLSEdgeConnector.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | [**EdgeConnectorStorageTypeProperties**](EdgeConnectorStorageTypeProperties.md) |  | 

## Methods

### NewEdgeConnectorStorage

`func NewEdgeConnectorStorage(id int64, name string, lastEditor string, lastModified time.Time, modules EdgeConnectorModules, productVersion string, type_ string, typeProperties EdgeConnectorStorageTypeProperties, ) *EdgeConnectorStorage`

NewEdgeConnectorStorage instantiates a new EdgeConnectorStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorStorageWithDefaults

`func NewEdgeConnectorStorageWithDefaults() *EdgeConnectorStorage`

NewEdgeConnectorStorageWithDefaults instantiates a new EdgeConnectorStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeConnectorStorage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeConnectorStorage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeConnectorStorage) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeConnectorStorage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorStorage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorStorage) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeConnectorStorage) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeConnectorStorage) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeConnectorStorage) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeConnectorStorage) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeConnectorStorage) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeConnectorStorage) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetModules

`func (o *EdgeConnectorStorage) GetModules() EdgeConnectorModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorStorage) GetModulesOk() (*EdgeConnectorModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorStorage) SetModules(v EdgeConnectorModules)`

SetModules sets Modules field to given value.


### GetActive

`func (o *EdgeConnectorStorage) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorStorage) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorStorage) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorStorage) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *EdgeConnectorStorage) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeConnectorStorage) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeConnectorStorage) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *EdgeConnectorStorage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorStorage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorStorage) SetType(v string)`

SetType sets Type field to given value.


### GetTls

`func (o *EdgeConnectorStorage) GetTls() TLSEdgeConnector`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *EdgeConnectorStorage) GetTlsOk() (*TLSEdgeConnector, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *EdgeConnectorStorage) SetTls(v TLSEdgeConnector)`

SetTls sets Tls field to given value.

### HasTls

`func (o *EdgeConnectorStorage) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *EdgeConnectorStorage) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *EdgeConnectorStorage) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *EdgeConnectorStorage) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *EdgeConnectorStorage) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *EdgeConnectorStorage) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *EdgeConnectorStorage) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *EdgeConnectorStorage) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *EdgeConnectorStorage) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EdgeConnectorStorage) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EdgeConnectorStorage) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EdgeConnectorStorage) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EdgeConnectorStorage) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *EdgeConnectorStorage) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *EdgeConnectorStorage) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *EdgeConnectorStorage) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *EdgeConnectorStorage) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EdgeConnectorStorage) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EdgeConnectorStorage) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EdgeConnectorStorage) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EdgeConnectorStorage) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *EdgeConnectorStorage) GetTypeProperties() EdgeConnectorStorageTypeProperties`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *EdgeConnectorStorage) GetTypePropertiesOk() (*EdgeConnectorStorageTypeProperties, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *EdgeConnectorStorage) SetTypeProperties(v EdgeConnectorStorageTypeProperties)`

SetTypeProperties sets TypeProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


