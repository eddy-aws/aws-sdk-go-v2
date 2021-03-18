// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs a command remotely on a container within a task.
func (c *Client) ExecuteCommand(ctx context.Context, params *ExecuteCommandInput, optFns ...func(*Options)) (*ExecuteCommandOutput, error) {
	if params == nil {
		params = &ExecuteCommandInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteCommand", params, optFns, addOperationExecuteCommandMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteCommandOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteCommandInput struct {

	// The command to run on the container.
	//
	// This member is required.
	Command *string

	// Use this flag to run your command in interactive mode.
	//
	// This member is required.
	Interactive bool

	// The Amazon Resource Name (ARN) or ID of the task the container is part of.
	//
	// This member is required.
	Task *string

	// The Amazon Resource Name (ARN) or short name of the cluster the task is running
	// in. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string

	// The name of the container to execute the command on. A container name only needs
	// to be specified for tasks containing multiple containers.
	Container *string
}

type ExecuteCommandOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the container.
	ContainerArn *string

	// The name of the container.
	ContainerName *string

	// Whether or not the execute command session is running in interactive mode.
	Interactive bool

	// The details of the SSM session that was created for this instance of
	// execute-command.
	Session *types.Session

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExecuteCommandMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExecuteCommand{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExecuteCommand{}, middleware.After)
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
	if err = addOpExecuteCommandValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteCommand(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteCommand(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ExecuteCommand",
	}
}
