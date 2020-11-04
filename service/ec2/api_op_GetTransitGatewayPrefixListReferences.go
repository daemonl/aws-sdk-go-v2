// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the prefix list references in a specified transit gateway
// route table.
func (c *Client) GetTransitGatewayPrefixListReferences(ctx context.Context, params *GetTransitGatewayPrefixListReferencesInput, optFns ...func(*Options)) (*GetTransitGatewayPrefixListReferencesOutput, error) {
	if params == nil {
		params = &GetTransitGatewayPrefixListReferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayPrefixListReferences", params, optFns, addOperationGetTransitGatewayPrefixListReferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayPrefixListReferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayPrefixListReferencesInput struct {

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	// * attachment.resource-id - The ID
	// of the resource for the attachment.
	//
	// * attachment.resource-type - The type of
	// resource for the attachment. Valid values are vpc | vpn | direct-connect-gateway
	// | peering.
	//
	// * attachment.transit-gateway-attachment-id - The ID of the
	// attachment.
	//
	// * is-blackhole - Whether traffic matching the route is blocked
	// (true | false).
	//
	// * prefix-list-id - The ID of the prefix list.
	//
	// *
	// prefix-list-owner-id - The ID of the owner of the prefix list.
	//
	// * state - The
	// state of the prefix list reference (pending | available | modifying | deleting).
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetTransitGatewayPrefixListReferencesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the prefix list references.
	TransitGatewayPrefixListReferences []*types.TransitGatewayPrefixListReference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTransitGatewayPrefixListReferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayPrefixListReferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayPrefixListReferences{}, middleware.After)
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
	if err = addOpGetTransitGatewayPrefixListReferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayPrefixListReferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTransitGatewayPrefixListReferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayPrefixListReferences",
	}
}
