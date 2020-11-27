// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves a sample of messages from the specified channel ingested during the
// specified timeframe. Up to 10 messages can be retrieved.
func (c *Client) SampleChannelData(ctx context.Context, params *SampleChannelDataInput, optFns ...func(*Options)) (*SampleChannelDataOutput, error) {
	if params == nil {
		params = &SampleChannelDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SampleChannelData", params, optFns, addOperationSampleChannelDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SampleChannelDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SampleChannelDataInput struct {

	// The name of the channel whose message samples are retrieved.
	//
	// This member is required.
	ChannelName *string

	// The end of the time window from which sample messages are retrieved.
	EndTime *time.Time

	// The number of sample messages to be retrieved. The limit is 10. The default is
	// also 10.
	MaxMessages *int32

	// The start of the time window from which sample messages are retrieved.
	StartTime *time.Time
}

type SampleChannelDataOutput struct {

	// The list of message samples. Each sample message is returned as a base64-encoded
	// string.
	Payloads [][]byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSampleChannelDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSampleChannelData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSampleChannelData{}, middleware.After)
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
	if err = addOpSampleChannelDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSampleChannelData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSampleChannelData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "SampleChannelData",
	}
}
