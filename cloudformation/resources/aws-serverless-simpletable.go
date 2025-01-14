package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lukedemi/goformation/cloudformation/policies"
)

// AWSServerlessSimpleTable AWS CloudFormation Resource (AWS::Serverless::SimpleTable)
// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesssimpletable
type AWSServerlessSimpleTable struct {

	// PrimaryKey AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#primary-key-object
	PrimaryKey *AWSServerlessSimpleTable_PrimaryKey `json:"PrimaryKey,omitempty"`

	// ProvisionedThroughput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
	ProvisionedThroughput *AWSServerlessSimpleTable_ProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`

	// SSESpecification AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesssimpletable
	SSESpecification *AWSServerlessSimpleTable_SSESpecification `json:"SSESpecification,omitempty"`

	// TableName AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesssimpletable
	TableName string `json:"TableName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesssimpletable
	Tags map[string]string `json:"Tags,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessSimpleTable) AWSCloudFormationType() string {
	return "AWS::Serverless::SimpleTable"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessSimpleTable) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessSimpleTable) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessSimpleTable) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessSimpleTable) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessSimpleTable) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessSimpleTable) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSServerlessSimpleTable) MarshalJSON() ([]byte, error) {
	type Properties AWSServerlessSimpleTable
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
func (r *AWSServerlessSimpleTable) UnmarshalJSON(b []byte) error {
	type Properties AWSServerlessSimpleTable
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
		*r = AWSServerlessSimpleTable(*res.Properties)
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
