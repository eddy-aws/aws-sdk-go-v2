// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a forecast for how much Amazon Web Services predicts that you will use
// over the forecast time period that you select, based on your past usage.
func (c *Client) GetUsageForecast(ctx context.Context, params *GetUsageForecastInput, optFns ...func(*Options)) (*GetUsageForecastOutput, error) {
	if params == nil {
		params = &GetUsageForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageForecast", params, optFns, addOperationGetUsageForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageForecastInput struct {

	// How granular you want the forecast to be. You can get 3 months of DAILY
	// forecasts or 12 months of MONTHLY forecasts. The GetUsageForecast operation
	// supports only DAILY and MONTHLY granularities.
	//
	// This member is required.
	Granularity types.Granularity

	// Which metric Cost Explorer uses to create your forecast. Valid values for a
	// GetUsageForecast call are the following:
	//
	// * USAGE_QUANTITY
	//
	// *
	// NORMALIZED_USAGE_AMOUNT
	//
	// This member is required.
	Metric types.Metric

	// The start and end dates of the period that you want to retrieve usage forecast
	// for. The start date is inclusive, but the end date is exclusive. For example, if
	// start is 2017-01-01 and end is 2017-05-01, then the cost and usage data is
	// retrieved from 2017-01-01 up to and including 2017-04-30 but not including
	// 2017-05-01. The start date must be equal to or later than the current date to
	// avoid a validation error.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The filters that you want to use to filter your forecast. The GetUsageForecast
	// API supports filtering by the following dimensions:
	//
	// * AZ
	//
	// * INSTANCE_TYPE
	//
	// *
	// LINKED_ACCOUNT
	//
	// * LINKED_ACCOUNT_NAME
	//
	// * OPERATION
	//
	// * PURCHASE_TYPE
	//
	// * REGION
	//
	// *
	// SERVICE
	//
	// * USAGE_TYPE
	//
	// * USAGE_TYPE_GROUP
	//
	// * RECORD_TYPE
	//
	// * OPERATING_SYSTEM
	//
	// *
	// TENANCY
	//
	// * SCOPE
	//
	// * PLATFORM
	//
	// * SUBSCRIPTION_ID
	//
	// * LEGAL_ENTITY_NAME
	//
	// *
	// DEPLOYMENT_OPTION
	//
	// * DATABASE_ENGINE
	//
	// * INSTANCE_TYPE_FAMILY
	//
	// *
	// BILLING_ENTITY
	//
	// * RESERVATION_ID
	//
	// * SAVINGS_PLAN_ARN
	Filter *types.Expression

	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence level.
	// The higher the confidence level, the more confident Cost Explorer is about the
	// actual value falling in the prediction interval. Higher confidence levels result
	// in wider prediction intervals.
	PredictionIntervalLevel *int32
}

type GetUsageForecastOutput struct {

	// The forecasts for your query, in order. For DAILY forecasts, this is a list of
	// days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []types.ForecastResult

	// How much you're forecasted to use over the forecast period.
	Total *types.MetricValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsageForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUsageForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUsageForecast{}, middleware.After)
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
	if err = addOpGetUsageForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUsageForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetUsageForecast",
	}
}
