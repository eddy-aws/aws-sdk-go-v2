// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Architecture string

// Enum values for Architecture
const (
	ArchitectureX8664 Architecture = "X86_64"
	ArchitectureArm64 Architecture = "ARM64"
	ArchitectureArmhf Architecture = "ARMHF"
)

// Values returns all known values for Architecture. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Architecture) Values() []Architecture {
	return []Architecture{
		"X86_64",
		"ARM64",
		"ARMHF",
	}
}

type DeploymentJobErrorCode string

// Enum values for DeploymentJobErrorCode
const (
	DeploymentJobErrorCodeResourceNotFound                    DeploymentJobErrorCode = "ResourceNotFound"
	DeploymentJobErrorCodeEnvironmentSetupError               DeploymentJobErrorCode = "EnvironmentSetupError"
	DeploymentJobErrorCodeEtagMismatch                        DeploymentJobErrorCode = "EtagMismatch"
	DeploymentJobErrorCodeFailureThresholdBreached            DeploymentJobErrorCode = "FailureThresholdBreached"
	DeploymentJobErrorCodeRobotDeploymentAborted              DeploymentJobErrorCode = "RobotDeploymentAborted"
	DeploymentJobErrorCodeRobotDeploymentNoResponse           DeploymentJobErrorCode = "RobotDeploymentNoResponse"
	DeploymentJobErrorCodeRobotAgentConnectionTimeout         DeploymentJobErrorCode = "RobotAgentConnectionTimeout"
	DeploymentJobErrorCodeGreengrassDeploymentFailed          DeploymentJobErrorCode = "GreengrassDeploymentFailed"
	DeploymentJobErrorCodeInvalidGreengrassGroup              DeploymentJobErrorCode = "InvalidGreengrassGroup"
	DeploymentJobErrorCodeMissingRobotArchitecture            DeploymentJobErrorCode = "MissingRobotArchitecture"
	DeploymentJobErrorCodeMissingRobotApplicationArchitecture DeploymentJobErrorCode = "MissingRobotApplicationArchitecture"
	DeploymentJobErrorCodeMissingRobotDeploymentResource      DeploymentJobErrorCode = "MissingRobotDeploymentResource"
	DeploymentJobErrorCodeGreengrassGroupVersionDoesNotExist  DeploymentJobErrorCode = "GreengrassGroupVersionDoesNotExist"
	DeploymentJobErrorCodeLambdaDeleted                       DeploymentJobErrorCode = "LambdaDeleted"
	DeploymentJobErrorCodeExtractingBundleFailure             DeploymentJobErrorCode = "ExtractingBundleFailure"
	DeploymentJobErrorCodePreLaunchFileFailure                DeploymentJobErrorCode = "PreLaunchFileFailure"
	DeploymentJobErrorCodePostLaunchFileFailure               DeploymentJobErrorCode = "PostLaunchFileFailure"
	DeploymentJobErrorCodeBadPermissionError                  DeploymentJobErrorCode = "BadPermissionError"
	DeploymentJobErrorCodeDownloadConditionFailed             DeploymentJobErrorCode = "DownloadConditionFailed"
	DeploymentJobErrorCodeInternalServerError                 DeploymentJobErrorCode = "InternalServerError"
)

// Values returns all known values for DeploymentJobErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentJobErrorCode) Values() []DeploymentJobErrorCode {
	return []DeploymentJobErrorCode{
		"ResourceNotFound",
		"EnvironmentSetupError",
		"EtagMismatch",
		"FailureThresholdBreached",
		"RobotDeploymentAborted",
		"RobotDeploymentNoResponse",
		"RobotAgentConnectionTimeout",
		"GreengrassDeploymentFailed",
		"InvalidGreengrassGroup",
		"MissingRobotArchitecture",
		"MissingRobotApplicationArchitecture",
		"MissingRobotDeploymentResource",
		"GreengrassGroupVersionDoesNotExist",
		"LambdaDeleted",
		"ExtractingBundleFailure",
		"PreLaunchFileFailure",
		"PostLaunchFileFailure",
		"BadPermissionError",
		"DownloadConditionFailed",
		"InternalServerError",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusPending    DeploymentStatus = "Pending"
	DeploymentStatusPreparing  DeploymentStatus = "Preparing"
	DeploymentStatusInProgress DeploymentStatus = "InProgress"
	DeploymentStatusFailed     DeploymentStatus = "Failed"
	DeploymentStatusSucceeded  DeploymentStatus = "Succeeded"
	DeploymentStatusCanceled   DeploymentStatus = "Canceled"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"Pending",
		"Preparing",
		"InProgress",
		"Failed",
		"Succeeded",
		"Canceled",
	}
}

