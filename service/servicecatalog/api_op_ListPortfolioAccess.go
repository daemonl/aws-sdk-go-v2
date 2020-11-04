// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the account IDs that have access to the specified portfolio. A delegated
// admin can list the accounts that have access to the shared portfolio. Note that
// if a delegated admin is de-registered, they can no longer perform this
// operation.
func (c *Client) ListPortfolioAccess(ctx context.Context, params *ListPortfolioAccessInput, optFns ...func(*Options)) (*ListPortfolioAccessOutput, error) {
	if params == nil {
		params = &ListPortfolioAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPortfolioAccess", params, optFns, addOperationListPortfolioAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPortfolioAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPortfolioAccessInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The ID of an organization node the portfolio is shared with. All children of
	// this node with an inherited portfolio share will be returned.
	OrganizationParentId *string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListPortfolioAccessOutput struct {

	// Information about the AWS accounts with access to the portfolio.
	AccountIds []*string

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPortfolioAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPortfolioAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPortfolioAccess{}, middleware.After)
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
	if err = addOpListPortfolioAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPortfolioAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPortfolioAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListPortfolioAccess",
	}
}
