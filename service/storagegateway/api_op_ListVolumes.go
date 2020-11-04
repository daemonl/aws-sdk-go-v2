// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the iSCSI stored volumes of a gateway. Results are sorted by volume ARN.
// The response includes only the volume ARNs. If you want additional volume
// information, use the DescribeStorediSCSIVolumes or the
// DescribeCachediSCSIVolumes API. The operation supports pagination. By default,
// the operation returns a maximum of up to 100 volumes. You can optionally specify
// the Limit field in the body to limit the number of volumes in the response. If
// the number of volumes returned in the response is truncated, the response
// includes a Marker field. You can use this Marker value in your subsequent
// request to retrieve the next set of volumes. This operation is only supported in
// the cached volume and stored volume gateway types.
func (c *Client) ListVolumes(ctx context.Context, params *ListVolumesInput, optFns ...func(*Options)) (*ListVolumesOutput, error) {
	if params == nil {
		params = &ListVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVolumes", params, optFns, addOperationListVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object that contains one or more of the following fields:
//
// *
// ListVolumesInput$Limit
//
// * ListVolumesInput$Marker
type ListVolumesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Specifies that the list of volumes returned be limited to the specified number
	// of items.
	Limit *int32

	// A string that indicates the position at which to begin the returned list of
	// volumes. Obtain the marker from the response of a previous List iSCSI Volumes
	// request.
	Marker *string
}

// A JSON object containing the following fields:
//
// * ListVolumesOutput$Marker
//
// *
// ListVolumesOutput$VolumeInfos
type ListVolumesOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Use the marker in your next request to continue pagination of iSCSI volumes. If
	// there are no more volumes to list, this field does not appear in the response
	// body.
	Marker *string

	// An array of VolumeInfo objects, where each object describes an iSCSI volume. If
	// no volumes are defined for the gateway, then VolumeInfos is an empty array "[]".
	VolumeInfos []*types.VolumeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumes",
	}
}
