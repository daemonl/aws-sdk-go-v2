// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Bucket lifecycle configuration now supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, or a combination of both.
// Accordingly, this section describes the latest API. The response describes the
// new filter element that you can use to specify a filter to select a subset of
// objects to which the rule applies. If you are using a previous version of the
// lifecycle configuration, it still works. For the earlier API description, see
// GetBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html).
// Returns the lifecycle configuration information set on the bucket. For
// information about lifecycle configuration, see Object Lifecycle Management
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html). To
// use this operation, you must have permission to perform the
// s3:GetLifecycleConfiguration action. The bucket owner has this permission, by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
// GetBucketLifecycleConfiguration has the following special error:
//
// * Error code:
// NoSuchLifecycleConfiguration
//
// * Description: The lifecycle configuration does
// not exist.
//
// * HTTP Status Code: 404 Not Found
//
// * SOAP Fault Code Prefix:
// Client
//
// The following operations are related to
// GetBucketLifecycleConfiguration:
//
// * GetBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html)
//
// *
// PutBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html)
//
// *
// DeleteBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketLifecycleConfiguration", params, optFns, addOperationGetBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketLifecycleConfigurationInput struct {

	// The name of the bucket for which to get the lifecycle information.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketLifecycleConfigurationOutput struct {

	// Container for a lifecycle rule.
	Rules []*types.LifecycleRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketLifecycleConfiguration{}, middleware.After)
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
	if err = addOpGetBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketLifecycleConfiguration",
	}
}
