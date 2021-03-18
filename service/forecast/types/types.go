// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Specifies a categorical hyperparameter and it's range of tunable values. This
// object is part of the ParameterRanges object.
type CategoricalParameterRange struct {

	// The name of the categorical hyperparameter to tune.
	//
	// This member is required.
	Name *string

	// A list of the tunable categories for the hyperparameter.
	//
	// This member is required.
	Values []string
}

// Specifies a continuous hyperparameter and it's range of tunable values. This
// object is part of the ParameterRanges object.
type ContinuousParameterRange struct {

	// The maximum tunable value of the hyperparameter.
	//
	// This member is required.
	MaxValue *float64

	// The minimum tunable value of the hyperparameter.
	//
	// This member is required.
	MinValue *float64

	// The name of the hyperparameter to tune.
	//
	// This member is required.
	Name *string

	// The scale that hyperparameter tuning uses to search the hyperparameter range.
	// Valid values: Auto Amazon Forecast hyperparameter tuning chooses the best scale
	// for the hyperparameter. Linear Hyperparameter tuning searches the values in the
	// hyperparameter range by using a linear scale. Logarithmic Hyperparameter tuning
	// searches the values in the hyperparameter range by using a logarithmic scale.
	// Logarithmic scaling works only for ranges that have values greater than 0.
	// ReverseLogarithmic hyperparameter tuning searches the values in the
	// hyperparameter range by using a reverse logarithmic scale. Reverse logarithmic
	// scaling works only for ranges that are entirely within the range 0 <= x < 1.0.
	// For information about choosing a hyperparameter scale, see Hyperparameter
	// Scaling
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html#scaling-type).
	// One of the following values:
	ScalingType ScalingType
}

// The destination for an export job. Provide an S3 path, an AWS Identity and
// Access Management (IAM) role that allows Amazon Forecast to access the location,
// and an AWS Key Management Service (KMS) key (optional).
type DataDestination struct {

	// The path to an Amazon Simple Storage Service (Amazon S3) bucket along with the
	// credentials to access the bucket.
	//
	// This member is required.
	S3Config *S3Config
}

// Provides a summary of the dataset group properties used in the ListDatasetGroups
// operation. To get the complete set of properties, call the DescribeDatasetGroup
// operation, and provide the DatasetGroupArn.
type DatasetGroupSummary struct {

	// When the dataset group was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// The name of the dataset group.
	DatasetGroupName *string

	// When the dataset group was created or last updated from a call to the
	// UpdateDatasetGroup operation. While the dataset group is being updated,
	// LastModificationTime is the current time of the ListDatasetGroups call.
	LastModificationTime *time.Time
}

// Provides a summary of the dataset import job properties used in the
// ListDatasetImportJobs operation. To get the complete set of properties, call the
// DescribeDatasetImportJob operation, and provide the DatasetImportJobArn.
type DatasetImportJobSummary struct {

	// When the dataset import job was created.
	CreationTime *time.Time

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data. The
	// training data must be stored in an Amazon S3 bucket. If encryption is used,
	// DataSource includes an AWS Key Management Service (KMS) key.
	DataSource *DataSource

	// The Amazon Resource Name (ARN) of the dataset import job.
	DatasetImportJobArn *string

	// The name of the dataset import job.
	DatasetImportJobName *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	// * CREATE_PENDING - The CreationTime.
	//
	// * CREATE_IN_PROGRESS - The
	// current timestamp.
	//
	// * CREATE_STOPPING - The current timestamp.
	//
	// * CREATE_STOPPED
	// - When the job stopped.
	//
	// * ACTIVE or CREATE_FAILED - When the job finished or
	// failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the dataset import job. States include:
	//
	// * ACTIVE
	//
	// *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// * CREATE_STOPPING, CREATE_STOPPED
	Status *string
}

