// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the state of fast snapshot restores for your snapshots.
func (c *Client) DescribeFastSnapshotRestores(ctx context.Context, params *DescribeFastSnapshotRestoresInput, optFns ...func(*Options)) (*DescribeFastSnapshotRestoresOutput, error) {
	if params == nil {
		params = &DescribeFastSnapshotRestoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFastSnapshotRestores", params, optFns, addOperationDescribeFastSnapshotRestoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFastSnapshotRestoresInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters. The possible values are:
	//
	// * availability-zone: The Availability
	// Zone of the snapshot.
	//
	// * owner-id: The ID of the AWS account that enabled fast
	// snapshot restore on the snapshot.
	//
	// * snapshot-id: The ID of the snapshot.
	//
	// *
	// state: The state of fast snapshot restores for the snapshot (enabling |
	// optimizing | enabled | disabling | disabled).
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeFastSnapshotRestoresOutput struct {

	// Information about the state of fast snapshot restores.
	FastSnapshotRestores []*types.DescribeFastSnapshotRestoreSuccessItem

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFastSnapshotRestoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFastSnapshotRestores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFastSnapshotRestores{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFastSnapshotRestores",
	}
}
