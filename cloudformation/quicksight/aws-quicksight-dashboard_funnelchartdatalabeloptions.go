// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_FunnelChartDataLabelOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FunnelChartDataLabelOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html
type Dashboard_FunnelChartDataLabelOptions struct {

	// CategoryLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-categorylabelvisibility
	CategoryLabelVisibility *string `json:"CategoryLabelVisibility,omitempty"`

	// LabelColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-labelcolor
	LabelColor *string `json:"LabelColor,omitempty"`

	// LabelFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-labelfontconfiguration
	LabelFontConfiguration *Dashboard_FontConfiguration `json:"LabelFontConfiguration,omitempty"`

	// MeasureDataLabelStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-measuredatalabelstyle
	MeasureDataLabelStyle *string `json:"MeasureDataLabelStyle,omitempty"`

	// MeasureLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-measurelabelvisibility
	MeasureLabelVisibility *string `json:"MeasureLabelVisibility,omitempty"`

	// Position AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-position
	Position *string `json:"Position,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-funnelchartdatalabeloptions.html#cfn-quicksight-dashboard-funnelchartdatalabeloptions-visibility
	Visibility *string `json:"Visibility,omitempty"`

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
func (r *Dashboard_FunnelChartDataLabelOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FunnelChartDataLabelOptions"
}