// Provides a summary of the dataset properties used in the ListDatasets operation.
// To get the complete set of properties, call the DescribeDataset operation, and
// provide the DatasetArn.
type DatasetSummary struct {

	// When the dataset was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string

	// The name of the dataset.
	DatasetName *string

	// The dataset type.
	DatasetType DatasetType

	// The domain associated with the dataset.
	Domain Domain

	// When you create a dataset, LastModificationTime is the same as CreationTime.
	// While data is being imported to the dataset, LastModificationTime is the current
	// time of the ListDatasets call. After a CreateDatasetImportJob operation has
	// finished, LastModificationTime is when the import job completed or failed.
	LastModificationTime *time.Time
}

// The source of your training data, an AWS Identity and Access Management (IAM)
// role that allows Amazon Forecast to access the data and, optionally, an AWS Key
// Management Service (KMS) key. This object is submitted in the
// CreateDatasetImportJob request.
type DataSource struct {

	// The path to the training data stored in an Amazon Simple Storage Service (Amazon
	// S3) bucket along with the credentials to access the data.
	//
	// This member is required.
	S3Config *S3Config
}

// An AWS Key Management Service (KMS) key and an AWS Identity and Access
// Management (IAM) role that Amazon Forecast can assume to access the key. You can
// specify this optional object in the CreateDataset and CreatePredictor requests.
type EncryptionConfig struct {

	// The Amazon Resource Name (ARN) of the KMS key.
	//
	// This member is required.
	KMSKeyArn *string

	// The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS
	// key. Passing a role across AWS accounts is not allowed. If you pass a role that
	// isn't in your account, you get an InvalidInputException error.
	//
	// This member is required.
	RoleArn *string
}

// Provides detailed error metrics to evaluate the performance of a predictor. This
// object is part of the Metrics object.
type ErrorMetric struct {

	// The Forecast type used to compute WAPE and RMSE.
	ForecastType *string

	// The root-mean-square error (RMSE).
	RMSE *float64

	// The weighted absolute percentage error (WAPE).
	WAPE *float64
}

// Parameters that define how to split a dataset into training data and testing
// data, and the number of iterations to perform. These parameters are specified in
// the predefined algorithms but you can override them in the CreatePredictor
// request.
type EvaluationParameters struct {

	// The point from the end of the dataset where you want to split the data for model
	// training and testing (evaluation). Specify the value as the number of data
	// points. The default is the value of the forecast horizon. BackTestWindowOffset
	// can be used to mimic a past virtual forecast start date. This value must be
	// greater than or equal to the forecast horizon and less than half of the
	// TARGET_TIME_SERIES dataset length. ForecastHorizon <= BackTestWindowOffset < 1/2
	// * TARGET_TIME_SERIES dataset length
	BackTestWindowOffset *int32

	// The number of times to split the input data. The default is 1. Valid values are
	// 1 through 5.
	NumberOfBacktestWindows *int32
}

// The results of evaluating an algorithm. Returned as part of the
// GetAccuracyMetrics response.
type EvaluationResult struct {

	// The Amazon Resource Name (ARN) of the algorithm that was evaluated.
	AlgorithmArn *string

	// The array of test windows used for evaluating the algorithm. The
	// NumberOfBacktestWindows from the EvaluationParameters object determines the
	// number of windows in the array.
	TestWindows []WindowSummary
}

// Provides featurization (transformation) information for a dataset field. This
// object is part of the FeaturizationConfig object. For example: {
//
// "AttributeName": "demand",
//
//     FeaturizationPipeline [ {
//
//
// "FeaturizationMethodName": "filling",
//
//     "FeaturizationMethodParameters":
// {"aggregation": "avg", "backfill": "nan"}
//
//     } ]
//
//     }
type Featurization struct {

	// The name of the schema attribute that specifies the data field to be featurized.
	// Amazon Forecast supports the target field of the TARGET_TIME_SERIES and the
	// RELATED_TIME_SERIES datasets. For example, for the RETAIL domain, the target is
	// demand, and for the CUSTOM domain, the target is target_value. For more
	// information, see howitworks-missing-values.
	//
	// This member is required.
	AttributeName *string

	// An array of one FeaturizationMethod object that specifies the feature
	// transformation method.
	FeaturizationPipeline []FeaturizationMethod
}

