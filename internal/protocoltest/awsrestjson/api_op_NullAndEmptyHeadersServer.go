// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Null and empty headers are not sent over the wire.
func (c *Client) NullAndEmptyHeadersServer(ctx context.Context, params *NullAndEmptyHeadersServerInput, optFns ...func(*Options)) (*NullAndEmptyHeadersServerOutput, error) {
	if params == nil {
		params = &NullAndEmptyHeadersServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NullAndEmptyHeadersServer", params, optFns, addOperationNullAndEmptyHeadersServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NullAndEmptyHeadersServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NullAndEmptyHeadersServerInput struct {
	A *string

	B *string

	C []*string
}

type NullAndEmptyHeadersServerOutput struct {
	A *string

	B *string

	C []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationNullAndEmptyHeadersServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpNullAndEmptyHeadersServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpNullAndEmptyHeadersServer{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNullAndEmptyHeadersServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNullAndEmptyHeadersServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NullAndEmptyHeadersServer",
	}
}
