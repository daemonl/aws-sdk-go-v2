// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Greengrass and Cloud Deployments Deploys the system instance to the target
// specified in CreateSystemInstance. Greengrass Deployments If the system or any
// workflows and entities have been updated before this action is called, then the
// deployment will create a new Amazon Simple Storage Service resource file and
// then deploy it. Since this action creates a Greengrass deployment on the
// caller's behalf, the calling identity must have write permissions to the
// specified Greengrass group. Otherwise, the call will fail with an authorization
// error. For information about the artifacts that get added to your Greengrass
// core device when you use this API, see AWS IoT Things Graph and AWS IoT
// Greengrass
// (https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-greengrass.html).
func (c *Client) DeploySystemInstance(ctx context.Context, params *DeploySystemInstanceInput, optFns ...func(*Options)) (*DeploySystemInstanceOutput, error) {
	if params == nil {
		params = &DeploySystemInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeploySystemInstance", params, optFns, addOperationDeploySystemInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeploySystemInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeploySystemInstanceInput struct {

	// The ID of the system instance. This value is returned by the
	// CreateSystemInstance action. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:deployment:DEPLOYMENTNAME
	Id *string
}

type DeploySystemInstanceOutput struct {

	// An object that contains summary information about a system instance that was
	// deployed.
	//
	// This member is required.
	Summary *types.SystemInstanceSummary

	// The ID of the Greengrass deployment used to deploy the system instance.
	GreengrassDeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeploySystemInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeploySystemInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeploySystemInstance{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeploySystemInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeploySystemInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "DeploySystemInstance",
	}
}
