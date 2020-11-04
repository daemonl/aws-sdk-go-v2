// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a cache parameter group. You can modify up to 20
// parameters in a single request by submitting a list parameter name and value
// pairs.
func (c *Client) ModifyCacheParameterGroup(ctx context.Context, params *ModifyCacheParameterGroupInput, optFns ...func(*Options)) (*ModifyCacheParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyCacheParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCacheParameterGroup", params, optFns, addOperationModifyCacheParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ModifyCacheParameterGroup operation.
type ModifyCacheParameterGroupInput struct {

	// The name of the cache parameter group to modify.
	//
	// This member is required.
	CacheParameterGroupName *string

	// An array of parameter names and values for the parameter update. You must supply
	// at least one parameter name and value; subsequent arguments are optional. A
	// maximum of 20 parameters may be modified per request.
	//
	// This member is required.
	ParameterNameValues []*types.ParameterNameValue
}

// Represents the output of one of the following operations:
//
// *
// ModifyCacheParameterGroup
//
// * ResetCacheParameterGroup
type ModifyCacheParameterGroupOutput struct {

	// The name of the cache parameter group.
	CacheParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyCacheParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCacheParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCacheParameterGroup{}, middleware.After)
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
	if err = addOpModifyCacheParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCacheParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCacheParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyCacheParameterGroup",
	}
}
