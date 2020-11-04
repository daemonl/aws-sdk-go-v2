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

// Adds an InputProcessingConfiguration to a SQL-based Kinesis Data Analytics
// application. An input processor pre-processes records on the input stream before
// the application's SQL code executes. Currently, the only input processor
// available is AWS Lambda (https://docs.aws.amazon.com/lambda/).
func (c *Client) AddApplicationInputProcessingConfiguration(ctx context.Context, params *AddApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*AddApplicationInputProcessingConfigurationOutput, error) {
	if params == nil {
		params = &AddApplicationInputProcessingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationInputProcessingConfiguration", params, optFns, addOperationAddApplicationInputProcessingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationInputProcessingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationInputProcessingConfigurationInput struct {

	// The name of the application to which you want to add the input processing
	// configuration.
	//
	// This member is required.
	ApplicationName *string

	// The version of the application to which you want to add the input processing
	// configuration. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the input configuration to add the input processing configuration to.
	// You can get a list of the input IDs for an application using the
	// DescribeApplication operation.
	//
	// This member is required.
	InputId *string

	// The InputProcessingConfiguration to add to the application.
	//
	// This member is required.
	InputProcessingConfiguration *types.InputProcessingConfiguration
}

type AddApplicationInputProcessingConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string

	// Provides the current application version.
	ApplicationVersionId *int64

	// The input ID that is associated with the application input. This is the ID that
	// Kinesis Data Analytics assigns to each input configuration that you add to your
	// application.
	InputId *string

	// The description of the preprocessor that executes on records in this input
	// before the application's code is run.
	InputProcessingConfigurationDescription *types.InputProcessingConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationInputProcessingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationInputProcessingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationInputProcessingConfiguration{}, middleware.After)
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
	if err = addOpAddApplicationInputProcessingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationInputProcessingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationInputProcessingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationInputProcessingConfiguration",
	}
}
