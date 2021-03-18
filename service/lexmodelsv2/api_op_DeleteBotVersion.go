// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specific version of a bot. To delete all version of a bot, use the
// DeleteBot operation.
func (c *Client) DeleteBotVersion(ctx context.Context, params *DeleteBotVersionInput, optFns ...func(*Options)) (*DeleteBotVersionOutput, error) {
	if params == nil {
		params = &DeleteBotVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBotVersion", params, optFns, addOperationDeleteBotVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotVersionInput struct {

	// The identifier of the bot that contains the version.
	//
	// This member is required.
	BotId *string

	// The version of the bot to delete.
	//
	// This member is required.
	BotVersion *string

	// By default, the DeleteBotVersion operations throws a ResourceInUseException
	// exception if you try to delete a bot version that has an alias pointing at it.
	// Set the skipResourceInUseCheck parameter to true to skip this check and remove
	// the version even if an alias points to it.
	SkipResourceInUseCheck bool
}

type DeleteBotVersionOutput struct {

	// The identifier of the bot that is being deleted.
	BotId *string

	// The current status of the bot.
	BotStatus types.BotStatus

	// The version of the bot that is being deleted.
	BotVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBotVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBotVersion{}, middleware.After)
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
	if err = addOpDeleteBotVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBotVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBotVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBotVersion",
	}
}
