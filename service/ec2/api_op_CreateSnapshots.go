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

// Creates crash-consistent snapshots of multiple EBS volumes and stores the data
// in S3. Volumes are chosen by specifying an instance. Any attached volumes will
// produce one snapshot each that is crash-consistent across the instance. Boot
// volumes can be excluded by changing the parameters.
func (c *Client) CreateSnapshots(ctx context.Context, params *CreateSnapshotsInput, optFns ...func(*Options)) (*CreateSnapshotsOutput, error) {
	if params == nil {
		params = &CreateSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshots", params, optFns, addOperationCreateSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotsInput struct {

	// The instance to specify which volumes should be included in the snapshots.
	//
	// This member is required.
	InstanceSpecification *types.InstanceSpecification

	// Copies the tags from the specified volume to corresponding snapshot.
	CopyTagsFromSource types.CopyTagsFromSource

	// A description propagated to every snapshot specified by the instance.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Tags to apply to every snapshot specified by the instance.
	TagSpecifications []*types.TagSpecification
}

type CreateSnapshotsOutput struct {

	// List of snapshots.
	Snapshots []*types.SnapshotInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSnapshots{}, middleware.After)
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
	if err = addOpCreateSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshots(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSnapshots",
	}
}
