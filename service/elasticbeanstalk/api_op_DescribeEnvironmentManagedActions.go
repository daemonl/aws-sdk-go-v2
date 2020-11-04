// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists an environment's upcoming and in-progress managed actions.
func (c *Client) DescribeEnvironmentManagedActions(ctx context.Context, params *DescribeEnvironmentManagedActionsInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentManagedActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentManagedActions", params, optFns, addOperationDescribeEnvironmentManagedActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentManagedActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list an environment's upcoming and in-progress managed actions.
type DescribeEnvironmentManagedActionsInput struct {

	// The environment ID of the target environment.
	EnvironmentId *string

	// The name of the target environment.
	EnvironmentName *string

	// To show only actions with a particular status, specify a status.
	Status types.ActionStatus
}

// The result message containing a list of managed actions.
type DescribeEnvironmentManagedActionsOutput struct {

	// A list of upcoming and in-progress managed actions.
	ManagedActions []*types.ManagedAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEnvironmentManagedActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentManagedActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentManagedActions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEnvironmentManagedActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentManagedActions",
	}
}
