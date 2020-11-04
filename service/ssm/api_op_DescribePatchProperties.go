// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You can use
// the reported properties in the filters you specify in requests for actions such
// as CreatePatchBaseline, UpdatePatchBaseline, DescribeAvailablePatches, and
// DescribePatchBaselines. The following section lists the properties that can be
// used in filters for each major operating system type: AMAZON_LINUX Valid
// properties: PRODUCT, CLASSIFICATION, SEVERITY AMAZON_LINUX_2 Valid properties:
// PRODUCT, CLASSIFICATION, SEVERITY CENTOS Valid properties: PRODUCT,
// CLASSIFICATION, SEVERITY DEBIAN Valid properties: PRODUCT, PRIORITY ORACLE_LINUX
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY REDHAT_ENTERPRISE_LINUX
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY SUSE Valid properties:
// PRODUCT, CLASSIFICATION, SEVERITY UBUNTU Valid properties: PRODUCT, PRIORITY
// WINDOWS Valid properties: PRODUCT, PRODUCT_FAMILY, CLASSIFICATION, MSRC_SEVERITY
func (c *Client) DescribePatchProperties(ctx context.Context, params *DescribePatchPropertiesInput, optFns ...func(*Options)) (*DescribePatchPropertiesOutput, error) {
	if params == nil {
		params = &DescribePatchPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchProperties", params, optFns, addOperationDescribePatchPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchPropertiesInput struct {

	// The operating system type for which to list patches.
	//
	// This member is required.
	OperatingSystem types.OperatingSystem

	// The patch property for which you want to view patch details.
	//
	// This member is required.
	Property types.PatchProperty

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Indicates whether to list patches for the Windows operating system or for
	// Microsoft applications. Not applicable for Linux operating systems.
	PatchSet types.PatchSet
}

type DescribePatchPropertiesOutput struct {

	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// A list of the properties for patches matching the filter request parameters.
	Properties []map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePatchPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchProperties{}, middleware.After)
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
	if err = addOpDescribePatchPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePatchProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchProperties",
	}
}
