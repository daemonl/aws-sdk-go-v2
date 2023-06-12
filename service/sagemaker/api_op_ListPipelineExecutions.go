// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of the pipeline executions.
func (c *Client) ListPipelineExecutions(ctx context.Context, params *ListPipelineExecutionsInput, optFns ...func(*Options)) (*ListPipelineExecutionsOutput, error) {
	if params == nil {
		params = &ListPipelineExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelineExecutions", params, optFns, c.addOperationListPipelineExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelineExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPipelineExecutionsInput struct {

	// The name or Amazon Resource Name (ARN) of the pipeline.
	//
	// This member is required.
	PipelineName *string

	// A filter that returns the pipeline executions that were created after a
	// specified time.
	CreatedAfter *time.Time

	// A filter that returns the pipeline executions that were created before a
	// specified time.
	CreatedBefore *time.Time

	// The maximum number of pipeline executions to return in the response.
	MaxResults *int32

	// If the result of the previous ListPipelineExecutions request was truncated, the
	// response includes a NextToken . To retrieve the next set of pipeline executions,
	// use the token in the next request.
	NextToken *string

	// The field by which to sort results. The default is CreatedTime .
	SortBy types.SortPipelineExecutionsBy

	// The sort order for results.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListPipelineExecutionsOutput struct {

	// If the result of the previous ListPipelineExecutions request was truncated, the
	// response includes a NextToken . To retrieve the next set of pipeline executions,
	// use the token in the next request.
	NextToken *string

	// Contains a sorted list of pipeline execution summary objects matching the
	// specified filters. Each run summary includes the Amazon Resource Name (ARN) of
	// the pipeline execution, the run date, and the status. This list can be empty.
	PipelineExecutionSummaries []types.PipelineExecutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipelineExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelineExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelineExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpListPipelineExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelineExecutions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListPipelineExecutionsAPIClient is a client that implements the
// ListPipelineExecutions operation.
type ListPipelineExecutionsAPIClient interface {
	ListPipelineExecutions(context.Context, *ListPipelineExecutionsInput, ...func(*Options)) (*ListPipelineExecutionsOutput, error)
}

var _ ListPipelineExecutionsAPIClient = (*Client)(nil)

// ListPipelineExecutionsPaginatorOptions is the paginator options for
// ListPipelineExecutions
type ListPipelineExecutionsPaginatorOptions struct {
	// The maximum number of pipeline executions to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipelineExecutionsPaginator is a paginator for ListPipelineExecutions
type ListPipelineExecutionsPaginator struct {
	options   ListPipelineExecutionsPaginatorOptions
	client    ListPipelineExecutionsAPIClient
	params    *ListPipelineExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListPipelineExecutionsPaginator returns a new ListPipelineExecutionsPaginator
func NewListPipelineExecutionsPaginator(client ListPipelineExecutionsAPIClient, params *ListPipelineExecutionsInput, optFns ...func(*ListPipelineExecutionsPaginatorOptions)) *ListPipelineExecutionsPaginator {
	if params == nil {
		params = &ListPipelineExecutionsInput{}
	}

	options := ListPipelineExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipelineExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipelineExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipelineExecutions page.
func (p *ListPipelineExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipelineExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPipelineExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPipelineExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListPipelineExecutions",
	}
}
