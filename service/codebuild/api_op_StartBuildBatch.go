// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a batch build for a project.
func (c *Client) StartBuildBatch(ctx context.Context, params *StartBuildBatchInput, optFns ...func(*Options)) (*StartBuildBatchOutput, error) {
	if params == nil {
		params = &StartBuildBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBuildBatch", params, optFns, addOperationStartBuildBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBuildBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBuildBatchInput struct {

	// The name of the project.
	//
	// This member is required.
	ProjectName *string

	// An array of ProjectArtifacts objects that contains information about the build
	// output artifact overrides for the build project.
	ArtifactsOverride *types.ProjectArtifacts

	// A BuildBatchConfigOverride object that contains batch build configuration
	// overrides.
	BuildBatchConfigOverride *types.ProjectBuildBatchConfig

	// Overrides the build timeout specified in the batch build project.
	BuildTimeoutInMinutesOverride *int32

	// A buildspec file declaration that overrides, for this build only, the latest one
	// already defined in the build project. If this value is set, it can be either an
	// inline buildspec definition, the path to an alternate buildspec file relative to
	// the value of the built-in CODEBUILD_SRC_DIR environment variable, or the path to
	// an S3 bucket. The bucket must be in the same AWS Region as the build project.
	// Specify the buildspec file using its ARN (for example,
	// arn:aws:s3:::my-codebuild-sample2/buildspec.yml). If this value is not provided
	// or is set to an empty string, the source code must contain a buildspec file in
	// its root directory. For more information, see Buildspec File Name and Storage
	// Location
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-name-storage).
	BuildspecOverride *string

	// A ProjectCache object that specifies cache overrides.
	CacheOverride *types.ProjectCache

	// The name of a certificate for this batch build that overrides the one specified
	// in the batch build project.
	CertificateOverride *string

	// The name of a compute type for this batch build that overrides the one specified
	// in the batch build project.
	ComputeTypeOverride types.ComputeType

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) that
	// overrides the one specified in the batch build project. The CMK key encrypts the
	// build output artifacts. You can use a cross-account KMS key to encrypt the build
	// output artifacts if your service role has permission to that key. You can
	// specify either the Amazon Resource Name (ARN) of the CMK or, if available, the
	// CMK's alias (using the format alias/).
	EncryptionKeyOverride *string

	// A container type for this batch build that overrides the one specified in the
	// batch build project.
	EnvironmentTypeOverride types.EnvironmentType

	// An array of EnvironmentVariable objects that override, or add to, the
	// environment variables defined in the batch build project.
	EnvironmentVariablesOverride []*types.EnvironmentVariable

	// The user-defined depth of history, with a minimum value of 0, that overrides,
	// for this batch build only, any previous depth of history defined in the batch
	// build project.
	GitCloneDepthOverride *int32

	// A GitSubmodulesConfig object that overrides the Git submodules configuration for
	// this batch build.
	GitSubmodulesConfigOverride *types.GitSubmodulesConfig

	// A unique, case sensitive identifier you provide to ensure the idempotency of the
	// StartBuildBatch request. The token is included in the StartBuildBatch request
	// and is valid for five minutes. If you repeat the StartBuildBatch request with
	// the same token, but change a parameter, AWS CodeBuild returns a parameter
	// mismatch error.
	IdempotencyToken *string

	// The name of an image for this batch build that overrides the one specified in
	// the batch build project.
	ImageOverride *string

	// The type of credentials AWS CodeBuild uses to pull images in your batch build.
	// There are two valid values: CODEBUILD Specifies that AWS CodeBuild uses its own
	// credentials. This requires that you modify your ECR repository policy to trust
	// AWS CodeBuild's service principal. SERVICE_ROLE Specifies that AWS CodeBuild
	// uses your build project's service role. When using a cross-account or private
	// registry image, you must use SERVICE_ROLE credentials. When using an AWS
	// CodeBuild curated image, you must use CODEBUILD credentials.
	ImagePullCredentialsTypeOverride types.ImagePullCredentialsType

	// Enable this flag to override the insecure SSL setting that is specified in the
	// batch build project. The insecure SSL setting determines whether to ignore SSL
	// warnings while connecting to the project source code. This override applies only
	// if the build's source is GitHub Enterprise.
	InsecureSslOverride *bool

	// A LogsConfig object that override the log settings defined in the batch build
	// project.
	LogsConfigOverride *types.LogsConfig

	// Enable this flag to override privileged mode in the batch build project.
	PrivilegedModeOverride *bool

	// The number of minutes a batch build is allowed to be queued before it times out.
	QueuedTimeoutInMinutesOverride *int32

	// A RegistryCredential object that overrides credentials for access to a private
	// registry.
	RegistryCredentialOverride *types.RegistryCredential

	// Set to true to report to your source provider the status of a batch build's
	// start and completion. If you use this option with a source provider other than
	// GitHub, GitHub Enterprise, or Bitbucket, an invalidInputException is thrown. The
	// status of a build triggered by a webhook is always reported to your source
	// provider.
	ReportBuildBatchStatusOverride *bool

	// An array of ProjectArtifacts objects that override the secondary artifacts
	// defined in the batch build project.
	SecondaryArtifactsOverride []*types.ProjectArtifacts

	// An array of ProjectSource objects that override the secondary sources defined in
	// the batch build project.
	SecondarySourcesOverride []*types.ProjectSource

	// An array of ProjectSourceVersion objects that override the secondary source
	// versions in the batch build project.
	SecondarySourcesVersionOverride []*types.ProjectSourceVersion

	// The name of a service role for this batch build that overrides the one specified
	// in the batch build project.
	ServiceRoleOverride *string

	// A SourceAuth object that overrides the one defined in the batch build project.
	// This override applies only if the build project's source is BitBucket or GitHub.
	SourceAuthOverride *types.SourceAuth

	// A location that overrides, for this batch build, the source location defined in
	// the batch build project.
	SourceLocationOverride *string

	// The source input type that overrides the source input defined in the batch build
	// project.
	SourceTypeOverride types.SourceType

	// The version of the batch build input to be built, for this build only. If not
	// specified, the latest version is used. If specified, the contents depends on the
	// source provider: AWS CodeCommit The commit ID, branch, or Git tag to use. GitHub
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the
	// version of the source code you want to build. If a pull request ID is specified,
	// it must use the format pr/pull-request-ID (for example pr/25). If a branch name
	// is specified, the branch's HEAD commit ID is used. If not specified, the default
	// branch's HEAD commit ID is used. Bitbucket The commit ID, branch name, or tag
	// name that corresponds to the version of the source code you want to build. If a
	// branch name is specified, the branch's HEAD commit ID is used. If not specified,
	// the default branch's HEAD commit ID is used. Amazon Simple Storage Service
	// (Amazon S3) The version ID of the object that represents the build input ZIP
	// file to use. If sourceVersion is specified at the project level, then this
	// sourceVersion (at the build level) takes precedence. For more information, see
	// Source Version Sample with CodeBuild
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the AWS CodeBuild User Guide.
	SourceVersion *string
}

type StartBuildBatchOutput struct {

	// A BuildBatch object that contains information about the batch build.
	BuildBatch *types.BuildBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartBuildBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartBuildBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartBuildBatch{}, middleware.After)
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
	if err = addOpStartBuildBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBuildBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBuildBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "StartBuildBatch",
	}
}
