// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a list of job definitions. You can specify a status (such as ACTIVE)
// to only return job definitions that match that status.
func (c *Client) DescribeJobDefinitions(ctx context.Context, params *DescribeJobDefinitionsInput, optFns ...func(*Options)) (*DescribeJobDefinitionsOutput, error) {
	if params == nil {
		params = &DescribeJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobDefinitions", params, optFns, addOperationDescribeJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeJobDefinitions.
type DescribeJobDefinitionsInput struct {

	// The name of the job definition to describe.
	JobDefinitionName *string

	// A list of up to 100 job definition names or full Amazon Resource Name (ARN)
	// entries.
	JobDefinitions []string

	// The maximum number of results returned by DescribeJobDefinitions in paginated
	// output. When this parameter is used, DescribeJobDefinitions only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeJobDefinitions request with the returned nextToken value. This value can
	// be between 1 and 100. If this parameter isn't used, then DescribeJobDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults int32

	// The nextToken value returned from a previous paginated DescribeJobDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return. This token should be treated as an opaque identifier that's only used
	// to retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// The status used to filter job definitions.
	Status *string
}

type DescribeJobDefinitionsOutput struct {

	// The list of job definitions.
	JobDefinitions []types.JobDefinition

	// The nextToken value to include in a future DescribeJobDefinitions request. When
	// the results of a DescribeJobDefinitions request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobDefinitions(options.Region), middleware.Before); err != nil {
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

// DescribeJobDefinitionsAPIClient is a client that implements the
// DescribeJobDefinitions operation.
type DescribeJobDefinitionsAPIClient interface {
	DescribeJobDefinitions(context.Context, *DescribeJobDefinitionsInput, ...func(*Options)) (*DescribeJobDefinitionsOutput, error)
}

var _ DescribeJobDefinitionsAPIClient = (*Client)(nil)

// DescribeJobDefinitionsPaginatorOptions is the paginator options for
// DescribeJobDefinitions
type DescribeJobDefinitionsPaginatorOptions struct {
	// The maximum number of results returned by DescribeJobDefinitions in paginated
	// output. When this parameter is used, DescribeJobDefinitions only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// DescribeJobDefinitions request with the returned nextToken value. This value can
	// be between 1 and 100. If this parameter isn't used, then DescribeJobDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobDefinitionsPaginator is a paginator for DescribeJobDefinitions
type DescribeJobDefinitionsPaginator struct {
	options   DescribeJobDefinitionsPaginatorOptions
	client    DescribeJobDefinitionsAPIClient
	params    *DescribeJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobDefinitionsPaginator returns a new DescribeJobDefinitionsPaginator
func NewDescribeJobDefinitionsPaginator(client DescribeJobDefinitionsAPIClient, params *DescribeJobDefinitionsInput, optFns ...func(*DescribeJobDefinitionsPaginatorOptions)) *DescribeJobDefinitionsPaginator {
	if params == nil {
		params = &DescribeJobDefinitionsInput{}
	}

	options := DescribeJobDefinitionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeJobDefinitions page.
func (p *DescribeJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobDefinitionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeJobDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DescribeJobDefinitions",
	}
}
