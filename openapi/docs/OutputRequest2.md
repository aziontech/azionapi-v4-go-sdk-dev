# OutputRequest2

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

### NewOutputRequest2

`func NewOutputRequest2(url string, headers map[string]string, type_ string, bootstrapServers string, kafkaTopic string, useTls bool, accessKey string, secretKey string, region string, bucketName string, contentType string, hostUrl string, datasetId string, projectId string, tableId string, serviceAccountKey string, apiKey string, streamName string, logType string, sharedKey string, workspaceId string, storageAccount string, containerName string, blobSasToken string, ) *OutputRequest2`

NewOutputRequest2 instantiates a new OutputRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRequest2WithDefaults

`func NewOutputRequest2WithDefaults() *OutputRequest2`

NewOutputRequest2WithDefaults instantiates a new OutputRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OutputRequest2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutputRequest2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutputRequest2) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLogLineSeparator

`func (o *OutputRequest2) GetLogLineSeparator() string`

GetLogLineSeparator returns the LogLineSeparator field if non-nil, zero value otherwise.

### GetLogLineSeparatorOk

`func (o *OutputRequest2) GetLogLineSeparatorOk() (*string, bool)`

GetLogLineSeparatorOk returns a tuple with the LogLineSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLineSeparator

`func (o *OutputRequest2) SetLogLineSeparator(v string)`

SetLogLineSeparator sets LogLineSeparator field to given value.

### HasLogLineSeparator

`func (o *OutputRequest2) HasLogLineSeparator() bool`

HasLogLineSeparator returns a boolean if a field has been set.

### GetPayloadFormat

`func (o *OutputRequest2) GetPayloadFormat() string`

GetPayloadFormat returns the PayloadFormat field if non-nil, zero value otherwise.

### GetPayloadFormatOk

`func (o *OutputRequest2) GetPayloadFormatOk() (*string, bool)`

GetPayloadFormatOk returns a tuple with the PayloadFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFormat

`func (o *OutputRequest2) SetPayloadFormat(v string)`

SetPayloadFormat sets PayloadFormat field to given value.

### HasPayloadFormat

`func (o *OutputRequest2) HasPayloadFormat() bool`

HasPayloadFormat returns a boolean if a field has been set.

### GetMaxSize

`func (o *OutputRequest2) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *OutputRequest2) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *OutputRequest2) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *OutputRequest2) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *OutputRequest2) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *OutputRequest2) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetHeaders

`func (o *OutputRequest2) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutputRequest2) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutputRequest2) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetType

`func (o *OutputRequest2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputRequest2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputRequest2) SetType(v string)`

SetType sets Type field to given value.


### GetBootstrapServers

`func (o *OutputRequest2) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *OutputRequest2) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *OutputRequest2) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.


### GetKafkaTopic

`func (o *OutputRequest2) GetKafkaTopic() string`

GetKafkaTopic returns the KafkaTopic field if non-nil, zero value otherwise.

### GetKafkaTopicOk

`func (o *OutputRequest2) GetKafkaTopicOk() (*string, bool)`

GetKafkaTopicOk returns a tuple with the KafkaTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopic

`func (o *OutputRequest2) SetKafkaTopic(v string)`

SetKafkaTopic sets KafkaTopic field to given value.


### GetUseTls

`func (o *OutputRequest2) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *OutputRequest2) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *OutputRequest2) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.


### GetAccessKey

`func (o *OutputRequest2) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OutputRequest2) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OutputRequest2) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *OutputRequest2) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OutputRequest2) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OutputRequest2) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetRegion

`func (o *OutputRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OutputRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OutputRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetObjectKeyPrefix

`func (o *OutputRequest2) GetObjectKeyPrefix() string`

GetObjectKeyPrefix returns the ObjectKeyPrefix field if non-nil, zero value otherwise.

### GetObjectKeyPrefixOk

`func (o *OutputRequest2) GetObjectKeyPrefixOk() (*string, bool)`

GetObjectKeyPrefixOk returns a tuple with the ObjectKeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKeyPrefix

`func (o *OutputRequest2) SetObjectKeyPrefix(v string)`

SetObjectKeyPrefix sets ObjectKeyPrefix field to given value.

### HasObjectKeyPrefix

`func (o *OutputRequest2) HasObjectKeyPrefix() bool`

HasObjectKeyPrefix returns a boolean if a field has been set.

### SetObjectKeyPrefixNil

`func (o *OutputRequest2) SetObjectKeyPrefixNil(b bool)`

 SetObjectKeyPrefixNil sets the value for ObjectKeyPrefix to be an explicit nil

### UnsetObjectKeyPrefix
`func (o *OutputRequest2) UnsetObjectKeyPrefix()`

UnsetObjectKeyPrefix ensures that no value is present for ObjectKeyPrefix, not even an explicit nil
### GetBucketName

`func (o *OutputRequest2) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *OutputRequest2) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *OutputRequest2) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetContentType

`func (o *OutputRequest2) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OutputRequest2) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OutputRequest2) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetHostUrl

