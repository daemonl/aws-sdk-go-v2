// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches users and lists the ones that meet a set of filter and sort criteria.
func (c *Client) SearchUsers(ctx context.Context, params *SearchUsersInput, optFns ...func(*Options)) (*SearchUsersOutput, error) {
	if params == nil {
		params = &SearchUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchUsers", params, optFns, addOperationSearchUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchUsersInput struct {

	// The filters to use for listing a specific set of users. Required. Supported
	// filter keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	Filters []*types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved. Required.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string

	// The sort order to use in listing the filtered set of users. Required. Supported
	// sort keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	SortCriteria []*types.Sort
}

type SearchUsersOutput struct {

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The total number of users returned.
	TotalCount *int32

	// The users that meet the specified set of filter criteria, in sort order.
	Users []*types.UserData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchUsers{}, middleware.After)
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
	if err = addOpSearchUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchUsers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchUsers",
	}
}
