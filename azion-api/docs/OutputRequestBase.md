# OutputRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;standard&#x60; - Standard HTTP/HTTPS POST * &#x60;kafka&#x60; - Apache Kafka * &#x60;s3&#x60; - Simple Storage Service (S3) * &#x60;big_query&#x60; - Google BigQuery * &#x60;elasticsearch&#x60; - Elasticsearch * &#x60;splunk&#x60; - Splunk * &#x60;aws_kinesis_firehose&#x60; - AWS Kinesis Data Firehose * &#x60;datadog&#x60; - Datadog * &#x60;qradar&#x60; - IBM QRadar * &#x60;azure_monitor&#x60; - Azure Monitor * &#x60;azure_blob_storage&#x60; - Azure Blob Storage | 
**Attributes** | [**OutputRequest**](OutputRequest.md) |  | 

## Methods

### NewOutputRequestBase

`func NewOutputRequestBase(type_ string, attributes OutputRequest, ) *OutputRequestBase`

NewOutputRequestBase instantiates a new OutputRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRequestBaseWithDefaults

`func NewOutputRequestBaseWithDefaults() *OutputRequestBase`

NewOutputRequestBaseWithDefaults instantiates a new OutputRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputRequestBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputRequestBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputRequestBase) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OutputRequestBase) GetAttributes() OutputRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OutputRequestBase) GetAttributesOk() (*OutputRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OutputRequestBase) SetAttributes(v OutputRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


