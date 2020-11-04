// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API operation gets an Amazon S3 on Outposts bucket's lifecycle
// configuration. To get an S3 bucket's lifecycle configuration, see
// GetBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html)
// in the Amazon Simple Storage Service API. Returns the lifecycle configuration
// information set on the Outposts bucket. For more information, see Using Amazon
// S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) and for
// information about lifecycle configuration, see  Object Lifecycle Management
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in
// Amazon Simple Storage Service Developer Guide. To use this operation, you must
// have permission to perform the s3outposts:GetLifecycleConfiguration action. The
// Outposts bucket owner has this permission, by default. The bucket owner can
// grant this permission to others. For more information about permissions, see
// Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). All
// Amazon S3 on Outposts REST API requests for this action require an additional
// parameter of outpost-id to be passed with the request and an S3 on Outposts
// endpoint hostname prefix instead of s3-control. For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucketLifecycleConfiguration.html#API_control_GetBucketLifecycleConfiguration_Examples)
// section below. GetBucketLifecycleConfiguration has the following special
// error:
//
// * Error code: NoSuchLifecycleConfiguration
//
// * Description: The lifecycle
// configuration does not exist.
//
// * HTTP Status Code: 404 Not Found
//
// * SOAP Fault
// Code Prefix: Client
//
// The following actions are related to
// GetBucketLifecycleConfiguration:
//
// * PutBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
//
// *
// DeleteBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html)
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

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of the bucket. For Amazon S3 on Outposts specify
	// the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type GetBucketLifecycleConfigurationOutput struct {

	// Container for the lifecycle rule of the Outposts bucket.
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
	if err = addEndpointPrefix_opGetBucketLifecycleConfigurationMiddleware(stack); err != nil {
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
	return nil
}

type endpointPrefix_opGetBucketLifecycleConfigurationMiddleware struct {
}

func (*endpointPrefix_opGetBucketLifecycleConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketLifecycleConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetBucketLifecycleConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetBucketLifecycleConfigurationMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketLifecycleConfiguration",
	}
}
