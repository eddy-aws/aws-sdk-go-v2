// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Information about a server's CPU.
type CPU struct {

	// The number of CPU cores.
	Cores int64

	// The model name of the CPU.
	ModelName *string

	noSmithyDocumentSerde
}

// Error in data replication.
type DataReplicationError struct {

	// Error in data replication.
	Error DataReplicationErrorString

	// Error in data replication.
	RawError *string

	noSmithyDocumentSerde
}

// Information about Data Replication
type DataReplicationInfo struct {

	// Error in data replication.
	DataReplicationError *DataReplicationError

	// Information about whether the data replication has been initiated.
	DataReplicationInitiation *DataReplicationInitiation

	// The state of the data replication.
	DataReplicationState DataReplicationState

	// An estimate of when the data replication will be completed.
	EtaDateTime *string

	// Data replication lag duration.
	LagDuration *string

	// The disks that should be replicated.
	ReplicatedDisks []DataReplicationInfoReplicatedDisk

	noSmithyDocumentSerde
}

// A disk that should be replicated.
type DataReplicationInfoReplicatedDisk struct {

	// The size of the replication backlog in bytes.
	BackloggedStorageBytes int64

	// The name of the device.
	DeviceName *string

	// The amount of data replicated so far in bytes.
	ReplicatedStorageBytes int64

	// The amount of data to be rescanned in bytes.
	RescannedStorageBytes int64

	// The total amount of data to be replicated in bytes.
	TotalStorageBytes int64

	noSmithyDocumentSerde
}

// Data replication initiation.
type DataReplicationInitiation struct {

	// The date and time of the next attempt to initiate data replication.
	NextAttemptDateTime *string

	// The date and time of the current attempt to initiate data replication.
	StartDateTime *string

	// The steps of the current attempt to initiate data replication.
	Steps []DataReplicationInitiationStep

	noSmithyDocumentSerde
}

// Data replication initiation step.
type DataReplicationInitiationStep struct {

	// The name of the step.
	Name DataReplicationInitiationStepName

	// The status of the step.
	Status DataReplicationInitiationStepStatus

	noSmithyDocumentSerde
}

// A set of filters by which to return Jobs.
type DescribeJobsRequestFilters struct {

	// The start date in a date range query.
	FromDate *string

	// An array of Job IDs that should be returned. An empty array means all jobs.
	JobIDs []string

	// The end date in a date range query.
	ToDate *string

	noSmithyDocumentSerde
}

// A set of filters by which to return Recovery Instances.
type DescribeRecoveryInstancesRequestFilters struct {

	// An array of Recovery Instance IDs that should be returned. An empty array means
	// all Recovery Instances.
	RecoveryInstanceIDs []string

	// An array of Source Server IDs for which associated Recovery Instances should be
	// returned.
	SourceServerIDs []string

	noSmithyDocumentSerde
}

// A set of filters by which to return Recovery Snapshots.
type DescribeRecoverySnapshotsRequestFilters struct {

	// The start date in a date range query.
	FromDateTime *string

	// The end date in a date range query.
	ToDateTime *string

	noSmithyDocumentSerde
}

// A set of filters by which to return Source Servers.
type DescribeSourceServersRequestFilters struct {

	// An ID that describes the hardware of the Source Server. This is either an EC2
	// instance id, a VMware uuid or a mac address.
	HardwareId *string

	// An array of Source Servers IDs that should be returned. An empty array means all
	// Source Servers.
	SourceServerIDs []string

	noSmithyDocumentSerde
}

// An object representing a data storage device on a server.
type Disk struct {

	// The amount of storage on the disk in bytes.
	Bytes int64

	// The disk or device name.
	DeviceName *string

	noSmithyDocumentSerde
}

// Hints used to uniquely identify a machine.
type IdentificationHints struct {

	// AWS Instance ID identification hint.
	AwsInstanceID *string

	// Fully Qualified Domain Name identification hint.
	Fqdn *string

	// Hostname identification hint.
	Hostname *string

	// vCenter VM path identification hint.
	VmWareUuid *string

	noSmithyDocumentSerde
}

// A job is an asynchronous workflow.
type Job struct {

	// The ID of the Job.
	//
	// This member is required.
	JobID *string

	// The ARN of a Job.
	Arn *string

	// The date and time of when the Job was created.
	CreationDateTime *string

	// The date and time of when the Job ended.
	EndDateTime *string

	// A string representing who initiated the Job.
	InitiatedBy InitiatedBy

	// A list of servers that the Job is acting upon.
	ParticipatingServers []ParticipatingServer

	// The status of the Job.
	Status JobStatus

	// A list of tags associated with the Job.
	Tags map[string]string

	// The type of the Job.
	Type JobType

	noSmithyDocumentSerde
}

