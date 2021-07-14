// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the SMS sandbox status for the calling account in the target Region.
// When you start using Amazon SNS to send SMS messages, your account is in the SMS
// sandbox. The SMS sandbox provides a safe environment for you to try Amazon SNS
// features without risking your reputation as an SMS sender. While your account is
// in the SMS sandbox, you can use all of the features of Amazon SNS. However, you
// can send SMS messages only to verified destination phone numbers. For more
// information, including how to move out of the sandbox to send messages without
// restrictions, see SMS sandbox
// (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html) in the Amazon
// SNS Developer Guide.
func (c *Client) GetSMSSandboxAccountStatus(ctx context.Context, params *GetSMSSandboxAccountStatusInput, optFns ...func(*Options)) (*GetSMSSandboxAccountStatusOutput, error) {
	if params == nil {
		params = &GetSMSSandboxAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSMSSandboxAccountStatus", params, optFns, c.addOperationGetSMSSandboxAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSMSSandboxAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSMSSandboxAccountStatusInput struct {
}

type GetSMSSandboxAccountStatusOutput struct {

	// Indicates whether the calling account is in the SMS sandbox.
	//
	// This member is required.
	IsInSandbox bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetSMSSandboxAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetSMSSandboxAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetSMSSandboxAccountStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSMSSandboxAccountStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSMSSandboxAccountStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "GetSMSSandboxAccountStatus",
	}
}
