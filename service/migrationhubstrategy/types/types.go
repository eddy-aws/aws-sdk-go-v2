// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains the summary of anti-patterns and their severity.
type AntipatternSeveritySummary struct {

	// Contains the count of anti-patterns.
	Count *int32

	// Contains the severity of anti-patterns.
	Severity Severity

	noSmithyDocumentSerde
}

// Contains detailed information about an application component.
type ApplicationComponentDetail struct {

	// The status of analysis, if the application component has source code or an
	// associated database.
	AnalysisStatus SrcCodeOrDbAnalysisStatus

	// The S3 bucket name and the Amazon S3 key name for the anti-pattern report.
	AntipatternReportS3Object *S3Object

	// The status of the anti-pattern report generation.
	AntipatternReportStatus AntipatternReportStatus

	// The status message for the anti-pattern.
	AntipatternReportStatusMessage *string

	// The type of application component.
	AppType AppType

	// The ID of the server that the application component is running on.
	AssociatedServerId *string

	// Configuration details for the database associated with the application
	// component.
	DatabaseConfigDetail *DatabaseConfigDetail

	// The ID of the application component.
	Id *string

	// Indicates whether the application component has been included for server
	// recommendation or not.
	InclusionStatus InclusionStatus

	// The timestamp of when the application component was assessed.
	LastAnalyzedTimestamp *time.Time

	// A list of anti-pattern severity summaries.
	ListAntipatternSeveritySummary []AntipatternSeveritySummary

	// Set to true if the application component is running on multiple servers.
	MoreServerAssociationExists *bool

	// The name of application component.
	Name *string

	// OS driver.
	OsDriver *string

	// OS version.
	OsVersion *string

	// The top recommendation set for the application component.
	RecommendationSet *RecommendationSet

	// The application component subtype.
	ResourceSubType ResourceSubType

	// Details about the source code repository associated with the application
	// component.
	SourceCodeRepositories []SourceCodeRepository

	// A detailed description of the analysis status and any failure message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Contains information about a strategy recommendation for an application
// component.
type ApplicationComponentStrategy struct {

	// Set to true if the recommendation is set as preferred.
	IsPreferred *bool

	// Strategy recommendation for the application component.
	Recommendation *RecommendationSet

	// The recommendation status of a strategy for an application component.
	Status StrategyRecommendation

	noSmithyDocumentSerde
}

// Contains the summary of application components.
type ApplicationComponentSummary struct {

	// Contains the name of application types.
	AppType AppType

	// Contains the count of application type.
	Count *int32

	noSmithyDocumentSerde
}

// Application preferences that you specify.
type ApplicationPreferences struct {

	// Application preferences that you specify to prefer managed environment.
	ManagementPreference ManagementPreference

	noSmithyDocumentSerde
}

// Contains the summary of the assessment results.
type AssessmentSummary struct {

	// The Amazon S3 object containing the anti-pattern report.
	AntipatternReportS3Object *S3Object

	// The status of the anti-pattern report.
	AntipatternReportStatus AntipatternReportStatus

	// The status message of the anti-pattern report.
	AntipatternReportStatusMessage *string

	// The time the assessment was performed.
	LastAnalyzedTimestamp *time.Time

	// List of AntipatternSeveritySummary.
	ListAntipatternSeveritySummary []AntipatternSeveritySummary

	// List of ApplicationComponentStrategySummary.
	ListApplicationComponentStrategySummary []StrategySummary

	// List of ApplicationComponentSummary.
	ListApplicationComponentSummary []ApplicationComponentSummary

	// List of ServerStrategySummary.
	ListServerStrategySummary []StrategySummary

	// List of ServerSummary.
	ListServerSummary []ServerSummary

	noSmithyDocumentSerde
}

// Object containing details about applications as defined in Application Discovery
// Service.
type AssociatedApplication struct {

	// ID of the application as defined in Application Discovery Service.
	Id *string

	// Name of the application as defined in Application Discovery Service.
	Name *string

	noSmithyDocumentSerde
}

// Object containing the choice of application destination that you specify.
type AwsManagedResources struct {

	// The choice of application destination that you specify.
	//
	// This member is required.
	TargetDestination []AwsManagedTargetDestination

	noSmithyDocumentSerde
}

// Business goals that you specify.
type BusinessGoals struct {

	// Business goal to reduce license costs.
	LicenseCostReduction *int32

	// Business goal to modernize infrastructure by moving to cloud native
	// technologies.
	ModernizeInfrastructureWithCloudNativeTechnologies *int32

	// Business goal to reduce the operational overhead on the team by moving into
	// managed services.
	ReduceOperationalOverheadWithManagedServices *int32

	// Business goal to achieve migration at a fast pace.
	SpeedOfMigration *int32

	noSmithyDocumentSerde
}

// Process data collector that runs in the environment that you specify.
type Collector struct {

	// Indicates the health of a collector.
	CollectorHealth CollectorHealth

	// The ID of the collector.
	CollectorId *string

	// Current version of the collector that is running in the environment that you
	// specify.
	CollectorVersion *string

	// Hostname of the server that is hosting the collector.
	HostName *string

	// IP address of the server that is hosting the collector.
	IpAddress *string

	// Time when the collector last pinged the service.
	LastActivityTimeStamp *string

	// Time when the collector registered with the service.
	RegisteredTimeStamp *string

	noSmithyDocumentSerde
}

// Configuration information used for assessing databases.
type DatabaseConfigDetail struct {

	// AWS Secrets Manager key that holds the credentials that you use to connect to a
	// database.
	SecretName *string

	noSmithyDocumentSerde
}

// Preferences for migrating a database to AWS.
//
// The following types satisfy this interface:
//  DatabaseMigrationPreferenceMemberHeterogeneous
//  DatabaseMigrationPreferenceMemberHomogeneous
//  DatabaseMigrationPreferenceMemberNoPreference
type DatabaseMigrationPreference interface {
	isDatabaseMigrationPreference()
}

// Indicates whether you are interested in moving from one type of database to
// another. For example, from SQL Server to Amazon Aurora MySQL-Compatible Edition.
type DatabaseMigrationPreferenceMemberHeterogeneous struct {
	Value Heterogeneous

	noSmithyDocumentSerde
}

func (*DatabaseMigrationPreferenceMemberHeterogeneous) isDatabaseMigrationPreference() {}

// Indicates whether you are interested in moving to the same type of database into
// AWS. For example, from SQL Server in your environment to SQL Server on AWS.
type DatabaseMigrationPreferenceMemberHomogeneous struct {
	Value Homogeneous

	noSmithyDocumentSerde
}

func (*DatabaseMigrationPreferenceMemberHomogeneous) isDatabaseMigrationPreference() {}

// Indicated that you do not prefer heterogeneous or homogeneous.
type DatabaseMigrationPreferenceMemberNoPreference struct {
	Value NoDatabaseMigrationPreference

	noSmithyDocumentSerde
}

func (*DatabaseMigrationPreferenceMemberNoPreference) isDatabaseMigrationPreference() {}

// Preferences on managing your databases on AWS.
type DatabasePreferences struct {

	// Specifies whether you're interested in self-managed databases or databases
	// managed by AWS.
	DatabaseManagementPreference DatabaseManagementPreference

	// Specifies your preferred migration path.
	DatabaseMigrationPreference DatabaseMigrationPreference

	noSmithyDocumentSerde
}

// Detailed information about an assessment.
type DataCollectionDetails struct {

	// The time the assessment completes.
	CompletionTime *time.Time

	// The number of failed servers in the assessment.
	Failed *int32

	// The number of servers with the assessment status IN_PROGESS.
	InProgress *int32

	// The total number of servers in the assessment.
	Servers *int32

	// The start time of assessment.
	StartTime *time.Time

	// The status of the assessment.
	Status AssessmentStatus

	// The number of successful servers in the assessment.
	Success *int32

	noSmithyDocumentSerde
}

// The object containing information about distinct imports or groups for Strategy
// Recommendations.
type Group struct {

	// The key of the specific import group.
	Name GroupName

	// The value of the specific import group.
	Value *string

	noSmithyDocumentSerde
}

// The object containing details about heterogeneous database preferences.
type Heterogeneous struct {

	// The target database engine for heterogeneous database migration preference.
	//
	// This member is required.
	TargetDatabaseEngine []HeterogeneousTargetDatabaseEngine

	noSmithyDocumentSerde
}

// The object containing details about homogeneous database preferences.
type Homogeneous struct {

	// The target database engine for homogeneous database migration preferences.
	TargetDatabaseEngine []HomogeneousTargetDatabaseEngine

	noSmithyDocumentSerde
}

// Information about the import file tasks you request.
type ImportFileTaskInformation struct {

	// The time that the import task completes.
	CompletionTime *time.Time

	// The ID of the import file task.
	Id *string

	// The name of the import task given in StartImportFileTask.
	ImportName *string

	// The S3 bucket where the import file is located.
	InputS3Bucket *string

	// The Amazon S3 key name of the import file.
	InputS3Key *string

	// The number of records that failed to be imported.
	NumberOfRecordsFailed *int32

	// The number of records successfully imported.
	NumberOfRecordsSuccess *int32

	// Start time of the import task.
	StartTime *time.Time

	// Status of import file task.
	Status ImportFileTaskStatus

	// The S3 bucket name for status report of import task.
	StatusReportS3Bucket *string

	// The Amazon S3 key name for status report of import task. The report contains
	// details about whether each record imported successfully or why it did not.
	StatusReportS3Key *string

	noSmithyDocumentSerde
}

// Preferences for migrating an application to AWS.
//
// The following types satisfy this interface:
//  ManagementPreferenceMemberAwsManagedResources
//  ManagementPreferenceMemberNoPreference
//  ManagementPreferenceMemberSelfManageResources
type ManagementPreference interface {
	isManagementPreference()
}

// Indicates interest in solutions that are managed by AWS.
type ManagementPreferenceMemberAwsManagedResources struct {
	Value AwsManagedResources

	noSmithyDocumentSerde
}

func (*ManagementPreferenceMemberAwsManagedResources) isManagementPreference() {}

// No specific preference.
type ManagementPreferenceMemberNoPreference struct {
	Value NoManagementPreference

	noSmithyDocumentSerde
}

func (*ManagementPreferenceMemberNoPreference) isManagementPreference() {}

// Indicates interest in managing your own resources on AWS.
type ManagementPreferenceMemberSelfManageResources struct {
	Value SelfManageResources

	noSmithyDocumentSerde
}

func (*ManagementPreferenceMemberSelfManageResources) isManagementPreference() {}

// Information about the server's network for which the assessment was run.
type NetworkInfo struct {

	// Information about the name of the interface of the server for which the
	// assessment was run.
	//
	// This member is required.
	InterfaceName *string

	// Information about the IP address of the server for which the assessment was run.
	//
	// This member is required.
	IpAddress *string

	// Information about the MAC address of the server for which the assessment was
	// run.
	//
	// This member is required.
	MacAddress *string

	// Information about the subnet mask of the server for which the assessment was
	// run.
	//
	// This member is required.
	NetMask *string

	noSmithyDocumentSerde
}

// The object containing details about database migration preferences, when you
// have no particular preference.
type NoDatabaseMigrationPreference struct {

	// The target database engine for database migration preference that you specify.
	//
	// This member is required.
	TargetDatabaseEngine []TargetDatabaseEngine

	noSmithyDocumentSerde
}

// Object containing the choice of application destination that you specify.
type NoManagementPreference struct {

	// The choice of application destination that you specify.
	//
	// This member is required.
	TargetDestination []NoPreferenceTargetDestination

	noSmithyDocumentSerde
}

// Information about the operating system.
type OSInfo struct {

	// Information about the type of operating system.
	Type OSType

	// Information about the version of operating system.
	Version *string

	noSmithyDocumentSerde
}

// Rank of business goals based on priority.
type PrioritizeBusinessGoals struct {

	// Rank of business goals based on priority.
	BusinessGoals *BusinessGoals

	noSmithyDocumentSerde
}

// Contains detailed information about a recommendation report.
type RecommendationReportDetails struct {

	// The time that the recommendation report generation task completes.
	CompletionTime *time.Time

	// The S3 bucket where the report file is located.
	S3Bucket *string

	// The Amazon S3 key name of the report file.
	S3Keys []string

	// The time that the recommendation report generation task starts.
	StartTime *time.Time

	// The status of the recommendation report generation task.
	Status RecommendationReportStatus

	// The status message for recommendation report generation.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Contains a recommendation set.
type RecommendationSet struct {

	// The recommended strategy.
	Strategy Strategy

	// The recommended target destination.
	TargetDestination TargetDestination

	// The target destination for the recommendation set.
	TransformationTool *TransformationTool

	noSmithyDocumentSerde
}

// Contains the S3 bucket name and the Amazon S3 key name.
type S3Object struct {

	// The S3 bucket name.
	S3Bucket *string

	// The Amazon S3 key name.
	S3key *string

	noSmithyDocumentSerde
}

// Self-managed resources.
type SelfManageResources struct {

	// Self-managed resources target destination.
	//
	// This member is required.
	TargetDestination []SelfManageTargetDestination

	noSmithyDocumentSerde
}

// Detailed information about a server.
type ServerDetail struct {

	// The S3 bucket name and Amazon S3 key name for anti-pattern report.
	AntipatternReportS3Object *S3Object

	// The status of the anti-pattern report generation.
	AntipatternReportStatus AntipatternReportStatus

	// A message about the status of the anti-pattern report generation.
	AntipatternReportStatusMessage *string

	// A list of strategy summaries.
	ApplicationComponentStrategySummary []StrategySummary

	// The status of assessment for the server.
	DataCollectionStatus RunTimeAssessmentStatus

	// The server ID.
	Id *string

	// The timestamp of when the server was assessed.
	LastAnalyzedTimestamp *time.Time

	// A list of anti-pattern severity summaries.
	ListAntipatternSeveritySummary []AntipatternSeveritySummary

	// The name of the server.
	Name *string

	// A set of recommendations.
	RecommendationSet *RecommendationSet

	// The type of server.
	ServerType *string

	// A message about the status of data collection, which contains detailed
	// descriptions of any error messages.
	StatusMessage *string

	// System information about the server.
	SystemInfo *SystemInfo

	noSmithyDocumentSerde
}

// Contains information about a strategy recommendation for a server.
type ServerStrategy struct {

	// Set to true if the recommendation is set as preferred.
	IsPreferred *bool

	// The number of application components with this strategy recommendation running
	// on the server.
	NumberOfApplicationComponents *int32

	// Strategy recommendation for the server.
	Recommendation *RecommendationSet

	// The recommendation status of the strategy for the server.
	Status StrategyRecommendation

	noSmithyDocumentSerde
}

// Object containing details about the servers imported by Application Discovery
// Service
type ServerSummary struct {

	// Number of servers.
	Count *int32

	// Type of operating system for the servers.
	ServerOsType ServerOsType

	noSmithyDocumentSerde
}

// Object containing source code information that is linked to an application
// component.
type SourceCode struct {

	// The repository name for the source code.
	Location *string

	// The branch of the source code.
	SourceVersion *string

	// The type of repository to use for the source code.
	VersionControl VersionControl

	noSmithyDocumentSerde
}

// Object containing source code information that is linked to an application
// component.
type SourceCodeRepository struct {

	// The branch of the source code.
	Branch *string

	// The repository name for the source code.
	Repository *string

	// The type of repository to use for the source code.
	VersionControlType *string

	noSmithyDocumentSerde
}

// Information about all the available strategy options for migrating and
// modernizing an application component.
type StrategyOption struct {

	// Indicates if a specific strategy is preferred for the application component.
	IsPreferred *bool

	// Type of transformation. For example, Rehost, Replatform, and so on.
	Strategy Strategy

	// Destination information about where the application component can migrate to.
	// For example, EC2, ECS, and so on.
	TargetDestination TargetDestination

	// The name of the tool that can be used to transform an application component
	// using this strategy.
	ToolName TransformationToolName

	noSmithyDocumentSerde
}

// Object containing the summary of the strategy recommendations.
type StrategySummary struct {

	// The count of recommendations per strategy.
	Count *int32

	// The name of recommended strategy.
	Strategy Strategy

	noSmithyDocumentSerde
}

// Information about the server that hosts application components.
type SystemInfo struct {

	// CPU architecture type for the server.
	CpuArchitecture *string

	// File system type for the server.
	FileSystemType *string

	// Networking information related to a server.
	NetworkInfoList []NetworkInfo

	// Operating system corresponding to a server.
	OsInfo *OSInfo

	noSmithyDocumentSerde
}

// Information of the transformation tool that can be used to migrate and modernize
// the application.
type TransformationTool struct {

	// Description of the tool.
	Description *string

	// Name of the tool.
	Name TransformationToolName

	// URL for installing the tool.
	TranformationToolInstallationLink *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isDatabaseMigrationPreference() {}
func (*UnknownUnionMember) isManagementPreference()        {}
