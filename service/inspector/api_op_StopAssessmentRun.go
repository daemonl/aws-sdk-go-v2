// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops the assessment run that is specified by the ARN of the assessment run.
func (c *Client) StopAssessmentRun(ctx context.Context, params *StopAssessmentRunInput, optFns ...func(*Options)) (*StopAssessmentRunOutput, error) {
	if params == nil {
		params = &StopAssessmentRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopAssessmentRun", params, optFns, addOperationStopAssessmentRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopAssessmentRunInput struct {

	// The ARN of the assessment run that you want to stop.
	//
	// This member is required.
	AssessmentRunArn *string

	// An input option that can be set to either START_EVALUATION or SKIP_EVALUATION.
	// START_EVALUATION (the default value), stops the AWS agent from collecting data
	// and begins the results evaluation and the findings generation process.
	// SKIP_EVALUATION cancels the assessment run immediately, after which no findings
	// are generated.
	StopAction types.StopAction
}

type StopAssessmentRunOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopAssessmentRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopAssessmentRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopAssessmentRun{}, middleware.After)
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
	if err = addOpStopAssessmentRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopAssessmentRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopAssessmentRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "StopAssessmentRun",
	}
}
