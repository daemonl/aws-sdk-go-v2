// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attempts to stop an ongoing deployment.
func (c *Client) StopDeployment(ctx context.Context, params *StopDeploymentInput, optFns ...func(*Options)) (*StopDeploymentOutput, error) {
	if params == nil {
		params = &StopDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDeployment", params, optFns, addOperationStopDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a StopDeployment operation.
type StopDeploymentInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// Indicates, when a deployment is stopped, whether instances that have been
	// updated should be rolled back to the previous version of the application
	// revision.
	AutoRollbackEnabled *bool
}

// Represents the output of a StopDeployment operation.
type StopDeploymentOutput struct {

	// The status of the stop deployment operation:
	//
	// * Pending: The stop operation is
	// pending.
	//
	// * Succeeded: The stop operation was successful.
	Status types.StopStatus

	// An accompanying status message.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDeployment{}, middleware.After)
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
	if err = addOpStopDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "StopDeployment",
	}
}