type FailureBehavior string

// Enum values for FailureBehavior
const (
	FailureBehaviorFail     FailureBehavior = "Fail"
	FailureBehaviorContinue FailureBehavior = "Continue"
)

// Values returns all known values for FailureBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailureBehavior) Values() []FailureBehavior {
	return []FailureBehavior{
		"Fail",
		"Continue",
	}
}

type RenderingEngineType string

// Enum values for RenderingEngineType
const (
	RenderingEngineTypeOgre RenderingEngineType = "OGRE"
)

// Values returns all known values for RenderingEngineType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RenderingEngineType) Values() []RenderingEngineType {
	return []RenderingEngineType{
		"OGRE",
	}
}

type RobotDeploymentStep string

// Enum values for RobotDeploymentStep
const (
	RobotDeploymentStepValidatingStep             RobotDeploymentStep = "Validating"
	RobotDeploymentStepDownloadingExtractingStep  RobotDeploymentStep = "DownloadingExtracting"
	RobotDeploymentStepExecutingDownloadCondition RobotDeploymentStep = "ExecutingDownloadCondition"
	RobotDeploymentStepPreLaunchStep              RobotDeploymentStep = "ExecutingPreLaunch"
	RobotDeploymentStepLaunchingStep              RobotDeploymentStep = "Launching"
	RobotDeploymentStepPostLaunchStep             RobotDeploymentStep = "ExecutingPostLaunch"
	RobotDeploymentStepFinishedStep               RobotDeploymentStep = "Finished"
)

// Values returns all known values for RobotDeploymentStep. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RobotDeploymentStep) Values() []RobotDeploymentStep {
	return []RobotDeploymentStep{
		"Validating",
		"DownloadingExtracting",
		"ExecutingDownloadCondition",
		"ExecutingPreLaunch",
		"Launching",
		"ExecutingPostLaunch",
		"Finished",
	}
}

type RobotSoftwareSuiteType string

// Enum values for RobotSoftwareSuiteType
const (
	RobotSoftwareSuiteTypeRos  RobotSoftwareSuiteType = "ROS"
	RobotSoftwareSuiteTypeRos2 RobotSoftwareSuiteType = "ROS2"
)

// Values returns all known values for RobotSoftwareSuiteType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RobotSoftwareSuiteType) Values() []RobotSoftwareSuiteType {
	return []RobotSoftwareSuiteType{
		"ROS",
		"ROS2",
	}
}

type RobotSoftwareSuiteVersionType string

// Enum values for RobotSoftwareSuiteVersionType
const (
	RobotSoftwareSuiteVersionTypeKinetic RobotSoftwareSuiteVersionType = "Kinetic"
	RobotSoftwareSuiteVersionTypeMelodic RobotSoftwareSuiteVersionType = "Melodic"
	RobotSoftwareSuiteVersionTypeDashing RobotSoftwareSuiteVersionType = "Dashing"
)

// Values returns all known values for RobotSoftwareSuiteVersionType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RobotSoftwareSuiteVersionType) Values() []RobotSoftwareSuiteVersionType {
	return []RobotSoftwareSuiteVersionType{
		"Kinetic",
		"Melodic",
		"Dashing",
	}
}

type RobotStatus string

// Enum values for RobotStatus
const (
	RobotStatusAvailable            RobotStatus = "Available"
	RobotStatusRegistered           RobotStatus = "Registered"
	RobotStatusPendingNewDeployment RobotStatus = "PendingNewDeployment"
	RobotStatusDeploying            RobotStatus = "Deploying"
	RobotStatusFailed               RobotStatus = "Failed"
	RobotStatusInSync               RobotStatus = "InSync"
	RobotStatusNoResponse           RobotStatus = "NoResponse"
)

// Values returns all known values for RobotStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RobotStatus) Values() []RobotStatus {
	return []RobotStatus{
		"Available",
		"Registered",
		"PendingNewDeployment",
		"Deploying",
		"Failed",
		"InSync",
		"NoResponse",
	}
}

type SimulationJobBatchErrorCode string

// Enum values for SimulationJobBatchErrorCode
const (
	SimulationJobBatchErrorCodeInternalServiceError SimulationJobBatchErrorCode = "InternalServiceError"
)

