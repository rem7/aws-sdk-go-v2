// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more Amazon Lightsail SSL/TLS certificates. To
// get a summary of a certificate, ommit includeCertificateDetails from your
// request. The response will include only the certificate Amazon Resource Name
// (ARN), certificate name, domain name, and tags.
func (c *Client) GetCertificates(ctx context.Context, params *GetCertificatesInput, optFns ...func(*Options)) (*GetCertificatesOutput, error) {
	if params == nil {
		params = &GetCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificates", params, optFns, addOperationGetCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificatesInput struct {

	// The name for the certificate for which to return information. When omitted, the
	// response includes all of your certificates in the AWS Region where the request
	// is made.
	CertificateName *string

	// The status of the certificates for which to return information. For example,
	// specify ISSUED to return only certificates with an ISSUED status. When omitted,
	// the response includes all of your certificates in the AWS Region where the
	// request is made, regardless of their current status.
	CertificateStatuses []types.CertificateStatus

	// Indicates whether to include detailed information about the certificates in the
	// response. When omitted, the response includes only the certificate names, Amazon
	// Resource Names (ARNs), domain names, and tags.
	IncludeCertificateDetails bool
}

type GetCertificatesOutput struct {

	// An object that describes certificates.
	Certificates []types.CertificateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificates{}, middleware.After)
	if err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificates(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetCertificates",
	}
}
