// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Syncrhonizes robots in a fleet to the latest deployment. This is helpful if
// robots were added after a deployment.
func (c *Client) SyncDeploymentJob(ctx context.Context, params *SyncDeploymentJobInput, optFns ...func(*Options)) (*SyncDeploymentJobOutput, error) {
	if params == nil {
		params = &SyncDeploymentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SyncDeploymentJob", params, optFns, addOperationSyncDeploymentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SyncDeploymentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SyncDeploymentJobInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientRequestToken *string

	// The target fleet for the synchronization.
	//
	// This member is required.
	Fleet *string
}

type SyncDeploymentJobOutput struct {

	// The Amazon Resource Name (ARN) of the synchronization request.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// Information about the deployment application configurations.
	DeploymentApplicationConfigs []*types.DeploymentApplicationConfig

	// Information about the deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// The failure code if the job fails: InternalServiceError Internal service error.
	// RobotApplicationCrash Robot application exited abnormally.
	// SimulationApplicationCrash Simulation application exited abnormally.
	// BadPermissionsRobotApplication Robot application bundle could not be downloaded.
	// BadPermissionsSimulationApplication Simulation application bundle could not be
	// downloaded. BadPermissionsS3Output Unable to publish outputs to
	// customer-provided S3 bucket. BadPermissionsCloudwatchLogs Unable to publish logs
	// to customer-provided CloudWatch Logs resource. SubnetIpLimitExceeded Subnet IP
	// limit exceeded. ENILimitExceeded ENI limit exceeded.
	// BadPermissionsUserCredentials Unable to use the Role provided.
	// InvalidBundleRobotApplication Robot bundle cannot be extracted (invalid format,
	// bundling error, or other issue). InvalidBundleSimulationApplication Simulation
	// bundle cannot be extracted (invalid format, bundling error, or other issue).
	// RobotApplicationVersionMismatchedEtag Etag for RobotApplication does not match
	// value during version creation. SimulationApplicationVersionMismatchedEtag Etag
	// for SimulationApplication does not match value during version creation.
	FailureCode types.DeploymentJobErrorCode

	// The failure reason if the job fails.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string

	// The status of the synchronization job.
	Status types.DeploymentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSyncDeploymentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSyncDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSyncDeploymentJob{}, middleware.After)
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
	if err = addIdempotencyToken_opSyncDeploymentJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSyncDeploymentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSyncDeploymentJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSyncDeploymentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSyncDeploymentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSyncDeploymentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SyncDeploymentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SyncDeploymentJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSyncDeploymentJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSyncDeploymentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSyncDeploymentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "SyncDeploymentJob",
	}
}
