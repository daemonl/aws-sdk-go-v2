// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of sending authorization policies that are attached to the given
// identity (an email address or a domain). This API returns only a list. If you
// want the actual policy content, you can use GetIdentityPolicies. This API is for
// the identity owner only. If you have not verified the identity, this API will
// return an error. Sending authorization is a feature that enables an identity
// owner to authorize other senders to use its identities. For information about
// using sending authorization, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// You can execute this operation no more than once per second.
func (c *Client) ListIdentityPolicies(ctx context.Context, params *ListIdentityPoliciesInput, optFns ...func(*Options)) (*ListIdentityPoliciesOutput, error) {
	if params == nil {
		params = &ListIdentityPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityPolicies", params, optFns, addOperationListIdentityPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return a list of sending authorization policies that are
// attached to an identity. Sending authorization is an Amazon SES feature that
// enables you to authorize other senders to use your identities. For information,
// see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type ListIdentityPoliciesInput struct {

	// The identity that is associated with the policy for which the policies will be
	// listed. You can specify an identity by using its name or by using its Amazon
	// Resource Name (ARN). Examples: user@example.com, example.com,
	// arn:aws:ses:us-east-1:123456789012:identity/example.com. To successfully call
	// this API, you must own the identity.
	//
	// This member is required.
	Identity *string
}

// A list of names of sending authorization policies that apply to an identity.
type ListIdentityPoliciesOutput struct {

	// A list of names of policies that apply to the specified identity.
	//
	// This member is required.
	PolicyNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIdentityPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListIdentityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListIdentityPolicies{}, middleware.After)
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
	if err = addOpListIdentityPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIdentityPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListIdentityPolicies",
	}
}