// In a CreatePredictor operation, the specified algorithm trains a model using the
// specified dataset group. You can optionally tell the operation to modify data
// fields prior to training a model. These modifications are referred to as
// featurization. You define featurization using the FeaturizationConfig object.
// You specify an array of transformations, one for each field that you want to
// featurize. You then include the FeaturizationConfig object in your
// CreatePredictor request. Amazon Forecast applies the featurization to the
// TARGET_TIME_SERIES and RELATED_TIME_SERIES datasets before model training. You
// can create multiple featurization configurations. For example, you might call
// the CreatePredictor operation twice by specifying different featurization
// configurations.
type FeaturizationConfig struct {

	// The frequency of predictions in a forecast. Valid intervals are Y (Year), M
	// (Month), W (Week), D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes),
	// 10min (10 minutes), 5min (5 minutes), and 1min (1 minute). For example, "Y"
	// indicates every year and "5min" indicates every five minutes. The frequency must
	// be greater than or equal to the TARGET_TIME_SERIES dataset frequency. When a
	// RELATED_TIME_SERIES dataset is provided, the frequency must be equal to the
	// RELATED_TIME_SERIES dataset frequency.
	//
	// This member is required.
	ForecastFrequency *string

	// An array of featurization (transformation) information for the fields of a
	// dataset.
	Featurizations []Featurization

	// An array of dimension (field) names that specify how to group the generated
	// forecast. For example, suppose that you are generating a forecast for item sales
	// across all of your stores, and your dataset contains a store_id field. If you
	// want the sales forecast for each item by store, you would specify store_id as
	// the dimension. All forecast dimensions specified in the TARGET_TIME_SERIES
	// dataset don't need to be specified in the CreatePredictor request. All forecast
	// dimensions specified in the RELATED_TIME_SERIES dataset must be specified in the
	// CreatePredictor request.
	ForecastDimensions []string
}

// Provides information about the method that featurizes (transforms) a dataset
// field. The method is part of the FeaturizationPipeline of the Featurization
// object. The following is an example of how you specify a FeaturizationMethod
// object. {
//     "FeaturizationMethodName": "filling",
//
//
// "FeaturizationMethodParameters": {"aggregation": "sum", "middlefill": "zero",
// "backfill": "zero"}
//
//     }
type FeaturizationMethod struct {

	// The name of the method. The "filling" method is the only supported method.
	//
	// This member is required.
	FeaturizationMethodName FeaturizationMethodName

	// The method parameters (key-value pairs), which are a map of override parameters.
	// Specify these parameters to override the default values. Related Time Series
	// attributes do not accept aggregation parameters. The following list shows the
	// parameters and their valid values for the "filling" featurization method for a
	// Target Time Series dataset. Bold signifies the default value.
	//
	// * aggregation:
	// sum, avg, first, min, max
	//
	// * frontfill: none
	//
	// * middlefill: zero, nan (not a
	// number), value, median, mean, min, max
	//
	// * backfill: zero, nan, value, median,
	// mean, min, max
	//
	// The following list shows the parameters and their valid values
	// for a Related Time Series featurization method (there are no defaults):
	//
	// *
	// middlefill: zero, value, median, mean, min, max
	//
	// * backfill: zero, value,
	// median, mean, min, max
	//
	// * futurefill: zero, value, median, mean, min, max
	//
	// To
	// set a filling method to a specific value, set the fill parameter to value and
	// define the value in a corresponding _value parameter. For example, to set
	// backfilling to a value of 2, include the following: "backfill": "value" and
	// "backfill_value":"2".
	FeaturizationMethodParameters map[string]string
}

