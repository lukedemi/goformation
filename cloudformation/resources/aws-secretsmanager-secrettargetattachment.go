package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lukedemi/goformation/cloudformation/policies"
)

// AWSSecretsManagerSecretTargetAttachment AWS CloudFormation Resource (AWS::SecretsManager::SecretTargetAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
type AWSSecretsManagerSecretTargetAttachment struct {

	// SecretId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-secretid
	SecretId string `json:"SecretId,omitempty"`

	// TargetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-targetid
	TargetId string `json:"TargetId,omitempty"`

	// TargetType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html#cfn-secretsmanager-secrettargetattachment-targettype
	TargetType string `json:"TargetType,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSecretsManagerSecretTargetAttachment) AWSCloudFormationType() string {
	return "AWS::SecretsManager::SecretTargetAttachment"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSecretsManagerSecretTargetAttachment) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSecretsManagerSecretTargetAttachment) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSecretsManagerSecretTargetAttachment) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSecretsManagerSecretTargetAttachment) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSecretsManagerSecretTargetAttachment) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSecretsManagerSecretTargetAttachment) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSSecretsManagerSecretTargetAttachment) MarshalJSON() ([]byte, error) {
	type Properties AWSSecretsManagerSecretTargetAttachment
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string                `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{}  `json:"Metadata,omitempty"`
		DeletionPolicy policies.DeletionPolicy `json:"DeletionPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r._dependsOn,
		Metadata:       r._metadata,
		DeletionPolicy: r._deletionPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSSecretsManagerSecretTargetAttachment) UnmarshalJSON(b []byte) error {
	type Properties AWSSecretsManagerSecretTargetAttachment
	res := &struct {
		Type           string
		Properties     *Properties
		DependsOn      []string
		Metadata       map[string]interface{}
		DeletionPolicy string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSSecretsManagerSecretTargetAttachment(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r._deletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	return nil
}
