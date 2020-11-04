// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the resource types, the number of each resource type, and the total
// number of resources that AWS Config is recording in this region for your AWS
// account. Example
//
// * AWS Config is recording three resource types in the US East
// (Ohio) Region for your account: 25 EC2 instances, 20 IAM users, and 15 S3
// buckets.
//
// * You make a call to the GetDiscoveredResourceCounts action and
// specify that you want all resource types.
//
// * AWS Config returns the
// following:
//
// * The resource types (EC2 instances, IAM users, and S3 buckets).
//
// *
// The number of each resource type (25, 20, and 15).
//
// * The total number of all
// resources (60).
//
// The response is paginated. By default, AWS Config lists 100
// ResourceCount objects on each page. You can customize this number with the limit
// parameter. The response includes a nextToken string. To get the next page of
// results, run the request again and specify the string for the nextToken
// parameter. If you make a call to the GetDiscoveredResourceCounts action, you
// might not immediately receive resource counts in the following situations:
//
// *
// You are a new AWS Config customer.
//
// * You just enabled resource recording.
//
// It
// might take a few minutes for AWS Config to record and count your resources. Wait
// a few minutes and then retry the GetDiscoveredResourceCounts action.
func (c *Client) GetDiscoveredResourceCounts(ctx context.Context, params *GetDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error) {
	if params == nil {
		params = &GetDiscoveredResourceCountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDiscoveredResourceCounts", params, optFns, addOperationGetDiscoveredResourceCountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDiscoveredResourceCountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDiscoveredResourceCountsInput struct {

	// The maximum number of ResourceCount objects returned on each page. The default
	// is 100. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The comma-separated list that specifies the resource types that you want AWS
	// Config to return (for example, "AWS::EC2::Instance", "AWS::IAM::User"). If a
	// value for resourceTypes is not specified, AWS Config returns all resource types
	// that AWS Config is recording in the region for your account. If the
	// configuration recorder is turned off, AWS Config returns an empty list of
	// ResourceCount objects. If the configuration recorder is not recording a specific
	// resource type (for example, S3 buckets), that resource type is not returned in
	// the list of ResourceCount objects.
	ResourceTypes []*string
}

type GetDiscoveredResourceCountsOutput struct {

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// The list of ResourceCount objects. Each object is listed in descending order by
	// the number of resources.
	ResourceCounts []*types.ResourceCount

	// The total number of resources that AWS Config is recording in the region for
	// your account. If you specify resource types in the request, AWS Config returns
	// only the total number of resources for those resource types. Example
	//
	// * AWS
	// Config is recording three resource types in the US East (Ohio) Region for your
	// account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets, for a total of 60
	// resources.
	//
	// * You make a call to the GetDiscoveredResourceCounts action and
	// specify the resource type, "AWS::EC2::Instances", in the request.
	//
	// * AWS Config
	// returns 25 for totalDiscoveredResources.
	TotalDiscoveredResources *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDiscoveredResourceCountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDiscoveredResourceCounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDiscoveredResourceCounts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetDiscoveredResourceCounts",
	}
}
