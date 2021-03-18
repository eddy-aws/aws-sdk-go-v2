// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently bans a member from a channel. Moderators can't add banned members to
// a channel. To undo a ban, you first have to DeleteChannelBan, and then
// CreateChannelMembership. Bans are cleaned up when you delete users or channels.
// If you ban a user who is already part of a channel, that user is automatically
// kicked from the channel. The x-amz-chime-bearer request header is mandatory. Use
// the AppInstanceUserArn of the user that makes the API call as the value in the
// header.
func (c *Client) CreateChannelBan(ctx context.Context, params *CreateChannelBanInput, optFns ...func(*Options)) (*CreateChannelBanOutput, error) {
	if params == nil {
		params = &CreateChannelBanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannelBan", params, optFns, addOperationCreateChannelBanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelBanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelBanInput struct {

	// The ARN of the ban request.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the member being banned.
	//
	// This member is required.
	MemberArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string
}

type CreateChannelBanOutput struct {

	// The ARN of the response to the ban request.
	ChannelArn *string

	// The ChannelArn and BannedIdentity of the member in the ban response.
	Member *types.Identity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateChannelBanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannelBan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannelBan{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateChannelBanMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateChannelBanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannelBan(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateChannelBanMiddleware struct {
}

func (*endpointPrefix_opCreateChannelBanMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateChannelBanMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateChannelBanMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateChannelBanMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateChannelBan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateChannelBan",
	}
}
