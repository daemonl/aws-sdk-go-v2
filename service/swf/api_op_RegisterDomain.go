// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a new domain. Access Control You can use IAM policies to control this
// action's access to Amazon SWF resources as follows:
//
// * You cannot use an IAM
// policy to control domain access for this action. The name of the domain being
// registered is available as the resource of this action.
//
// * Use an Action element
// to allow or deny permission to call this action.
//
// * You cannot use an IAM policy
// to constrain this action's parameters.
//
// If the caller doesn't have sufficient
// permissions to invoke the action, or the parameter values fall outside the
// specified constraints, the action fails. The associated event attribute's cause
// parameter is set to OPERATION_NOT_PERMITTED. For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RegisterDomain(ctx context.Context, params *RegisterDomainInput, optFns ...func(*Options)) (*RegisterDomainOutput, error) {
	if params == nil {
		params = &RegisterDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDomain", params, optFns, addOperationRegisterDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDomainInput struct {

	// Name of the domain to register. The name must be unique in the region that the
	// domain is registered in. The specified string must not start or end with
	// whitespace. It must not contain a : (colon), / (slash), | (vertical bar), or any
	// control characters (\u0000-\u001f | \u007f-\u009f). Also, it must not be the
	// literal string arn.
	//
	// This member is required.
	Name *string

	// The duration (in days) that records and histories of workflow executions on the
	// domain should be kept by the service. After the retention period, the workflow
	// execution isn't available in the results of visibility calls. If you pass the
	// value NONE or 0 (zero), then the workflow execution history isn't retained. As
	// soon as the workflow execution completes, the execution record and its history
	// are deleted. The maximum workflow execution retention period is 90 days. For
	// more information about Amazon SWF service limits, see: Amazon SWF Service Limits
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html)
	// in the Amazon SWF Developer Guide.
	//
	// This member is required.
	WorkflowExecutionRetentionPeriodInDays *string

	// A text description of the domain.
	Description *string

	// Tags to be added when registering a domain. Tags may only contain unicode
	// letters, digits, whitespace, or these symbols: _ . : / = + - @.
	Tags []*types.ResourceTag
}

type RegisterDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRegisterDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRegisterDomain{}, middleware.After)
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
	if err = addOpRegisterDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RegisterDomain",
	}
}
