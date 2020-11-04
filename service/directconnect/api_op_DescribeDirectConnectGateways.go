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

// Lists all your Direct Connect gateways or only the specified Direct Connect
// gateway. Deleted Direct Connect gateways are not returned.
func (c *Client) DescribeDirectConnectGateways(ctx context.Context, params *DescribeDirectConnectGatewaysInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewaysOutput, error) {
	if params == nil {
		params = &DescribeDirectConnectGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectConnectGateways", params, optFns, addOperationDescribeDirectConnectGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectConnectGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewaysInput struct {

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token provided in the previous call to retrieve the next page.
	NextToken *string
}

type DescribeDirectConnectGatewaysOutput struct {

	// The Direct Connect gateways.
	DirectConnectGateways []*types.DirectConnectGateway

	// The token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDirectConnectGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGateways(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDirectConnectGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGateways",
	}
}
