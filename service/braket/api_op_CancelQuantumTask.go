// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels the specified task.
func (c *Client) CancelQuantumTask(ctx context.Context, params *CancelQuantumTaskInput, optFns ...func(*Options)) (*CancelQuantumTaskOutput, error) {
	if params == nil {
		params = &CancelQuantumTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelQuantumTask", params, optFns, addOperationCancelQuantumTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelQuantumTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelQuantumTaskInput struct {

	// The client token associated with the request.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the task to cancel.
	//
	// This member is required.
	QuantumTaskArn *string
}

type CancelQuantumTaskOutput struct {

	// The status of the cancellation request.
	//
	// This member is required.
	CancellationStatus types.CancellationStatus

	// The ARN of the task.
	//
	// This member is required.
	QuantumTaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelQuantumTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelQuantumTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelQuantumTask{}, middleware.After)
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
	if err = addOpCancelQuantumTaskValidationMiddleware(stack); err != nil {
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
