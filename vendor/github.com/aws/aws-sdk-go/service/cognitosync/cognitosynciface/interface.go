// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitosynciface provides an interface for the Amazon Cognito Sync.
package cognitosynciface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitosync"
)

// CognitoSyncAPI is the interface type for cognitosync.CognitoSync.
type CognitoSyncAPI interface {
	BulkPublishRequest(*cognitosync.BulkPublishInput) (*aws.Request, *cognitosync.BulkPublishOutput)

	BulkPublish(*cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)

	DeleteDatasetRequest(*cognitosync.DeleteDatasetInput) (*aws.Request, *cognitosync.DeleteDatasetOutput)

	DeleteDataset(*cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)

	DescribeDatasetRequest(*cognitosync.DescribeDatasetInput) (*aws.Request, *cognitosync.DescribeDatasetOutput)

	DescribeDataset(*cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)

	DescribeIdentityPoolUsageRequest(*cognitosync.DescribeIdentityPoolUsageInput) (*aws.Request, *cognitosync.DescribeIdentityPoolUsageOutput)

	DescribeIdentityPoolUsage(*cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)

	DescribeIdentityUsageRequest(*cognitosync.DescribeIdentityUsageInput) (*aws.Request, *cognitosync.DescribeIdentityUsageOutput)

	DescribeIdentityUsage(*cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)

	GetBulkPublishDetailsRequest(*cognitosync.GetBulkPublishDetailsInput) (*aws.Request, *cognitosync.GetBulkPublishDetailsOutput)

	GetBulkPublishDetails(*cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)

	GetCognitoEventsRequest(*cognitosync.GetCognitoEventsInput) (*aws.Request, *cognitosync.GetCognitoEventsOutput)

	GetCognitoEvents(*cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)

	GetIdentityPoolConfigurationRequest(*cognitosync.GetIdentityPoolConfigurationInput) (*aws.Request, *cognitosync.GetIdentityPoolConfigurationOutput)

	GetIdentityPoolConfiguration(*cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)

	ListDatasetsRequest(*cognitosync.ListDatasetsInput) (*aws.Request, *cognitosync.ListDatasetsOutput)

	ListDatasets(*cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)

	ListIdentityPoolUsageRequest(*cognitosync.ListIdentityPoolUsageInput) (*aws.Request, *cognitosync.ListIdentityPoolUsageOutput)

	ListIdentityPoolUsage(*cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)

	ListRecordsRequest(*cognitosync.ListRecordsInput) (*aws.Request, *cognitosync.ListRecordsOutput)

	ListRecords(*cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)

	RegisterDeviceRequest(*cognitosync.RegisterDeviceInput) (*aws.Request, *cognitosync.RegisterDeviceOutput)

	RegisterDevice(*cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)

	SetCognitoEventsRequest(*cognitosync.SetCognitoEventsInput) (*aws.Request, *cognitosync.SetCognitoEventsOutput)

	SetCognitoEvents(*cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)

	SetIdentityPoolConfigurationRequest(*cognitosync.SetIdentityPoolConfigurationInput) (*aws.Request, *cognitosync.SetIdentityPoolConfigurationOutput)

	SetIdentityPoolConfiguration(*cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)

	SubscribeToDatasetRequest(*cognitosync.SubscribeToDatasetInput) (*aws.Request, *cognitosync.SubscribeToDatasetOutput)

	SubscribeToDataset(*cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)

	UnsubscribeFromDatasetRequest(*cognitosync.UnsubscribeFromDatasetInput) (*aws.Request, *cognitosync.UnsubscribeFromDatasetOutput)

	UnsubscribeFromDataset(*cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)

	UpdateRecordsRequest(*cognitosync.UpdateRecordsInput) (*aws.Request, *cognitosync.UpdateRecordsOutput)

	UpdateRecords(*cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
}
