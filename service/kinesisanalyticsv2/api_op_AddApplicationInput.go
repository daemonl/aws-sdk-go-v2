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

// Adds a streaming source to your SQL-based Kinesis Data Analytics application.
// You can add a streaming source when you create an application, or you can use
// this operation to add a streaming source after you create an application. For
// more information, see CreateApplication. Any configuration update, including
// adding a streaming source using this operation, results in a new version of the
// application. You can use the DescribeApplication operation to find the current
// application version.
func (c *Client) AddApplicationInput(ctx context.Context, params *AddApplicationInputInput, optFns ...func(*Options)) (*AddApplicationInputOutput, error) {
	if params == nil {
		params = &AddApplicationInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationInput", params, optFns, addOperationAddApplicationInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationInputInput struct {

	// The name of your existing application to which you want to add the streaming
	// source.
	//
	// This member is required.
	ApplicationName *string

	// The current version of your application. You can use the DescribeApplication
	// operation to find the current application version.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The Input to add.
	//
	// This member is required.
	Input *types.Input
}

type AddApplicationInputOutput struct {

	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string

	// Provides the current application version.
	ApplicationVersionId *int64

	// Describes the application input configuration.
	InputDescriptions []*types.InputDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationInput{}, middleware.After)
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
	if err = addOpAddApplicationInputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationInput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationInput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationInput",
	}
}
