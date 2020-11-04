// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// Versions that have been deleted aren't listed. Specify a runtime identifier
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list only
// versions that indicate that they're compatible with that runtime.
func (c *Client) ListLayerVersions(ctx context.Context, params *ListLayerVersionsInput, optFns ...func(*Options)) (*ListLayerVersionsOutput, error) {
	if params == nil {
		params = &ListLayerVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLayerVersions", params, optFns, addOperationListLayerVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLayerVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLayerVersionsInput struct {

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// This member is required.
	LayerName *string

	// A runtime identifier. For example, go1.x.
	CompatibleRuntime types.Runtime

	// A pagination token returned by a previous call.
	Marker *string

	// The maximum number of versions to return.
	MaxItems *int32
}

type ListLayerVersionsOutput struct {

	// A list of versions.
	LayerVersions []*types.LayerVersionsListItem

	// A pagination token returned when the response doesn't contain all versions.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLayerVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLayerVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLayerVersions{}, middleware.After)
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
	if err = addOpListLayerVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLayerVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListLayerVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListLayerVersions",
	}
}
