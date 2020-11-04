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

// Describes the customizations associated with the provided AWS account and Amazon
// QuickSight namespace in an AWS Region. The QuickSight console evaluates which
// customizations to apply by running this API operation with the Resolved flag
// included. To determine what customizations display when you run this command, it
// can help to visualize the relationship of the entities involved.
//
// * AWS Account
// - The AWS account exists at the top of the hierarchy. It has the potential to
// use all of the AWS Regions and AWS Services. When you subscribe to QuickSight,
// you choose one AWS Region to use as your home Region. That's where your free
// SPICE capacity is located. You can use QuickSight in any supported AWS
// Region.
//
// * AWS Region - In each AWS Region where you sign in to QuickSight at
// least once, QuickSight acts as a separate instance of the same service. If you
// have a user directory, it resides in us-east-1, which is the US East (N.
// Virginia). Generally speaking, these users have access to QuickSight in any AWS
// Region, unless they are constrained to a namespace. To run the command in a
// different AWS Region, you change your Region settings. If you're using the AWS
// CLI, you can use one of the following options:
//
// * Use command line options
// (https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-options.html).
//
// *
// Use named profiles
// (https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html).
//
// *
// Run aws configure to change your default AWS Region. Use Enter to key the same
// settings for your keys. For more information, see Configuring the AWS CLI
// (https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html).
//
// *
// Namespace - A QuickSight namespace is a partition that contains users and assets
// (data sources, datasets, dashboards, and so on). To access assets that are in a
// specific namespace, users and groups must also be part of the same namespace.
// People who share a namespace are completely isolated from users and assets in
// other namespaces, even if they are in the same AWS account and AWS Region.
//
// *
// Applied customizations - Within an AWS Region, a set of QuickSight
// customizations can apply to an AWS account or to a namespace. Settings that you
// apply to a namespace override settings that you apply to an AWS account. All
// settings are isolated to a single AWS Region. To apply them in other AWS
// Regions, run the CreateAccountCustomization command in each AWS Region where you
// want to apply the same customizations.
func (c *Client) DescribeAccountCustomization(ctx context.Context, params *DescribeAccountCustomizationInput, optFns ...func(*Options)) (*DescribeAccountCustomizationOutput, error) {
	if params == nil {
		params = &DescribeAccountCustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountCustomization", params, optFns, addOperationDescribeAccountCustomizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountCustomizationInput struct {

	// The ID for the AWS account that you want to describe QuickSight customizations
	// for.
	//
	// This member is required.
	AwsAccountId *string

	// The QuickSight namespace that you want to describe QuickSight customizations
	// for.
	Namespace *string

	// The Resolved flag works with the other parameters to determine which view of
	// QuickSight customizations is returned. You can add this flag to your command to
	// use the same view that QuickSight uses to identify which customizations to apply
	// to the console. Omit this flag, or set it to no-resolved, to reveal
	// customizations that are configured at different levels.
	Resolved *bool
}

type DescribeAccountCustomizationOutput struct {

	// The QuickSight customizations that exist in the current AWS Region.
	AccountCustomization *types.AccountCustomization

	// The Amazon Resource Name (ARN) of the customization that's associated with this
	// AWS account.
	Arn *string

	// The ID for the AWS account that you're describing.
	AwsAccountId *string

	// The QuickSight namespace that you're describing.
	Namespace *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountCustomizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountCustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountCustomization{}, middleware.After)
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
	if err = addOpDescribeAccountCustomizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountCustomization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountCustomization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAccountCustomization",
	}
}
