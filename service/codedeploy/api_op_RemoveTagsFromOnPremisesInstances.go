// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more tags from one or more on-premises instances.
func (c *Client) RemoveTagsFromOnPremisesInstances(ctx context.Context, params *RemoveTagsFromOnPremisesInstancesInput, optFns ...func(*Options)) (*RemoveTagsFromOnPremisesInstancesOutput, error) {
	if params == nil {
		params = &RemoveTagsFromOnPremisesInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTagsFromOnPremisesInstances", params, optFns, addOperationRemoveTagsFromOnPremisesInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsFromOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RemoveTagsFromOnPremisesInstances operation.
type RemoveTagsFromOnPremisesInstancesInput struct {

	// The names of the on-premises instances from which to remove tags.
	//
	// This member is required.
	InstanceNames []*string

	// The tag key-value pairs to remove from the on-premises instances.
	//
	// This member is required.
	Tags []*types.Tag
}

type RemoveTagsFromOnPremisesInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveTagsFromOnPremisesInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromOnPremisesInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromOnPremisesInstances{}, middleware.After)
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
	if err = addOpRemoveTagsFromOnPremisesInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromOnPremisesInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveTagsFromOnPremisesInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "RemoveTagsFromOnPremisesInstances",
	}
}
