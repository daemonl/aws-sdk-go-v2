// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML name, changing the
// wrapper name.
func (c *Client) HttpPayloadWithXmlName(ctx context.Context, params *HttpPayloadWithXmlNameInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNameOutput, error) {
	if params == nil {
		params = &HttpPayloadWithXmlNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadWithXmlName", params, optFns, addOperationHttpPayloadWithXmlNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadWithXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNameInput struct {
	Nested *types.PayloadWithXmlName
}

type HttpPayloadWithXmlNameOutput struct {
	Nested *types.PayloadWithXmlName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpPayloadWithXmlNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlName{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlName(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpPayloadWithXmlName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadWithXmlName",
	}
}
