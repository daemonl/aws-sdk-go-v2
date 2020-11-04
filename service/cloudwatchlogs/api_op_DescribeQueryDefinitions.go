// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns a paginated list of your saved CloudWatch Logs Insights
// query definitions. You can use the queryDefinitionNamePrefix parameter to limit
// the results to only the query definitions that have names that start with a
// certain string.
func (c *Client) DescribeQueryDefinitions(ctx context.Context, params *DescribeQueryDefinitionsInput, optFns ...func(*Options)) (*DescribeQueryDefinitionsOutput, error) {
	if params == nil {
		params = &DescribeQueryDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQueryDefinitions", params, optFns, addOperationDescribeQueryDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQueryDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQueryDefinitionsInput struct {

	// Limits the number of returned query definitions to the specified number.
	MaxResults *int32

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Use this parameter to filter your results to only the query definitions that
	// have names that start with the prefix you specify.
	QueryDefinitionNamePrefix *string
}

type DescribeQueryDefinitionsOutput struct {

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// The list of query definitions that match your request.
	QueryDefinitions []*types.QueryDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeQueryDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQueryDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQueryDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQueryDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQueryDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeQueryDefinitions",
	}
}
