// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The RejectQualificationRequest operation rejects a user's request for a
// Qualification. You can provide a text message explaining why the request was
// rejected. The Worker who made the request can see this message.
func (c *Client) RejectQualificationRequest(ctx context.Context, params *RejectQualificationRequestInput, optFns ...func(*Options)) (*RejectQualificationRequestOutput, error) {
	if params == nil {
		params = &RejectQualificationRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectQualificationRequest", params, optFns, addOperationRejectQualificationRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectQualificationRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectQualificationRequestInput struct {

	// The ID of the Qualification request, as returned by the
	// ListQualificationRequests operation.
	//
	// This member is required.
	QualificationRequestId *string

	// A text message explaining why the request was rejected, to be shown to the
	// Worker who made the request.
	Reason *string
}

type RejectQualificationRequestOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRejectQualificationRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRejectQualificationRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRejectQualificationRequest{}, middleware.After)
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
	if err = addOpRejectQualificationRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectQualificationRequest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRejectQualificationRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "RejectQualificationRequest",
	}
}
