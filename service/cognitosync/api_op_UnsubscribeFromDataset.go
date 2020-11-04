// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unsubscribes from receiving notifications when a dataset is modified by another
// device. This API can only be called with temporary credentials provided by
// Cognito Identity. You cannot call this API with developer credentials.
func (c *Client) UnsubscribeFromDataset(ctx context.Context, params *UnsubscribeFromDatasetInput, optFns ...func(*Options)) (*UnsubscribeFromDatasetOutput, error) {
	if params == nil {
		params = &UnsubscribeFromDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnsubscribeFromDataset", params, optFns, addOperationUnsubscribeFromDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnsubscribeFromDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to UnsubscribeFromDataset.
type UnsubscribeFromDatasetInput struct {

	// The name of the dataset from which to unsubcribe.
	//
	// This member is required.
	DatasetName *string

	// The unique ID generated for this device by Cognito.
	//
	// This member is required.
	DeviceId *string

	// Unique ID for this identity.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. The ID of the pool to which this identity belongs.
	//
	// This member is required.
	IdentityPoolId *string
}

// Response to an UnsubscribeFromDataset request.
type UnsubscribeFromDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnsubscribeFromDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUnsubscribeFromDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUnsubscribeFromDataset{}, middleware.After)
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
	if err = addOpUnsubscribeFromDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribeFromDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnsubscribeFromDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "UnsubscribeFromDataset",
	}
}
