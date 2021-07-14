// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the Amazon Lightsail resources that can access the specified Lightsail
// bucket. Lightsail buckets currently support setting access for Lightsail
// instances in the same AWS Region.
func (c *Client) SetResourceAccessForBucket(ctx context.Context, params *SetResourceAccessForBucketInput, optFns ...func(*Options)) (*SetResourceAccessForBucketOutput, error) {
	if params == nil {
		params = &SetResourceAccessForBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetResourceAccessForBucket", params, optFns, c.addOperationSetResourceAccessForBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetResourceAccessForBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetResourceAccessForBucketInput struct {

	// The access setting. The following access settings are available:
	//
	// * allow -
	// Allows access to the bucket and its objects.
	//
	// * deny - Denies access to the
	// bucket and its objects. Use this setting to remove access for a resource
	// previously set to allow.
	//
	// This member is required.
	Access types.ResourceBucketAccess

	// The name of the bucket for which to set access to another Lightsail resource.
	//
	// This member is required.
	BucketName *string

	// The name of the Lightsail instance for which to set bucket access. The instance
	// must be in a running or stopped state.
	//
	// This member is required.
	ResourceName *string
}

type SetResourceAccessForBucketOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationSetResourceAccessForBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetResourceAccessForBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetResourceAccessForBucket{}, middleware.After)
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
	if err = addOpSetResourceAccessForBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetResourceAccessForBucket(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetResourceAccessForBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "SetResourceAccessForBucket",
	}
}
