// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an IAM instance profile with a running or stopped instance. You
// cannot associate more than one IAM instance profile with an instance.
func (c *Client) AssociateIamInstanceProfile(ctx context.Context, params *AssociateIamInstanceProfileInput, optFns ...func(*Options)) (*AssociateIamInstanceProfileOutput, error) {
	if params == nil {
		params = &AssociateIamInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateIamInstanceProfile", params, optFns, addOperationAssociateIamInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateIamInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateIamInstanceProfileInput struct {

	// The IAM instance profile.
	//
	// This member is required.
	IamInstanceProfile *types.IamInstanceProfileSpecification

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string
}

type AssociateIamInstanceProfileOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateIamInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateIamInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateIamInstanceProfile{}, middleware.After)
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
	if err = addOpAssociateIamInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateIamInstanceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateIamInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateIamInstanceProfile",
	}
}
