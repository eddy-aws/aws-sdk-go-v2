// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about your channel's schedule.
func (c *Client) GetChannelSchedule(ctx context.Context, params *GetChannelScheduleInput, optFns ...func(*Options)) (*GetChannelScheduleOutput, error) {
	if params == nil {
		params = &GetChannelScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChannelSchedule", params, optFns, addOperationGetChannelScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChannelScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChannelScheduleInput struct {

	// The identifier for the channel you are working on.
	//
	// This member is required.
	ChannelName *string

	// The schedule duration in minutes. The maximum duration is 4320 minutes (three
	// days).
	DurationMinutes *string

	// Upper bound on number of records to return. The maximum number of results is
	// 100.
	MaxResults int32

	// Pagination token from the GET list request. Use the token to fetch the next page
	// of results.
	NextToken *string
}

type GetChannelScheduleOutput struct {

	// An array of schedule entries for the channel.
	Items []types.ScheduleEntry

	// Pagination token from the GET list request. Use the token to fetch the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetChannelScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChannelSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChannelSchedule{}, middleware.After)
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
	if err = addOpGetChannelScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChannelSchedule(options.Region), middleware.Before); err != nil {
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

// GetChannelScheduleAPIClient is a client that implements the GetChannelSchedule
// operation.
type GetChannelScheduleAPIClient interface {
	GetChannelSchedule(context.Context, *GetChannelScheduleInput, ...func(*Options)) (*GetChannelScheduleOutput, error)
}

var _ GetChannelScheduleAPIClient = (*Client)(nil)

// GetChannelSchedulePaginatorOptions is the paginator options for
// GetChannelSchedule
type GetChannelSchedulePaginatorOptions struct {
	// Upper bound on number of records to return. The maximum number of results is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetChannelSchedulePaginator is a paginator for GetChannelSchedule
type GetChannelSchedulePaginator struct {
	options   GetChannelSchedulePaginatorOptions
	client    GetChannelScheduleAPIClient
	params    *GetChannelScheduleInput
	nextToken *string
	firstPage bool
}

// NewGetChannelSchedulePaginator returns a new GetChannelSchedulePaginator
func NewGetChannelSchedulePaginator(client GetChannelScheduleAPIClient, params *GetChannelScheduleInput, optFns ...func(*GetChannelSchedulePaginatorOptions)) *GetChannelSchedulePaginator {
	if params == nil {
		params = &GetChannelScheduleInput{}
	}

	options := GetChannelSchedulePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetChannelSchedulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetChannelSchedulePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetChannelSchedule page.
func (p *GetChannelSchedulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetChannelScheduleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetChannelSchedule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetChannelSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "GetChannelSchedule",
	}
}
