// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Aurora global database spread across multiple AWS Regions. The global
// database contains a single primary cluster with read-write capability, and a
// read-only secondary cluster that receives data from the primary cluster through
// high-speed replication performed by the Aurora storage subsystem. You can create
// a global database that is initially empty, and then add a primary cluster and a
// secondary cluster to it. Or you can specify an existing Aurora cluster during
// the create operation, and this cluster becomes the primary cluster of the global
// database. This action only applies to Aurora DB clusters.
func (c *Client) CreateGlobalCluster(ctx context.Context, params *CreateGlobalClusterInput, optFns ...func(*Options)) (*CreateGlobalClusterOutput, error) {
	if params == nil {
		params = &CreateGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlobalCluster", params, optFns, addOperationCreateGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalClusterInput struct {

	// The name for your database of up to 64 alpha-numeric characters. If you do not
	// provide a name, Amazon Aurora will not create a database in the global database
	// cluster you are creating.
	DatabaseName *string

	// The deletion protection setting for the new global database. The global database
	// can't be deleted when deletion protection is enabled.
	DeletionProtection *bool

	// The name of the database engine to be used for this DB cluster.
	Engine *string

	// The engine version of the Aurora global database.
	EngineVersion *string

	// The cluster identifier of the new global database cluster.
	GlobalClusterIdentifier *string

	// The Amazon Resource Name (ARN) to use as the primary cluster of the global
	// database. This parameter is optional.
	SourceDBClusterIdentifier *string

	// The storage encryption setting for the new global database cluster.
	StorageEncrypted *bool
}

type CreateGlobalClusterOutput struct {

	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateGlobalCluster{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateGlobalCluster",
	}
}
