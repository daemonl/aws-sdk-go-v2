// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a type or type version from active use in the CloudFormation registry.
// If a type or type version is deregistered, it cannot be used in CloudFormation
// operations. To deregister a type, you must individually deregister all
// registered versions of that type. If a type has only a single registered
// version, deregistering that version results in the type itself being
// deregistered. You cannot deregister the default version of a type, unless it is
// the only registered version of that type, in which case the type itself is
// deregistered as well.
func (c *Client) DeregisterType(ctx context.Context, params *DeregisterTypeInput, optFns ...func(*Options)) (*DeregisterTypeOutput, error) {
	if params == nil {
		params = &DeregisterTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterType", params, optFns, addOperationDeregisterTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTypeInput struct {

	// The Amazon Resource Name (ARN) of the type. Conditional: You must specify either
	// TypeName and Type, or Arn.
	Arn *string

	// The kind of type. Currently the only valid value is RESOURCE. Conditional: You
	// must specify either TypeName and Type, or Arn.
	Type types.RegistryType

	// The name of the type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	TypeName *string

	// The ID of a specific version of the type. The version ID is the value at the end
	// of the Amazon Resource Name (ARN) assigned to the type version when it is
	// registered.
	VersionId *string
}

type DeregisterTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeregisterType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeregisterType{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeregisterType",
	}
}
