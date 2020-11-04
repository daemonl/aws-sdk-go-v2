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

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Retrieves information on all game servers that are currently active in a
// specified game server group. You can opt to sort the list by game server age.
// Use the pagination parameters to retrieve results in a set of sequential
// segments. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
// * RegisterGameServer
//
// * ListGameServers
//
// *
// ClaimGameServer
//
// * DescribeGameServer
//
// * UpdateGameServer
//
// *
// DeregisterGameServer
func (c *Client) ListGameServers(ctx context.Context, params *ListGameServersInput, optFns ...func(*Options)) (*ListGameServersOutput, error) {
	if params == nil {
		params = &ListGameServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGameServers", params, optFns, addOperationListGameServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGameServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGameServersInput struct {

	// An identifier for the game server group to retrieve a list of game servers from.
	// Use either the GameServerGroup name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential segments.
	Limit *int32

	// A token that indicates the start of the next sequential segment of results. Use
	// the token returned with the previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Indicates how to sort the returned data based on game server registration
	// timestamp. Use ASCENDING to retrieve oldest game servers first, or use
	// DESCENDING to retrieve newest game servers first. If this parameter is left
	// empty, game servers are returned in no particular order.
	SortOrder types.SortOrder
}

type ListGameServersOutput struct {

	// A collection of game server objects that match the request.
	GameServers []*types.GameServer

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGameServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGameServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGameServers{}, middleware.After)
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
	if err = addOpListGameServersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGameServers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListGameServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListGameServers",
	}
}