// Describes a filter for choosing a subset of objects. Each filter consists of a
// condition and a match statement. The condition is either IS or IS_NOT, which
// specifies whether to include or exclude the objects that match the statement,
// respectively. The match statement consists of a key and a value.
type Filter struct {

	// The condition to apply. To include the objects that match the statement, specify
	// IS. To exclude matching objects, specify IS_NOT.
	//
	// This member is required.
	Condition FilterConditionString

	// The name of the parameter to filter on.
	//
	// This member is required.
	Key *string

	// The value to match.
	//
	// This member is required.
	Value *string
}

// Provides a summary of the forecast export job properties used in the
// ListForecastExportJobs operation. To get the complete set of properties, call
// the DescribeForecastExportJob operation, and provide the listed
// ForecastExportJobArn.
type ForecastExportJobSummary struct {

	// When the forecast export job was created.
	CreationTime *time.Time

	// The path to the Amazon Simple Storage Service (Amazon S3) bucket where the
	// forecast is exported.
	Destination *DataDestination

	// The Amazon Resource Name (ARN) of the forecast export job.
	ForecastExportJobArn *string

	// The name of the forecast export job.
	ForecastExportJobName *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	// * CREATE_PENDING - The CreationTime.
	//
	// * CREATE_IN_PROGRESS - The
	// current timestamp.
	//
	// * CREATE_STOPPING - The current timestamp.
	//
	// * CREATE_STOPPED
	// - When the job stopped.
	//
	// * ACTIVE or CREATE_FAILED - When the job finished or
	// failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the forecast export job. States include:
	//
	// * ACTIVE
	//
	// *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * CREATE_STOPPING,
	// CREATE_STOPPED
	//
	// * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// The Status
	// of the forecast export job must be ACTIVE before you can access the forecast in
	// your S3 bucket.
	Status *string
}

// Provides a summary of the forecast properties used in the ListForecasts
// operation. To get the complete set of properties, call the DescribeForecast
// operation, and provide the ForecastArn that is listed in the summary.
type ForecastSummary struct {

	// When the forecast creation task was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group that provided the data used
	// to train the predictor.
	DatasetGroupArn *string

	// The ARN of the forecast.
	ForecastArn *string

	// The name of the forecast.
	ForecastName *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	// * CREATE_PENDING - The CreationTime.
	//
	// * CREATE_IN_PROGRESS - The
	// current timestamp.
	//
	// * CREATE_STOPPING - The current timestamp.
	//
	// * CREATE_STOPPED
	// - When the job stopped.
	//
	// * ACTIVE or CREATE_FAILED - When the job finished or
	// failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The ARN of the predictor used to generate the forecast.
	PredictorArn *string

	// The status of the forecast. States include:
	//
	// * ACTIVE
	//
	// * CREATE_PENDING,
	// CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * CREATE_STOPPING, CREATE_STOPPED
	//
	// *
	// DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	//
	// The Status of the forecast
	// must be ACTIVE before you can query or export the forecast.
	Status *string
}

// Configuration information for a hyperparameter tuning job. You specify this
// object in the CreatePredictor request. A hyperparameter is a parameter that
// governs the model training process. You set hyperparameters before training
// starts, unlike model parameters, which are determined during training. The
// values of the hyperparameters effect which values are chosen for the model
// parameters. In a hyperparameter tuning job, Amazon Forecast chooses the set of
// hyperparameter values that optimize a specified metric. Forecast accomplishes
// this by running many training jobs over a range of hyperparameter values. The
// optimum set of values depends on the algorithm, the training data, and the
// specified metric objective.
type HyperParameterTuningJobConfig struct {

	// Specifies the ranges of valid values for the hyperparameters.
	ParameterRanges *ParameterRanges
}

// The data used to train a predictor. The data includes a dataset group and any
// supplementary features. You specify this object in the CreatePredictor request.
type InputDataConfig struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	//
	// This member is required.
	DatasetGroupArn *string

	// An array of supplementary features. The only supported feature is a holiday
	// calendar.
	SupplementaryFeatures []SupplementaryFeature
}

