// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation sets the transfer lock on the domain (specifically the
// clientTransferProhibited status) to prevent domain transfers. Successful
// submission returns an operation ID that you can use to track the progress and
// completion of the action. If the request is not completed successfully, the
// domain registrant will be notified by email.
func (c *Client) EnableDomainTransferLock(ctx context.Context, params *EnableDomainTransferLockInput, optFns ...func(*Options)) (*EnableDomainTransferLockOutput, error) {
	if params == nil {
		params = &EnableDomainTransferLockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableDomainTransferLock", params, optFns, addOperationEnableDomainTransferLockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableDomainTransferLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to set the transfer lock for the specified domain.
type EnableDomainTransferLockInput struct {

	// The name of the domain that you want to set the transfer lock for.
	//
	// This member is required.
	DomainName *string
}

// The EnableDomainTransferLock response includes the following elements.
type EnableDomainTransferLockOutput struct {

	// Identifier for tracking the progress of the request. To use this ID to query the
	// operation status, use GetOperationDetail.
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableDomainTransferLockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableDomainTransferLock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableDomainTransferLock{}, middleware.After)
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
	if err = addOpEnableDomainTransferLockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableDomainTransferLock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableDomainTransferLock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "EnableDomainTransferLock",
	}
}
