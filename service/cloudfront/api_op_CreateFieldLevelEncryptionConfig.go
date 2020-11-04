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

// Create a new field-level encryption configuration.
func (c *Client) CreateFieldLevelEncryptionConfig(ctx context.Context, params *CreateFieldLevelEncryptionConfigInput, optFns ...func(*Options)) (*CreateFieldLevelEncryptionConfigOutput, error) {
	if params == nil {
		params = &CreateFieldLevelEncryptionConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFieldLevelEncryptionConfig", params, optFns, addOperationCreateFieldLevelEncryptionConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFieldLevelEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFieldLevelEncryptionConfigInput struct {

	// The request to create a new field-level encryption configuration.
	//
	// This member is required.
	FieldLevelEncryptionConfig *types.FieldLevelEncryptionConfig
}

type CreateFieldLevelEncryptionConfigOutput struct {

	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Returned when you create a new field-level encryption configuration.
	FieldLevelEncryption *types.FieldLevelEncryption

	// The fully qualified URI of the new configuration resource just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFieldLevelEncryptionConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateFieldLevelEncryptionConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateFieldLevelEncryptionConfig{}, middleware.After)
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
	if err = addOpCreateFieldLevelEncryptionConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFieldLevelEncryptionConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFieldLevelEncryptionConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateFieldLevelEncryptionConfig",
	}
}
