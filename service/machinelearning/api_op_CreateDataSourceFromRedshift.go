// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a DataSource from a database hosted on an Amazon Redshift cluster. A
// DataSource references data that can be used to perform either CreateMLModel,
// CreateEvaluation, or CreateBatchPrediction operations.
// CreateDataSourceFromRedshift is an asynchronous operation. In response to
// CreateDataSourceFromRedshift, Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING. After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED.
// DataSource in COMPLETED or PENDING states can be used to perform only
// CreateMLModel, CreateEvaluation, or CreateBatchPrediction operations. If Amazon
// ML can't accept the input source, it sets the Status parameter to FAILED and
// includes an error message in the Message attribute of the GetDataSource
// operation response. The observations should be contained in the database hosted
// on an Amazon Redshift cluster and should be specified by a SelectSqlQuery query.
// Amazon ML executes an Unload command in Amazon Redshift to transfer the result
// set of the SelectSqlQuery query to S3StagingLocation. After the DataSource has
// been created, it's ready for use in evaluations and batch predictions. If you
// plan to use the DataSource to train an MLModel, the DataSource also requires a
// recipe. A recipe describes how each input variable will be used in training an
// MLModel. Will the variable be included or excluded from training? Will the
// variable be manipulated; for example, will it be combined with another variable
// or will it be split apart into word combinations? The recipe provides answers to
// these questions. You can't change an existing datasource, but you can copy and
// modify the settings from an existing Amazon Redshift datasource to create a new
// datasource. To do so, call GetDataSource for an existing datasource and copy the
// values to a CreateDataSource call. Change the settings that you want to change
// and make sure that all required fields have the appropriate values.
func (c *Client) CreateDataSourceFromRedshift(ctx context.Context, params *CreateDataSourceFromRedshiftInput, optFns ...func(*Options)) (*CreateDataSourceFromRedshiftOutput, error) {
	if params == nil {
		params = &CreateDataSourceFromRedshiftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSourceFromRedshift", params, optFns, addOperationCreateDataSourceFromRedshiftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceFromRedshiftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceFromRedshiftInput struct {

	// A user-supplied ID that uniquely identifies the DataSource.
	//
	// This member is required.
	DataSourceId *string

	// The data specification of an Amazon Redshift DataSource:
	//
	// * DatabaseInformation
	// -
	//
	// * DatabaseName - The name of the Amazon Redshift database.
	//
	// *
	// ClusterIdentifier - The unique ID for the Amazon Redshift cluster.
	//
	// *
	// DatabaseCredentials - The AWS Identity and Access Management (IAM) credentials
	// that are used to connect to the Amazon Redshift database.
	//
	// * SelectSqlQuery -
	// The query that is used to retrieve the observation data for the Datasource.
	//
	// *
	// S3StagingLocation - The Amazon Simple Storage Service (Amazon S3) location for
	// staging Amazon Redshift data. The data retrieved from Amazon Redshift using the
	// SelectSqlQuery query is stored in this location.
	//
	// * DataSchemaUri - The Amazon
	// S3 location of the DataSchema.
	//
	// * DataSchema - A JSON string representing the
	// schema. This is not required if DataSchemaUri is specified.
	//
	// * DataRearrangement
	// - A JSON string that represents the splitting and rearrangement requirements for
	// the DataSource. Sample -
	// "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// This member is required.
	DataSpec *types.RedshiftDataSpec

	// A fully specified role Amazon Resource Name (ARN). Amazon ML assumes the role on
	// behalf of the user to create the following:
	//
	// * A security group to allow Amazon
	// ML to execute the SelectSqlQuery query on an Amazon Redshift cluster
	//
	// * An
	// Amazon S3 bucket policy to grant Amazon ML read/write permissions on the
	// S3StagingLocation
	//
	// This member is required.
	RoleARN *string

	// The compute statistics for a DataSource. The statistics are generated from the
	// observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if the
	// DataSource needs to be used for MLModel training.
	ComputeStatistics *bool

	// A user-supplied name or description of the DataSource.
	DataSourceName *string
}

// Represents the output of a CreateDataSourceFromRedshift operation, and is an
// acknowledgement that Amazon ML received the request. The
// CreateDataSourceFromRedshift operation is asynchronous. You can poll for updates
// by using the GetBatchPrediction operation and checking the Status parameter.
type CreateDataSourceFromRedshiftOutput struct {

	// A user-supplied ID that uniquely identifies the datasource. This value should be
	// identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDataSourceFromRedshiftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSourceFromRedshift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSourceFromRedshift{}, middleware.After)
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
	if err = addOpCreateDataSourceFromRedshiftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSourceFromRedshift(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataSourceFromRedshift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateDataSourceFromRedshift",
	}
}
