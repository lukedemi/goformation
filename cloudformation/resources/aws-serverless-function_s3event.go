package resources

import "github.com/lukedemi/goformation/cloudformation/policies"

// AWSServerlessFunction_S3Event AWS CloudFormation Resource (AWS::Serverless::Function.S3Event)
// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#s3
type AWSServerlessFunction_S3Event struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#s3
	Bucket string `json:"Bucket,omitempty"`

	// Events AWS CloudFormation Property
	// Required: true
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#s3
	Events *AWSServerlessFunction_Events `json:"Events,omitempty"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#s3
	Filter *AWSServerlessFunction_S3NotificationFilter `json:"Filter,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_S3Event) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.S3Event"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_S3Event) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_S3Event) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_S3Event) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_S3Event) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_S3Event) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_S3Event) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
