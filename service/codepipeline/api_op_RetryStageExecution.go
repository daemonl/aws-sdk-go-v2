// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resumes the pipeline execution by retrying the last failed actions in a stage.
// You can retry a stage immediately if any of the actions in the stage fail. When
// you retry, all actions that are still in progress continue working, and failed
// actions are triggered again.
func (c *Client) RetryStageExecution(ctx context.Context, params *RetryStageExecutionInput, optFns ...func(*Options)) (*RetryStageExecutionOutput, error) {
	if params == nil {
		params = &RetryStageExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetryStageExecution", params, optFns, addOperationRetryStageExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetryStageExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RetryStageExecution action.
type RetryStageExecutionInput struct {

	// The ID of the pipeline execution in the failed stage to be retried. Use the
	// GetPipelineState action to retrieve the current pipelineExecutionId of the
	// failed stage
	//
	// This member is required.
	PipelineExecutionId *string

	// The name of the pipeline that contains the failed stage.
	//
	// This member is required.
	PipelineName *string

	// The scope of the retry attempt. Currently, the only supported value is
	// FAILED_ACTIONS.
	//
	// This member is required.
	RetryMode types.StageRetryMode

	// The name of the failed stage to be retried.
	//
	// This member is required.
	StageName *string
}

// Represents the output of a RetryStageExecution action.
type RetryStageExecutionOutput struct {

	// The ID of the current workflow execution in the failed stage.
	PipelineExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRetryStageExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetryStageExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetryStageExecution{}, middleware.After)
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
	if err = addOpRetryStageExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetryStageExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetryStageExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "RetryStageExecution",
	}
}
