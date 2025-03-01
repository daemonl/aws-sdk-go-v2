// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about an execution of your DataSync task. You can use this
// operation to help monitor the progress of an ongoing transfer or check the
// results of the transfer.
func (c *Client) DescribeTaskExecution(ctx context.Context, params *DescribeTaskExecutionInput, optFns ...func(*Options)) (*DescribeTaskExecutionOutput, error) {
	if params == nil {
		params = &DescribeTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTaskExecution", params, optFns, c.addOperationDescribeTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskExecutionRequest
type DescribeTaskExecutionInput struct {

	// Specifies the Amazon Resource Name (ARN) of the task execution that you want
	// information about.
	//
	// This member is required.
	TaskExecutionArn *string

	noSmithyDocumentSerde
}

// DescribeTaskExecutionResponse
type DescribeTaskExecutionOutput struct {

	// The physical number of bytes transferred over the network after compression was
	// applied. In most cases, this number is less than BytesTransferred unless the
	// data isn't compressible.
	BytesCompressed int64

	// The total number of bytes that are involved in the transfer. For the number of
	// bytes sent over the network, see BytesCompressed .
	BytesTransferred int64

	// The number of logical bytes written to the destination location.
	BytesWritten int64

	// The estimated physical number of bytes that will transfer over the network.
	EstimatedBytesToTransfer int64

	// The expected number of files, objects, and directories that DataSync will
	// delete in your destination location. If you don't configure your task (https://docs.aws.amazon.com/datasync/latest/userguide/configure-metadata.html)
	// to delete data in the destination that isn't in the source, the value is always
	// 0 .
	EstimatedFilesToDelete int64

	// The expected number of files, objects, and directories that DataSync will
	// transfer over the network. This value is calculated during the task execution's
	// PREPARING phase before the TRANSFERRING phase. The calculation is based on
	// comparing the content of the source and destination locations and finding the
	// difference that needs to be transferred.
	EstimatedFilesToTransfer int64

	// A list of filter rules that exclude specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Excludes []types.FilterRule

	// The number of files, objects, and directories that DataSync deleted in your
	// destination location. If you don't configure your task (https://docs.aws.amazon.com/datasync/latest/userguide/configure-metadata.html)
	// to delete data in the destination that isn't in the source, the value is always
	// 0 .
	FilesDeleted int64

	// The number of files, objects, and directories that DataSync skipped during your
	// transfer.
	FilesSkipped int64

	// The actual number of files, objects, and directories that DataSync transferred
	// over the network. This value is updated periodically during the task execution's
	// TRANSFERRING phase when something is read from the source and sent over the
	// network. If DataSync fails to transfer something, this value can be less than
	// EstimatedFilesToTransfer . In some cases, this value can also be greater than
	// EstimatedFilesToTransfer . This element is implementation-specific for some
	// location types, so don't use it as an exact indication of what transferred or to
	// monitor your task execution.
	FilesTransferred int64

	// The number of files, objects, and directories that DataSync verified during
	// your transfer.
	FilesVerified int64

	// A list of filter rules that include specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Includes []types.FilterRule

	// Indicates how your transfer task is configured. These options include how
	// DataSync handles files, objects, and their associated metadata during your
	// transfer. You also can specify how to verify data integrity, set bandwidth
	// limits for your task, among other options. Each option has a default value.
	// Unless you need to, you don't have to configure any of these options before
	// starting your task.
	Options *types.Options

	// Indicates whether DataSync generated a complete task report (https://docs.aws.amazon.com/datasync/latest/userguide/creating-task-reports.html)
	// for your transfer.
	ReportResult *types.ReportResult

	// The result of the task execution.
	Result *types.TaskExecutionResultDetail

	// The time when the task execution started.
	StartTime *time.Time

	// The status of the task execution.
	Status types.TaskExecutionStatus

	// The ARN of the task execution that you wanted information about.
	// TaskExecutionArn is hierarchical and includes TaskArn for the task that was
	// executed. For example, a TaskExecution value with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2/execution/exec-08ef1e88ec491019b
	// executed the task with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2 .
	TaskExecutionArn *string

	// The configuration of your task report, which provides detailed information
	// about for your DataSync transfer.
	TaskReportConfig *types.TaskReportConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDescribeTaskExecutionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskExecution(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTaskExecution",
	}
}

type opDescribeTaskExecutionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeTaskExecutionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeTaskExecutionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "datasync"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "datasync"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("datasync")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeTaskExecutionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeTaskExecutionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
