// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists endpoint configurations.
func (c *Client) ListEndpointConfigs(ctx context.Context, params *ListEndpointConfigsInput, optFns ...func(*Options)) (*ListEndpointConfigsOutput, error) {
	if params == nil {
		params = &ListEndpointConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointConfigs", params, optFns, addOperationListEndpointConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointConfigsInput struct {

	// A filter that returns only endpoint configurations with a creation time greater
	// than or equal to the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only endpoint configurations created before the specified
	// time (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of training jobs to return in the response.
	MaxResults *int32

	// A string in the endpoint configuration name. This filter returns only endpoint
	// configurations whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListEndpointConfig request was truncated, the
	// response includes a NextToken. To retrieve the next set of endpoint
	// configurations, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.EndpointConfigSortKey

	// The sort order for results. The default is Descending.
	SortOrder types.OrderKey
}

type ListEndpointConfigsOutput struct {

	// An array of endpoint configurations.
	//
	// This member is required.
	EndpointConfigs []*types.EndpointConfigSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of endpoint configurations, use it in the subsequent request
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEndpointConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointConfigs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointConfigs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEndpointConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEndpointConfigs",
	}
}
