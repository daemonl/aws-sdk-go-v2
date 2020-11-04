// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets information about an audit mitigation task that is used to apply mitigation
// actions to a set of audit findings. Properties include the actions being
// applied, the audit checks to which they're being applied, the task status, and
// aggregated task statistics.
func (c *Client) DescribeAuditMitigationActionsTask(ctx context.Context, params *DescribeAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*DescribeAuditMitigationActionsTaskOutput, error) {
	if params == nil {
		params = &DescribeAuditMitigationActionsTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAuditMitigationActionsTask", params, optFns, addOperationDescribeAuditMitigationActionsTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAuditMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAuditMitigationActionsTaskInput struct {

	// The unique identifier for the audit mitigation task.
	//
	// This member is required.
	TaskId *string
}

type DescribeAuditMitigationActionsTaskOutput struct {

	// Specifies the mitigation actions and their parameters that are applied as part
	// of this task.
	ActionsDefinition []*types.MitigationAction

	// Specifies the mitigation actions that should be applied to specific audit
	// checks.
	AuditCheckToActionsMapping map[string][]*string

	// The date and time when the task was completed or canceled.
	EndTime *time.Time

	// The date and time when the task was started.
	StartTime *time.Time

	// Identifies the findings to which the mitigation actions are applied. This can be
	// by audit checks, by audit task, or a set of findings.
	Target *types.AuditMitigationActionsTaskTarget

	// Aggregate counts of the results when the mitigation tasks were applied to the
	// findings for this audit mitigation actions task.
	TaskStatistics map[string]*types.TaskStatisticsForAuditCheck

	// The current status of the task.
	TaskStatus types.AuditMitigationActionsTaskStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAuditMitigationActionsTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAuditMitigationActionsTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAuditMitigationActionsTask{}, middleware.After)
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
	if err = addOpDescribeAuditMitigationActionsTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAuditMitigationActionsTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAuditMitigationActionsTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeAuditMitigationActionsTask",
	}
}
