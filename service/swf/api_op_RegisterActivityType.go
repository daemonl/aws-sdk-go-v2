// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a new activity type along with its configuration settings in the
// specified domain. A TypeAlreadyExists fault is returned if the type already
// exists in the domain. You cannot change any configuration settings of the type
// after its registration, and it must be registered as a new version. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// * Use a Resource element with the domain name to limit
// the action to only specified domains.
//
// * Use an Action element to allow or deny
// permission to call this action.
//
// * Constrain the following parameters by using a
// Condition element with the appropriate keys.
//
// * defaultTaskList.name: String
// constraint. The key is swf:defaultTaskList.name.
//
// * name: String constraint. The
// key is swf:name.
//
// * version: String constraint. The key is swf:version.
//
// If the
// caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RegisterActivityType(ctx context.Context, params *RegisterActivityTypeInput, optFns ...func(*Options)) (*RegisterActivityTypeOutput, error) {
	if params == nil {
		params = &RegisterActivityTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterActivityType", params, optFns, addOperationRegisterActivityTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterActivityTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterActivityTypeInput struct {

	// The name of the domain in which this activity is to be registered.
	//
	// This member is required.
	Domain *string

	// The name of the activity type within the domain. The specified string must not
	// start or end with whitespace. It must not contain a : (colon), / (slash), |
	// (vertical bar), or any control characters (\u0000-\u001f | \u007f-\u009f). Also,
	// it must not be the literal string arn.
	//
	// This member is required.
	Name *string

	// The version of the activity type. The activity type consists of the name and
	// version, the combination of which must be unique within the domain. The
	// specified string must not start or end with whitespace. It must not contain a :
	// (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f |
	// \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// This member is required.
	Version *string

	// If set, specifies the default maximum time before which a worker processing a
	// task of this type must report progress by calling RecordActivityTaskHeartbeat.
	// If the timeout is exceeded, the activity task is automatically timed out. This
	// default can be overridden when scheduling an activity task using the
	// ScheduleActivityTaskDecision. If the activity worker subsequently attempts to
	// record a heartbeat or returns a result, the activity worker receives an
	// UnknownResource fault. In this case, Amazon SWF no longer considers the activity
	// task to be valid; the activity worker should clean up the activity task. The
	// duration is specified in seconds, an integer greater than or equal to 0. You can
	// use NONE to specify unlimited duration.
	DefaultTaskHeartbeatTimeout *string

	// If set, specifies the default task list to use for scheduling tasks of this
	// activity type. This default task list is used if a task list isn't provided when
	// a task is scheduled through the ScheduleActivityTaskDecision.
	DefaultTaskList *types.TaskList

	// The default task priority to assign to the activity type. If not assigned, then
	// 0 is used. Valid values are integers that range from Java's Integer.MIN_VALUE
	// (-2147483648) to Integer.MAX_VALUE (2147483647). Higher numbers indicate higher
	// priority. For more information about setting task priority, see Setting Task
	// Priority
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the in the Amazon SWF Developer Guide..
	DefaultTaskPriority *string

	// If set, specifies the default maximum duration for a task of this activity type.
	// This default can be overridden when scheduling an activity task using the
	// ScheduleActivityTaskDecision. The duration is specified in seconds, an integer
	// greater than or equal to 0. You can use NONE to specify unlimited duration.
	DefaultTaskScheduleToCloseTimeout *string

	// If set, specifies the default maximum duration that a task of this activity type
	// can wait before being assigned to a worker. This default can be overridden when
	// scheduling an activity task using the ScheduleActivityTaskDecision. The duration
	// is specified in seconds, an integer greater than or equal to 0. You can use NONE
	// to specify unlimited duration.
	DefaultTaskScheduleToStartTimeout *string

	// If set, specifies the default maximum duration that a worker can take to process
	// tasks of this activity type. This default can be overridden when scheduling an
	// activity task using the ScheduleActivityTaskDecision. The duration is specified
	// in seconds, an integer greater than or equal to 0. You can use NONE to specify
	// unlimited duration.
	DefaultTaskStartToCloseTimeout *string

	// A textual description of the activity type.
	Description *string
}

type RegisterActivityTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterActivityTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRegisterActivityType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRegisterActivityType{}, middleware.After)
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
	if err = addOpRegisterActivityTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterActivityType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterActivityType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RegisterActivityType",
	}
}
