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

// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// squash merge strategy. If the merge is successful, it closes the pull request.
func (c *Client) MergePullRequestBySquash(ctx context.Context, params *MergePullRequestBySquashInput, optFns ...func(*Options)) (*MergePullRequestBySquashOutput, error) {
	if params == nil {
		params = &MergePullRequestBySquashInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergePullRequestBySquash", params, optFns, addOperationMergePullRequestBySquashMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergePullRequestBySquashOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MergePullRequestBySquashInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests.
	//
	// This member is required.
	PullRequestId *string

	// The name of the repository where the pull request was created.
	//
	// This member is required.
	RepositoryName *string

	// The name of the author who created the commit. This information is used as both
	// the author and committer for the commit.
	AuthorName *string

	// The commit message to include in the commit information for the merge.
	CommitMessage *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when
	// resolving conflicts during a merge.
	ConflictResolution *types.ConflictResolution

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The email address of the person merging the branches. This information is used
	// in the commit information for the merge.
	Email *string

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If true, a .gitkeep file is created for
	// empty folders. The default is false.
	KeepEmptyFolders *bool

	// The full commit ID of the original or updated commit in the pull request source
	// branch. Pass this value if you want an exception thrown if the current commit ID
	// of the tip of the source branch does not match this commit ID.
	SourceCommitId *string
}

type MergePullRequestBySquashOutput struct {

	// Returns information about a pull request.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMergePullRequestBySquashMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergePullRequestBySquash{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergePullRequestBySquash{}, middleware.After)
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
	if err = addOpMergePullRequestBySquashValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergePullRequestBySquash(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergePullRequestBySquash(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "MergePullRequestBySquash",
	}
}
