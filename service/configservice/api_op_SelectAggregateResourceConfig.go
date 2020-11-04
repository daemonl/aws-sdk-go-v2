// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts a structured query language (SQL) SELECT command and an aggregator to
// query configuration state of AWS resources across multiple accounts and regions,
// performs the corresponding search, and returns resource configurations matching
// the properties. For more information about query components, see the  Query
// Components
// (https://docs.aws.amazon.com/config/latest/developerguide/query-components.html)
// section in the AWS Config Developer Guide.
func (c *Client) SelectAggregateResourceConfig(ctx context.Context, params *SelectAggregateResourceConfigInput, optFns ...func(*Options)) (*SelectAggregateResourceConfigOutput, error) {
	if params == nil {
		params = &SelectAggregateResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SelectAggregateResourceConfig", params, optFns, addOperationSelectAggregateResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SelectAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SelectAggregateResourceConfigInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The SQL query SELECT command.
	//
	// This member is required.
	Expression *string

	// The maximum number of query results returned on each page.
	Limit *int32

	// The maximum number of query results returned on each page. AWS Config also
	// allows the Limit request parameter.
	MaxResults *int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string
}

type SelectAggregateResourceConfigOutput struct {

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Details about the query.
	QueryInfo *types.QueryInfo

	// Returns the results for the SQL query.
	Results []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSelectAggregateResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSelectAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSelectAggregateResourceConfig{}, middleware.After)
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
	if err = addOpSelectAggregateResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSelectAggregateResourceConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSelectAggregateResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "SelectAggregateResourceConfig",
	}
}
