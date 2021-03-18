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

// Creates a custom slot type To create a custom slot type, specify a name for the
// slot type and a set of enumeration values, the values that a slot of this type
// can assume.
func (c *Client) CreateSlotType(ctx context.Context, params *CreateSlotTypeInput, optFns ...func(*Options)) (*CreateSlotTypeOutput, error) {
	if params == nil {
		params = &CreateSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSlotType", params, optFns, addOperationCreateSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSlotTypeInput struct {

	// The identifier of the bot associated with this slot type.
	//
	// This member is required.
	BotId *string

	// The identifier of the bot version associated with this slot type.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that the slot type will be used in.
	// The string must match one of the supported locales. All of the bots, intents,
	// and slots used by the slot type must have the same locale. For more information,
	// see https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html
	// (https://docs.aws.amazon.com/lex/latest/dg/supported-locales.html).
	//
	// This member is required.
	LocaleId *string

	// The name for the slot. A slot type name must be unique within the account.
	//
	// This member is required.
	SlotTypeName *string

	// Determines the strategy that Amazon Lex uses to select a value from the list of
	// possible values. The field can be set to one of the following values:
	//
	// *
	// OriginalValue - Returns the value entered by the user, if the user value is
	// similar to the slot value.
	//
	// * TopResolution - If there is a resolution list for
	// the slot, return the first value in the resolution list. If there is no
	// resolution list, return null.
	//
	// If you don't specify the valueSelectionSetting
	// parameter, the default is OriginalValue.
	//
	// This member is required.
	ValueSelectionSetting *types.SlotValueSelectionSetting

	// A description of the slot type. Use the description to help identify the slot
	// type in lists.
	Description *string

	// The built-in slot type used as a parent of this slot type. When you define a
	// parent slot type, the new slot type has the configuration of the parent slot
	// type. Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string

	// A list of SlotTypeValue objects that defines the values that the slot type can
	// take. Each value can have a list of synonyms, additional values that help train
	// the machine learning model about the values that it resolves for a slot.
	SlotTypeValues []types.SlotTypeValue
}

type CreateSlotTypeOutput struct {

	// The identifier for the bot associated with the slot type.
	BotId *string

	// The version of the bot associated with the slot type.
	BotVersion *string

	// A timestamp of the date and time that the slot type was created.
	CreationDateTime *time.Time

	// The description specified for the slot type.
	Description *string

	// The specified language and local specified for the slot type.
	LocaleId *string

	// The signature of the base slot type specified for the slot type.
	ParentSlotTypeSignature *string

	// The unique identifier assigned to the slot type. Use this to identify the slot
	// type in the UpdateSlotType and DeleteSlotType operations.
	SlotTypeId *string

	// The name specified for the slot type.
	SlotTypeName *string

	// The list of values that the slot type can assume.
	SlotTypeValues []types.SlotTypeValue

	// The strategy that Amazon Lex uses to select a value from the list of possible
	// values.
	ValueSelectionSetting *types.SlotValueSelectionSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSlotType{}, middleware.After)
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
	if err = addOpCreateSlotTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSlotType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSlotType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateSlotType",
	}
}
