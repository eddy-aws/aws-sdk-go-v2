// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a model compilation job. To create a model compilation
// job, use CreateCompilationJob. To get information about multiple model
// compilation jobs, use ListCompilationJobs.
func (c *Client) DescribeCompilationJob(ctx context.Context, params *DescribeCompilationJobInput, optFns ...func(*Options)) (*DescribeCompilationJobOutput, error) {
	if params == nil {
		params = &DescribeCompilationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCompilationJob", params, optFns, addOperationDescribeCompilationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCompilationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCompilationJobInput struct {

	// The name of the model compilation job that you want information about.
	//
	// This member is required.
	CompilationJobName *string
}

type DescribeCompilationJobOutput struct {

	// The Amazon Resource Name (ARN) of the model compilation job.
	//
	// This member is required.
	CompilationJobArn *string

	// The name of the model compilation job.
	//
	// This member is required.
	CompilationJobName *string

	// The status of the model compilation job.
	//
	// This member is required.
	CompilationJobStatus types.CompilationJobStatus

	// The time that the model compilation job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// If a model compilation job failed, the reason it failed.
	//
	// This member is required.
	FailureReason *string

	// Information about the location in Amazon S3 of the input model artifacts, the
	// name and shape of the expected data inputs, and the framework in which the model
	// was trained.
	//
	// This member is required.
	InputConfig *types.InputConfig

	// The time that the status of the model compilation job was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Information about the location in Amazon S3 that has been configured for storing
	// the model artifacts used in the compilation job.
	//
	// This member is required.
	ModelArtifacts *types.ModelArtifacts

	// Information about the output location for the compiled model and the target
	// device that the model runs on.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker assumes to
	// perform the model compilation job.
	//
	// This member is required.
	RoleArn *string

	// Specifies a limit to how long a model compilation job can run. When the job
	// reaches the time limit, Amazon SageMaker ends the compilation job. Use this API
	// to cap model training costs.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// The time when the model compilation job on a compilation job instance ended. For
	// a successful or stopped job, this is when the job's model artifacts have
	// finished uploading. For a failed job, this is when Amazon SageMaker detected
	// that the job failed.
	CompilationEndTime *time.Time

	// The time when the model compilation job started the CompilationJob instances.
	// You are billed for the time between this timestamp and the timestamp in the
	// DescribeCompilationJobResponse$CompilationEndTime field. In Amazon CloudWatch
	// Logs, the start time might be later than this time. That's because it takes time
	// to download the compilation job, which depends on the size of the compilation
	// job container.
	CompilationStartTime *time.Time

	// Provides a BLAKE2 hash value that identifies the compiled model artifacts in
	// Amazon S3.
	ModelDigests *types.ModelDigests

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCompilationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCompilationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCompilationJob{}, middleware.After)
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
	if err = addOpDescribeCompilationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCompilationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCompilationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeCompilationJob",
	}
}
