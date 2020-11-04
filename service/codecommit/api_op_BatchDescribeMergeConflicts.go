// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more merge conflicts in the attempted merge of
// two commit specifiers using the squash or three-way merge strategy.
func (c *Client) BatchDescribeMergeConflicts(ctx context.Context, params *BatchDescribeMergeConflictsInput, optFns ...func(*Options)) (*BatchDescribeMergeConflictsOutput, error) {
	if params == nil {
		params = &BatchDescribeMergeConflictsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDescribeMergeConflicts", params, optFns, addOperationBatchDescribeMergeConflictsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDescribeMergeConflictsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDescribeMergeConflictsInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The merge option or strategy you want to use to merge the code.
	//
	// This member is required.
	MergeOption types.MergeOptionTypeEnum

	// The name of the repository that contains the merge conflicts you want to review.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The path of the target files used to describe the conflicts. If not specified,
	// the default is all conflict files.
	FilePaths []*string

	// The maximum number of files to include in the output.
	MaxConflictFiles *int32

	// The maximum number of merge hunks to include in the output.
	MaxMergeHunks *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type BatchDescribeMergeConflictsOutput struct {

	// A list of conflicts for each file, including the conflict metadata and the hunks
	// of the differences between the files.
	//
	// This member is required.
	Conflicts []*types.Conflict

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	DestinationCommitId *string

	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	SourceCommitId *string

	// The commit ID of the merge base.
	BaseCommitId *string

	// A list of any errors returned while describing the merge conflicts for each
	// file.
	Errors []*types.BatchDescribeMergeConflictsError

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDescribeMergeConflictsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDescribeMergeConflicts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDescribeMergeConflicts{}, middleware.After)
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
	if err = addOpBatchDescribeMergeConflictsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDescribeMergeConflicts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDescribeMergeConflicts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchDescribeMergeConflicts",
	}
}
