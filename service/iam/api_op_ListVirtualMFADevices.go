// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the virtual MFA devices defined in the AWS account by assignment status.
// If you do not specify an assignment status, the operation returns a list of all
// virtual MFA devices. Assignment status can be Assigned, Unassigned, or Any. You
// can paginate the results using the MaxItems and Marker parameters.
func (c *Client) ListVirtualMFADevices(ctx context.Context, params *ListVirtualMFADevicesInput, optFns ...func(*Options)) (*ListVirtualMFADevicesOutput, error) {
	if params == nil {
		params = &ListVirtualMFADevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualMFADevices", params, optFns, addOperationListVirtualMFADevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualMFADevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualMFADevicesInput struct {

	// The status (Unassigned or Assigned) of the devices to list. If you do not
	// specify an AssignmentStatus, the operation defaults to Any, which lists both
	// assigned and unassigned virtual MFA devices.,
	AssignmentStatus types.AssignmentStatusType

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32
}

// Contains the response to a successful ListVirtualMFADevices request.
type ListVirtualMFADevicesOutput struct {

	// The list of virtual MFA devices in the current account that match the
	// AssignmentStatus value that was passed in the request.
	//
	// This member is required.
	VirtualMFADevices []*types.VirtualMFADevice

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualMFADevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListVirtualMFADevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListVirtualMFADevices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualMFADevices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVirtualMFADevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListVirtualMFADevices",
	}
}
