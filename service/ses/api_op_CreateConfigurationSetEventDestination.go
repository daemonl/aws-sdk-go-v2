// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a configuration set event destination. When you create or update an
// event destination, you must provide one, and only one, destination. The
// destination can be CloudWatch, Amazon Kinesis Firehose, or Amazon Simple
// Notification Service (Amazon SNS). An event destination is the AWS service to
// which Amazon SES publishes the email sending events associated with a
// configuration set. For information about using configuration sets, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
// You can execute this operation no more than once per second.
func (c *Client) CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &CreateConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationSetEventDestination", params, optFns, addOperationCreateConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a configuration set event destination. A
// configuration set event destination, which can be either Amazon CloudWatch or
// Amazon Kinesis Firehose, describes an AWS service in which Amazon SES publishes
// the email sending events associated with a configuration set. For information
// about using configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type CreateConfigurationSetEventDestinationInput struct {

	// The name of the configuration set that the event destination should be
	// associated with.
	//
	// This member is required.
	ConfigurationSetName *string

	// An object that describes the AWS service that email sending event information
	// will be published to.
	//
	// This member is required.
	EventDestination *types.EventDestination
}

// An empty element returned on a successful request.
type CreateConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateConfigurationSetEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateConfigurationSetEventDestination{}, middleware.After)
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
	if err = addOpCreateConfigurationSetEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationSetEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationSetEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateConfigurationSetEventDestination",
	}
}
