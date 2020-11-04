// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the products for the specified portfolio or all products.
func (c *Client) SearchProductsAsAdmin(ctx context.Context, params *SearchProductsAsAdminInput, optFns ...func(*Options)) (*SearchProductsAsAdminOutput, error) {
	if params == nil {
		params = &SearchProductsAsAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProductsAsAdmin", params, optFns, addOperationSearchProductsAsAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProductsAsAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProductsAsAdminInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The search filters. If no search filters are specified, the output includes all
	// products to which the administrator has access.
	Filters map[string][]*string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The portfolio identifier.
	PortfolioId *string

	// Access level of the source of the product.
	ProductSource types.ProductSource

	// The sort field. If no value is specified, the results are not sorted.
	SortBy types.ProductViewSortBy

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder
}

type SearchProductsAsAdminOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the product views.
	ProductViewDetails []*types.ProductViewDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchProductsAsAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProductsAsAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProductsAsAdmin{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProductsAsAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchProductsAsAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SearchProductsAsAdmin",
	}
}
