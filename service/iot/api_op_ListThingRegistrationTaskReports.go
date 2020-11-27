// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Information about the thing registration tasks.
func (c *Client) ListThingRegistrationTaskReports(ctx context.Context, params *ListThingRegistrationTaskReportsInput, optFns ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error) {
	if params == nil {
		params = &ListThingRegistrationTaskReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingRegistrationTaskReports", params, optFns, addOperationListThingRegistrationTaskReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingRegistrationTaskReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingRegistrationTaskReportsInput struct {

	// The type of task report.
	//
	// This member is required.
	ReportType types.ReportType

	// The id of the task.
	//
	// This member is required.
	TaskId *string

	// The maximum number of results to return per request.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string
}

type ListThingRegistrationTaskReportsOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The type of task report.
	ReportType types.ReportType

	// Links to the task resources.
	ResourceLinks []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingRegistrationTaskReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingRegistrationTaskReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingRegistrationTaskReports{}, middleware.After)
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
	if err = addOpListThingRegistrationTaskReportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingRegistrationTaskReports(options.Region), middleware.Before); err != nil {
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

// ListThingRegistrationTaskReportsAPIClient is a client that implements the
// ListThingRegistrationTaskReports operation.
type ListThingRegistrationTaskReportsAPIClient interface {
	ListThingRegistrationTaskReports(context.Context, *ListThingRegistrationTaskReportsInput, ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error)
}

var _ ListThingRegistrationTaskReportsAPIClient = (*Client)(nil)

// ListThingRegistrationTaskReportsPaginatorOptions is the paginator options for
// ListThingRegistrationTaskReports
type ListThingRegistrationTaskReportsPaginatorOptions struct {
	// The maximum number of results to return per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingRegistrationTaskReportsPaginator is a paginator for
// ListThingRegistrationTaskReports
type ListThingRegistrationTaskReportsPaginator struct {
	options   ListThingRegistrationTaskReportsPaginatorOptions
	client    ListThingRegistrationTaskReportsAPIClient
	params    *ListThingRegistrationTaskReportsInput
	nextToken *string
	firstPage bool
}

// NewListThingRegistrationTaskReportsPaginator returns a new
// ListThingRegistrationTaskReportsPaginator
func NewListThingRegistrationTaskReportsPaginator(client ListThingRegistrationTaskReportsAPIClient, params *ListThingRegistrationTaskReportsInput, optFns ...func(*ListThingRegistrationTaskReportsPaginatorOptions)) *ListThingRegistrationTaskReportsPaginator {
	options := ListThingRegistrationTaskReportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListThingRegistrationTaskReportsInput{}
	}

	return &ListThingRegistrationTaskReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingRegistrationTaskReportsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListThingRegistrationTaskReports page.
func (p *ListThingRegistrationTaskReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListThingRegistrationTaskReports(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListThingRegistrationTaskReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingRegistrationTaskReports",
	}
}
