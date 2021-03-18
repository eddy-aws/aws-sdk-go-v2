// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new schedule for one or more DataBrew jobs. Jobs can be run at a
// specific date and time, or at regular intervals.
func (c *Client) CreateSchedule(ctx context.Context, params *CreateScheduleInput, optFns ...func(*Options)) (*CreateScheduleOutput, error) {
	if params == nil {
		params = &CreateScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSchedule", params, optFns, addOperationCreateScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduleInput struct {

	// The date or dates and time or times when the jobs are to be run. For more
	// information, see Cron expressions
	// (https://docs.aws.amazon.com/databrew/latest/dg/jobs.cron.html) in the AWS Glue
	// DataBrew Developer Guide.
	//
	// This member is required.
	CronExpression *string

	// A unique name for the schedule. Valid characters are alphanumeric (A-Z, a-z,
	// 0-9), hyphen (-), period (.), and space.
	//
	// This member is required.
	Name *string

	// The name or names of one or more jobs to be run.
	JobNames []string

	// Metadata tags to apply to this schedule.
	Tags map[string]string
}

type CreateScheduleOutput struct {

	// The name of the schedule that was created.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSchedule{}, middleware.After)
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
	if err = addOpCreateScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "CreateSchedule",
	}
}
