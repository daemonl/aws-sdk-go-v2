// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a simulation application with a specific revision id.
func (c *Client) CreateSimulationApplicationVersion(ctx context.Context, params *CreateSimulationApplicationVersionInput, optFns ...func(*Options)) (*CreateSimulationApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateSimulationApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSimulationApplicationVersion", params, optFns, addOperationCreateSimulationApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSimulationApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSimulationApplicationVersionInput struct {

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string

	// The current revision id for the simulation application. If you provide a value
	// and it matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string
}

type CreateSimulationApplicationVersionOutput struct {

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The revision ID of the simulation application.
	RevisionId *string

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The sources of the simulation application.
	Sources []*types.Source

	// The version of the simulation application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSimulationApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSimulationApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSimulationApplicationVersion{}, middleware.After)
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
	if err = addOpCreateSimulationApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSimulationApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSimulationApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateSimulationApplicationVersion",
	}
}
