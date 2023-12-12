// Code generated by "go generate". Please don't change this file directly.

package osis

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Pipeline_EncryptionAtRestOptions AWS CloudFormation Resource (AWS::OSIS::Pipeline.EncryptionAtRestOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-encryptionatrestoptions.html
type Pipeline_EncryptionAtRestOptions struct {

	// KmsKeyArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-encryptionatrestoptions.html#cfn-osis-pipeline-encryptionatrestoptions-kmskeyarn
	KmsKeyArn string `json:"KmsKeyArn"`

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
func (r *Pipeline_EncryptionAtRestOptions) AWSCloudFormationType() string {
	return "AWS::OSIS::Pipeline.EncryptionAtRestOptions"
}
