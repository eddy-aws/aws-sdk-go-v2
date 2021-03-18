// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Creates a new queue for the specified Amazon Connect instance.
func (c *Client) CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) {
	if params == nil {
		params = &CreateQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQueue", params, optFns, addOperationCreateQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQueueInput struct {

	// The identifier for the hours of operation.
	//
	// This member is required.
	HoursOfOperationId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The name of the queue.
	//
	// This member is required.
	Name *string

	// The description of the queue.
	Description *string

	// The maximum number of contacts that can be in the queue before it is considered
	// full.
	MaxContacts int32

	// The outbound caller ID name, number, and outbound whisper flow.
	OutboundCallerConfig *types.OutboundCallerConfig

	// The quick connects available to agents who are working the queue.
	QuickConnectIds []string

	// One or more tags.
	Tags map[string]string
}

type CreateQueueOutput struct {

	// The Amazon Resource Name (ARN) of the queue.
	QueueArn *string

	// The identifier for the queue.
	QueueId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQueue{}, middleware.After)
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
	if err = addOpCreateQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateQueue",
	}
}