// Specifies an integer hyperparameter and it's range of tunable values. This
// object is part of the ParameterRanges object.
type IntegerParameterRange struct {

	// The maximum tunable value of the hyperparameter.
	//
	// This member is required.
	MaxValue *int32

	// The minimum tunable value of the hyperparameter.
	//
	// This member is required.
	MinValue *int32

	// The name of the hyperparameter to tune.
	//
	// This member is required.
	Name *string

	// The scale that hyperparameter tuning uses to search the hyperparameter range.
	// Valid values: Auto Amazon Forecast hyperparameter tuning chooses the best scale
	// for the hyperparameter. Linear Hyperparameter tuning searches the values in the
	// hyperparameter range by using a linear scale. Logarithmic Hyperparameter tuning
	// searches the values in the hyperparameter range by using a logarithmic scale.
	// Logarithmic scaling works only for ranges that have values greater than 0.
	// ReverseLogarithmic Not supported for IntegerParameterRange. Reverse logarithmic
	// scaling works only for ranges that are entirely within the range 0 <= x < 1.0.
	// For information about choosing a hyperparameter scale, see Hyperparameter
	// Scaling
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html#scaling-type).
	// One of the following values:
	ScalingType ScalingType
}

// Provides metrics that are used to evaluate the performance of a predictor. This
// object is part of the WindowSummary object.
type Metrics struct {

	// Provides detailed error metrics on forecast type, root-mean square-error (RMSE),
	// and weighted average percentage error (WAPE).
	ErrorMetrics []ErrorMetric

	// The root-mean-square error (RMSE).
	//
	// Deprecated: This property is deprecated, please refer to ErrorMetrics for both
	// RMSE and WAPE
	RMSE *float64

	// An array of weighted quantile losses. Quantiles divide a probability
	// distribution into regions of equal probability. The distribution in this case is
	// the loss function.
	WeightedQuantileLosses []WeightedQuantileLoss
}

// Specifies the categorical, continuous, and integer hyperparameters, and their
// ranges of tunable values. The range of tunable values determines which values
// that a hyperparameter tuning job can choose for the specified hyperparameter.
// This object is part of the HyperParameterTuningJobConfig object.
type ParameterRanges struct {

	// Specifies the tunable range for each categorical hyperparameter.
	CategoricalParameterRanges []CategoricalParameterRange

	// Specifies the tunable range for each continuous hyperparameter.
	ContinuousParameterRanges []ContinuousParameterRange

	// Specifies the tunable range for each integer hyperparameter.
	IntegerParameterRanges []IntegerParameterRange
}

// Provides a summary of the predictor backtest export job properties used in the
// ListPredictorBacktestExportJobs operation. To get a complete set of properties,
// call the DescribePredictorBacktestExportJob operation, and provide the listed
// PredictorBacktestExportJobArn.
type PredictorBacktestExportJobSummary struct {

	// When the predictor backtest export job was created.
	CreationTime *time.Time

	// The destination for an export job. Provide an S3 path, an AWS Identity and
	// Access Management (IAM) role that allows Amazon Forecast to access the location,
	// and an AWS Key Management Service (KMS) key (optional).
	Destination *DataDestination

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	// * CREATE_PENDING - The CreationTime.
	//
	// * CREATE_IN_PROGRESS - The
	// current timestamp.
	//
	// * CREATE_STOPPING - The current timestamp.
	//
	// * CREATE_STOPPED
	// - When the job stopped.
	//
	// * ACTIVE or CREATE_FAILED - When the job finished or
	// failed.
	LastModificationTime *time.Time

	// Information about any errors that may have occurred during the backtest export.
	Message *string

	// The Amazon Resource Name (ARN) of the predictor backtest export job.
	PredictorBacktestExportJobArn *string

	// The name of the predictor backtest export job.
	PredictorBacktestExportJobName *string

	// The status of the predictor backtest export job. States include:
	//
	// * ACTIVE
	//
	// *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * CREATE_STOPPING,
	// CREATE_STOPPED
	//
	// * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	Status *string
}

