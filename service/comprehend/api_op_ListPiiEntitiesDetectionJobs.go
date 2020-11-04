// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the PII entity detection jobs that you have submitted.
func (c *Client) ListPiiEntitiesDetectionJobs(ctx context.Context, params *ListPiiEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListPiiEntitiesDetectionJobsOutput, error) {
	if params == nil {
		params = &ListPiiEntitiesDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPiiEntitiesDetectionJobs", params, optFns, addOperationListPiiEntitiesDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPiiEntitiesDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPiiEntitiesDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.PiiEntitiesDetectionJobFilter

	// The maximum number of results to return in each page.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListPiiEntitiesDetectionJobsOutput struct {

	// Identifies the next page of results to return.
	NextToken *string

	// A list containing the properties of each job that is returned.
	PiiEntitiesDetectionJobPropertiesList []*types.PiiEntitiesDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPiiEntitiesDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPiiEntitiesDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPiiEntitiesDetectionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPiiEntitiesDetectionJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPiiEntitiesDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListPiiEntitiesDetectionJobs",
	}
}
