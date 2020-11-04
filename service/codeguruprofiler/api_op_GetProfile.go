// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the aggregated profile of a profiling group for the specified time range.
// If the requested time range does not align with the available aggregated
// profiles, it is expanded to attain alignment. If aggregated profiles are
// available only for part of the period requested, the profile is returned from
// the earliest available to the latest within the requested time range. For
// example, if the requested time range is from 00:00 to 00:20 and the available
// profiles are from 00:15 to 00:25, the returned profile will be from 00:15 to
// 00:20. You must specify exactly two of the following parameters: startTime,
// period, and endTime.
func (c *Client) GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) {
	if params == nil {
		params = &GetProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProfile", params, optFns, addOperationGetProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the getProfileRequest.
type GetProfileInput struct {

	// The name of the profiling group to get.
	//
	// This member is required.
	ProfilingGroupName *string

	// The format of the profile to return. You can choose application/json or the
	// default application/x-amzn-ion.
	Accept *string

	// You must specify exactly two of the following parameters: startTime, period, and
	// endTime.
	EndTime *time.Time

	// The maximum depth of the graph.
	MaxDepth *int32

	// The period of the profile to get. The time range must be in the past and not
	// longer than one week. You must specify exactly two of the following parameters:
	// startTime, period, and endTime.
	Period *string

	// The start time of the profile to get. You must specify exactly two of the
	// following parameters: startTime, period, and endTime.
	StartTime *time.Time
}

// The structure representing the getProfileResponse.
type GetProfileOutput struct {

	// The content type of the profile in the payload. It is either application/json or
	// the default application/x-amzn-ion.
	//
	// This member is required.
	ContentType *string

	// Information about the profile.
	//
	// This member is required.
	Profile []byte

	// The content encoding of the profile.
	ContentEncoding *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProfile{}, middleware.After)
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
	if err = addOpGetProfileValidationMiddleware(stack); err != nil {
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
