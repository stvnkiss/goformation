// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Topic_TopicCategoryFilter AWS CloudFormation Resource (AWS::QuickSight::Topic.TopicCategoryFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html
type Topic_TopicCategoryFilter struct {

	// CategoryFilterFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-categoryfilterfunction
	CategoryFilterFunction *string `json:"CategoryFilterFunction,omitempty"`

	// CategoryFilterType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-categoryfiltertype
	CategoryFilterType *string `json:"CategoryFilterType,omitempty"`

	// Constant AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-constant
	Constant *Topic_TopicCategoryFilterConstant `json:"Constant,omitempty"`

	// Inverse AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccategoryfilter.html#cfn-quicksight-topic-topiccategoryfilter-inverse
	Inverse *bool `json:"Inverse,omitempty"`

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
func (r *Topic_TopicCategoryFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.TopicCategoryFilter"
}