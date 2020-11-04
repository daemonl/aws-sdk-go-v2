// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata associated with a restore job that is specified by a job ID.
func (c *Client) DescribeRestoreJob(ctx context.Context, params *DescribeRestoreJobInput, optFns ...func(*Options)) (*DescribeRestoreJobOutput, error) {
	if params == nil {
		params = &DescribeRestoreJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRestoreJob", params, optFns, addOperationDescribeRestoreJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRestoreJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRestoreJobInput struct {

	// Uniquely identifies the job that restores a recovery point.
	//
	// This member is required.
	RestoreJobId *string
}

type DescribeRestoreJobOutput struct {

	// Returns the account ID that owns the restore job.
	AccountId *string

	// The size, in bytes, of the restored resource.
	BackupSizeInBytes *int64

	// The date and time that a job to restore a recovery point is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CompletionDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a resource whose recovery
	// point is being restored. The format of the ARN depends on the resource type of
	// the backed-up resource.
	CreatedResourceArn *string

	// The date and time that a restore job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The amount of time in minutes that a job restoring a recovery point is expected
	// to take.
	ExpectedCompletionTimeMinutes *int64

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// Contains an estimated percentage that is complete of a job at the time the job
	// status was queried.
	PercentDone *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// Returns metadata associated with a restore job listed by resource type.
	ResourceType *string

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string

	// Status code specifying the state of the job that is initiated by AWS Backup to
	// restore a recovery point.
	Status types.RestoreJobStatus

	// A message showing the status of a job to restore a recovery point.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRestoreJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRestoreJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRestoreJob{}, middleware.After)
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
	if err = addOpDescribeRestoreJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRestoreJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRestoreJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeRestoreJob",
	}
}
