// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Direct Connect gateway. You must first delete all virtual
// interfaces that are attached to the Direct Connect gateway and disassociate all
// virtual private gateways associated with the Direct Connect gateway.
func (c *Client) DeleteDirectConnectGateway(ctx context.Context, params *DeleteDirectConnectGatewayInput, optFns ...func(*Options)) (*DeleteDirectConnectGatewayOutput, error) {
	if params == nil {
		params = &DeleteDirectConnectGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDirectConnectGateway", params, optFns, addOperationDeleteDirectConnectGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDirectConnectGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectConnectGatewayInput struct {

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string
}

type DeleteDirectConnectGatewayOutput struct {

	// The Direct Connect gateway.
	DirectConnectGateway *types.DirectConnectGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDirectConnectGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectConnectGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectConnectGateway{}, middleware.After)
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
	if err = addOpDeleteDirectConnectGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectConnectGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDirectConnectGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteDirectConnectGateway",
	}
}
