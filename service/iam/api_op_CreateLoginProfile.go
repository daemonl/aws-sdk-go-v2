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

// Creates a password for the specified user, giving the user the ability to access
// AWS services through the AWS Management Console. For more information about
// managing passwords, see Managing Passwords
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in
// the IAM User Guide.
func (c *Client) CreateLoginProfile(ctx context.Context, params *CreateLoginProfileInput, optFns ...func(*Options)) (*CreateLoginProfileOutput, error) {
	if params == nil {
		params = &CreateLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoginProfile", params, optFns, addOperationCreateLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoginProfileInput struct {

	// The new password for the user. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of characters. That string can include almost any printable ASCII
	// character from the space (\u0020) through the end of the ASCII character range
	// (\u00FF). You can also include the tab (\u0009), line feed (\u000A), and
	// carriage return (\u000D) characters. Any of these characters are valid in a
	// password. However, many tools, such as the AWS Management Console, might
	// restrict the ability to type certain characters because they have special
	// meaning within that tool.
	//
	// This member is required.
	Password *string

	// The name of the IAM user to create a password for. The user must already exist.
	// This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	// Specifies whether the user is required to set a new password on next sign-in.
	PasswordResetRequired *bool
}

// Contains the response to a successful CreateLoginProfile request.
type CreateLoginProfileOutput struct {

	// A structure containing the user name and password create date.
	//
	// This member is required.
	LoginProfile *types.LoginProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoginProfile{}, middleware.After)
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
	if err = addOpCreateLoginProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoginProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoginProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateLoginProfile",
	}
}
