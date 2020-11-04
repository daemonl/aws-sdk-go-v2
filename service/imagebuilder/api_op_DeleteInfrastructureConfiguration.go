// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an infrastructure configuration.
func (c *Client) DeleteInfrastructureConfiguration(ctx context.Context, params *DeleteInfrastructureConfigurationInput, optFns ...func(*Options)) (*DeleteInfrastructureConfigurationOutput, error) {
	if params == nil {
		params = &DeleteInfrastructureConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInfrastructureConfiguration", params, optFns, addOperationDeleteInfrastructureConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInfrastructureConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInfrastructureConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration to delete.
	//
	// This member is required.
	InfrastructureConfigurationArn *string
}

type DeleteInfrastructureConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration that was
	// deleted.
	InfrastructureConfigurationArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteInfrastructureConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteInfrastructureConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteInfrastructureConfiguration{}, middleware.After)
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
	if err = addOpDeleteInfrastructureConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInfrastructureConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteInfrastructureConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "DeleteInfrastructureConfiguration",
	}
}
