// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a vocabulary filter.
func (c *Client) GetVocabularyFilter(ctx context.Context, params *GetVocabularyFilterInput, optFns ...func(*Options)) (*GetVocabularyFilterOutput, error) {
	if params == nil {
		params = &GetVocabularyFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVocabularyFilter", params, optFns, addOperationGetVocabularyFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVocabularyFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVocabularyFilterInput struct {

	// The name of the vocabulary filter for which to return information.
	//
	// This member is required.
	VocabularyFilterName *string
}

type GetVocabularyFilterOutput struct {

	// The URI of the list of words in the vocabulary filter. You can use this URI to
	// get the list of words.
	DownloadUri *string

	// The language code of the words in the vocabulary filter.
	LanguageCode types.LanguageCode

	// The date and time that the contents of the vocabulary filter were updated.
	LastModifiedTime *time.Time

	// The name of the vocabulary filter.
	VocabularyFilterName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVocabularyFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVocabularyFilter{}, middleware.After)
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
	if err = addOpGetVocabularyFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVocabularyFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVocabularyFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetVocabularyFilter",
	}
}
