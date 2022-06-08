// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type LogExport string

// Enum values for LogExport
const (
	LogExportUserActivityLog LogExport = "useractivitylog"
	LogExportUserLog         LogExport = "userlog"
	LogExportConnectionLog   LogExport = "connectionlog"
)

// Values returns all known values for LogExport. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogExport) Values() []LogExport {
	return []LogExport{
		"useractivitylog",
		"userlog",
		"connectionlog",
	}
}

type NamespaceStatus string

// Enum values for NamespaceStatus
const (
	NamespaceStatusAvailable NamespaceStatus = "AVAILABLE"
	NamespaceStatusModifying NamespaceStatus = "MODIFYING"
	NamespaceStatusDeleting  NamespaceStatus = "DELETING"
)

// Values returns all known values for NamespaceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NamespaceStatus) Values() []NamespaceStatus {
	return []NamespaceStatus{
		"AVAILABLE",
		"MODIFYING",
		"DELETING",
	}
}

type SnapshotStatus string

// Enum values for SnapshotStatus
const (
	SnapshotStatusAvailable SnapshotStatus = "AVAILABLE"
	SnapshotStatusCreating  SnapshotStatus = "CREATING"
	SnapshotStatusDeleted   SnapshotStatus = "DELETED"
	SnapshotStatusCancelled SnapshotStatus = "CANCELLED"
	SnapshotStatusFailed    SnapshotStatus = "FAILED"
	SnapshotStatusCopying   SnapshotStatus = "COPYING"
)

// Values returns all known values for SnapshotStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SnapshotStatus) Values() []SnapshotStatus {
	return []SnapshotStatus{
		"AVAILABLE",
		"CREATING",
		"DELETED",
		"CANCELLED",
		"FAILED",
		"COPYING",
	}
}

type UsageLimitBreachAction string

// Enum values for UsageLimitBreachAction
const (
	UsageLimitBreachActionLog        UsageLimitBreachAction = "log"
	UsageLimitBreachActionEmitMetric UsageLimitBreachAction = "emit-metric"
	UsageLimitBreachActionDeactivate UsageLimitBreachAction = "deactivate"
)

// Values returns all known values for UsageLimitBreachAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitBreachAction) Values() []UsageLimitBreachAction {
	return []UsageLimitBreachAction{
		"log",
		"emit-metric",
		"deactivate",
	}
}

type UsageLimitPeriod string

// Enum values for UsageLimitPeriod
const (
	UsageLimitPeriodDaily   UsageLimitPeriod = "daily"
	UsageLimitPeriodWeekly  UsageLimitPeriod = "weekly"
	UsageLimitPeriodMonthly UsageLimitPeriod = "monthly"
)

// Values returns all known values for UsageLimitPeriod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitPeriod) Values() []UsageLimitPeriod {
	return []UsageLimitPeriod{
		"daily",
		"weekly",
		"monthly",
	}
}

type UsageLimitUsageType string

// Enum values for UsageLimitUsageType
const (
	UsageLimitUsageTypeServerlessCompute      UsageLimitUsageType = "serverless-compute"
	UsageLimitUsageTypeCrossRegionDatasharing UsageLimitUsageType = "cross-region-datasharing"
)

// Values returns all known values for UsageLimitUsageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitUsageType) Values() []UsageLimitUsageType {
	return []UsageLimitUsageType{
		"serverless-compute",
		"cross-region-datasharing",
	}
}

type WorkgroupStatus string

// Enum values for WorkgroupStatus
const (
	WorkgroupStatusCreating  WorkgroupStatus = "CREATING"
	WorkgroupStatusAvailable WorkgroupStatus = "AVAILABLE"
	WorkgroupStatusModifying WorkgroupStatus = "MODIFYING"
	WorkgroupStatusDeleting  WorkgroupStatus = "DELETING"
)

// Values returns all known values for WorkgroupStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkgroupStatus) Values() []WorkgroupStatus {
	return []WorkgroupStatus{
		"CREATING",
		"AVAILABLE",
		"MODIFYING",
		"DELETING",
	}
}
