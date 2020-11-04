// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates a Stage.
func (c *Client) UpdateStage(ctx context.Context, params *UpdateStageInput, optFns ...func(*Options)) (*UpdateStageOutput, error) {
	if params == nil {
		params = &UpdateStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStage", params, optFns, addOperationUpdateStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a Stage.
type UpdateStageInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The stage name. Stage names can only contain alphanumeric characters, hyphens,
	// and underscores. Maximum length is 128 characters.
	//
	// This member is required.
	StageName *string

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// The identifier of a client certificate for a Stage.
	ClientCertificateId *string

	// The default route settings for the stage.
	DefaultRouteSettings *types.RouteSettings

	// The deployment identifier for the API stage. Can't be updated if autoDeploy is
	// enabled.
	DeploymentId *string

	// The description for the API stage.
	Description *string

	// Route settings for the stage.
	RouteSettings map[string]*types.RouteSettings

	// A map that defines the stage variables for a Stage. Variable names can have
	// alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]*string
}

type UpdateStageOutput struct {

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether a stage is managed by API Gateway. If you created an API using
	// quick create, the $default stage is managed by API Gateway. You can't modify the
	// $default stage.
	ApiGatewayManaged *bool

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// The identifier of a client certificate for a Stage. Supported only for WebSocket
	// APIs.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// Default route settings for the stage.
	DefaultRouteSettings *types.RouteSettings

	// The identifier of the Deployment that the Stage is associated with. Can't be
	// updated if autoDeploy is enabled.
	DeploymentId *string

	// The description of the stage.
	Description *string

	// Describes the status of the last deployment of a stage. Supported only for
	// stages with autoDeploy enabled.
	LastDeploymentStatusMessage *string

	// The timestamp when the stage was last updated.
	LastUpdatedDate *time.Time

	// Route settings for the stage, by routeKey.
	RouteSettings map[string]*types.RouteSettings

	// The name of the stage.
	StageName *string

	// A map that defines the stage variables for a stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]*string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStage{}, middleware.After)
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
	if err = addOpUpdateStageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateStage",
	}
}
