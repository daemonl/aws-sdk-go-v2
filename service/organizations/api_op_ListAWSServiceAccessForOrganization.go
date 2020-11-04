// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the AWS services that you enabled to integrate with your
// organization. After a service on this list creates the resources that it
// requires for the integration, it can perform operations on your organization and
// its accounts. For more information about integrating other services with AWS
// Organizations, including the list of services that currently work with
// Organizations, see Integrating AWS Organizations with Other AWS Services
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) ListAWSServiceAccessForOrganization(ctx context.Context, params *ListAWSServiceAccessForOrganizationInput, optFns ...func(*Options)) (*ListAWSServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &ListAWSServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAWSServiceAccessForOrganization", params, optFns, addOperationListAWSServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAWSServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAWSServiceAccessForOrganizationInput struct {

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListAWSServiceAccessForOrganizationOutput struct {

	// A list of the service principals for the services that are enabled to integrate
	// with your organization. Each principal is a structure that includes the name and
	// the date that it was enabled for integration with AWS Organizations.
	EnabledServicePrincipals []*types.EnabledServicePrincipal

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAWSServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAWSServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAWSServiceAccessForOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAWSServiceAccessForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAWSServiceAccessForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListAWSServiceAccessForOrganization",
	}
}
