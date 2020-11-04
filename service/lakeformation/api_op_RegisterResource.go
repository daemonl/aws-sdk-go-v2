// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers the resource as managed by the Data Catalog. To add or update data,
// Lake Formation needs read/write access to the chosen Amazon S3 path. Choose a
// role that you know has permission to do this, or choose the
// AWSServiceRoleForLakeFormationDataAccess service-linked role. When you register
// the first Amazon S3 path, the service-linked role and a new inline policy are
// created on your behalf. Lake Formation adds the first path to the inline policy
// and attaches it to the service-linked role. When you register subsequent paths,
// Lake Formation adds the path to the existing policy. The following request
// registers a new location and gives AWS Lake Formation permission to use the
// service-linked role to access that location. ResourceArn =
// arn:aws:s3:::my-bucket UseServiceLinkedRole = true If UseServiceLinkedRole is
// not set to true, you must provide or set the RoleArn:
// arn:aws:iam::12345:role/my-data-access-role
func (c *Client) RegisterResource(ctx context.Context, params *RegisterResourceInput, optFns ...func(*Options)) (*RegisterResourceOutput, error) {
	if params == nil {
		params = &RegisterResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterResource", params, optFns, addOperationRegisterResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource that you want to register.
	//
	// This member is required.
	ResourceArn *string

	// The identifier for the role that registers the resource.
	RoleArn *string

	// Designates an AWS Identity and Access Management (IAM) service-linked role by
	// registering this role with the Data Catalog. A service-linked role is a unique
	// type of IAM role that is linked directly to Lake Formation. For more
	// information, see Using Service-Linked Roles for Lake Formation
	// (https://docs-aws.amazon.com/lake-formation/latest/dg/service-linked-roles.html).
	UseServiceLinkedRole *bool
}

type RegisterResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterResource{}, middleware.After)
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
	if err = addOpRegisterResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "RegisterResource",
	}
}
