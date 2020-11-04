// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a set of permissions for the specified folder or document. The resource
// permissions are overwritten if the principals already have different
// permissions.
func (c *Client) AddResourcePermissions(ctx context.Context, params *AddResourcePermissionsInput, optFns ...func(*Options)) (*AddResourcePermissionsOutput, error) {
	if params == nil {
		params = &AddResourcePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddResourcePermissions", params, optFns, addOperationAddResourcePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddResourcePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddResourcePermissionsInput struct {

	// The users, groups, or organization being granted permission.
	//
	// This member is required.
	Principals []*types.SharePrincipal

	// The ID of the resource.
	//
	// This member is required.
	ResourceId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The notification options.
	NotificationOptions *types.NotificationOptions
}

type AddResourcePermissionsOutput struct {

	// The share results.
	ShareResults []*types.ShareResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddResourcePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddResourcePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddResourcePermissions{}, middleware.After)
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
	if err = addOpAddResourcePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddResourcePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddResourcePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "AddResourcePermissions",
	}
}
