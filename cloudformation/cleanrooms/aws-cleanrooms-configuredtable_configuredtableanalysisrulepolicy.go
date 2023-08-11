// Code generated by "go generate". Please don't change this file directly.

package cleanrooms

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ConfiguredTable_ConfiguredTableAnalysisRulePolicy AWS CloudFormation Resource (AWS::CleanRooms::ConfiguredTable.ConfiguredTableAnalysisRulePolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html
type ConfiguredTable_ConfiguredTableAnalysisRulePolicy struct {

	// V1 AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.html#cfn-cleanrooms-configuredtable-configuredtableanalysisrulepolicy-v1
	V1 *ConfiguredTable_ConfiguredTableAnalysisRulePolicyV1 `json:"V1"`

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
func (r *ConfiguredTable_ConfiguredTableAnalysisRulePolicy) AWSCloudFormationType() string {
	return "AWS::CleanRooms::ConfiguredTable.ConfiguredTableAnalysisRulePolicy"
}