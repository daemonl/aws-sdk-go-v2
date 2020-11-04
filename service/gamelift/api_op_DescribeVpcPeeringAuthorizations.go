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

// Retrieves valid VPC peering authorizations that are pending for the AWS account.
// This operation returns all VPC peering authorizations and requests for peering.
// This includes those initiated and received by this account.
//
// *
// CreateVpcPeeringAuthorization
//
// * DescribeVpcPeeringAuthorizations
//
// *
// DeleteVpcPeeringAuthorization
//
// * CreateVpcPeeringConnection
//
// *
// DescribeVpcPeeringConnections
//
// * DeleteVpcPeeringConnection
func (c *Client) DescribeVpcPeeringAuthorizations(ctx context.Context, params *DescribeVpcPeeringAuthorizationsInput, optFns ...func(*Options)) (*DescribeVpcPeeringAuthorizationsOutput, error) {
	if params == nil {
		params = &DescribeVpcPeeringAuthorizationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcPeeringAuthorizations", params, optFns, addOperationDescribeVpcPeeringAuthorizationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcPeeringAuthorizationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcPeeringAuthorizationsInput struct {
}

type DescribeVpcPeeringAuthorizationsOutput struct {

	// A collection of objects that describe all valid VPC peering operations for the
	// current AWS account.
	VpcPeeringAuthorizations []*types.VpcPeeringAuthorization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcPeeringAuthorizationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringAuthorizations",
	}
}
