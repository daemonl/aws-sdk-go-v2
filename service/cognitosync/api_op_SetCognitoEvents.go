// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the AWS Lambda function for a given event type for an identity pool. This
// request only updates the key/value pair specified. Other key/values pairs are
// not updated. To remove a key value pair, pass a empty value for the particular
// key. This API can only be called with developer credentials. You cannot call
// this API with the temporary user credentials provided by Cognito Identity.
func (c *Client) SetCognitoEvents(ctx context.Context, params *SetCognitoEventsInput, optFns ...func(*Options)) (*SetCognitoEventsOutput, error) {
	if params == nil {
		params = &SetCognitoEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetCognitoEvents", params, optFns, addOperationSetCognitoEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetCognitoEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to configure Cognito Events.
type SetCognitoEventsInput struct {

	// The events to configure
	//
	// This member is required.
	Events map[string]*string

	// The Cognito Identity Pool to use when configuring Cognito Events
	//
	// This member is required.
	IdentityPoolId *string
}

type SetCognitoEventsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetCognitoEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetCognitoEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetCognitoEvents{}, middleware.After)
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
	if err = addOpSetCognitoEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetCognitoEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetCognitoEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "SetCognitoEvents",
	}
}
