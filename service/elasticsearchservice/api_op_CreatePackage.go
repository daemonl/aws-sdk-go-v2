// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a package for use with Amazon ES domains.
func (c *Client) CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) {
	if params == nil {
		params = &CreatePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePackage", params, optFns, addOperationCreatePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to CreatePackage operation.
type CreatePackageInput struct {

	// Unique identifier for the package.
	//
	// This member is required.
	PackageName *string

	// The customer S3 location PackageSource for importing the package.
	//
	// This member is required.
	PackageSource *types.PackageSource

	// Type of package. Currently supports only TXT-DICTIONARY.
	//
	// This member is required.
	PackageType types.PackageType

	// Description of the package.
	PackageDescription *string
}

// Container for response returned by CreatePackage operation.
type CreatePackageOutput struct {

	// Information about the package PackageDetails.
	PackageDetails *types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackage{}, middleware.After)
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
	if err = addOpCreatePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreatePackage",
	}
}
