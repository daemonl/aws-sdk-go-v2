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

// Changes the route table associated with a given subnet, internet gateway, or
// virtual private gateway in a VPC. After the operation completes, the subnet or
// gateway uses the routes in the new route table. For more information about route
// tables, see Route Tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
// Amazon Virtual Private Cloud User Guide. You can also use this operation to
// change which table is the main route table in the VPC. Specify the main route
// table's association ID and the route table ID of the new main route table.
func (c *Client) ReplaceRouteTableAssociation(ctx context.Context, params *ReplaceRouteTableAssociationInput, optFns ...func(*Options)) (*ReplaceRouteTableAssociationOutput, error) {
	if params == nil {
		params = &ReplaceRouteTableAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceRouteTableAssociation", params, optFns, addOperationReplaceRouteTableAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceRouteTableAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceRouteTableAssociationInput struct {

	// The association ID.
	//
	// This member is required.
	AssociationId *string

	// The ID of the new route table to associate with the subnet.
	//
	// This member is required.
	RouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ReplaceRouteTableAssociationOutput struct {

	// The state of the association.
	AssociationState *types.RouteTableAssociationState

	// The ID of the new association.
	NewAssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReplaceRouteTableAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReplaceRouteTableAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceRouteTableAssociation{}, middleware.After)
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
	if err = addOpReplaceRouteTableAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceRouteTableAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplaceRouteTableAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceRouteTableAssociation",
	}
}
