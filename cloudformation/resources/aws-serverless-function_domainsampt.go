package resources

import "github.com/lukedemi/goformation/cloudformation/policies"

// AWSServerlessFunction_DomainSAMPT AWS CloudFormation Resource (AWS::Serverless::Function.DomainSAMPT)
// See: https://github.com/lukedemi/serverless-application-model/blob/master/docs/policy_templates.rst
type AWSServerlessFunction_DomainSAMPT struct {

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/docs/policy_templates.rst
	DomainName string `json:"DomainName,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_DomainSAMPT) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.DomainSAMPT"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_DomainSAMPT) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessFunction_DomainSAMPT) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_DomainSAMPT) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessFunction_DomainSAMPT) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_DomainSAMPT) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessFunction_DomainSAMPT) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
