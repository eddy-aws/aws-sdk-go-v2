// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of image pipelines.
func (c *Client) ListImagePipelines(ctx context.Context, params *ListImagePipelinesInput, optFns ...func(*Options)) (*ListImagePipelinesOutput, error) {
	if params == nil {
		params = &ListImagePipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImagePipelines", params, optFns, addOperationListImagePipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagePipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagePipelinesInput struct {

	// The filters.
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type ListImagePipelinesOutput struct {

	// The list of image pipelines.
	ImagePipelineList []types.ImagePipeline

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListImagePipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImagePipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImagePipelines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImagePipelines(options.Region), middleware.Before); err != nil {
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

// ListImagePipelinesAPIClient is a client that implements the ListImagePipelines
// operation.
type ListImagePipelinesAPIClient interface {
	ListImagePipelines(context.Context, *ListImagePipelinesInput, ...func(*Options)) (*ListImagePipelinesOutput, error)
}

var _ ListImagePipelinesAPIClient = (*Client)(nil)

// ListImagePipelinesPaginatorOptions is the paginator options for
// ListImagePipelines
type ListImagePipelinesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImagePipelinesPaginator is a paginator for ListImagePipelines
type ListImagePipelinesPaginator struct {
	options   ListImagePipelinesPaginatorOptions
	client    ListImagePipelinesAPIClient
	params    *ListImagePipelinesInput
	nextToken *string
	firstPage bool
}

// NewListImagePipelinesPaginator returns a new ListImagePipelinesPaginator
func NewListImagePipelinesPaginator(client ListImagePipelinesAPIClient, params *ListImagePipelinesInput, optFns ...func(*ListImagePipelinesPaginatorOptions)) *ListImagePipelinesPaginator {
	if params == nil {
		params = &ListImagePipelinesInput{}
	}

	options := ListImagePipelinesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImagePipelinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImagePipelinesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListImagePipelines page.
func (p *ListImagePipelinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImagePipelinesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListImagePipelines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImagePipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImagePipelines",
	}
}
