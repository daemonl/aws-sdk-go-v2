// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the automatic tape creation policies for a gateway. If there are no
// automatic tape creation policies for the gateway, it returns an empty list. This
// operation is only supported for tape gateways.
func (c *Client) ListAutomaticTapeCreationPolicies(ctx context.Context, params *ListAutomaticTapeCreationPoliciesInput, optFns ...func(*Options)) (*ListAutomaticTapeCreationPoliciesOutput, error) {
	if params == nil {
		params = &ListAutomaticTapeCreationPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAutomaticTapeCreationPolicies", params, optFns, addOperationListAutomaticTapeCreationPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAutomaticTapeCreationPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAutomaticTapeCreationPoliciesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type ListAutomaticTapeCreationPoliciesOutput struct {

	// Gets a listing of information about the gateway's automatic tape creation
	// policies, including the automatic tape creation rules and the gateway that is
	// using the policies.
	AutomaticTapeCreationPolicyInfos []*types.AutomaticTapeCreationPolicyInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAutomaticTapeCreationPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAutomaticTapeCreationPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAutomaticTapeCreationPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAutomaticTapeCreationPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAutomaticTapeCreationPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListAutomaticTapeCreationPolicies",
	}
}
