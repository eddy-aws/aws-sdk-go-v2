// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Fleet Advisor databases in your account.
func (c *Client) DescribeFleetAdvisorDatabases(ctx context.Context, params *DescribeFleetAdvisorDatabasesInput, optFns ...func(*Options)) (*DescribeFleetAdvisorDatabasesOutput, error) {
	if params == nil {
		params = &DescribeFleetAdvisorDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAdvisorDatabases", params, optFns, c.addOperationDescribeFleetAdvisorDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAdvisorDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetAdvisorDatabasesInput struct {

	// If you specify any of the following filters, the output includes information for
	// only those databases that meet the filter criteria:
	//
	// * database-id – The ID of
	// the database, for example d4610ac5-e323-4ad9-bc50-eaf7249dfe9d.
	//
	// * database-name
	// – The name of the database.
	//
	// * database-engine – The name of the database
	// engine.
	//
	// * server-ip-address – The IP address of the database server.
	//
	// *
	// database-ip-address – The IP address of the database.
	//
	// * collector-name – The
	// name of the associated Fleet Advisor collector.
	//
	// An example is:
	// describe-fleet-advisor-databases --filter
	// Name="database-id",Values="d4610ac5-e323-4ad9-bc50-eaf7249dfe9d"
	Filters []types.Filter

	// Sets the maximum number of records returned in the response.
	MaxRecords *int32

	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetAdvisorDatabasesOutput struct {

	// Provides descriptions of the Fleet Advisor collector databases, including the
	// database's collector, ID, and name.
	Databases []types.DatabaseResponse

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetAdvisorDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAdvisorDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAdvisorDatabases{}, middleware.After)
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
	if err = addOpDescribeFleetAdvisorDatabasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAdvisorDatabases(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAdvisorDatabasesAPIClient is a client that implements the
// DescribeFleetAdvisorDatabases operation.
type DescribeFleetAdvisorDatabasesAPIClient interface {
	DescribeFleetAdvisorDatabases(context.Context, *DescribeFleetAdvisorDatabasesInput, ...func(*Options)) (*DescribeFleetAdvisorDatabasesOutput, error)
}

var _ DescribeFleetAdvisorDatabasesAPIClient = (*Client)(nil)

// DescribeFleetAdvisorDatabasesPaginatorOptions is the paginator options for
// DescribeFleetAdvisorDatabases
type DescribeFleetAdvisorDatabasesPaginatorOptions struct {
	// Sets the maximum number of records returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAdvisorDatabasesPaginator is a paginator for
// DescribeFleetAdvisorDatabases
type DescribeFleetAdvisorDatabasesPaginator struct {
	options   DescribeFleetAdvisorDatabasesPaginatorOptions
	client    DescribeFleetAdvisorDatabasesAPIClient
	params    *DescribeFleetAdvisorDatabasesInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAdvisorDatabasesPaginator returns a new
// DescribeFleetAdvisorDatabasesPaginator
func NewDescribeFleetAdvisorDatabasesPaginator(client DescribeFleetAdvisorDatabasesAPIClient, params *DescribeFleetAdvisorDatabasesInput, optFns ...func(*DescribeFleetAdvisorDatabasesPaginatorOptions)) *DescribeFleetAdvisorDatabasesPaginator {
	if params == nil {
		params = &DescribeFleetAdvisorDatabasesInput{}
	}

	options := DescribeFleetAdvisorDatabasesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetAdvisorDatabasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAdvisorDatabasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetAdvisorDatabases page.
func (p *DescribeFleetAdvisorDatabasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAdvisorDatabasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeFleetAdvisorDatabases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetAdvisorDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeFleetAdvisorDatabases",
	}
}
