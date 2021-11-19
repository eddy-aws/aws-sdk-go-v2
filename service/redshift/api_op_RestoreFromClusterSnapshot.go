// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cluster from a snapshot. By default, Amazon Redshift creates the
// resulting cluster with the same configuration as the original cluster from which
// the snapshot was created, except that the new cluster is created with the
// default cluster security and parameter groups. After Amazon Redshift creates the
// cluster, you can use the ModifyCluster API to associate a different security
// group and different parameter group with the restored cluster. If you are using
// a DS node type, you can also choose to change to another DS node type of the
// same size during restore. If you restore a cluster into a VPC, you must provide
// a cluster subnet group where you want the cluster restored. For more information
// about working with snapshots, go to Amazon Redshift Snapshots
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) RestoreFromClusterSnapshot(ctx context.Context, params *RestoreFromClusterSnapshotInput, optFns ...func(*Options)) (*RestoreFromClusterSnapshotOutput, error) {
	if params == nil {
		params = &RestoreFromClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreFromClusterSnapshot", params, optFns, c.addOperationRestoreFromClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreFromClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreFromClusterSnapshotInput struct {

	// The identifier of the cluster that will be created from restoring the snapshot.
	// Constraints:
	//
	// * Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	// *
	// Alphabetic characters must be lowercase.
	//
	// * First character must be a letter.
	//
	// *
	// Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// * Must be unique
	// for all clusters within an Amazon Web Services account.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the snapshot from which to create the new cluster. This parameter
	// isn't case sensitive. Example: my-snapshot-id
	//
	// This member is required.
	SnapshotIdentifier *string

	// Reserved.
	AdditionalInfo *string

	// If true, major version upgrades can be applied during the maintenance window to
	// the Amazon Redshift engine that is running on the cluster. Default: true
	AllowVersionUpgrade *bool

	// The value represents how the cluster is configured to use AQUA (Advanced Query
	// Accelerator) after the cluster is restored. Possible values include the
	// following.
	//
	// * enabled - Use AQUA if it is available for the current Amazon Web
	// Services Region and Amazon Redshift node type.
	//
	// * disabled - Don't use AQUA.
	//
	// *
	// auto - Amazon Redshift determines whether to use AQUA.
	AquaConfigurationStatus types.AquaConfigurationStatus

	// The number of days that automated snapshots are retained. If the value is 0,
	// automated snapshots are disabled. Even if automated snapshots are disabled, you
	// can still create manual snapshots when you want with CreateClusterSnapshot. You
	// can't disable automated snapshots for RA3 node types. Set the automated
	// retention period from 1-35 days. Default: The value selected for the cluster
	// from which the snapshot was taken. Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *int32

	// The Amazon EC2 Availability Zone in which to restore the cluster. Default: A
	// random, system-chosen Availability Zone. Example: us-east-2a
	AvailabilityZone *string

	// The option to enable relocation for an Amazon Redshift cluster between
	// Availability Zones after the cluster is restored.
	AvailabilityZoneRelocation *bool

	// The name of the parameter group to be associated with this cluster. Default: The
	// default Amazon Redshift cluster parameter group. For information about the
	// default parameter group, go to Working with Amazon Redshift Parameter Groups
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html).
	// Constraints:
	//
	// * Must be 1 to 255 alphanumeric characters or hyphens.
	//
	// * First
	// character must be a letter.
	//
	// * Cannot end with a hyphen or contain two
	// consecutive hyphens.
	ClusterParameterGroupName *string

	// A list of security groups to be associated with this cluster. Default: The
	// default cluster security group for Amazon Redshift. Cluster security groups only
	// apply to clusters outside of VPCs.
	ClusterSecurityGroups []string

	// The name of the subnet group where you want to cluster restored. A snapshot of
	// cluster in VPC can be restored only in VPC. Therefore, you must provide subnet
	// group name where you want the cluster restored.
	ClusterSubnetGroupName *string

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the
	// cluster when the cluster was last modified while it was restored from a
	// snapshot.
	DefaultIamRoleArn *string

	// The elastic IP (EIP) address for the cluster.
	ElasticIp *string

	// An option that specifies whether to create the cluster with enhanced VPC routing
	// enabled. To create a cluster that uses enhanced VPC routing, the cluster must be
	// in a VPC. For more information, see Enhanced VPC Routing
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in
	// the Amazon Redshift Cluster Management Guide. If this option is true, enhanced
	// VPC routing is enabled. Default: false
	EnhancedVpcRouting *bool

	// Specifies the name of the HSM client certificate the Amazon Redshift cluster
	// uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string

	// Specifies the name of the HSM configuration that contains the information the
	// Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string

	// A list of Identity and Access Management (IAM) roles that can be used by the
	// cluster to access other Amazon Web Services services. You must supply the IAM
	// roles in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM
	// roles in a single request. A cluster can have up to 10 IAM roles associated at
	// any time.
	IamRoles []string

	// The Key Management Service (KMS) key ID of the encryption key that you want to
	// use to encrypt data in the cluster that you restore from a shared snapshot.
	KmsKeyId *string

	// The name of the maintenance track for the restored cluster. When you take a
	// snapshot, the snapshot inherits the MaintenanceTrack value from the cluster. The
	// snapshot might be on a different track than the cluster that was the source for
	// the snapshot. For example, suppose that you take a snapshot of a cluster that is
	// on the current track and then change the cluster to be on the trailing track. In
	// this case, the snapshot and the source cluster are on different tracks.
	MaintenanceTrackName *string

	// The default number of days to retain a manual snapshot. If the value is -1, the
	// snapshot is retained indefinitely. This setting doesn't change the retention
	// period of existing snapshots. The value must be either -1 or an integer between
	// 1 and 3,653.
	ManualSnapshotRetentionPeriod *int32

	// The node type that the restored cluster will be provisioned with. Default: The
	// node type of the cluster from which the snapshot was taken. You can modify this
	// if you are using any DS node type. In that case, you can choose to restore into
	// another DS node type of the same size. For example, you can restore ds1.8xlarge
	// into ds2.8xlarge, or ds1.xlarge into ds2.xlarge. If you have a DC instance type,
	// you must restore into that same instance type and size. In other words, you can
	// only restore a dc1.large instance type into another dc1.large instance type or
	// dc2.large instance type. You can't restore dc1.8xlarge to dc2.8xlarge. First
	// restore to a dc1.8xlarge cluster, then resize to a dc2.8large cluster. For more
	// information about node types, see  About Clusters and Nodes
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-about-clusters-and-nodes)
	// in the Amazon Redshift Cluster Management Guide.
	NodeType *string

	// The number of nodes specified when provisioning the restored cluster.
	NumberOfNodes *int32

	// The Amazon Web Services account used to create or copy the snapshot. Required if
	// you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string

	// The port number on which the cluster accepts connections. Default: The same port
	// as the original cluster. Constraints: Must be between 1115 and 65535.
	Port *int32

	// The weekly time range (in UTC) during which automated cluster maintenance can
	// occur. Format: ddd:hh24:mi-ddd:hh24:mi Default: The value selected for the
	// cluster from which the snapshot was taken. For more information about the time
	// blocks for each region, see Maintenance Windows
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows)
	// in Amazon Redshift Cluster Management Guide. Valid Days: Mon | Tue | Wed | Thu |
	// Fri | Sat | Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// If true, the cluster can be accessed from a public network.
	PubliclyAccessible *bool

	// The name of the cluster the source snapshot was created from. This parameter is
	// required if your IAM user has a policy containing a snapshot resource element
	// that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string

	// A unique identifier for the snapshot schedule.
	SnapshotScheduleIdentifier *string

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the
	// cluster. Default: The default VPC security group is associated with the cluster.
	// VPC security groups only apply to clusters in VPCs.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreFromClusterSnapshotOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreFromClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreFromClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreFromClusterSnapshot{}, middleware.After)
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
	if err = addOpRestoreFromClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreFromClusterSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreFromClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RestoreFromClusterSnapshot",
	}
}
