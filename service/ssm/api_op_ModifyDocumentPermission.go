// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Shares a Systems Manager document publicly or privately. If you share a document
// privately, you must specify the AWS user account IDs for those people who can
// use the document. If you share a document publicly, you must specify All as the
// account ID.
func (c *Client) ModifyDocumentPermission(ctx context.Context, params *ModifyDocumentPermissionInput, optFns ...func(*Options)) (*ModifyDocumentPermissionOutput, error) {
	if params == nil {
		params = &ModifyDocumentPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDocumentPermission", params, optFns, addOperationModifyDocumentPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDocumentPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDocumentPermissionInput struct {

	// The name of the document that you want to share.
	//
	// This member is required.
	Name *string

	// The permission type for the document. The permission type can be Share.
	//
	// This member is required.
	PermissionType types.DocumentPermissionType

	// The AWS user accounts that should have access to the document. The account IDs
	// can either be a group of account IDs or All.
	AccountIdsToAdd []*string

	// The AWS user accounts that should no longer have access to the document. The AWS
	// user account can either be a group of account IDs or All. This action has a
	// higher priority than AccountIdsToAdd. If you specify an account ID to add and
	// the same ID to remove, the system removes access to the document.
	AccountIdsToRemove []*string

	// (Optional) The version of the document to share. If it's not specified, the
	// system choose the Default version to share.
	SharedDocumentVersion *string
}

type ModifyDocumentPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDocumentPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyDocumentPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyDocumentPermission{}, middleware.After)
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
	if err = addOpModifyDocumentPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDocumentPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDocumentPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ModifyDocumentPermission",
	}
}