// Values returns all known values for SimulationJobBatchErrorCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SimulationJobBatchErrorCode) Values() []SimulationJobBatchErrorCode {
	return []SimulationJobBatchErrorCode{
		"InternalServiceError",
	}
}

type SimulationJobBatchStatus string

// Enum values for SimulationJobBatchStatus
const (
	SimulationJobBatchStatusPending    SimulationJobBatchStatus = "Pending"
	SimulationJobBatchStatusInProgress SimulationJobBatchStatus = "InProgress"
	SimulationJobBatchStatusFailed     SimulationJobBatchStatus = "Failed"
	SimulationJobBatchStatusCompleted  SimulationJobBatchStatus = "Completed"
	SimulationJobBatchStatusCanceled   SimulationJobBatchStatus = "Canceled"
	SimulationJobBatchStatusCanceling  SimulationJobBatchStatus = "Canceling"
	SimulationJobBatchStatusCompleting SimulationJobBatchStatus = "Completing"
	SimulationJobBatchStatusTimingOut  SimulationJobBatchStatus = "TimingOut"
	SimulationJobBatchStatusTimedOut   SimulationJobBatchStatus = "TimedOut"
)

// Values returns all known values for SimulationJobBatchStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SimulationJobBatchStatus) Values() []SimulationJobBatchStatus {
	return []SimulationJobBatchStatus{
		"Pending",
		"InProgress",
		"Failed",
		"Completed",
		"Canceled",
		"Canceling",
		"Completing",
		"TimingOut",
		"TimedOut",
	}
}

type SimulationJobErrorCode string

// Enum values for SimulationJobErrorCode
const (
	SimulationJobErrorCodeInternalServiceError                       SimulationJobErrorCode = "InternalServiceError"
	SimulationJobErrorCodeRobotApplicationCrash                      SimulationJobErrorCode = "RobotApplicationCrash"
	SimulationJobErrorCodeSimulationApplicationCrash                 SimulationJobErrorCode = "SimulationApplicationCrash"
	SimulationJobErrorCodeBadPermissionsRobotApplication             SimulationJobErrorCode = "BadPermissionsRobotApplication"
	SimulationJobErrorCodeBadPermissionsSimulationApplication        SimulationJobErrorCode = "BadPermissionsSimulationApplication"
	SimulationJobErrorCodeBadPermissionsS3Object                     SimulationJobErrorCode = "BadPermissionsS3Object"
	SimulationJobErrorCodeBadPermissionsS3Output                     SimulationJobErrorCode = "BadPermissionsS3Output"
	SimulationJobErrorCodeBadPermissionsCloudwatchLogs               SimulationJobErrorCode = "BadPermissionsCloudwatchLogs"
	SimulationJobErrorCodeSubnetIpLimitExceeded                      SimulationJobErrorCode = "SubnetIpLimitExceeded"
	SimulationJobErrorCodeENILimitExceeded                           SimulationJobErrorCode = "ENILimitExceeded"
	SimulationJobErrorCodeBadPermissionsUserCredentials              SimulationJobErrorCode = "BadPermissionsUserCredentials"
	SimulationJobErrorCodeInvalidBundleRobotApplication              SimulationJobErrorCode = "InvalidBundleRobotApplication"
	SimulationJobErrorCodeInvalidBundleSimulationApplication         SimulationJobErrorCode = "InvalidBundleSimulationApplication"
	SimulationJobErrorCodeInvalidS3Resource                          SimulationJobErrorCode = "InvalidS3Resource"
	SimulationJobErrorCodeLimitExceeded                              SimulationJobErrorCode = "LimitExceeded"
	SimulationJobErrorCodeMismatchedEtag                             SimulationJobErrorCode = "MismatchedEtag"
	SimulationJobErrorCodeRobotApplicationVersionMismatchedEtag      SimulationJobErrorCode = "RobotApplicationVersionMismatchedEtag"
	SimulationJobErrorCodeSimulationApplicationVersionMismatchedEtag SimulationJobErrorCode = "SimulationApplicationVersionMismatchedEtag"
	SimulationJobErrorCodeResourceNotFound                           SimulationJobErrorCode = "ResourceNotFound"
	SimulationJobErrorCodeRequestThrottled                           SimulationJobErrorCode = "RequestThrottled"
	SimulationJobErrorCodeBatchTimedOut                              SimulationJobErrorCode = "BatchTimedOut"
	SimulationJobErrorCodeBatchCanceled                              SimulationJobErrorCode = "BatchCanceled"
	SimulationJobErrorCodeInvalidInput                               SimulationJobErrorCode = "InvalidInput"
	SimulationJobErrorCodeWrongRegionS3Bucket                        SimulationJobErrorCode = "WrongRegionS3Bucket"
	SimulationJobErrorCodeWrongRegionS3Output                        SimulationJobErrorCode = "WrongRegionS3Output"
	SimulationJobErrorCodeWrongRegionRobotApplication                SimulationJobErrorCode = "WrongRegionRobotApplication"
	SimulationJobErrorCodeWrongRegionSimulationApplication           SimulationJobErrorCode = "WrongRegionSimulationApplication"
	SimulationJobErrorCodeUploadContentMismatchError                 SimulationJobErrorCode = "UploadContentMismatchError"
)

