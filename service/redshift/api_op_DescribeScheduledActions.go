// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes properties of scheduled actions.
func (c *Client) DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledActions", params, optFns, addOperationDescribeScheduledActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledActionsInput struct {

	// If true, retrieve only active scheduled actions. If false, retrieve only
	// disabled scheduled actions.
	Active *bool

	// The end time in UTC of the scheduled action to retrieve. Only active scheduled
	// actions that have invocations before this time are retrieved.
	EndTime *time.Time

	// List of scheduled action filters.
	Filters []*types.ScheduledActionFilter

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request exceed
	// the value specified in MaxRecords, AWS returns a value in the Marker field of
	// the response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the scheduled action to retrieve.
	ScheduledActionName *string

	// The start time in UTC of the scheduled actions to retrieve. Only active
	// scheduled actions that have invocations after this time are retrieved.
	StartTime *time.Time

	// The type of the scheduled actions to retrieve.
	TargetActionType types.ScheduledActionTypeValues
}

type DescribeScheduledActionsOutput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request exceed
	// the value specified in MaxRecords, AWS returns a value in the Marker field of
	// the response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// List of retrieved scheduled actions.
	ScheduledActions []*types.ScheduledAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeScheduledActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScheduledActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScheduledActions{}, middleware.After)
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
	if err = addOpDescribeScheduledActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeScheduledActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeScheduledActions",
	}
}
