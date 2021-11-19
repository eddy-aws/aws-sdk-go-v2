// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Configuration of statistics that are allowed to be run on columns that contain
// detected entities. When undefined, no statistics will be computed on columns
// that contain detected entities.
type AllowedStatistics struct {

	// One or more column statistics to allow for columns that contain detected
	// entities.
	//
	// This member is required.
	Statistics []string

	noSmithyDocumentSerde
}

// Selector of a column from a dataset for profile job configuration. One selector
// includes either a column name or a regular expression.
type ColumnSelector struct {

	// The name of a column from a dataset.
	Name *string

	// A regular expression for selecting a column from a dataset.
	Regex *string

	noSmithyDocumentSerde
}

// Configuration for column evaluations for a profile job.
// ColumnStatisticsConfiguration can be used to select evaluations and override
// parameters of evaluations for particular columns.
type ColumnStatisticsConfiguration struct {

	// Configuration for evaluations. Statistics can be used to select evaluations and
	// override parameters of evaluations.
	//
	// This member is required.
	Statistics *StatisticsConfiguration

	// List of column selectors. Selectors can be used to select columns from the
	// dataset. When selectors are undefined, configuration will be applied to all
	// supported columns.
	Selectors []ColumnSelector

	noSmithyDocumentSerde
}

// Represents an individual condition that evaluates to true or false. Conditions
// are used with recipe actions. The action is only performed for column values
// where the condition evaluates to true. If a recipe requires more than one
// condition, then the recipe must specify multiple ConditionExpression elements.
// Each condition is applied to the rows in a dataset first, before the recipe
// action is performed.
type ConditionExpression struct {

	// A specific condition to apply to a recipe action. For more information, see
	// Recipe structure
	// (https://docs.aws.amazon.com/databrew/latest/dg/recipes.html#recipes.structure)
	// in the Glue DataBrew Developer Guide.
	//
	// This member is required.
	Condition *string

	// A column to apply this condition to.
	//
	// This member is required.
	TargetColumn *string

	// A value that the condition must evaluate to for the condition to succeed.
	Value *string

	noSmithyDocumentSerde
}

// Represents a set of options that define how DataBrew will read a comma-separated
// value (CSV) file when creating a dataset from that file.
type CsvOptions struct {

	// A single character that specifies the delimiter being used in the CSV file.
	Delimiter *string

	// A variable that specifies whether the first row in the file is parsed as the
	// header. If this value is false, column names are auto-generated.
	HeaderRow *bool

	noSmithyDocumentSerde
}

// Represents a set of options that define how DataBrew will write a
// comma-separated value (CSV) file.
type CsvOutputOptions struct {

	// A single character that specifies the delimiter used to create CSV job output.
	Delimiter *string

	noSmithyDocumentSerde
}

// Connection information for dataset input files stored in a database.
type DatabaseInputDefinition struct {

	// The Glue Connection that stores the connection information for the target
	// database.
	//
	// This member is required.
	GlueConnectionName *string

	// The table within the target database.
	DatabaseTableName *string

	// Custom SQL to run against the provided Glue connection. This SQL will be used as
	// the input for DataBrew projects and jobs.
	QueryString *string

	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
	// read input data, or write output from a job.
	TempDirectory *S3Location

	noSmithyDocumentSerde
}

// Represents a JDBC database output object which defines the output destination
// for a DataBrew recipe job to write into.
type DatabaseOutput struct {

	// Represents options that specify how and where DataBrew writes the database
	// output generated by recipe jobs.
	//
	// This member is required.
	DatabaseOptions *DatabaseTableOutputOptions

	// The Glue connection that stores the connection information for the target
	// database.
	//
	// This member is required.
	GlueConnectionName *string

	// The output mode to write into the database. Currently supported option:
	// NEW_TABLE.
	DatabaseOutputMode DatabaseOutputMode

	noSmithyDocumentSerde
}

