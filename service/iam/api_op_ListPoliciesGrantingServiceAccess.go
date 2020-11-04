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

// Retrieves a list of policies that the IAM identity (user, group, or role) can
// use to access each specified service. This operation does not use other policy
// types when determining whether a resource could access a service. These other
// policy types include resource-based policies, access control lists, AWS
// Organizations policies, IAM permissions boundaries, and AWS STS assume role
// policies. It only applies permissions policy logic. For more about the
// evaluation of policy types, see Evaluating Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide. The list of policies returned by the operation depends on
// the ARN of the identity that you provide.
//
// * User – The list of policies
// includes the managed and inline policies that are attached to the user directly.
// The list also includes any additional managed and inline policies that are
// attached to the group to which the user belongs.
//
// * Group – The list of policies
// includes only the managed and inline policies that are attached to the group
// directly. Policies that are attached to the group’s user are not included.
//
// *
// Role – The list of policies includes only the managed and inline policies that
// are attached to the role.
//
// For each managed policy, this operation returns the
// ARN and policy name. For each inline policy, it returns the policy name and the
// entity to which it is attached. Inline policies do not have an ARN. For more
// information about these policy types, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html)
// in the IAM User Guide. Policies that are attached to users and roles as
// permissions boundaries are not returned. To view which managed policy is
// currently used to set the permissions boundary for a user or role, use the
// GetUser or GetRole operations.
func (c *Client) ListPoliciesGrantingServiceAccess(ctx context.Context, params *ListPoliciesGrantingServiceAccessInput, optFns ...func(*Options)) (*ListPoliciesGrantingServiceAccessOutput, error) {
	if params == nil {
		params = &ListPoliciesGrantingServiceAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPoliciesGrantingServiceAccess", params, optFns, addOperationListPoliciesGrantingServiceAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesGrantingServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesGrantingServiceAccessInput struct {

	// The ARN of the IAM identity (user, group, or role) whose policies you want to
	// list.
	//
	// This member is required.
	Arn *string

	// The service namespace for the AWS services whose policies you want to list. To
	// learn the service namespace for a service, go to Actions, Resources, and
	// Condition Keys for AWS Services
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_actions-resources-contextkeys.html)
	// in the IAM User Guide. Choose the name of the service to view details for that
	// service. In the first paragraph, find the service prefix. For example, (service
	// prefix: a4b). For more information about service namespaces, see AWS Service
	// Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the AWS General Reference.
	//
	// This member is required.
	ServiceNamespaces []*string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string
}

type ListPoliciesGrantingServiceAccessOutput struct {

	// A ListPoliciesGrantingServiceAccess object that contains details about the
	// permissions policies attached to the specified identity (user, group, or role).
	//
	// This member is required.
	PoliciesGrantingServiceAccess []*types.ListPoliciesGrantingServiceAccessEntry

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. We recommend that you check
	// IsTruncated after every call to ensure that you receive all your results.
	IsTruncated *bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPoliciesGrantingServiceAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPoliciesGrantingServiceAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPoliciesGrantingServiceAccess{}, middleware.After)
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
	if err = addOpListPoliciesGrantingServiceAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPoliciesGrantingServiceAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPoliciesGrantingServiceAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListPoliciesGrantingServiceAccess",
	}
}
