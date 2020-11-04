// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Query is a synchronous operation that enables you to execute a query. Query will
// timeout after 60 seconds. You must update the default timeout in the SDK to
// support a timeout of 60 seconds. The result set will be truncated to 1MB.
// Service quotas apply. For more information, see Quotas in the Timestream
// Developer Guide.
func (c *Client) Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) {
	if params == nil {
		params = &QueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Query", params, optFns, addOperationQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryInput struct {

	// The query to be executed by Timestream.
	//
	// This member is required.
	QueryString *string

	// Unique, case-sensitive string of up to 64 ASCII characters that you specify when
	// you make a Query request. Providing a ClientToken makes the call to Query
	// idempotent, meaning that multiple identical calls have the same effect as one
	// single call. Your query request will fail in the following cases:
	//
	// * If you
	// submit a request with the same client token outside the 5-minute idepotency
	// window.
	//
	// * If you submit a request with the same client token but a change in
	// other parameters within the 5-minute idempotency window.
	//
	// After 4 hours, any
	// request with the same client token is treated as a new request.
	ClientToken *string

	// The total number of rows to return in the output. If the total number of rows
	// available is more than the value specified, a NextToken is provided in the
	// command's output. To resume pagination, provide the NextToken value in the
	// starting-token argument of a subsequent command.
	MaxRows *int32

	// A pagination token passed to get a set of results.
	NextToken *string
}

type QueryOutput struct {

	// The column data types of the returned result set.
	//
	// This member is required.
	ColumnInfo []*types.ColumnInfo

	// A unique ID for the given query.
	//
	// This member is required.
	QueryId *string

	// The result set rows returned by the query.
	//
	// This member is required.
	Rows []*types.Row

	// A pagination token that can be used again on a Query call to get the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpQuery{}, middleware.After)
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
	if err = addIdempotencyToken_opQueryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQuery(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*QueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *QueryInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opQueryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "Query",
	}
}
