// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns an array of recovery points.
func (c *Client) ListRecoveryPoints(ctx context.Context, params *ListRecoveryPointsInput, optFns ...func(*Options)) (*ListRecoveryPointsOutput, error) {
	if params == nil {
		params = &ListRecoveryPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecoveryPoints", params, optFns, c.addOperationListRecoveryPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecoveryPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecoveryPointsInput struct {

	// The time when creation of the recovery point finished.
	EndTime *time.Time

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults *int32

	// The name of the namespace to list recovery points for.
	NamespaceName *string

	// If your initial ListRecoveryPoints operation returns a nextToken, you can
	// include the returned nextToken in subsequent ListRecoveryPoints operations,
	// which returns results in the next page.
	NextToken *string

	// The time when the recovery point's creation was initiated.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListRecoveryPointsOutput struct {

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// The returned recovery point objects.
	RecoveryPoints []types.RecoveryPoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecoveryPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRecoveryPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRecoveryPoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecoveryPoints(options.Region), middleware.Before); err != nil {
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

// ListRecoveryPointsAPIClient is a client that implements the ListRecoveryPoints
// operation.
type ListRecoveryPointsAPIClient interface {
	ListRecoveryPoints(context.Context, *ListRecoveryPointsInput, ...func(*Options)) (*ListRecoveryPointsOutput, error)
}

var _ ListRecoveryPointsAPIClient = (*Client)(nil)

// ListRecoveryPointsPaginatorOptions is the paginator options for
// ListRecoveryPoints
type ListRecoveryPointsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecoveryPointsPaginator is a paginator for ListRecoveryPoints
type ListRecoveryPointsPaginator struct {
	options   ListRecoveryPointsPaginatorOptions
	client    ListRecoveryPointsAPIClient
	params    *ListRecoveryPointsInput
	nextToken *string
	firstPage bool
}

// NewListRecoveryPointsPaginator returns a new ListRecoveryPointsPaginator
func NewListRecoveryPointsPaginator(client ListRecoveryPointsAPIClient, params *ListRecoveryPointsInput, optFns ...func(*ListRecoveryPointsPaginatorOptions)) *ListRecoveryPointsPaginator {
	if params == nil {
		params = &ListRecoveryPointsInput{}
	}

	options := ListRecoveryPointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecoveryPointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecoveryPointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecoveryPoints page.
func (p *ListRecoveryPointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecoveryPointsOutput, error) {
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

	result, err := p.client.ListRecoveryPoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecoveryPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "ListRecoveryPoints",
	}
}
