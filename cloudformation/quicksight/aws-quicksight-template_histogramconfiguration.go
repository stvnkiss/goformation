// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_HistogramConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.HistogramConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html
type Template_HistogramConfiguration struct {

	// BinOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-binoptions
	BinOptions *Template_HistogramBinOptions `json:"BinOptions,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-datalabels
	DataLabels *Template_DataLabelOptions `json:"DataLabels,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-fieldwells
	FieldWells *Template_HistogramFieldWells `json:"FieldWells,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-tooltip
	Tooltip *Template_TooltipOptions `json:"Tooltip,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-visualpalette
	VisualPalette *Template_VisualPalette `json:"VisualPalette,omitempty"`

	// XAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-xaxisdisplayoptions
	XAxisDisplayOptions *Template_AxisDisplayOptions `json:"XAxisDisplayOptions,omitempty"`

	// XAxisLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-xaxislabeloptions
	XAxisLabelOptions *Template_ChartAxisLabelOptions `json:"XAxisLabelOptions,omitempty"`

	// YAxisDisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-histogramconfiguration.html#cfn-quicksight-template-histogramconfiguration-yaxisdisplayoptions
	YAxisDisplayOptions *Template_AxisDisplayOptions `json:"YAxisDisplayOptions,omitempty"`

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
func (r *Template_HistogramConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.HistogramConfiguration"
}
