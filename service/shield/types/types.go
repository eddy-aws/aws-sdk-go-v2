// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The details of a DDoS attack.
type AttackDetail struct {

	// List of counters that describe the attack for the specified time period.
	AttackCounters []SummarizedCounter

	// The unique identifier (ID) of the attack.
	AttackId *string

	// The array of AttackProperty objects.
	AttackProperties []AttackProperty

	// The time the attack ended, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time

	// List of mitigation actions taken for the attack.
	Mitigations []Mitigation

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string

	// The time the attack started, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time

	// If applicable, additional detail about the resource being attacked, for example,
	// IP address or URL.
	SubResources []SubResourceSummary
}

// Details of the described attack.
type AttackProperty struct {

	// The type of distributed denial of service (DDoS) event that was observed.
	// NETWORK indicates layer 3 and layer 4 events and APPLICATION indicates layer 7
	// events.
	AttackLayer AttackLayer

	// Defines the DDoS attack property information that is provided. The
	// WORDPRESS_PINGBACK_REFLECTOR and WORDPRESS_PINGBACK_SOURCE values are valid only
	// for WordPress reflective pingback DDoS attacks.
	AttackPropertyIdentifier AttackPropertyIdentifier

	// The array of contributor objects that includes the top five contributors to an
	// attack.
	TopContributors []Contributor

	// The total contributions made to this attack by all contributors, not just the
	// five listed in the TopContributors list.
	Total int64

	// The unit of the Value of the contributions.
	Unit Unit
}

// A single attack statistics data record. This is returned by
// DescribeAttackStatistics along with a time range indicating the time period that
// the attack statistics apply to.
type AttackStatisticsDataItem struct {

	// The number of attacks detected during the time period. This is always present,
	// but might be zero.
	//
	// This member is required.
	AttackCount int64

	// Information about the volume of attacks during the time period. If the
	// accompanying AttackCount is zero, this setting might be empty.
	AttackVolume *AttackVolume
}

// Summarizes all DDoS attacks for a specified time period.
type AttackSummary struct {

	// The unique identifier (ID) of the attack.
	AttackId *string

	// The list of attacks for a specified time period.
	AttackVectors []AttackVectorDescription

	// The end time of the attack, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string

	// The start time of the attack, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time
}

// Describes the attack.
type AttackVectorDescription struct {

	// The attack type. Valid values:
	//
	// * UDP_TRAFFIC
	//
	// * UDP_FRAGMENT
	//
	// *
	// GENERIC_UDP_REFLECTION
	//
	// * DNS_REFLECTION
	//
	// * NTP_REFLECTION
	//
	// *
	// CHARGEN_REFLECTION
	//
	// * SSDP_REFLECTION
	//
	// * PORT_MAPPER
	//
	// * RIP_REFLECTION
	//
	// *
	// SNMP_REFLECTION
	//
	// * MSSQL_REFLECTION
	//
	// * NET_BIOS_REFLECTION
	//
	// * SYN_FLOOD
	//
	// *
	// ACK_FLOOD
	//
	// * REQUEST_FLOOD
	//
	// * HTTP_REFLECTION
	//
	// * UDS_REFLECTION
	//
	// *
	// MEMCACHED_REFLECTION
	//
	// This member is required.
	VectorType *string
}

// Information about the volume of attacks during the time period, included in an
// AttackStatisticsDataItem. If the accompanying AttackCount in the statistics
// object is zero, this setting might be empty.
type AttackVolume struct {

	// A statistics object that uses bits per second as the unit. This is included for
	// network level attacks.
	BitsPerSecond *AttackVolumeStatistics

	// A statistics object that uses packets per second as the unit. This is included
	// for network level attacks.
	PacketsPerSecond *AttackVolumeStatistics

	// A statistics object that uses requests per second as the unit. This is included
	// for application level attacks, and is only available for accounts that are
	// subscribed to Shield Advanced.
	RequestsPerSecond *AttackVolumeStatistics
}

// Statistics objects for the various data types in AttackVolume.
type AttackVolumeStatistics struct {

	// The maximum attack volume observed for the given unit.
	//
	// This member is required.
	Max float64
}

