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

// Restarts the stopped cluster that is specified by DBClusterIdentifier. For more
// information, see Stopping and Starting an Amazon DocumentDB Cluster
// (https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-stop-start.html).
func (c *Client) StartDBCluster(ctx context.Context, params *StartDBClusterInput, optFns ...func(*Options)) (*StartDBClusterOutput, error) {
	if params == nil {
		params = &StartDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDBCluster", params, optFns, addOperationStartDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDBClusterInput struct {

	// The identifier of the cluster to restart. Example: docdb-2019-05-28-15-24-52
	//
	// This member is required.
	DBClusterIdentifier *string
}

type StartDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartDBCluster{}, middleware.After)
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
	if err = addOpStartDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartDBCluster",
	}
}
