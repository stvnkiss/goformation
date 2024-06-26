// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_PercentageDisplayFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.PercentageDisplayFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html
type Analysis_PercentageDisplayFormatConfiguration struct {

	// DecimalPlacesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-decimalplacesconfiguration
	DecimalPlacesConfiguration *Analysis_DecimalPlacesConfiguration `json:"DecimalPlacesConfiguration,omitempty"`

	// NegativeValueConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-negativevalueconfiguration
	NegativeValueConfiguration *Analysis_NegativeValueConfiguration `json:"NegativeValueConfiguration,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Analysis_NullValueFormatConfiguration `json:"NullValueFormatConfiguration,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-prefix
	Prefix *string `json:"Prefix,omitempty"`

	// SeparatorConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-separatorconfiguration
	SeparatorConfiguration *Analysis_NumericSeparatorConfiguration `json:"SeparatorConfiguration,omitempty"`

	// Suffix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-percentagedisplayformatconfiguration.html#cfn-quicksight-analysis-percentagedisplayformatconfiguration-suffix
	Suffix *string `json:"Suffix,omitempty"`

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
func (r *Analysis_PercentageDisplayFormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PercentageDisplayFormatConfiguration"
}
