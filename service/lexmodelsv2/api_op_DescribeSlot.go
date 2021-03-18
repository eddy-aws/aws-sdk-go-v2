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

// Gets metadata information about a slot.
func (c *Client) DescribeSlot(ctx context.Context, params *DescribeSlotInput, optFns ...func(*Options)) (*DescribeSlotOutput, error) {
	if params == nil {
		params = &DescribeSlotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSlot", params, optFns, addOperationDescribeSlotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSlotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSlotInput struct {

	// The identifier of the bot associated with the slot.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with the slot.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the intent that contains the slot.
	//
	// This member is required.
	IntentId *string

	// The identifier of the language and locale of the slot to describe. The string
	// must match one of the supported locales. For more information, see
	// https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html
	// (https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html).
	//
	// This member is required.
	LocaleId *string

	// The unique identifier for the slot.
	//
	// This member is required.
	SlotId *string
}

type DescribeSlotOutput struct {

	// The identifier of the bot associated with the slot.
	BotId *string

	// The version of the bot associated with the slot.
	BotVersion *string

	// A timestamp of the date and time that the slot was created.
	CreationDateTime *time.Time

	// The description specified for the slot.
	Description *string

	// The identifier of the intent associated with the slot.
	IntentId *string

	// A timestamp of the date and time that the slot was last updated.
	LastUpdatedDateTime *time.Time

	// The language and locale specified for the slot.
	LocaleId *string

	// Whether slot values are shown in Amazon CloudWatch logs. If the value is None,
	// the actual value of the slot is shown in logs.
	ObfuscationSetting *types.ObfuscationSetting

	// The unique identifier generated for the slot.
	SlotId *string

	// The name specified for the slot.
	SlotName *string

	// The identifier of the slot type that determines the values entered into the
	// slot.
	SlotTypeId *string

	// Prompts that Amazon Lex uses to elicit a value for the slot.
	ValueElicitationSetting *types.SlotValueElicitationSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSlotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSlot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSlot{}, middleware.After)
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
	if err = addOpDescribeSlotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSlot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSlot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeSlot",
	}
}
