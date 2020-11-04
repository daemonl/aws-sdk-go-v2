// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the roles for an identity pool. These roles are used when making calls to
// GetCredentialsForIdentity action. You must use AWS Developer credentials to call
// this API.
func (c *Client) SetIdentityPoolRoles(ctx context.Context, params *SetIdentityPoolRolesInput, optFns ...func(*Options)) (*SetIdentityPoolRolesOutput, error) {
	if params == nil {
		params = &SetIdentityPoolRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityPoolRoles", params, optFns, addOperationSetIdentityPoolRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityPoolRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the SetIdentityPoolRoles action.
type SetIdentityPoolRolesInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// The map of roles associated with this pool. For a given role, the key will be
	// either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	//
	// This member is required.
	Roles map[string]*string

	// How users for a specific identity provider are to mapped to roles. This is a
	// string to RoleMapping object map. The string identifies the identity provider,
	// for example, "graph.facebook.com" or
	// "cognito-idp-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Up to 25
	// rules can be specified per identity provider.
	RoleMappings map[string]*types.RoleMapping
}

type SetIdentityPoolRolesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetIdentityPoolRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetIdentityPoolRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetIdentityPoolRoles{}, middleware.After)
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
	if err = addOpSetIdentityPoolRolesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityPoolRoles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityPoolRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "SetIdentityPoolRoles",
	}
}
