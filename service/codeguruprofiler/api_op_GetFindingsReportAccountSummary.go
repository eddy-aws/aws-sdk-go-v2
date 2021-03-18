// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of FindingsReportSummary
// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html)
// objects that contain analysis results for all profiling groups in your AWS
// account.
func (c *Client) GetFindingsReportAccountSummary(ctx context.Context, params *GetFindingsReportAccountSummaryInput, optFns ...func(*Options)) (*GetFindingsReportAccountSummaryOutput, error) {
	if params == nil {
		params = &GetFindingsReportAccountSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingsReportAccountSummary", params, optFns, addOperationGetFindingsReportAccountSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingsReportAccountSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the GetFindingsReportAccountSummaryRequest.
type GetFindingsReportAccountSummaryInput struct {

	// A Boolean value indicating whether to only return reports from daily profiles.
	// If set to True, only analysis data from daily profiles is returned. If set to
	// False, analysis data is returned from smaller time windows (for example, one
	// hour).
	DailyReportsOnly *bool

	// The maximum number of results returned by  GetFindingsReportAccountSummary in
	// paginated output. When this parameter is used, GetFindingsReportAccountSummary
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetFindingsReportAccountSummary request with the returned nextToken
	// value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// GetFindingsReportAccountSummary request where maxResults was used and the
	// results exceeded the value of that parameter. Pagination continues from the end
	// of the previous results that returned the nextToken value. This token should be
	// treated as an opaque identifier that is only used to retrieve the next items in
	// a list and not for other programmatic purposes.
	NextToken *string
}

// The structure representing the GetFindingsReportAccountSummaryResponse.
type GetFindingsReportAccountSummaryOutput struct {

	// The return list of FindingsReportSummary
	// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html)
	// objects taht contain summaries of analysis results for all profiling groups in
	// your AWS account.
	//
	// This member is required.
	ReportSummaries []types.FindingsReportSummary

	// The nextToken value to include in a future GetFindingsReportAccountSummary
	// request. When the results of a GetFindingsReportAccountSummary request exceed
	// maxResults, this value can be used to retrieve the next page of results. This
	// value is null when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFindingsReportAccountSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingsReportAccountSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingsReportAccountSummary{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingsReportAccountSummary(options.Region), middleware.Before); err != nil {
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

// GetFindingsReportAccountSummaryAPIClient is a client that implements the
// GetFindingsReportAccountSummary operation.
type GetFindingsReportAccountSummaryAPIClient interface {
	GetFindingsReportAccountSummary(context.Context, *GetFindingsReportAccountSummaryInput, ...func(*Options)) (*GetFindingsReportAccountSummaryOutput, error)
}

var _ GetFindingsReportAccountSummaryAPIClient = (*Client)(nil)

// GetFindingsReportAccountSummaryPaginatorOptions is the paginator options for
// GetFindingsReportAccountSummary
type GetFindingsReportAccountSummaryPaginatorOptions struct {
	// The maximum number of results returned by  GetFindingsReportAccountSummary in
	// paginated output. When this parameter is used, GetFindingsReportAccountSummary
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetFindingsReportAccountSummary request with the returned nextToken
	// value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFindingsReportAccountSummaryPaginator is a paginator for
// GetFindingsReportAccountSummary
type GetFindingsReportAccountSummaryPaginator struct {
	options   GetFindingsReportAccountSummaryPaginatorOptions
	client    GetFindingsReportAccountSummaryAPIClient
	params    *GetFindingsReportAccountSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetFindingsReportAccountSummaryPaginator returns a new
// GetFindingsReportAccountSummaryPaginator
func NewGetFindingsReportAccountSummaryPaginator(client GetFindingsReportAccountSummaryAPIClient, params *GetFindingsReportAccountSummaryInput, optFns ...func(*GetFindingsReportAccountSummaryPaginatorOptions)) *GetFindingsReportAccountSummaryPaginator {
	if params == nil {
		params = &GetFindingsReportAccountSummaryInput{}
	}

	options := GetFindingsReportAccountSummaryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFindingsReportAccountSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFindingsReportAccountSummaryPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetFindingsReportAccountSummary page.
func (p *GetFindingsReportAccountSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFindingsReportAccountSummaryOutput, error) {
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

	result, err := p.client.GetFindingsReportAccountSummary(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFindingsReportAccountSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "GetFindingsReportAccountSummary",
	}
}
