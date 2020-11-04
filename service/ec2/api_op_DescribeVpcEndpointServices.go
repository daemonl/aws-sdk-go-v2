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

// Describes available services to which you can create a VPC endpoint.
func (c *Client) DescribeVpcEndpointServices(ctx context.Context, params *DescribeVpcEndpointServicesInput, optFns ...func(*Options)) (*DescribeVpcEndpointServicesOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpointServices", params, optFns, addOperationDescribeVpcEndpointServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeVpcEndpointServices.
type DescribeVpcEndpointServicesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * service-name - The name of the service.
	//
	// * tag: - The
	// key/value combination of a tag assigned to the resource. Use the tag key in the
	// filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key - The
	// key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	Filters []*types.Filter

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// Constraint: If the value is greater than 1,000, we return only 1,000 items.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// prior call.)
	NextToken *string

	// One or more service names.
	ServiceNames []*string
}

// Contains the output of DescribeVpcEndpointServices.
type DescribeVpcEndpointServicesOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the service.
	ServiceDetails []*types.ServiceDetail

	// A list of supported services.
	ServiceNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcEndpointServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointServices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointServices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcEndpointServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointServices",
	}
}
