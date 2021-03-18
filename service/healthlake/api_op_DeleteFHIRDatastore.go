// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Data Store.
func (c *Client) DeleteFHIRDatastore(ctx context.Context, params *DeleteFHIRDatastoreInput, optFns ...func(*Options)) (*DeleteFHIRDatastoreOutput, error) {
	if params == nil {
		params = &DeleteFHIRDatastoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFHIRDatastore", params, optFns, addOperationDeleteFHIRDatastoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFHIRDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFHIRDatastoreInput struct {

	// The AWS-generated ID for the Data Store to be deleted.
	DatastoreId *string
}

type DeleteFHIRDatastoreOutput struct {

	// The Amazon Resource Name (ARN) that gives Amazon HealthLake access permission.
	//
	// This member is required.
	DatastoreArn *string

	// The AWS endpoint for the Data Store the user has requested to be deleted.
	//
	// This member is required.
	DatastoreEndpoint *string

	// The AWS-generated ID for the Data Store to be deleted.
	//
	// This member is required.
	DatastoreId *string

	// The status of the Data Store that the user has requested to be deleted.
	//
	// This member is required.
	DatastoreStatus types.DatastoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFHIRDatastoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFHIRDatastore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFHIRDatastore{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFHIRDatastore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFHIRDatastore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "DeleteFHIRDatastore",
	}
}