// A contributor to the attack and their contribution.
type Contributor struct {

	// The name of the contributor. This is dependent on the AttackPropertyIdentifier.
	// For example, if the AttackPropertyIdentifier is SOURCE_COUNTRY, the Name could
	// be United States.
	Name *string

	// The contribution of this contributor expressed in Protection units. For example
	// 10,000.
	Value int64
}

// Contact information that the DRT can use to contact you if you have proactive
// engagement enabled, for escalations to the DRT and to initiate proactive
// customer support.
type EmergencyContact struct {

	// The email address for the contact.
	//
	// This member is required.
	EmailAddress *string

	// Additional notes regarding the contact.
	ContactNotes *string

	// The phone number for the contact.
	PhoneNumber *string
}

// Specifies how many protections of a given type you can create.
type Limit struct {

	// The maximum number of protections that can be created for the specified Type.
	Max int64

	// The type of protection.
	Type *string
}

// The mitigation applied to a DDoS attack.
type Mitigation struct {

	// The name of the mitigation taken for this attack.
	MitigationName *string
}

// An object that represents a resource that is under DDoS protection.
type Protection struct {

	// The unique identifier (ID) for the Route 53 health check that's associated with
	// the protection.
	HealthCheckIds []string

	// The unique identifier (ID) of the protection.
	Id *string

	// The name of the protection. For example, My CloudFront distributions.
	Name *string

	// The ARN (Amazon Resource Name) of the protection.
	ProtectionArn *string

	// The ARN (Amazon Resource Name) of the AWS resource that is protected.
	ResourceArn *string
}

// A grouping of protected resources that you and AWS Shield Advanced can monitor
// as a collective. This resource grouping improves the accuracy of detection and
// reduces false positives.
type ProtectionGroup struct {

	// Defines how AWS Shield combines resource data for the group in order to detect,
	// mitigate, and report events.
	//
	// * Sum - Use the total traffic across the group.
	// This is a good choice for most cases. Examples include Elastic IP addresses for
	// EC2 instances that scale manually or automatically.
	//
	// * Mean - Use the average of
	// the traffic across the group. This is a good choice for resources that share
	// traffic uniformly. Examples include accelerators and load balancers.
	//
	// * Max -
	// Use the highest traffic from each resource. This is useful for resources that
	// don't share traffic and for resources that share that traffic in a non-uniform
	// way. Examples include CloudFront distributions and origin resources for
	// CloudFront distributions.
	//
	// This member is required.
	Aggregation ProtectionGroupAggregation

	// The Amazon Resource Names (ARNs) of the resources to include in the protection
	// group. You must set this when you set Pattern to ARBITRARY and you must not set
	// it for any other Pattern setting.
	//
	// This member is required.
	Members []string

	// The criteria to use to choose the protected resources for inclusion in the
	// group. You can include all resources that have protections, provide a list of
	// resource Amazon Resource Names (ARNs), or include all resources of a specified
	// resource type.
	//
	// This member is required.
	Pattern ProtectionGroupPattern

	// The name of the protection group. You use this to identify the protection group
	// in lists and to manage the protection group, for example to update, delete, or
	// describe it.
	//
	// This member is required.
	ProtectionGroupId *string

	// The ARN (Amazon Resource Name) of the protection group.
	ProtectionGroupArn *string

	// The resource type to include in the protection group. All protected resources of
	// this type are included in the protection group. You must set this when you set
	// Pattern to BY_RESOURCE_TYPE and you must not set it for any other Pattern
	// setting.
	ResourceType ProtectedResourceType
}

// Limits settings on protection groups with arbitrary pattern type.
type ProtectionGroupArbitraryPatternLimits struct {

	// The maximum number of resources you can specify for a single arbitrary pattern
	// in a protection group.
	//
	// This member is required.
	MaxMembers int64
}

// Limits settings on protection groups for your subscription.
type ProtectionGroupLimits struct {

	// The maximum number of protection groups that you can have at one time.
	//
	// This member is required.
	MaxProtectionGroups int64

	// Limits settings by pattern type in the protection groups for your subscription.
	//
	// This member is required.
	PatternTypeLimits *ProtectionGroupPatternTypeLimits
}

// Limits settings by pattern type in the protection groups for your subscription.
type ProtectionGroupPatternTypeLimits struct {

	// Limits settings on protection groups with arbitrary pattern type.
	//
	// This member is required.
	ArbitraryPatternLimits *ProtectionGroupArbitraryPatternLimits
}

