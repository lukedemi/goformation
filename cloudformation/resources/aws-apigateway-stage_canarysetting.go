package resources

import "github.com/lukedemi/goformation/cloudformation/policies"

// AWSApiGatewayStage_CanarySetting AWS CloudFormation Resource (AWS::ApiGateway::Stage.CanarySetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html
type AWSApiGatewayStage_CanarySetting struct {

	// DeploymentId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-deploymentid
	DeploymentId string `json:"DeploymentId,omitempty"`

	// PercentTraffic AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-percenttraffic
	PercentTraffic float64 `json:"PercentTraffic,omitempty"`

	// StageVariableOverrides AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-stagevariableoverrides
	StageVariableOverrides map[string]string `json:"StageVariableOverrides,omitempty"`

	// UseStageCache AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-stage-canarysetting.html#cfn-apigateway-stage-canarysetting-usestagecache
	UseStageCache bool `json:"UseStageCache,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayStage_CanarySetting) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Stage.CanarySetting"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSApiGatewayStage_CanarySetting) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSApiGatewayStage_CanarySetting) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSApiGatewayStage_CanarySetting) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSApiGatewayStage_CanarySetting) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApiGatewayStage_CanarySetting) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApiGatewayStage_CanarySetting) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
