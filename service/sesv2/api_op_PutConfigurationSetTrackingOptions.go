// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specify a custom domain to use for open and click tracking elements in email
// that you send.
func (c *Client) PutConfigurationSetTrackingOptions(ctx context.Context, params *PutConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetTrackingOptionsOutput, error) {
	if params == nil {
		params = &PutConfigurationSetTrackingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationSetTrackingOptions", params, optFns, addOperationPutConfigurationSetTrackingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add a custom domain for tracking open and click events to a
// configuration set.
type PutConfigurationSetTrackingOptionsInput struct {

	// The name of the configuration set that you want to add a custom tracking domain
	// to.
	//
	// This member is required.
	ConfigurationSetName *string

	// The domain that you want to use to track open and click events.
	CustomRedirectDomain *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetTrackingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetTrackingOptions{}, middleware.After)
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
	if err = addOpPutConfigurationSetTrackingOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetTrackingOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationSetTrackingOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetTrackingOptions",
	}
}
