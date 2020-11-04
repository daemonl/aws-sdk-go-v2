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

// Lists the resource groups to which a Capacity Reservation has been added.
func (c *Client) GetGroupsForCapacityReservation(ctx context.Context, params *GetGroupsForCapacityReservationInput, optFns ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error) {
	if params == nil {
		params = &GetGroupsForCapacityReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupsForCapacityReservation", params, optFns, addOperationGetGroupsForCapacityReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupsForCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupsForCapacityReservationInput struct {

	// The ID of the Capacity Reservation.
	//
	// This member is required.
	CapacityReservationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string
}

type GetGroupsForCapacityReservationOutput struct {

	// Information about the resource groups to which the Capacity Reservation has been
	// added.
	CapacityReservationGroups []*types.CapacityReservationGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetGroupsForCapacityReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetGroupsForCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetGroupsForCapacityReservation{}, middleware.After)
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
	if err = addOpGetGroupsForCapacityReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetGroupsForCapacityReservation",
	}
}
