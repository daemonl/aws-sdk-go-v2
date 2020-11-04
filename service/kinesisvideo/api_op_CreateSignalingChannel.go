// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a signaling channel. CreateSignalingChannel is an asynchronous
// operation.
func (c *Client) CreateSignalingChannel(ctx context.Context, params *CreateSignalingChannelInput, optFns ...func(*Options)) (*CreateSignalingChannelOutput, error) {
	if params == nil {
		params = &CreateSignalingChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSignalingChannel", params, optFns, addOperationCreateSignalingChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSignalingChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSignalingChannelInput struct {

	// A name for the signaling channel that you are creating. It must be unique for
	// each AWS account and AWS Region.
	//
	// This member is required.
	ChannelName *string

	// A type of the signaling channel that you are creating. Currently, SINGLE_MASTER
	// is the only supported channel type.
	ChannelType types.ChannelType

	// A structure containing the configuration for the SINGLE_MASTER channel type.
	SingleMasterConfiguration *types.SingleMasterConfiguration

	// A set of tags (key-value pairs) that you want to associate with this channel.
	Tags []*types.Tag
}

type CreateSignalingChannelOutput struct {

	// The Amazon Resource Name (ARN) of the created channel.
	ChannelARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSignalingChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSignalingChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSignalingChannel{}, middleware.After)
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
	if err = addOpCreateSignalingChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSignalingChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSignalingChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "CreateSignalingChannel",
	}
}
