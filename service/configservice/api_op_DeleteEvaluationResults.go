// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the evaluation results for the specified AWS Config rule. You can
// specify one AWS Config rule per request. After you delete the evaluation
// results, you can call the StartConfigRulesEvaluation API to start evaluating
// your AWS resources against the rule.
func (c *Client) DeleteEvaluationResults(ctx context.Context, params *DeleteEvaluationResultsInput, optFns ...func(*Options)) (*DeleteEvaluationResultsOutput, error) {
	if params == nil {
		params = &DeleteEvaluationResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEvaluationResults", params, optFns, addOperationDeleteEvaluationResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEvaluationResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteEvaluationResultsInput struct {

	// The name of the AWS Config rule for which you want to delete the evaluation
	// results.
	//
	// This member is required.
	ConfigRuleName *string
}

// The output when you delete the evaluation results for the specified AWS Config
// rule.
type DeleteEvaluationResultsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEvaluationResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEvaluationResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEvaluationResults{}, middleware.After)
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
	if err = addOpDeleteEvaluationResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEvaluationResults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEvaluationResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteEvaluationResults",
	}
}
