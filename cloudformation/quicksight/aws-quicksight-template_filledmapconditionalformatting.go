// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_FilledMapConditionalFormatting AWS CloudFormation Resource (AWS::QuickSight::Template.FilledMapConditionalFormatting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filledmapconditionalformatting.html
type Template_FilledMapConditionalFormatting struct {

	// ConditionalFormattingOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filledmapconditionalformatting.html#cfn-quicksight-template-filledmapconditionalformatting-conditionalformattingoptions
	ConditionalFormattingOptions []Template_FilledMapConditionalFormattingOption `json:"ConditionalFormattingOptions"`

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
func (r *Template_FilledMapConditionalFormatting) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FilledMapConditionalFormatting"
}
