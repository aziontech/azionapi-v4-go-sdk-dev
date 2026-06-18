# OutputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**LogLineSeparator** | Pointer to **string** |  | [optional] 
**PayloadFormat** | Pointer to **string** |  | [optional] 
**MaxSize** | Pointer to **NullableInt64** |  | [optional] 
**Headers** | **map[string]string** |  | 
**Type** | **string** | Type identifier for this endpoint (splunk) | 
**BootstrapServers** | **string** |  | 
**KafkaTopic** | **string** |  | 
**UseTls** | **bool** |  | 
**AccessKey** | **string** |  | 
**SecretKey** | **string** |  | 
**Region** | **string** |  | 
**ObjectKeyPrefix** | Pointer to **NullableString** |  | [optional] 
**BucketName** | **string** |  | 
**ContentType** | **string** | * &#x60;plain/text&#x60; - plain/text * &#x60;application/gzip&#x60; - application/gzip | 
**HostUrl** | **string** |  | 
**DatasetId** | **string** |  | 
**ProjectId** | **string** |  | 
**TableId** | **string** |  | 
**ServiceAccountKey** | **string** |  | 
**ApiKey** | **string** |  | 
**StreamName** | **string** |  | 
**LogType** | **string** |  | 
**SharedKey** | **string** |  | 
**TimeGeneratedField** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | **string** |  | 
**StorageAccount** | **string** |  | 
**ContainerName** | **string** |  | 
**BlobSasToken** | **string** |  | 

## Methods

### NewOutputRequest

`func NewOutputRequest(url string, headers map[string]string, type_ string, bootstrapServers string, kafkaTopic string, useTls bool, accessKey string, secretKey string, region string, bucketName string, contentType string, hostUrl string, datasetId string, projectId string, tableId string, serviceAccountKey string, apiKey string, streamName string, logType string, sharedKey string, workspaceId string, storageAccount string, containerName string, blobSasToken string, ) *OutputRequest`

NewOutputRequest instantiates a new OutputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRequestWithDefaults

`func NewOutputRequestWithDefaults() *OutputRequest`

NewOutputRequestWithDefaults instantiates a new OutputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OutputRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLogLineSeparator

`func (o *OutputRequest) GetLogLineSeparator() string`

GetLogLineSeparator returns the LogLineSeparator field if non-nil, zero value otherwise.

### GetLogLineSeparatorOk

`func (o *OutputRequest) GetLogLineSeparatorOk() (*string, bool)`

GetLogLineSeparatorOk returns a tuple with the LogLineSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLineSeparator

`func (o *OutputRequest) SetLogLineSeparator(v string)`

SetLogLineSeparator sets LogLineSeparator field to given value.

### HasLogLineSeparator

`func (o *OutputRequest) HasLogLineSeparator() bool`

HasLogLineSeparator returns a boolean if a field has been set.

### GetPayloadFormat

`func (o *OutputRequest) GetPayloadFormat() string`

GetPayloadFormat returns the PayloadFormat field if non-nil, zero value otherwise.

### GetPayloadFormatOk

`func (o *OutputRequest) GetPayloadFormatOk() (*string, bool)`

GetPayloadFormatOk returns a tuple with the PayloadFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFormat

`func (o *OutputRequest) SetPayloadFormat(v string)`

SetPayloadFormat sets PayloadFormat field to given value.

### HasPayloadFormat

`func (o *OutputRequest) HasPayloadFormat() bool`

HasPayloadFormat returns a boolean if a field has been set.

### GetMaxSize

`func (o *OutputRequest) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *OutputRequest) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *OutputRequest) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *OutputRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *OutputRequest) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *OutputRequest) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetHeaders

`func (o *OutputRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutputRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutputRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetType

`func (o *OutputRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputRequest) SetType(v string)`

SetType sets Type field to given value.


### GetBootstrapServers

`func (o *OutputRequest) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *OutputRequest) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *OutputRequest) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.


### GetKafkaTopic

`func (o *OutputRequest) GetKafkaTopic() string`

GetKafkaTopic returns the KafkaTopic field if non-nil, zero value otherwise.

### GetKafkaTopicOk

`func (o *OutputRequest) GetKafkaTopicOk() (*string, bool)`

GetKafkaTopicOk returns a tuple with the KafkaTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopic

`func (o *OutputRequest) SetKafkaTopic(v string)`

SetKafkaTopic sets KafkaTopic field to given value.


### GetUseTls

`func (o *OutputRequest) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *OutputRequest) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *OutputRequest) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.


### GetAccessKey

