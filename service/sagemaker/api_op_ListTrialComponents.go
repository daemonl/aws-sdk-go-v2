// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the trial components in your account. You can sort the list by trial
// component name or creation time. You can filter the list to show only components
// that were created in a specific time range. You can also filter on one of the
// following:
//
// * ExperimentName
//
// * SourceArn
//
// * TrialName
func (c *Client) ListTrialComponents(ctx context.Context, params *ListTrialComponentsInput, optFns ...func(*Options)) (*ListTrialComponentsOutput, error) {
	if params == nil {
		params = &ListTrialComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrialComponents", params, optFns, addOperationListTrialComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrialComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrialComponentsInput struct {

	// A filter that returns only components created after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only components created before the specified time.
	CreatedBefore *time.Time

	// A filter that returns only components that are part of the specified experiment.
	// If you specify ExperimentName, you can't filter by SourceArn or TrialName.
	ExperimentName *string

	// The maximum number of components to return in the response. The default value is
	// 10.
	MaxResults *int32

	// If the previous call to ListTrialComponents didn't return the full set of
	// components, the call returns a token for getting the next set of components.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortTrialComponentsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	// A filter that returns only components that have the specified source Amazon
	// Resource Name (ARN). If you specify SourceArn, you can't filter by
	// ExperimentName or TrialName.
	SourceArn *string

	// A filter that returns only components that are part of the specified trial. If
	// you specify TrialName, you can't filter by ExperimentName or SourceArn.
	TrialName *string
}

type ListTrialComponentsOutput struct {

	// A token for getting the next set of components, if there are any.
	NextToken *string

	// A list of the summaries of your trial components.
	TrialComponentSummaries []*types.TrialComponentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrialComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrialComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrialComponents{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrialComponents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrialComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrialComponents",
	}
}
