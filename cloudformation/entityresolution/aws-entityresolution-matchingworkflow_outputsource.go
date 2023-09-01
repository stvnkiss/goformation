// Code generated by "go generate". Please don't change this file directly.

package entityresolution

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// MatchingWorkflow_OutputSource AWS CloudFormation Resource (AWS::EntityResolution::MatchingWorkflow.OutputSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html
type MatchingWorkflow_OutputSource struct {

	// ApplyNormalization AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-applynormalization
	ApplyNormalization *bool `json:"ApplyNormalization,omitempty"`

	// KMSArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-kmsarn
	KMSArn *string `json:"KMSArn,omitempty"`

	// Output AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-output
	Output []MatchingWorkflow_OutputAttribute `json:"Output"`

	// OutputS3Path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-outputs3path
	OutputS3Path string `json:"OutputS3Path"`

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
func (r *MatchingWorkflow_OutputSource) AWSCloudFormationType() string {
	return "AWS::EntityResolution::MatchingWorkflow.OutputSource"
}