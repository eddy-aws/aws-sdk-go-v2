// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the phone numbers for the specified Amazon Connect
// instance. For more information about phone numbers, see Set Up Phone Numbers for
// Your Contact Center
// (https://docs.aws.amazon.com/connect/latest/adminguide/contact-center-phone-number.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumbers", params, optFns, addOperationListPhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumbersInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The ISO country code.
	PhoneNumberCountryCodes []types.PhoneNumberCountryCode

	// The type of phone number.
	PhoneNumberTypes []types.PhoneNumberType
}

type ListPhoneNumbersOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the phone numbers.
	PhoneNumberSummaryList []types.PhoneNumberSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumbers{}, middleware.After)
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
	if err = addOpListPhoneNumbersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbers(options.Region), middleware.Before); err != nil {
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

// ListPhoneNumbersAPIClient is a client that implements the ListPhoneNumbers
// operation.
type ListPhoneNumbersAPIClient interface {
	ListPhoneNumbers(context.Context, *ListPhoneNumbersInput, ...func(*Options)) (*ListPhoneNumbersOutput, error)
}

var _ ListPhoneNumbersAPIClient = (*Client)(nil)

// ListPhoneNumbersPaginatorOptions is the paginator options for ListPhoneNumbers
type ListPhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPhoneNumbersPaginator is a paginator for ListPhoneNumbers
type ListPhoneNumbersPaginator struct {
	options   ListPhoneNumbersPaginatorOptions
	client    ListPhoneNumbersAPIClient
	params    *ListPhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewListPhoneNumbersPaginator returns a new ListPhoneNumbersPaginator
func NewListPhoneNumbersPaginator(client ListPhoneNumbersAPIClient, params *ListPhoneNumbersInput, optFns ...func(*ListPhoneNumbersPaginatorOptions)) *ListPhoneNumbersPaginator {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	options := ListPhoneNumbersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPhoneNumbers page.
func (p *ListPhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListPhoneNumbers",
	}
}
