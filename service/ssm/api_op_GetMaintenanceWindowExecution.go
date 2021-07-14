// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about a specific a maintenance window execution.
func (c *Client) GetMaintenanceWindowExecution(ctx context.Context, params *GetMaintenanceWindowExecutionInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecution", params, optFns, c.addOperationGetMaintenanceWindowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionInput struct {

	// The ID of the maintenance window execution that includes the task.
	//
	// This member is required.
	WindowExecutionId *string
}

type GetMaintenanceWindowExecutionOutput struct {

	// The time the maintenance window finished running.
	EndTime *time.Time

	// The time the maintenance window started running.
	StartTime *time.Time

	// The status of the maintenance window execution.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the status. Not available for all status values.
	StatusDetails *string

	// The ID of the task executions from the maintenance window execution.
	TaskIds []string

	// The ID of the maintenance window execution.
	WindowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetMaintenanceWindowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecution{}, middleware.After)
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
	if err = addOpGetMaintenanceWindowExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetMaintenanceWindowExecution",
	}
}
