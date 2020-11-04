// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Creates a WebACL per the specifications provided. A Web ACL defines a collection
// of rules to use to inspect and control web requests. Each rule has an action
// defined (allow, block, or count) for requests that match the statement of the
// rule. In the Web ACL, you assign a default action to take (allow, block) for any
// request that does not match any of the rules. The rules in a Web ACL can be a
// combination of the types Rule, RuleGroup, and managed rule group. You can
// associate a Web ACL with one or more AWS resources to protect. The resources can
// be Amazon CloudFront, an Amazon API Gateway REST API, an Application Load
// Balancer, or an AWS AppSync GraphQL API.
func (c *Client) CreateWebACL(ctx context.Context, params *CreateWebACLInput, optFns ...func(*Options)) (*CreateWebACLOutput, error) {
	if params == nil {
		params = &CreateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebACL", params, optFns, addOperationCreateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebACLInput struct {

	// The action to perform if none of the Rules contained in the WebACL match.
	//
	// This member is required.
	DefaultAction *types.DefaultAction

	// The name of the Web ACL. You cannot change the name of a Web ACL after you
	// create it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	// * CLI - Specify
	// the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// This member is required.
	VisibilityConfig *types.VisibilityConfig

	// A description of the Web ACL that helps with identification. You cannot change
	// the description of a Web ACL after you create it.
	Description *string

	// The Rule statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that AWS WAF uses to
	// identify matching web requests, and parameters that govern how AWS WAF handles
	// them.
	Rules []*types.Rule

	// An array of key:value pairs to associate with the resource.
	Tags []*types.Tag
}

type CreateWebACLOutput struct {

	// High-level information about a WebACL, returned by operations like create and
	// list. This provides information like the ID, that you can use to retrieve and
	// manage a WebACL, and the ARN, that you provide to operations like
	// AssociateWebACL.
	Summary *types.WebACLSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWebACL{}, middleware.After)
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
	if err = addOpCreateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateWebACL",
	}
}
