// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists historical, current, and scheduled batch job executions for a specific
// application.
func (c *Client) ListBatchJobExecutions(ctx context.Context, params *ListBatchJobExecutionsInput, optFns ...func(*Options)) (*ListBatchJobExecutionsOutput, error) {
	if params == nil {
		params = &ListBatchJobExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchJobExecutions", params, optFns, c.addOperationListBatchJobExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchJobExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchJobExecutionsInput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier of each batch job execution.
	ExecutionIds []string

	// The name of each batch job execution.
	JobName *string

	// The maximum number of batch job executions to return.
	MaxResults *int32

	// A pagination token to control the number of batch job executions displayed in
	// the list.
	NextToken *string

	// The time after which the batch job executions started.
	StartedAfter *time.Time

	// The time before the batch job executions started.
	StartedBefore *time.Time

	// The status of the batch job executions.
	Status types.BatchJobExecutionStatus

	noSmithyDocumentSerde
}

type ListBatchJobExecutionsOutput struct {

	// Returns a list of batch job executions for an application.
	//
	// This member is required.
	BatchJobExecutions []types.BatchJobExecutionSummary

	// A pagination token that's returned when the response doesn't contain all batch
	// job executions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBatchJobExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBatchJobExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBatchJobExecutions{}, middleware.After)
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
	if err = addOpListBatchJobExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchJobExecutions(options.Region), middleware.Before); err != nil {
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

// ListBatchJobExecutionsAPIClient is a client that implements the
// ListBatchJobExecutions operation.
type ListBatchJobExecutionsAPIClient interface {
	ListBatchJobExecutions(context.Context, *ListBatchJobExecutionsInput, ...func(*Options)) (*ListBatchJobExecutionsOutput, error)
}

var _ ListBatchJobExecutionsAPIClient = (*Client)(nil)

// ListBatchJobExecutionsPaginatorOptions is the paginator options for
// ListBatchJobExecutions
type ListBatchJobExecutionsPaginatorOptions struct {
	// The maximum number of batch job executions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBatchJobExecutionsPaginator is a paginator for ListBatchJobExecutions
type ListBatchJobExecutionsPaginator struct {
	options   ListBatchJobExecutionsPaginatorOptions
	client    ListBatchJobExecutionsAPIClient
	params    *ListBatchJobExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListBatchJobExecutionsPaginator returns a new ListBatchJobExecutionsPaginator
func NewListBatchJobExecutionsPaginator(client ListBatchJobExecutionsAPIClient, params *ListBatchJobExecutionsInput, optFns ...func(*ListBatchJobExecutionsPaginatorOptions)) *ListBatchJobExecutionsPaginator {
	if params == nil {
		params = &ListBatchJobExecutionsInput{}
	}

	options := ListBatchJobExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBatchJobExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBatchJobExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBatchJobExecutions page.
func (p *ListBatchJobExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBatchJobExecutionsOutput, error) {
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

	result, err := p.client.ListBatchJobExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListBatchJobExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "ListBatchJobExecutions",
	}
}
