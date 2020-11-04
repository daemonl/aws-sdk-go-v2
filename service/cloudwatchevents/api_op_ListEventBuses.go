// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the event buses in your account, including the default event bus,
// custom event buses, and partner event buses.
func (c *Client) ListEventBuses(ctx context.Context, params *ListEventBusesInput, optFns ...func(*Options)) (*ListEventBusesOutput, error) {
	if params == nil {
		params = &ListEventBusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventBuses", params, optFns, addOperationListEventBusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventBusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventBusesInput struct {

	// Specifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// Specifying this limits the results to only those event buses with names that
	// start with the specified prefix.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListEventBusesOutput struct {

	// This list of event buses.
	EventBuses []*types.EventBus

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEventBusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventBuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventBuses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventBuses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEventBuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListEventBuses",
	}
}
