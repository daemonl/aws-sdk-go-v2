// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the APNs VoIP sandbox channel for an application or updates the status
// and settings of the APNs VoIP sandbox channel for an application.
func (c *Client) UpdateApnsVoipSandboxChannel(ctx context.Context, params *UpdateApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*UpdateApnsVoipSandboxChannelOutput, error) {
	if params == nil {
		params = &UpdateApnsVoipSandboxChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApnsVoipSandboxChannel", params, optFns, addOperationUpdateApnsVoipSandboxChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApnsVoipSandboxChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApnsVoipSandboxChannelInput struct {

	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// VoIP sandbox channel for an application.
	//
	// This member is required.
	APNSVoipSandboxChannelRequest *types.APNSVoipSandboxChannelRequest

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type UpdateApnsVoipSandboxChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP sandbox channel for an application.
	//
	// This member is required.
	APNSVoipSandboxChannelResponse *types.APNSVoipSandboxChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateApnsVoipSandboxChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApnsVoipSandboxChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApnsVoipSandboxChannel{}, middleware.After)
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
	if err = addOpUpdateApnsVoipSandboxChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApnsVoipSandboxChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApnsVoipSandboxChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateApnsVoipSandboxChannel",
	}
}
