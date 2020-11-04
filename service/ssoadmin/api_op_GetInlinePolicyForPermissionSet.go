// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Obtains the inline policy assigned to the permission set.
func (c *Client) GetInlinePolicyForPermissionSet(ctx context.Context, params *GetInlinePolicyForPermissionSetInput, optFns ...func(*Options)) (*GetInlinePolicyForPermissionSetOutput, error) {
	if params == nil {
		params = &GetInlinePolicyForPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInlinePolicyForPermissionSet", params, optFns, addOperationGetInlinePolicyForPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInlinePolicyForPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInlinePolicyForPermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set.
	//
	// This member is required.
	PermissionSetArn *string
}

type GetInlinePolicyForPermissionSetOutput struct {

	// The IAM inline policy that is attached to the permission set.
	InlinePolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInlinePolicyForPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInlinePolicyForPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInlinePolicyForPermissionSet{}, middleware.After)
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
	if err = addOpGetInlinePolicyForPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInlinePolicyForPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInlinePolicyForPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "GetInlinePolicyForPermissionSet",
	}
}
