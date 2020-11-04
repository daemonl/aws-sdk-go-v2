// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an Amazon CloudWatch log stream to monitor application configuration
// errors.
func (c *Client) AddApplicationCloudWatchLoggingOption(ctx context.Context, params *AddApplicationCloudWatchLoggingOptionInput, optFns ...func(*Options)) (*AddApplicationCloudWatchLoggingOptionOutput, error) {
	if params == nil {
		params = &AddApplicationCloudWatchLoggingOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationCloudWatchLoggingOption", params, optFns, addOperationAddApplicationCloudWatchLoggingOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationCloudWatchLoggingOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationCloudWatchLoggingOptionInput struct {

	// The Kinesis Data Analytics application name.
	//
	// This member is required.
	ApplicationName *string

	// Provides the Amazon CloudWatch log stream Amazon Resource Name (ARN).
	//
	// This member is required.
	CloudWatchLoggingOption *types.CloudWatchLoggingOption

	// The version ID of the Kinesis Data Analytics application. You can retrieve the
	// application version ID using DescribeApplication.
	//
	// This member is required.
	CurrentApplicationVersionId *int64
}

type AddApplicationCloudWatchLoggingOptionOutput struct {

	// The application's ARN.
	ApplicationARN *string

	// The new version ID of the Kinesis Data Analytics application. Kinesis Data
	// Analytics updates the ApplicationVersionId each time you change the CloudWatch
	// logging options.
	ApplicationVersionId *int64

	// The descriptions of the current CloudWatch logging options for the Kinesis Data
	// Analytics application.
	CloudWatchLoggingOptionDescriptions []*types.CloudWatchLoggingOptionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationCloudWatchLoggingOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationCloudWatchLoggingOption{}, middleware.After)
	if err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAddApplicationCloudWatchLoggingOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAddApplicationCloudWatchLoggingOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationCloudWatchLoggingOption",
	}
}