// Values returns all known values for SimulationJobErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SimulationJobErrorCode) Values() []SimulationJobErrorCode {
	return []SimulationJobErrorCode{
		"InternalServiceError",
		"RobotApplicationCrash",
		"SimulationApplicationCrash",
		"BadPermissionsRobotApplication",
		"BadPermissionsSimulationApplication",
		"BadPermissionsS3Object",
		"BadPermissionsS3Output",
		"BadPermissionsCloudwatchLogs",
		"SubnetIpLimitExceeded",
		"ENILimitExceeded",
		"BadPermissionsUserCredentials",
		"InvalidBundleRobotApplication",
		"InvalidBundleSimulationApplication",
		"InvalidS3Resource",
		"LimitExceeded",
		"MismatchedEtag",
		"RobotApplicationVersionMismatchedEtag",
		"SimulationApplicationVersionMismatchedEtag",
		"ResourceNotFound",
		"RequestThrottled",
		"BatchTimedOut",
		"BatchCanceled",
		"InvalidInput",
		"WrongRegionS3Bucket",
		"WrongRegionS3Output",
		"WrongRegionRobotApplication",
		"WrongRegionSimulationApplication",
		"UploadContentMismatchError",
	}
}

type SimulationJobStatus string

// Enum values for SimulationJobStatus
const (
	SimulationJobStatusPending       SimulationJobStatus = "Pending"
	SimulationJobStatusPreparing     SimulationJobStatus = "Preparing"
	SimulationJobStatusRunning       SimulationJobStatus = "Running"
	SimulationJobStatusRestarting    SimulationJobStatus = "Restarting"
	SimulationJobStatusCompleted     SimulationJobStatus = "Completed"
	SimulationJobStatusFailed        SimulationJobStatus = "Failed"
	SimulationJobStatusRunningFailed SimulationJobStatus = "RunningFailed"
	SimulationJobStatusTerminating   SimulationJobStatus = "Terminating"
	SimulationJobStatusTerminated    SimulationJobStatus = "Terminated"
	SimulationJobStatusCanceled      SimulationJobStatus = "Canceled"
)

// Values returns all known values for SimulationJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SimulationJobStatus) Values() []SimulationJobStatus {
	return []SimulationJobStatus{
		"Pending",
		"Preparing",
		"Running",
		"Restarting",
		"Completed",
		"Failed",
		"RunningFailed",
		"Terminating",
		"Terminated",
		"Canceled",
	}
}

type SimulationSoftwareSuiteType string

// Enum values for SimulationSoftwareSuiteType
const (
	SimulationSoftwareSuiteTypeGazebo     SimulationSoftwareSuiteType = "Gazebo"
	SimulationSoftwareSuiteTypeRosbagPlay SimulationSoftwareSuiteType = "RosbagPlay"
)

// Values returns all known values for SimulationSoftwareSuiteType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SimulationSoftwareSuiteType) Values() []SimulationSoftwareSuiteType {
	return []SimulationSoftwareSuiteType{
		"Gazebo",
		"RosbagPlay",
	}
}

type UploadBehavior string

// Enum values for UploadBehavior
const (
	UploadBehaviorUploadOnTerminate       UploadBehavior = "UPLOAD_ON_TERMINATE"
	UploadBehaviorUploadRollingAutoRemove UploadBehavior = "UPLOAD_ROLLING_AUTO_REMOVE"
)