// The algorithm used to perform a backtest and the status of those tests.
type PredictorExecution struct {

	// The ARN of the algorithm used to test the predictor.
	AlgorithmArn *string

	// An array of test windows used to evaluate the algorithm. The
	// NumberOfBacktestWindows from the object determines the number of windows in the
	// array.
	TestWindows []TestWindowSummary
}

// Contains details on the backtests performed to evaluate the accuracy of the
// predictor. The tests are returned in descending order of accuracy, with the most
// accurate backtest appearing first. You specify the number of backtests to
// perform when you call the operation.
type PredictorExecutionDetails struct {

	// An array of the backtests performed to evaluate the accuracy of the predictor
	// against a particular algorithm. The NumberOfBacktestWindows from the object
	// determines the number of windows in the array.
	PredictorExecutions []PredictorExecution
}

// Provides a summary of the predictor properties that are used in the
// ListPredictors operation. To get the complete set of properties, call the
// DescribePredictor operation, and provide the listed PredictorArn.
type PredictorSummary struct {

	// When the model training task was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group that contains the data used
	// to train the predictor.
	DatasetGroupArn *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	// * CREATE_PENDING - The CreationTime.
	//
	// * CREATE_IN_PROGRESS - The
	// current timestamp.
	//
	// * CREATE_STOPPING - The current timestamp.
	//
	// * CREATE_STOPPED
	// - When the job stopped.
	//
	// * ACTIVE or CREATE_FAILED - When the job finished or
	// failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The ARN of the predictor.
	PredictorArn *string

	// The name of the predictor.
	PredictorName *string

	// The status of the predictor. States include:
	//
	// * ACTIVE
	//
	// * CREATE_PENDING,
	// CREATE_IN_PROGRESS, CREATE_FAILED
	//
	// * DELETE_PENDING, DELETE_IN_PROGRESS,
	// DELETE_FAILED
	//
	// * CREATE_STOPPING, CREATE_STOPPED
	//
	// The Status of the predictor
	// must be ACTIVE before you can use the predictor to create a forecast.
	Status *string
}

// The path to the file(s) in an Amazon Simple Storage Service (Amazon S3) bucket,
// and an AWS Identity and Access Management (IAM) role that Amazon Forecast can
// assume to access the file(s). Optionally, includes an AWS Key Management Service
// (KMS) key. This object is part of the DataSource object that is submitted in the
// CreateDatasetImportJob request, and part of the DataDestination object.
type S3Config struct {

	// The path to an Amazon Simple Storage Service (Amazon S3) bucket or file(s) in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Path *string

	// The ARN of the AWS Identity and Access Management (IAM) role that Amazon
	// Forecast can assume to access the Amazon S3 bucket or files. If you provide a
	// value for the KMSKeyArn key, the role must allow access to the key. Passing a
	// role across AWS accounts is not allowed. If you pass a role that isn't in your
	// account, you get an InvalidInputException error.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key.
	KMSKeyArn *string
}

// Defines the fields of a dataset. You specify this object in the CreateDataset
// request.
type Schema struct {

	// An array of attributes specifying the name and type of each field in a dataset.
	Attributes []SchemaAttribute
}

// An attribute of a schema, which defines a dataset field. A schema attribute is
// required for every field in a dataset. The Schema object contains an array of
// SchemaAttribute objects.
type SchemaAttribute struct {

	// The name of the dataset field.
	AttributeName *string

	// The data type of the field.
	AttributeType AttributeType
}

// Provides statistics for each data field imported into to an Amazon Forecast
// dataset with the CreateDatasetImportJob operation.
type Statistics struct {

	// For a numeric field, the average value in the field.
	Avg *float64

	// The number of values in the field.
	Count *int32

	// The number of distinct values in the field.
	CountDistinct *int32

	// The number of NAN (not a number) values in the field.
	CountNan *int32

	// The number of null values in the field.
	CountNull *int32

	// For a numeric field, the maximum value in the field.
	Max *string

	// For a numeric field, the minimum value in the field.
	Min *string

	// For a numeric field, the standard deviation.
	Stddev *float64
}

