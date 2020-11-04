// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation. Returns a list of Amazon Kinesis Analytics
// applications in your account. For each application, the response includes the
// application name, Amazon Resource Name (ARN), and status. If the response
// returns the HasMoreApplications value as true,
//
// you can send another request by
// adding the ExclusiveStartApplicationName in the request body, and set the value
// of this to the last application name from the previous response. If you want
// detailed information about a specific application, use DescribeApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html).
// This operation requires permissions to perform the
// kinesisanalytics:ListApplications action.
func (c *Client) ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) {
	if params == nil {
		params = &ListApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplications", params, optFns, addOperationListApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListApplicationsInput struct {

	// Name of the application to start the list with. When using pagination to
	// retrieve the list, you don't need to specify this parameter in the first
	// request. However, in subsequent requests, you add the last application name from
	// the previous response to get the next page of applications.
	ExclusiveStartApplicationName *string

	// Maximum number of applications to list.
	Limit *int32
}

//
type ListApplicationsOutput struct {

	// List of ApplicationSummary objects.
	//
	// This member is required.
	ApplicationSummaries []*types.ApplicationSummary

	// Returns true if there are more applications to retrieve.
	//
	// This member is required.
	HasMoreApplications *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "ListApplications",
	}
}
