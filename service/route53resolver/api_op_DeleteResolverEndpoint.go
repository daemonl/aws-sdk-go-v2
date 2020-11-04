// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a Resolver endpoint. The effect of deleting a Resolver endpoint depends
// on whether it's an inbound or an outbound Resolver endpoint:
//
// * Inbound: DNS
// queries from your network are no longer routed to the DNS service for the
// specified VPC.
//
// * Outbound: DNS queries from a VPC are no longer routed to your
// network.
func (c *Client) DeleteResolverEndpoint(ctx context.Context, params *DeleteResolverEndpointInput, optFns ...func(*Options)) (*DeleteResolverEndpointOutput, error) {
	if params == nil {
		params = &DeleteResolverEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResolverEndpoint", params, optFns, addOperationDeleteResolverEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResolverEndpointInput struct {

	// The ID of the Resolver endpoint that you want to delete.
	//
	// This member is required.
	ResolverEndpointId *string
}

type DeleteResolverEndpointOutput struct {

	// Information about the DeleteResolverEndpoint request, including the status of
	// the request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteResolverEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResolverEndpoint{}, middleware.After)
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
	if err = addOpDeleteResolverEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResolverEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteResolverEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "DeleteResolverEndpoint",
	}
}
