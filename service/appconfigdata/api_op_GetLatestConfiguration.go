// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfigdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the latest deployed configuration. This API may return empty
// Configuration data if the client already has the latest version. See
// StartConfigurationSession to obtain an InitialConfigurationToken to call this
// API. Each call to GetLatestConfiguration returns a new ConfigurationToken
// (NextPollConfigurationToken in the response). This new token MUST be provided to
// the next call to GetLatestConfiguration when polling for configuration updates.
// To avoid excess charges, we recommend that you include the
// ClientConfigurationVersion value with every call to GetConfiguration. This value
// must be saved on your client. Subsequent calls to GetConfiguration must pass
// this value by using the ClientConfigurationVersion parameter.
func (c *Client) GetLatestConfiguration(ctx context.Context, params *GetLatestConfigurationInput, optFns ...func(*Options)) (*GetLatestConfigurationOutput, error) {
	if params == nil {
		params = &GetLatestConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLatestConfiguration", params, optFns, c.addOperationGetLatestConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLatestConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request parameters for the GetLatestConfiguration API
type GetLatestConfigurationInput struct {

	// Token describing the current state of the configuration session. To obtain a
	// token, first call the StartConfigurationSession API. Note that every call to
	// GetLatestConfiguration will return a new ConfigurationToken
	// (NextPollConfigurationToken in the response) and MUST be provided to subsequent
	// GetLatestConfiguration API calls.
	//
	// This member is required.
	ConfigurationToken *string

	noSmithyDocumentSerde
}

// Response parameters for the GetLatestConfiguration API
type GetLatestConfigurationOutput struct {

	// The data of the configuration. Note that this may be empty if the client already
	// has the latest version of configuration.
	Configuration []byte

	// A standard MIME type describing the format of the configuration content.
	ContentType *string

	// The latest token describing the current state of the configuration session. This
	// MUST be provided to the next call to GetLatestConfiguration.
	NextPollConfigurationToken *string

	// The amount of time the client should wait before polling for configuration
	// updates again. See RequiredMinimumPollIntervalInSeconds to set the desired poll
	// interval.
	NextPollIntervalInSeconds int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLatestConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLatestConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLatestConfiguration{}, middleware.After)
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
	if err = addOpGetLatestConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLatestConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLatestConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetLatestConfiguration",
	}
}
