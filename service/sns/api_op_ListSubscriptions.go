// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the requester's subscriptions. Each call returns a limited
// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
// is also returned. Use the NextToken parameter in a new ListSubscriptions call to
// get further results. This action is throttled at 30 transactions per second
// (TPS).
func (c *Client) ListSubscriptions(ctx context.Context, params *ListSubscriptionsInput, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) {
	if params == nil {
		params = &ListSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptions", params, optFns, addOperationListSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListSubscriptions action.
type ListSubscriptionsInput struct {

	// Token returned by the previous ListSubscriptions request.
	NextToken *string
}

// Response for ListSubscriptions action
type ListSubscriptionsOutput struct {

	// Token to pass along to the next ListSubscriptions request. This element is
	// returned if there are more subscriptions to retrieve.
	NextToken *string

	// A list of subscriptions.
	Subscriptions []*types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSubscriptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListSubscriptions",
	}
}
