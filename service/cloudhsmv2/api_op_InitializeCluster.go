// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Claims an AWS CloudHSM cluster by submitting the cluster certificate issued by
// your issuing certificate authority (CA) and the CA's root certificate. Before
// you can claim a cluster, you must sign the cluster's certificate signing request
// (CSR) with your issuing CA. To get the cluster's CSR, use DescribeClusters.
func (c *Client) InitializeCluster(ctx context.Context, params *InitializeClusterInput, optFns ...func(*Options)) (*InitializeClusterOutput, error) {
	if params == nil {
		params = &InitializeClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitializeCluster", params, optFns, addOperationInitializeClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InitializeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InitializeClusterInput struct {

	// The identifier (ID) of the cluster that you are claiming. To find the cluster
	// ID, use DescribeClusters.
	//
	// This member is required.
	ClusterId *string

	// The cluster certificate issued (signed) by your issuing certificate authority
	// (CA). The certificate must be in PEM format and can contain a maximum of 5000
	// characters.
	//
	// This member is required.
	SignedCert *string

	// The issuing certificate of the issuing certificate authority (CA) that issued
	// (signed) the cluster certificate. You must use a self-signed certificate. The
	// certificate used to sign the HSM CSR must be directly available, and thus must
	// be the root certificate. The certificate must be in PEM format and can contain a
	// maximum of 5000 characters.
	//
	// This member is required.
	TrustAnchor *string
}

type InitializeClusterOutput struct {

	// The cluster's state.
	State types.ClusterState

	// A description of the cluster's state.
	StateMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInitializeClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInitializeCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInitializeCluster{}, middleware.After)
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
	if err = addOpInitializeClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitializeCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInitializeCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "InitializeCluster",
	}
}
