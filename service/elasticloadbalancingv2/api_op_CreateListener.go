// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a listener for the specified Application Load Balancer or Network Load
// Balancer. To update a listener, use ModifyListener. When you are finished with a
// listener, you can delete it using DeleteListener. If you are finished with both
// the listener and the load balancer, you can delete them both using
// DeleteLoadBalancer. This operation is idempotent, which means that it completes
// at most one time. If you attempt to create multiple listeners with the same
// settings, each call succeeds. For more information, see Listeners for Your
// Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)
// in the Application Load Balancers Guide and Listeners for Your Network Load
// Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html)
// in the Network Load Balancers Guide.
func (c *Client) CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) {
	if params == nil {
		params = &CreateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateListener", params, optFns, addOperationCreateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateListenerInput struct {

	// The actions for the default rule.
	//
	// This member is required.
	DefaultActions []*types.Action

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	// The port on which the load balancer is listening.
	//
	// This member is required.
	Port *int32

	// The protocol for connections from clients to the load balancer. For Application
	// Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP.
	//
	// This member is required.
	Protocol types.ProtocolEnum

	// [TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN)
	// policy. You can specify one policy name. The following are the possible
	// values:
	//
	// * HTTP1Only
	//
	// * HTTP2Only
	//
	// * HTTP2Optional
	//
	// * HTTP2Preferred
	//
	// *
	// None
	//
	// For more information, see ALPN Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#alpn-policies)
	// in the Network Load Balancers Guide.
	AlpnPolicy []*string

	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN but
	// do not set IsDefault. To create a certificate list for the listener, use
	// AddListenerCertificates.
	Certificates []*types.Certificate

	// [HTTPS and TLS listeners] The security policy that defines which protocols and
	// ciphers are supported. The following are the possible values:
	//
	// *
	// ELBSecurityPolicy-2016-08
	//
	// * ELBSecurityPolicy-TLS-1-0-2015-04
	//
	// *
	// ELBSecurityPolicy-TLS-1-1-2017-01
	//
	// * ELBSecurityPolicy-TLS-1-2-2017-01
	//
	// *
	// ELBSecurityPolicy-TLS-1-2-Ext-2018-06
	//
	// * ELBSecurityPolicy-FS-2018-06
	//
	// *
	// ELBSecurityPolicy-FS-1-1-2019-08
	//
	// * ELBSecurityPolicy-FS-1-2-2019-08
	//
	// *
	// ELBSecurityPolicy-FS-1-2-Res-2019-08
	//
	// For more information, see Security
	// Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
	// in the Application Load Balancers Guide and Security Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies)
	// in the Network Load Balancers Guide.
	SslPolicy *string

	// The tags to assign to the listener.
	Tags []*types.Tag
}

type CreateListenerOutput struct {

	// Information about the listener.
	Listeners []*types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateListener{}, middleware.After)
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
	if err = addOpCreateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateListener",
	}
}
