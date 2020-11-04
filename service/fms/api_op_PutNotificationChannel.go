// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Designates the IAM role and Amazon Simple Notification Service (SNS) topic that
// AWS Firewall Manager uses to record SNS logs.
func (c *Client) PutNotificationChannel(ctx context.Context, params *PutNotificationChannelInput, optFns ...func(*Options)) (*PutNotificationChannelOutput, error) {
	if params == nil {
		params = &PutNotificationChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutNotificationChannel", params, optFns, addOperationPutNotificationChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutNotificationChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutNotificationChannelInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record
	// AWS Firewall Manager activity.
	//
	// This member is required.
	SnsRoleName *string

	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications from
	// AWS Firewall Manager.
	//
	// This member is required.
	SnsTopicArn *string
}

type PutNotificationChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutNotificationChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutNotificationChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutNotificationChannel{}, middleware.After)
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
	if err = addOpPutNotificationChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutNotificationChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutNotificationChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "PutNotificationChannel",
	}
}
