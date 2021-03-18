// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides metadata about a version of a bot.
func (c *Client) DescribeBotVersion(ctx context.Context, params *DescribeBotVersionInput, optFns ...func(*Options)) (*DescribeBotVersionOutput, error) {
	if params == nil {
		params = &DescribeBotVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotVersion", params, optFns, addOperationDescribeBotVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotVersionInput struct {

	// The identifier of the bot containing the version to return metadata for.
	//
	// This member is required.
	BotId *string

	// The version of the bot to return metadata for.
	//
	// This member is required.
	BotVersion *string
}

type DescribeBotVersionOutput struct {

	// The identifier of the bot that contains the version.
	BotId *string

	// The name of the bot that contains the version.
	BotName *string

	// The current status of the bot. When the status is Available, the bot version is
	// ready for use.
	BotStatus types.BotStatus

	// The version of the bot to describe.
	BotVersion *string

	// A timestamp of the date and time that the bot version was created.
	CreationDateTime *time.Time

	// Data privacy settings for the bot version.
	DataPrivacy *types.DataPrivacy

	// The description specified for the bot.
	Description *string

	// If the botStatus is Failed, this contains a list of reasons that the version
	// couldn't be built.
	FailureReasons []string

	// The number of seconds that a session with the bot remains active before it is
	// discarded by Amazon Lex.
	IdleSessionTTLInSeconds *int32

	// The Amazon Resource Name (ARN) of an IAM role that has permission to access the
	// bot version.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBotVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotVersion{}, middleware.After)
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
	if err = addOpDescribeBotVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBotVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeBotVersion",
	}
}
