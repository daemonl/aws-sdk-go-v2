// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List streaming distributions.
func (c *Client) ListStreamingDistributions(ctx context.Context, params *ListStreamingDistributionsInput, optFns ...func(*Options)) (*ListStreamingDistributionsOutput, error) {
	if params == nil {
		params = &ListStreamingDistributionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingDistributions", params, optFns, addOperationListStreamingDistributionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingDistributionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list your streaming distributions.
type ListStreamingDistributionsInput struct {

	// The value that you provided for the Marker request parameter.
	Marker *string

	// The value that you provided for the MaxItems request parameter.
	MaxItems *string
}

// The returned result of the corresponding request.
type ListStreamingDistributionsOutput struct {

	// The StreamingDistributionList type.
	StreamingDistributionList *types.StreamingDistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStreamingDistributionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListStreamingDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListStreamingDistributions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingDistributions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListStreamingDistributions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListStreamingDistributions",
	}
}
