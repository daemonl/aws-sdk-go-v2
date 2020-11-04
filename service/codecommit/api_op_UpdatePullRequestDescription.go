// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the contents of the description of a pull request.
func (c *Client) UpdatePullRequestDescription(ctx context.Context, params *UpdatePullRequestDescriptionInput, optFns ...func(*Options)) (*UpdatePullRequestDescriptionOutput, error) {
	if params == nil {
		params = &UpdatePullRequestDescriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePullRequestDescription", params, optFns, addOperationUpdatePullRequestDescriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePullRequestDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePullRequestDescriptionInput struct {

	// The updated content of the description for the pull request. This content
	// replaces the existing description.
	//
	// This member is required.
	Description *string

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests.
	//
	// This member is required.
	PullRequestId *string
}

type UpdatePullRequestDescriptionOutput struct {

	// Information about the updated pull request.
	//
	// This member is required.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdatePullRequestDescriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePullRequestDescription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePullRequestDescription{}, middleware.After)
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
	if err = addOpUpdatePullRequestDescriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePullRequestDescription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePullRequestDescription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdatePullRequestDescription",
	}
}