// Represents options that specify how and where DataBrew writes the database
// output generated by recipe jobs.
type DatabaseTableOutputOptions struct {

	// A prefix for the name of a table DataBrew will create in the database.
	//
	// This member is required.
	TableName *string

	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
	// store intermediate results.
	TempDirectory *S3Location

	noSmithyDocumentSerde
}

// Represents how metadata stored in the Glue Data Catalog is defined in a DataBrew
// dataset.
type DataCatalogInputDefinition struct {

	// The name of a database in the Data Catalog.
	//
	// This member is required.
	DatabaseName *string

	// The name of a database table in the Data Catalog. This table corresponds to a
	// DataBrew dataset.
	//
	// This member is required.
	TableName *string

	// The unique identifier of the Amazon Web Services account that holds the Data
	// Catalog that stores the data.
	CatalogId *string

	// Represents an Amazon location where DataBrew can store intermediate results.
	TempDirectory *S3Location

	noSmithyDocumentSerde
}

// Represents options that specify how and where in the Glue Data Catalog DataBrew
// writes the output generated by recipe jobs.
type DataCatalogOutput struct {

	// The name of a database in the Data Catalog.
	//
	// This member is required.
	DatabaseName *string

	// The name of a table in the Data Catalog.
	//
	// This member is required.
	TableName *string

	// The unique identifier of the Amazon Web Services account that holds the Data
	// Catalog that stores the data.
	CatalogId *string

	// Represents options that specify how and where DataBrew writes the database
	// output generated by recipe jobs.
	DatabaseOptions *DatabaseTableOutputOptions

	// A value that, if true, means that any data in the location specified for output
	// is overwritten with new output. Not supported with DatabaseOptions.
	Overwrite bool

	// Represents options that specify how and where DataBrew writes the Amazon S3
	// output generated by recipe jobs.
	S3Options *S3TableOutputOptions

	noSmithyDocumentSerde
}

