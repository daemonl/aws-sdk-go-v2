// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the details of the list of email addresses and phone numbers that the
// DDoS Response Team (DRT) can use to contact you if you have proactive engagement
// enabled, for escalations to the DRT and to initiate proactive customer support.
func (c *Client) UpdateEmergencyContactSettings(ctx context.Context, params *UpdateEmergencyContactSettingsInput, optFns ...func(*Options)) (*UpdateEmergencyContactSettingsOutput, error) {
	if params == nil {
		params = &UpdateEmergencyContactSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEmergencyContactSettings", params, optFns, addOperationUpdateEmergencyContactSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEmergencyContactSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEmergencyContactSettingsInput struct {

	// A list of email addresses and phone numbers that the DDoS Response Team (DRT)
	// can use to contact you if you have proactive engagement enabled, for escalations
	// to the DRT and to initiate proactive customer support. If you have proactive
	// engagement enabled, the contact list must include at least one phone number.
	EmergencyContactList []*types.EmergencyContact
}

type UpdateEmergencyContactSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEmergencyContactSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEmergencyContactSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEmergencyContactSettings{}, middleware.After)
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
	if err = addOpUpdateEmergencyContactSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEmergencyContactSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEmergencyContactSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "UpdateEmergencyContactSettings",
	}
}
