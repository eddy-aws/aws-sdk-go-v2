// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of identity provider configurations.
func (c *Client) ListIdentityProviderConfigs(ctx context.Context, params *ListIdentityProviderConfigsInput, optFns ...func(*Options)) (*ListIdentityProviderConfigsOutput, error) {
	if params == nil {
		params = &ListIdentityProviderConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityProviderConfigs", params, optFns, addOperationListIdentityProviderConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityProviderConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIdentityProviderConfigsInput struct {

	// The cluster name that you want to list identity provider configurations for.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of identity provider configurations returned by
	// ListIdentityProviderConfigs in paginated output. When you use this parameter,
	// ListIdentityProviderConfigs returns only maxResults results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListIdentityProviderConfigs request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, ListIdentityProviderConfigs returns up to 100 results and a
	// nextToken value, if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// IdentityProviderConfigsRequest where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value.
	NextToken *string
}

type ListIdentityProviderConfigsOutput struct {

	// The identity provider configurations for the cluster.
	IdentityProviderConfigs []types.IdentityProviderConfig

	// The nextToken value returned from a previous paginated
	// ListIdentityProviderConfigsResponse where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIdentityProviderConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIdentityProviderConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdentityProviderConfigs{}, middleware.After)
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
	if err = addOpListIdentityProviderConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityProviderConfigs(options.Region), middleware.Before); err != nil {
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

// ListIdentityProviderConfigsAPIClient is a client that implements the
// ListIdentityProviderConfigs operation.
type ListIdentityProviderConfigsAPIClient interface {
	ListIdentityProviderConfigs(context.Context, *ListIdentityProviderConfigsInput, ...func(*Options)) (*ListIdentityProviderConfigsOutput, error)
}

var _ ListIdentityProviderConfigsAPIClient = (*Client)(nil)

// ListIdentityProviderConfigsPaginatorOptions is the paginator options for
// ListIdentityProviderConfigs
type ListIdentityProviderConfigsPaginatorOptions struct {
	// The maximum number of identity provider configurations returned by
	// ListIdentityProviderConfigs in paginated output. When you use this parameter,
	// ListIdentityProviderConfigs returns only maxResults results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListIdentityProviderConfigs request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, ListIdentityProviderConfigs returns up to 100 results and a
	// nextToken value, if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIdentityProviderConfigsPaginator is a paginator for
// ListIdentityProviderConfigs
type ListIdentityProviderConfigsPaginator struct {
	options   ListIdentityProviderConfigsPaginatorOptions
	client    ListIdentityProviderConfigsAPIClient
	params    *ListIdentityProviderConfigsInput
	nextToken *string
	firstPage bool
}

// NewListIdentityProviderConfigsPaginator returns a new
// ListIdentityProviderConfigsPaginator
func NewListIdentityProviderConfigsPaginator(client ListIdentityProviderConfigsAPIClient, params *ListIdentityProviderConfigsInput, optFns ...func(*ListIdentityProviderConfigsPaginatorOptions)) *ListIdentityProviderConfigsPaginator {
	if params == nil {
		params = &ListIdentityProviderConfigsInput{}
	}

	options := ListIdentityProviderConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIdentityProviderConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIdentityProviderConfigsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListIdentityProviderConfigs page.
func (p *ListIdentityProviderConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIdentityProviderConfigsOutput, error) {
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

	result, err := p.client.ListIdentityProviderConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIdentityProviderConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListIdentityProviderConfigs",
	}
}
