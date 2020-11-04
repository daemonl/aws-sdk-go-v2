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

// Creates a replication configuration or replaces an existing one. For more
// information, see Replication
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the Amazon
// S3 Developer Guide. To perform this operation, the user or role performing the
// operation must have the iam:PassRole
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html)
// permission. Specify the replication configuration in the request body. In the
// replication configuration, you provide the name of the destination bucket where
// you want Amazon S3 to replicate objects, the IAM role that Amazon S3 can assume
// to replicate objects on your behalf, and other relevant information. A
// replication configuration must include at least one rule, and can contain a
// maximum of 1,000. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source bucket. To choose additional subsets of
// objects to replicate, add a rule for each subset. All rules must specify the
// same destination bucket. To specify a subset of the objects in the source bucket
// to apply a replication rule to, add the Filter element as a child of the Rule
// element. You can filter objects based on an object key prefix, one or more
// object tags, or both. When you add the Filter element in the configuration, you
// must also add the following elements: DeleteMarkerReplication, Status, and
// Priority. The latest version of the replication configuration XML is V2. XML V2
// replication configurations are those that contain the Filter element for rules,
// and rules that specify S3 Replication Time Control (S3 RTC). In XML V2
// replication configurations, Amazon S3 doesn't replicate delete markers.
// Therefore, you must set the DeleteMarkerReplication element to Disabled. For
// backward compatibility, Amazon S3 continues to support the XML V1 replication
// configuration. For information about enabling versioning on a bucket, see Using
// Versioning (https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html). By
// default, a resource owner, in this case the AWS account that created the bucket,
// can perform this operation. The resource owner can also grant others permissions
// to perform the operation. For more information about permissions, see Specifying
// Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html) and
// Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
// Handling Replication of Encrypted Objects By default, Amazon S3 doesn't
// replicate objects that are stored at rest using server-side encryption with CMKs
// stored in AWS KMS. To replicate AWS KMS-encrypted objects, add the following:
// SourceSelectionCriteria, SseKmsEncryptedObjects, Status,
// EncryptionConfiguration, and ReplicaKmsKeyID. For information about replication
// configuration, see Replicating Objects Created with SSE Using CMKs stored in AWS
// KMS
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-config-for-kms-objects.html).
// For information on PutBucketReplication errors, see List of replication-related
// error codes
// (https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList)
// The following operations are related to PutBucketReplication:
//
// *
// GetBucketReplication
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html)
//
// *
// DeleteBucketReplication
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html)
func (c *Client) PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error) {
	if params == nil {
		params = &PutBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketReplication", params, optFns, addOperationPutBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketReplicationInput struct {

	// The name of the bucket
	//
	// This member is required.
	Bucket *string

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	//
	// This member is required.
	ReplicationConfiguration *types.ReplicationConfiguration

	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt).
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	//
	Token *string
}

type PutBucketReplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketReplication{}, middleware.After)
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
	if err = addOpPutBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketReplication(options.Region), middleware.Before); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketReplication",
	}
}