// A log outputted by a Job.
type JobLog struct {

	// The event represents the type of a log.
	Event JobLogEvent

	// Metadata associated with a Job log.
	EventData *JobLogEventData

	// The date and time the log was taken.
	LogDateTime *string

	noSmithyDocumentSerde
}

// Metadata associated with a Job log.
type JobLogEventData struct {

	// The ID of a conversion server.
	ConversionServerID *string

	// A string representing a job error.
	RawError *string

	// The ID of a Source Server.
	SourceServerID *string

	// The ID of a Recovery Instance.
	TargetInstanceID *string

	noSmithyDocumentSerde
}

// Configuration of a machine's license.
type Licensing struct {

	// Whether to enable "Bring your own license" or not.
	OsByol *bool

	noSmithyDocumentSerde
}

// An object representing the Source Server Lifecycle.
type LifeCycle struct {

	// The date and time of when the Source Server was added to the service.
	AddedToServiceDateTime *string

	// The amount of time that the Source Server has been replicating for.
	ElapsedReplicationDuration *string

	// The date and time of the first byte that was replicated from the Source Server.
	FirstByteDateTime *string

	// An object containing information regarding the last launch of the Source Server.
	LastLaunch *LifeCycleLastLaunch

	// The date and time this Source Server was last seen by the service.
	LastSeenByServiceDateTime *string

	noSmithyDocumentSerde
}

// An object containing information regarding the last launch of a Source Server.
type LifeCycleLastLaunch struct {

	// An object containing information regarding the initiation of the last launch of
	// a Source Server.
	Initiated *LifeCycleLastLaunchInitiated

	noSmithyDocumentSerde
}

// An object containing information regarding the initiation of the last launch of
// a Source Server.
type LifeCycleLastLaunchInitiated struct {

	// The date and time the last Source Server launch was initiated.
	ApiCallDateTime *string

	// The ID of the Job that was used to last launch the Source Server.
	JobID *string

	// The Job type that was used to last launch the Source Server.
	Type LastLaunchType

	noSmithyDocumentSerde
}

// Network interface.
type NetworkInterface struct {

	// Network interface IPs.
	Ips []string

	// Whether this is the primary network interface.
	IsPrimary *bool

	// The MAC address of the network interface.
	MacAddress *string

	noSmithyDocumentSerde
}

// Operating System.
type OS struct {

	// The long name of the Operating System.
	FullString *string

	noSmithyDocumentSerde
}

// Represents a server participating in an asynchronous Job.
type ParticipatingServer struct {

	// The launch status of a participating server.
	LaunchStatus LaunchStatus

	// The Recovery Instance ID of a participating server.
	RecoveryInstanceID *string

	// The Source Server ID of a participating server.
	SourceServerID *string

	noSmithyDocumentSerde
}

// A rule in the Point in Time (PIT) policy representing when to take snapshots and
// how long to retain them for.
type PITPolicyRule struct {

	// How often, in the chosen units, a snapshot should be taken.
	//
	// This member is required.
	Interval int32

	// The duration to retain a snapshot for, in the chosen units.
	//
	// This member is required.
	RetentionDuration int32

	// The units used to measure the interval and retentionDuration.
	//
	// This member is required.
	Units PITPolicyRuleUnits

	// Whether this rule is enabled or not.
	Enabled *bool

	// The ID of the rule.
	RuleID int64

	noSmithyDocumentSerde
}

