// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a permission to a queue for a specific principal
// (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P). This allows
// sharing access to the queue. When you create a queue, you have full control
// access rights for the queue. Only you, the owner of the queue, can grant or deny
// permissions to the queue. For more information about these permissions, see
// Allow Developers to Write Messages to a Shared Queue
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
//
// * AddPermission generates a
// policy for you. You can use SetQueueAttributes to upload your policy. For more
// information, see Using Custom Policies with the Amazon SQS Access Policy
// Language
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// * An Amazon SQS policy can
// have a maximum of 7 actions.
//
// * To remove the ability to change queue
// permissions, you must deny permission to the AddPermission, RemovePermission,
// and SetQueueAttributes actions in your IAM policy.
//
// Some actions take lists of
// parameters. These lists are specified using the param.n notation. Values of n
// are integers starting from 1. For example, a parameter list with two elements
// looks like this: &AttributeName.1=first&AttributeName.2=second Cross-account
// permissions don't apply to this action. For more information, see Grant
// Cross-Account Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) {
	if params == nil {
		params = &AddPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPermission", params, optFns, addOperationAddPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type AddPermissionInput struct {

	// The AWS account number of the principal
	// (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P) who is given
	// permission. The principal must have an AWS account, but does not need to be
	// signed up for Amazon SQS. For information about locating the AWS account
	// identification, see Your AWS Identifiers
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-making-api-requests.html#sqs-api-request-authentication)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// This member is required.
	AWSAccountIds []*string

	// The action the client wants to allow for the specified principal. Valid values:
	// the name of any action or *. For more information about these actions, see
	// Overview of Managing Access Permissions to Your Amazon Simple Queue Service
	// Resource
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-overview-of-managing-access.html)
	// in the Amazon Simple Queue Service Developer Guide. Specifying SendMessage,
	// DeleteMessage, or ChangeMessageVisibility for ActionName.n also grants
	// permissions for the corresponding batch versions of those actions:
	// SendMessageBatch, DeleteMessageBatch, and ChangeMessageVisibilityBatch.
	//
	// This member is required.
	Actions []*string

	// The unique identification of the permission you're setting (for example,
	// AliceSendMessage). Maximum 80 characters. Allowed characters include
	// alphanumeric characters, hyphens (-), and underscores (_).
	//
	// This member is required.
	Label *string

	// The URL of the Amazon SQS queue to which permissions are added. Queue URLs and
	// names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

type AddPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddPermission{}, middleware.After)
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
	if err = addOpAddPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "AddPermission",
	}
}
