// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates an alias for the specified version of the bot or replaces an alias for
// the specified bot. To change the version of the bot that the alias points to,
// replace the alias. For more information about aliases, see versioning-aliases.
// This operation requires permissions for the lex:PutBotAlias action.
func (c *Client) PutBotAlias(ctx context.Context, params *PutBotAliasInput, optFns ...func(*Options)) (*PutBotAliasOutput, error) {
	if params == nil {
		params = &PutBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBotAlias", params, optFns, addOperationPutBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBotAliasInput struct {

	// The name of the bot.
	//
	// This member is required.
	BotName *string

	// The version of the bot.
	//
	// This member is required.
	BotVersion *string

	// The name of the alias. The name is not case sensitive.
	//
	// This member is required.
	Name *string

	// Identifies a specific revision of the $LATEST version. When you create a new bot
	// alias, leave the checksum field blank. If you specify a checksum you get a
	// BadRequestException exception. When you want to update a bot alias, set the
	// checksum field to the checksum of the most recent revision of the $LATEST
	// version. If you don't specify the  checksum field, or if the checksum does not
	// match the $LATEST version, you get a PreconditionFailedException exception.
	Checksum *string

	// Settings for conversation logs for the alias.
	ConversationLogs *types.ConversationLogsRequest

	// A description of the alias.
	Description *string

	// A list of tags to add to the bot alias. You can only add tags when you create an
	// alias, you can't use the PutBotAlias operation to update the tags on a bot
	// alias. To update tags, use the TagResource operation.
	Tags []*types.Tag
}

type PutBotAliasOutput struct {

	// The name of the bot that the alias points to.
	BotName *string

	// The version of the bot that the alias points to.
	BotVersion *string

	// The checksum for the current version of the alias.
	Checksum *string

	// The settings that determine how Amazon Lex uses conversation logs for the alias.
	ConversationLogs *types.ConversationLogsResponse

	// The date that the bot alias was created.
	CreatedDate *time.Time

	// A description of the alias.
	Description *string

	// The date that the bot alias was updated. When you create a resource, the
	// creation date and the last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the alias.
	Name *string

	// A list of tags associated with a bot.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBotAlias{}, middleware.After)
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
	if err = addOpPutBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBotAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "PutBotAlias",
	}
}
