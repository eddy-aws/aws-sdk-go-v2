// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EndpointTypesElement string

// Enum values for EndpointTypesElement
const (
	EndpointTypesElementPush            EndpointTypesElement = "PUSH"
	EndpointTypesElementGcm             EndpointTypesElement = "GCM"
	EndpointTypesElementApns            EndpointTypesElement = "APNS"
	EndpointTypesElementApnsSandbox     EndpointTypesElement = "APNS_SANDBOX"
	EndpointTypesElementApnsVoip        EndpointTypesElement = "APNS_VOIP"
	EndpointTypesElementApnsVoipSandbox EndpointTypesElement = "APNS_VOIP_SANDBOX"
	EndpointTypesElementAdm             EndpointTypesElement = "ADM"
	EndpointTypesElementSms             EndpointTypesElement = "SMS"
	EndpointTypesElementVoice           EndpointTypesElement = "VOICE"
	EndpointTypesElementEmail           EndpointTypesElement = "EMAIL"
	EndpointTypesElementBaidu           EndpointTypesElement = "BAIDU"
	EndpointTypesElementCustom          EndpointTypesElement = "CUSTOM"
)

// Values returns all known values for EndpointTypesElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointTypesElement) Values() []EndpointTypesElement {
	return []EndpointTypesElement{
		"PUSH",
		"GCM",
		"APNS",
		"APNS_SANDBOX",
		"APNS_VOIP",
		"APNS_VOIP_SANDBOX",
		"ADM",
		"SMS",
		"VOICE",
		"EMAIL",
		"BAIDU",
		"CUSTOM",
	}
}

type Action string

// Enum values for Action
const (
	ActionOpenApp  Action = "OPEN_APP"
	ActionDeepLink Action = "DEEP_LINK"
	ActionUrl      Action = "URL"
)

// Values returns all known values for Action. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Action) Values() []Action {
	return []Action{
		"OPEN_APP",
		"DEEP_LINK",
		"URL",
	}
}

type AttributeType string

// Enum values for AttributeType
const (
	AttributeTypeInclusive AttributeType = "INCLUSIVE"
	AttributeTypeExclusive AttributeType = "EXCLUSIVE"
	AttributeTypeContains  AttributeType = "CONTAINS"
	AttributeTypeBefore    AttributeType = "BEFORE"
	AttributeTypeAfter     AttributeType = "AFTER"
	AttributeTypeOn        AttributeType = "ON"
	AttributeTypeBetween   AttributeType = "BETWEEN"
)

// Values returns all known values for AttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttributeType) Values() []AttributeType {
	return []AttributeType{
		"INCLUSIVE",
		"EXCLUSIVE",
		"CONTAINS",
		"BEFORE",
		"AFTER",
		"ON",
		"BETWEEN",
	}
}

type CampaignStatus string

// Enum values for CampaignStatus
const (
	CampaignStatusScheduled      CampaignStatus = "SCHEDULED"
	CampaignStatusExecuting      CampaignStatus = "EXECUTING"
	CampaignStatusPendingNextRun CampaignStatus = "PENDING_NEXT_RUN"
	CampaignStatusCompleted      CampaignStatus = "COMPLETED"
	CampaignStatusPaused         CampaignStatus = "PAUSED"
	CampaignStatusDeleted        CampaignStatus = "DELETED"
	CampaignStatusInvalid        CampaignStatus = "INVALID"
)

// Values returns all known values for CampaignStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CampaignStatus) Values() []CampaignStatus {
	return []CampaignStatus{
		"SCHEDULED",
		"EXECUTING",
		"PENDING_NEXT_RUN",
		"COMPLETED",
		"PAUSED",
		"DELETED",
		"INVALID",
	}
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypePush            ChannelType = "PUSH"
	ChannelTypeGcm             ChannelType = "GCM"
	ChannelTypeApns            ChannelType = "APNS"
	ChannelTypeApnsSandbox     ChannelType = "APNS_SANDBOX"
	ChannelTypeApnsVoip        ChannelType = "APNS_VOIP"
	ChannelTypeApnsVoipSandbox ChannelType = "APNS_VOIP_SANDBOX"
	ChannelTypeAdm             ChannelType = "ADM"
	ChannelTypeSms             ChannelType = "SMS"
	ChannelTypeVoice           ChannelType = "VOICE"
	ChannelTypeEmail           ChannelType = "EMAIL"
	ChannelTypeBaidu           ChannelType = "BAIDU"
	ChannelTypeCustom          ChannelType = "CUSTOM"
)

// Values returns all known values for ChannelType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChannelType) Values() []ChannelType {
	return []ChannelType{
		"PUSH",
		"GCM",
		"APNS",
		"APNS_SANDBOX",
		"APNS_VOIP",
		"APNS_VOIP_SANDBOX",
		"ADM",
		"SMS",
		"VOICE",
		"EMAIL",
		"BAIDU",
		"CUSTOM",
	}
}

type DeliveryStatus string

// Enum values for DeliveryStatus
const (
	DeliveryStatusSuccessful       DeliveryStatus = "SUCCESSFUL"
	DeliveryStatusThrottled        DeliveryStatus = "THROTTLED"
	DeliveryStatusTemporaryFailure DeliveryStatus = "TEMPORARY_FAILURE"
	DeliveryStatusPermanentFailure DeliveryStatus = "PERMANENT_FAILURE"
	DeliveryStatusUnknownFailure   DeliveryStatus = "UNKNOWN_FAILURE"
	DeliveryStatusOptOut           DeliveryStatus = "OPT_OUT"
	DeliveryStatusDuplicate        DeliveryStatus = "DUPLICATE"
)

