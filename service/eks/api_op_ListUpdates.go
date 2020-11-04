// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the updates associated with an Amazon EKS cluster or managed node group in
// your AWS account, in the specified Region.
func (c *Client) ListUpdates(ctx context.Context, params *ListUpdatesInput, optFns ...func(*Options)) (*ListUpdatesOutput, error) {
	if params == nil {
		params = &ListUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUpdates", params, optFns, addOperationListUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUpdatesInput struct {

	// The name of the Amazon EKS cluster to list updates for.
	//
	// This member is required.
	Name *string

	// The maximum number of update results returned by ListUpdates in paginated
	// output. When you use this parameter, ListUpdates returns only maxResults results
	// in a single page along with a nextToken response element. You can see the
	// remaining results of the initial request by sending another ListUpdates request
	// with the returned nextToken value. This value can be between 1 and 100. If you
	// don't use this parameter, ListUpdates returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListUpdates request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string

	// The name of the Amazon EKS managed node group to list updates for.
	NodegroupName *string
}

type ListUpdatesOutput struct {

	// The nextToken value to include in a future ListUpdates request. When the results
	// of a ListUpdates request exceed maxResults, you can use this value to retrieve
	// the next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// A list of all the updates for the specified cluster and Region.
	UpdateIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUpdates{}, middleware.After)
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
	if err = addOpListUpdatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUpdates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListUpdates",
	}
}
