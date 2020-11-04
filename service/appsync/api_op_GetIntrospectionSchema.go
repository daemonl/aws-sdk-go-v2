// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the introspection schema for a GraphQL API.
func (c *Client) GetIntrospectionSchema(ctx context.Context, params *GetIntrospectionSchemaInput, optFns ...func(*Options)) (*GetIntrospectionSchemaOutput, error) {
	if params == nil {
		params = &GetIntrospectionSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntrospectionSchema", params, optFns, addOperationGetIntrospectionSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntrospectionSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntrospectionSchemaInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The schema format: SDL or JSON.
	//
	// This member is required.
	Format types.OutputType

	// A flag that specifies whether the schema introspection should contain
	// directives.
	IncludeDirectives *bool
}

type GetIntrospectionSchemaOutput struct {

	// The schema, in GraphQL Schema Definition Language (SDL) format. For more
	// information, see the GraphQL SDL documentation
	// (http://graphql.org/learn/schema/).
	Schema []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIntrospectionSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntrospectionSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntrospectionSchema{}, middleware.After)
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
	if err = addOpGetIntrospectionSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntrospectionSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIntrospectionSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "GetIntrospectionSchema",
	}
}
