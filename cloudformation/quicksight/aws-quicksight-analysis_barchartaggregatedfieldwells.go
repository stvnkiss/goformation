// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_BarChartAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Analysis.BarChartAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartaggregatedfieldwells.html
type Analysis_BarChartAggregatedFieldWells struct {

	// Category AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartaggregatedfieldwells.html#cfn-quicksight-analysis-barchartaggregatedfieldwells-category
	Category []Analysis_DimensionField `json:"Category,omitempty"`

	// Colors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartaggregatedfieldwells.html#cfn-quicksight-analysis-barchartaggregatedfieldwells-colors
	Colors []Analysis_DimensionField `json:"Colors,omitempty"`

	// SmallMultiples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartaggregatedfieldwells.html#cfn-quicksight-analysis-barchartaggregatedfieldwells-smallmultiples
	SmallMultiples []Analysis_DimensionField `json:"SmallMultiples,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartaggregatedfieldwells.html#cfn-quicksight-analysis-barchartaggregatedfieldwells-values
	Values []Analysis_MeasureField `json:"Values,omitempty"`

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
func (r *Analysis_BarChartAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.BarChartAggregatedFieldWells"
}
