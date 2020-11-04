// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of behavior graphs that the calling account is a master of.
// This operation can only be called by a master account. Because an account can
// currently only be the master of one behavior graph within a Region, the results
// always contain a single graph.
func (c *Client) ListGraphs(ctx context.Context, params *ListGraphsInput, optFns ...func(*Options)) (*ListGraphsOutput, error) {
	if params == nil {
		params = &ListGraphsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGraphs", params, optFns, addOperationListGraphsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGraphsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGraphsInput struct {

	// The maximum number of graphs to return at a time. The total must be less than
	// the overall limit on the number of results to return, which is currently 200.
	MaxResults *int32

	// For requests to get the next page of results, the pagination token that was
	// returned with the previous set of results. The initial request does not include
	// a pagination token.
	NextToken *string
}

type ListGraphsOutput struct {

	// A list of behavior graphs that the account is a master for.
	GraphList []*types.Graph

	// If there are more behavior graphs remaining in the results, then this is the
	// pagination token to use to request the next page of behavior graphs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGraphsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGraphs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGraphs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGraphs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListGraphs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "ListGraphs",
	}
}
