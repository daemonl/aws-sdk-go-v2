// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the projected utilization metrics of Amazon EC2 instance
// recommendations. The Cpu and Memory metrics are the only projected utilization
// metrics returned when you run this action. Additionally, the Memory metric is
// returned only for resources that have the unified CloudWatch agent installed on
// them. For more information, see Enabling Memory Utilization with the CloudWatch
// Agent
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/metrics.html#cw-agent).
func (c *Client) GetEC2RecommendationProjectedMetrics(ctx context.Context, params *GetEC2RecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetEC2RecommendationProjectedMetricsOutput, error) {
	if params == nil {
		params = &GetEC2RecommendationProjectedMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEC2RecommendationProjectedMetrics", params, optFns, addOperationGetEC2RecommendationProjectedMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEC2RecommendationProjectedMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEC2RecommendationProjectedMetricsInput struct {

	// The time stamp of the last projected metrics data point to return.
	//
	// This member is required.
	EndTime *time.Time

	// The Amazon Resource Name (ARN) of the instances for which to return
	// recommendation projected metrics.
	//
	// This member is required.
	InstanceArn *string

	// The granularity, in seconds, of the projected metrics data points.
	//
	// This member is required.
	Period *int32

	// The time stamp of the first projected metrics data point to return.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic of the projected metrics.
	//
	// This member is required.
	Stat types.MetricStatistic
}

type GetEC2RecommendationProjectedMetricsOutput struct {

	// An array of objects that describe a projected metrics.
	RecommendedOptionProjectedMetrics []*types.RecommendedOptionProjectedMetric

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEC2RecommendationProjectedMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEC2RecommendationProjectedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEC2RecommendationProjectedMetrics{}, middleware.After)
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
	if err = addOpGetEC2RecommendationProjectedMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEC2RecommendationProjectedMetrics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEC2RecommendationProjectedMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetEC2RecommendationProjectedMetrics",
	}
}
