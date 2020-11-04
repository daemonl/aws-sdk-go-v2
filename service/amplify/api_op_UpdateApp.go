// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing Amplify app.
func (c *Client) UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) {
	if params == nil {
		params = &UpdateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApp", params, optFns, addOperationUpdateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the update app request.
type UpdateAppInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The personal access token for a third-party source control system for an Amplify
	// app. The token is used to create webhook and a read-only deploy key. The token
	// is not stored.
	AccessToken *string

	// The automated branch creation configuration for the Amplify app.
	AutoBranchCreationConfig *types.AutoBranchCreationConfig

	// Describes the automated branch creation glob patterns for the Amplify app.
	AutoBranchCreationPatterns []*string

	// The basic authorization credentials for an Amplify app.
	BasicAuthCredentials *string

	// The build specification (build spec) for an Amplify app.
	BuildSpec *string

	// The custom redirect and rewrite rules for an Amplify app.
	CustomRules []*types.CustomRule

	// The description for an Amplify app.
	Description *string

	// Enables automated branch creation for the Amplify app.
	EnableAutoBranchCreation *bool

	// Enables basic authorization for an Amplify app.
	EnableBasicAuth *bool

	// Enables branch auto-building for an Amplify app.
	EnableBranchAutoBuild *bool

	// Automatically disconnects a branch in the Amplify Console when you delete a
	// branch from your Git repository.
	EnableBranchAutoDeletion *bool

	// The environment variables for an Amplify app.
	EnvironmentVariables map[string]*string

	// The AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn *string

	// The name for an Amplify app.
	Name *string

	// The OAuth token for a third-party source control system for an Amplify app. The
	// token is used to create a webhook and a read-only deploy key. The OAuth token is
	// not stored.
	OauthToken *string

	// The platform for an Amplify app.
	Platform types.Platform

	// The name of the repository for an Amplify app
	Repository *string
}

// The result structure for an Amplify app update request.
type UpdateAppOutput struct {

	// Represents the updated Amplify app.
	//
	// This member is required.
	App *types.App

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApp{}, middleware.After)
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
	if err = addOpUpdateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "UpdateApp",
	}
}
