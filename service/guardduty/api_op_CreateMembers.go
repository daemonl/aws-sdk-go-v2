// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates member accounts of the current AWS account by specifying a list of AWS
// account IDs. This step is a prerequisite for managing the associated member
// accounts either by invitation or through an organization. When using Create
// Members as an organizations delegated administrator this action will enable
// GuardDuty in the added member accounts, with the exception of the organization
// master account, which must enable GuardDuty prior to being added as a member. If
// you are adding accounts by invitation use this action after GuardDuty has been
// enabled in potential member accounts and before using Invite Members
// (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html).
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	if params == nil {
		params = &CreateMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembers", params, optFns, addOperationCreateMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// A list of account ID and email address pairs of the accounts that you want to
	// associate with the master GuardDuty account.
	//
	// This member is required.
	AccountDetails []*types.AccountDetail

	// The unique ID of the detector of the GuardDuty account that you want to
	// associate member accounts with.
	//
	// This member is required.
	DetectorId *string
}

type CreateMembersOutput struct {

	// A list of objects that include the accountIds of the unprocessed accounts and a
	// result string that explains why each was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
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
	if err = addOpCreateMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateMembers",
	}
}
