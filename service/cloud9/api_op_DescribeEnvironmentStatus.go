// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets status information for an AWS Cloud9 development environment.
func (c *Client) DescribeEnvironmentStatus(ctx context.Context, params *DescribeEnvironmentStatusInput, optFns ...func(*Options)) (*DescribeEnvironmentStatusOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentStatus", params, optFns, addOperationDescribeEnvironmentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEnvironmentStatusInput struct {

	// The ID of the environment to get status information about.
	//
	// This member is required.
	EnvironmentId *string
}

type DescribeEnvironmentStatusOutput struct {

	// Any informational message about the status of the environment.
	Message *string

	// The status of the environment. Available values include:
	//
	// * connecting: The
	// environment is connecting.
	//
	// * creating: The environment is being created.
	//
	// *
	// deleting: The environment is being deleted.
	//
	// * error: The environment is in an
	// error state.
	//
	// * ready: The environment is ready.
	//
	// * stopped: The environment is
	// stopped.
	//
	// * stopping: The environment is stopping.
	Status types.EnvironmentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEnvironmentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEnvironmentStatus{}, middleware.After)
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
	if err = addOpDescribeEnvironmentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "DescribeEnvironmentStatus",
	}
}
