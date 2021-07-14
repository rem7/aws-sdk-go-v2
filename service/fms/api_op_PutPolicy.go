// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Firewall Manager policy. Firewall Manager provides the following
// types of policies:
//
// * An WAF policy (type WAFV2), which defines rule groups to
// run first in the corresponding WAF web ACL and rule groups to run last in the
// web ACL.
//
// * An WAF Classic policy (type WAF), which defines a rule group.
//
// * A
// Shield Advanced policy, which applies Shield Advanced protection to specified
// accounts and resources.
//
// * A security group policy, which manages VPC security
// groups across your Amazon Web Services organization.
//
// * An Network Firewall
// policy, which provides firewall rules to filter network traffic in specified
// Amazon VPCs.
//
// * A DNS Firewall policy, which provides Route 53 Resolver DNS
// Firewall rules to filter DNS queries for specified VPCs.
//
// Each policy is
// specific to one of the types. If you want to enforce more than one policy type
// across accounts, create multiple policies. You can create multiple policies for
// each type. You must be subscribed to Shield Advanced to create a Shield Advanced
// policy. For more information about subscribing to Shield Advanced, see
// CreateSubscription
// (https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_CreateSubscription.html).
func (c *Client) PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) {
	if params == nil {
		params = &PutPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPolicy", params, optFns, c.addOperationPutPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPolicyInput struct {

	// The details of the Firewall Manager policy to be created.
	//
	// This member is required.
	Policy *types.Policy

	// The tags to add to the Amazon Web Services resource.
	TagList []types.Tag
}

type PutPolicyOutput struct {

	// The details of the Firewall Manager policy.
	Policy *types.Policy

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPolicy{}, middleware.After)
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
	if err = addOpPutPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "PutPolicy",
	}
}
