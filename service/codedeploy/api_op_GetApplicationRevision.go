// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about an application revision.
func (c *Client) GetApplicationRevision(ctx context.Context, params *GetApplicationRevisionInput, optFns ...func(*Options)) (*GetApplicationRevisionOutput, error) {
	if params == nil {
		params = &GetApplicationRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplicationRevision", params, optFns, addOperationGetApplicationRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetApplicationRevision operation.
type GetApplicationRevisionInput struct {

	// The name of the application that corresponds to the revision.
	//
	// This member is required.
	ApplicationName *string

	// Information about the application revision to get, including type and location.
	//
	// This member is required.
	Revision *types.RevisionLocation
}

// Represents the output of a GetApplicationRevision operation.
type GetApplicationRevisionOutput struct {

	// The name of the application that corresponds to the revision.
	ApplicationName *string

	// Additional information about the revision, including type and location.
	Revision *types.RevisionLocation

	// General information about the revision.
	RevisionInfo *types.GenericRevisionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetApplicationRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetApplicationRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetApplicationRevision{}, middleware.After)
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
	if err = addOpGetApplicationRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationRevision(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplicationRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetApplicationRevision",
	}
}
