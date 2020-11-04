// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all of the contact lists available.
func (c *Client) ListContactLists(ctx context.Context, params *ListContactListsInput, optFns ...func(*Options)) (*ListContactListsOutput, error) {
	if params == nil {
		params = &ListContactListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContactLists", params, optFns, addOperationListContactListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactListsInput struct {

	// A string token indicating that there might be additional contact lists available
	// to be listed. Use the token provided in the Response to use in the subsequent
	// call to ListContactLists with the same parameters to retrieve the next page of
	// contact lists.
	NextToken *string

	// Maximum number of contact lists to return at once. Use this parameter to
	// paginate results. If additional contact lists exist beyond the specified limit,
	// the NextToken element is sent in the response. Use the NextToken value in
	// subsequent requests to retrieve additional lists.
	PageSize *int32
}

type ListContactListsOutput struct {

	// The available contact lists.
	ContactLists []*types.ContactList

	// A string token indicating that there might be additional contact lists available
	// to be listed. Copy this token to a subsequent call to ListContactLists with the
	// same parameters to retrieve the next page of contact lists.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListContactListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContactLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactLists{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContactLists(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListContactLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListContactLists",
	}
}
