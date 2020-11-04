// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details of the service quota increase request in your template.
func (c *Client) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *GetServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	if params == nil {
		params = &GetServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceQuotaIncreaseRequestFromTemplate", params, optFns, addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceQuotaIncreaseRequestFromTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceQuotaIncreaseRequestFromTemplateInput struct {

	// Specifies the AWS Region for the quota that you want to use.
	//
	// This member is required.
	AwsRegion *string

	// Specifies the quota you want.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string
}

type GetServiceQuotaIncreaseRequestFromTemplateOutput struct {

	// This object contains the details about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
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
	if err = addOpGetServiceQuotaIncreaseRequestFromTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetServiceQuotaIncreaseRequestFromTemplate",
	}
}
