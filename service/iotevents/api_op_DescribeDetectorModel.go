// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a detector model. If the version parameter is not specified,
// information about the latest version is returned.
func (c *Client) DescribeDetectorModel(ctx context.Context, params *DescribeDetectorModelInput, optFns ...func(*Options)) (*DescribeDetectorModelOutput, error) {
	if params == nil {
		params = &DescribeDetectorModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDetectorModel", params, optFns, addOperationDescribeDetectorModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDetectorModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDetectorModelInput struct {

	// The name of the detector model.
	//
	// This member is required.
	DetectorModelName *string

	// The version of the detector model.
	DetectorModelVersion *string
}

type DescribeDetectorModelOutput struct {

	// Information about the detector model.
	DetectorModel *types.DetectorModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDetectorModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDetectorModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDetectorModel{}, middleware.After)
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
	if err = addOpDescribeDetectorModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDetectorModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDetectorModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "DescribeDetectorModel",
	}
}
