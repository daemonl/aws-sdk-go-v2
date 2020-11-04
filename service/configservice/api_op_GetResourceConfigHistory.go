// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a list of configuration items for the specified resource. The list
// contains details about each state of the resource during the specified time
// interval. If you specified a retention period to retain your ConfigurationItems
// between a minimum of 30 days and a maximum of 7 years (2557 days), AWS Config
// returns the ConfigurationItems for the specified retention period. The response
// is paginated. By default, AWS Config returns a limit of 10 configuration items
// per page. You can customize this number with the limit parameter. The response
// includes a nextToken string. To get the next page of results, run the request
// again and specify the string for the nextToken parameter. Each call to the API
// is limited to span a duration of seven days. It is likely that the number of
// records returned is smaller than the specified limit. In such cases, you can
// make another call, using the nextToken.
func (c *Client) GetResourceConfigHistory(ctx context.Context, params *GetResourceConfigHistoryInput, optFns ...func(*Options)) (*GetResourceConfigHistoryOutput, error) {
	if params == nil {
		params = &GetResourceConfigHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceConfigHistory", params, optFns, addOperationGetResourceConfigHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceConfigHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetResourceConfigHistory action.
type GetResourceConfigHistoryInput struct {

	// The ID of the resource (for example., sg-xxxxxx).
	//
	// This member is required.
	ResourceId *string

	// The resource type.
	//
	// This member is required.
	ResourceType types.ResourceType

	// The chronological order for configuration items listed. By default, the results
	// are listed in reverse chronological order.
	ChronologicalOrder types.ChronologicalOrder

	// The time stamp that indicates an earlier time. If not specified, the action
	// returns paginated results that contain configuration items that start when the
	// first configuration item was recorded.
	EarlierTime *time.Time

	// The time stamp that indicates a later time. If not specified, current time is
	// taken.
	LaterTime *time.Time

	// The maximum number of configuration items returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

// The output for the GetResourceConfigHistory action.
type GetResourceConfigHistoryOutput struct {

	// A list that contains the configuration history of one or more resources.
	ConfigurationItems []*types.ConfigurationItem

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResourceConfigHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourceConfigHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourceConfigHistory{}, middleware.After)
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
	if err = addOpGetResourceConfigHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceConfigHistory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceConfigHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetResourceConfigHistory",
	}
}
