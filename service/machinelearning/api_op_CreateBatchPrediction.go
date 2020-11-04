// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates predictions for a group of observations. The observations to process
// exist in one or more data files referenced by a DataSource. This operation
// creates a new BatchPrediction, and uses an MLModel and the data files referenced
// by the DataSource as information sources. CreateBatchPrediction is an
// asynchronous operation. In response to CreateBatchPrediction, Amazon Machine
// Learning (Amazon ML) immediately returns and sets the BatchPrediction status to
// PENDING. After the BatchPrediction completes, Amazon ML sets the status to
// COMPLETED. You can poll for status updates by using the GetBatchPrediction
// operation and checking the Status parameter of the result. After the COMPLETED
// status appears, the results are available in the location specified by the
// OutputUri parameter.
func (c *Client) CreateBatchPrediction(ctx context.Context, params *CreateBatchPredictionInput, optFns ...func(*Options)) (*CreateBatchPredictionOutput, error) {
	if params == nil {
		params = &CreateBatchPredictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchPrediction", params, optFns, addOperationCreateBatchPredictionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchPredictionInput struct {

	// The ID of the DataSource that points to the group of observations to predict.
	//
	// This member is required.
	BatchPredictionDataSourceId *string

	// A user-supplied ID that uniquely identifies the BatchPrediction.
	//
	// This member is required.
	BatchPredictionId *string

	// The ID of the MLModel that will generate predictions for the group of
	// observations.
	//
	// This member is required.
	MLModelId *string

	// The location of an Amazon Simple Storage Service (Amazon S3) bucket or directory
	// to store the batch prediction results. The following substrings are not allowed
	// in the s3 key portion of the outputURI field: ':', '//', '/./', '/../'. Amazon
	// ML needs permissions to store and retrieve the logs on your behalf. For
	// information about how to set permissions, see the Amazon Machine Learning
	// Developer Guide (https://docs.aws.amazon.com/machine-learning/latest/dg).
	//
	// This member is required.
	OutputUri *string

	// A user-supplied name or description of the BatchPrediction. BatchPredictionName
	// can only use the UTF-8 character set.
	BatchPredictionName *string
}

// Represents the output of a CreateBatchPrediction operation, and is an
// acknowledgement that Amazon ML received the request. The CreateBatchPrediction
// operation is asynchronous. You can poll for status updates by using the
// >GetBatchPrediction operation and checking the Status parameter of the result.
type CreateBatchPredictionOutput struct {

	// A user-supplied ID that uniquely identifies the BatchPrediction. This value is
	// identical to the value of the BatchPredictionId in the request.
	BatchPredictionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBatchPredictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchPrediction{}, middleware.After)
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
	if err = addOpCreateBatchPredictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchPrediction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchPrediction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateBatchPrediction",
	}
}
