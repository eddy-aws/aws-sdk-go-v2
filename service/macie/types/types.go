// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The classification type that Amazon Macie Classic applies to the associated S3
// resources.
type ClassificationType struct {

	// A continuous classification of the objects that are added to a specified S3
	// bucket. Amazon Macie Classic begins performing continuous classification after a
	// bucket is successfully associated with Macie Classic.
	//
	// This member is required.
	Continuous S3ContinuousClassificationType

	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	//
	// This member is required.
	OneTime S3OneTimeClassificationType
}

// The classification type that Amazon Macie Classic applies to the associated S3
// resources. At least one of the classification types (oneTime or continuous) must
// be specified.
type ClassificationTypeUpdate struct {

	// A continuous classification of the objects that are added to a specified S3
	// bucket. Amazon Macie Classic begins performing continuous classification after a
	// bucket is successfully associated with Macie Classic.
	Continuous S3ContinuousClassificationType

	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	OneTime S3OneTimeClassificationType
}

// Includes details about the failed S3 resources.
type FailedS3Resource struct {

	// The status code of a failed item.
	ErrorCode *string

	// The error message of a failed item.
	ErrorMessage *string

	// The failed S3 resources.
	FailedItem *S3Resource
}

// Contains information about the Amazon Macie Classic member account.
type MemberAccount struct {

	// The AWS account ID of the Amazon Macie Classic member account.
	AccountId *string
}

// Contains information about the S3 resource. This data type is used as a request
// parameter in the DisassociateS3Resources action and can be used as a response
// parameter in the AssociateS3Resources and UpdateS3Resources actions.
type S3Resource struct {

	// The name of the S3 bucket.
	//
	// This member is required.
	BucketName *string

	// The prefix of the S3 bucket.
	Prefix *string
}

// The S3 resources that you want to associate with Amazon Macie Classic for
// monitoring and data classification. This data type is used as a request
// parameter in the AssociateS3Resources action and a response parameter in the
// ListS3Resources action.
type S3ResourceClassification struct {

	// The name of the S3 bucket that you want to associate with Amazon Macie Classic.
	//
	// This member is required.
	BucketName *string

	// The classification type that you want to specify for the resource associated
	// with Amazon Macie Classic.
	//
	// This member is required.
	ClassificationType *ClassificationType

	// The prefix of the S3 bucket that you want to associate with Amazon Macie
	// Classic.
	Prefix *string
}

// The S3 resources whose classification types you want to update. This data type
// is used as a request parameter in the UpdateS3Resources action.
type S3ResourceClassificationUpdate struct {

	// The name of the S3 bucket whose classification types you want to update.
	//
	// This member is required.
	BucketName *string

	// The classification type that you want to update for the resource associated with
	// Amazon Macie Classic.
	//
	// This member is required.
	ClassificationTypeUpdate *ClassificationTypeUpdate

	// The prefix of the S3 bucket whose classification types you want to update.
	Prefix *string
}
