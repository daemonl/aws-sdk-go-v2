// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an account setting. Account settings are set on a per-Region basis. If
// you change the account setting for the root user, the default settings for all
// of the IAM users and roles for which no individual account setting has been
// specified are reset. For more information, see Account Settings
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html)
// in the Amazon Elastic Container Service Developer Guide. When
// serviceLongArnFormat, taskLongArnFormat, or containerInstanceLongArnFormat are
// specified, the Amazon Resource Name (ARN) and resource ID format of the resource
// type for a specified IAM user, IAM role, or the root user for an account is
// affected. The opt-in and opt-out account setting must be set for each Amazon ECS
// resource separately. The ARN and resource ID format of a resource will be
// defined by the opt-in status of the IAM user or role that created the resource.
// You must enable this setting to use Amazon ECS features such as resource
// tagging. When awsvpcTrunking is specified, the elastic network interface (ENI)
// limit for any new container instances that support the feature is changed. If
// awsvpcTrunking is enabled, any new container instances that support the feature
// are launched have the increased ENI limits available to them. For more
// information, see Elastic Network Interface Trunking
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-eni.html)
// in the Amazon Elastic Container Service Developer Guide. When containerInsights
// is specified, the default setting indicating whether CloudWatch Container
// Insights is enabled for your clusters is changed. If containerInsights is
// enabled, any new clusters that are created will have Container Insights enabled
// unless you disable it during cluster creation. For more information, see
// CloudWatch Container Insights
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) PutAccountSetting(ctx context.Context, params *PutAccountSettingInput, optFns ...func(*Options)) (*PutAccountSettingOutput, error) {
	if params == nil {
		params = &PutAccountSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountSetting", params, optFns, addOperationPutAccountSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountSettingInput struct {

	// The Amazon ECS resource name for which to modify the account setting. If
	// serviceLongArnFormat is specified, the ARN for your Amazon ECS services is
	// affected. If taskLongArnFormat is specified, the ARN and resource ID for your
	// Amazon ECS tasks is affected. If containerInstanceLongArnFormat is specified,
	// the ARN and resource ID for your Amazon ECS container instances is affected. If
	// awsvpcTrunking is specified, the elastic network interface (ENI) limit for your
	// Amazon ECS container instances is affected. If containerInsights is specified,
	// the default setting for CloudWatch Container Insights for your clusters is
	// affected.
	//
	// This member is required.
	Name types.SettingName

	// The account setting value for the specified principal ARN. Accepted values are
	// enabled and disabled.
	//
	// This member is required.
	Value *string

	// The ARN of the principal, which can be an IAM user, IAM role, or the root user.
	// If you specify the root user, it modifies the account setting for all IAM users,
	// IAM roles, and the root user of the account unless an IAM user or role
	// explicitly overrides these settings. If this field is omitted, the setting is
	// changed only for the authenticated user.
	PrincipalArn *string
}

type PutAccountSettingOutput struct {

	// The current account setting for a resource.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAccountSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountSetting{}, middleware.After)
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
	if err = addOpPutAccountSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "PutAccountSetting",
	}
}
