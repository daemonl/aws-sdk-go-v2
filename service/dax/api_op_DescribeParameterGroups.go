// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of parameter group descriptions. If a parameter group name is
// specified, the list will contain only the descriptions for that group.
func (c *Client) DescribeParameterGroups(ctx context.Context, params *DescribeParameterGroupsInput, optFns ...func(*Options)) (*DescribeParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeParameterGroups", params, optFns, addOperationDescribeParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeParameterGroupsInput struct {

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved. The value for MaxResults must be between
	// 20 and 100.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The names of the parameter groups.
	ParameterGroupNames []*string
}

type DescribeParameterGroupsOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	// An array of parameter groups. Each element in the array represents one parameter
	// group.
	ParameterGroups []*types.ParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeParameterGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeParameterGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeParameterGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "DescribeParameterGroups",
	}
}
