// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes all packages available to Amazon ES. Includes options for filtering,
// limiting the number of results, and pagination.
func (c *Client) DescribePackages(ctx context.Context, params *DescribePackagesInput, optFns ...func(*Options)) (*DescribePackagesOutput, error) {
	if params == nil {
		params = &DescribePackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackages", params, optFns, addOperationDescribePackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to DescribePackage operation.
type DescribePackagesInput struct {

	// Only returns packages that match the DescribePackagesFilterList values.
	Filters []*types.DescribePackagesFilter

	// Limits results to a maximum number of packages.
	MaxResults *int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response returned by DescribePackages operation.
type DescribePackagesOutput struct {
	NextToken *string

	// List of PackageDetails objects.
	PackageDetailsList []*types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackages(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribePackages",
	}
}
