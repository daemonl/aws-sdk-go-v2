// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the destination domain owner to reject an inbound cross-cluster search
// connection request.
func (c *Client) RejectInboundCrossClusterSearchConnection(ctx context.Context, params *RejectInboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*RejectInboundCrossClusterSearchConnectionOutput, error) {
	if params == nil {
		params = &RejectInboundCrossClusterSearchConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectInboundCrossClusterSearchConnection", params, optFns, addOperationRejectInboundCrossClusterSearchConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectInboundCrossClusterSearchConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the RejectInboundCrossClusterSearchConnection
// operation.
type RejectInboundCrossClusterSearchConnectionInput struct {

	// The id of the inbound connection that you want to reject.
	//
	// This member is required.
	CrossClusterSearchConnectionId *string
}

// The result of a RejectInboundCrossClusterSearchConnection operation. Contains
// details of rejected inbound connection.
type RejectInboundCrossClusterSearchConnectionOutput struct {

	// Specifies the InboundCrossClusterSearchConnection of rejected inbound
	// connection.
	CrossClusterSearchConnection *types.InboundCrossClusterSearchConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRejectInboundCrossClusterSearchConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRejectInboundCrossClusterSearchConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRejectInboundCrossClusterSearchConnection{}, middleware.After)
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
	if err = addOpRejectInboundCrossClusterSearchConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectInboundCrossClusterSearchConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRejectInboundCrossClusterSearchConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "RejectInboundCrossClusterSearchConnection",
	}
}
