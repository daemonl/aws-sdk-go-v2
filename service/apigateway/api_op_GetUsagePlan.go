// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a usage plan of a given plan identifier.
func (c *Client) GetUsagePlan(ctx context.Context, params *GetUsagePlanInput, optFns ...func(*Options)) (*GetUsagePlanOutput, error) {
	if params == nil {
		params = &GetUsagePlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlan", params, optFns, addOperationGetUsagePlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get a usage plan of a given plan identifier.
type GetUsagePlanInput struct {

	// [Required] The identifier of the UsagePlan resource to be retrieved.
	//
	// This member is required.
	UsagePlanId *string
}

// Represents a usage plan than can specify who can assess associated API stages
// with specified request limits and quotas. In a usage plan, you associate an API
// by specifying the API's Id and a stage name of the specified API. You add plan
// customers by adding API keys to the plan. Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanOutput struct {

	// The associated API stages of a usage plan.
	ApiStages []*types.ApiStage

	// The description of a usage plan.
	Description *string

	// The identifier of a UsagePlan resource.
	Id *string

	// The name of a usage plan.
	Name *string

	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS
	// product on AWS Marketplace.
	ProductCode *string

	// The maximum number of permitted requests per a given unit time interval.
	Quota *types.QuotaSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The request throttle limits of a usage plan.
	Throttle *types.ThrottleSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsagePlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlan{}, middleware.After)
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
	if err = addOpGetUsagePlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlan(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetUsagePlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlan",
	}
}
