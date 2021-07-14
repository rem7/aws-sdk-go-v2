// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the default patch baseline for the relevant operating system. To reset
// the Amazon Web Services-predefined patch baseline as the default, specify the
// full patch baseline Amazon Resource Name (ARN) as the baseline ID value. For
// example, for CentOS, specify
// arn:aws:ssm:us-east-2:733109147000:patchbaseline/pb-0574b43a65ea646ed instead of
// pb-0574b43a65ea646ed.
func (c *Client) RegisterDefaultPatchBaseline(ctx context.Context, params *RegisterDefaultPatchBaselineInput, optFns ...func(*Options)) (*RegisterDefaultPatchBaselineOutput, error) {
	if params == nil {
		params = &RegisterDefaultPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDefaultPatchBaseline", params, optFns, c.addOperationRegisterDefaultPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDefaultPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDefaultPatchBaselineInput struct {

	// The ID of the patch baseline that should be the default patch baseline.
	//
	// This member is required.
	BaselineId *string
}

type RegisterDefaultPatchBaselineOutput struct {

	// The ID of the default patch baseline.
	BaselineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRegisterDefaultPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterDefaultPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterDefaultPatchBaseline{}, middleware.After)
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
	if err = addOpRegisterDefaultPatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDefaultPatchBaseline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterDefaultPatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterDefaultPatchBaseline",
	}
}
