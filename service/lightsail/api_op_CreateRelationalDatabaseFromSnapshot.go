// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new database from an existing database snapshot in Amazon Lightsail.
// You can create a new database from a snapshot in if something goes wrong with
// your original database, or to change it to a different plan, such as a high
// availability or standard plan. The create relational database from snapshot
// operation supports tag-based access control via request tags and resource tags
// applied to the resource identified by relationalDatabaseSnapshotName. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateRelationalDatabaseFromSnapshot(ctx context.Context, params *CreateRelationalDatabaseFromSnapshotInput, optFns ...func(*Options)) (*CreateRelationalDatabaseFromSnapshotOutput, error) {
	if params == nil {
		params = &CreateRelationalDatabaseFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRelationalDatabaseFromSnapshot", params, optFns, addOperationCreateRelationalDatabaseFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRelationalDatabaseFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRelationalDatabaseFromSnapshotInput struct {

	// The name to use for your new database. Constraints:
	//
	// * Must contain from 2 to
	// 255 alphanumeric characters, or hyphens.
	//
	// * The first and last character must be
	// a letter or number.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The Availability Zone in which to create your new database. Use the us-east-2a
	// case-sensitive format. You can get a list of Availability Zones by using the get
	// regions operation. Be sure to add the include relational database Availability
	// Zones parameter to your request.
	AvailabilityZone *string

	// Specifies the accessibility options for your new database. A value of true
	// specifies a database that is available to resources outside of your Lightsail
	// account. A value of false specifies a database that is available only to your
	// Lightsail resources in the same region as your database.
	PubliclyAccessible *bool

	// The bundle ID for your new database. A bundle describes the performance
	// specifications for your database. You can get a list of database bundle IDs by
	// using the get relational database bundles operation. When creating a new
	// database from a snapshot, you cannot choose a bundle that is smaller than the
	// bundle of the source database.
	RelationalDatabaseBundleId *string

	// The name of the database snapshot from which to create your new database.
	RelationalDatabaseSnapshotName *string

	// The date and time to restore your database from. Constraints:
	//
	// * Must be before
	// the latest restorable time for the database.
	//
	// * Cannot be specified if the use
	// latest restorable time parameter is true.
	//
	// * Specified in Coordinated Universal
	// Time (UTC).
	//
	// * Specified in the Unix time format. For example, if you wish to
	// use a restore time of October 1, 2018, at 8 PM UTC, then you input 1538424000 as
	// the restore time.
	RestoreTime *time.Time

	// The name of the source database.
	SourceRelationalDatabaseName *string

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag

	// Specifies whether your database is restored from the latest backup time. A value
	// of true restores from the latest backup time. Default: false Constraints: Cannot
	// be specified if the restore time parameter is provided.
	UseLatestRestorableTime *bool
}

type CreateRelationalDatabaseFromSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRelationalDatabaseFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRelationalDatabaseFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRelationalDatabaseFromSnapshot{}, middleware.After)
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
	if err = addOpCreateRelationalDatabaseFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRelationalDatabaseFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRelationalDatabaseFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateRelationalDatabaseFromSnapshot",
	}
}
