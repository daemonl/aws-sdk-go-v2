// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the recovery point for the specified virtual tape. This operation is
// only supported in the tape gateway type. A recovery point is a point in time
// view of a virtual tape at which all the data on the tape is consistent. If your
// gateway crashes, virtual tapes that have recovery points can be recovered to a
// new gateway. The virtual tape can be retrieved to only one gateway. The
// retrieved tape is read-only. The virtual tape can be retrieved to only a tape
// gateway. There is no charge for retrieving recovery points.
func (c *Client) RetrieveTapeRecoveryPoint(ctx context.Context, params *RetrieveTapeRecoveryPointInput, optFns ...func(*Options)) (*RetrieveTapeRecoveryPointOutput, error) {
	if params == nil {
		params = &RetrieveTapeRecoveryPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetrieveTapeRecoveryPoint", params, optFns, addOperationRetrieveTapeRecoveryPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveTapeRecoveryPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// RetrieveTapeRecoveryPointInput
type RetrieveTapeRecoveryPointInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the virtual tape for which you want to
	// retrieve the recovery point.
	//
	// This member is required.
	TapeARN *string
}

// RetrieveTapeRecoveryPointOutput
type RetrieveTapeRecoveryPointOutput struct {

	// The Amazon Resource Name (ARN) of the virtual tape for which the recovery point
	// was retrieved.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRetrieveTapeRecoveryPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetrieveTapeRecoveryPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetrieveTapeRecoveryPoint{}, middleware.After)
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
	if err = addOpRetrieveTapeRecoveryPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveTapeRecoveryPoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetrieveTapeRecoveryPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "RetrieveTapeRecoveryPoint",
	}
}
