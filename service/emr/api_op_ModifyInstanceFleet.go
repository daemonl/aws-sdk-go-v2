// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the target On-Demand and target Spot capacities for the instance fleet
// with the specified InstanceFleetID within the cluster specified using ClusterID.
// The call either succeeds or fails atomically. The instance fleet configuration
// is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x
// versions.
func (c *Client) ModifyInstanceFleet(ctx context.Context, params *ModifyInstanceFleetInput, optFns ...func(*Options)) (*ModifyInstanceFleetOutput, error) {
	if params == nil {
		params = &ModifyInstanceFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceFleet", params, optFns, addOperationModifyInstanceFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceFleetInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The unique identifier of the instance fleet.
	//
	// This member is required.
	InstanceFleet *types.InstanceFleetModifyConfig
}

type ModifyInstanceFleetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyInstanceFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyInstanceFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyInstanceFleet{}, middleware.After)
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
	if err = addOpModifyInstanceFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ModifyInstanceFleet",
	}
}
