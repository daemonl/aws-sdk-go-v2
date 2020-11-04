// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the artifact info that corresponds to an artifact id.
func (c *Client) GetArtifactUrl(ctx context.Context, params *GetArtifactUrlInput, optFns ...func(*Options)) (*GetArtifactUrlOutput, error) {
	if params == nil {
		params = &GetArtifactUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetArtifactUrl", params, optFns, addOperationGetArtifactUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetArtifactUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Returns the request structure for the get artifact request.
type GetArtifactUrlInput struct {

	// The unique ID for an artifact.
	//
	// This member is required.
	ArtifactId *string
}

// Returns the result structure for the get artifact request.
type GetArtifactUrlOutput struct {

	// The unique ID for an artifact.
	//
	// This member is required.
	ArtifactId *string

	// The presigned URL for the artifact.
	//
	// This member is required.
	ArtifactUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetArtifactUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetArtifactUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetArtifactUrl{}, middleware.After)
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
	if err = addOpGetArtifactUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetArtifactUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetArtifactUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "GetArtifactUrl",
	}
}
