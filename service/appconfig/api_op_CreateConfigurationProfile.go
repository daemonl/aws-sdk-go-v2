// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Information that enables AppConfig to access the configuration source. Valid
// configuration sources include Systems Manager (SSM) documents, SSM Parameter
// Store parameters, and Amazon S3 objects. A configuration profile includes the
// following information.
//
// * The Uri location of the configuration data.
//
// * The AWS
// Identity and Access Management (IAM) role that provides access to the
// configuration data.
//
// * A validator for the configuration data. Available
// validators include either a JSON Schema or an AWS Lambda function.
//
// For more
// information, see Create a Configuration and a Configuration Profile
// (http://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-creating-configuration-and-profile.html)
// in the AWS AppConfig User Guide.
func (c *Client) CreateConfigurationProfile(ctx context.Context, params *CreateConfigurationProfileInput, optFns ...func(*Options)) (*CreateConfigurationProfileOutput, error) {
	if params == nil {
		params = &CreateConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationProfile", params, optFns, addOperationCreateConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfigurationProfileInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// A URI to locate the configuration. You can specify a Systems Manager (SSM)
	// document, an SSM Parameter Store parameter, or an Amazon S3 object. For an SSM
	// document, specify either the document name in the format ssm-document:// or the
	// Amazon Resource Name (ARN). For a parameter, specify either the parameter name
	// in the format ssm-parameter:// or the ARN. For an Amazon S3 object, specify the
	// URI in the following format: s3:/// . Here is an example:
	// s3://my-bucket/my-app/us-east-1/my-config.json
	//
	// This member is required.
	LocationUri *string

	// A name for the configuration profile.
	//
	// This member is required.
	Name *string

	// A description of the configuration profile.
	Description *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri.
	RetrievalRoleArn *string

	// Metadata to assign to the configuration profile. Tags help organize and
	// categorize your AppConfig resources. Each tag consists of a key and an optional
	// value, both of which you define.
	Tags map[string]*string

	// A list of methods for validating the configuration.
	Validators []*types.Validator
}

type CreateConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile description.
	Description *string

	// The configuration profile ID.
	Id *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri.
	RetrievalRoleArn *string

	// A list of methods for validating the configuration.
	Validators []*types.Validator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfigurationProfile{}, middleware.After)
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
	if err = addOpCreateConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateConfigurationProfile",
	}
}
