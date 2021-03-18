// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the details of FlexMatch matchmaking configurations. This operation
// offers the following options: (1) retrieve all matchmaking configurations, (2)
// retrieve configurations for a specified list, or (3) retrieve all configurations
// that use a specified rule set name. When requesting multiple items, use the
// pagination parameters to retrieve results as a set of sequential pages. If
// successful, a configuration is returned for each requested name. When specifying
// a list of names, only configurations that currently exist are returned. Learn
// more  Setting up FlexMatch matchmakers
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/matchmaker-build.html)
// Related actions CreateMatchmakingConfiguration |
// DescribeMatchmakingConfigurations | UpdateMatchmakingConfiguration |
// DeleteMatchmakingConfiguration | CreateMatchmakingRuleSet |
// DescribeMatchmakingRuleSets | ValidateMatchmakingRuleSet |
// DeleteMatchmakingRuleSet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeMatchmakingConfigurations(ctx context.Context, params *DescribeMatchmakingConfigurationsInput, optFns ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeMatchmakingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMatchmakingConfigurations", params, optFns, addOperationDescribeMatchmakingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMatchmakingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeMatchmakingConfigurationsInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is limited to 10.
	Limit *int32

	// A unique identifier for the matchmaking configuration(s) to retrieve. You can
	// use either the configuration name or ARN value. To request all existing
	// configurations, leave this parameter empty.
	Names []string

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// A unique identifier for the matchmaking rule set. You can use either the rule
	// set name or ARN value. Use this parameter to retrieve all matchmaking
	// configurations that use this rule set.
	RuleSetName *string
}

// Represents the returned data in response to a request operation.
type DescribeMatchmakingConfigurationsOutput struct {

	// A collection of requested matchmaking configurations.
	Configurations []types.MatchmakingConfiguration

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMatchmakingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMatchmakingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMatchmakingConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMatchmakingConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeMatchmakingConfigurationsAPIClient is a client that implements the
// DescribeMatchmakingConfigurations operation.
type DescribeMatchmakingConfigurationsAPIClient interface {
	DescribeMatchmakingConfigurations(context.Context, *DescribeMatchmakingConfigurationsInput, ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error)
}

var _ DescribeMatchmakingConfigurationsAPIClient = (*Client)(nil)

// DescribeMatchmakingConfigurationsPaginatorOptions is the paginator options for
// DescribeMatchmakingConfigurations
type DescribeMatchmakingConfigurationsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is limited to 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMatchmakingConfigurationsPaginator is a paginator for
// DescribeMatchmakingConfigurations
type DescribeMatchmakingConfigurationsPaginator struct {
	options   DescribeMatchmakingConfigurationsPaginatorOptions
	client    DescribeMatchmakingConfigurationsAPIClient
	params    *DescribeMatchmakingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMatchmakingConfigurationsPaginator returns a new
// DescribeMatchmakingConfigurationsPaginator
func NewDescribeMatchmakingConfigurationsPaginator(client DescribeMatchmakingConfigurationsAPIClient, params *DescribeMatchmakingConfigurationsInput, optFns ...func(*DescribeMatchmakingConfigurationsPaginatorOptions)) *DescribeMatchmakingConfigurationsPaginator {
	if params == nil {
		params = &DescribeMatchmakingConfigurationsInput{}
	}

	options := DescribeMatchmakingConfigurationsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMatchmakingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMatchmakingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeMatchmakingConfigurations page.
func (p *DescribeMatchmakingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeMatchmakingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMatchmakingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeMatchmakingConfigurations",
	}
}
