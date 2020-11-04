// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the template alias for a template.
func (c *Client) DescribeTemplateAlias(ctx context.Context, params *DescribeTemplateAliasInput, optFns ...func(*Options)) (*DescribeTemplateAliasOutput, error) {
	if params == nil {
		params = &DescribeTemplateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTemplateAlias", params, optFns, addOperationDescribeTemplateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTemplateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTemplateAliasInput struct {

	// The name of the template alias that you want to describe. If you name a specific
	// alias, you describe the version that the alias points to. You can specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter. The keyword $PUBLISHED doesn't apply to templates.
	//
	// This member is required.
	AliasName *string

	// The ID of the AWS account that contains the template alias that you're
	// describing.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string
}

type DescribeTemplateAliasOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Information about the template alias.
	TemplateAlias *types.TemplateAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTemplateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTemplateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTemplateAlias{}, middleware.After)
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
	if err = addOpDescribeTemplateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTemplateAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTemplateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeTemplateAlias",
	}
}
