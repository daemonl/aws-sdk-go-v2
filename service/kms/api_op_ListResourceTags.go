// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all tags for the specified customer master key (CMK). You
// cannot perform this operation on a CMK in a different AWS account.
func (c *Client) ListResourceTags(ctx context.Context, params *ListResourceTagsInput, optFns ...func(*Options)) (*ListResourceTagsOutput, error) {
	if params == nil {
		params = &ListResourceTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceTags", params, optFns, addOperationListResourceTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceTagsInput struct {

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, AWS KMS does not return more than the specified number of
	// items, but it might return fewer. This value is optional. If you include a
	// value, it must be between 1 and 50, inclusive. If you do not include a value, it
	// defaults to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received. Do not attempt to construct this value. Use only the value of
	// NextMarker from the truncated response you just received.
	Marker *string
}

type ListResourceTagsOutput struct {

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request. Do not assume or infer any
	// information from this value.
	NextMarker *string

	// A list of tags. Each tag consists of a tag key and a tag value.
	Tags []*types.Tag

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResourceTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceTags{}, middleware.After)
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
	if err = addOpListResourceTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListResourceTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListResourceTags",
	}
}
