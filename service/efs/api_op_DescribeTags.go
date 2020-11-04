// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the tags associated with a file system. The order of tags returned in
// the response of one DescribeTags call and the order of tags returned across the
// responses of a multiple-call iteration (when using pagination) is unspecified.
// This operation requires permissions for the elasticfilesystem:DescribeTags
// action.
func (c *Client) DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTags", params, optFns, addOperationDescribeTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeTagsInput struct {

	// The ID of the file system whose tag set you want to retrieve.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) An opaque pagination token returned from a previous DescribeTags
	// operation (String). If present, it specifies to continue the list from where the
	// previous call left off.
	Marker *string

	// (Optional) The maximum number of file system tags to return in the response.
	// Currently, this number is automatically set to 100, and other values are
	// ignored. The response is paginated at 100 per page if you have more than 100
	// tags.
	MaxItems *int32
}

//
type DescribeTagsOutput struct {

	// Returns tags associated with the file system as an array of Tag objects.
	//
	// This member is required.
	Tags []*types.Tag

	// If the request included a Marker, the response returns that value in this field.
	Marker *string

	// If a value is present, there are more tags to return. In a subsequent request,
	// you can provide the value of NextMarker as the value of the Marker parameter in
	// your next request to retrieve the next set of tags.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTags{}, middleware.After)
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
	if err = addOpDescribeTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeTags",
	}
}
