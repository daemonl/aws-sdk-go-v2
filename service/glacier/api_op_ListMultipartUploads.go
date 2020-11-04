// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists in-progress multipart uploads for the specified vault. An
// in-progress multipart upload is a multipart upload that has been initiated by an
// InitiateMultipartUpload request, but has not yet been completed or aborted. The
// list returned in the List Multipart Upload response has no guaranteed order. The
// List Multipart Uploads operation supports pagination. By default, this operation
// returns up to 50 multipart uploads in the response. You should always check the
// response for a marker at which to continue the list; if there are no more items
// the marker is null. To return a list of multipart uploads that begins at a
// specific upload, set the marker request parameter to the value you obtained from
// a previous List Multipart Upload request. You can also limit the number of
// uploads returned in the response by specifying the limit parameter in the
// request. Note the difference between this operation and listing parts
// (ListParts). The List Multipart Uploads operation lists all multipart uploads
// for a vault and does not require a multipart upload ID. The List Parts operation
// requires a multipart upload ID since parts are associated with a single upload.
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions by
// default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access
// Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and the underlying REST API, see Working with
// Archives in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// and List Multipart Uploads
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-uploads.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) ListMultipartUploads(ctx context.Context, params *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error) {
	if params == nil {
		params = &ListMultipartUploadsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMultipartUploads", params, optFns, addOperationListMultipartUploadsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMultipartUploadsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving list of in-progress multipart uploads for an
// Amazon Glacier vault.
type ListMultipartUploadsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// Specifies the maximum number of uploads returned in the response body. If this
	// value is not specified, the List Uploads operation returns up to 50 uploads.
	Limit *string

	// An opaque string used for pagination. This value specifies the upload at which
	// the listing of uploads should begin. Get the marker value from a previous List
	// Uploads response. You need only include the marker if you are continuing the
	// pagination of results started in a previous List Uploads request.
	Marker *string
}

// Contains the Amazon S3 Glacier response to your request.
type ListMultipartUploadsOutput struct {

	// An opaque string that represents where to continue pagination of the results.
	// You use the marker in a new List Multipart Uploads request to obtain more
	// uploads in the list. If there are no more uploads, this value is null.
	Marker *string

	// A list of in-progress multipart uploads.
	UploadsList []*types.UploadListElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMultipartUploadsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMultipartUploads{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMultipartUploads{}, middleware.After)
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
	if err = addOpListMultipartUploadsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMultipartUploads(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListMultipartUploads(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListMultipartUploads",
	}
}
