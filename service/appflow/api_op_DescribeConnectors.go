// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the connectors vended by Amazon AppFlow for specified connector types.
// If you don't specify a connector type, this operation describes all connectors
// vended by Amazon AppFlow. If there are more connectors than can be returned in
// one page, the response contains a nextToken object, which can be be passed in to
// the next call to the DescribeConnectors API operation to retrieve the next page.
func (c *Client) DescribeConnectors(ctx context.Context, params *DescribeConnectorsInput, optFns ...func(*Options)) (*DescribeConnectorsOutput, error) {
	if params == nil {
		params = &DescribeConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectors", params, optFns, addOperationDescribeConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorsInput struct {

	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorTypes []types.ConnectorType

	// The pagination token for the next page of data.
	NextToken *string
}

type DescribeConnectorsOutput struct {

	// The configuration that is applied to the connectors used in the flow.
	ConnectorConfigurations map[string]*types.ConnectorConfiguration

	// The pagination token for the next page of data.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeConnectors",
	}
}
