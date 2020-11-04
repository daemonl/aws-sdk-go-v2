// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the ObjectIdentifiers to which a given policy is attached.
func (c *Client) ListPolicyAttachments(ctx context.Context, params *ListPolicyAttachmentsInput, optFns ...func(*Options)) (*ListPolicyAttachmentsOutput, error) {
	if params == nil {
		params = &ListPolicyAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicyAttachments", params, optFns, addOperationListPolicyAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPolicyAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyAttachmentsInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// objects reside. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the policy object.
	//
	// This member is required.
	PolicyReference *types.ObjectReference

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListPolicyAttachmentsOutput struct {

	// The pagination token.
	NextToken *string

	// A list of ObjectIdentifiers to which the policy is attached.
	ObjectIdentifiers []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPolicyAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyAttachments{}, middleware.After)
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
	if err = addOpListPolicyAttachmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyAttachments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPolicyAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListPolicyAttachments",
	}
}
