// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds, updates, or removes the description of an experiment. Updates the display
// name of an experiment.
func (c *Client) UpdateExperiment(ctx context.Context, params *UpdateExperimentInput, optFns ...func(*Options)) (*UpdateExperimentOutput, error) {
	if params == nil {
		params = &UpdateExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateExperiment", params, optFns, addOperationUpdateExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExperimentInput struct {

	// The name of the experiment to update.
	//
	// This member is required.
	ExperimentName *string

	// The description of the experiment.
	Description *string

	// The name of the experiment as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, ExperimentName is displayed.
	DisplayName *string
}

type UpdateExperimentOutput struct {

	// The Amazon Resource Name (ARN) of the experiment.
	ExperimentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateExperiment{}, middleware.After)
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
	if err = addOpUpdateExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExperiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateExperiment",
	}
}
