// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets summary information about a domain configuration. The domain configuration
// feature is in public preview and is subject to change.
func (c *Client) DescribeDomainConfiguration(ctx context.Context, params *DescribeDomainConfigurationInput, optFns ...func(*Options)) (*DescribeDomainConfigurationOutput, error) {
	if params == nil {
		params = &DescribeDomainConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainConfiguration", params, optFns, addOperationDescribeDomainConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainConfigurationInput struct {

	// The name of the domain configuration.
	//
	// This member is required.
	DomainConfigurationName *string
}

type DescribeDomainConfigurationOutput struct {

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *types.AuthorizerConfig

	// The ARN of the domain configuration.
	DomainConfigurationArn *string

	// The name of the domain configuration.
	DomainConfigurationName *string

	// A Boolean value that specifies the current state of the domain configuration.
	DomainConfigurationStatus types.DomainConfigurationStatus

	// The name of the domain.
	DomainName *string

	// The type of the domain.
	DomainType types.DomainType

	// The date and time the domain configuration's status was last changed.
	LastStatusChangeDate *time.Time

	// A list containing summary information about the server certificate included in
	// the domain configuration.
	ServerCertificates []*types.ServerCertificateSummary

	// The type of service delivered by the endpoint.
	ServiceType types.ServiceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDomainConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDomainConfiguration{}, middleware.After)
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
	if err = addOpDescribeDomainConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomainConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeDomainConfiguration",
	}
}
