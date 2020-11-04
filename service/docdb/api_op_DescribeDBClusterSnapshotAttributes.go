// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of cluster snapshot attribute names and values for a manual DB
// cluster snapshot. When you share snapshots with other AWS accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the AWS accounts that are authorized to copy or restore the manual
// cluster snapshot. If all is included in the list of values for the restore
// attribute, then the manual cluster snapshot is public and can be copied or
// restored by all AWS accounts.
func (c *Client) DescribeDBClusterSnapshotAttributes(ctx context.Context, params *DescribeDBClusterSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotAttributesOutput, error) {
	if params == nil {
		params = &DescribeDBClusterSnapshotAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterSnapshotAttributes", params, optFns, addOperationDescribeDBClusterSnapshotAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterSnapshotAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterSnapshotAttributes.
type DescribeDBClusterSnapshotAttributesInput struct {

	// The identifier for the cluster snapshot to describe the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string
}

type DescribeDBClusterSnapshotAttributesOutput struct {

	// Detailed information about the attributes that are associated with a cluster
	// snapshot.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterSnapshotAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
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
	if err = addOpDescribeDBClusterSnapshotAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterSnapshotAttributes",
	}
}