`func (o *OutputRequest2) GetHostUrl() string`

GetHostUrl returns the HostUrl field if non-nil, zero value otherwise.

### GetHostUrlOk

`func (o *OutputRequest2) GetHostUrlOk() (*string, bool)`

GetHostUrlOk returns a tuple with the HostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUrl

`func (o *OutputRequest2) SetHostUrl(v string)`

SetHostUrl sets HostUrl field to given value.


### GetDatasetId

`func (o *OutputRequest2) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *OutputRequest2) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *OutputRequest2) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetProjectId

`func (o *OutputRequest2) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OutputRequest2) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OutputRequest2) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTableId

`func (o *OutputRequest2) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *OutputRequest2) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *OutputRequest2) SetTableId(v string)`

SetTableId sets TableId field to given value.


### GetServiceAccountKey

`func (o *OutputRequest2) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *OutputRequest2) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *OutputRequest2) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.


### GetApiKey

`func (o *OutputRequest2) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OutputRequest2) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OutputRequest2) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetStreamName

`func (o *OutputRequest2) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *OutputRequest2) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *OutputRequest2) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.


### GetLogType

`func (o *OutputRequest2) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *OutputRequest2) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *OutputRequest2) SetLogType(v string)`

SetLogType sets LogType field to given value.


### GetSharedKey

`func (o *OutputRequest2) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *OutputRequest2) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *OutputRequest2) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.


### GetTimeGeneratedField

`func (o *OutputRequest2) GetTimeGeneratedField() string`

GetTimeGeneratedField returns the TimeGeneratedField field if non-nil, zero value otherwise.

### GetTimeGeneratedFieldOk

`func (o *OutputRequest2) GetTimeGeneratedFieldOk() (*string, bool)`

GetTimeGeneratedFieldOk returns a tuple with the TimeGeneratedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGeneratedField

`func (o *OutputRequest2) SetTimeGeneratedField(v string)`

SetTimeGeneratedField sets TimeGeneratedField field to given value.

### HasTimeGeneratedField

`func (o *OutputRequest2) HasTimeGeneratedField() bool`

HasTimeGeneratedField returns a boolean if a field has been set.

### SetTimeGeneratedFieldNil

`func (o *OutputRequest2) SetTimeGeneratedFieldNil(b bool)`

 SetTimeGeneratedFieldNil sets the value for TimeGeneratedField to be an explicit nil

### UnsetTimeGeneratedField
`func (o *OutputRequest2) UnsetTimeGeneratedField()`

UnsetTimeGeneratedField ensures that no value is present for TimeGeneratedField, not even an explicit nil
### GetWorkspaceId

`func (o *OutputRequest2) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OutputRequest2) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OutputRequest2) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetStorageAccount

`func (o *OutputRequest2) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *OutputRequest2) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *OutputRequest2) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetContainerName

`func (o *OutputRequest2) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *OutputRequest2) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *OutputRequest2) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetBlobSasToken

`func (o *OutputRequest2) GetBlobSasToken() string`

GetBlobSasToken returns the BlobSasToken field if non-nil, zero value otherwise.

### GetBlobSasTokenOk

`func (o *OutputRequest2) GetBlobSasTokenOk() (*string, bool)`

GetBlobSasTokenOk returns a tuple with the BlobSasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSasToken

`func (o *OutputRequest2) SetBlobSasToken(v string)`

SetBlobSasToken sets BlobSasToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


