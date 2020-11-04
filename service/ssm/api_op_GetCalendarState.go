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

// Gets the state of the AWS Systems Manager Change Calendar at an optional,
// specified time. If you specify a time, GetCalendarState returns the state of the
// calendar at a specific time, and returns the next time that the Change Calendar
// state will transition. If you do not specify a time, GetCalendarState assumes
// the current time. Change Calendar entries have two possible states: OPEN or
// CLOSED. If you specify more than one calendar in a request, the command returns
// the status of OPEN only if all calendars in the request are open. If one or more
// calendars in the request are closed, the status returned is CLOSED. For more
// information about Systems Manager Change Calendar, see AWS Systems Manager
// Change Calendar
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar.html)
// in the AWS Systems Manager User Guide.
func (c *Client) GetCalendarState(ctx context.Context, params *GetCalendarStateInput, optFns ...func(*Options)) (*GetCalendarStateOutput, error) {
	if params == nil {
		params = &GetCalendarStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCalendarState", params, optFns, addOperationGetCalendarStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCalendarStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCalendarStateInput struct {

	// The names or Amazon Resource Names (ARNs) of the Systems Manager documents that
	// represent the calendar entries for which you want to get the state.
	//
	// This member is required.
	CalendarNames []*string

	// (Optional) The specific time for which you want to get calendar state
	// information, in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) format. If you
	// do not add AtTime, the current time is assumed.
	AtTime *string
}

type GetCalendarStateOutput struct {

	// The time, as an ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) string, that
	// you specified in your command. If you did not specify a time, GetCalendarState
	// uses the current time.
	AtTime *string

	// The time, as an ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) string, that
	// the calendar state will change. If the current calendar state is OPEN,
	// NextTransitionTime indicates when the calendar state changes to CLOSED, and
	// vice-versa.
	NextTransitionTime *string

	// The state of the calendar. An OPEN calendar indicates that actions are allowed
	// to proceed, and a CLOSED calendar indicates that actions are not allowed to
	// proceed.
	State types.CalendarState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCalendarStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCalendarState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCalendarState{}, middleware.After)
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
	if err = addOpGetCalendarStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCalendarState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCalendarState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetCalendarState",
	}
}
