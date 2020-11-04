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

// Creates a new OpsItem. You must have permission in AWS Identity and Access
// Management (IAM) to create a new OpsItem. For more information, see Getting
// started with OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide. Operations engineers and IT professionals
// use OpsCenter to view, investigate, and remediate operational issues impacting
// the performance and health of their AWS resources. For more information, see AWS
// Systems Manager OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html) in
// the AWS Systems Manager User Guide.
func (c *Client) CreateOpsItem(ctx context.Context, params *CreateOpsItemInput, optFns ...func(*Options)) (*CreateOpsItemOutput, error) {
	if params == nil {
		params = &CreateOpsItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOpsItem", params, optFns, addOperationCreateOpsItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOpsItemInput struct {

	// Information about the OpsItem.
	//
	// This member is required.
	Description *string

	// The origin of the OpsItem, such as Amazon EC2 or Systems Manager. The source
	// name can't contain the following strings: aws, amazon, and amzn.
	//
	// This member is required.
	Source *string

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	//
	// This member is required.
	Title *string

	// Specify a category to assign to an OpsItem.
	Category *string

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent when
	// this OpsItem is edited or changed.
	Notifications []*types.OpsItemNotification

	// Operational data is custom data that provides useful reference details about the
	// OpsItem. For example, you can specify log files, error strings, license keys,
	// troubleshooting tips, or other relevant data. You enter operational data as
	// key-value pairs. The key has a maximum length of 128 characters. The value has a
	// maximum size of 20 KB. Operational data keys can't begin with the following:
	// amazon, aws, amzn, ssm, /amazon, /aws, /amzn, /ssm. You can choose to make the
	// data searchable by other users in the account or you can restrict search access.
	// Searchable data means that all users with access to the OpsItem Overview page
	// (as provided by the DescribeOpsItems API action) can view and search on the
	// specified data. Operational data that is not searchable is only viewable by
	// users who have access to the OpsItem (as provided by the GetOpsItem API action).
	// Use the /aws/resources key in OperationalData to specify a related resource in
	// the request. Use the /aws/automations key in OperationalData to associate an
	// Automation runbook with the OpsItem. To view AWS CLI example commands that use
	// these keys, see Creating OpsItems manually
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-creating-OpsItems.html#OpsCenter-manually-create-OpsItems)
	// in the AWS Systems Manager User Guide.
	OperationalData map[string]*types.OpsItemDataValue

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int32

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []*types.RelatedOpsItem

	// Specify a severity to assign to an OpsItem.
	Severity *string

	// Optional metadata that you assign to a resource. You can restrict access to
	// OpsItems by using an inline IAM policy that specifies tags. For more
	// information, see Getting started with OpsCenter
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html#OpsCenter-getting-started-user-permissions)
	// in the AWS Systems Manager User Guide. Tags use a key-value pair. For example:
	// Key=Department,Value=Finance To add tags to an existing OpsItem, use the
	// AddTagsToResource action.
	Tags []*types.Tag
}

type CreateOpsItemOutput struct {

	// The ID of the OpsItem.
	OpsItemId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateOpsItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOpsItem{}, middleware.After)
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
	if err = addOpCreateOpsItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOpsItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOpsItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateOpsItem",
	}
}
