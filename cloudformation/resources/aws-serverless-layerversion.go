package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lukedemi/goformation/cloudformation/policies"
)

// AWSServerlessLayerVersion AWS CloudFormation Resource (AWS::Serverless::LayerVersion)
// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
type AWSServerlessLayerVersion struct {

	// CompatibleRuntimes AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	CompatibleRuntimes []string `json:"CompatibleRuntimes,omitempty"`

	// ContentUri AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	ContentUri string `json:"ContentUri,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	Description string `json:"Description,omitempty"`

	// LayerName AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	LayerName string `json:"LayerName,omitempty"`

	// LicenseInfo AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	LicenseInfo string `json:"LicenseInfo,omitempty"`

	// RetentionPolicy AWS CloudFormation Property
	// Required: false
	// See: https://github.com/lukedemi/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlesslayerversion
	RetentionPolicy string `json:"RetentionPolicy,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessLayerVersion) AWSCloudFormationType() string {
	return "AWS::Serverless::LayerVersion"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessLayerVersion) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessLayerVersion) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessLayerVersion) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessLayerVersion) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessLayerVersion) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessLayerVersion) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSServerlessLayerVersion) MarshalJSON() ([]byte, error) {
	type Properties AWSServerlessLayerVersion
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
func (r *AWSServerlessLayerVersion) UnmarshalJSON(b []byte) error {
	type Properties AWSServerlessLayerVersion
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
		*r = AWSServerlessLayerVersion(*res.Properties)
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
