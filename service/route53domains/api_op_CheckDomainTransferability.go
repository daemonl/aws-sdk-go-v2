// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Checks whether a domain name can be transferred to Amazon Route 53.
func (c *Client) CheckDomainTransferability(ctx context.Context, params *CheckDomainTransferabilityInput, optFns ...func(*Options)) (*CheckDomainTransferabilityOutput, error) {
	if params == nil {
		params = &CheckDomainTransferabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckDomainTransferability", params, optFns, addOperationCheckDomainTransferabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckDomainTransferabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CheckDomainTransferability request contains the following elements.
type CheckDomainTransferabilityInput struct {

	// The name of the domain that you want to transfer to Route 53. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list of
	// supported TLDs, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. The domain name can contain only the
	// following characters:
	//
	// * Letters a through z. Domain names are not case
	// sensitive.
	//
	// * Numbers 0 through 9.
	//
	// * Hyphen (-). You can't specify a hyphen at
	// the beginning or end of a label.
	//
	// * Period (.) to separate the labels in the
	// name, such as the . in example.com.
	//
	// This member is required.
	DomainName *string

	// If the registrar for the top-level domain (TLD) requires an authorization code
	// to transfer the domain, the code that you got from the current registrar for the
	// domain.
	AuthCode *string
}

// The CheckDomainTransferability response includes the following elements.
type CheckDomainTransferabilityOutput struct {

	// A complex type that contains information about whether the specified domain can
	// be transferred to Route 53.
	//
	// This member is required.
	Transferability *types.DomainTransferability

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCheckDomainTransferabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCheckDomainTransferability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckDomainTransferability{}, middleware.After)
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
	if err = addOpCheckDomainTransferabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckDomainTransferability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckDomainTransferability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "CheckDomainTransferability",
	}
}
