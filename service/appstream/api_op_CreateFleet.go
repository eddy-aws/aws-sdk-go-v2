// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a fleet. A fleet consists of streaming instances that run a specified
// image.
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// The desired capacity for the fleet.
	//
	// This member is required.
	ComputeCapacity *types.ComputeCapacity

	// The instance type to use when launching fleet instances. The following instance
	// types are available:
	//
	// * stream.standard.small
	//
	// * stream.standard.medium
	//
	// *
	// stream.standard.large
	//
	// * stream.compute.large
	//
	// * stream.compute.xlarge
	//
	// *
	// stream.compute.2xlarge
	//
	// * stream.compute.4xlarge
	//
	// * stream.compute.8xlarge
	//
	// *
	// stream.memory.large
	//
	// * stream.memory.xlarge
	//
	// * stream.memory.2xlarge
	//
	// *
	// stream.memory.4xlarge
	//
	// * stream.memory.8xlarge
	//
	// * stream.memory.z1d.large
	//
	// *
	// stream.memory.z1d.xlarge
	//
	// * stream.memory.z1d.2xlarge
	//
	// *
	// stream.memory.z1d.3xlarge
	//
	// * stream.memory.z1d.6xlarge
	//
	// *
	// stream.memory.z1d.12xlarge
	//
	// * stream.graphics-design.large
	//
	// *
	// stream.graphics-design.xlarge
	//
	// * stream.graphics-design.2xlarge
	//
	// *
	// stream.graphics-design.4xlarge
	//
	// * stream.graphics-desktop.2xlarge
	//
	// *
	// stream.graphics.g4dn.xlarge
	//
	// * stream.graphics.g4dn.2xlarge
	//
	// *
	// stream.graphics.g4dn.4xlarge
	//
	// * stream.graphics.g4dn.8xlarge
	//
	// *
	// stream.graphics.g4dn.12xlarge
	//
	// * stream.graphics.g4dn.16xlarge
	//
	// *
	// stream.graphics-pro.4xlarge
	//
	// * stream.graphics-pro.8xlarge
	//
	// *
	// stream.graphics-pro.16xlarge
	//
	// This member is required.
	InstanceType *string

	// A unique name for the fleet.
	//
	// This member is required.
	Name *string

	// The description to display.
	Description *string

	// The amount of time that a streaming session remains active after users
	// disconnect. If users try to reconnect to the streaming session after a
	// disconnection or network interruption within this time interval, they are
	// connected to their previous session. Otherwise, they are connected to a new
	// session with a new streaming instance. Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *int32

	// The fleet name to display.
	DisplayName *string

	// The name of the directory and organizational unit (OU) to use to join the fleet
	// to a Microsoft Active Directory domain.
	DomainJoinInfo *types.DomainJoinInfo

	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess *bool

	// The fleet type. ALWAYS_ON Provides users with instant-on access to their apps.
	// You are charged for all running instances in your fleet, even if no users are
	// streaming apps. ON_DEMAND Provide users with access to applications after they
	// connect, which takes one to two minutes. You are charged for instance streaming
	// when users are connected and a small hourly fee for instances that are not
	// streaming apps.
	FleetType types.FleetType

	// The Amazon Resource Name (ARN) of the IAM role to apply to the fleet. To assume
	// a role, a fleet instance calls the AWS Security Token Service (STS) AssumeRole
	// API operation and passes the ARN of the role to use. The operation creates a new
	// session with temporary credentials. AppStream 2.0 retrieves the temporary
	// credentials and creates the appstream_machine_role credential profile on the
	// instance. For more information, see Using an IAM Role to Grant Permissions to
	// Applications and Scripts Running on AppStream 2.0 Streaming Instances
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string

	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the DisconnectTimeoutInSeconds
	// time interval begins. Users are notified before they are disconnected due to
	// inactivity. If they try to reconnect to the streaming session before the time
	// interval specified in DisconnectTimeoutInSeconds elapses, they are connected to
	// their previous session. Users are considered idle when they stop providing
	// keyboard or mouse input during their streaming session. File uploads and
	// downloads, audio in, audio out, and pixels changing do not qualify as user
	// activity. If users continue to be idle after the time interval in
	// IdleDisconnectTimeoutInSeconds elapses, they are disconnected. To prevent users
	// from being disconnected due to inactivity, specify a value of 0. Otherwise,
	// specify a value between 60 and 3600. The default value is 0. If you enable this
	// feature, we recommend that you specify a value that corresponds exactly to a
	// whole number of minutes (for example, 60, 120, and 180). If you don't do this,
	// the value is rounded to the nearest minute. For example, if you specify a value
	// of 70, users are disconnected after 1 minute of inactivity. If you specify a
	// value that is at the midpoint between two different minutes, the value is
	// rounded up. For example, if you specify a value of 90, users are disconnected
	// after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *int32

	// The ARN of the public, private, or shared image to use.
	ImageArn *string

	// The name of the image used to create the fleet.
	ImageName *string

	// The maximum amount of time that a streaming session can remain active, in
	// seconds. If users are still connected to a streaming instance five minutes
	// before this limit is reached, they are prompted to save any open documents
	// before being disconnected. After this time elapses, the instance is terminated
	// and replaced by a new instance. Specify a value between 600 and 360000.
	MaxUserDurationInSeconds *int32

	// The AppStream 2.0 view that is displayed to your users when they stream from the
	// fleet. When APP is specified, only the windows of applications opened by users
	// display. When DESKTOP is specified, the standard desktop that is provided by the
	// operating system displays. The default value is APP.
	StreamView types.StreamView

	// The tags to associate with the fleet. A tag is a key-value pair, and the value
	// is optional. For example, Environment=Test. If you do not specify a value,
	// Environment=. If you do not specify a value, the value is set to an empty
	// string. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: _ . : / = + \ - @
	// For more information, see Tagging Your Resources
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string

	// The VPC configuration for the fleet.
	VpcConfig *types.VpcConfig
}

type CreateFleetOutput struct {

	// Information about the fleet.
	Fleet *types.Fleet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleet{}, middleware.After)
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
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateFleet",
	}
}
