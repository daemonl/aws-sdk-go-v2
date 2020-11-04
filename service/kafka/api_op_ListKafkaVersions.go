// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Kafka versions.
func (c *Client) ListKafkaVersions(ctx context.Context, params *ListKafkaVersionsInput, optFns ...func(*Options)) (*ListKafkaVersionsOutput, error) {
	if params == nil {
		params = &ListKafkaVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKafkaVersions", params, optFns, addOperationListKafkaVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKafkaVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKafkaVersionsInput struct {

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults *int32

	// The paginated results marker. When the result of the operation is truncated, the
	// call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string
}

type ListKafkaVersionsOutput struct {
	KafkaVersions []*types.KafkaVersion

	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListKafkaVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKafkaVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKafkaVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKafkaVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListKafkaVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListKafkaVersions",
	}
}
