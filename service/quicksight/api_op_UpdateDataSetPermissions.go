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

// Updates the permissions on a dataset. The permissions resource is
// arn:aws:quicksight:region:aws-account-id:dataset/data-set-id.
func (c *Client) UpdateDataSetPermissions(ctx context.Context, params *UpdateDataSetPermissionsInput, optFns ...func(*Options)) (*UpdateDataSetPermissionsOutput, error) {
	if params == nil {
		params = &UpdateDataSetPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSetPermissions", params, optFns, addOperationUpdateDataSetPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSetPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSetPermissionsInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	//
	// This member is required.
	DataSetId *string

	// The resource permissions that you want to grant to the dataset.
	GrantPermissions []*types.ResourcePermission

	// The resource permissions that you want to revoke from the dataset.
	RevokePermissions []*types.ResourcePermission
}

type UpdateDataSetPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	DataSetArn *string

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	DataSetId *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataSetPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSetPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSetPermissions{}, middleware.After)
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
	if err = addOpUpdateDataSetPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSetPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSetPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDataSetPermissions",
	}
}
