// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates variant weight of one or more variants associated with an existing
// endpoint, or capacity of one variant associated with an existing endpoint. When
// it receives the request, Amazon SageMaker sets the endpoint status to Updating.
// After updating the endpoint, it sets the status to InService. To check the
// status of an endpoint, use the DescribeEndpoint API.
func (c *Client) UpdateEndpointWeightsAndCapacities(ctx context.Context, params *UpdateEndpointWeightsAndCapacitiesInput, optFns ...func(*Options)) (*UpdateEndpointWeightsAndCapacitiesOutput, error) {
	if params == nil {
		params = &UpdateEndpointWeightsAndCapacitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpointWeightsAndCapacities", params, optFns, addOperationUpdateEndpointWeightsAndCapacitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointWeightsAndCapacitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointWeightsAndCapacitiesInput struct {

	// An object that provides new capacity and weight values for a variant.
	//
	// This member is required.
	DesiredWeightsAndCapacities []*types.DesiredWeightAndCapacity

	// The name of an existing Amazon SageMaker endpoint.
	//
	// This member is required.
	EndpointName *string
}

type UpdateEndpointWeightsAndCapacitiesOutput struct {

	// The Amazon Resource Name (ARN) of the updated endpoint.
	//
	// This member is required.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEndpointWeightsAndCapacitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpointWeightsAndCapacities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpointWeightsAndCapacities{}, middleware.After)
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
	if err = addOpUpdateEndpointWeightsAndCapacitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpointWeightsAndCapacities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpointWeightsAndCapacities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateEndpointWeightsAndCapacities",
	}
}
