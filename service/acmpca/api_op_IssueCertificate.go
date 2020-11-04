// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uses your private certificate authority (CA), or one that has been shared with
// you, to issue a client certificate. This action returns the Amazon Resource Name
// (ARN) of the certificate. You can retrieve the certificate by calling the
// GetCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_GetCertificate.html)
// action and specifying the ARN. You cannot use the ACM ListCertificateAuthorities
// action to retrieve the ARNs of the certificates that you issue by using ACM
// Private CA.
func (c *Client) IssueCertificate(ctx context.Context, params *IssueCertificateInput, optFns ...func(*Options)) (*IssueCertificateOutput, error) {
	if params == nil {
		params = &IssueCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IssueCertificate", params, optFns, addOperationIssueCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IssueCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IssueCertificateInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html).
	// This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The certificate signing request (CSR) for the certificate you want to issue. You
	// can use the following OpenSSL command to create the CSR and a 2048 bit RSA
	// private key. openssl req -new -newkey rsa:2048 -days 365 -keyout
	// private/test_cert_priv_key.pem -out csr/test_cert_.csr If you have a
	// configuration file, you can use the following OpenSSL command. The usr_cert
	// block in the configuration file contains your X509 version 3 extensions. openssl
	// req -new -config openssl_rsa.cnf -extensions usr_cert -newkey rsa:2048 -days
	// -365 -keyout private/test_cert_priv_key.pem -out csr/test_cert_.csr Note: A CSR
	// must provide either a subject name or a subject alternative name or the request
	// will be rejected.
	//
	// This member is required.
	Csr []byte

	// The name of the algorithm that will be used to sign the certificate to be
	// issued. This parameter should not be confused with the SigningAlgorithm
	// parameter used to sign a CSR.
	//
	// This member is required.
	SigningAlgorithm types.SigningAlgorithm

	// Information describing the validity period of the certificate. When issuing a
	// certificate, ACM Private CA sets the "Not Before" date in the validity field to
	// date and time minus 60 minutes. This is intended to compensate for time
	// inconsistencies across systems of 60 minutes or less. The validity period
	// configured on a certificate must not exceed the limit set by its parents in the
	// CA hierarchy.
	//
	// This member is required.
	Validity *types.Validity

	// Custom string that can be used to distinguish between calls to the
	// IssueCertificate action. Idempotency tokens time out after one hour. Therefore,
	// if you call IssueCertificate multiple times with the same idempotency token
	// within 5 minutes, ACM Private CA recognizes that you are requesting only one
	// certificate and will issue only one. If you change the idempotency token for
	// each call, PCA recognizes that you are requesting multiple certificates.
	IdempotencyToken *string

	// Specifies a custom configuration template to use when issuing a certificate. If
	// this parameter is not provided, ACM Private CA defaults to the
	// EndEntityCertificate/V1 template. For CA certificates, you should choose the
	// shortest path length that meets your needs. The path length is indicated by the
	// PathLenN portion of the ARN, where N is the CA depth
	// (https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaTerms.html#terms-cadepth).
	// Note: The CA depth configured on a subordinate CA certificate must not exceed
	// the limit set by its parents in the CA hierarchy. The following service-owned
	// TemplateArn values are supported by ACM Private CA:
	//
	// *
	// arn:aws:acm-pca:::template/CodeSigningCertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/CodeSigningCertificate_CSRPassthrough/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityCertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityCertificate_CSRPassthrough/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityClientAuthCertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityClientAuthCertificate_CSRPassthrough/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityServerAuthCertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/EndEntityServerAuthCertificate_CSRPassthrough/V1
	//
	// *
	// arn:aws:acm-pca:::template/OCSPSigningCertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/OCSPSigningCertificate_CSRPassthrough/V1
	//
	// *
	// arn:aws:acm-pca:::template/RootCACertificate/V1
	//
	// *
	// arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen0/V1
	//
	// *
	// arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen1/V1
	//
	// *
	// arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen2/V1
	//
	// *
	// arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen3/V1
	//
	// For more
	// information, see Using Templates
	// (https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html).
	TemplateArn *string
}

type IssueCertificateOutput struct {

	// The Amazon Resource Name (ARN) of the issued certificate and the certificate
	// serial number. This is of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/286535153982981100925020015808220737245
	CertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIssueCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpIssueCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpIssueCertificate{}, middleware.After)
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
	if err = addOpIssueCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIssueCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIssueCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "IssueCertificate",
	}
}
