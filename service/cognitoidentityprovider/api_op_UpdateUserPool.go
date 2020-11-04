// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified user pool with the specified attributes. You can get a
// list of the current user pool settings using DescribeUserPool
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html).
// If you don't provide a value for an attribute, it will be set to the default
// value.
func (c *Client) UpdateUserPool(ctx context.Context, params *UpdateUserPoolInput, optFns ...func(*Options)) (*UpdateUserPoolOutput, error) {
	if params == nil {
		params = &UpdateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserPool", params, optFns, addOperationUpdateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user pool.
type UpdateUserPoolInput struct {

	// The user pool ID for the user pool you want to update.
	//
	// This member is required.
	UserPoolId *string

	// Use this setting to define which verified available method a user can use to
	// recover their password when they call ForgotPassword. It allows you to define a
	// preferred method when a user has more than one method available. With this
	// setting, SMS does not qualify for a valid password recovery mechanism if the
	// user also has SMS MFA enabled. In the absence of this setting, Cognito uses the
	// legacy behavior to determine the recovery method where SMS is preferred over
	// email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// The attributes that are automatically verified when the Amazon Cognito service
	// makes a request to update user pools.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// Device configuration.
	DeviceConfiguration *types.DeviceConfigurationType

	// Email configuration.
	EmailConfiguration *types.EmailConfigurationType

	// The contents of the email verification message.
	EmailVerificationMessage *string

	// The subject of the email verification message.
	EmailVerificationSubject *string

	// The AWS Lambda configuration information from the request to update the user
	// pool.
	LambdaConfig *types.LambdaConfigType

	// Can be one of the following values:
	//
	// * OFF - MFA tokens are not required and
	// cannot be specified during user registration.
	//
	// * ON - MFA tokens are required
	// for all user registrations. You can only specify required when you are initially
	// creating a user pool.
	//
	// * OPTIONAL - Users have the option when registering to
	// create an MFA token.
	MfaConfiguration types.UserPoolMfaType

	// A container with the policies you wish to update in a user pool.
	Policies *types.UserPoolPolicyType

	// The contents of the SMS authentication message.
	SmsAuthenticationMessage *string

	// SMS configuration.
	SmsConfiguration *types.SmsConfigurationType

	// A container with information about the SMS verification message.
	SmsVerificationMessage *string

	// Used to enable advanced security risk detection. Set the key
	// AdvancedSecurityMode to the value "AUDIT".
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]*string

	// The template for verification messages.
	VerificationMessageTemplate *types.VerificationMessageTemplateType
}

// Represents the response from the server when you make a request to update the
// user pool.
type UpdateUserPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPool{}, middleware.After)
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
	if err = addOpUpdateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateUserPool",
	}
}
