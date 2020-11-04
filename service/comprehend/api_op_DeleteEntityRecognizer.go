// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an entity recognizer. Only those recognizers that are in terminated
// states (IN_ERROR, TRAINED) will be deleted. If an active inference job is using
// the model, a ResourceInUseException will be returned. This is an asynchronous
// action that puts the recognizer into a DELETING state, and it is then removed by
// a background job. Once removed, the recognizer disappears from your account and
// is no longer available for use.
func (c *Client) DeleteEntityRecognizer(ctx context.Context, params *DeleteEntityRecognizerInput, optFns ...func(*Options)) (*DeleteEntityRecognizerOutput, error) {
	if params == nil {
		params = &DeleteEntityRecognizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEntityRecognizer", params, optFns, addOperationDeleteEntityRecognizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEntityRecognizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEntityRecognizerInput struct {

	// The Amazon Resource Name (ARN) that identifies the entity recognizer.
	//
	// This member is required.
	EntityRecognizerArn *string
}

type DeleteEntityRecognizerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEntityRecognizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEntityRecognizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEntityRecognizer{}, middleware.After)
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
	if err = addOpDeleteEntityRecognizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEntityRecognizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEntityRecognizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DeleteEntityRecognizer",
	}
}
