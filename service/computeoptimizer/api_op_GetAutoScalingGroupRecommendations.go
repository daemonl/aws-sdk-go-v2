// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns Auto Scaling group recommendations. AWS Compute Optimizer generates
// recommendations for Amazon EC2 Auto Scaling groups that meet a specific set of
// requirements. For more information, see the Supported resources and requirements
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html) in
// the AWS Compute Optimizer User Guide.
func (c *Client) GetAutoScalingGroupRecommendations(ctx context.Context, params *GetAutoScalingGroupRecommendationsInput, optFns ...func(*Options)) (*GetAutoScalingGroupRecommendationsOutput, error) {
	if params == nil {
		params = &GetAutoScalingGroupRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAutoScalingGroupRecommendations", params, optFns, addOperationGetAutoScalingGroupRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAutoScalingGroupRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAutoScalingGroupRecommendationsInput struct {

	// The IDs of the AWS accounts for which to return Auto Scaling group
	// recommendations. If your account is the master account of an organization, use
	// this parameter to specify the member accounts for which you want to return Auto
	// Scaling group recommendations. Only one account ID can be specified per request.
	AccountIds []*string

	// The Amazon Resource Name (ARN) of the Auto Scaling groups for which to return
	// recommendations.
	AutoScalingGroupArns []*string

	// An array of objects that describe a filter that returns a more specific list of
	// Auto Scaling group recommendations.
	Filters []*types.Filter

	// The maximum number of Auto Scaling group recommendations to return with a single
	// request. To retrieve the remaining results, make another request with the
	// returned NextToken value.
	MaxResults *int32

	// The token to advance to the next page of Auto Scaling group recommendations.
	NextToken *string
}

type GetAutoScalingGroupRecommendationsOutput struct {

	// An array of objects that describe Auto Scaling group recommendations.
	AutoScalingGroupRecommendations []*types.AutoScalingGroupRecommendation

	// An array of objects that describe errors of the request. For example, an error
	// is returned if you request recommendations for an unsupported Auto Scaling
	// group.
	Errors []*types.GetRecommendationError

	// The token to use to advance to the next page of Auto Scaling group
	// recommendations. This value is null when there are no more pages of Auto Scaling
	// group recommendations to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAutoScalingGroupRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetAutoScalingGroupRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetAutoScalingGroupRecommendations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAutoScalingGroupRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAutoScalingGroupRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetAutoScalingGroupRecommendations",
	}
}
