// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
func (c *Client) DescribeAccountAuditConfiguration(ctx context.Context, params *DescribeAccountAuditConfigurationInput, optFns ...func(*Options)) (*DescribeAccountAuditConfigurationOutput, error) {
	if params == nil {
		params = &DescribeAccountAuditConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAuditConfiguration", params, optFns, addOperationDescribeAccountAuditConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAuditConfigurationInput struct {
}

type DescribeAccountAuditConfigurationOutput struct {

	// Which audit checks are enabled and disabled for this account.
	AuditCheckConfigurations map[string]*types.AuditCheckConfiguration

	// Information about the targets to which audit notifications are sent for this
	// account.
	AuditNotificationTargetConfigurations map[string]*types.AuditNotificationTarget

	// The ARN of the role that grants permission to AWS IoT to access information
	// about your devices, policies, certificates, and other items as required when
	// performing an audit. On the first call to UpdateAccountAuditConfiguration, this
	// parameter is required.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountAuditConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountAuditConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountAuditConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeAccountAuditConfiguration",
	}
}
