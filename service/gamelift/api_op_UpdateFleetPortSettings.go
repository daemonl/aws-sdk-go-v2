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

// Updates port settings for a fleet. To update settings, specify the fleet ID to
// be updated and list the permissions you want to update. List the permissions you
// want to add in InboundPermissionAuthorizations, and permissions you want to
// remove in InboundPermissionRevocations. Permissions to be removed must match
// existing fleet permissions. If successful, the fleet ID for the updated fleet is
// returned. Learn more Setting up GameLift Fleets
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
// * Update fleets:
//
// * UpdateFleetAttributes
//
// *
// UpdateFleetCapacity
//
// * UpdateFleetPortSettings
//
// * UpdateRuntimeConfiguration
//
// *
// StartFleetActions or StopFleetActions
func (c *Client) UpdateFleetPortSettings(ctx context.Context, params *UpdateFleetPortSettingsInput, optFns ...func(*Options)) (*UpdateFleetPortSettingsOutput, error) {
	if params == nil {
		params = &UpdateFleetPortSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetPortSettings", params, optFns, addOperationUpdateFleetPortSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetPortSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type UpdateFleetPortSettingsInput struct {

	// A unique identifier for a fleet to update port settings for. You can use either
	// the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// A collection of port settings to be added to the fleet resource.
	InboundPermissionAuthorizations []*types.IpPermission

	// A collection of port settings to be removed from the fleet resource.
	InboundPermissionRevocations []*types.IpPermission
}

// Represents the returned data in response to a request operation.
type UpdateFleetPortSettingsOutput struct {

	// A unique identifier for a fleet that was updated.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFleetPortSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetPortSettings{}, middleware.After)
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
	if err = addOpUpdateFleetPortSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetPortSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetPortSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateFleetPortSettings",
	}
}
