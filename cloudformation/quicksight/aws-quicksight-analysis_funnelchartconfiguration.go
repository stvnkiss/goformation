// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_FunnelChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.FunnelChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html
type Analysis_FunnelChartConfiguration struct {

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-categorylabeloptions
	CategoryLabelOptions *Analysis_ChartAxisLabelOptions `json:"CategoryLabelOptions,omitempty"`

	// DataLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-datalabeloptions
	DataLabelOptions *Analysis_FunnelChartDataLabelOptions `json:"DataLabelOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-fieldwells
	FieldWells *Analysis_FunnelChartFieldWells `json:"FieldWells,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-sortconfiguration
	SortConfiguration *Analysis_FunnelChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-tooltip
	Tooltip *Analysis_TooltipOptions `json:"Tooltip,omitempty"`

	// ValueLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-valuelabeloptions
	ValueLabelOptions *Analysis_ChartAxisLabelOptions `json:"ValueLabelOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartconfiguration.html#cfn-quicksight-analysis-funnelchartconfiguration-visualpalette
	VisualPalette *Analysis_VisualPalette `json:"VisualPalette,omitempty"`

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
func (r *Analysis_FunnelChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.FunnelChartConfiguration"
}