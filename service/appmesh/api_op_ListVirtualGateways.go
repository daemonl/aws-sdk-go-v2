// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of existing virtual gateways in a service mesh.
func (c *Client) ListVirtualGateways(ctx context.Context, params *ListVirtualGatewaysInput, optFns ...func(*Options)) (*ListVirtualGatewaysOutput, error) {
	if params == nil {
		params = &ListVirtualGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualGateways", params, optFns, addOperationListVirtualGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualGatewaysInput struct {

	// The name of the service mesh to list virtual gateways in.
	//
	// This member is required.
	MeshName *string

	// The maximum number of results returned by ListVirtualGateways in paginated
	// output. When you use this parameter, ListVirtualGateways returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualGateways request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualGateways returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListVirtualGateways
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

type ListVirtualGatewaysOutput struct {

	// The list of existing virtual gateways for the specified service mesh.
	//
	// This member is required.
	VirtualGateways []*types.VirtualGatewayRef

	// The nextToken value to include in a future ListVirtualGateways request. When the
	// results of a ListVirtualGateways request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualGateways{}, middleware.After)
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
	if err = addOpListVirtualGatewaysValidationMiddleware(stack); err != nil {
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
