// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of JobListEntry objects of the specified length. Each
// JobListEntry object is for a job in the specified cluster and contains a job's
// state, a job's ID, and other information.
func (c *Client) ListClusterJobs(ctx context.Context, params *ListClusterJobsInput, optFns ...func(*Options)) (*ListClusterJobsOutput, error) {
	if params == nil {
		params = &ListClusterJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterJobs", params, optFns, addOperationListClusterJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterJobsInput struct {

	// The 39-character ID for the cluster that you want to list, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	ClusterId *string

	// The number of JobListEntry objects to return.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// JobListEntry objects, you have the option of specifying NextToken as the
	// starting point for your returned list.
	NextToken *string
}

type ListClusterJobsOutput struct {

	// Each JobListEntry object contains a job's state, a job's ID, and a value that
	// indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []*types.JobListEntry

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next ListClusterJobsResult call, your list of returned jobs will
	// start from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListClusterJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusterJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusterJobs{}, middleware.After)
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
	if err = addOpListClusterJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListClusterJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListClusterJobs",
	}
}
