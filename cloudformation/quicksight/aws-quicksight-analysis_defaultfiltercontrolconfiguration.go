// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_DefaultFilterControlConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.DefaultFilterControlConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html
type Analysis_DefaultFilterControlConfiguration struct {

	// ControlOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html#cfn-quicksight-analysis-defaultfiltercontrolconfiguration-controloptions
	ControlOptions *Analysis_DefaultFilterControlOptions `json:"ControlOptions"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.html#cfn-quicksight-analysis-defaultfiltercontrolconfiguration-title
	Title string `json:"Title"`

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
func (r *Analysis_DefaultFilterControlConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DefaultFilterControlConfiguration"
}
