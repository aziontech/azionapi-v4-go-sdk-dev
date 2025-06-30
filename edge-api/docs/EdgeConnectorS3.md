# EdgeConnectorS3

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
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**Tls** | Pointer to [**TLSEdgeConnector**](TLSEdgeConnector.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | [**EdgeConnectorS3TypeProperties**](EdgeConnectorS3TypeProperties.md) |  | 

## Methods

### NewEdgeConnectorS3

`func NewEdgeConnectorS3(id int64, name string, lastEditor string, lastModified time.Time, modules EdgeConnectorModules, productVersion string, type_ string, typeProperties EdgeConnectorS3TypeProperties, ) *EdgeConnectorS3`

NewEdgeConnectorS3 instantiates a new EdgeConnectorS3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorS3WithDefaults

`func NewEdgeConnectorS3WithDefaults() *EdgeConnectorS3`

NewEdgeConnectorS3WithDefaults instantiates a new EdgeConnectorS3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeConnectorS3) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeConnectorS3) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeConnectorS3) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeConnectorS3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorS3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorS3) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeConnectorS3) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeConnectorS3) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeConnectorS3) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeConnectorS3) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeConnectorS3) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeConnectorS3) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetModules

`func (o *EdgeConnectorS3) GetModules() EdgeConnectorModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorS3) GetModulesOk() (*EdgeConnectorModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorS3) SetModules(v EdgeConnectorModules)`

SetModules sets Modules field to given value.


### GetActive

`func (o *EdgeConnectorS3) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorS3) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorS3) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorS3) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *EdgeConnectorS3) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeConnectorS3) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeConnectorS3) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *EdgeConnectorS3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorS3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorS3) SetType(v string)`

SetType sets Type field to given value.


### GetAddresses

`func (o *EdgeConnectorS3) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorS3) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorS3) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *EdgeConnectorS3) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *EdgeConnectorS3) GetTls() TLSEdgeConnector`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *EdgeConnectorS3) GetTlsOk() (*TLSEdgeConnector, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *EdgeConnectorS3) SetTls(v TLSEdgeConnector)`

SetTls sets Tls field to given value.

### HasTls

`func (o *EdgeConnectorS3) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *EdgeConnectorS3) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *EdgeConnectorS3) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *EdgeConnectorS3) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *EdgeConnectorS3) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *EdgeConnectorS3) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *EdgeConnectorS3) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *EdgeConnectorS3) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *EdgeConnectorS3) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EdgeConnectorS3) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EdgeConnectorS3) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EdgeConnectorS3) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EdgeConnectorS3) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *EdgeConnectorS3) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *EdgeConnectorS3) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *EdgeConnectorS3) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *EdgeConnectorS3) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EdgeConnectorS3) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EdgeConnectorS3) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EdgeConnectorS3) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EdgeConnectorS3) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *EdgeConnectorS3) GetTypeProperties() EdgeConnectorS3TypeProperties`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *EdgeConnectorS3) GetTypePropertiesOk() (*EdgeConnectorS3TypeProperties, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *EdgeConnectorS3) SetTypeProperties(v EdgeConnectorS3TypeProperties)`

SetTypeProperties sets TypeProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


