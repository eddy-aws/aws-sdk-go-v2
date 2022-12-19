// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Returned if the access point that you are trying to create already exists, with
// the creation token you provided in the request.
type AccessPointAlreadyExists struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_    *string
	AccessPointId *string

	noSmithyDocumentSerde
}

func (e *AccessPointAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessPointAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessPointAlreadyExists) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AccessPointAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessPointAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the Amazon Web Services account has already created the maximum
// number of access points allowed per file system. For more informaton, see
// https://docs.aws.amazon.com/efs/latest/ug/limits.html#limits-efs-resources-per-account-per-region
// (https://docs.aws.amazon.com/efs/latest/ug/limits.html#limits-efs-resources-per-account-per-region).
type AccessPointLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *AccessPointLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessPointLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessPointLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AccessPointLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessPointLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the specified AccessPointId value doesn't exist in the requester's
// Amazon Web Services account.
type AccessPointNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *AccessPointNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessPointNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessPointNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AccessPointNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessPointNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the Availability Zone that was specified for a mount target is
// different from the Availability Zone that was specified for One Zone storage.
// For more information, see Regional and One Zone storage redundancy
// (https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html).
type AvailabilityZonesMismatch struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *AvailabilityZonesMismatch) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AvailabilityZonesMismatch) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AvailabilityZonesMismatch) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AvailabilityZonesMismatch"
	}
	return *e.ErrorCodeOverride
}
func (e *AvailabilityZonesMismatch) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the request is malformed or contains an error such as an invalid
// parameter value or a missing required parameter.
type BadRequest struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequest) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "BadRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service timed out trying to fulfill the request, and the client should try
// the call again.
type DependencyTimeout struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *DependencyTimeout) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DependencyTimeout) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DependencyTimeout) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DependencyTimeout"
	}
	return *e.ErrorCodeOverride
}
func (e *DependencyTimeout) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Returned if the file system you are trying to create already exists, with the
// creation token you provided.
type FileSystemAlreadyExists struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_   *string
	FileSystemId *string

	noSmithyDocumentSerde
}

func (e *FileSystemAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemAlreadyExists) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "FileSystemAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *FileSystemAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if a file system has mount targets.
type FileSystemInUse struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *FileSystemInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemInUse) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "FileSystemInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *FileSystemInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the Amazon Web Services account has already created the maximum
// number of file systems allowed per account.
type FileSystemLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *FileSystemLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "FileSystemLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *FileSystemLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the specified FileSystemId value doesn't exist in the requester's
// Amazon Web Services account.
type FileSystemNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *FileSystemNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "FileSystemNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *FileSystemNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the file system's lifecycle state is not "available".
type IncorrectFileSystemLifeCycleState struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *IncorrectFileSystemLifeCycleState) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncorrectFileSystemLifeCycleState) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncorrectFileSystemLifeCycleState) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "IncorrectFileSystemLifeCycleState"
	}
	return *e.ErrorCodeOverride
}
func (e *IncorrectFileSystemLifeCycleState) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the mount target is not in the correct state for the operation.
type IncorrectMountTargetState struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *IncorrectMountTargetState) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncorrectMountTargetState) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncorrectMountTargetState) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "IncorrectMountTargetState"
	}
	return *e.ErrorCodeOverride
}
func (e *IncorrectMountTargetState) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if there's not enough capacity to provision additional throughput. This
// value might be returned when you try to create a file system in provisioned
// throughput mode, when you attempt to increase the provisioned throughput of an
// existing file system, or when you attempt to change an existing file system from
// Bursting Throughput to Provisioned Throughput mode. Try again later.
type InsufficientThroughputCapacity struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *InsufficientThroughputCapacity) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientThroughputCapacity) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientThroughputCapacity) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InsufficientThroughputCapacity"
	}
	return *e.ErrorCodeOverride
}
func (e *InsufficientThroughputCapacity) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Returned if an error occurred on the server side.
type InternalServerError struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Returned if the FileSystemPolicy is malformed or contains an error such as a
// parameter value that is not valid or a missing required parameter. Returned in
// the case of a policy lockout safety check error.
type InvalidPolicyException struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *InvalidPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPolicyException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidPolicyException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the request specified an IpAddress that is already in use in the
// subnet.
type IpAddressInUse struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *IpAddressInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IpAddressInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IpAddressInUse) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "IpAddressInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *IpAddressInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the mount target would violate one of the specified restrictions
// based on the file system's existing mount targets.
type MountTargetConflict struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *MountTargetConflict) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MountTargetConflict) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MountTargetConflict) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "MountTargetConflict"
	}
	return *e.ErrorCodeOverride
}
func (e *MountTargetConflict) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if there is no mount target with the specified ID found in the caller's
// Amazon Web Services account.
type MountTargetNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *MountTargetNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MountTargetNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MountTargetNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "MountTargetNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *MountTargetNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The calling account has reached the limit for elastic network interfaces for the
// specific Amazon Web Services Region. Either delete some network interfaces or
// request that the account quota be raised. For more information, see Amazon VPC
// Quotas
// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Appendix_Limits.html)
// in the Amazon VPC User Guide (see the Network interfaces per Region entry in the
// Network interfaces table).
type NetworkInterfaceLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *NetworkInterfaceLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NetworkInterfaceLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NetworkInterfaceLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NetworkInterfaceLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *NetworkInterfaceLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if IpAddress was not specified in the request and there are no free IP
// addresses in the subnet.
type NoFreeAddressesInSubnet struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *NoFreeAddressesInSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoFreeAddressesInSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoFreeAddressesInSubnet) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NoFreeAddressesInSubnet"
	}
	return *e.ErrorCodeOverride
}
func (e *NoFreeAddressesInSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the default file system policy is in effect for the EFS file system
// specified.
type PolicyNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *PolicyNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "PolicyNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *PolicyNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the specified file system does not have a replication configuration.
type ReplicationNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *ReplicationNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ReplicationNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ReplicationNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the size of SecurityGroups specified in the request is greater than
// five.
type SecurityGroupLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *SecurityGroupLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SecurityGroupLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SecurityGroupLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "SecurityGroupLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *SecurityGroupLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if one of the specified security groups doesn't exist in the subnet's
// virtual private cloud (VPC).
type SecurityGroupNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *SecurityGroupNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SecurityGroupNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SecurityGroupNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "SecurityGroupNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *SecurityGroupNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if there is no subnet with ID SubnetId provided in the request.
type SubnetNotFound struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *SubnetNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "SubnetNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned when the CreateAccessPoint API action is called too quickly and the
// number of Access Points on the file system is nearing the limit of 120
// (https://docs.aws.amazon.com/efs/latest/ug/limits.html#limits-efs-resources-per-account-per-region).
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the throughput mode or amount of provisioned throughput can't be
// changed because the throughput limit of 1024 MiB/s has been reached.
type ThroughputLimitExceeded struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *ThroughputLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThroughputLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThroughputLimitExceeded) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ThroughputLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *ThroughputLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if you don’t wait at least 24 hours before either changing the
// throughput mode, or decreasing the Provisioned Throughput value.
type TooManyRequests struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *TooManyRequests) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequests) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequests) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "TooManyRequests"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRequests) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the requested Amazon EFS functionality is not available in the
// specified Availability Zone.
type UnsupportedAvailabilityZone struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *UnsupportedAvailabilityZone) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedAvailabilityZone) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedAvailabilityZone) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UnsupportedAvailabilityZone"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedAvailabilityZone) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the Backup service is not available in the Amazon Web Services
// Region in which the request was made.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
