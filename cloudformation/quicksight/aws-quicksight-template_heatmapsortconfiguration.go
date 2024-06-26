// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_HeatMapSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.HeatMapSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapsortconfiguration.html
type Template_HeatMapSortConfiguration struct {

	// HeatMapColumnItemsLimitConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapsortconfiguration.html#cfn-quicksight-template-heatmapsortconfiguration-heatmapcolumnitemslimitconfiguration
	HeatMapColumnItemsLimitConfiguration *Template_ItemsLimitConfiguration `json:"HeatMapColumnItemsLimitConfiguration,omitempty"`

	// HeatMapColumnSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapsortconfiguration.html#cfn-quicksight-template-heatmapsortconfiguration-heatmapcolumnsort
	HeatMapColumnSort []Template_FieldSortOptions `json:"HeatMapColumnSort,omitempty"`

	// HeatMapRowItemsLimitConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapsortconfiguration.html#cfn-quicksight-template-heatmapsortconfiguration-heatmaprowitemslimitconfiguration
	HeatMapRowItemsLimitConfiguration *Template_ItemsLimitConfiguration `json:"HeatMapRowItemsLimitConfiguration,omitempty"`

	// HeatMapRowSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapsortconfiguration.html#cfn-quicksight-template-heatmapsortconfiguration-heatmaprowsort
	HeatMapRowSort []Template_FieldSortOptions `json:"HeatMapRowSort,omitempty"`

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
func (r *Template_HeatMapSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.HeatMapSortConfiguration"
}
