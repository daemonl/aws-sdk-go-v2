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

// Asynchronously uploads one or more entity definitions to the user's namespace.
// The document parameter is required if syncWithPublicNamespace and
// deleteExistingEntites are false. If the syncWithPublicNamespace parameter is set
// to true, the user's namespace will synchronize with the latest version of the
// public namespace. If deprecateExistingEntities is set to true, all entities in
// the latest version will be deleted before the new DefinitionDocument is
// uploaded. When a user uploads entity definitions for the first time, the service
// creates a new namespace for the user. The new namespace tracks the public
// namespace. Currently users can have only one namespace. The namespace version
// increments whenever a user uploads entity definitions that are
// backwards-incompatible and whenever a user sets the syncWithPublicNamespace
// parameter or the deprecateExistingEntities parameter to true. The IDs for all of
// the entities should be in URN format. Each entity must be in the user's
// namespace. Users can't create entities in the public namespace, but entity
// definitions can refer to entities in the public namespace. Valid entities are
// Device, DeviceModel, Service, Capability, State, Action, Event, Property,
// Mapping, Enum.
func (c *Client) UploadEntityDefinitions(ctx context.Context, params *UploadEntityDefinitionsInput, optFns ...func(*Options)) (*UploadEntityDefinitionsOutput, error) {
	if params == nil {
		params = &UploadEntityDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadEntityDefinitions", params, optFns, addOperationUploadEntityDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadEntityDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadEntityDefinitionsInput struct {

	// A Boolean that specifies whether to deprecate all entities in the latest version
	// before uploading the new DefinitionDocument. If set to true, the upload will
	// create a new namespace version.
	DeprecateExistingEntities *bool

	// The DefinitionDocument that defines the updated entities.
	Document *types.DefinitionDocument

	// A Boolean that specifies whether to synchronize with the latest version of the
	// public namespace. If set to true, the upload will create a new namespace
	// version.
	SyncWithPublicNamespace *bool
}

type UploadEntityDefinitionsOutput struct {

	// The ID that specifies the upload action. You can use this to track the status of
	// the upload.
	//
	// This member is required.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadEntityDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUploadEntityDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUploadEntityDefinitions{}, middleware.After)
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
	if err = addOpUploadEntityDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadEntityDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUploadEntityDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "UploadEntityDefinitions",
	}
}
