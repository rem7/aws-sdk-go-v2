// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a description of a Server Message Block (SMB) file share settings from a
// file gateway. This operation is only supported for file gateways.
func (c *Client) DescribeSMBSettings(ctx context.Context, params *DescribeSMBSettingsInput, optFns ...func(*Options)) (*DescribeSMBSettingsOutput, error) {
	if params == nil {
		params = &DescribeSMBSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSMBSettings", params, optFns, c.addOperationDescribeSMBSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSMBSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSMBSettingsInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Region.
	//
	// This member is required.
	GatewayARN *string
}

type DescribeSMBSettingsOutput struct {

	// Indicates the status of a gateway that is a member of the Active Directory
	// domain.
	//
	// * ACCESS_DENIED: Indicates that the JoinDomain operation failed due to
	// an authentication error.
	//
	// * DETACHED: Indicates that gateway is not joined to a
	// domain.
	//
	// * JOINED: Indicates that the gateway has successfully joined a
	// domain.
	//
	// * JOINING: Indicates that a JoinDomain operation is in progress.
	//
	// *
	// NETWORK_ERROR: Indicates that JoinDomain operation failed due to a network or
	// connectivity error.
	//
	// * TIMEOUT: Indicates that the JoinDomain operation failed
	// because the operation didn't complete within the allotted time.
	//
	// *
	// UNKNOWN_ERROR: Indicates that the JoinDomain operation failed due to another
	// type of error.
	ActiveDirectoryStatus types.ActiveDirectoryStatus

	// The name of the domain that the gateway is joined to.
	DomainName *string

	// The shares on this gateway appear when listing shares. Only supported for S3
	// File Gateways.
	FileSharesVisible *bool

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Region.
	GatewayARN *string

	// This value is true if a password for the guest user smbguest is set, otherwise
	// false. Only supported for S3 File Gateways. Valid Values: true | false
	SMBGuestPasswordSet *bool

	// The type of security strategy that was specified for file gateway.
	//
	// *
	// ClientSpecified: If you use this option, requests are established based on what
	// is negotiated by the client. This option is recommended when you want to
	// maximize compatibility across different clients in your environment. Only
	// supported for S3 File Gateways.
	//
	// * MandatorySigning: If you use this option,
	// file gateway only allows connections from SMBv2 or SMBv3 clients that have
	// signing enabled. This option works with SMB clients on Microsoft Windows Vista,
	// Windows Server 2008 or newer.
	//
	// * MandatoryEncryption: If you use this option,
	// file gateway only allows connections from SMBv3 clients that have encryption
	// enabled. This option is highly recommended for environments that handle
	// sensitive data. This option works with SMB clients on Microsoft Windows 8,
	// Windows Server 2012 or newer.
	SMBSecurityStrategy types.SMBSecurityStrategy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeSMBSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSMBSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSMBSettings{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpDescribeSMBSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSMBSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSMBSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeSMBSettings",
	}
}
