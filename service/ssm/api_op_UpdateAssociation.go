// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an association. You can update the association name and version, the
// document version, schedule, parameters, and Amazon Simple Storage Service
// (Amazon S3) output. In order to call this API operation, your Identity and
// Access Management (IAM) user account, group, or role must be configured with
// permission to call the DescribeAssociation API operation. If you don't have
// permission to call DescribeAssociation, then you receive the following error: An
// error occurred (AccessDeniedException) when calling the UpdateAssociation
// operation: User: isn't authorized to perform: ssm:DescribeAssociation on
// resource:  When you update an association, the association immediately runs
// against the specified targets.
func (c *Client) UpdateAssociation(ctx context.Context, params *UpdateAssociationInput, optFns ...func(*Options)) (*UpdateAssociationOutput, error) {
	if params == nil {
		params = &UpdateAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssociation", params, optFns, c.addOperationUpdateAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssociationInput struct {

	// The ID of the association you want to update.
	//
	// This member is required.
	AssociationId *string

	// By default, when you update an association, the system runs it immediately after
	// it is updated and then according to the schedule you specified. Specify this
	// option if you don't want an association to run immediately after you update it.
	// This parameter isn't supported for rate expressions. Also, if you specified this
	// option when you created the association, you can reset it. To do so, specify the
	// no-apply-only-at-cron-interval parameter when you update the association from
	// the command line. This parameter forces the association to run immediately after
	// updating it and according to the interval specified.
	ApplyOnlyAtCronInterval bool

	// The name of the association that you want to update.
	AssociationName *string

	// This parameter is provided for concurrency control purposes. You must specify
	// the latest association version in the service. If you want to ensure that this
	// request succeeds, either specify $LATEST, or omit this parameter.
	AssociationVersion *string

	// Choose the parameter that will define how your automation will branch out. This
	// target is required for associations that use an Automation runbook and target
	// resources by using rate controls. Automation is a capability of Amazon Web
	// Services Systems Manager.
	AutomationTargetParameterName *string

	// The names or Amazon Resource Names (ARNs) of the Change Calendar type documents
	// you want to gate your associations under. The associations only run when that
	// change calendar is open. For more information, see Amazon Web Services Systems
	// Manager Change Calendar
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar).
	CalendarNames []string

	// The severity level to assign to the association.
	ComplianceSeverity types.AssociationComplianceSeverity

	// The document version you want update for the association.
	DocumentVersion *string

	// The maximum number of targets allowed to run the association at the same time.
	// You can specify a number, for example 10, or a percentage of the target set, for
	// example 10%. The default value is 100%, which means all targets run the
	// association at the same time. If a new instance starts and attempts to run an
	// association while Systems Manager is running MaxConcurrency associations, the
	// association is allowed to run. During the next association interval, the new
	// instance will process its association within the limit specified for
	// MaxConcurrency.
	MaxConcurrency *string

	// The number of errors that are allowed before the system stops sending requests
	// to run the association on additional targets. You can specify either an absolute
	// number of errors, for example 10, or a percentage of the target set, for example
	// 10%. If you specify 3, for example, the system stops sending requests when the
	// fourth error is received. If you specify 0, then the system stops sending
	// requests after the first error is returned. If you run an association on 50
	// instances and set MaxError to 10%, then the system stops sending the request
	// when the sixth error is received. Executions that are already running an
	// association when MaxErrors is reached are allowed to complete, but some of these
	// executions may fail as well. If you need to ensure that there won't be more than
	// max-errors failed executions, set MaxConcurrency to 1 so that executions proceed
	// one at a time.
	MaxErrors *string

	// The name of the SSM Command document or Automation runbook that contains the
	// configuration information for the instance. You can specify Amazon Web
	// Services-predefined documents, documents you created, or a document that is
	// shared with you from another account. For Systems Manager document (SSM
	// document) that are shared with you from other Amazon Web Services accounts, you
	// must specify the complete SSM document ARN, in the following format:
	// arn:aws:ssm:region:account-id:document/document-name  For example:
	// arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document For Amazon Web
	// Services-predefined documents and SSM documents you created in your account, you
	// only need to specify the document name. For example, AWS-ApplyPatchBaseline or
	// My-Document.
	Name *string

	// An S3 bucket where you want to store the results of this request.
	OutputLocation *types.InstanceAssociationOutputLocation

	// The parameters you want to update for the association. If you create a parameter
	// using Parameter Store, a capability of Amazon Web Services Systems Manager, you
	// can reference the parameter using {{ssm:parameter-name}}.
	Parameters map[string][]string

	// The cron expression used to schedule the association that you want to update.
	ScheduleExpression *string

	// The mode for generating association compliance. You can specify AUTO or MANUAL.
	// In AUTO mode, the system uses the status of the association execution to
	// determine the compliance status. If the association execution runs successfully,
	// then the association is COMPLIANT. If the association execution doesn't run
	// successfully, the association is NON-COMPLIANT. In MANUAL mode, you must specify
	// the AssociationId as a parameter for the PutComplianceItems API operation. In
	// this case, compliance data isn't managed by State Manager, a capability of
	// Amazon Web Services Systems Manager. It is managed by your direct call to the
	// PutComplianceItems API operation. By default, all associations use AUTO mode.
	SyncCompliance types.AssociationSyncCompliance

	// A location is a combination of Amazon Web Services Regions and Amazon Web
	// Services accounts where you want to run the association. Use this action to
	// update an association in multiple Regions and multiple accounts.
	TargetLocations []types.TargetLocation

	// The targets of the association.
	Targets []types.Target

	noSmithyDocumentSerde
}

type UpdateAssociationOutput struct {

	// The description of the association that was updated.
	AssociationDescription *types.AssociationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAssociation{}, middleware.After)
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
	if err = addOpUpdateAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateAssociation",
	}
}