`func (o *OutputRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OutputRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OutputRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *OutputRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OutputRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OutputRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetRegion

`func (o *OutputRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OutputRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OutputRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetObjectKeyPrefix

`func (o *OutputRequest) GetObjectKeyPrefix() string`

GetObjectKeyPrefix returns the ObjectKeyPrefix field if non-nil, zero value otherwise.

### GetObjectKeyPrefixOk

`func (o *OutputRequest) GetObjectKeyPrefixOk() (*string, bool)`

GetObjectKeyPrefixOk returns a tuple with the ObjectKeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKeyPrefix

`func (o *OutputRequest) SetObjectKeyPrefix(v string)`

SetObjectKeyPrefix sets ObjectKeyPrefix field to given value.

### HasObjectKeyPrefix

`func (o *OutputRequest) HasObjectKeyPrefix() bool`

HasObjectKeyPrefix returns a boolean if a field has been set.

### SetObjectKeyPrefixNil

`func (o *OutputRequest) SetObjectKeyPrefixNil(b bool)`

 SetObjectKeyPrefixNil sets the value for ObjectKeyPrefix to be an explicit nil

### UnsetObjectKeyPrefix
`func (o *OutputRequest) UnsetObjectKeyPrefix()`

UnsetObjectKeyPrefix ensures that no value is present for ObjectKeyPrefix, not even an explicit nil
### GetBucketName

`func (o *OutputRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *OutputRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *OutputRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetContentType

`func (o *OutputRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OutputRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OutputRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetHostUrl

`func (o *OutputRequest) GetHostUrl() string`

GetHostUrl returns the HostUrl field if non-nil, zero value otherwise.

### GetHostUrlOk

`func (o *OutputRequest) GetHostUrlOk() (*string, bool)`

GetHostUrlOk returns a tuple with the HostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUrl

`func (o *OutputRequest) SetHostUrl(v string)`

SetHostUrl sets HostUrl field to given value.


### GetDatasetId

`func (o *OutputRequest) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *OutputRequest) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *OutputRequest) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetProjectId

`func (o *OutputRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OutputRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OutputRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTableId

`func (o *OutputRequest) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *OutputRequest) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *OutputRequest) SetTableId(v string)`

SetTableId sets TableId field to given value.


### GetServiceAccountKey

`func (o *OutputRequest) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *OutputRequest) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *OutputRequest) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.


### GetApiKey

`func (o *OutputRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OutputRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OutputRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetStreamName

`func (o *OutputRequest) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *OutputRequest) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *OutputRequest) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.


### GetLogType

`func (o *OutputRequest) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *OutputRequest) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *OutputRequest) SetLogType(v string)`

SetLogType sets LogType field to given value.


### GetSharedKey

`func (o *OutputRequest) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *OutputRequest) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *OutputRequest) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.


### GetTimeGeneratedField

`func (o *OutputRequest) GetTimeGeneratedField() string`

GetTimeGeneratedField returns the TimeGeneratedField field if non-nil, zero value otherwise.

### GetTimeGeneratedFieldOk

`func (o *OutputRequest) GetTimeGeneratedFieldOk() (*string, bool)`

GetTimeGeneratedFieldOk returns a tuple with the TimeGeneratedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGeneratedField

`func (o *OutputRequest) SetTimeGeneratedField(v string)`

SetTimeGeneratedField sets TimeGeneratedField field to given value.

### HasTimeGeneratedField

`func (o *OutputRequest) HasTimeGeneratedField() bool`

HasTimeGeneratedField returns a boolean if a field has been set.

### SetTimeGeneratedFieldNil

`func (o *OutputRequest) SetTimeGeneratedFieldNil(b bool)`

 SetTimeGeneratedFieldNil sets the value for TimeGeneratedField to be an explicit nil

### UnsetTimeGeneratedField
`func (o *OutputRequest) UnsetTimeGeneratedField()`

UnsetTimeGeneratedField ensures that no value is present for TimeGeneratedField, not even an explicit nil
### GetWorkspaceId

`func (o *OutputRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OutputRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OutputRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetStorageAccount

`func (o *OutputRequest) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *OutputRequest) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *OutputRequest) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetContainerName

`func (o *OutputRequest) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *OutputRequest) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *OutputRequest) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetBlobSasToken

`func (o *OutputRequest) GetBlobSasToken() string`

GetBlobSasToken returns the BlobSasToken field if non-nil, zero value otherwise.

### GetBlobSasTokenOk

`func (o *OutputRequest) GetBlobSasTokenOk() (*string, bool)`

GetBlobSasTokenOk returns a tuple with the BlobSasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSasToken

`func (o *OutputRequest) SetBlobSasToken(v string)`

SetBlobSasToken sets BlobSasToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


