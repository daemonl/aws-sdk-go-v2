// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DataflowEndpoint groups.
func (c *Client) ListDataflowEndpointGroups(ctx context.Context, params *ListDataflowEndpointGroupsInput, optFns ...func(*Options)) (*ListDataflowEndpointGroupsOutput, error) {
	if params == nil {
		params = &ListDataflowEndpointGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataflowEndpointGroups", params, optFns, addOperationListDataflowEndpointGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataflowEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListDataflowEndpointGroupsInput struct {

	// Maximum number of dataflow endpoint groups returned.
	MaxResults *int32

	// Next token returned in the request of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string
}

//
type ListDataflowEndpointGroupsOutput struct {

	// A list of dataflow endpoint groups.
	DataflowEndpointGroupList []*types.DataflowEndpointListItem

	// Next token returned in the response of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDataflowEndpointGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataflowEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataflowEndpointGroups{}, middleware.After)
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
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}
