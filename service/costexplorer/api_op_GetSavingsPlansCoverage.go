// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the Savings Plans covered for your account. This enables you to see
// how much of your cost is covered by a Savings Plan. An organization’s master
// account can see the coverage of the associated member accounts. This supports
// dimensions, Cost Categories, and nested expressions. For any time period, you
// can filter data for Savings Plans usage with the following dimensions:
//
// *
// LINKED_ACCOUNT
//
// * REGION
//
// * SERVICE
//
// * INSTANCE_FAMILY
//
// To determine valid
// values for a dimension, use the GetDimensionValues operation.
func (c *Client) GetSavingsPlansCoverage(ctx context.Context, params *GetSavingsPlansCoverageInput, optFns ...func(*Options)) (*GetSavingsPlansCoverageOutput, error) {
	if params == nil {
		params = &GetSavingsPlansCoverageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansCoverage", params, optFns, addOperationGetSavingsPlansCoverageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansCoverageInput struct {

	// The time period that you want the usage and costs for. The Start date must be
	// within 13 months. The End date must be after the Start date, and before the
	// current date. Future dates can't be used as an End date.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters Savings Plans coverage data by dimensions. You can filter data for
	// Savings Plans usage with the following dimensions:
	//
	// * LINKED_ACCOUNT
	//
	// *
	// REGION
	//
	// * SERVICE
	//
	// * INSTANCE_FAMILY
	//
	// GetSavingsPlansCoverage uses the same
	// Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	// If there are multiple values for a dimension, they are OR'd together. Cost
	// category is also supported.
	Filter *types.Expression

	// The granularity of the Amazon Web Services cost data for your Savings Plans.
	// Granularity can't be set if GroupBy is set. The GetSavingsPlansCoverage
	// operation supports only DAILY and MONTHLY granularities.
	Granularity types.Granularity

	// You can group the data using the attributes INSTANCE_FAMILY, REGION, or SERVICE.
	GroupBy []*types.GroupDefinition

	// The number of items to be returned in a response. The default is 20, with a
	// minimum value of 1.
	MaxResults *int32

	// The measurement that you want your Savings Plans coverage reported in. The only
	// valid value is SpendCoveredBySavingsPlans.
	Metrics []*string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string
}

type GetSavingsPlansCoverageOutput struct {

	// The amount of spend that your Savings Plans covered.
	//
	// This member is required.
	SavingsPlansCoverages []*types.SavingsPlansCoverage

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSavingsPlansCoverageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansCoverage{}, middleware.After)
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
	if err = addOpGetSavingsPlansCoverageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansCoverage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSavingsPlansCoverage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansCoverage",
	}
}
