// Code generated by "go generate". Please don't change this file directly.

package s3

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Bucket_EventBridgeConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.EventBridgeConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-eventbridgeconfiguration.html
type Bucket_EventBridgeConfiguration struct {

	// EventBridgeEnabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-eventbridgeconfiguration.html#cfn-s3-bucket-eventbridgeconfiguration-eventbridgeenabled
	EventBridgeEnabled bool `json:"EventBridgeEnabled"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Bucket_EventBridgeConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.EventBridgeConfiguration"
}
