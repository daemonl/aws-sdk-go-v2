// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Changes information about an ClientCertificate resource.
func (c *Client) UpdateClientCertificate(ctx context.Context, params *UpdateClientCertificateInput, optFns ...func(*Options)) (*UpdateClientCertificateOutput, error) {
	if params == nil {
		params = &UpdateClientCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClientCertificate", params, optFns, addOperationUpdateClientCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change information about an ClientCertificate resource.
type UpdateClientCertificateInput struct {

	// [Required] The identifier of the ClientCertificate resource to be updated.
	//
	// This member is required.
	ClientCertificateId *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint. Client certificates are used
// to authenticate an API by the backend server. To authenticate an API client (or
// user), use IAM roles and policies, a custom Authorizer or an Amazon Cognito user
// pool. Use Client-Side Certificate
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type UpdateClientCertificateOutput struct {

	// The identifier of the client certificate.
	ClientCertificateId *string

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time

	// The description of the client certificate.
	Description *string

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateClientCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClientCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClientCertificate{}, middleware.After)
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
	if err = addOpUpdateClientCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClientCertificate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateClientCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateClientCertificate",
	}
}