// Describes a supplementary feature of a dataset group. This object is part of the
// InputDataConfig object. Forecast supports the Weather Index and Holidays
// built-in featurizations. Weather Index The Amazon Forecast Weather Index is a
// built-in featurization that incorporates historical and projected weather
// information into your model. The Weather Index supplements your datasets with
// over two years of historical weather data and up to 14 days of projected weather
// data. For more information, see Amazon Forecast Weather Index
// (https://docs.aws.amazon.com/forecast/latest/dg/weather.html). Holidays Holidays
// is a built-in featurization that incorporates a feature-engineered dataset of
// national holiday information into your model. It provides native support for the
// holiday calendars of 66 countries. To view the holiday calendars, refer to the
// Jollyday (http://jollyday.sourceforge.net/data.html) library. For more
// information, see Holidays Featurization
// (https://docs.aws.amazon.com/forecast/latest/dg/holidays.html).
type SupplementaryFeature struct {

	// The name of the feature. Valid values: "holiday" and "weather".
	//
	// This member is required.
	Name *string

	// Weather Index To enable the Weather Index, set the value to "true" Holidays To
	// enable Holidays, specify a country with one of the following two-letter country
	// codes:
	//
	// * "AL" - ALBANIA
	//
	// * "AR" - ARGENTINA
	//
	// * "AT" - AUSTRIA
	//
	// * "AU" -
	// AUSTRALIA
	//
	// * "BA" - BOSNIA HERZEGOVINA
	//
	// * "BE" - BELGIUM
	//
	// * "BG" - BULGARIA
	//
	// *
	// "BO" - BOLIVIA
	//
	// * "BR" - BRAZIL
	//
	// * "BY" - BELARUS
	//
	// * "CA" - CANADA
	//
	// * "CL" -
	// CHILE
	//
	// * "CO" - COLOMBIA
	//
	// * "CR" - COSTA RICA
	//
	// * "HR" - CROATIA
	//
	// * "CZ" - CZECH
	// REPUBLIC
	//
	// * "DK" - DENMARK
	//
	// * "EC" - ECUADOR
	//
	// * "EE" - ESTONIA
	//
	// * "ET" -
	// ETHIOPIA
	//
	// * "FI" - FINLAND
	//
	// * "FR" - FRANCE
	//
	// * "DE" - GERMANY
	//
	// * "GR" -
	// GREECE
	//
	// * "HU" - HUNGARY
	//
	// * "IS" - ICELAND
	//
	// * "IN" - INDIA
	//
	// * "IE" - IRELAND
	//
	// *
	// "IT" - ITALY
	//
	// * "JP" - JAPAN
	//
	// * "KZ" - KAZAKHSTAN
	//
	// * "KR" - KOREA
	//
	// * "LV" -
	// LATVIA
	//
	// * "LI" - LIECHTENSTEIN
	//
	// * "LT" - LITHUANIA
	//
	// * "LU" - LUXEMBOURG
	//
	// * "MK"
	// - MACEDONIA
	//
	// * "MT" - MALTA
	//
	// * "MX" - MEXICO
	//
	// * "MD" - MOLDOVA
	//
	// * "ME" -
	// MONTENEGRO
	//
	// * "NL" - NETHERLANDS
	//
	// * "NZ" - NEW ZEALAND
	//
	// * "NI" - NICARAGUA
	//
	// *
	// "NG" - NIGERIA
	//
	// * "NO" - NORWAY
	//
	// * "PA" - PANAMA
	//
	// * "PY" - PARAGUAY
	//
	// * "PE" -
	// PERU
	//
	// * "PL" - POLAND
	//
	// * "PT" - PORTUGAL
	//
	// * "RO" - ROMANIA
	//
	// * "RU" - RUSSIA
	//
	// *
	// "RS" - SERBIA
	//
	// * "SK" - SLOVAKIA
	//
	// * "SI" - SLOVENIA
	//
	// * "ZA" - SOUTH AFRICA
	//
	// *
	// "ES" - SPAIN
	//
	// * "SE" - SWEDEN
	//
	// * "CH" - SWITZERLAND
	//
	// * "UA" - UKRAINE
	//
	// * "AE" -
	// UNITED ARAB EMIRATES
	//
	// * "US" - UNITED STATES
	//
	// * "UK" - UNITED KINGDOM
	//
	// * "UY" -
	// URUGUAY
	//
	// * "VE" - VENEZUELA
	//
	// This member is required.
	Value *string
}

