// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an IAM role that is linked to a specific AWS service. The service
// controls the attached policies and when the role can be deleted. This helps
// ensure that the service is not broken by an unexpectedly changed or deleted
// role, which could put your AWS resources into an unknown state. Allowing the
// service to control the role helps improve service stability and proper cleanup
// when a service and its role are no longer needed. For more information, see
// Using Service-Linked Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in the IAM User Guide. To attach a policy to this service-linked role, you must
// make the request using the AWS service that depends on this role.
func (c *Client) CreateServiceLinkedRole(ctx context.Context, params *CreateServiceLinkedRoleInput, optFns ...func(*Options)) (*CreateServiceLinkedRoleOutput, error) {
	if params == nil {
		params = &CreateServiceLinkedRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceLinkedRole", params, optFns, addOperationCreateServiceLinkedRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceLinkedRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceLinkedRoleInput struct {

	// The service principal for the AWS service to which this role is attached. You
	// use a string similar to a URL but without the http:// in front. For example:
	// elasticbeanstalk.amazonaws.com. Service principals are unique and
	// case-sensitive. To find the exact service principal for your service-linked
	// role, see AWS Services That Work with IAM
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html)
	// in the IAM User Guide. Look for the services that have Yes in the Service-Linked
	// Role column. Choose the Yes link to view the service-linked role documentation
	// for that service.
	//
	// This member is required.
	AWSServiceName *string

	// A string that you provide, which is combined with the service-provided prefix to
	// form the complete role name. If you make multiple requests for the same service,
	// then you must supply a different CustomSuffix for each request. Otherwise the
	// request fails with a duplicate role name error. For example, you could add -1 or
	// -debug to the suffix. Some services do not support the CustomSuffix parameter.
	// If you provide an optional suffix and the operation fails, try the operation
	// again without the suffix.
	CustomSuffix *string

	// The description of the role.
	Description *string
}

type CreateServiceLinkedRoleOutput struct {

	// A Role object that contains details about the newly created role.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateServiceLinkedRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateServiceLinkedRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateServiceLinkedRole{}, middleware.After)
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
	if err = addOpCreateServiceLinkedRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceLinkedRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServiceLinkedRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateServiceLinkedRole",
	}
}