// Values returns all known values for DeliveryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryStatus) Values() []DeliveryStatus {
	return []DeliveryStatus{
		"SUCCESSFUL",
		"THROTTLED",
		"TEMPORARY_FAILURE",
		"PERMANENT_FAILURE",
		"UNKNOWN_FAILURE",
		"OPT_OUT",
		"DUPLICATE",
	}
}

type DimensionType string

// Enum values for DimensionType
const (
	DimensionTypeInclusive DimensionType = "INCLUSIVE"
	DimensionTypeExclusive DimensionType = "EXCLUSIVE"
)

// Values returns all known values for DimensionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DimensionType) Values() []DimensionType {
	return []DimensionType{
		"INCLUSIVE",
		"EXCLUSIVE",
	}
}

type Duration string

// Enum values for Duration
const (
	DurationHr24  Duration = "HR_24"
	DurationDay7  Duration = "DAY_7"
	DurationDay14 Duration = "DAY_14"
	DurationDay30 Duration = "DAY_30"
)

// Values returns all known values for Duration. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Duration) Values() []Duration {
	return []Duration{
		"HR_24",
		"DAY_7",
		"DAY_14",
		"DAY_30",
	}
}

type FilterType string

// Enum values for FilterType
const (
	FilterTypeSystem   FilterType = "SYSTEM"
	FilterTypeEndpoint FilterType = "ENDPOINT"
)

// Values returns all known values for FilterType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FilterType) Values() []FilterType {
	return []FilterType{
		"SYSTEM",
		"ENDPOINT",
	}
}

type Format string

// Enum values for Format
const (
	FormatCsv  Format = "CSV"
	FormatJson Format = "JSON"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"CSV",
		"JSON",
	}
}

type Frequency string

// Enum values for Frequency
const (
	FrequencyOnce    Frequency = "ONCE"
	FrequencyHourly  Frequency = "HOURLY"
	FrequencyDaily   Frequency = "DAILY"
	FrequencyWeekly  Frequency = "WEEKLY"
	FrequencyMonthly Frequency = "MONTHLY"
	FrequencyEvent   Frequency = "EVENT"
)

// Values returns all known values for Frequency. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Frequency) Values() []Frequency {
	return []Frequency{
		"ONCE",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"EVENT",
	}
}

type Include string

// Enum values for Include
const (
	IncludeAll  Include = "ALL"
	IncludeAny  Include = "ANY"
	IncludeNone Include = "NONE"
)

// Values returns all known values for Include. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Include) Values() []Include {
	return []Include{
		"ALL",
		"ANY",
		"NONE",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated                    JobStatus = "CREATED"
	JobStatusPreparingForInitialization JobStatus = "PREPARING_FOR_INITIALIZATION"
	JobStatusInitializing               JobStatus = "INITIALIZING"
	JobStatusProcessing                 JobStatus = "PROCESSING"
	JobStatusPendingJob                 JobStatus = "PENDING_JOB"
	JobStatusCompleting                 JobStatus = "COMPLETING"
	JobStatusCompleted                  JobStatus = "COMPLETED"
	JobStatusFailing                    JobStatus = "FAILING"
	JobStatusFailed                     JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"CREATED",
		"PREPARING_FOR_INITIALIZATION",
		"INITIALIZING",
		"PROCESSING",
		"PENDING_JOB",
		"COMPLETING",
		"COMPLETED",
		"FAILING",
		"FAILED",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeTransactional MessageType = "TRANSACTIONAL"
	MessageTypePromotional   MessageType = "PROMOTIONAL"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"TRANSACTIONAL",
		"PROMOTIONAL",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeDelivery Mode = "DELIVERY"
	ModeFilter   Mode = "FILTER"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"DELIVERY",
		"FILTER",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorAll Operator = "ALL"
	OperatorAny Operator = "ANY"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"ALL",
		"ANY",
	}
}

type RecencyType string

// Enum values for RecencyType
const (
	RecencyTypeActive   RecencyType = "ACTIVE"
	RecencyTypeInactive RecencyType = "INACTIVE"
)

// Values returns all known values for RecencyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RecencyType) Values() []RecencyType {
	return []RecencyType{
		"ACTIVE",
		"INACTIVE",
	}
}

type SegmentType string

// Enum values for SegmentType
const (
	SegmentTypeDimensional SegmentType = "DIMENSIONAL"
	SegmentTypeImport      SegmentType = "IMPORT"
)

// Values returns all known values for SegmentType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SegmentType) Values() []SegmentType {
	return []SegmentType{
		"DIMENSIONAL",
		"IMPORT",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeAll  SourceType = "ALL"
	SourceTypeAny  SourceType = "ANY"
	SourceTypeNone SourceType = "NONE"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"ALL",
		"ANY",
		"NONE",
	}
}

type State string

// Enum values for State
const (
	StateDraft     State = "DRAFT"
	StateActive    State = "ACTIVE"
	StateCompleted State = "COMPLETED"
	StateCancelled State = "CANCELLED"
	StateClosed    State = "CLOSED"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"DRAFT",
		"ACTIVE",
		"COMPLETED",
		"CANCELLED",
		"CLOSED",
	}
}

type TemplateType string

// Enum values for TemplateType
const (
	TemplateTypeEmail TemplateType = "EMAIL"
	TemplateTypeSms   TemplateType = "SMS"
	TemplateTypeVoice TemplateType = "VOICE"
	TemplateTypePush  TemplateType = "PUSH"
)

// Values returns all known values for TemplateType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TemplateType) Values() []TemplateType {
	return []TemplateType{
		"EMAIL",
		"SMS",
		"VOICE",
		"PUSH",
	}
}

type Type string

// Enum values for Type
const (
	TypeAll  Type = "ALL"
	TypeAny  Type = "ANY"
	TypeNone Type = "NONE"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"ALL",
		"ANY",
		"NONE",
	}
}
