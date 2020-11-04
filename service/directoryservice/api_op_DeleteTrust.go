// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing trust relationship between your AWS Managed Microsoft AD
// directory and an external domain.
func (c *Client) DeleteTrust(ctx context.Context, params *DeleteTrustInput, optFns ...func(*Options)) (*DeleteTrustOutput, error) {
	if params == nil {
		params = &DeleteTrustInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrust", params, optFns, addOperationDeleteTrustMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrustOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Deletes the local side of an existing trust relationship between the AWS Managed
// Microsoft AD directory and the external domain.
type DeleteTrustInput struct {

	// The Trust ID of the trust relationship to be deleted.
	//
	// This member is required.
	TrustId *string

	// Delete a conditional forwarder as part of a DeleteTrustRequest.
	DeleteAssociatedConditionalForwarder *bool
}

// The result of a DeleteTrust request.
type DeleteTrustOutput struct {

	// The Trust ID of the trust relationship that was deleted.
	TrustId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTrustMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTrust{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTrust{}, middleware.After)
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
	if err = addOpDeleteTrustValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrust(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTrust(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeleteTrust",
	}
}
