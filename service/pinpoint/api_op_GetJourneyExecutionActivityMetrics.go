// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves (queries) pre-aggregated data for a standard execution metric that
// applies to a journey activity.
func (c *Client) GetJourneyExecutionActivityMetrics(ctx context.Context, params *GetJourneyExecutionActivityMetricsInput, optFns ...func(*Options)) (*GetJourneyExecutionActivityMetricsOutput, error) {
	if params == nil {
		params = &GetJourneyExecutionActivityMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJourneyExecutionActivityMetrics", params, optFns, addOperationGetJourneyExecutionActivityMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJourneyExecutionActivityMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJourneyExecutionActivityMetricsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the journey activity.
	//
	// This member is required.
	JourneyActivityId *string

	// The unique identifier for the journey.
	//
	// This member is required.
	JourneyId *string

	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string
}

type GetJourneyExecutionActivityMetricsOutput struct {

	// Provides the results of a query that retrieved the data for a standard execution
	// metric that applies to a journey activity, and provides information about that
	// query.
	//
	// This member is required.
	JourneyExecutionActivityMetricsResponse *types.JourneyExecutionActivityMetricsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetJourneyExecutionActivityMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetJourneyExecutionActivityMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJourneyExecutionActivityMetrics{}, middleware.After)
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
	if err = addOpGetJourneyExecutionActivityMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJourneyExecutionActivityMetrics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetJourneyExecutionActivityMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetJourneyExecutionActivityMetrics",
	}
}
