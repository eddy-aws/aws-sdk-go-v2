// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes the properties of a security policy that was specified. For more
// information about security policies, see Working with security policies
// (https://docs.aws.amazon.com/transfer/latest/userguide/security-policies.html).
type DescribedSecurityPolicy struct {

	// Specifies the name of the security policy that is attached to the server.
	//
	// This member is required.
	SecurityPolicyName *string

	// Specifies whether this policy enables Federal Information Processing Standards
	// (FIPS).
	Fips *bool

	// Specifies the enabled Secure Shell (SSH) cipher encryption algorithms in the
	// security policy that is attached to the server.
	SshCiphers []string

	// Specifies the enabled SSH key exchange (KEX) encryption algorithms in the
	// security policy that is attached to the server.
	SshKexs []string

	// Specifies the enabled SSH message authentication code (MAC) encryption
	// algorithms in the security policy that is attached to the server.
	SshMacs []string

	// Specifies the enabled Transport Layer Security (TLS) cipher encryption
	// algorithms in the security policy that is attached to the server.
	TlsCiphers []string
}

// Describes the properties of a file transfer protocol-enabled server that was
// specified.
type DescribedServer struct {

	// Specifies the unique Amazon Resource Name (ARN) of the server.
	//
	// This member is required.
	Arn *string

	// Specifies the ARN of the AWS Certificate Manager (ACM) certificate. Required
	// when Protocols is set to FTPS.
	Certificate *string

	Domain Domain

	// Specifies the virtual private cloud (VPC) endpoint settings that you configured
	// for your server.
	EndpointDetails *EndpointDetails

	// Defines the type of endpoint that your server is connected to. If your server is
	// connected to a VPC endpoint, your server isn't accessible over the public
	// internet.
	EndpointType EndpointType

	// Specifies the Base64-encoded SHA256 fingerprint of the server's host key. This
	// value is equivalent to the output of the ssh-keygen -l -f my-new-server-key
	// command.
	HostKeyFingerprint *string

	// Specifies information to call a customer-supplied authentication API. This field
	// is not populated when the IdentityProviderType of a server is SERVICE_MANAGED.
	IdentityProviderDetails *IdentityProviderDetails

	// Specifies the mode of authentication method enabled for this service. A value of
	// SERVICE_MANAGED means that you are using this server to store and access user
	// credentials within the service. A value of API_GATEWAY indicates that you have
	// integrated an API Gateway endpoint that will be invoked for authenticating your
	// user into the service.
	IdentityProviderType IdentityProviderType

	// Specifies the AWS Identity and Access Management (IAM) role that allows a server
	// to turn on Amazon CloudWatch logging for Amazon S3 events. When set, user
	// activity can be viewed in your CloudWatch logs.
	LoggingRole *string

	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:
	//
	// * SFTP (Secure Shell (SSH) File Transfer Protocol): File transfer over
	// SSH
	//
	// * FTPS (File Transfer Protocol Secure): File transfer with TLS
	// encryption
	//
	// * FTP (File Transfer Protocol): Unencrypted file transfer
	Protocols []Protocol

	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName *string

	// Specifies the unique system-assigned identifier for a server that you
	// instantiate.
	ServerId *string

	// Specifies the condition of a server for the server that was described. A value
	// of ONLINE indicates that the server can accept jobs and transfer files. A State
	// value of OFFLINE means that the server cannot perform file transfer operations.
	// The states of STARTING and STOPPING indicate that the server is in an
	// intermediate state, either not fully able to respond, or not fully offline. The
	// values of START_FAILED or STOP_FAILED can indicate an error condition.
	State State

	// Specifies the key-value pairs that you can use to search for and group servers
	// that were assigned to the server that was described.
	Tags []Tag

	// Specifies the number of users that are assigned to a server you specified with
	// the ServerId.
	UserCount *int32
}

