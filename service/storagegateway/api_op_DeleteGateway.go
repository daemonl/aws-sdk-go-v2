// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a gateway. To specify which gateway to delete, use the Amazon Resource
// Name (ARN) of the gateway in your request. The operation deletes the gateway;
// however, it does not delete the gateway virtual machine (VM) from your host
// computer. After you delete a gateway, you cannot reactivate it. Completed
// snapshots of the gateway volumes are not deleted upon deleting the gateway,
// however, pending snapshots will not complete. After you delete a gateway, your
// next step is to remove it from your environment. You no longer pay software
// charges after the gateway is deleted; however, your existing Amazon EBS
// snapshots persist and you will continue to be billed for these snapshots. You
// can choose to remove all remaining Amazon EBS snapshots by canceling your Amazon
// EC2 subscription. If you prefer not to cancel your Amazon EC2 subscription, you
// can delete your snapshots using the Amazon EC2 console. For more information,
// see the AWS Storage Gateway detail page (http://aws.amazon.com/storagegateway).
func (c *Client) DeleteGateway(ctx context.Context, params *DeleteGatewayInput, optFns ...func(*Options)) (*DeleteGatewayOutput, error) {
	if params == nil {
		params = &DeleteGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGateway", params, optFns, addOperationDeleteGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the ID of the gateway to delete.
type DeleteGatewayInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the ID of the deleted gateway.
type DeleteGatewayOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGateway{}, middleware.After)
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
	if err = addOpDeleteGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteGateway",
	}
}
