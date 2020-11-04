// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends activity on a fleet. Currently, this operation is used to stop a
// fleet's auto-scaling activity. It is used to temporarily stop triggering scaling
// events. The policies can be retained and auto-scaling activity can be restarted
// using StartFleetActions. You can view a fleet's stopped actions using
// DescribeFleetAttributes. To stop fleet actions, specify the fleet ID and the
// type of actions to suspend. When auto-scaling fleet actions are stopped, Amazon
// GameLift no longer initiates scaling events except in response to manual changes
// using UpdateFleetCapacity. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// *
// DescribeFleetAttributes
//
// * UpdateFleetAttributes
//
// * StartFleetActions or
// StopFleetActions
func (c *Client) StopFleetActions(ctx context.Context, params *StopFleetActionsInput, optFns ...func(*Options)) (*StopFleetActionsOutput, error) {
	if params == nil {
		params = &StopFleetActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopFleetActions", params, optFns, addOperationStopFleetActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopFleetActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopFleetActionsInput struct {

	// List of actions to suspend on the fleet.
	//
	// This member is required.
	Actions []types.FleetAction

	// A unique identifier for a fleet to stop actions on. You can use either the fleet
	// ID or ARN value.
	//
	// This member is required.
	FleetId *string
}

type StopFleetActionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopFleetActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopFleetActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopFleetActions{}, middleware.After)
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
	if err = addOpStopFleetActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopFleetActions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopFleetActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopFleetActions",
	}
}
