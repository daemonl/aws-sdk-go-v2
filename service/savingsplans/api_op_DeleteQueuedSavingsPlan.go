// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the queued purchase for the specified Savings Plan.
func (c *Client) DeleteQueuedSavingsPlan(ctx context.Context, params *DeleteQueuedSavingsPlanInput, optFns ...func(*Options)) (*DeleteQueuedSavingsPlanOutput, error) {
	if params == nil {
		params = &DeleteQueuedSavingsPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQueuedSavingsPlan", params, optFns, addOperationDeleteQueuedSavingsPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQueuedSavingsPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueuedSavingsPlanInput struct {

	// The ID of the Savings Plan.
	//
	// This member is required.
	SavingsPlanId *string
}

type DeleteQueuedSavingsPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteQueuedSavingsPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteQueuedSavingsPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteQueuedSavingsPlan{}, middleware.After)
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
	if err = addOpDeleteQueuedSavingsPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueuedSavingsPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteQueuedSavingsPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "savingsplans",
		OperationName: "DeleteQueuedSavingsPlan",
	}
}
