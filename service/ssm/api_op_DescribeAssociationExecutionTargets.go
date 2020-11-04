// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this API action to view information about a specific execution of a specific
// association.
func (c *Client) DescribeAssociationExecutionTargets(ctx context.Context, params *DescribeAssociationExecutionTargetsInput, optFns ...func(*Options)) (*DescribeAssociationExecutionTargetsOutput, error) {
	if params == nil {
		params = &DescribeAssociationExecutionTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssociationExecutionTargets", params, optFns, addOperationDescribeAssociationExecutionTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssociationExecutionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssociationExecutionTargetsInput struct {

	// The association ID that includes the execution for which you want to view
	// details.
	//
	// This member is required.
	AssociationId *string

	// The execution ID for which you want to view details.
	//
	// This member is required.
	ExecutionId *string

	// Filters for the request. You can specify the following filters and values.
	// Status (EQUAL) ResourceId (EQUAL) ResourceType (EQUAL)
	Filters []*types.AssociationExecutionTargetsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type DescribeAssociationExecutionTargetsOutput struct {

	// Information about the execution.
	AssociationExecutionTargets []*types.AssociationExecutionTarget

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAssociationExecutionTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssociationExecutionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssociationExecutionTargets{}, middleware.After)
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
	if err = addOpDescribeAssociationExecutionTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssociationExecutionTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAssociationExecutionTargets",
	}
}
