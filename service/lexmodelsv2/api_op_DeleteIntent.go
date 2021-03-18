// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified intent. Deleting an intent also deletes the slots
// associated with the intent.
func (c *Client) DeleteIntent(ctx context.Context, params *DeleteIntentInput, optFns ...func(*Options)) (*DeleteIntentOutput, error) {
	if params == nil {
		params = &DeleteIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIntent", params, optFns, addOperationDeleteIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIntentInput struct {

	// The identifier of the bot associated with the intent.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with the intent.
	//
	// This member is required.
	BotVersion *string

	// The unique identifier of the intent to delete.
	//
	// This member is required.
	IntentId *string

	// The identifier of the language and locale where the bot will be deleted. The
	// string must match one of the supported locales. For more information, see
	// https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html
	// (https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html).
	//
	// This member is required.
	LocaleId *string
}

type DeleteIntentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIntent{}, middleware.After)
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
	if err = addOpDeleteIntentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIntent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteIntent",
	}
}
