// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the current status of a mailbox export job.
func (c *Client) DescribeMailboxExportJob(ctx context.Context, params *DescribeMailboxExportJobInput, optFns ...func(*Options)) (*DescribeMailboxExportJobOutput, error) {
	if params == nil {
		params = &DescribeMailboxExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMailboxExportJob", params, optFns, addOperationDescribeMailboxExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMailboxExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMailboxExportJobInput struct {

	// The mailbox export job ID.
	//
	// This member is required.
	JobId *string

	// The organization ID.
	//
	// This member is required.
	OrganizationId *string
}

type DescribeMailboxExportJobOutput struct {

	// The mailbox export job description.
	Description *string

	// The mailbox export job end timestamp.
	EndTime *time.Time

	// The identifier of the user or resource associated with the mailbox.
	EntityId *string

	// Error information for failed mailbox export jobs.
	ErrorInfo *string

	// The estimated progress of the mailbox export job, in percentage points.
	EstimatedProgress *int32

	// The Amazon Resource Name (ARN) of the symmetric AWS Key Management Service (AWS
	// KMS) key that encrypts the exported mailbox content.
	KmsKeyArn *string

	// The ARN of the AWS Identity and Access Management (IAM) role that grants write
	// permission to the Amazon Simple Storage Service (Amazon S3) bucket.
	RoleArn *string

	// The name of the S3 bucket.
	S3BucketName *string

	// The path to the S3 bucket and file that the mailbox export job is exporting to.
	S3Path *string

	// The S3 bucket prefix.
	S3Prefix *string

	// The mailbox export job start timestamp.
	StartTime *time.Time

	// The state of the mailbox export job.
	State types.MailboxExportJobState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMailboxExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMailboxExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMailboxExportJob{}, middleware.After)
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
	if err = addOpDescribeMailboxExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMailboxExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMailboxExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DescribeMailboxExportJob",
	}
}
