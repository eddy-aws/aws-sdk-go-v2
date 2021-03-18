// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all FHIR Data Stores that are in the user’s account, regardless of Data
// Store status.
func (c *Client) ListFHIRDatastores(ctx context.Context, params *ListFHIRDatastoresInput, optFns ...func(*Options)) (*ListFHIRDatastoresOutput, error) {
	if params == nil {
		params = &ListFHIRDatastoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFHIRDatastores", params, optFns, addOperationListFHIRDatastoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFHIRDatastoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFHIRDatastoresInput struct {

	// Lists all filters associated with a FHIR Data Store request.
	Filter *types.DatastoreFilter

	// The maximum number of Data Stores returned in a single page of a
	// ListFHIRDatastoresRequest call.
	MaxResults *int32

	// Fetches the next page of Data Stores when results are paginated.
	NextToken *string
}

type ListFHIRDatastoresOutput struct {

	// All properties associated with the listed Data Stores.
	//
	// This member is required.
	DatastorePropertiesList []types.DatastoreProperties

	// Pagination token that can be used to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFHIRDatastoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFHIRDatastores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFHIRDatastores{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFHIRDatastores(options.Region), middleware.Before); err != nil {
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

// ListFHIRDatastoresAPIClient is a client that implements the ListFHIRDatastores
// operation.
type ListFHIRDatastoresAPIClient interface {
	ListFHIRDatastores(context.Context, *ListFHIRDatastoresInput, ...func(*Options)) (*ListFHIRDatastoresOutput, error)
}

var _ ListFHIRDatastoresAPIClient = (*Client)(nil)

// ListFHIRDatastoresPaginatorOptions is the paginator options for
// ListFHIRDatastores
type ListFHIRDatastoresPaginatorOptions struct {
	// The maximum number of Data Stores returned in a single page of a
	// ListFHIRDatastoresRequest call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFHIRDatastoresPaginator is a paginator for ListFHIRDatastores
type ListFHIRDatastoresPaginator struct {
	options   ListFHIRDatastoresPaginatorOptions
	client    ListFHIRDatastoresAPIClient
	params    *ListFHIRDatastoresInput
	nextToken *string
	firstPage bool
}

// NewListFHIRDatastoresPaginator returns a new ListFHIRDatastoresPaginator
func NewListFHIRDatastoresPaginator(client ListFHIRDatastoresAPIClient, params *ListFHIRDatastoresInput, optFns ...func(*ListFHIRDatastoresPaginatorOptions)) *ListFHIRDatastoresPaginator {
	if params == nil {
		params = &ListFHIRDatastoresInput{}
	}

	options := ListFHIRDatastoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFHIRDatastoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFHIRDatastoresPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFHIRDatastores page.
func (p *ListFHIRDatastoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFHIRDatastoresOutput, error) {
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

	result, err := p.client.ListFHIRDatastores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFHIRDatastores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "ListFHIRDatastores",
	}
}