// Values returns all known values for UploadBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UploadBehavior) Values() []UploadBehavior {
	return []UploadBehavior{
		"UPLOAD_ON_TERMINATE",
		"UPLOAD_ROLLING_AUTO_REMOVE",
	}
}

type WorldExportJobErrorCode string

// Enum values for WorldExportJobErrorCode
const (
	WorldExportJobErrorCodeInternalServiceError WorldExportJobErrorCode = "InternalServiceError"
	WorldExportJobErrorCodeLimitExceeded        WorldExportJobErrorCode = "LimitExceeded"
	WorldExportJobErrorCodeResourceNotFound     WorldExportJobErrorCode = "ResourceNotFound"
	WorldExportJobErrorCodeRequestThrottled     WorldExportJobErrorCode = "RequestThrottled"
	WorldExportJobErrorCodeInvalidInput         WorldExportJobErrorCode = "InvalidInput"
	WorldExportJobErrorCodeAccessDenied         WorldExportJobErrorCode = "AccessDenied"
)

// Values returns all known values for WorldExportJobErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorldExportJobErrorCode) Values() []WorldExportJobErrorCode {
	return []WorldExportJobErrorCode{
		"InternalServiceError",
		"LimitExceeded",
		"ResourceNotFound",
		"RequestThrottled",
		"InvalidInput",
		"AccessDenied",
	}
}

type WorldExportJobStatus string

// Enum values for WorldExportJobStatus
const (
	WorldExportJobStatusPending   WorldExportJobStatus = "Pending"
	WorldExportJobStatusRunning   WorldExportJobStatus = "Running"
	WorldExportJobStatusCompleted WorldExportJobStatus = "Completed"
	WorldExportJobStatusFailed    WorldExportJobStatus = "Failed"
	WorldExportJobStatusCanceling WorldExportJobStatus = "Canceling"
	WorldExportJobStatusCanceled  WorldExportJobStatus = "Canceled"
)

// Values returns all known values for WorldExportJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorldExportJobStatus) Values() []WorldExportJobStatus {
	return []WorldExportJobStatus{
		"Pending",
		"Running",
		"Completed",
		"Failed",
		"Canceling",
		"Canceled",
	}
}

type WorldGenerationJobErrorCode string

// Enum values for WorldGenerationJobErrorCode
const (
	WorldGenerationJobErrorCodeInternalServiceError     WorldGenerationJobErrorCode = "InternalServiceError"
	WorldGenerationJobErrorCodeLimitExceeded            WorldGenerationJobErrorCode = "LimitExceeded"
	WorldGenerationJobErrorCodeResourceNotFound         WorldGenerationJobErrorCode = "ResourceNotFound"
	WorldGenerationJobErrorCodeRequestThrottled         WorldGenerationJobErrorCode = "RequestThrottled"
	WorldGenerationJobErrorCodeInvalidInput             WorldGenerationJobErrorCode = "InvalidInput"
	WorldGenerationJobErrorCodeAllWorldGenerationFailed WorldGenerationJobErrorCode = "AllWorldGenerationFailed"
)

// Values returns all known values for WorldGenerationJobErrorCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorldGenerationJobErrorCode) Values() []WorldGenerationJobErrorCode {
	return []WorldGenerationJobErrorCode{
		"InternalServiceError",
		"LimitExceeded",
		"ResourceNotFound",
		"RequestThrottled",
		"InvalidInput",
		"AllWorldGenerationFailed",
	}
}

type WorldGenerationJobStatus string

// Enum values for WorldGenerationJobStatus
const (
	WorldGenerationJobStatusPending       WorldGenerationJobStatus = "Pending"
	WorldGenerationJobStatusRunning       WorldGenerationJobStatus = "Running"
	WorldGenerationJobStatusCompleted     WorldGenerationJobStatus = "Completed"
	WorldGenerationJobStatusFailed        WorldGenerationJobStatus = "Failed"
	WorldGenerationJobStatusPartialFailed WorldGenerationJobStatus = "PartialFailed"
	WorldGenerationJobStatusCanceling     WorldGenerationJobStatus = "Canceling"
	WorldGenerationJobStatusCanceled      WorldGenerationJobStatus = "Canceled"
)

// Values returns all known values for WorldGenerationJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorldGenerationJobStatus) Values() []WorldGenerationJobStatus {
	return []WorldGenerationJobStatus{
		"Pending",
		"Running",
		"Completed",
		"Failed",
		"PartialFailed",
		"Canceling",
		"Canceled",
	}
}
