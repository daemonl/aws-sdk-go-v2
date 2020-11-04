// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Notifies Migration Hub of the current status, progress, or other detail
// regarding a migration task. This API has the following traits:
//
// * Migration
// tools will call the NotifyMigrationTaskState API to share the latest progress
// and status.
//
// * MigrationTaskName is used for addressing updates to the correct
// target.
//
// * ProgressUpdateStream is used for access control and to provide a
// namespace for each migration tool.
func (c *Client) NotifyMigrationTaskState(ctx context.Context, params *NotifyMigrationTaskStateInput, optFns ...func(*Options)) (*NotifyMigrationTaskStateOutput, error) {
	if params == nil {
		params = &NotifyMigrationTaskStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyMigrationTaskState", params, optFns, addOperationNotifyMigrationTaskStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyMigrationTaskStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyMigrationTaskStateInput struct {

	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// Number of seconds after the UpdateDateTime within which the Migration Hub can
	// expect an update. If Migration Hub does not receive an update within the
	// specified interval, then the migration task will be considered stale.
	//
	// This member is required.
	NextUpdateSeconds *int32

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// Information about the task's progress and status.
	//
	// This member is required.
	Task *types.Task

	// The timestamp when the task was gathered.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type NotifyMigrationTaskStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationNotifyMigrationTaskStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyMigrationTaskState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyMigrationTaskState{}, middleware.After)
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
	if err = addOpNotifyMigrationTaskStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyMigrationTaskState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyMigrationTaskState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "NotifyMigrationTaskState",
	}
}
