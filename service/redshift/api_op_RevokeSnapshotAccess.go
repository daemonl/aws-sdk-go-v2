// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the ability of the specified AWS customer account to restore the
// specified snapshot. If the account is currently restoring the snapshot, the
// restore will run to completion. For more information about working with
// snapshots, go to Amazon Redshift Snapshots
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) RevokeSnapshotAccess(ctx context.Context, params *RevokeSnapshotAccessInput, optFns ...func(*Options)) (*RevokeSnapshotAccessOutput, error) {
	if params == nil {
		params = &RevokeSnapshotAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeSnapshotAccess", params, optFns, addOperationRevokeSnapshotAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeSnapshotAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RevokeSnapshotAccessInput struct {

	// The identifier of the AWS customer account that can no longer restore the
	// specified snapshot.
	//
	// This member is required.
	AccountWithRestoreAccess *string

	// The identifier of the snapshot that the account can no longer access.
	//
	// This member is required.
	SnapshotIdentifier *string

	// The identifier of the cluster the snapshot was created from. This parameter is
	// required if your IAM user has a policy containing a snapshot resource element
	// that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string
}

type RevokeSnapshotAccessOutput struct {

	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRevokeSnapshotAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRevokeSnapshotAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeSnapshotAccess{}, middleware.After)
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
	if err = addOpRevokeSnapshotAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSnapshotAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeSnapshotAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RevokeSnapshotAccess",
	}
}
