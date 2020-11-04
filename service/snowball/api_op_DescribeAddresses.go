// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a specified number of ADDRESS objects. Calling this API in one of the US
// regions will return addresses from the list of all addresses associated with
// this account in all US regions.
func (c *Client) DescribeAddresses(ctx context.Context, params *DescribeAddressesInput, optFns ...func(*Options)) (*DescribeAddressesOutput, error) {
	if params == nil {
		params = &DescribeAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddresses", params, optFns, addOperationDescribeAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressesInput struct {

	// The number of ADDRESS objects to return.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// ADDRESS objects, you have the option of specifying a value for NextToken as the
	// starting point for your list of returned addresses.
	NextToken *string
}

type DescribeAddressesOutput struct {

	// The Snow device shipping addresses that were created for this account.
	Addresses []*types.Address

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next DescribeAddresses call, your list of returned addresses will
	// start from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAddresses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddresses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "DescribeAddresses",
	}
}