// Describes the properties of a user that was specified.
type DescribedUser struct {

	// Specifies the unique Amazon Resource Name (ARN) for the user that was requested
	// to be described.
	//
	// This member is required.
	Arn *string

	// Specifies the landing directory (or folder), which is the location that files
	// are written to or read from in an Amazon S3 bucket, for the described user. An
	// example is your-Amazon-S3-bucket-name>/home/username .
	HomeDirectory *string

	// Specifies the logical directory mappings that specify what Amazon S3 paths and
	// keys should be visible to your user and how you want to make them visible. You
	// will need to specify the "Entry" and "Target" pair, where Entry shows how the
	// path is made visible and Target is the actual Amazon S3 path. If you only
	// specify a target, it will be displayed as is. You will need to also make sure
	// that your AWS Identity and Access Management (IAM) role provides access to paths
	// in Target. In most cases, you can use this value instead of the scope-down
	// policy to lock your user down to the designated home directory ("chroot"). To do
	// this, you can set Entry to '/' and set Target to the HomeDirectory parameter
	// value.
	HomeDirectoryMappings []HomeDirectoryMapEntry

	// Specifies the type of landing directory (folder) you mapped for your users to
	// see when they log into the file transfer protocol-enabled server. If you set it
	// to PATH, the user will see the absolute Amazon S3 bucket paths as is in their
	// file transfer protocol clients. If you set it LOGICAL, you will need to provide
	// mappings in the HomeDirectoryMappings for how you want to make Amazon S3 paths
	// visible to your users.
	HomeDirectoryType HomeDirectoryType

	// Specifies the name of the policy in use for the described user.
	Policy *string

	PosixProfile *PosixProfile

	// Specifies the IAM role that controls your users' access to your Amazon S3
	// bucket. The policies attached to this role will determine the level of access
	// you want to provide your users when transferring files into and out of your
	// Amazon S3 bucket or buckets. The IAM role should also contain a trust
	// relationship that allows a server to access your resources when servicing your
	// users' transfer requests.
	Role *string

	// Specifies the public key portion of the Secure Shell (SSH) keys stored for the
	// described user.
	SshPublicKeys []SshPublicKey

	// Specifies the key-value pairs for the user requested. Tag can be used to search
	// for and group users for a variety of purposes.
	Tags []Tag

	// Specifies the name of the user that was requested to be described. User names
	// are used for authentication purposes. This is the string that will be used by
	// your user when they log in to your server.
	UserName *string
}

// The virtual private cloud (VPC) endpoint settings that are configured for your
// file transfer protocol-enabled server. With a VPC endpoint, you can restrict
// access to your server and resources only within your VPC. To control incoming
// internet traffic, invoke the UpdateServer API and attach an Elastic IP to your
// server's endpoint.
type EndpointDetails struct {

	// A list of address allocation IDs that are required to attach an Elastic IP
	// address to your server's endpoint. This property can only be set when
	// EndpointType is set to VPC and it is only valid in the UpdateServer API.
	AddressAllocationIds []string

	// A list of security groups IDs that are available to attach to your server's
	// endpoint. This property can only be set when EndpointType is set to VPC. You can
	// only edit the SecurityGroupIds property in the UpdateServer API and only if you
	// are changing the EndpointType from PUBLIC or VPC_ENDPOINT to VPC.
	SecurityGroupIds []string

	// A list of subnet IDs that are required to host your server endpoint in your VPC.
	// This property can only be set when EndpointType is set to VPC.
	SubnetIds []string

	// The ID of the VPC endpoint. This property can only be set when EndpointType is
	// set to VPC_ENDPOINT.
	VpcEndpointId *string

	// The VPC ID of the VPC in which a server's endpoint will be hosted. This property
	// can only be set when EndpointType is set to VPC.
	VpcId *string
}

// Represents an object that contains entries and targets for
// HomeDirectoryMappings.
type HomeDirectoryMapEntry struct {

	// Represents an entry and a target for HomeDirectoryMappings.
	//
	// This member is required.
	Entry *string

	// Represents the map target that is used in a HomeDirectorymapEntry.
	//
	// This member is required.
	Target *string
}

// Returns information related to the type of user authentication that is in use
// for a file transfer protocol-enabled server's users. A server can have only one
// method of authentication.
type IdentityProviderDetails struct {

	// Provides the type of InvocationRole used to authenticate the user account.
	InvocationRole *string

	// Provides the location of the service endpoint used to authenticate users.
	Url *string
}

