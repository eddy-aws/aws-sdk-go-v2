// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Channels created under a single Chime App as a paginated list. You can
// specify filters to narrow results. Functionality & restrictions
//
// * Use privacy =
// PUBLIC to retrieve all public channels in the account
//
// * Only an
// AppInstanceAdmin can set privacy = PRIVATE to list the private channels in an
// account.
//
// The x-amz-chime-bearer request header is mandatory. Use the
// AppInstanceUserArn of the user that makes the API call as the value in the
// header.
func (c *Client) ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) {
	if params == nil {
		params = &ListChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannels", params, optFns, addOperationListChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelsInput struct {

	// The ARN of the AppInstance.
	//
	// This member is required.
	AppInstanceArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	// The maximum number of channels that you want to return.
	MaxResults *int32

	// The token passed by previous API calls until all requested channels are
	// returned.
	NextToken *string

	// The privacy setting. PUBLIC retrieves all the public channels. PRIVATE retrieves
	// private channels. Only an AppInstanceAdmin can retrieve private channels.
	Privacy types.ChannelPrivacy
}

type ListChannelsOutput struct {

	// The information about each channel.
	Channels []types.ChannelSummary

	// The token returned from previous API requests until the number of channels is
	// reached.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannels{}, middleware.After)
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
	if err = addEndpointPrefix_opListChannelsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListChannelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannels(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListChannelsMiddleware struct {
}

func (*endpointPrefix_opListChannelsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListChannelsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListChannelsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListChannelsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListChannelsAPIClient is a client that implements the ListChannels operation.
type ListChannelsAPIClient interface {
	ListChannels(context.Context, *ListChannelsInput, ...func(*Options)) (*ListChannelsOutput, error)
}

var _ ListChannelsAPIClient = (*Client)(nil)

// ListChannelsPaginatorOptions is the paginator options for ListChannels
type ListChannelsPaginatorOptions struct {
	// The maximum number of channels that you want to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelsPaginator is a paginator for ListChannels
type ListChannelsPaginator struct {
	options   ListChannelsPaginatorOptions
	client    ListChannelsAPIClient
	params    *ListChannelsInput
	nextToken *string
	firstPage bool
}

// NewListChannelsPaginator returns a new ListChannelsPaginator
func NewListChannelsPaginator(client ListChannelsAPIClient, params *ListChannelsInput, optFns ...func(*ListChannelsPaginatorOptions)) *ListChannelsPaginator {
	if params == nil {
		params = &ListChannelsInput{}
	}

	options := ListChannelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListChannels page.
func (p *ListChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelsOutput, error) {
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

	result, err := p.client.ListChannels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListChannels",
	}
}
