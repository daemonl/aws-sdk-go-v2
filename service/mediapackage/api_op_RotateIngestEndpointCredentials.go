// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rotate the IngestEndpoint's username and password, as specified by the
// IngestEndpoint's id.
func (c *Client) RotateIngestEndpointCredentials(ctx context.Context, params *RotateIngestEndpointCredentialsInput, optFns ...func(*Options)) (*RotateIngestEndpointCredentialsOutput, error) {
	if params == nil {
		params = &RotateIngestEndpointCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RotateIngestEndpointCredentials", params, optFns, addOperationRotateIngestEndpointCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RotateIngestEndpointCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateIngestEndpointCredentialsInput struct {

	// The ID of the channel the IngestEndpoint is on.
	//
	// This member is required.
	Id *string

	// The id of the IngestEndpoint whose credentials should be rotated
	//
	// This member is required.
	IngestEndpointId *string
}

type RotateIngestEndpointCredentialsOutput struct {

	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// A short text description of the Channel.
	Description *string

	// Configure egress access logging.
	EgressAccessLogs *types.EgressAccessLogs

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest

	// The ID of the Channel.
	Id *string

	// Configure ingress access logging.
	IngressAccessLogs *types.IngressAccessLogs

	// A collection of tags associated with a resource
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRotateIngestEndpointCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRotateIngestEndpointCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRotateIngestEndpointCredentials{}, middleware.After)
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
	if err = addOpRotateIngestEndpointCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRotateIngestEndpointCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRotateIngestEndpointCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "RotateIngestEndpointCredentials",
	}
}
