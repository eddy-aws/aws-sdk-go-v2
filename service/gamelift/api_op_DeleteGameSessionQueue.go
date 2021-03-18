// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a game session queue. Once a queue is successfully deleted, unfulfilled
// StartGameSessionPlacement requests that reference the queue will fail. To delete
// a queue, specify the queue name. Learn more  Using Multi-Region Queues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-intro.html)
// Related actions CreateGameSessionQueue | DescribeGameSessionQueues |
// UpdateGameSessionQueue | DeleteGameSessionQueue | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteGameSessionQueue(ctx context.Context, params *DeleteGameSessionQueueInput, optFns ...func(*Options)) (*DeleteGameSessionQueueOutput, error) {
	if params == nil {
		params = &DeleteGameSessionQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGameSessionQueue", params, optFns, addOperationDeleteGameSessionQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGameSessionQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteGameSessionQueueInput struct {

	// A descriptive label that is associated with game session queue. Queue names must
	// be unique within each Region. You can use either the queue ID or ARN value.
	//
	// This member is required.
	Name *string
}

type DeleteGameSessionQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGameSessionQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGameSessionQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGameSessionQueue{}, middleware.After)
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
	if err = addOpDeleteGameSessionQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGameSessionQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGameSessionQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteGameSessionQueue",
	}
}
