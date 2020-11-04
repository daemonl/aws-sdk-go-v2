// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a stack's Elastic Load Balancing instances. This call accepts only one
// resource-identifying parameter. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeElasticLoadBalancers(ctx context.Context, params *DescribeElasticLoadBalancersInput, optFns ...func(*Options)) (*DescribeElasticLoadBalancersOutput, error) {
	if params == nil {
		params = &DescribeElasticLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticLoadBalancers", params, optFns, addOperationDescribeElasticLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeElasticLoadBalancersInput struct {

	// A list of layer IDs. The action describes the Elastic Load Balancing instances
	// for the specified layers.
	LayerIds []*string

	// A stack ID. The action describes the stack's Elastic Load Balancing instances.
	StackId *string
}

// Contains the response to a DescribeElasticLoadBalancers request.
type DescribeElasticLoadBalancersOutput struct {

	// A list of ElasticLoadBalancer objects that describe the specified Elastic Load
	// Balancing instances.
	ElasticLoadBalancers []*types.ElasticLoadBalancer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeElasticLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeElasticLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeElasticLoadBalancers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticLoadBalancers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeElasticLoadBalancers",
	}
}
