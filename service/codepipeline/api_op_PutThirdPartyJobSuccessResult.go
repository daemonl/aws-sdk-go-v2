// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents the success of a third party job as returned to the pipeline by a job
// worker. Used for partner actions only.
func (c *Client) PutThirdPartyJobSuccessResult(ctx context.Context, params *PutThirdPartyJobSuccessResultInput, optFns ...func(*Options)) (*PutThirdPartyJobSuccessResultOutput, error) {
	if params == nil {
		params = &PutThirdPartyJobSuccessResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutThirdPartyJobSuccessResult", params, optFns, addOperationPutThirdPartyJobSuccessResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutThirdPartyJobSuccessResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutThirdPartyJobSuccessResult action.
type PutThirdPartyJobSuccessResultInput struct {

	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	//
	// This member is required.
	ClientToken *string

	// The ID of the job that successfully completed. This is the same ID returned from
	// PollForThirdPartyJobs.
	//
	// This member is required.
	JobId *string

	// A token generated by a job worker, such as an AWS CodeDeploy deployment ID, that
	// a successful job provides to identify a partner action in progress. Future jobs
	// use this token to identify the running instance of the action. It can be reused
	// to return more information about the progress of the partner action. When the
	// action is complete, no continuation token should be supplied.
	ContinuationToken *string

	// Represents information about a current revision.
	CurrentRevision *types.CurrentRevision

	// The details of the actions taken and results produced on an artifact as it
	// passes through stages in the pipeline.
	ExecutionDetails *types.ExecutionDetails
}

type PutThirdPartyJobSuccessResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutThirdPartyJobSuccessResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutThirdPartyJobSuccessResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutThirdPartyJobSuccessResult{}, middleware.After)
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
	if err = addOpPutThirdPartyJobSuccessResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutThirdPartyJobSuccessResult(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutThirdPartyJobSuccessResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutThirdPartyJobSuccessResult",
	}
}
