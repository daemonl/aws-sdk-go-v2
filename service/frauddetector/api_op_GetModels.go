// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets one or more models. Gets all models for the AWS account if no model type
// and no model id provided. Gets all models for the AWS account and model type, if
// the model type is specified but model id is not provided. Gets a specific model
// if (model type, model id) tuple is specified. This is a paginated API. If you
// provide a null maxResults, this action retrieves a maximum of 10 records per
// page. If you provide a maxResults, the value must be between 1 and 10. To get
// the next page results, provide the pagination token from the response as part of
// your request. A null pagination token fetches the records from the beginning.
func (c *Client) GetModels(ctx context.Context, params *GetModelsInput, optFns ...func(*Options)) (*GetModelsOutput, error) {
	if params == nil {
		params = &GetModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModels", params, optFns, addOperationGetModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelsInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The model ID.
	ModelId *string

	// The model type.
	ModelType types.ModelTypeEnum

	// The next token for the subsequent request.
	NextToken *string
}

type GetModelsOutput struct {

	// The array of models.
	Models []*types.Model

	// The next page token to be used in subsequent requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetModels",
	}
}
