// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports a component and transforms its data into a component document.
func (c *Client) ImportComponent(ctx context.Context, params *ImportComponentInput, optFns ...func(*Options)) (*ImportComponentOutput, error) {
	if params == nil {
		params = &ImportComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportComponent", params, optFns, addOperationImportComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportComponentInput struct {

	// The idempotency token of the component.
	//
	// This member is required.
	ClientToken *string

	// The format of the resource that you want to import as a component.
	//
	// This member is required.
	Format types.ComponentFormat

	// The name of the component.
	//
	// This member is required.
	Name *string

	// The platform of the component.
	//
	// This member is required.
	Platform types.Platform

	// The semantic version of the component. This version follows the semantic version
	// syntax. For example, major.minor.patch. This could be versioned like software
	// (2.0.1) or like a date (2019.12.01).
	//
	// This member is required.
	SemanticVersion *string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	//
	// This member is required.
	Type types.ComponentType

	// The change description of the component. Describes what change has been made in
	// this version, or what makes this version different from other versions of this
	// component.
	ChangeDescription *string

	// The data of the component. Used to specify the data inline. Either data or uri
	// can be used to specify the data within the component.
	Data *string

	// The description of the component. Describes the contents of the component.
	Description *string

	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string

	// The tags of the component.
	Tags map[string]*string

	// The uri of the component. Must be an S3 URL and the requester must have
	// permission to access the S3 bucket. If you use S3, you can specify component
	// content up to your service quota. Either data or uri can be used to specify the
	// data within the component.
	Uri *string
}

type ImportComponentOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the imported component.
	ComponentBuildVersionArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportComponent{}, middleware.After)
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
	if err = addIdempotencyToken_opImportComponentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportComponent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpImportComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportComponentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opImportComponentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ImportComponent",
	}
}
