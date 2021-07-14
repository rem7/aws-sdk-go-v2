// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the Identity and Access Management (IAM) role that is assigned to the
// on-premises instance or virtual machines (VM). IAM roles are first assigned to
// these hybrid instances during the activation process. For more information, see
// CreateActivation.
func (c *Client) UpdateManagedInstanceRole(ctx context.Context, params *UpdateManagedInstanceRoleInput, optFns ...func(*Options)) (*UpdateManagedInstanceRoleOutput, error) {
	if params == nil {
		params = &UpdateManagedInstanceRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateManagedInstanceRole", params, optFns, c.addOperationUpdateManagedInstanceRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateManagedInstanceRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateManagedInstanceRoleInput struct {

	// The IAM role you want to assign or change.
	//
	// This member is required.
	IamRole *string

	// The ID of the managed instance where you want to update the role.
	//
	// This member is required.
	InstanceId *string
}

type UpdateManagedInstanceRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateManagedInstanceRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateManagedInstanceRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateManagedInstanceRole{}, middleware.After)
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
	if err = addOpUpdateManagedInstanceRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateManagedInstanceRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateManagedInstanceRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateManagedInstanceRole",
	}
}
