// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the endpoint attributes for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
// information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
func (c *Client) GetEndpointAttributes(ctx context.Context, params *GetEndpointAttributesInput, optFns ...func(*Options)) (*GetEndpointAttributesOutput, error) {
	if params == nil {
		params = &GetEndpointAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEndpointAttributes", params, optFns, addOperationGetEndpointAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEndpointAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetEndpointAttributes action.
type GetEndpointAttributesInput struct {

	// EndpointArn for GetEndpointAttributes input.
	//
	// This member is required.
	EndpointArn *string
}

// Response from GetEndpointAttributes of the EndpointArn.
type GetEndpointAttributesOutput struct {

	// Attributes include the following:
	//
	// * CustomUserData – arbitrary user data to
	// associate with the endpoint. Amazon SNS does not use this data. The data must be
	// in UTF-8 format and less than 2KB.
	//
	// * Enabled – flag that enables/disables
	// delivery to the endpoint. Amazon SNS will set this to false when a notification
	// service indicates to Amazon SNS that the endpoint is invalid. Users can set it
	// back to true, typically after updating Token.
	//
	// * Token – device token, also
	// referred to as a registration id, for an app and mobile device. This is returned
	// from the notification service when an app and mobile device are registered with
	// the notification service. The device token for the iOS platform is returned in
	// lowercase.
	Attributes map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEndpointAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetEndpointAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetEndpointAttributes{}, middleware.After)
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
	if err = addOpGetEndpointAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEndpointAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEndpointAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "GetEndpointAttributes",
	}
}
