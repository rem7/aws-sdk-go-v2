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

// Returns the CloudFormation stack record created as a result of the create cloud
// formation stack operation. An AWS CloudFormation stack is used to create a new
// Amazon EC2 instance from an exported Lightsail snapshot.
func (c *Client) GetCloudFormationStackRecords(ctx context.Context, params *GetCloudFormationStackRecordsInput, optFns ...func(*Options)) (*GetCloudFormationStackRecordsOutput, error) {
	if params == nil {
		params = &GetCloudFormationStackRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCloudFormationStackRecords", params, optFns, addOperationGetCloudFormationStackRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCloudFormationStackRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCloudFormationStackRecordsInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetClouFormationStackRecords request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetCloudFormationStackRecordsOutput struct {

	// A list of objects describing the CloudFormation stack records.
	CloudFormationStackRecords []types.CloudFormationStackRecord

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetCloudFormationStackRecords request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCloudFormationStackRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCloudFormationStackRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCloudFormationStackRecords{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFormationStackRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCloudFormationStackRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetCloudFormationStackRecords",
	}
}
