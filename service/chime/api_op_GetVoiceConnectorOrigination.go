// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves origination setting details for the specified Amazon Chime Voice
// Connector.
func (c *Client) GetVoiceConnectorOrigination(ctx context.Context, params *GetVoiceConnectorOriginationInput, optFns ...func(*Options)) (*GetVoiceConnectorOriginationOutput, error) {
	if params == nil {
		params = &GetVoiceConnectorOriginationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVoiceConnectorOrigination", params, optFns, addOperationGetVoiceConnectorOriginationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVoiceConnectorOriginationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorOriginationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type GetVoiceConnectorOriginationOutput struct {

	// The origination setting details.
	Origination *types.Origination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVoiceConnectorOriginationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorOrigination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorOrigination{}, middleware.After)
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
	if err = addOpGetVoiceConnectorOriginationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorOrigination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVoiceConnectorOrigination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorOrigination",
	}
}