// Represents a dataset that can be processed by DataBrew.
type Dataset struct {

	// Information on how DataBrew can find the dataset, in either the Glue Data
	// Catalog or Amazon S3.
	//
	// This member is required.
	Input *Input

	// The unique name of the dataset.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon Web Services account that owns the dataset.
	AccountId *string

	// The date and time that the dataset was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the dataset.
	CreatedBy *string

	// The file format of a dataset that is created from an Amazon S3 file or folder.
	Format InputFormat

	// A set of options that define how DataBrew interprets the data in the dataset.
	FormatOptions *FormatOptions

	// The Amazon Resource Name (ARN) of the user who last modified the dataset.
	LastModifiedBy *string

	// The last modification date and time of the dataset.
	LastModifiedDate *time.Time

	// A set of options that defines how DataBrew interprets an Amazon S3 path of the
	// dataset.
	PathOptions *PathOptions

	// The unique Amazon Resource Name (ARN) for the dataset.
	ResourceArn *string

	// The location of the data for the dataset, either Amazon S3 or the Glue Data
	// Catalog.
	Source Source

	// Metadata tags that have been applied to the dataset.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents a dataset paramater that defines type and conditions for a parameter
// in the Amazon S3 path of the dataset.
type DatasetParameter struct {

	// The name of the parameter that is used in the dataset's Amazon S3 path.
	//
	// This member is required.
	Name *string

	// The type of the dataset parameter, can be one of a 'String', 'Number' or
	// 'Datetime'.
	//
	// This member is required.
	Type ParameterType

	// Optional boolean value that defines whether the captured value of this parameter
	// should be used to create a new column in a dataset.
	CreateColumn bool

	// Additional parameter options such as a format and a timezone. Required for
	// datetime parameters.
	DatetimeOptions *DatetimeOptions

	// The optional filter expression structure to apply additional matching criteria
	// to the parameter.
	Filter *FilterExpression

	noSmithyDocumentSerde
}

// Represents additional options for correct interpretation of datetime parameters
// used in the Amazon S3 path of a dataset.
type DatetimeOptions struct {

	// Required option, that defines the datetime format used for a date parameter in
	// the Amazon S3 path. Should use only supported datetime specifiers and separation
	// characters, all literal a-z or A-Z characters should be escaped with single
	// quotes. E.g. "MM.dd.yyyy-'at'-HH:mm".
	//
	// This member is required.
	Format *string

	// Optional value for a non-US locale code, needed for correct interpretation of
	// some date formats.
	LocaleCode *string

	// Optional value for a timezone offset of the datetime parameter value in the
	// Amazon S3 path. Shouldn't be used if Format for this parameter includes timezone
	// fields. If no offset specified, UTC is assumed.
	TimezoneOffset *string

	noSmithyDocumentSerde
}

// Configuration of entity detection for a profile job. When undefined, entity
// detection is disabled.
type EntityDetectorConfiguration struct {

	// Entity types to detect. Can be any of the following:
	//
	// * USA_SSN
	//
	// * EMAIL
	//
	// *
	// USA_ITIN
	//
	// * USA_PASSPORT_NUMBER
	//
	// * PHONE_NUMBER
	//
	// * USA_DRIVING_LICENSE
	//
	// *
	// BANK_ACCOUNT
	//
	// * CREDIT_CARD
	//
	// * IP_ADDRESS
	//
	// * MAC_ADDRESS
	//
	// * USA_DEA_NUMBER
	//
	// *
	// USA_HCPCS_CODE
	//
	// * USA_NATIONAL_PROVIDER_IDENTIFIER
	//
	// * USA_NATIONAL_DRUG_CODE
	//
	// *
	// USA_HEALTH_INSURANCE_CLAIM_NUMBER
	//
	// * USA_MEDICARE_BENEFICIARY_IDENTIFIER
	//
	// *
	// USA_CPT_CODE
	//
	// * PERSON_NAME
	//
	// * DATE
	//
	// The Entity type group USA_ALL is also
	// supported, and includes all of the above entity types except PERSON_NAME and
	// DATE.
	//
	// This member is required.
	EntityTypes []string

	// Configuration of statistics that are allowed to be run on columns that contain
	// detected entities. When undefined, no statistics will be computed on columns
	// that contain detected entities.
	AllowedStatistics []AllowedStatistics

	noSmithyDocumentSerde
}

// Represents a set of options that define how DataBrew will interpret a Microsoft
// Excel file when creating a dataset from that file.
type ExcelOptions struct {

	// A variable that specifies whether the first row in the file is parsed as the
	// header. If this value is false, column names are auto-generated.
	HeaderRow *bool

	// One or more sheet numbers in the Excel file that will be included in the
	// dataset.
	SheetIndexes []int32

	// One or more named sheets in the Excel file that will be included in the dataset.
	SheetNames []string

	noSmithyDocumentSerde
}

// Represents a limit imposed on number of Amazon S3 files that should be selected
// for a dataset from a connected Amazon S3 path.
type FilesLimit struct {

	// The number of Amazon S3 files to select.
	//
	// This member is required.
	MaxFiles int32

	// A criteria to use for Amazon S3 files sorting before their selection. By default
	// uses DESCENDING order, i.e. most recent files are selected first.
	// Anotherpossible value is ASCENDING.
	Order Order

	// A criteria to use for Amazon S3 files sorting before their selection. By default
	// uses LAST_MODIFIED_DATE as a sorting criteria. Currently it's the only allowed
	// value.
	OrderedBy OrderedBy

	noSmithyDocumentSerde
}

// Represents a structure for defining parameter conditions. Supported conditions
// are described here: Supported conditions for dynamic datasets
// (https://docs.aws.amazon.com/databrew/latest/dg/datasets.multiple-files.html#conditions.for.dynamic.datasets)
// in the Glue DataBrew Developer Guide.
type FilterExpression struct {

	// The expression which includes condition names followed by substitution
	// variables, possibly grouped and combined with other conditions. For example,
	// "(starts_with :prefix1 or starts_with :prefix2) and (ends_with :suffix1 or
	// ends_with :suffix2)". Substitution variables should start with ':' symbol.
	//
	// This member is required.
	Expression *string

	// The map of substitution variable names to their values used in this filter
	// expression.
	//
	// This member is required.
	ValuesMap map[string]string

	noSmithyDocumentSerde
}

// Represents a set of options that define the structure of either comma-separated
// value (CSV), Excel, or JSON input.
type FormatOptions struct {

	// Options that define how CSV input is to be interpreted by DataBrew.
	Csv *CsvOptions

	// Options that define how Excel input is to be interpreted by DataBrew.
	Excel *ExcelOptions

	// Options that define how JSON input is to be interpreted by DataBrew.
	Json *JsonOptions

	noSmithyDocumentSerde
}

// Represents information on how DataBrew can find data, in either the Glue Data
// Catalog or Amazon S3.
type Input struct {

	// The Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition *DataCatalogInputDefinition

	// Connection information for dataset input files stored in a database.
	DatabaseInputDefinition *DatabaseInputDefinition

	// Contains additional resource information needed for specific datasets.
	Metadata *Metadata

	// The Amazon S3 location where the data is stored.
	S3InputDefinition *S3Location

	noSmithyDocumentSerde
}

// Represents all of the attributes of a DataBrew job.
type Job struct {

	// The unique name of the job.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon Web Services account that owns the job.
	AccountId *string

	// The date and time that the job was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the job.
	CreatedBy *string

	// One or more artifacts that represent the Glue Data Catalog output from running
	// the job.
	DataCatalogOutputs []DataCatalogOutput

	// Represents a list of JDBC database output objects which defines the output
	// destination for a DataBrew recipe job to write into.
	DatabaseOutputs []DatabaseOutput

	// A dataset that the job is to process.
	DatasetName *string

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job output. For more information, see Encrypting data written by DataBrew jobs
	// (https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//
	// * SSE-KMS -
	// Server-side encryption with keys managed by KMS.
	//
	// * SSE-S3 - Server-side
	// encryption with keys managed by Amazon S3.
	EncryptionMode EncryptionMode

	// A sample configuration for profile jobs only, which determines the number of
	// rows on which the profile job is run. If a JobSample value isn't provided, the
	// default value is used. The default value is CUSTOM_ROWS for the mode parameter
	// and 20,000 for the size parameter.
	JobSample *JobSample

	// The Amazon Resource Name (ARN) of the user who last modified the job.
	LastModifiedBy *string

	// The modification date and time of the job.
	LastModifiedDate *time.Time

	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription LogSubscription

	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// One or more artifacts that represent output from running the job.
	Outputs []Output

	// The name of the project that the job is associated with.
	ProjectName *string

	// A set of steps that the job runs.
	RecipeReference *RecipeReference

	// The unique Amazon Resource Name (ARN) for the job.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn *string

	// Metadata tags that have been applied to the job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32

	// The job type of the job, which must be one of the following:
	//
	// * PROFILE - A job
	// to analyze a dataset, to determine its size, data types, data distribution, and
	// more.
	//
	// * RECIPE - A job to apply one or more transformations to a dataset.
	Type JobType

	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations []ValidationConfiguration

	noSmithyDocumentSerde
}

// Represents one run of a DataBrew job.
type JobRun struct {

	// The number of times that DataBrew has attempted to run the job.
	Attempt int32

	// The date and time when the job completed processing.
	CompletedOn *time.Time

	// One or more artifacts that represent the Glue Data Catalog output from running
	// the job.
	DataCatalogOutputs []DataCatalogOutput

	// Represents a list of JDBC database output objects which defines the output
	// destination for a DataBrew recipe job to write into.
	DatabaseOutputs []DatabaseOutput

	// The name of the dataset for the job to process.
	DatasetName *string

	// A message indicating an error (if any) that was encountered when the job ran.
	ErrorMessage *string

	// The amount of time, in seconds, during which a job run consumed resources.
	ExecutionTime int32

	// The name of the job being processed during this run.
	JobName *string

	// A sample configuration for profile jobs only, which determines the number of
	// rows on which the profile job is run. If a JobSample value isn't provided, the
	// default is used. The default value is CUSTOM_ROWS for the mode parameter and
	// 20,000 for the size parameter.
	JobSample *JobSample

	// The name of an Amazon CloudWatch log group, where the job writes diagnostic
	// messages when it runs.
	LogGroupName *string

	// The current status of Amazon CloudWatch logging for the job run.
	LogSubscription LogSubscription

	// One or more output artifacts from a job run.
	Outputs []Output

	// The set of steps processed by the job.
	RecipeReference *RecipeReference

	// The unique identifier of the job run.
	RunId *string

	// The Amazon Resource Name (ARN) of the user who initiated the job run.
	StartedBy *string

	// The date and time when the job run began.
	StartedOn *time.Time

	// The current state of the job run entity itself.
	State JobRunState

	// List of validation configurations that are applied to the profile job run.
	ValidationConfigurations []ValidationConfiguration

	noSmithyDocumentSerde
}

// A sample configuration for profile jobs only, which determines the number of
// rows on which the profile job is run. If a JobSample value isn't provided, the
// default is used. The default value is CUSTOM_ROWS for the mode parameter and
// 20,000 for the size parameter.
type JobSample struct {

	// A value that determines whether the profile job is run on the entire dataset or
	// a specified number of rows. This value must be one of the following:
	//
	// *
	// FULL_DATASET - The profile job is run on the entire dataset.
	//
	// * CUSTOM_ROWS -
	// The profile job is run on the number of rows specified in the Size parameter.
	Mode SampleMode

	// The Size parameter is only required when the mode is CUSTOM_ROWS. The profile
	// job is run on the specified number of rows. The maximum value for size is
	// Long.MAX_VALUE. Long.MAX_VALUE = 9223372036854775807
	Size *int64

	noSmithyDocumentSerde
}

// Represents the JSON-specific options that define how input is to be interpreted
// by Glue DataBrew.
type JsonOptions struct {

	// A value that specifies whether JSON input contains embedded new line characters.
	MultiLine bool

	noSmithyDocumentSerde
}

// Contains additional resource information needed for specific datasets.
type Metadata struct {

	// The Amazon Resource Name (ARN) associated with the dataset. Currently, DataBrew
	// only supports ARNs from Amazon AppFlow.
	SourceArn *string

	noSmithyDocumentSerde
}

// Represents options that specify how and where in Amazon S3 DataBrew writes the
// output generated by recipe jobs or profile jobs.
type Output struct {

	// The location in Amazon S3 where the job writes its output.
	//
	// This member is required.
	Location *S3Location

	// The compression algorithm used to compress the output text of the job.
	CompressionFormat CompressionFormat

	// The data format of the output of the job.
	Format OutputFormat

	// Represents options that define how DataBrew formats job output files.
	FormatOptions *OutputFormatOptions

	// A value that, if true, means that any data in the location specified for output
	// is overwritten with new output.
	Overwrite bool

	// The names of one or more partition columns for the output of the job.
	PartitionColumns []string

	noSmithyDocumentSerde
}

// Represents a set of options that define the structure of comma-separated (CSV)
// job output.
type OutputFormatOptions struct {

	// Represents a set of options that define the structure of comma-separated value
	// (CSV) job output.
	Csv *CsvOutputOptions

	noSmithyDocumentSerde
}

// Represents a set of options that define how DataBrew selects files for a given
// Amazon S3 path in a dataset.
type PathOptions struct {

	// If provided, this structure imposes a limit on a number of files that should be
	// selected.
	FilesLimit *FilesLimit

	// If provided, this structure defines a date range for matching Amazon S3 objects
	// based on their LastModifiedDate attribute in Amazon S3.
	LastModifiedDateCondition *FilterExpression

	// A structure that maps names of parameters used in the Amazon S3 path of a
	// dataset to their definitions.
	Parameters map[string]DatasetParameter

	noSmithyDocumentSerde
}

// Configuration for profile jobs. Configuration can be used to select columns, do
// evaluations, and override default parameters of evaluations. When configuration
// is undefined, the profile job will apply default settings to all supported
// columns.
type ProfileConfiguration struct {

	// List of configurations for column evaluations. ColumnStatisticsConfigurations
	// are used to select evaluations and override parameters of evaluations for
	// particular columns. When ColumnStatisticsConfigurations is undefined, the
	// profile job will profile all supported columns and run all supported
	// evaluations.
	ColumnStatisticsConfigurations []ColumnStatisticsConfiguration

	// Configuration for inter-column evaluations. Configuration can be used to select
	// evaluations and override parameters of evaluations. When configuration is
	// undefined, the profile job will run all supported inter-column evaluations.
	DatasetStatisticsConfiguration *StatisticsConfiguration

	// Configuration of entity detection for a profile job. When undefined, entity
	// detection is disabled.
	EntityDetectorConfiguration *EntityDetectorConfiguration

	// List of column selectors. ProfileColumns can be used to select columns from the
	// dataset. When ProfileColumns is undefined, the profile job will profile all
	// supported columns.
	ProfileColumns []ColumnSelector

	noSmithyDocumentSerde
}

// Represents all of the attributes of a DataBrew project.
type Project struct {

	// The unique name of a project.
	//
	// This member is required.
	Name *string

	// The name of a recipe that will be developed during a project session.
	//
	// This member is required.
	RecipeName *string

	// The ID of the Amazon Web Services account that owns the project.
	AccountId *string

	// The date and time that the project was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who crated the project.
	CreatedBy *string

	// The dataset that the project is to act upon.
	DatasetName *string

	// The Amazon Resource Name (ARN) of the user who last modified the project.
	LastModifiedBy *string

	// The last modification date and time for the project.
	LastModifiedDate *time.Time

	// The date and time when the project was opened.
	OpenDate *time.Time

	// The Amazon Resource Name (ARN) of the user that opened the project for use.
	OpenedBy *string

	// The Amazon Resource Name (ARN) for the project.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the role that will be assumed for this
	// project.
	RoleArn *string

	// The sample size and sampling type to apply to the data. If this parameter isn't
	// specified, then the sample consists of the first 500 rows from the dataset.
	Sample *Sample

	// Metadata tags that have been applied to the project.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents one or more actions to be performed on a DataBrew dataset.
type Recipe struct {

	// The unique name for the recipe.
	//
	// This member is required.
	Name *string

	// The date and time that the recipe was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the recipe.
	CreatedBy *string

	// The description of the recipe.
	Description *string

	// The Amazon Resource Name (ARN) of the user who last modified the recipe.
	LastModifiedBy *string

	// The last modification date and time of the recipe.
	LastModifiedDate *time.Time

	// The name of the project that the recipe is associated with.
	ProjectName *string

	// The Amazon Resource Name (ARN) of the user who published the recipe.
	PublishedBy *string

	// The date and time when the recipe was published.
	PublishedDate *time.Time

	// The identifier for the version for the recipe. Must be one of the following:
	//
	// *
	// Numeric version (X.Y) - X and Y stand for major and minor version numbers. The
	// maximum length of each is 6 digits, and neither can be negative values. Both X
	// and Y are required, and "0.0" isn't a valid version.
	//
	// * LATEST_WORKING - the
	// most recent valid version being developed in a DataBrew project.
	//
	// *
	// LATEST_PUBLISHED - the most recent published version.
	RecipeVersion *string

	// The Amazon Resource Name (ARN) for the recipe.
	ResourceArn *string

	// A list of steps that are defined by the recipe.
	Steps []RecipeStep

	// Metadata tags that have been applied to the recipe.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents a transformation and associated parameters that are used to apply a
// change to a DataBrew dataset. For more information, see Recipe actions reference
// (https://docs.aws.amazon.com/databrew/latest/dg/recipe-actions-reference.html).
type RecipeAction struct {

	// The name of a valid DataBrew transformation to be performed on the data.
	//
	// This member is required.
	Operation *string

	// Contextual parameters for the transformation.
	Parameters map[string]string

	noSmithyDocumentSerde
}

// Represents the name and version of a DataBrew recipe.
type RecipeReference struct {

	// The name of the recipe.
	//
	// This member is required.
	Name *string

	// The identifier for the version for the recipe.
	RecipeVersion *string

	noSmithyDocumentSerde
}

// Represents a single step from a DataBrew recipe to be performed.
type RecipeStep struct {

	// The particular action to be performed in the recipe step.
	//
	// This member is required.
	Action *RecipeAction

	// One or more conditions that must be met for the recipe step to succeed. All of
	// the conditions in the array must be met. In other words, all of the conditions
	// must be combined using a logical AND operation.
	ConditionExpressions []ConditionExpression

	noSmithyDocumentSerde
}

// Represents any errors encountered when attempting to delete multiple recipe
// versions.
type RecipeVersionErrorDetail struct {

	// The HTTP status code for the error.
	ErrorCode *string

	// The text of the error message.
	ErrorMessage *string

	// The identifier for the recipe version associated with this error.
	RecipeVersion *string

	noSmithyDocumentSerde
}

// Represents a single data quality requirement that should be validated in the
// scope of this dataset.
type Rule struct {

	// The expression which includes column references, condition names followed by
	// variable references, possibly grouped and combined with other conditions. For
	// example, (:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1
	// ends_with :suffix1 or :col1 ends_with :suffix2). Column and value references are
	// substitution variables that should start with the ':' symbol. Depending on the
	// context, substitution variables' values can be either an actual value or a
	// column name. These values are defined in the SubstitutionMap. If a
	// CheckExpression starts with a column reference, then ColumnSelectors in the rule
	// should be null. If ColumnSelectors has been defined, then there should be no
	// columnn reference in the left side of a condition, for example, is_between :val1
	// and :val2.
	//
	// This member is required.
	CheckExpression *string

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// List of column selectors. Selectors can be used to select columns using a name
	// or regular expression from the dataset. Rule will be applied to selected
	// columns.
	ColumnSelectors []ColumnSelector

	// A value that specifies whether the rule is disabled. Once a rule is disabled, a
	// profile job will not validate it during a job run. Default value is false.
	Disabled bool

	// The map of substitution variable names to their values used in a check
	// expression. Variable names should start with a ':' (colon). Variable values can
	// either be actual values or column names. To differentiate between the two,
	// column names should be enclosed in backticks, for example, ":col1": "`Column
	// A`".
	SubstitutionMap map[string]string

	// The threshold used with a non-aggregate check expression. Non-aggregate check
	// expressions will be applied to each row in a specific column, and the threshold
	// will be used to determine whether the validation succeeds.
	Threshold *Threshold

	noSmithyDocumentSerde
}

// Contains metadata about the ruleset.
type RulesetItem struct {

	// The name of the ruleset.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is
	// associated with.
	//
	// This member is required.
	TargetArn *string

	// The ID of the Amazon Web Services account that owns the ruleset.
	AccountId *string

	// The date and time that the ruleset was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the ruleset.
	CreatedBy *string

	// The description of the ruleset.
	Description *string

	// The Amazon Resource Name (ARN) of the user who last modified the ruleset.
	LastModifiedBy *string

	// The modification date and time of the ruleset.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) for the ruleset.
	ResourceArn *string

	// The number of rules that are defined in the ruleset.
	RuleCount int32

	// Metadata tags that have been applied to the ruleset.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
// read input data, or write output from a job.
type S3Location struct {

	// The Amazon S3 bucket name.
	//
	// This member is required.
	Bucket *string

	// The unique name of the object in the bucket.
	Key *string

	noSmithyDocumentSerde
}

// Represents options that specify how and where DataBrew writes the Amazon S3
// output generated by recipe jobs.
type S3TableOutputOptions struct {

	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
	// write output from a job.
	//
	// This member is required.
	Location *S3Location

	noSmithyDocumentSerde
}

// Represents the sample size and sampling type for DataBrew to use for interactive
// data analysis.
type Sample struct {

	// The way in which DataBrew obtains rows from a dataset.
	//
	// This member is required.
	Type SampleType

	// The number of rows in the sample.
	Size *int32

	noSmithyDocumentSerde
}

// Represents one or more dates and times when a job is to run.
type Schedule struct {

	// The name of the schedule.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon Web Services account that owns the schedule.
	AccountId *string

	// The date and time that the schedule was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the schedule.
	CreatedBy *string

	// The dates and times when the job is to run. For more information, see Cron
	// expressions (https://docs.aws.amazon.com/databrew/latest/dg/jobs.cron.html) in
	// the Glue DataBrew Developer Guide.
	CronExpression *string

	// A list of jobs to be run, according to the schedule.
	JobNames []string

	// The Amazon Resource Name (ARN) of the user who last modified the schedule.
	LastModifiedBy *string

	// The date and time when the schedule was last modified.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) of the schedule.
	ResourceArn *string

	// Metadata tags that have been applied to the schedule.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Override of a particular evaluation for a profile job.
type StatisticOverride struct {

	// A map that includes overrides of an evaluation’s parameters.
	//
	// This member is required.
	Parameters map[string]string

	// The name of an evaluation
	//
	// This member is required.
	Statistic *string

	noSmithyDocumentSerde
}

// Configuration of evaluations for a profile job. This configuration can be used
// to select evaluations and override the parameters of selected evaluations.
type StatisticsConfiguration struct {

	// List of included evaluations. When the list is undefined, all supported
	// evaluations will be included.
	IncludedStatistics []string

	// List of overrides for evaluations.
	Overrides []StatisticOverride

	noSmithyDocumentSerde
}

// The threshold used with a non-aggregate check expression. The non-aggregate
// check expression will be applied to each row in a specific column. Then the
// threshold will be used to determine whether the validation succeeds.
type Threshold struct {

	// The value of a threshold.
	//
	// This member is required.
	Value float64

	// The type of a threshold. Used for comparison of an actual count of rows that
	// satisfy the rule to the threshold value.
	Type ThresholdType

	// Unit of threshold value. Can be either a COUNT or PERCENTAGE of the full sample
	// size used for validation.
	Unit ThresholdUnit

	noSmithyDocumentSerde
}

// Configuration for data quality validation. Used to select the Rulesets and
// Validation Mode to be used in the profile job. When ValidationConfiguration is
// null, the profile job will run without data quality validation.
type ValidationConfiguration struct {

	// The Amazon Resource Name (ARN) for the ruleset to be validated in the profile
	// job. The TargetArn of the selected ruleset should be the same as the Amazon
	// Resource Name (ARN) of the dataset that is associated with the profile job.
	//
	// This member is required.
	RulesetArn *string

	// Mode of data quality validation. Default mode is “CHECK_ALL” which verifies all
	// rules defined in the selected ruleset.
	ValidationMode ValidationMode

	noSmithyDocumentSerde
}

// Represents the data being transformed during an action.
type ViewFrame struct {

	// The starting index for the range of columns to return in the view frame.
	//
	// This member is required.
	StartColumnIndex *int32

	// Controls if analytics computation is enabled or disabled. Enabled by default.
	Analytics AnalyticsMode

	// The number of columns to include in the view frame, beginning with the
	// StartColumnIndex value and ignoring any columns in the HiddenColumns list.
	ColumnRange *int32

	// A list of columns to hide in the view frame.
	HiddenColumns []string

	// The number of rows to include in the view frame, beginning with the
	// StartRowIndex value.
	RowRange *int32

	// The starting index for the range of rows to return in the view frame.
	StartRowIndex *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
