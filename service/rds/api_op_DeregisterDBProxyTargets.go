// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove the association between one or more DBProxyTarget data structures and a
// DBProxyTargetGroup.
func (c *Client) DeregisterDBProxyTargets(ctx context.Context, params *DeregisterDBProxyTargetsInput, optFns ...func(*Options)) (*DeregisterDBProxyTargetsOutput, error) {
	if params == nil {
		params = &DeregisterDBProxyTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterDBProxyTargets", params, optFns, addOperationDeregisterDBProxyTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterDBProxyTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDBProxyTargetsInput struct {

	// The identifier of the DBProxy that is associated with the DBProxyTargetGroup.
	//
	// This member is required.
	DBProxyName *string

	// One or more DB cluster identifiers.
	DBClusterIdentifiers []*string

	// One or more DB instance identifiers.
	DBInstanceIdentifiers []*string

	// The identifier of the DBProxyTargetGroup.
	TargetGroupName *string
}

type DeregisterDBProxyTargetsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterDBProxyTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeregisterDBProxyTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeregisterDBProxyTargets{}, middleware.After)
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
	if err = addOpDeregisterDBProxyTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDBProxyTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterDBProxyTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeregisterDBProxyTargets",
	}
}
