// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a DB security group. The specified DB security group must not be
// associated with any DB instances.
func (c *Client) DeleteDBSecurityGroup(ctx context.Context, params *DeleteDBSecurityGroupInput, optFns ...func(*Options)) (*DeleteDBSecurityGroupOutput, error) {
	if params == nil {
		params = &DeleteDBSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBSecurityGroup", params, optFns, addOperationDeleteDBSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteDBSecurityGroupInput struct {

	// The name of the DB security group to delete. You can't delete the default DB
	// security group. Constraints:
	//
	// * Must be 1 to 255 letters, numbers, or
	// hyphens.
	//
	// * First character must be a letter
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens
	//
	// * Must not be "Default"
	//
	// This member is required.
	DBSecurityGroupName *string
}

type DeleteDBSecurityGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDBSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSecurityGroup{}, middleware.After)
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
	if err = addOpDeleteDBSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSecurityGroup",
	}
}
