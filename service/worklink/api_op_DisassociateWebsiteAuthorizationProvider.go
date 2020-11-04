// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a website authorization provider from a specified fleet. After the
// disassociation, users can't load any associated websites that require this
// authorization provider.
func (c *Client) DisassociateWebsiteAuthorizationProvider(ctx context.Context, params *DisassociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*DisassociateWebsiteAuthorizationProviderOutput, error) {
	if params == nil {
		params = &DisassociateWebsiteAuthorizationProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateWebsiteAuthorizationProvider", params, optFns, addOperationDisassociateWebsiteAuthorizationProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateWebsiteAuthorizationProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateWebsiteAuthorizationProviderInput struct {

	// A unique identifier for the authorization provider.
	//
	// This member is required.
	AuthorizationProviderId *string

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string
}

type DisassociateWebsiteAuthorizationProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateWebsiteAuthorizationProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
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
	if err = addOpDisassociateWebsiteAuthorizationProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DisassociateWebsiteAuthorizationProvider",
	}
}