// Returns properties of a file transfer protocol-enabled server that was
// specified.
type ListedServer struct {

	// Specifies the unique Amazon Resource Name (ARN) for a server to be listed.
	//
	// This member is required.
	Arn *string

	Domain Domain

	// Specifies the type of VPC endpoint that your server is connected to. If your
	// server is connected to a VPC endpoint, your server isn't accessible over the
	// public internet.
	EndpointType EndpointType

	// Specifies the authentication method used to validate a user for a server that
	// was specified. This can include Secure Shell (SSH), user name and password
	// combinations, or your own custom authentication method. Valid values include
	// SERVICE_MANAGED or API_GATEWAY.
	IdentityProviderType IdentityProviderType

	// Specifies the AWS Identity and Access Management (IAM) role that allows a server
	// to turn on Amazon CloudWatch logging.
	LoggingRole *string

	// Specifies the unique system assigned identifier for the servers that were
	// listed.
	ServerId *string

	// Specifies the condition of a server for the server that was described. A value
	// of ONLINE indicates that the server can accept jobs and transfer files. A State
	// value of OFFLINE means that the server cannot perform file transfer operations.
	// The states of STARTING and STOPPING indicate that the server is in an
	// intermediate state, either not fully able to respond, or not fully offline. The
	// values of START_FAILED or STOP_FAILED can indicate an error condition.
	State State

	// Specifies the number of users that are assigned to a server you specified with
	// the ServerId.
	UserCount *int32
}

// Returns properties of the user that you specify.
type ListedUser struct {

	// Provides the unique Amazon Resource Name (ARN) for the user that you want to
	// learn about.
	//
	// This member is required.
	Arn *string

	// Specifies the location that files are written to or read from an Amazon S3
	// bucket for the user you specify by their ARN.
	HomeDirectory *string

	// Specifies the type of landing directory (folder) you mapped for your users' home
	// directory. If you set it to PATH, the user will see the absolute Amazon S3
	// bucket paths as is in their file transfer protocol clients. If you set it
	// LOGICAL, you will need to provide mappings in the HomeDirectoryMappings for how
	// you want to make Amazon S3 paths visible to your users.
	HomeDirectoryType HomeDirectoryType

	// Specifies the role that is in use by this user. A role is an AWS Identity and
	// Access Management (IAM) entity that, in this case, allows a file transfer
	// protocol-enabled server to act on a user's behalf. It allows the server to
	// inherit the trust relationship that enables that user to perform file operations
	// to their Amazon S3 bucket.
	Role *string

	// Specifies the number of SSH public keys stored for the user you specified.
	SshPublicKeyCount *int32

	// Specifies the name of the user whose ARN was specified. User names are used for
	// authentication purposes.
	UserName *string
}

type PosixProfile struct {

	// This member is required.
	Gid *int64

	// This member is required.
	Uid *int64

	SecondaryGids []int64
}

// Provides information about the public Secure Shell (SSH) key that is associated
// with a user account for the specific file transfer protocol-enabled server (as
// identified by ServerId). The information returned includes the date the key was
// imported, the public key contents, and the public key ID. A user can store more
// than one SSH public key associated with their user name on a specific server.
type SshPublicKey struct {

	// Specifies the date that the public key was added to the user account.
	//
	// This member is required.
	DateImported *time.Time

	// Specifies the content of the SSH public key as specified by the PublicKeyId.
	//
	// This member is required.
	SshPublicKeyBody *string

	// Specifies the SshPublicKeyId parameter contains the identifier of the public
	// key.
	//
	// This member is required.
	SshPublicKeyId *string
}

// Creates a key-value pair for a specific resource. Tags are metadata that you can
// use to search for and group a resource for various purposes. You can apply tags
// to servers, users, and roles. A tag key can take more than one value. For
// example, to group servers for accounting purposes, you might create a tag called
// Group and assign the values Research and Accounting to that group.
type Tag struct {

	// The name assigned to the tag that you create.
	//
	// This member is required.
	Key *string

	// Contains one or more values that you assigned to the key name you create.
	//
	// This member is required.
	Value *string
}