// The optional metadata that you apply to a resource to help you categorize and
// organize them. Each tag consists of a key and an optional value, both of which
// you define. The following basic restrictions apply to tags:
//
// * Maximum number of
// tags per resource - 50.
//
// * For each resource, each tag key must be unique, and
// each tag key can have only one value.
//
// * Maximum key length - 128 Unicode
// characters in UTF-8.
//
// * Maximum value length - 256 Unicode characters in
// UTF-8.
//
// * If your tagging schema is used across multiple services and resources,
// remember that other services may have restrictions on allowed characters.
// Generally allowed characters are: letters, numbers, and spaces representable in
// UTF-8, and the following characters: + - = . _ : / @.
//
// * Tag keys and values are
// case sensitive.
//
// * Do not use aws:, AWS:, or any upper or lowercase combination
// of such as a prefix for keys as it is reserved for AWS use. You cannot edit or
// delete tag keys with this prefix. Values can have this prefix. If a tag value
// has aws as its prefix but the key does not, then Forecast considers it to be a
// user tag and will count against the limit of 50 tags. Tags with only the key
// prefix of aws do not count against your tags per resource limit.
type Tag struct {

	// One part of a key-value pair that makes up a tag. A key is a general label that
	// acts like a category for more specific tag values.
	//
	// This member is required.
	Key *string

	// The optional part of a key-value pair that makes up a tag. A value acts as a
	// descriptor within a tag category (key).
	//
	// This member is required.
	Value *string
}

// The status, start time, and end time of a backtest, as well as a failure reason
// if applicable.
type TestWindowSummary struct {

	// If the test failed, the reason why it failed.
	Message *string

	// The status of the test. Possible status values are:
	//
	// * ACTIVE
	//
	// *
	// CREATE_IN_PROGRESS
	//
	// * CREATE_FAILED
	Status *string

	// The time at which the test ended.
	TestWindowEnd *time.Time

	// The time at which the test began.
	TestWindowStart *time.Time
}

// The weighted loss value for a quantile. This object is part of the Metrics
// object.
type WeightedQuantileLoss struct {

	// The difference between the predicted value and the actual value over the
	// quantile, weighted (normalized) by dividing by the sum over all quantiles.
	LossValue *float64

	// The quantile. Quantiles divide a probability distribution into regions of equal
	// probability. For example, if the distribution was divided into 5 regions of
	// equal probability, the quantiles would be 0.2, 0.4, 0.6, and 0.8.
	Quantile *float64
}

// The metrics for a time range within the evaluation portion of a dataset. This
// object is part of the EvaluationResult object. The TestWindowStart and
// TestWindowEnd parameters are determined by the BackTestWindowOffset parameter of
// the EvaluationParameters object.
type WindowSummary struct {

	// The type of evaluation.
	//
	// * SUMMARY - The average metrics across all windows.
	//
	// *
	// COMPUTED - The metrics for the specified window.
	EvaluationType EvaluationType

	// The number of data points within the window.
	ItemCount *int32

	// Provides metrics used to evaluate the performance of a predictor.
	Metrics *Metrics

	// The timestamp that defines the end of the window.
	TestWindowEnd *time.Time

	// The timestamp that defines the start of the window.
	TestWindowStart *time.Time
}
