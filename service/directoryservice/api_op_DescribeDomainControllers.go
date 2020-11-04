// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about any domain controllers in your directory.
func (c *Client) DescribeDomainControllers(ctx context.Context, params *DescribeDomainControllersInput, optFns ...func(*Options)) (*DescribeDomainControllersOutput, error) {
	if params == nil {
		params = &DescribeDomainControllersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainControllers", params, optFns, addOperationDescribeDomainControllersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainControllersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainControllersInput struct {

	// Identifier of the directory for which to retrieve the domain controller
	// information.
	//
	// This member is required.
	DirectoryId *string

	// A list of identifiers for the domain controllers whose information will be
	// provided.
	DomainControllerIds []*string

	// The maximum number of items to return.
	Limit *int32

	// The DescribeDomainControllers.NextToken value from a previous call to
	// DescribeDomainControllers. Pass null if this is the first call.
	NextToken *string
}

type DescribeDomainControllersOutput struct {

	// List of the DomainController objects that were retrieved.
	DomainControllers []*types.DomainController

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeDomainControllers retrieve the next
	// set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDomainControllersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDomainControllers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDomainControllers{}, middleware.After)
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
	if err = addOpDescribeDomainControllersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainControllers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomainControllers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeDomainControllers",
	}
}
