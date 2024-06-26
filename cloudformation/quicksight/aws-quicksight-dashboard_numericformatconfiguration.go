// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_NumericFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.NumericFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html
type Dashboard_NumericFormatConfiguration struct {

	// CurrencyDisplayFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-currencydisplayformatconfiguration
	CurrencyDisplayFormatConfiguration *Dashboard_CurrencyDisplayFormatConfiguration `json:"CurrencyDisplayFormatConfiguration,omitempty"`

	// NumberDisplayFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-numberdisplayformatconfiguration
	NumberDisplayFormatConfiguration *Dashboard_NumberDisplayFormatConfiguration `json:"NumberDisplayFormatConfiguration,omitempty"`

	// PercentageDisplayFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericformatconfiguration.html#cfn-quicksight-dashboard-numericformatconfiguration-percentagedisplayformatconfiguration
	PercentageDisplayFormatConfiguration *Dashboard_PercentageDisplayFormatConfiguration `json:"PercentageDisplayFormatConfiguration,omitempty"`

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
func (r *Dashboard_NumericFormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.NumericFormatConfiguration"
}
