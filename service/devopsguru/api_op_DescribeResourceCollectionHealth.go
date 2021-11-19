// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of open proactive insights, open reactive insights, and the
// Mean Time to Recover (MTTR) for all closed insights in resource collections in
// your account. You specify the type of Amazon Web Services resources collection.
// The one type of Amazon Web Services resource collection supported is Amazon Web
// Services CloudFormation stacks. DevOps Guru can be configured to analyze only
// the Amazon Web Services resources that are defined in the stacks. You can
// specify up to 500 Amazon Web Services CloudFormation stacks.
func (c *Client) DescribeResourceCollectionHealth(ctx context.Context, params *DescribeResourceCollectionHealthInput, optFns ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error) {
	if params == nil {
		params = &DescribeResourceCollectionHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeResourceCollectionHealth", params, optFns, c.addOperationDescribeResourceCollectionHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeResourceCollectionHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResourceCollectionHealthInput struct {

	// An Amazon Web Services resource collection type. This type specifies how
	// analyzed Amazon Web Services resources are defined. The one type of Amazon Web
	// Services resource collection supported is Amazon Web Services CloudFormation
	// stacks. DevOps Guru can be configured to analyze only the Amazon Web Services
	// resources that are defined in the stacks. You can specify up to 500 Amazon Web
	// Services CloudFormation stacks.
	//
	// This member is required.
	ResourceCollectionType types.ResourceCollectionType

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeResourceCollectionHealthOutput struct {

	// The returned CloudFormationHealthOverview object that contains an
	// InsightHealthOverview object with the requested system health information.
	//
	// This member is required.
	CloudFormation []types.CloudFormationHealth

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// An array of ServiceHealth objects that describes the health of the Amazon Web
	// Services services associated with the resources in the collection.
	Service []types.ServiceHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeResourceCollectionHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeResourceCollectionHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeResourceCollectionHealth{}, middleware.After)
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
	if err = addOpDescribeResourceCollectionHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResourceCollectionHealth(options.Region), middleware.Before); err != nil {
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

// DescribeResourceCollectionHealthAPIClient is a client that implements the
// DescribeResourceCollectionHealth operation.
type DescribeResourceCollectionHealthAPIClient interface {
	DescribeResourceCollectionHealth(context.Context, *DescribeResourceCollectionHealthInput, ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error)
}

var _ DescribeResourceCollectionHealthAPIClient = (*Client)(nil)

// DescribeResourceCollectionHealthPaginatorOptions is the paginator options for
// DescribeResourceCollectionHealth
type DescribeResourceCollectionHealthPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeResourceCollectionHealthPaginator is a paginator for
// DescribeResourceCollectionHealth
type DescribeResourceCollectionHealthPaginator struct {
	options   DescribeResourceCollectionHealthPaginatorOptions
	client    DescribeResourceCollectionHealthAPIClient
	params    *DescribeResourceCollectionHealthInput
	nextToken *string
	firstPage bool
}

// NewDescribeResourceCollectionHealthPaginator returns a new
// DescribeResourceCollectionHealthPaginator
func NewDescribeResourceCollectionHealthPaginator(client DescribeResourceCollectionHealthAPIClient, params *DescribeResourceCollectionHealthInput, optFns ...func(*DescribeResourceCollectionHealthPaginatorOptions)) *DescribeResourceCollectionHealthPaginator {
	if params == nil {
		params = &DescribeResourceCollectionHealthInput{}
	}

	options := DescribeResourceCollectionHealthPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeResourceCollectionHealthPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeResourceCollectionHealthPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeResourceCollectionHealth page.
func (p *DescribeResourceCollectionHealthPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeResourceCollectionHealthOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeResourceCollectionHealth(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeResourceCollectionHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeResourceCollectionHealth",
	}
}
