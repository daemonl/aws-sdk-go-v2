// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists origin access identities.
func (c *Client) ListCloudFrontOriginAccessIdentities(ctx context.Context, params *ListCloudFrontOriginAccessIdentitiesInput, optFns ...func(*Options)) (*ListCloudFrontOriginAccessIdentitiesOutput, error) {
	if params == nil {
		params = &ListCloudFrontOriginAccessIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCloudFrontOriginAccessIdentities", params, optFns, addOperationListCloudFrontOriginAccessIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCloudFrontOriginAccessIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list origin access identities.
type ListCloudFrontOriginAccessIdentitiesInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// origin access identities. The results include identities in the list that occur
	// after the marker. To get the next page of results, set the Marker to the value
	// of the NextMarker from the current page's response (which is also the ID of the
	// last identity on that page).
	Marker *string

	// The maximum number of origin access identities you want in the response body.
	MaxItems *string
}

// The returned result of the corresponding request.
type ListCloudFrontOriginAccessIdentitiesOutput struct {

	// The CloudFrontOriginAccessIdentityList type.
	CloudFrontOriginAccessIdentityList *types.CloudFrontOriginAccessIdentityList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCloudFrontOriginAccessIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListCloudFrontOriginAccessIdentities",
	}
}
