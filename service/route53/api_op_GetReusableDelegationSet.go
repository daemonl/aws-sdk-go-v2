// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a specified reusable delegation set, including the
// four name servers that are assigned to the delegation set.
func (c *Client) GetReusableDelegationSet(ctx context.Context, params *GetReusableDelegationSetInput, optFns ...func(*Options)) (*GetReusableDelegationSetOutput, error) {
	if params == nil {
		params = &GetReusableDelegationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReusableDelegationSet", params, optFns, addOperationGetReusableDelegationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReusableDelegationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a specified reusable delegation set.
type GetReusableDelegationSetInput struct {

	// The ID of the reusable delegation set that you want to get a list of name
	// servers for.
	//
	// This member is required.
	Id *string
}

// A complex type that contains the response to the GetReusableDelegationSet
// request.
type GetReusableDelegationSetOutput struct {

	// A complex type that contains information about the reusable delegation set.
	//
	// This member is required.
	DelegationSet *types.DelegationSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetReusableDelegationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetReusableDelegationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetReusableDelegationSet{}, middleware.After)
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
	if err = addOpGetReusableDelegationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReusableDelegationSet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetReusableDelegationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetReusableDelegationSet",
	}
}