// A Recovery Instance is a replica of a Source Server running on EC2.
type RecoveryInstance struct {

	// The ARN of the Recovery Instance.
	Arn *string

	// The Data Replication Info of the Recovery Instance.
	DataReplicationInfo *RecoveryInstanceDataReplicationInfo

	// The EC2 instance ID of the Recovery Instance.
	Ec2InstanceID *string

	// The state of the EC2 instance for this Recovery Instance.
	Ec2InstanceState EC2InstanceState

	// An object representing failback related information of the Recovery Instance.
	Failback *RecoveryInstanceFailback

	// Whether this Recovery Instance was created for a drill or for an actual Recovery
	// event.
	IsDrill *bool

	// The ID of the Job that created the Recovery Instance.
	JobID *string

	// The date and time of the Point in Time (PIT) snapshot that this Recovery
	// Instance was launched from.
	PointInTimeSnapshotDateTime *string

	// The ID of the Recovery Instance.
	RecoveryInstanceID *string

	// Properties of the Recovery Instance machine.
	RecoveryInstanceProperties *RecoveryInstanceProperties

	// The Source Server ID that this Recovery Instance is associated with.
	SourceServerID *string

	// An array of tags that are associated with the Recovery Instance.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Error in data replication.
type RecoveryInstanceDataReplicationError struct {

	// Error in data replication.
	Error FailbackReplicationError

	// Error in data replication.
	RawError *string

	noSmithyDocumentSerde
}

// Information about Data Replication
type RecoveryInstanceDataReplicationInfo struct {

	// Information about Data Replication
	DataReplicationError *RecoveryInstanceDataReplicationError

	// Information about whether the data replication has been initiated.
	DataReplicationInitiation *RecoveryInstanceDataReplicationInitiation

	// The state of the data replication.
	DataReplicationState RecoveryInstanceDataReplicationState

	// An estimate of when the data replication will be completed.
	EtaDateTime *string

	// Data replication lag duration.
	LagDuration *string

	// The disks that should be replicated.
	ReplicatedDisks []RecoveryInstanceDataReplicationInfoReplicatedDisk

	noSmithyDocumentSerde
}

// A disk that should be replicated.
type RecoveryInstanceDataReplicationInfoReplicatedDisk struct {

	// The size of the replication backlog in bytes.
	BackloggedStorageBytes int64

	// The name of the device.
	DeviceName *string

	// The amount of data replicated so far in bytes.
	ReplicatedStorageBytes int64

	// The amount of data to be rescanned in bytes.
	RescannedStorageBytes int64

	// The total amount of data to be replicated in bytes.
	TotalStorageBytes int64

	noSmithyDocumentSerde
}

// Data replication initiation.
type RecoveryInstanceDataReplicationInitiation struct {

	// The date and time of the current attempt to initiate data replication.
	StartDateTime *string

	// The steps of the current attempt to initiate data replication.
	Steps []RecoveryInstanceDataReplicationInitiationStep

	noSmithyDocumentSerde
}

// Data replication initiation step.
type RecoveryInstanceDataReplicationInitiationStep struct {

	// The name of the step.
	Name RecoveryInstanceDataReplicationInitiationStepName

	// The status of the step.
	Status RecoveryInstanceDataReplicationInitiationStepStatus

	noSmithyDocumentSerde
}

// An object representing a block storage device on the Recovery Instance.
type RecoveryInstanceDisk struct {

	// The amount of storage on the disk in bytes.
	Bytes int64

	// The EBS Volume ID of this disk.
	EbsVolumeID *string

	// The internal device name of this disk. This is the name that is visible on the
	// machine itself and not from the EC2 console.
	InternalDeviceName *string

	noSmithyDocumentSerde
}

// An object representing failback related information of the Recovery Instance.
type RecoveryInstanceFailback struct {

	// The date and time the agent on the Recovery Instance was last seen by the
	// service.
	AgentLastSeenByServiceDateTime *string

	// The amount of time that the Recovery Instance has been replicating for.
	ElapsedReplicationDuration *string

	// The ID of the failback client that this Recovery Instance is associated with.
	FailbackClientID *string

	// The date and time that the failback client was last seen by the service.
	FailbackClientLastSeenByServiceDateTime *string

	// The date and time that the failback initiation started.
	FailbackInitiationTime *string

	// The Job ID of the last failback log for this Recovery Instance.
	FailbackJobID *string

	// Whether we are failing back to the original Source Server for this Recovery
	// Instance.
	FailbackToOriginalServer *bool

	// The date and time of the first byte that was replicated from the Recovery
	// Instance.
	FirstByteDateTime *string

	// The state of the failback process that this Recovery Instance is in.
	State FailbackState

	noSmithyDocumentSerde
}

// Properties of the Recovery Instance machine.
type RecoveryInstanceProperties struct {

	// An array of CPUs.
	Cpus []CPU

	// An array of disks.
	Disks []RecoveryInstanceDisk

	// Hints used to uniquely identify a machine.
	IdentificationHints *IdentificationHints

	// The date and time the Recovery Instance properties were last updated on.
	LastUpdatedDateTime *string

	// An array of network interfaces.
	NetworkInterfaces []NetworkInterface

	// Operating system.
	Os *OS

	// The amount of RAM in bytes.
	RamBytes int64

	noSmithyDocumentSerde
}

// A snapshot of a Source Server used during recovery.
type RecoverySnapshot struct {

	// The timestamp of when we expect the snapshot to be taken.
	//
	// This member is required.
	ExpectedTimestamp *string

	// The ID of the Recovery Snapshot.
	//
	// This member is required.
	SnapshotID *string

	// The ID of the Source Server that the snapshot was taken for.
	//
	// This member is required.
	SourceServerID *string

	// A list of EBS snapshots.
	EbsSnapshots []string

	// The actual timestamp that the snapshot was taken.
	Timestamp *string

	noSmithyDocumentSerde
}

// The configuration of a disk of the Source Server to be replicated.
type ReplicationConfigurationReplicatedDisk struct {

	// The name of the device.
	DeviceName *string

	// The requested number of I/O operations per second (IOPS).
	Iops int64

	// Whether to boot from this disk or not.
	IsBootDisk *bool

	// The Staging Disk EBS volume type to be used during replication.
	StagingDiskType ReplicationConfigurationReplicatedDiskStagingDiskType

	// The throughput to use for the EBS volume in MiB/s. This parameter is valid only
	// for gp3 volumes.
	Throughput int64

	noSmithyDocumentSerde
}

type ReplicationConfigurationTemplate struct {

	// The Replication Configuration Template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// The Replication Configuration Template ARN.
	Arn *string

	// Whether to associate the default Elastic Disaster Recovery Security group with
	// the Replication Configuration Template.
	AssociateDefaultSecurityGroup *bool

	// Configure bandwidth throttling for the outbound data transfer rate of the Source
	// Server in Mbps.
	BandwidthThrottling int64

	// Whether to create a Public IP for the Recovery Instance by default.
	CreatePublicIP *bool

	// The data plane routing mechanism that will be used for replication.
	DataPlaneRouting ReplicationConfigurationDataPlaneRouting

	// The Staging Disk EBS volume type to be used during replication.
	DefaultLargeStagingDiskType ReplicationConfigurationDefaultLargeStagingDiskType

	// The type of EBS encryption to be used during replication.
	EbsEncryption ReplicationConfigurationEbsEncryption

	// The ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn *string

	// The Point in time (PIT) policy to manage snapshots taken during replication.
	PitPolicy []PITPolicyRule

	// The instance type to be used for the replication server.
	ReplicationServerInstanceType *string

	// The security group IDs that will be used by the replication server.
	ReplicationServersSecurityGroupsIDs []string

	// The subnet to be used by the replication staging area.
	StagingAreaSubnetId *string

	// A set of tags to be associated with all resources created in the replication
	// staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	StagingAreaTags map[string]string

	// A set of tags to be associated with the Replication Configuration Template
	// resource.
	Tags map[string]string

	// Whether to use a dedicated Replication Server in the replication staging area.
	UseDedicatedReplicationServer *bool

	noSmithyDocumentSerde
}

// Properties of the Source Server machine.
type SourceProperties struct {

	// An array of CPUs.
	Cpus []CPU

	// An array of disks.
	Disks []Disk

	// Hints used to uniquely identify a machine.
	IdentificationHints *IdentificationHints

	// The date and time the Source Properties were last updated on.
	LastUpdatedDateTime *string

	// An array of network interfaces.
	NetworkInterfaces []NetworkInterface

	// Operating system.
	Os *OS

	// The amount of RAM in bytes.
	RamBytes int64

	// The recommended EC2 instance type that will be used when recovering the Source
	// Server.
	RecommendedInstanceType *string

	noSmithyDocumentSerde
}

type SourceServer struct {

	// The ARN of the Source Server.
	Arn *string

	// The Data Replication Info of the Source Server.
	DataReplicationInfo *DataReplicationInfo

	// The status of the last recovery launch of this Source Server.
	LastLaunchResult LastLaunchResult

	// The lifecycle information of this Source Server.
	LifeCycle *LifeCycle

	// The ID of the Recovery Instance associated with this Source Server.
	RecoveryInstanceId *string

	// The source properties of the Source Server.
	SourceProperties *SourceProperties

	// The ID of the Source Server.
	SourceServerID *string

	// The tags associated with the Source Server.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object representing the Source Server to recover.
type StartRecoveryRequestSourceServer struct {

	// The ID of the Source Server you want to recover.
	//
	// This member is required.
	SourceServerID *string

	// The ID of a Recovery Snapshot we want to recover from. Omit this field to launch
	// from the latest data by taking an on-demand snapshot.
	RecoverySnapshotID *string

	noSmithyDocumentSerde
}

// Validate exception field.
type ValidationExceptionField struct {

	// Validate exception field message.
	Message *string

	// Validate exception field name.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
