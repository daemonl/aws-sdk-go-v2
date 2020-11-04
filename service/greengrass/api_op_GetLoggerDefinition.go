// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a logger definition.
func (c *Client) GetLoggerDefinition(ctx context.Context, params *GetLoggerDefinitionInput, optFns ...func(*Options)) (*GetLoggerDefinitionOutput, error) {
	if params == nil {
		params = &GetLoggerDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoggerDefinition", params, optFns, addOperationGetLoggerDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoggerDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggerDefinitionInput struct {

	// The ID of the logger definition.
	//
	// This member is required.
	LoggerDefinitionId *string
}

type GetLoggerDefinitionOutput struct {

	// The ARN of the definition.
	Arn *string

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string

	// The ID of the definition.
	Id *string

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string

	// The ID of the latest version associated with the definition.
	LatestVersion *string

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// The name of the definition.
	Name *string

	// Tag(s) attached to the resource arn.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLoggerDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLoggerDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLoggerDefinition{}, middleware.After)
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
	if err = addOpGetLoggerDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoggerDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoggerDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetLoggerDefinition",
	}
}
