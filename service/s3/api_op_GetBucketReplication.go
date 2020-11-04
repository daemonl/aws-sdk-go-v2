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

// Returns the replication configuration of a bucket. It can take a while to
// propagate the put or delete a replication configuration to all Amazon S3
// systems. Therefore, a get request soon after put or delete can return a wrong
// result. For information about replication configuration, see Replication
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the Amazon
// Simple Storage Service Developer Guide. This operation requires permissions for
// the s3:GetReplicationConfiguration action. For more information about
// permissions, see Using Bucket Policies and User Policies
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). If
// you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication and Priority elements. The response also
// returns those elements. For information about GetBucketReplication errors, see
// List of replication-related error codes
// (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList)
// The following operations are related to GetBucketReplication:
//
// *
// PutBucketReplication
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html)
//
// *
// DeleteBucketReplication
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html)
func (c *Client) GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) {
	if params == nil {
		params = &GetBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketReplication", params, optFns, addOperationGetBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketReplicationInput struct {

	// The bucket name for which to get the replication information.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketReplicationOutput struct {

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	ReplicationConfiguration *types.ReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketReplication{}, middleware.After)
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
	if err = addOpGetBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketReplication",
	}
}
