// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about operations performed on a stack set.
func (c *Client) ListStackSetOperations(ctx context.Context, params *ListStackSetOperationsInput, optFns ...func(*Options)) (*ListStackSetOperationsOutput, error) {
	if params == nil {
		params = &ListStackSetOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackSetOperations", params, optFns, addOperationListStackSetOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackSetOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackSetOperationsInput struct {

	// The name or unique ID of the stack set that you want to get operation summaries
	// for.
	//
	// This member is required.
	StackSetName *string

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListStackSetOperations again and assign that token
	// to the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListStackSetOperationsOutput struct {

	// If the request doesn't return all results, NextToken is set to a token. To
	// retrieve the next set of results, call ListOperationResults again and assign
	// that token to the request object's NextToken parameter. If there are no
	// remaining results, NextToken is set to null.
	NextToken *string

	// A list of StackSetOperationSummary structures that contain summary information
	// about operations for the specified stack set.
	Summaries []*types.StackSetOperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStackSetOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackSetOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackSetOperations{}, middleware.After)
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
	if err = addOpListStackSetOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackSetOperations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListStackSetOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStackSetOperations",
	}
}
