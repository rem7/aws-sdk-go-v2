// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Maps users to their groups. You can also map sub groups to groups. For example,
// the group "Company Intellectual Property Teams" includes sub groups "Research"
// and "Engineering". These sub groups include their own list of users or people
// who work in these teams. Only users who work in research and engineering, and
// therefore belong in the intellectual property group, can see top-secret company
// documents in their search results. You map users to their groups when you want
// to filter search results for different users based on their group’s access to
// documents. For more information on filtering search results for different users,
// see Filtering on user context
// (https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html). If more
// than five PUT actions for a group are currently processing, a validation
// exception is thrown.
func (c *Client) PutPrincipalMapping(ctx context.Context, params *PutPrincipalMappingInput, optFns ...func(*Options)) (*PutPrincipalMappingOutput, error) {
	if params == nil {
		params = &PutPrincipalMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPrincipalMapping", params, optFns, c.addOperationPutPrincipalMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPrincipalMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPrincipalMappingInput struct {

	// The identifier of the group you want to map its users to.
	//
	// This member is required.
	GroupId *string

	// The list that contains your users or sub groups that belong the same group. For
	// example, the group "Company" includes the user "CEO" and the sub groups
	// "Research", "Engineering", and "Sales and Marketing". If you have more than 1000
	// users and/or sub groups for a single group, you need to provide the path to the
	// S3 file that lists your users and sub groups for a group. Your sub groups can
	// contain more than 1000 users, but the list of sub groups that belong to a group
	// (and/or users) must be no more than 1000.
	//
	// This member is required.
	GroupMembers *types.GroupMembers

	// The identifier of the index you want to map users to their groups.
	//
	// This member is required.
	IndexId *string

	// The identifier of the data source you want to map users to their groups. This is
	// useful if a group is tied to multiple data sources, but you only want the group
	// to access documents of a certain data source. For example, the groups
	// "Research", "Engineering", and "Sales and Marketing" are all tied to the
	// company's documents stored in the data sources Confluence and Salesforce.
	// However, "Sales and Marketing" team only needs access to customer-related
	// documents stored in Salesforce.
	DataSourceId *string

	// The timestamp identifier you specify to ensure Amazon Kendra does not override
	// the latest PUT action with previous actions. The highest number ID, which is the
	// ordering ID, is the latest action you want to process and apply on top of other
	// actions with lower number IDs. This prevents previous actions with lower number
	// IDs from possibly overriding the latest action. The ordering ID can be the UNIX
	// time of the last update you made to a group members list. You would then provide
	// this list when calling PutPrincipalMapping. This ensures your PUT action for
	// that updated group with the latest members list doesn't get overwritten by
	// earlier PUT actions for the same group which are yet to be processed. The
	// default ordering ID is the current UNIX time in milliseconds that the action was
	// received by Amazon Kendra.
	OrderingId *int64

	// The Amazon Resource Name (ARN) of a role that has access to the S3 file that
	// contains your list of users or sub groups that belong to a group. For more
	// information, see IAM roles for Amazon Kendra
	// (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html#iam-roles-ds).
	RoleArn *string
}

type PutPrincipalMappingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutPrincipalMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPrincipalMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPrincipalMapping{}, middleware.After)
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
	if err = addOpPutPrincipalMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPrincipalMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPrincipalMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "PutPrincipalMapping",
	}
}
