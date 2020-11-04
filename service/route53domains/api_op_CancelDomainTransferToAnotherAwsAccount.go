// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels the transfer of a domain from the current AWS account to another AWS
// account. You initiate a transfer between AWS accounts using
// TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
// You must cancel the transfer before the other AWS account accepts the transfer
// using AcceptDomainTransferFromAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html).
// Use either ListOperations
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html)
// or GetOperationDetail
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to determine whether the operation succeeded. GetOperationDetail
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// provides additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled.
func (c *Client) CancelDomainTransferToAnotherAwsAccount(ctx context.Context, params *CancelDomainTransferToAnotherAwsAccountInput, optFns ...func(*Options)) (*CancelDomainTransferToAnotherAwsAccountOutput, error) {
	if params == nil {
		params = &CancelDomainTransferToAnotherAwsAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelDomainTransferToAnotherAwsAccount", params, optFns, addOperationCancelDomainTransferToAnotherAwsAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelDomainTransferToAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CancelDomainTransferToAnotherAwsAccount request includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountInput struct {

	// The name of the domain for which you want to cancel the transfer to another AWS
	// account.
	//
	// This member is required.
	DomainName *string
}

// The CancelDomainTransferToAnotherAwsAccount response includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountOutput struct {

	// The identifier that TransferDomainToAnotherAwsAccount returned to track the
	// progress of the request. Because the transfer request was canceled, the value is
	// no longer valid, and you can't use GetOperationDetail to query the operation
	// status.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelDomainTransferToAnotherAwsAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelDomainTransferToAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelDomainTransferToAnotherAwsAccount{}, middleware.After)
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
	if err = addOpCancelDomainTransferToAnotherAwsAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelDomainTransferToAnotherAwsAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelDomainTransferToAnotherAwsAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "CancelDomainTransferToAnotherAwsAccount",
	}
}
