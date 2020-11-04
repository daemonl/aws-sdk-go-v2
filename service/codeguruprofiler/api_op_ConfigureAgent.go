// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

//
func (c *Client) ConfigureAgent(ctx context.Context, params *ConfigureAgentInput, optFns ...func(*Options)) (*ConfigureAgentOutput, error) {
	if params == nil {
		params = &ConfigureAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfigureAgent", params, optFns, addOperationConfigureAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfigureAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the configureAgentRequest.
type ConfigureAgentInput struct {

	//
	//
	// This member is required.
	ProfilingGroupName *string

	//
	FleetInstanceId *string

	Metadata map[string]*string
}

// The structure representing the configureAgentResponse.
type ConfigureAgentOutput struct {

	//
	//
	// This member is required.
	Configuration *types.AgentConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConfigureAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConfigureAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConfigureAgent{}, middleware.After)
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
	if err = addOpConfigureAgentValidationMiddleware(stack); err != nil {
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
