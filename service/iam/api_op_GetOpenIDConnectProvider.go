// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified OpenID Connect (OIDC) provider resource
// object in IAM.
func (c *Client) GetOpenIDConnectProvider(ctx context.Context, params *GetOpenIDConnectProviderInput, optFns ...func(*Options)) (*GetOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &GetOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpenIDConnectProvider", params, optFns, addOperationGetOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpenIDConnectProviderInput struct {

	// The Amazon Resource Name (ARN) of the OIDC provider resource object in IAM to
	// get information for. You can get a list of OIDC provider resource ARNs by using
	// the ListOpenIDConnectProviders operation. For more information about ARNs, see
	// Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	OpenIDConnectProviderArn *string
}

// Contains the response to a successful GetOpenIDConnectProvider request.
type GetOpenIDConnectProviderOutput struct {

	// A list of client IDs (also known as audiences) that are associated with the
	// specified IAM OIDC provider resource object. For more information, see
	// CreateOpenIDConnectProvider.
	ClientIDList []string

	// The date and time when the IAM OIDC provider resource object was created in the
	// AWS account.
	CreateDate *time.Time

	// A list of tags that are attached to the specified IAM OIDC provider. The
	// returned list of tags is sorted by tag key. For more information about tagging,
	// see Tagging IAM resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
	// Guide.
	Tags []types.Tag

	// A list of certificate thumbprints that are associated with the specified IAM
	// OIDC provider resource object. For more information, see
	// CreateOpenIDConnectProvider.
	ThumbprintList []string

	// The URL that the IAM OIDC provider resource object is associated with. For more
	// information, see CreateOpenIDConnectProvider.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetOpenIDConnectProvider{}, middleware.After)
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
	if err = addOpGetOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetOpenIDConnectProvider",
	}
}
