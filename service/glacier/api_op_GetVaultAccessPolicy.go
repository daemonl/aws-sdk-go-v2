// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation retrieves the access-policy subresource set on the vault; for
// more information on setting this subresource, see Set Vault Access Policy (PUT
// access-policy)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-SetVaultAccessPolicy.html).
// If there is no access policy set on the vault, the operation returns a 404 Not
// found error. For more information about vault access policies, see Amazon
// Glacier Access Control with Vault Access Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html).
func (c *Client) GetVaultAccessPolicy(ctx context.Context, params *GetVaultAccessPolicyInput, optFns ...func(*Options)) (*GetVaultAccessPolicyOutput, error) {
	if params == nil {
		params = &GetVaultAccessPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVaultAccessPolicy", params, optFns, addOperationGetVaultAccessPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVaultAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetVaultAccessPolicy.
type GetVaultAccessPolicyInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

// Output for GetVaultAccessPolicy.
type GetVaultAccessPolicyOutput struct {

	// Contains the returned vault access policy as a JSON string.
	Policy *types.VaultAccessPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVaultAccessPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVaultAccessPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVaultAccessPolicy{}, middleware.After)
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
	if err = addOpGetVaultAccessPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVaultAccessPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetVaultAccessPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "GetVaultAccessPolicy",
	}
}
