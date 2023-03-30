// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_NumberDisplayFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.NumberDisplayFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html
type Dashboard_NumberDisplayFormatConfiguration struct {

	// DecimalPlacesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-decimalplacesconfiguration
	DecimalPlacesConfiguration *Dashboard_DecimalPlacesConfiguration `json:"DecimalPlacesConfiguration,omitempty"`

	// NegativeValueConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-negativevalueconfiguration
	NegativeValueConfiguration *Dashboard_NegativeValueConfiguration `json:"NegativeValueConfiguration,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Dashboard_NullValueFormatConfiguration `json:"NullValueFormatConfiguration,omitempty"`

	// NumberScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-numberscale
	NumberScale *string `json:"NumberScale,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-prefix
	Prefix *string `json:"Prefix,omitempty"`

	// SeparatorConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-separatorconfiguration
	SeparatorConfiguration *Dashboard_NumericSeparatorConfiguration `json:"SeparatorConfiguration,omitempty"`

	// Suffix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-suffix
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
func (r *Dashboard_NumberDisplayFormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.NumberDisplayFormatConfiguration"
}