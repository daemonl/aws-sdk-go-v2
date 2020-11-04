// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a data source.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// An ID for the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// This member is required.
	DataSourceId *string

	// A display name for the data source.
	//
	// This member is required.
	Name *string

	// The type of the data source. Currently, the supported types for this operation
	// are: ATHENA, AURORA, AURORA_POSTGRESQL, MARIADB, MYSQL, POSTGRESQL, PRESTO,
	// REDSHIFT, S3, SNOWFLAKE, SPARK, SQLSERVER, TERADATA. Use ListDataSources to
	// return a list of all data sources.
	//
	// This member is required.
	Type types.DataSourceType

	// The credentials QuickSight that uses to connect to your underlying source.
	// Currently, only credentials based on user name and password are supported.
	Credentials *types.DataSourceCredentials

	// The parameters that QuickSight uses to connect to your underlying source.
	DataSourceParameters *types.DataSourceParameters

	// A list of resource permissions on the data source.
	Permissions []*types.ResourcePermission

	// Secure Socket Layer (SSL) properties that apply when QuickSight connects to your
	// underlying source.
	SslProperties *types.SslProperties

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the data source.
	Tags []*types.Tag

	// Use this parameter only when you want QuickSight to use a VPC connection when
	// connecting to your underlying source.
	VpcConnectionProperties *types.VpcConnectionProperties
}

type CreateDataSourceOutput struct {

	// The Amazon Resource Name (ARN) of the data source.
	Arn *string

	// The status of creating the data source.
	CreationStatus types.ResourceStatus

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
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
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateDataSource",
	}
}
