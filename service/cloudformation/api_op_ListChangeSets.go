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

// Returns the ID and status of each active change set for a stack. For example,
// AWS CloudFormation lists change sets that are in the CREATE_IN_PROGRESS or
// CREATE_PENDING state.
func (c *Client) ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
	if params == nil {
		params = &ListChangeSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChangeSets", params, optFns, addOperationListChangeSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChangeSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListChangeSets action.
type ListChangeSetsInput struct {

	// The name or the Amazon Resource Name (ARN) of the stack for which you want to
	// list change sets.
	//
	// This member is required.
	StackName *string

	// A string (provided by the ListChangeSets response output) that identifies the
	// next page of change sets that you want to retrieve.
	NextToken *string
}

// The output for the ListChangeSets action.
type ListChangeSetsOutput struct {

	// If the output exceeds 1 MB, a string that identifies the next page of change
	// sets. If there is no additional page, this value is null.
	NextToken *string

	// A list of ChangeSetSummary structures that provides the ID and status of each
	// change set for the specified stack.
	Summaries []*types.ChangeSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListChangeSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListChangeSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListChangeSets{}, middleware.After)
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
	if err = addOpListChangeSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChangeSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListChangeSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListChangeSets",
	}
}
