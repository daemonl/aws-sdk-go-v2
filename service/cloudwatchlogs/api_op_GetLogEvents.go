// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists log events from the specified log stream. You can list all of the log
// events or filter using a time range. By default, this operation returns as many
// log events as can fit in a response size of 1MB (up to 10,000 log events). You
// can get additional log events by specifying one of the tokens in a subsequent
// call. This operation can return empty results while there are more log events
// available through the token.
func (c *Client) GetLogEvents(ctx context.Context, params *GetLogEventsInput, optFns ...func(*Options)) (*GetLogEventsOutput, error) {
	if params == nil {
		params = &GetLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLogEvents", params, optFns, addOperationGetLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLogEventsInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The name of the log stream.
	//
	// This member is required.
	LogStreamName *string

	// The end of the time range, expressed as the number of milliseconds after Jan 1,
	// 1970 00:00:00 UTC. Events with a timestamp equal to or later than this time are
	// not included.
	EndTime *int64

	// The maximum number of log events returned. If you don't specify a value, the
	// maximum is as many log events as can fit in a response size of 1 MB, up to
	// 10,000 log events.
	Limit *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.) Using this token works only when you specify true for
	// startFromHead.
	NextToken *string

	// If the value is true, the earliest log events are returned first. If the value
	// is false, the latest log events are returned first. The default value is false.
	// If you are using nextToken in this operation, you must specify true for
	// startFromHead.
	StartFromHead *bool

	// The start of the time range, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC. Events with a timestamp equal to this time or later than
	// this time are included. Events with a timestamp earlier than this time are not
	// included.
	StartTime *int64
}

type GetLogEventsOutput struct {

	// The events.
	Events []*types.OutputLogEvent

	// The token for the next set of items in the backward direction. The token expires
	// after 24 hours. This token is never null. If you have reached the end of the
	// stream, it returns the same token you passed in.
	NextBackwardToken *string

	// The token for the next set of items in the forward direction. The token expires
	// after 24 hours. If you have reached the end of the stream, it returns the same
	// token you passed in.
	NextForwardToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLogEvents{}, middleware.After)
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
	if err = addOpGetLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLogEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "GetLogEvents",
	}
}