// Limits settings on protections for your subscription.
type ProtectionLimits struct {

	// The maximum number of resource types that you can specify in a protection.
	//
	// This member is required.
	ProtectedResourceTypeLimits []Limit
}

// The attack information for the specified SubResource.
type SubResourceSummary struct {

	// The list of attack types and associated counters.
	AttackVectors []SummarizedAttackVector

	// The counters that describe the details of the attack.
	Counters []SummarizedCounter

	// The unique identifier (ID) of the SubResource.
	Id *string

	// The SubResource type.
	Type SubResourceType
}

// Information about the AWS Shield Advanced subscription for an account.
type Subscription struct {

	// Limits settings for your subscription.
	//
	// This member is required.
	SubscriptionLimits *SubscriptionLimits

	// If ENABLED, the subscription will be automatically renewed at the end of the
	// existing subscription period. When you initally create a subscription, AutoRenew
	// is set to ENABLED. You can change this by submitting an UpdateSubscription
	// request. If the UpdateSubscription request does not included a value for
	// AutoRenew, the existing value for AutoRenew remains unchanged.
	AutoRenew AutoRenew

	// The date and time your subscription will end.
	EndTime *time.Time

	// Specifies how many protections of a given type you can create.
	Limits []Limit

	// If ENABLED, the DDoS Response Team (DRT) will use email and phone to notify
	// contacts about escalations to the DRT and to initiate proactive customer
	// support. If PENDING, you have requested proactive engagement and the request is
	// pending. The status changes to ENABLED when your request is fully processed. If
	// DISABLED, the DRT will not proactively notify contacts about escalations or to
	// initiate proactive customer support.
	ProactiveEngagementStatus ProactiveEngagementStatus

	// The start time of the subscription, in Unix time in seconds. For more
	// information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time

	// The ARN (Amazon Resource Name) of the subscription.
	SubscriptionArn *string

	// The length, in seconds, of the AWS Shield Advanced subscription for the account.
	TimeCommitmentInSeconds int64
}

// Limits settings for your subscription.
type SubscriptionLimits struct {

	// Limits settings on protection groups for your subscription.
	//
	// This member is required.
	ProtectionGroupLimits *ProtectionGroupLimits

	// Limits settings on protections for your subscription.
	//
	// This member is required.
	ProtectionLimits *ProtectionLimits
}

// A summary of information about the attack.
type SummarizedAttackVector struct {

	// The attack type, for example, SNMP reflection or SYN flood.
	//
	// This member is required.
	VectorType *string

	// The list of counters that describe the details of the attack.
	VectorCounters []SummarizedCounter
}

// The counter that describes a DDoS attack.
type SummarizedCounter struct {

	// The average value of the counter for a specified time period.
	Average float64

	// The maximum value of the counter for a specified time period.
	Max float64

	// The number of counters for a specified time period.
	N int32

	// The counter name.
	Name *string

	// The total of counter values for a specified time period.
	Sum float64

	// The unit of the counters.
	Unit *string
}

// A tag associated with an AWS resource. Tags are key:value pairs that you can use
// to categorize and manage your resources, for purposes like billing or other
// management. Typically, the tag key represents a category, such as "environment",
// and the tag value represents a specific value within that category, such as
// "test," "development," or "production". Or you might set the tag key to
// "customer" and the value to the customer name or ID. You can specify one or more
// tags to add to each AWS resource, up to 50 tags for a resource.
type Tag struct {

	// Part of the key:value pair that defines a tag. You can use a tag key to describe
	// a category of information, such as "customer." Tag keys are case-sensitive.
	Key *string

	// Part of the key:value pair that defines a tag. You can use a tag value to
	// describe a specific value within a category, such as "companyA" or "companyB."
	// Tag values are case-sensitive.
	Value *string
}

// The time range.
type TimeRange struct {

	// The start time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	FromInclusive *time.Time

	// The end time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	ToExclusive *time.Time
}

// Provides information about a particular parameter passed inside a request that
// resulted in an exception.
type ValidationExceptionField struct {

	// The message describing why the parameter failed validation.
	//
	// This member is required.
	Message *string

	// The name of the parameter that failed validation.
	//
	// This member is required.
	Name *string
}
