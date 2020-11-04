// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an Amazon S3 presigned URL for an update file associated with a
// specified JobId.
func (c *Client) GetSoftwareUpdates(ctx context.Context, params *GetSoftwareUpdatesInput, optFns ...func(*Options)) (*GetSoftwareUpdatesOutput, error) {
	if params == nil {
		params = &GetSoftwareUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSoftwareUpdates", params, optFns, addOperationGetSoftwareUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSoftwareUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSoftwareUpdatesInput struct {

	// The ID for a job that you want to get the software update file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string
}

type GetSoftwareUpdatesOutput struct {

	// The Amazon S3 presigned URL for the update file associated with the specified
	// JobId value. The software update will be available for 2 days after this request
	// is made. To access an update after the 2 days have passed, you'll have to make
	// another call to GetSoftwareUpdates.
	UpdatesURI *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSoftwareUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSoftwareUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSoftwareUpdates{}, middleware.After)
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
	if err = addOpGetSoftwareUpdatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSoftwareUpdates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSoftwareUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetSoftwareUpdates",
	}
}
